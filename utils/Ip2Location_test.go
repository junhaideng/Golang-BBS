package utils

import "testing"

func TestGetIpLocation(t *testing.T) {
	ip := "202.120.10.101"
	t.Log(GetLocationByIp(ip))
}
