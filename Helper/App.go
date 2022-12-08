package Helper

import (
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

func AppPath() string {
	dir, err := os.Getwd()
	if err != nil {
		println(err)
		return ""
	}
	return dir
}

func RandomText(prefix, suffix string) string {
	return prefix + uuid.New().String() + suffix
}

func GetFileExtension(Filename string) string {
	return filepath.Ext(Filename)
}
