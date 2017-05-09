package stack

import (
	"go_playground/list"
)

const (
	dog string = "dog"
	cat string = "cat"
)

type Animal struct {
	kind string
	name string
}

func CreateAnimal(k string, n string) Animal {
	return Animal{kind: k, name: n}
}

func (animal Animal) GetKind() string {
	return animal.kind
}

func (animal Animal) GetName() string {
	return animal.name
}

type animalInList struct {
	animal Animal
	id     uint64
}

//uint64 is more than enough to simulate a single animal shelter
type animalshelter struct {
	animal  map[string]*list.List
	startId uint64
}

func NewAnimalShelter() animalshelter {
	as := animalshelter{animal: map[string]*list.List{}}
	as.animal[dog] = new(list.List)
	as.animal[cat] = new(list.List)
	as.startId = 1
	return as
}

func assertValidType(kind string) {
	if !(kind == dog || kind == cat) {
		panic("not a valid type: only dog and cat allowed")
	}
}

func (as *animalshelter) Insert(a Animal) {
	assertValidType(a.kind)
	animalWithAdmittedTime := animalInList{animal: a, id: as.startId}
	as.startId++
	as.animal[a.kind].Append(animalWithAdmittedTime)
}

func (as animalshelter) GetAnimal(kind string) Animal {
	assertValidType(kind)
	if as.animal[kind].Size() == 0 {
		return Animal{}
	}
	animalWithAdmittedTime := as.animal[kind].Get(0).(animalInList)
	as.animal[kind].Delete(animalWithAdmittedTime)
	return animalWithAdmittedTime.animal
}

func (as animalshelter) GetAnyAnimal() Animal {

	if as.animal[cat].Size() == 0 {
		return as.GetAnimal(dog)
	}

	if as.animal[dog].Size() == 0 {
		return as.GetAnimal(cat)
	}

	animalCat := as.animal[cat].Get(0).(animalInList)
	animalDog := as.animal[dog].Get(0).(animalInList)

	if animalCat.id < animalDog.id {
		return as.GetAnimal(cat)
	} else {
		return as.GetAnimal(dog)
	}
}
