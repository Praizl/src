package main

/*import (
	//"bufio"
	//"io"
	/*"bufio"
	"fmt"
	"log"
	"os"
)*/

//func main() {
/*var newFile *os.File
fmt.Println(newFile)
var err error
newFile, err = os.Create("f1la.exe")
if err != nil {
	fmt.Println(err, newFile)
	os.Exit(1)
	log.Fatal(err)
}
err = os.Truncate("f1la.txt", 50)
if err != nil {
	fmt.Println(err)
	log.Fatal(err)
}
file, err := os.Open("file.txt")
if err != nil {
	fmt.Println(err)
	log.Fatal(err)
}
file.Close()
var fileInfo os.FileInfo
fileInfo, err = os.Stat("file.txt")
fmt.Println(err)
fmt.Println(fileInfo.Name())

file, err = os.Open("file.txt")
if err != nil {
	if os.IsNotExist(err) {
		fmt.Println(file, "file doesn't exist")
	}
}
file.Close()
err = os.Rename("file.txt", "flaiii.txt")
if err != nil {
	log.Fatal(err)
}
err = os.Remove("f1la.txt")
if err != nil {
	log.Fatal(err)
}

file, err = os.OpenFile(
	"b.txt",
	os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
	0644,
)
if err != nil {
	log.Fatal(err)
}
defer file.Close()

sliceOfByte := []byte("EA Sports! To the game")
noOfBytes, err := file.Write(sliceOfByte)
if err != nil {
	log.Fatal(err)
}
fmt.Println(noOfBytes)
byteSlice := []byte("Activision!!!!")
err = ioutil.WriteFile("c.txt", byteSlice, 0644)
if err != nil {
	log.Fatal(err)
}*/

//using buffer
/*file3, err3 := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
if err3 != nil {
	log.Fatal(err3)
}
defer file3.Close()

bufferedWriter := bufio.NewWriter(file3)
btSlyce := []byte{1, 2, 3}
checknBtSlyce, err := bufferedWriter.Write(btSlyce)
if err != nil {
	log.Fatal(err)
}
log.Printf("no of bytes currently in buffer:%d", checknBtSlyce)

log.Printf("%d", bufferedWriter.Available())

addedString, err := bufferedWriter.WriteString("\nadding this string to the buffer")
if err != nil {
	log.Fatal(err)
}
_ = addedString
log.Print(bufferedWriter.Buffered())
bufferedWriter.Flush()*/

//reading a file
/*file, err := os.Open("test.txt")
if err != nil {
	log.Fatal(err)
}
defer file.Close()

sliceToRead := make([]byte, 5)
bytesRead, err := io.ReadFull(file, sliceToRead)
if err != nil {
	log.Fatal(err)
}
log.Print(bytesRead)
log.Printf("%s\n", sliceToRead)

file, err = os.Open("files.go")
if err != nil {
	log.Fatal(err)
}
defer file.Close()

data, err := ioutil.ReadAll(file)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("data %s\n", data)
fmt.Println("bytes read:", len(data))

data, err = ioutil.ReadFile("test.txt")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("data as string:%s\n", data)*/

//using scanner

/*file, err := os.Open("my_file.txt")
if err != nil {
	log.Fatal(err)
}
defer file.Close()

scanned := bufio.NewScanner(file)

scanSuccesful := scanned.Scan()
if !scanSuccesful {
	scnError := scanned.Err()
	if scnError == nil {
		log.Println("Whole file scanned")
	} else {
		log.Fatal(scnError)
	}
}
fmt.Println("first line:", scanned.Text())

//to print the whole file
for scanned.Scan() {
	fmt.Println(scanned.Text())
}
if scnError := scanned.Err(); err != nil {
	log.Fatal(scnError)
}*/

//taking input from command line & scanning

/*scanner := bufio.NewScanner(os.Stdin)

scanner.Scan()

text := scanner.Text()
bytes := scanner.Bytes()
fmt.Println("text input:", text)
fmt.Println("bytes input:", bytes)

//getting it to scan through the whole file
for scanner.Scan() {
	text = scanner.Text()
	fmt.Println("text entered:", text)
	if text == "end scan" {
		fmt.Println("closing scan program...")
		break
	}
}
fmt.Println("program ended")
if err := scanner.Err(); err != nil {
	log.Println(err)
}*/
//}
