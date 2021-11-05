package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lireChaine() string {

	var donneeLue string
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	scanner.Scan()                        // lancement du scanner
	donneeLue = scanner.Text()            // stockage du résultat du scanner dans une variable
	return donneeLue
}
func lireEntier() int {
	var nombre int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nombre, _ = strconv.Atoi(scanner.Text()) // conversion du type string en int
	return nombre
}

func lireReel() float64 {
	var donneeLue string
	var nombreLu float64
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	scanner.Scan()                        // lancement du scanner
	donneeLue = scanner.Text()            // stockage du résultat du scanner dans une variable
	nombreLu, _ = strconv.ParseFloat(donneeLue, 8)
	return nombreLu
}

func muain() {

	var p1, p2 int
	fmt.Printf("Donner la première valeur : ")
	p1 = lireEntier() // ou fmt.scanln(ax) pour lire l'enteree
	fmt.Printf("Donner la deuxième valeur: ")
	p2 = lireEntier()

	if p1 <= 0 && 0 <= p2 {
		fmt.Println("la reponse est negative")
	} else {
		fmt.Println("la reponse est positive")
	}

}
