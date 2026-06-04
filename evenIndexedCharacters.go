package main
import (
	"fmt"
	"strings"
)

func main(){
	name := "vitumbiko kamanga"
	name2 := strings.Split(name,"")
	var name3 []string
	for i:=0; i< len(name2); i++{
		if i%2 == 0{
			name2[i] = strings.ToUpper(name2[i])
		}
		name3 = append(name3,name2[i])
		fmt.Print(name3[i])
	}
}