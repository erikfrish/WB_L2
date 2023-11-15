package main

/*
=== Утилита telnet ===

Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Требования:
1. 	Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
	После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
2. 	Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
3. 	При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
	Если сокет закрывается со стороны сервера, программа должна также завершаться.
	При подключении к несуществующему сервер, программа должна завершаться через timeout
*/
import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"unicode"
)

type conn_cfg struct {
	host, port string
	timeout    time.Duration
}

func (c *conn_cfg) tcpConnect() (net.Conn, error) {
	var err error
	var conn net.Conn
	timeout := time.After(c.timeout)
	connected := make(chan bool)
	err_ch := make(chan bool)
	go func() {
		conn, err = net.Dial("tcp", c.host+":"+c.port)
		if err == nil {
			connected <- true
		} else {
			err_ch <- true
		}
	}()
	select {
	case <-timeout:
		if err == nil {
			return nil, fmt.Errorf("connection timed out")
		} else {
			return nil, fmt.Errorf("connection timed out: %s", err)
		}
	case <-err_ch:
		return nil, fmt.Errorf("connection timed out: %s", err)
	case <-connected:
		return conn, nil
	}
}

func main() {
	timeout_s := flag.String("timeout", "10s", "timeout duration: 5s, 100ms, 1h, etc.")

	flag.Parse()
	host := flag.Arg(0)
	for len(host) == 0 {
		fmt.Println("Please type hostname (or ip):")
		fmt.Scan(&host)
	}
	port := flag.Arg(1)
	for len(port) == 0 {
		fmt.Println("Please type port number:")
		fmt.Scan(&host)
	}
	timeout, err := ConvertStringToTimeDuration(*timeout_s)
	if err != nil {
		fmt.Println("converting timeout failed: ", err)
		fmt.Println("it is set on 10s by default")
	}

	c_cfg := conn_cfg{
		host:    host,
		port:    port,
		timeout: timeout,
	}

	go startTestServer4242()

	tcp_conn, err := c_cfg.tcpConnect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer tcp_conn.Close()
	fmt.Printf("connected to %s:%s\n", c_cfg.host, c_cfg.port)

	go func() {
		r := bufio.NewScanner(tcp_conn)
		for r.Scan() {
			fmt.Println(r.Text())
		}
	}()

	sig_chan := make(chan os.Signal, 1)
	signal.Notify(sig_chan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for sig := range sig_chan {
			fmt.Printf("\nReceived signal: %v, exiting...\n", sig)
			tcp_conn.Close()
			os.Exit(0)
		}
	}()

	s := bufio.NewScanner(os.Stdin)
	for {
		if s.Scan() {
			text := s.Text()
			fmt.Println(text)
			_, err := fmt.Fprintln(tcp_conn, text)
			// При закрытии сокета сервером выходим
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("you typed ctrl+d, exiting...")
			tcp_conn.Close()
			return
		}
	}
}

func ConvertStringToTimeDuration(str string) (time.Duration, error) {
	digits := make([]rune, 0)
	suffix := make([]rune, 0)
	var result time.Duration

	for _, char := range str {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
		} else {
			suffix = append(suffix, char)
		}
	}

	num, err := strconv.Atoi(string(digits))
	if err != nil {
		return time.Duration(10 * time.Second), err
	}
	switch string(suffix) {
	case "h":
		result = time.Hour * time.Duration(num)
	case "m":
		result = time.Minute * time.Duration(num)
	case "s":
		result = time.Second * time.Duration(num)
	case "ms":
		result = time.Millisecond * time.Duration(num)
	case "ns":
		result = time.Nanosecond * time.Duration(num)
	default:
		result = time.Second * time.Duration(num)
	}

	return result, nil
}

func startTestServer4242() {
	ln, err := net.Listen("tcp", "127.0.0.1:4242")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		fmt.Println("connected to server")
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("handling connection", conn.RemoteAddr())
	r := bufio.NewScanner(conn)
	for r.Scan() {
		_, err := fmt.Fprintln(conn, "echo: ", r.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
