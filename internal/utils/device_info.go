package utils

import (
	"AuthService/internal/schemas"
	"net/http"
)

func GetDeviceInfo(r *http.Request) schemas.DeviceInfo {
	ipAddress := getIPAddress(r)
	userAgent := r.UserAgent()

	return schemas.DeviceInfo{
		IPAddress: ipAddress,
		UserAgent: userAgent,
	}
}

func getIPAddress(r *http.Request) string {
	ipAddress := r.Header.Get("X-Forwarded-For")

	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}

	return ipAddress
}
