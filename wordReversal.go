package main
import(
	"fmt"
	"strings"
)

func main (){
	name := "this is vitumbiko kamanga"
	name2 := strings.Split(name, " ")
	var name3 []string
	var name4 []string
	var name5 []string
	for i:=0; i< len(name2); i++{
		if len(name2[i]) >= 5{
			for j := (len(name2[i])-1); j>=0; j--{
				name4 = strings.Split(name2[i],"")
				name5 = append(name5, name4[j])
			}
			strings.Join(name5, "")
			name3 = append(name3, name5[i])
		}else{
			name3 = append(name3, name2[i])
		}
	}
	
	fmt.Println(name2)
}