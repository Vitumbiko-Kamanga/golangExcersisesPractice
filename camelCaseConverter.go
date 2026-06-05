package main

import (
	"fmt"
	"strings"
)

func main(){
	name := "vitumbiko-kamanga malawi"
	name2 := strings.Split(name,"")
	var name3 []string
	for i:= 0; i< len(name); i++{
		if name2[i] == "_" || name2[i] == "-" || name2[i] == " "{
			continue
		}
		
		if i >= 1 {
			j := i-1
			if name2[j] == "_" || name2[j] == "-" || name2[j] == " "{
				name2[i] = strings.ToUpper(name2[i])
			}
		}
		name3 = append(name3, name2[i])
	}
	fmt.Println(strings.Join(name3,""))
}