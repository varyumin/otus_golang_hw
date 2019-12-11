package main

import "testing"

func TestGetNtpTime(t *testing.T) {
	ntpServer := "0.beevik-ntp.pool.ntp.org"
	if len(ntpServer) < 1 {
		t.Fatalf("Name NTP server is empty")
	}
	_, e := GetNtpTime(ntpServer)
	if e != nil {
		t.Fatalf("bad return value for ntp server %s", ntpServer)
	}

}
