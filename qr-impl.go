package qr

import (
	"github.com/skip2/go-qrcode"
)

// Client manage all QR action
type Client struct{}

// Default function will generate an QR based on content and size
func (c *Client) Default(content string, size int) ([]byte, error) {
	return qrcode.Encode(content, qrcode.Medium, size)
}

// Export function will export QR to a file based on content, path and size. Path example: "C:\demo.png"
func (c *Client) Export(content, path string, size int) error {
	return qrcode.WriteFile(content, qrcode.Medium, size, path)
}
