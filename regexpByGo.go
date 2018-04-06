
package main

import (


	"fmt"	
)

func intopost(infix string)string{

	postfix:=""

	return postfix
}

func main() {

	//Should print: ab.c*
	fmt.Println("Infix: ","a.b.c*")
	fmt.Println("Postfix: ",intopost("a.b.c*"))

}
