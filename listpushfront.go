package piscine

// structs are already initialized in the previous git

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL // completely useless here because of the main func
}*/

func ListPushFront(l *List, data interface{}) {
	end := &NodeL{Data: data} //  on initialise un noeud et on met nos données à l'intérieur {....}
	if l.Head == nil {        // l.head = first node
		l.Head = end
	} else {
		end.Next = l.Head // on dit que la variable pour passer au suivant devient le premier noeud
		l.Head = end      // the first node become the last
	}
}

/*
func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}
*/
