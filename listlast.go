package piscine

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = n
	}
}
*/

func ListLast(l *List) interface{} {
	for l.Head != nil { // pour l'actuel différent de vide
		if l.Head.Next == nil { // si le suivant est vide
			return l.Head.Data // retourne ce qu'il y a dans l'actuel
		}
		l.Head = l.Head.Next // l'actuel est remplacé par le suivant pour boucler, si l'actuel est different de vide. La boucle s'arrete lorsque le suivant est vide.
	}
	return nil
}

/*
func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, 3)
	ListPushBack(link, "1")

	fmt.Println(ListLast(link))
	fmt.Println(ListLast(link2))
}
*/
