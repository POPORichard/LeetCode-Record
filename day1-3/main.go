package main

import (
	"fmt"
)

func main(){
	s := "a1c1e1"
	fmt.Println(replaceDigits(s))
}

func replaceDigits(s string) string {
	length := len(s)
	new := make([]uint8,0,length)
	for i:=1;i<len(s);i+=2{
		new = append(new,s[i-1])
		new = append(new,s[i]-48+s[i-1])
	}
	if len(new)<length{
		new = append(new,s[length-1])
	}
	return string(new)

}
