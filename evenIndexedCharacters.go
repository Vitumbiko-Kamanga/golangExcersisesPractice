package main
import (
	"fmt"
	"strings"
)

func strConv(str string) string{
	//converting the string to a slice of letters
	str2 := strings.Split(str,"")

	for i:=0; i< len(str2); i++{
		if i%2 == 0{
			// changing the even character to upper case
			str2[i] = strings.ToUpper(str2[i])
		}else{
			str2[i] = strings.ToLower(str2[i])
		}
	}
	// returning the single string
	return strings.Join(str2, "")
}

func main(){
	fmt.Println(strConv("VITUMBIKO kamanga"))
	
}