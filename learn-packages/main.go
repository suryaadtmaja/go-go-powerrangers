package main

import (
	"fmt"
	"learn-packages/math"
	"learn-packages/readz"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
	readz.ReadFile("file.txt")
	readz.ReadFileWithShorter("file.txt")
	readz.CreateFile("Randomhehe")
	readz.ListFiles()
	readz.WalkingOnPath()
}
