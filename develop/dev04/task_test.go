package main

import "testing"

type findAnagramsTest struct {
	data     []string
	expected *map[string][]string
}

var findAnagramsTests = []findAnagramsTest{
	{
		[]string{"ПЯТАК", "пЯтКа", "Тяпка", "Листок", "СЛИТОК", "столиК", "каждый", "охотник", "Желает"},
		&map[string][]string{
			"пятак":  {"пятак", "пятка", "тяпка"},
			"листок": {"листок", "слиток", "столик"},
		},
	},
	{
		[]string{"Воля", "вяЛо", "воля", "светобоязнь", "обезьянство", "каа", "аак", "ака", "каа", "каа", "каа"},
		&map[string][]string{
			"воля":        {"воля", "вяло"},
			"каа":         {"аак", "ака", "каа"},
			"светобоязнь": {"обезьянство", "светобоязнь"},
		},
	},
}

func TestFindAnagrams(t *testing.T) {
	for _, test := range findAnagramsTests {
		result := findAnagrams(&test.data)
		if len(*result) != len(*test.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", *test.expected, *result)
		}
		for k, v := range *result {
			expected := *test.expected
			if !Equal(v, expected[k]) {
				t.Errorf("Incorrect result. Expect: %v, Got %v\n", expected[k], v)
			}
		}
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
