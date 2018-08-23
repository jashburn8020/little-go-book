package c02structures

type SaiyanBasic struct {
	Name  string
	Power int
}

type SaiyanFather struct {
	Name   string
	Power  int
	Father *SaiyanFather
}

type Person struct {
	Name string
}

// SaiyanPerson is composed of a Person
type SaiyanPerson struct {
	*Person
	Power int
}

type SaiyanOverride struct {
	*Person
	Power int
}
