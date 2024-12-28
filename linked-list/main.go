package main

import (
	"fmt"
)

func main(){
	ll := newMyLinkedListFomValues(1, 2, 3, 4, 5)

	fmt.Printf(ll.myLikedListToString())
}