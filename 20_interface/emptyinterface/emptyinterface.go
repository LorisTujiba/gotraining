package emptyinterface

//Animals is exported
type Animals interface{} //no method just empty interface

//Animal is exported
type Animal struct {
	Breed string
}

//Horse is exported
type Horse struct {
	Animal
	Power int
}

//Dog is exported
type Dog struct {
	Animal
	Speed int
}

//Fish is exported
type Fish struct {
	Animal
	Mammal bool
}
