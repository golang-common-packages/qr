package qr

// IQR interface for qr package
type IQR interface {
	Default(content string, size int) ([]byte, error)
	Export(content, path string, size int) error
}
