package stack

import (
	"go_playground/list"
	"strings"
	"time"
)

const (
	dog string = "dog"
	cat string = "cat"
)

type Animal struct {
	kind string
	name string
}

type animalInList struct {
	animal Animal
	now    string
}

type animalshelter struct {
	animal map[string]*list.List
}

func NewAnimalShelter() animalshelter {
	as := animalshelter{animal: map[string]*list.List{}}
	as.animal[dog] = new(list.List)
	as.animal[cat] = new(list.List)
	return as
}

func assertValidType(kind string) {
	if kind != dog || kind != cat {
		panic("not a valid type: only dog and cat allowed")
	}
}

func (as animalshelter) Insert(a Animal) {
	assertValidType(a.kind)
	curTime := time.Now().Local().String()
	animalWithAdmittedTime := animalInList{animal: a, now: curTime}
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

	if as.animal[cat].Size() == 0 && as.animal[dog].Size() == 0 {
		return Animal{}
	}

	if as.animal[cat].Size() == 0 {
		animalDog := as.animal[dog].Get(0).(animalInList)
		as.animal[cat].Delete(animalDog)
		return animalDog.animal
	}

	if as.animal[dog].Size() == 0 {
		animalCat := as.animal[cat].Get(0).(animalInList)
		as.animal[dog].Delete(animalCat)
		return animalCat.animal

	}

	animalCat := as.animal[dog].Get(0).(animalInList)
	animalDog := as.animal[cat].Get(0).(animalInList)

	if strings.Compare(animalCat.now, animalDog.now) < 0 {
		as.animal[cat].Delete(animalDog)
		return animalCat.animal
	} else {
		as.animal[dog].Delete(animalDog)
		return animalDog.animal
	}
}
