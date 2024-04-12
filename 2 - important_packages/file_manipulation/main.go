package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("CREATING FILE")
	createFile()

	fmt.Println("\nREADING FULL FILE")
	readFile()

	fmt.Println("\nREADING FILE BY BLOCKS")
	readFileByBlocks()

	removeFile()
}

func createFile() {
	file, error := os.Create("file.txt")
	if error != nil {
		panic(error)
	}

	// size, error := file.WriteString("Hello World!")
	size, error := file.Write([]byte("Writing data on file"))
	if error != nil {
		panic(error)
	}

	fmt.Printf("File created successfuly! It's size is %d bytes\n", size)

	file.Close()
}

func readFile() {
	file, error := os.ReadFile("file.txt")
	if error != nil {
		panic(error)
	}

	fmt.Println(string(file))
}

func readFileByBlocks() {
	file, error := os.Open("file.txt")
	if error != nil {
		panic(error)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 3)

	for {
		n, error := reader.Read(buffer)
		if error != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}

func removeFile() {
	error := os.Remove("file.txt")
	if error != nil {
		panic(error)
	}
}
