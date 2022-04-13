package piscine

/*
import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL // the only node that we use and modify here
	Tail *NodeL // completely useless here because of the main func
}
*/
func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	for l != nil && count < pos {
		count++
		l = l.Next
	}
	return l
}

/*
func ListPushBack(l *List, data interface{}) { // "l" is  ; "*List" is the receiver ; "data interface {}" is the prepend(ajouté) element ; (...) called a method receiver
	n := &NodeL{Data: data} // a temporary var to store the first element
	if l.Head == nil {      // if Head is empty ( nil: non initialised / zero value )
		l.Head = n // the element inside &NodeL ( {Data: data} ) is now stored in the Head
	} else {
		current := l.Head         // current(actuel)
		for current.Next != nil { // if NodeL's Data of the next Node
			current = current.Next
		}
		current.Next = n // return the three case for which we have modified the Head because the node after is empty
	}
}

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 7))
}
*/
