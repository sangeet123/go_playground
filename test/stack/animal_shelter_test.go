package getmintest

import (
	"go_playground/stack"
	"testing"
)

func TestGetMinIntOp(t *testing.T) {
	s := stack.NewAnimalShelter()
	animals := []stack.Animal{stack.CreateAnimal("cat", "kitty2"), stack.CreateAnimal("dog", "bob2"),
		stack.CreateAnimal("dog", "bob3"), stack.CreateAnimal("cat", "kitty1"),
		stack.CreateAnimal("dog", "bob5"), stack.CreateAnimal("dog", "bob4"),
		stack.CreateAnimal("dog", "bob1"), stack.CreateAnimal("cat", "kitty3"),
		stack.CreateAnimal("cat", "kitty4"), stack.CreateAnimal("cat", "kitty5")}

	for _, animal := range animals {
		s.Insert(animal)
	}

	animal := s.GetAnyAnimal()

	if animal.GetKind() != "cat" {
		t.Error("expected cat but got", animal.GetKind())
	}

	if animal.GetName() != "kitty2" {
		t.Error("expected kitty2 but got", animal.GetName())
	}

	animal = s.GetAnyAnimal()

	if animal.GetName() != "bob2" {
		t.Error("expected bob2 but got", animal.GetName())
	}

	animal = s.GetAnimal("dog")
	animal = s.GetAnimal("dog")
	animal = s.GetAnimal("dog")
	animal = s.GetAnimal("dog")

	if animal.GetName() != "bob1" {
		t.Error("expected bob1 but got", animal.GetName())
	}

	animal = s.GetAnyAnimal()
	animal = s.GetAnyAnimal()
	animal = s.GetAnyAnimal()
	animal = s.GetAnyAnimal()

	if animal.GetName() != "kitty5" {
		t.Error("expected kitty5 but got", animal.GetName())
	}

}
