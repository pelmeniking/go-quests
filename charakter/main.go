package main

import "fmt"

func main() {
	var name string = "GoKeule"
	level := 1
	hp := 100
	kritChance := 0.15
	lebt := true

	fmt.Println("=== Charakterbogen ===")
	fmt.Println("Name:", name)
	fmt.Println("Level:", level)
	fmt.Println("HP:", hp)
	fmt.Println("Krit-Chance:", kritChance)
	fmt.Println("Lebt:", lebt)

	fmt.Println("*** LEVEL UP! ***")
	level = level + 1
	hp = hp + 20
	kritChance = kritChance + 0.02

	fmt.Println("Level:", level)
	fmt.Println("HP:", hp)
	fmt.Printf("Krit-Chance: %.2f\n", kritChance)

	goblinHP := 80
	goblinSchaden := 15

	fmt.Println("Ein wilder Goblin erscheint", goblinHP, "HP")
	fmt.Println("Der Goblin greift mit", goblinSchaden, "an")
	fmt.Println(name, "erhält", goblinSchaden, "Schaden")
	hp = hp - goblinSchaden
	fmt.Println(name, "Hat nun", hp)
}
