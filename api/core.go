package persona

type CardInfo struct {
	Name string
	Job  string
	Cost int
}

func DeckMake() CardInfo {
	var ci CardInfo

	ci.Name = "func DeckMake() "

	return ci
}
