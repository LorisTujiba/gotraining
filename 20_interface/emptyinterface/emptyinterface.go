package emptyinterface

type Animals interface{} //no method just empty interface

type Animal struct {
	Breed string
}

type Horse struct {
	Animal
	Power int
}

type Dog struct {
	Animal
	Speed int
}

type Fish struct {
	Animal
	Mammal bool
}
