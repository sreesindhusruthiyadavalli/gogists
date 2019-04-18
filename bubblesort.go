package main

import
("fmt")

func SortString(s string) string {
	r := []rune(s)
	for i:=0; i < len(r); i++{
		for j:=0; j < len(r)-1-i;j++{
			if r[j] > r[j+1]{
				r[j], r[j+1] = r[j+1], r[j]
				fmt.Println(string(r))
		    }
		}
	}	
	return string(r)		
}

func main(){
	SortString("aangrams")
}