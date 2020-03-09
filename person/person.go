package person

type Person struct {
	id int //unexported
	Name string //exported
	Age uint8 //exported
}