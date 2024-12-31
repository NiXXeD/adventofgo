package util

import "os"

func LoadFile(path string) string {
	data, err := os.ReadFile(path)
	Check(err)
	return string(data)
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
