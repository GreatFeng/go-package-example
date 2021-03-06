package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("12345678"))
	r := bufio.NewReader(rb)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Printf("%d, %v\n", n, err) // 8, <nil>
	fmt.Println(string(buf[:n]))   // 12345678
}
