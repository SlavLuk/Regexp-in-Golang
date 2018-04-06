
package main

import (

	"fmt"	
)

func intopost(infix string)string{

	specialschar:=map[rune]int{'*':10,'.':9,'|':8}

	pofix:=[]rune{}

	s:=[]rune{}

	for _,r :=range infix{

		switch{

			case r == '(':

				s=append(s,r)

			case r==')':

				for s[len(s)-1]!='(' {
					pofix = append(pofix,s[len(s)-1])
					s = s[:len(s)-1]

				}

				s=s[:len(s)-1]

			case specialschar[r]>0:

				for len(s)>0 && specialschar[r]<=specialschar[s[len(s)-1]]{
					pofix,s = append(pofix,s[len(s)-1]),s[:len(s)-1]
				}
				s=append(s,r)

		default:
			pofix=append(pofix,r)


		}
	}

  for len(s)>0{

	pofix,s = append(pofix,s[len(s)-1]),s[:len(s)-1]
  }

	return string(pofix)
}

func main() {

	//Should print: ab.c*.
	fmt.Println("Infix: ","a.b.c*")
	fmt.Println("Pofix: ",intopost("a.b.c*"))

	//Should print: abd|.*
	fmt.Println("Infix: ","(a.(b|d))*")
	fmt.Println("Pofix: ",intopost("(a.(b|d))*"))

	//Should print: abd|.c*.
	fmt.Println("Infix: ","a.(b|d).c*")
	fmt.Println("Pofix: ",intopost("a.(b|d).c*"))

	//Should print: abb.+.c.
	fmt.Println("Infix: ","a.(b.b)+.c")
	fmt.Println("Pofix: ",intopost("a.(b.b)+.c"))

}
