package main

import (
	"fmt"
	"reflect"
)

func pr(a interface{}){
	//fmt.Println(a)
	fmt.Println(reflect.ValueOf(a).Kind())
	fmt.Printf("%+v\n",reflect.ValueOf(a))
	fmt.Printf("%+v\n",reflect.Indirect(reflect.ValueOf(a)))
	fmt.Printf("%+v\n",reflect.Indirect(reflect.ValueOf(a)).FieldByName("Z").Interface())
	fmt.Printf("%+v\n",reflect.Indirect(reflect.ValueOf(a)).FieldByName("Z").Kind())
	fmt.Printf("%+v\n",reflect.Indirect(reflect.ValueOf(a)).FieldByName("Z").FieldByName("Id").Interface())
	
}

type(
	Parent struct{
		Z Child
	}
	
	Child struct{
		Name string
		Id string
	}

)

func main() {
	 x := &Parent{Child{Name:"Apple", Id:"1234" }   } 
			 
			
	//fmt.Println(x)
	pr(x)
	//fmt.Println(reflect.ValueOf(x))
}
