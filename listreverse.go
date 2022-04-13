package piscine

/*
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
func ListReverse(l *List) {
	current := l.Head   // La variable sert de stockage
	var previous *NodeL // necéssaire pour définir le type de "nil"
	previous = nil
	for current != nil {
		next := current.Next    // stock le l.Head suivant
		current.Next = previous // efface le suivant
		previous = current
		current = next // l'actuel prends la place du suivant
	}
	l.Head = previous
}

/*
func main() {
	link := &List{}

	ListPushBack(link, 1)
	ListPushBack(link, 2)
	ListPushBack(link, 3)
	ListPushBack(link, 4)

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}
*/
