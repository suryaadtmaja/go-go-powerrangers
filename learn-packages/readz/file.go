package readz

import (
	"fmt"
	"os"
)

func ReadFile(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()

	if err != nil {
		fmt.Printf("Error reading file stat: %v\n", err)
		return
		// In Go, the return statement after checking for an error is used to exit the current function
		// immediately if an error occurs. This is important for several reasons:
		// - Prevent Further Execution:
		// - Simplify Error Handling
		// - Resource Management
	}

	bs := make([]byte, stat.Size())

	_, err = file.Read(bs)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	str := string(bs)
	fmt.Println(str)
}
