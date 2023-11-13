package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var (
	options     = make(map[string]any, 0)
	visitedURLs = make(map[string]bool)
)

func main() {
	options["mirror"] = flag.Bool("mirror", false, "mirror entire site")
	options["l"] = flag.Int("l", 5, "Depth of recursive downloading site, negative is no limit")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("use with at least one url!")
		os.Exit(1)
	}

	for _, url_s := range args {
		url, err := url.Parse(url_s)
		if err != nil {
			fmt.Printf("Error parsing URL: %s, err: %s\n", url_s, err)
			continue
		}
		if err := getPage(url); err != nil {
			fmt.Println(err)
		}
	}
}

func getPage(URL *url.URL) error {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Разрешаем редиректы
			return nil
		},
	}
	resp, err := client.Get(URL.String())
	if err != nil {
		return fmt.Errorf("\nget: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP \nerror: %s", resp.Status)
	}

	if *options["mirror"].(*bool) {
		dirName := URL.Host
		getPageR(client, URL, URL.Host, 0)
		fmt.Println("\nDownload complete. Saved as", dirName)
	} else {
		dirName := strings.Trim(URL.Host+URL.Path, "/")
		err := os.Mkdir(dirName, 0777)
		if err != nil {
			return fmt.Errorf("\nmkdir: %s", err)
		}
		fileName := dirName + "/index.html"
		file, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("\ncreate: %s", err)
		}
		defer file.Close()
		var content []byte
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return fmt.Errorf("\ncopy: %s", err)
		}
		file.Write(content)
		fmt.Println("\nDownload complete. Saved as", fileName)
	}

	return nil
}
func getPageR(client *http.Client, URL *url.URL, rootDir string, depth int) error {
	if depth == *options["l"].(*int) {
		return nil
	}
	if visitedURLs[URL.String()] {
		return nil
	}
	visitedURLs[URL.String()] = true

	filename := strings.Trim(rootDir+"/"+URL.Host+URL.Path, "/")
	// filename := rootDir + "/" + URL.Host + URL.Path
	file, err := createFile(filename)
	if err != nil {
		return fmt.Errorf("\ncreate: %s", err)
	}
	defer file.Close()

	resp, err := client.Get(URL.String())
	if err != nil {
		return fmt.Errorf("\nget: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("\nHTTP error: %s", resp.Status)
	}

	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)

	links, err := extractLinks(tee)
	if err != nil {
		return fmt.Errorf("\nextractLinks: %s", err)
	}

	_, err = io.Copy(file, &buf)
	if err != nil {
		return fmt.Errorf("\ncopy: %s", err)
	}
	for _, link := range links {
		linkURL, err := url.Parse(link)
		if err != nil {
			fmt.Println("Error parsing link:", err)
			continue
		}
		err = getPageR(client, linkURL, rootDir, depth+1)
		if err != nil {
			fmt.Println("Error downloading", linkURL, ":", err)
		}
	}

	return nil
}
func createFile(fileName string) (file *os.File, err error) {
	path := strings.Split(fileName, "/")
	var i int
	for i = 0; i < len(path)-1; i++ {
		os.Mkdir(path[i], 0777)
		os.Chdir(path[i])
	}
	if i == 0 {
		file, err = os.Create(fileName)
	} else {
		file, err = os.Create(path[i])
	}

	if err != nil {
		fmt.Printf("could not create file: %s", err)
		file, err = os.Create("index.html")
	}

	return
}

func extractLinks(r io.Reader) ([]string, error) {
	var links []string
	tokenizer := html.NewTokenizer(r)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			if tokenizer.Err() == io.EOF {
				return links, nil
			}
			return nil, tokenizer.Err()

		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" || attr.Key == "src" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}
