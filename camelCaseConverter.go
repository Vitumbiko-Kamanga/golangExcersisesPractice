package main

import (
	"fmt"
	"strings"
)

func camelCaseConverter(str string) string{
	// converting the string str to a slice of strings
	name2 := strings.Split(str,"")
	var name3 []string
	for i:= 0; i< len(str); i++{
		// checking for a particular character to be ignored to come up with the camel case
		if name2[i] == "_" || name2[i] == "-" || name2[i] == " "{
			continue
		}

		if i >= 1 {
			j := i-1
			// checking if the previous character was a character to be ignored
			if name2[j] == "_" || name2[j] == "-" || name2[j] == " "{
				// converting the character followed by the ignored character to Upper case
				name2[i] = strings.ToUpper(name2[i])
			}
		}
		// appending the characters to a slice excluding the ignored character
		name3 = append(name3, name2[i])
	}
	// returning back one string from the slice
	return strings.Join(name3,"")
}
func main(){
	name := "vitumbiko-kamanga malawi"
	fmt.Println(camelCaseConverter(name))
	
}