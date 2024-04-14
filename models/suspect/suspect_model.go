package suspect

type Suspect struct {
	Name string
}

var suspects []Suspect

func getAllSuspects() []Suspect {
	return suspects
}

func fillSuspects(filledSuspects []Suspect) []Suspect {
	suspects = filledSuspects
	return suspects
}

func addSuspect(suspect Suspect) {
	suspects = append(suspects, suspect)
}

func removeAllSuspects() {
	suspects = []Suspect{}
}
