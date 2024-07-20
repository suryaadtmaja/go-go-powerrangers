package readz

import (
	"fmt"
	"os"
	"path/filepath"
)

func WalkingOnPath() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, "Path")
		return nil
	})

	if err != nil {
		fmt.Printf("error walking down the path %v\n", err)
		return
	}
}
