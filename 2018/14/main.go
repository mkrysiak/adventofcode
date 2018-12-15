package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %s\n", part1(864801))
	fmt.Printf("Part 2: %d\n", part2("864801"))
}

func part1(recipes int) string {
	head := ring.New(2)
	elf1 := head
	elf2 := head.Next()
	tail := head.Next()

	elf1.Value = 3
	elf2.Value = 7

	iterations := recipes + 10

	for i := 0; i < iterations; i++ {
		nextRecipes := elf1.Value.(int) + elf2.Value.(int)
		if nextRecipes > 9 {
			r1 := nextRecipes / 10
			r2 := nextRecipes % 10
			newRecipe := ring.New(2)
			newRecipe.Value = r1
			newRecipe.Next().Value = r2
			tail.Link(newRecipe)
			tail = tail.Move(2)
		} else {
			newRecipe := ring.New(1)
			newRecipe.Value = nextRecipes
			tail.Link(newRecipe)
			tail = tail.Next()
		}
		elf1 = elf1.Move(elf1.Value.(int) + 1)
		elf2 = elf2.Move(elf2.Value.(int) + 1)
	}

	scores := ""
	head = head.Move(recipes)
	for i := 0; i < 10; i++ {
		scores = scores + fmt.Sprintf("%s", strconv.Itoa(head.Value.(int)))
		head = head.Next()
	}

	return scores
}

func part2(recipes string) int {
	head := ring.New(2)
	elf1 := head
	elf2 := head.Next()
	tail := head.Next()

	elf1.Value = 3
	elf2.Value = 7

	recipeCount := 2
	lastDigit, _ := strconv.Atoi(recipes)
	lastDigit = lastDigit % 10
	recipeLen := len(recipes)

	for {
		// fmt.Println(tail.Value)
		rs := 0
		nextRecipes := elf1.Value.(int) + elf2.Value.(int)
		if nextRecipes > 9 {
			r1 := nextRecipes / 10
			r2 := nextRecipes % 10
			newRecipe := ring.New(2)
			newRecipe.Value = r1
			newRecipe.Next().Value = r2
			tail.Link(newRecipe)
			tail = tail.Move(2)
			recipeCount += 2
			rs = 2
		} else {
			newRecipe := ring.New(1)
			newRecipe.Value = nextRecipes
			tail.Link(newRecipe)
			tail = tail.Next()
			recipeCount++
			rs = 1
		}
		elf1 = elf1.Move(elf1.Value.(int) + 1)
		elf2 = elf2.Move(elf2.Value.(int) + 1)

		if tail.Value.(int) == lastDigit {
			t := tail.Move((-1 * recipeLen) + 1)
			s := ""
			for i := 0; i < recipeLen; i++ {
				s = s + fmt.Sprintf("%s", strconv.Itoa(t.Value.(int)))
				t = t.Next()
			}
			if recipes == s {
				break
			}

		}

		if rs == 2 && tail.Move(-1).Value.(int) == lastDigit {
			t := tail.Move((-1 * recipeLen))
			s := ""
			for i := 0; i < recipeLen; i++ {
				s = s + fmt.Sprintf("%s", strconv.Itoa(t.Value.(int)))
				t = t.Next()
			}
			if recipes == s {
				recipeCount--
				break
			}

		}
	}

	return recipeCount - recipeLen
}
