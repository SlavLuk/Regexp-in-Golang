
package main

import (

	"fmt"	
)

func intopost(infix string)string{

	specialschar:=map[rune]int{'*':10,'.':9,'|':8}

	pofix:=[]rune{}

	s:=[]rune{}





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
