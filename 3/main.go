package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type args struct {
	k *int
	n *bool
	r *bool
	u *bool
}

func newArgs() *args {
	flags := new(args)
	flags.k = flag.Int("k", 0, "column to sort")
	flags.n = flag.Bool("n", false, "numeric sort")
	flags.r = flag.Bool("r", false, "reverse")
	flags.u = flag.Bool("u", false, "unique")
	return flags
}

func main() {
	args := newArgs()
	flag.Parse()
	filename := flag.Arg(0)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("No file found")
		os.Exit(1)
	}
	data := make([]string, 0, 3)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	res := Sort(data, args)
	resFile, err := os.Create("out")
	if err != nil {
		fmt.Println("Error create file")
		os.Exit(1)
	}
	for _, str := range res {
		resFile.WriteString(str + "\n")
	}
}

func Sort(text []string, args *args) []string {
	columnSort := args.k
	isDefSort := true

	if *args.u {
		tmp := make([]string, 0, len(text))
		isDub := false
		for i := 0; i < len(text); i++ {
			for j := i + 1; j < len(text); j++ {
				if text[i] == text[j] {
					isDub = true
					break
				}
			}
			if !isDub {
				tmp = append(tmp, text[i])
			}
			isDub = false
		}
		text = tmp
	}

	if *columnSort > 0 {
		isDefSort = false
		data := make([][]string, len(text))
		for i, j := range text {
			data[i] = strings.Split(j, " ")
		}
		sort.Slice(data, func(i, j int) bool {
			var si string
			var sj string
			if len(data[i]) < *columnSort {
				si = data[i][len(data[i])-1]

			} else {
				si = data[i][*columnSort-1]
			}
			if len(data[j]) < *columnSort {
				sj = data[j][len(data[j])-1]
			} else {
				sj = data[j][*columnSort-1]
			}
			var siLower = strings.ToLower(si)
			var sjLower = strings.ToLower(sj)
			if siLower == sjLower {
				return si > sj
			}
			return siLower < sjLower

		})
		tmp := make([]string, 0, len(text))
		var s string
		for _, j := range data {
			s = strings.Join(j, " ")
			tmp = append(tmp, s)
		}
		text = tmp

	}

	if *args.n {
		//сортировать по числовому значению
		isDefSort = false
		num := make([]int, 0, len(text))
		var tmpNum int
		var err error
		for _, j := range text {
			tmpNum, err = strconv.Atoi(j)
			if err != nil {
				fmt.Println("Error convert string")
				os.Exit(2)
			}
			num = append(num, tmpNum)
		}
		if *args.r {
			sort.Sort(sort.Reverse(sort.IntSlice(num)))
		} else {
			sort.Ints(num)
		}
		tmp := make([]string, 0, len(num))
		for _, n := range num {

			tmp = append(tmp, strconv.Itoa(n))
		}
		text = tmp
	}

	if isDefSort {
		sort.Slice(text, func(i, j int) bool {
			var si string = text[i]
			var sj string = text[j]
			var si_lower = strings.ToLower(si)
			var sj_lower = strings.ToLower(sj)
			if si_lower == sj_lower {
				return si < sj
			}
			return si_lower < sj_lower
		})
	}

	if *args.r && !*args.n {
		sort.Sort(sort.Reverse(sort.StringSlice(text)))
	}

	return text
}
