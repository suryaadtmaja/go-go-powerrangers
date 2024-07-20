package readz

import (
	"fmt"
	"os"
)

func CreateFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Can't create a file %v\n", err)
		return
	}

	defer file.Close()

	file.WriteString("test")
}
