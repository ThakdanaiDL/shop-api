package exception

type ItemCounting struct{}

func (e *ItemCounting) Error() string {
	return "Item Counting faliled"

}
