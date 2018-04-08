
package main

import (

	"fmt"	
)

type state struct{

	symbol rune
	edge1 *state
	edge2 *state

}


type nfa struct{
	initial *state
	accept *state
}

func addState(list []*state, s *state, a *state)[]*state{

	list = append(list,s)

	if s!=a && s.symbol == 0 {
		list = addState(list,s.edge1,a)
		if s.edge2 !=nil {
			list = addState(list,s.edge2,a)
		}
	}
 return list
}

func pomatch(p string,s string)bool{

	po :=in2post(p)

	ismatch :=false

	ponfa :=poregtonfa(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:],ponfa.initial,ponfa.accept)


	for _,r := range s{

		for _,c := range current{

			if c.symbol == r{

			next = addState(next[:],c.edge1,ponfa.accept)

			}

		}

		current,next = next,[]*state{}
	}

	for _,c := range current {

		if c == ponfa.accept{
			ismatch = true
			break
		}
	}

	return ismatch
}

func poregtonfa(pofix string)*nfa{
	
	nfastack :=[]*nfa{}

	for _,r := range pofix{

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
			frag2:=nfastack[len(nfastack)-1]
			nfastack=nfastack[:len(nfastack)-1]
			frag1:=nfastack[len(nfastack)-1]
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

func in2post(infix string)string{

	specialschar:=map[rune]int{'*':10,'+':10,'?':10,'.':9,'|':8}

	pofix:=[]rune{}

	s:=[]rune{}

	for _,r :=range infix{

		switch {

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

fmt.Println(pomatch("a.b?","a"))

}
