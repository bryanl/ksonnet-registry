package store

// Store manages files
type Store interface {
	Write(filename string, data []byte) error
	Read(filename string) ([]byte, error)
	Delete(filename string) error
}
