package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {

	{

		fmt.Println("Hello World")
		fmt.Println("Hello World")

	}
	countingLettersFunction()
}

func countingLettersFunction() {

	fmt.Println("counting letters function")

	s := "The quick brown fox jumps over the lazy dog"
	sReader := strings.NewReader(s)
	counts, err := countLetters(sReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(counts)
}

func countLetters(r io.Reader) (map[string]int, error) {

	fmt.Println("Counting letters")
	buff := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buff)
		for _, b := range buff[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
