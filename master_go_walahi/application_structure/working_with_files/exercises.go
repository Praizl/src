package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	/*file, err := os.OpenFile("ïnfo.txt", os.O_CREATE, 0644)
	_ = file
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	err = os.Remove("ïnfo.txt")
	if err != nil {
		log.Fatal(err)
	}*/
	//use os.stat to check if file exists
	/*_, err = os.Stat("info.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Rename("info.txt", "data.txt")
	if err != nil {
		log.Fatal(err)
	}*/

	file, err := os.OpenFile("info.txt", os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	read, err1 := io.ReadAll(file)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("data read as string %s", read)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println("text in file:", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
