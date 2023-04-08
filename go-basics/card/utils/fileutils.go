package utils

import "os"

func CreateDirectory(dirName string) error {
	return os.MkdirAll(dirName, 0755)
}

func SaveToFile(fileName string, bytes []byte) error {
	return os.WriteFile(fileName, bytes, 0644)
}
