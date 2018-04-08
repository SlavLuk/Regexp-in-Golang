
package main

import (

	"fmt"	
	"strings"
)

//create struct for each state
type state struct{

	symbol rune
	edge1 *state
	edge2 *state

}

//create struct nfa containing all states
type nfa struct{
	initial *state
	accept *state
}
//helper recursive function to add states to list
func addState(list []*state, s *state, a *state)[]*state{

	list = append(list,s)

	if s!=a && s.symbol == 0 {
		//recursive call
		list = addState(list,s.edge1,a)

	if s.edge2 !=nil {
		//recursive call
		list = addState(list,s.edge2,a)

		}
	}
 return list
}

func match(infix string,str string)bool{

	//convert infix into postfix notation
	postfix :=in2post(infix)

	ismatch :=false

	//construct non-deterministic finite automata
	postnfa :=postregtonfa(postfix)

	current := []*state{}
	next := []*state{}

	current = addState(current[:],postnfa.initial,postnfa.accept)

	for _,r := range str{

		for _,c := range current{

			if c.symbol == r{

			next = addState(next[:],c.edge1,postnfa.accept)

			}
		}

		current,next = next,[]*state{}
	}

	for _,c := range current {

		if c == postnfa.accept{
			ismatch = true
			break
		}
	}

	return ismatch
}

func postregtonfa(postfix string)*nfa{
	
	nfastack :=[]*nfa{}

	for _,r := range postfix{

		switch r{

		case'.':

			frag2:=nfastack[len(nfastack)-1]
			nfastack=nfastack[:len(nfastack)-1]
			frag1:=nfastack[len(nfastack)-1]
			nfastack=nfastack[:len(nfastack)-1]

			frag1.accept.edge1 = frag2.initial

			nfastack = append(nfastack,&nfa{initial:frag1.initial,accept:frag2.accept})	

		case '+':
			//popping one fragment off nfastack
			frag := nfastack[len(nfastack)-1]
			nfastack := nfastack[:len(nfastack)-1]
			//creating new accept state
			accept := state{}
			//create new initial state with edge to fragment
			initial := state{edge1: frag.initial}
			//setting fragment accept state arrows back to start of fragment and to the new accept state
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			//appending new concatenated fragment to stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case'|':
			//popping one fragment off nfastack
			frag1:=nfastack[len(nfastack)-1]
			nfastack=nfastack[:len(nfastack)-1]

			frag2:=nfastack[len(nfastack)-1]
			nfastack=nfastack[:len(nfastack)-1]
			
			//creating new accept state
			accept :=state{}
			//create new initial state with edge to fragment
			initial:=state{edge1:frag1.initial,edge2:frag2.initial}
			//setting fragment accept state arrows back to start of fragment and to the new accept state
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept
			//appending new concatenated fragment to stack
			nfastack = append(nfastack,&nfa{initial:&initial,accept:&accept})

		case '?':
			//popping one fragment off nfastack
			frag:=nfastack[len(nfastack)-1]
			nfastack=nfastack[:len(nfastack)-1]

			accept :=state{}
			initial :=state{edge1:frag.initial,edge2:&accept}
			frag.accept.edge1 = &accept
			frag.accept.edge2 = &accept

			nfastack = append(nfastack,&nfa{initial:&initial,accept:&accept})
		
		case'*':
			frag :=nfastack[len(nfastack)-1]
			nfastack=nfastack[:len(nfastack)-1]

			accept :=state{}
			initial :=state{edge1:frag.initial,edge2:&accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept


			nfastack = append(nfastack,&nfa{initial:&initial,accept:&accept})

			
		default:

			accept :=state{}
			initial:=state{symbol:r,edge1:&accept}

			nfastack=append(nfastack,&nfa{initial:&initial,accept:&accept})
		}
	}

	return nfastack[0]

}

//shunting yard algorithm used to convert infix into postfix notation
func in2post(infix string)string{

	//precedence of special chars
	specialschar:=map[rune]int{'*':10,'+':10,'?':10,'.':9,'|':8}

	postfix:=[]rune{}

	s:=[]rune{}

	for _,r :=range infix{

		switch {

			case r == '(':

				s=append(s,r)

			case r==')':

				for s[len(s)-1]!='(' {

					postfix = append(postfix,s[len(s)-1])
					s = s[:len(s)-1]

				}

				s=s[:len(s)-1]

			case specialschar[r]>0:

				for len(s)>0 && specialschar[r]<=specialschar[s[len(s)-1]]{
					postfix,s = append(postfix,s[len(s)-1]),s[:len(s)-1]
				}
				s=append(s,r)

			default:

				postfix=append(postfix,r)
		}
	}
 			for len(s)>0{

				postfix,s = append(postfix,s[len(s)-1]),s[:len(s)-1]
 		 }

	return string(postfix)
}



//validate first input char
func validateChar(input string)bool{

	input = strings.ToLower(input)

	if input[0] >=97 && input[0] <= 122{


		return true
	}

	return false
}

func main() {

	//
	fmt.Println("Please enter  a reg expression (a.b.c|d+.e):")	

	var regexp string
	//read user input
	fmt.Scan(&regexp)

	for !validateChar(regexp){

			fmt.Println("Please enter valid reg exp...")
			fmt.Scan(&regexp)
	}


	fmt.Println("Please enter a string to check against (abc):")

		var str string
		fmt.Scan(&str)

	matched :=match(regexp,str)

	//output result to user
	if matched {
	fmt.Println("Full match")
	} else {
	fmt.Println("No match")
	}

	fmt.Println("Please enter  1 to continue or -1 to exit:")	
	var exit int
	fmt.Scan(&exit)


for exit !=-1 {

	fmt.Println("Please enter  a reg expression (a.b.c|d+.e):")	

	var regexp string
	fmt.Scan(&regexp)

	for !validateChar(regexp){

			fmt.Println("Please enter valid reg exp...")
			fmt.Scan(&regexp)
	}


fmt.Println("Please enter a string to check against (abc):")

	var str string
	fmt.Scan(&str)

matched :=match(regexp,str)

	//output result to user
		if matched {

			fmt.Println("Full match")

		} else {
			fmt.Println("No match")
		}

fmt.Println("Please enter  1 to continue or -1 to exit:")	
	
	fmt.Scan(&exit)

}


}
 