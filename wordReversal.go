package main
import(
	"fmt"
	"strings"
)

func toReverse(str string) string{
	name2 := strings.Split(str, " ")
	var name3 []string
	var name4 []string
	var name5 []string
	for i:=0; i< len(name2); i++{
		if len(name2[i]) >= 5{
			for j := (len(name2[i])-1); j>=0; j--{
				name4 = strings.Split(name2[i],"")
				name5 = append(name5, name4[j])
			}
			name2[i] = strings.Join(name5, "")
			name3 = append(name3, name2[i])
			name3 = append(name3, " ")
			name5 = nil
		}else{
			name3 = append(name3, name2[i])
			name3 = append(name3, " ")
		}
	}
	return strings.Join(name3, "")
}

func main (){
	name := "this is vitumbiko from blantyre malawi"
	fmt.Println(toReverse(name))
}
