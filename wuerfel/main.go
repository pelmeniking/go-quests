package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	wurf := rand.IntN(20) + 1
	fmt.Println("Du würfelst einen W20...")
	fmt.Println("Ergebnis:", wurf)

	switch wurf {
	case 20:
		fmt.Println("KRITTISCHER TREFFER!")
	case 1:
		fmt.Println("Patzer! Du stolperst über dein eigenes Schwert.")
	default:
		fmt.Println("Ein normaler Wurf.")
	}
}
