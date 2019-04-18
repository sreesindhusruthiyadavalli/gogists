package main

import(
"io/ioutil"
"fmt"
"strings"
)

func readFromFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error while reading file " + err.Error())
		return "", err
	}
	return string(data), nil
}


func main(){
	data, err := readFromFile("./cipher.txt")
	if err != nil{
		fmt.Println(err)
	}
	//fmt.Println(data)
	s := strings.Split(data, ",")

	for i, char := range s {
		fmt.Println('%c', s[i])
	}
	
}