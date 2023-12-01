package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sorted(s string) string {
	runes := []rune(s)
	sort.Sort(RuneSlice(runes))
	return string(runes)
}

func part1() {
	fmt.Println("**************  Part 1 **************")

	f, err := os.Open("./input.txt")
	check(err)

	r := bufio.NewReader(f)

	total := 0

	for {
		s, err := Readln(r)
		if err != nil {
			break
		}
		// s = "aAAbcdefxyzABCDEFXYZabAAcd"

		fmt.Println(s)
		fmt.Println([]rune(s))

		a := s[0 : len(s)/2]
		b := s[len(s)/2 : len(s)]
		a = sorted(a)
		b = sorted(b)

		var c []byte

		// Check the two strings for matches
		fmt.Println("a: ", a)
		fmt.Println("b: ", b)

		y := 0 // floating index for b
		for x := range a {
			for y < len(b) {
				if a[x] == b[y] {
					fmt.Println("Match: ", a[x])
					// check for duplicates
					dup := false
					for z := range c {
						if c[z] == a[x] {
							fmt.Println("Duplicate: ", a[x])
							dup = true
							break
						}
					}
					if !dup {
						c = append(c, a[x])
					}
					break
				} else if a[x] > b[y] {
					fmt.Println("y++: ", y, a[x], b[y])
					y++
				} else if a[x] < b[y] {
					fmt.Println("x++: ", x, a[x], b[y])
					break
				}
			}
		}

		fmt.Println("c: ", c)

		ruckTotal := 0

		for x := range c {
			// Use unicode to get the priority number
			// a-z = 1-26
			// A-Z = 27-52
			n := c[x] - '`'
			//fmt.Println("n: ", n)

			// If the number is greater than 26, then it's a lowercase letter - reverse the priority
			if n > 26 {
				n -= 198
			}

			fmt.Print(n, " ")

			ruckTotal += int(n)
		}
		fmt.Println("Ruck Total: ", ruckTotal)
		total += ruckTotal
	}

	fmt.Println("Total: ", total)

	f.Close()
}
