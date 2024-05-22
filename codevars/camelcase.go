package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ToJadenCase(str string) string {
	word := strings.Split(str, " ")
	JadenString := ""
	for i, first := range word {
		f := string(unicode.ToUpper(rune(first[0])))
		JadenString += f + first[1:]

		if i != len(word)-1 {
			JadenString += " "
		}
	}
	return JadenString
}
func main() {

	fmt.Println(ToJadenCase("XXI"))
}

// 	package kata
// import (

//   "strings"
//   )
// func ToJadenCase(str string) string {
//   for i,first := range str{
//     if(first==' '){
//       ru:=str[i+1]
//       str[i+1]=strings.ToUpper(ru)
//     }
//     }
//   return str
// }
