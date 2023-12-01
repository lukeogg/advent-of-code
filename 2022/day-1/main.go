package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

	f, err := os.Open("./input.txt")
	check(err)

	r := bufio.NewReader(f)
	s, e := Readln(r)
	i, err := strconv.Atoi(s)
	check(err)

	count := i
	var counts []int

	for e == nil {
		fmt.Println(i)
		s, e = Readln(r)
		if e != nil {
			break
		}

		if s != "" {
			i, err := strconv.Atoi(s)
			check(err)
			count += i
			fmt.Println(i)
		} else {
			counts = append(counts, count)
			// Reset the counter
			count = 0
		}
	}
	fmt.Println(counts)

	// Print the answer - Part 1
	sort.Ints(counts)
	fmt.Println(counts)
	fmt.Println("Part 1 - Answer: ", counts[len(counts)-1])

	// Part 2
	top_n := 3
	count = 0
	var top []int = counts[len(counts)-top_n:]

	fmt.Println(top)

	for i := range top {
		count += top[i]
	}
	// Print the answer - Part 2
	fmt.Println("Part 2 - Answer: ", count)

	f.Close()
}
