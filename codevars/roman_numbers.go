package main

import "fmt"

package kata
//import "os"
func Decode(roman string) int {
  num :=0
  for i,val :=range roman{
    switch (val) {
      case 'I' : num+=1;break
      case 'V' : if (i!=0 && roman[i-1]=='I'){
                    num+=3
                  } else{
                  num+=5;
                  }
      case 'X' : if (i!=0 && roman[i-1]=='I'){
                    num+=8
                  } else{
                    num+=10;
                  }
      case 'L' : if (i!=0 && roman[i-1]=='X'){
                    num+=30
                  } else{
                  num+=50;
                  }
      case 'C' : if (i!=0 && roman[i-1]=='X'){
                    num+=80;
                  } else{
                    num+=100;
                  }
      case 'D' : if (i!=0 && roman[i-1]=='C'){
                    num+=300
                  } else{
                  num+=500;
                  }
      case 'M' : if (i!=0 && roman[i-1]=='C'){
                    num+=800
                  } else{
                    num+=1000;
                  }
      //default: os.Exit(0)
    }
    }
  return num
}
func main() {

	fmt.Println(Decode("XXI"))
}
