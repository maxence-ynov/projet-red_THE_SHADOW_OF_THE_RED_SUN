package character

import "fmt"

func DisplayInfo(c Character) {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom :", c.Nom)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Niveau)
	fmt.Printf("Points de vie : %d / %d\n", c.HpActuels, c.HpMax)
	fmt.Println("Or :", c.Or)
	fmt.Println("Inventaire :", c.Inventaire)
	fmt.Println("==================================")
}

func AccessInventory(c Character) {
	fmt.Println("=== Inventaire ===")
	if len(c.Inventaire) == 0 {
		fmt.Println("Inventaire vide !")
		return
	}
	for i, item := range c.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("=================")
}

func UseItem(c *Character) {
	if len(c.Inventaire) == 0 {
		fmt.Println("Votre Inventaire est vide !")
		return
	}

	fmt.Println("Quel objet voulez-vous utiliser ?")
	for i, item := range c.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	var choix int
	fmt.Print("Choix : ")
	fmt.Scanln(&choix)

	if choix < 1 || choix > len(c.Inventaire) {
		fmt.Println("Choix invalide.")
		return
	}

	objet := c.Inventaire[choix-1]

	switch objet {
	case "Thé du Kitsune (狐茶)", "Potion de vie":
		c.HpActuels += 50
		if c.HpActuels > c.HpMax {
			c.HpActuels = c.HpMax
		}
		fmt.Println("Objet utilisé ! PV restaurés.")
	default:
		fmt.Println("Cet objet ne peut pas être utilisé.")
	}
	c.Inventaire = append(c.Inventaire[:choix-1], c.Inventaire[choix:]...)
}
