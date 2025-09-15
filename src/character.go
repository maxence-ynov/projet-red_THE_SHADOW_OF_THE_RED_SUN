package main

import "fmt"

// Définition du personnage
type Character struct {
	nom        string
	classe     string
	niveau     int
	hpMax      int
	hpActuels  int
	inventaire []string
	or         int
}

// Fonction pour initialiser un personnage
func initCharacter(nom string, classe string, niveau int, hpMax int, hpActuels int, inventaire []string, or int) Character {
	return Character{
		nom:        nom,
		classe:     classe,
		niveau:     niveau,
		hpMax:      hpMax,
		hpActuels:  hpActuels,
		inventaire: inventaire,
		or:         or,
	}
}

// Afficher les infos du personnage
func displayInfo(c Character) {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom :", c.nom)
	fmt.Println("Classe :", c.classe)
	fmt.Println("Niveau :", c.niveau)
	fmt.Printf("Points de vie : %d / %d\n", c.hpActuels, c.hpMax)
	fmt.Println("Or :", c.or)
	fmt.Println("Inventaire :", c.inventaire)
	fmt.Println("==================================")
}

// Afficher l’inventaire
func accessInventory(c Character) {
	fmt.Println("=== Inventaire ===")
	if len(c.inventaire) == 0 {
		fmt.Println("Inventaire vide !")
		return
	}
	for i, item := range c.inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("=================")
}

// Utiliser un objet
func useItem(c *Character) {
	if len(c.inventaire) == 0 {
		fmt.Println("Votre Inventaire est vide !")
		return
	}
	fmt.Println("Quel objet voulez-vous utiliser ?")
	for i, item := range c.inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	var choix int
	fmt.Print("Choix : ")
	fmt.Scanln(&choix)
	if choix < 1 || choix > len(c.inventaire) {
		fmt.Println("Choix invalide.")
		return
	}
	objet := c.inventaire[choix-1]
	switch objet {
	case "Thé du Kitsune (狐茶)":
		c.hpActuels += 50
		if c.hpActuels > c.hpMax {
			c.hpActuels = c.hpMax
		}
		fmt.Println(" Tu as utilisé Thé du Kitsune ! PV restaurés.")
	case "Potion de vie":
		c.hpActuels += 50
		if c.hpActuels > c.hpMax {
			c.hpActuels = c.hpMax
		}
		fmt.Println(" Tu as utilisé une Potion de vie !")
	default:
		fmt.Println("Cet objet ne peut pas être utilisé pour l’instant.")
	}
	// Retirer l’objet après utilisation
	c.inventaire = append(c.inventaire[:choix-1], c.inventaire[choix:]...)
}
