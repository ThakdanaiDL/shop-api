package exception

type Itemlisting struct {
}

func (e *Itemlisting) Error() string {
	return "Item listing faliled"

}
