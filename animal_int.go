package main

type AnimalList struct {
	id    int
	name  string
	voice string
	leg   int
}

type AnimalListint interface {
	Add(a AnimalList) error
	Update(name string, voice string, leg int, id int) error
	Delete(id int) error
	GetTask(id int)  error
}
