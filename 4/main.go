package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func Anagram(arr []string) map[string][]string {
	countSym := make(map[string]map[rune]int)
	for _, j := range arr {
		j = strings.ToLower(j)
		countSym[j] = make(map[rune]int)
		for _, v := range j {
			countSym[j][v]++
		}
	}
	firstRes := make(map[string][]string)
	for i, j := range countSym {
		for str, val := range countSym {
			if reflect.DeepEqual(j, val) {
				firstRes[i] = append(firstRes[i], str)
				delete(countSym, str)
			}
		}

	}
	finalResult := make(map[string][]string)
	for _, j := range firstRes {
		if len(j) > 1 {
			sort.Strings(j)
			tmp := j
			finalResult[j[0]] = tmp
		}
	}

	return finalResult
}

func main() {
	arr := []string{"пятак", "пятка", "тяпка", "листок", "слиток",
		"столик", "рост", "СОРТ", "тРоС"}

	res := Anagram(arr)
	fmt.Println(res)
}
