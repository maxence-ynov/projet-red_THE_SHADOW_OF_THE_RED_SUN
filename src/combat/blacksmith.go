package combat

import (
	"YAMA/src/character"
	"fmt"
)

func BlacksmithJP(c *character.Character) {
	for {
		fmt.Println("\n🔥 Forgeron 🏯")
		fmt.Println("1. Kabuto du guerrier (+10 PV max, 5 or)")
		fmt.Println("2. Do-Maru léger (+25 PV max, 5 or)")
		fmt.Println("3. Tabi du vent (+15 PV max, 5 or)")
		fmt.Println("0. Quitter")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)
		if choix == 0 {
			return
		}

		cost := 5
		items := map[int]string{
			1: "Kabuto du guerrier", 2: "Do-Maru léger", 3: "Tabi du vent",
		}

		if c.Or < cost {
			fmt.Println("Pas assez d’or !")
			continue
		}
		c.Or -= cost
		c.Inventaire = append(c.Inventaire, items[choix])
		fmt.Printf("%s ajouté ! Il vous reste %d or.\n", items[choix], c.Or)
	}
}
