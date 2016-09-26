package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"io"
)

func main() {
	file,err:= os.Create("samp.txt")

	if err !=nil {
		log.Fatal(io.ErrNoProgress)
	}

	file.WriteString("This is some random text")

	file.Close()

	stream, err := ioutil.ReadFile("samp.txt")

	if(err!=nil){
		log.Fatal(err)
	}

	readString :=string(stream)
	fmt.Println(readString)


}
