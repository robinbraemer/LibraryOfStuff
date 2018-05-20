package main

import (
	"fmt"
)

func Function(s *string, fun func(string) string) string {
	*s = fun(*s)
	return *s
}


func main() {
	function := func(s string) string {
		return s + " nice"
	}
	
	v1 := "Yo"
	v2 := "Cool"
	
	fmt.Println(Function(&v1, function))
	fmt.Println(Function(&v2, function))
	
	fmt.Println(v1)
}
