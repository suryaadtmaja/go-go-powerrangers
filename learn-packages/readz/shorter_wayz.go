package readz

import (
	"fmt"
	"os"
)

func ReadFileWithShorter(n string) {
	bs, err := os.ReadFile(n)

	if err != nil {
		ReturnError("Error opening file")
		// fmt.Printf("Error opening file: %v\n", err)
		return
	}

	str := string(bs)

	fmt.Println(str)
}
