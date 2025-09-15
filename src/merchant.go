package src

import "fmt"

// Fonction du marchand
func Merchant(c *Character) {
	fmt.Println("\n=== Marchand ===")
	fmt.Println("1. Potion de vie (3 pièces d'or)")
	fmt.Println("2. Thé du Kitsune (5 pièces d'or)")
	fmt.Println("3. Quitter")
	fmt.Print("Choix : ")
	var choix int
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		if c.or >= 3 {
			c.or -= 3
			c.inventaire = append(c.inventaire, "Potion de vie")
			fmt.Println(" Potion de vie ajoutée à l'inventaire !")
		} else {
			fmt.Println("Pas assez d'or !")
		}
	case 2:
		if c.or >= 5 {
			c.or -= 5
			c.inventaire = append(c.inventaire, "Thé du Kitsune (狐茶)")
			fmt.Println(" Thé du Kitsune ajouté à l'inventaire !")
		} else {
			fmt.Println("Pas assez d'or !")
		}
	case 3:
		fmt.Println("Vous quittez le marchand.")
	default:
		fmt.Println("Choix invalide.")
	}
}
