package sorting

import "testing"

type defaultSortTest struct {
	lines    []string
	expected []string
}

var defaultSortTests = []defaultSortTest{
	{
		[]string{"garlic", "world", "dell", "zebra", "Borgia", "abc"},
		[]string{"Borgia", "abc", "dell", "garlic", "world", "zebra"},
	},
}

type numSortTest struct {
	lines    []string
	expected []string
}

var numSortTests = []numSortTest{
	{
		[]string{"0", "22", "1", "125", "30", "350"},
		[]string{"0", "1", "22", "30", "125", "350"},
	},
}

type monthsSortTest struct {
	lines    []string
	expected []string
}

var monthsSortTests = []monthsSortTest{
	{
		[]string{"December", "november", "July", "february", "January", "May"},
		[]string{"January", "february", "May", "July", "november", "December"},
	},
}

type suffixSortTest struct {
	lines    []string
	expected []string
}

var suffixSortTests = []suffixSortTest{
	{
		[]string{"103", "100000B", "0.5P", "1.5T", "2K", "100M"},
		[]string{"103", "2K", "100000B", "100M", "1.5T", "0.5P"},
	},
}

func TestDefaultSort(t *testing.T) {
	for _, test := range defaultSortTests {
		units := groupBySortKeys(test.lines, 0)
		res := DefaultSort(units)
		for i := range res {
			if res[i] != test.expected[i] {
				t.Errorf("Output %v; Expected %v;", res, test.expected)
			}
		}
	}
}

func TestNumericSort(t *testing.T) {
	for _, test := range numSortTests {
		units := groupBySortKeys(test.lines, 0)
		res := NumSort(units)
		for i := range res {
			if res[i] != test.expected[i] {
				t.Errorf("Output %v; Expected %v;", res, test.expected)
			}
		}
	}
}

func TestMonthSort(t *testing.T) {
	for _, test := range monthsSortTests {
		units := groupBySortKeys(test.lines, 0)
		res := MonthSort(units)
		for i := range res {
			if res[i] != test.expected[i] {
				t.Errorf("Output %v; Expected %v;", res, test.expected)
			}
		}
	}
}

func TestHumanSort(t *testing.T) {
	for _, test := range suffixSortTests {
		units := groupBySortKeys(test.lines, 0)
		res := SuffixSort(units)
		for i := range res {
			if res[i] != test.expected[i] {
				t.Errorf("Output %v; Expected %v;", res, test.expected)
			}
		}
	}
}
