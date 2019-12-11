package main

import (
	"github.com/beevik/ntp"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func errorWrapper(e error, message string) {
	if e != nil {
		log.Error(e)
		os.Exit(1)
	}
}

func GetNtpTime(host string) (time.Time, error) {
	t, err := ntp.Time(host)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func main() {
	times, err := GetNtpTime("0.beevik-ntp.pool.ntp.org")
	errorWrapper(err, "")
	log.Info(times)
}
