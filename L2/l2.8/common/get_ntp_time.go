package common

import (
	"time"

	"github.com/beevik/ntp"
)

// GetNtpTime returns current time according to ntp server
func GetNtpTime() (time.Time, error) {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, err
	}
	return ntpTime, nil
}
