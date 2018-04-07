
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

		case'|':
			frag2:=nfastack[len(nfastack)-1]
			nfastack=nfastack[:len(nfastack)-1]
			frag1:=nfastack[len(nfastack)-1]
			nfastack=nfastack[:len(nfastack)-1]

			accept :=state{}
			initial:=state{edge1:frag1.initial,edge2:frag2.initial}

			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept
			
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



}
