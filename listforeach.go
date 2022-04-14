package piscine

/*
import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

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
*/
func ListForEach(l *List, f func(*NodeL)) {
	iter := l.Head
	for iter != nil {
		f(iter) // la func appliquée içi est déterminée dans la func main
		iter = iter.Next
	}
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}

/*
func main() {
	link := &List{}

	ListPushBack(link, "1")
	ListPushBack(link, "2")
	ListPushBack(link, "3")
	ListPushBack(link, "5")

	ListForEach(link, Add2_node)

	it := link.Head
	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}
}
*/
