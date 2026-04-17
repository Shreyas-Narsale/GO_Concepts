package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
/*
os → low-level file & system operations
io → generic input/output abstractions
os = “open the door” (file system access)
io = “move the data through the door , handles data flow (read/write streams)
*/
func main() {
	// os package 
	file, err := os.Open("file.txt")       // read-only
	file, err := os.Create("file.txt")     // create or truncate
	file, err := os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE, 0644)
	/*
	os.O_RDONLY → read only
	os.O_WRONLY → write only
	os.O_RDWR → read + write
	os.O_APPEND → append
	os.O_CREATE → create if not exists
	os.O_TRUNC → truncate
	*/
	// 0644  // 4 - read , 2 = write , 1 excute

	file.Write([]byte("Hello"))
	file.WriteString("World")

	data := make([]byte, 100)
	n, err := file.Read(data)
	defer file.Close()
	
	info, _ := os.Stat("file.txt")
	fmt.Println(info.Name())
	fmt.Println(info.Size())

	// read dir
	files, _ := os.ReadDir(".")

	for _, f := range files {
	    fmt.Println(f.Name())
	}
	// io
	// Core interface in Go
	// Represents any readable stream: Everything is a stream of bytes.
	/*Used for:
	files
	network
	buffers
	*/
	src, _ := os.Open("a.txt")
	dst, _ := os.Create("b.txt")

	io.Copy(dst, src)

	data, _ := io.ReadAll(file)
	io.WriteString(file, "Hello")
	/*
	o.ReadAll → reads entire content into memory
	file.Read → reads chunk by chunk
	*/
}

func ReadFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Unable to read file")
	}

	fmt.Println("Data is", string(data))
}
