package main

import (
	"fmt"
)

func main() {
	var heure, minute int
	fmt.Printf("Donne moi l'heure: ")
	fmt.Scanln(&heure)

	fmt.Printf("Donne moi la minute ")
	fmt.Scanln(&minute)

	if minute < 59 {
		minute = minute + 1
	} else {
		minute = 0
		if heure < 24 {
			heure = heure + 1
		} else {
			heure = 0
		}
	}
	fmt.Println(" Dans une minute il sera  ", heure, " heures ", minute)
}
