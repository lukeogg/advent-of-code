package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func main() {
	n := [][]rune{
		{'A', 'X', '3', '1'},
		{'A', 'Y', '6', '2'},
		{'A', 'Z', '0', '3'},
		{'B', 'X', '0', '1'},
		{'B', 'Y', '3', '2'},
		{'B', 'Z', '6', '3'},
		{'C', 'X', '6', '1'},
		{'C', 'Y', '0', '2'},
		{'C', 'Z', '3', '3'}}

	f, err := os.Open("./input.txt")
	check(err)

	r := bufio.NewReader(f)

	total := 0

	for {
		s, err := Readln(r)
		if err != nil {
			break
		}

		for x := range n {
			if s[0:1] == string(n[x][0]) && s[2:3] == string(+n[x][1]) {
				a := n[x][2] - '0'
				b := n[x][3] - '0'
				fmt.Println("Match! Score: ", s[0:1], s[2:3], a, b, a+b)
				total += int(a + b)
			}
		}
	}

	fmt.Println("Total: ", total)

	f.Close()
}
