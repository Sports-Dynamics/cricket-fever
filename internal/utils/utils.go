package utils

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sports-dynamics/cricket-fever/internal/validation"
)

func GetUUIDFromRequest(r *http.Request, urlParam string) (string, error) {
	paramUUID := chi.URLParam(r, urlParam)
	if _, err := uuid.Parse(paramUUID); err != nil {
		return "", validation.NewAppError(validation.ErrCodeInvalidInput, "invalid uuid, should be a valid uuid", urlParam)
	}
	return paramUUID, nil
}

func FetchLocalMachineIpAddress() *string {
	var result string
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return &result
	}

	// Iterate through each interface
	for _, iface := range interfaces {
		// Get the addresses for the current interface
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Iterate through each address
		for _, addr := range addrs {
			// Check if the address is an IP address
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				// Check if it's an IPv4 address (you can modify this condition if you need IPv6)
				if ipnet.IP.To4() != nil {
					fmt.Println("Local Machine IP Address:", ipnet.IP)
					result = ipnet.IP.String()
				}
			}
		}
	}
	return &result
}
