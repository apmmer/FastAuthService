package handlers_utils

import (
	"auth_service_api/internal/schemas"
	"net/http"
)

// GetDeviceInfo extracts the device information from the HTTP request.
func GetDeviceInfo(r *http.Request) schemas.DeviceInfo {
	ipAddress := getIPAddress(r)
	userAgent := r.UserAgent()

	return schemas.DeviceInfo{
		IPAddress: ipAddress,
		UserAgent: userAgent,
	}
}

// getIPAddress extracts the client's IP address from the HTTP request.
// It first checks the X-Forwarded-For header which can be set by proxy servers
// to indicate the original client's IP address. If the X-Forwarded-For header is empty,
// it falls back to using the RemoteAddr field of the http.Request which gives the network address
// that sent the request.
func getIPAddress(r *http.Request) string {
	ipAddress := r.Header.Get("X-Forwarded-For")

	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}

	return ipAddress
}
