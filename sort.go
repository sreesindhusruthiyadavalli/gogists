package main

import (
	"fmt"
	"sort"
)

type TempRune []rune

func (r TempRune) Len() int           { return len(r) }
func (r TempRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r TempRune) Less(i, j int) bool { return r[i] < r[j] }

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(TempRune(r))
	return string(r)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(SortString("aangrams"))
}
