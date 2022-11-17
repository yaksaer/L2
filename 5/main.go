package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type matchString struct {
	str   string
	print bool
}

type flags struct {
	F *bool
	t *string
	c *bool

	A *int
	B *int
	C *int

	v *bool
	i *bool
	n *bool
}

func newFlags() *flags {
	flags := new(flags)
	flags.F = flag.Bool("F", false, "fixed")
	flags.t = flag.String("t", "", "template")
	flags.c = flag.Bool("c", false, "count")

	flags.A = flag.Int("A", 0, "after")
	flags.B = flag.Int("B", 0, "before")
	flags.C = flag.Int("C", 0, "context")

	flags.v = flag.Bool("v", false, "invert")
	flags.i = flag.Bool("i", false, "ignore-case")
	flags.n = flag.Bool("n", false, "line num")
	return flags
}

func main() {
	flags := newFlags()
	flag.Parse()

	fileName := flag.Arg(0)

	//open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var strs []matchString

	//read file line by line
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		strs = append(strs, matchString{fileScanner.Text(), false})
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	if *flags.t == "" {
		log.Fatal("need template (-t)")
	}

	strs, err = match(strs, *flags.t, *flags.F, *flags.i) //match
	if err != nil {
		log.Fatal(err)
	}
	countStrs := count(strs) //count strings with match

	strs = context(strs, *flags.A, *flags.B, *flags.C) //if you need more strings for output

	if *flags.v { //invert
		strs = invert(strs)
	}

	if *flags.c { //print result
		fmt.Println(countStrs)
	} else {
		print(strs, *flags.n)
	}
}

// match - looking for lines that match
func match(strs []matchString, template string, fixed bool, ignoreCase bool) ([]matchString, error) {
	if ignoreCase {
		template = strings.ToLower(template)
	}
	regex := regexp.MustCompile(template)
	var temp string

	if fixed {
		for i, val := range strs {
			if ignoreCase {
				if strings.ToLower(val.str) == template {
					strs[i].print = true
				}
			} else {
				if val.str == template {
					strs[i].print = true
				}
			}
		}
	} else {
		for i, val := range strs {
			if ignoreCase {
				temp = strings.ToLower(val.str)
				if regex.Match([]byte(temp)) {
					strs[i].print = true
				}
			} else {
				if regex.Match([]byte(val.str)) {
					strs[i].print = true
				}
			}
		}
	}

	return strs, nil
}

// invert - invert marks line for output
func invert(strs []matchString) []matchString {
	for i, val := range strs {
		if val.print {
			strs[i].print = false
		} else {
			strs[i].print = true
		}
	}
	return strs
}

// contex - marks lines for output before and after the match
func context(strs []matchString, after int, before int, context int) []matchString {
	if context > 0 {
		after = context
		before = context
	}

	if after > 0 {
		for i, val := range strs {
			if val.print {
				for j, c := 1, after; c > 0; j, c = j+1, c-1 {
					if i-j < 0 {
						break
					}
					strs[i-j].print = true
				}
			}
		}
	}

	if before > 0 {
		for i := len(strs) - 1; i >= 0; i-- {
			if strs[i].print {
				for j, c := 1, before; c > 0; j, c = j+1, c-1 {
					if i+j > len(strs)-1 {
						break
					}
					strs[i+j].print = true
				}
			}
		}
	}

	return strs
}

// print - print all strings
func print(strs []matchString, lineNum bool) {
	for i, val := range strs {
		if val.print {
			if lineNum {
				fmt.Println(i+1, val.str)
			} else {
				fmt.Println(val.str)
			}
		}
	}
}

// count - number of matches
func count(strs []matchString) int {
	var countRes int
	for _, val := range strs {
		if val.print {
			countRes++
		}
	}
	return countRes
}
