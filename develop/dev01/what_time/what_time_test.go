package what_time

import "testing"

func IsErrorOccurred(out error, want bool) bool {
	if out != nil && want == false {
		return true
	}

	if out == nil && want == true {
		return true
	}

	return false
}

type getCurrentTest struct {
	hostname string
	wantNil  bool
}

var getCurrentTests = []getCurrentTest{
	{"0.ru.pool.ntp.org", true},
	{"1.ru.pool.ntp.org", true},
	{"2.ru.pool.ntp.org", true},
	{"3.ru.pool.ntp.org", true},
	{"ntp1.stratum2.ru", true},
	{"ntp2.stratum2.ru", true},
	{"ntp3.stratum2.ru", true},
	{"example.com", false},
}

func TestGetCurrent(t *testing.T) {
	dt := New("")
	for _, test := range getCurrentTests {
		dt.ChangeHost(test.hostname)
		if _, err := dt.GetTime(); !IsErrorOccurred(err, test.wantNil) {
			t.Errorf("Expected value to be error: %v; Actual value of error: %v; Hostname: %v;",
				!test.wantNil, err, test.hostname)
		}
	}
}
