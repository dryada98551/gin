package lib

import (
	"net/http"
	"github.com/ua-parser/uap-go/uaparser"
	"fmt"
)

// getClientIP tries to determine the client's IP address from the HTTP request.
func GetClientIP(r *http.Request) string {
	// Check if the X-Forwarded-For header is set.
	if xForwardedFor := r.Header.Get("X-Forwarded-For"); xForwardedFor != "" {
		return xForwardedFor
	}

	// If X-Forwarded-For is not set, check the X-Real-IP header.
	if xRealIP := r.Header.Get("X-Real-IP"); xRealIP != "" {
		return xRealIP
	}

	// If neither header is set, fall back to using the remote address from the request.
	// Note: This might not be the actual client IP if the request went through a proxy.
	return r.RemoteAddr
}

func GetDevice(r *http.Request) (string, error) {
	// Check if the X-Forwarded-For header is set.
	if userAgent := r.Header.Get("User-Agent"); userAgent != "" {
		parser := uaparser.NewFromSaved()
		client := parser.Parse(userAgent)
		Device := client.Device.Family
		return Device, nil
	} else {
		return "", fmt.Errorf("no Device")
	}
}