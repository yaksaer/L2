package main

import (
	"errors"
	"fmt"
	"unicode"
)

func Unpack(str string) (*string, error) {
	inStr := []rune(str)
	var res []rune

	var slashFlag bool

	for i, sym := range inStr {
		count := 0
		if (unicode.IsNumber(sym) && i == 0) || (i > 2 && unicode.IsNumber(sym) &&
			unicode.IsNumber(inStr[i-1]) && inStr[i-2] != '\\') {
			return nil, errors.New("wrong string\n")
		} else if sym == '\\' && !slashFlag {
			slashFlag = true
			continue
		} else if slashFlag {
			res = append(res, sym)
			slashFlag = false
			continue
		} else if unicode.IsNumber(sym) {
			count = int(sym - '0')
			for j := 0; j < count-1; j++ {
				res = append(res, inStr[i-1])
			}
			continue
		}
		res = append(res, sym)
	}
	ret := string(res)
	return &ret, nil
}

func main() {
	str, _ := Unpack("a4bc2d5e")
	fmt.Println(*str)

	str, _ = Unpack("abcd")
	fmt.Println(*str)

	_, err := Unpack("45")
	fmt.Println(err)

	str, _ = Unpack("")
	fmt.Println(*str)

	str, _ = Unpack("qwe\\45 ")
	fmt.Println(*str)

	str, _ = Unpack("he3r9m8l2\\4\\5")
	fmt.Println(*str)

	str, _ = Unpack("qwe\\\\5")
	fmt.Println(*str)

	str, _ = Unpack("he3r8m8l2\\4\\5")
	fmt.Println(*str)
}
