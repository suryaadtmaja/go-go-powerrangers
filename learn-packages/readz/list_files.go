package readz

import (
	"fmt"
	"os"
)

func ListFiles() {
	dir, err := os.Open(".")
	if err != nil {
		fmt.Printf("Cannot open the directory: %v\n", err)
		return
	}
	defer dir.Close()

	fileInfos, err := dir.ReadDir(-1)

	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
