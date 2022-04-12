package piscine

/*import (
	"fmt"
)

type NodeL struct {
	Data interface{} // is an element of the node
	Next *NodeL      // indicate where/what is the next node
}

type List struct {
	Head *NodeL // the only node that we use and modify here
	Tail *NodeL // completely useless here because of the main func
}

func ListPushFront(l *List, data interface{}) {
	end := &NodeL{Data: data} //  on initialise un noeud et on met nos données à l'intérieur {....}
	if l.Head == nil {        // l.head = first node
		l.Head = end
	} else {
		end.Next = l.Head // on dit que la variable pour passer au suivant devient le premier noeud
		l.Head = end      // the first node become the last
	}
}*/

func ListSize(l *List) int {
	count := 0
	for l.Head != nil {
		count++
		l.Head = l.Head.Next // le suivant remplace le precedent
	}
	return count
}

/*
func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}
*/
