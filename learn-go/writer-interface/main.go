package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	name string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.name))
	return err
}

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Throws error %s \n", err)
	}

	defer file.Close()

	person1 := person{
		name: "raja",
	}

	var newBuffer bytes.Buffer

	person1.writeOut(file)
	person1.writeOut(&newBuffer)

	fmt.Println(newBuffer.String())

	// str := []byte("Hello gophers!!!")

	// numOfBytes, err := file.Write(str)

	// if err != nil {
	// 	log.Fatalf("Throws error %s \n", err)
	// }

	// fmt.Println(numOfBytes)

	// by := bytes.NewBufferString("Hello ")
	// fmt.Println(by.String())
	// by.WriteString("Gophers !")
	// fmt.Println(by)

	// by.Reset()

	// by.WriteString("Today is tuesday ")
	// fmt.Println(by.String())
	// by.Write([]byte{72, 101, 108, 108, 111})
	// fmt.Println(by)
}
