// Package isp resolves the internet service provider of a client IP address.
package isp

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

// Detector resolves the internet service provider for a client IP address.
type Detector interface {
	// GetISP returns the ISP string for the given IP, or an empty string when
	// the IP is local, private or detection fails.
	GetISP(ip string) string
}

// IPInfoDetector resolves ISPs using the ipinfo.io JSON API.
type IPInfoDetector struct {
	client  *http.Client
	baseURL string
}

// NewIPInfoDetector creates an IPInfoDetector using the provided HTTP client.
// When client is nil, a default client with a short timeout is used.
func NewIPInfoDetector(client *http.Client) *IPInfoDetector {
	if client == nil {
		client = &http.Client{Timeout: 3 * time.Second}
	}
	return &IPInfoDetector{client: client, baseURL: "https://ipinfo.io"}
}

type ipInfoResponse struct {
	Org string `json:"org"`
}

// GetISP queries ipinfo.io for the given IP and returns its org value
// (e.g. "AS2501 Iran Cell"). Errors and non-public IPs yield an empty string.
func (d *IPInfoDetector) GetISP(ip string) string {
	parsed := net.ParseIP(strings.TrimSpace(ip))
	if parsed == nil || parsed.IsPrivate() || parsed.IsLoopback() || parsed.IsUnspecified() || parsed.IsLinkLocalUnicast() {
		return ""
	}

	url := fmt.Sprintf("%s/%s/json", d.baseURL, parsed.String())
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ""
	}
	req.Header.Set("Accept", "application/json")

	resp, err := d.client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ""
	}

	var data ipInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return ""
	}
	return data.Org
}
