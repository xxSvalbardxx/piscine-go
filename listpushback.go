package piscine

type NodeL struct {
	Data interface{} // is an element of the node
	Next *NodeL      // indicate where/what is the next node
}

type List struct {
	Head *NodeL // the only node that we use and modify here
	Tail *NodeL // completely useless here because of the main func
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

/*
func main() {
	link := &List{}

	ListPushBack(link, "Hello")
	ListPushBack(link, "man")
	ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}*/
