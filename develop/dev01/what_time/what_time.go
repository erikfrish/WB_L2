package what_time

import (
	"time"

	"github.com/beevik/ntp"
)

type NTP_server struct {
	host string
}

func New(hostname string) NTP_server {
	return NTP_server{
		host: hostname,
	}
}
func (s *NTP_server) ChangeHost(hostname string) {
	s.host = hostname
}

func (s *NTP_server) GetTime() (time.Time, error) {
	response, err := ntp.Query(s.host)
	if err != nil {
		return time.Time{}, err
	}
	err = response.Validate()
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().Add(response.ClockOffset), nil
}

func (s *NTP_server) GetTimeF(format string) (string, error) {
	response, err := ntp.Query(s.host)
	if err != nil {
		return "", err
	}
	err = response.Validate()
	if err != nil {
		return "", err
	}
	return time.Now().Add(response.ClockOffset).Format(format), nil
}
