package src

import "fmt"

// Fonction du forgeron
func BlacksmithJP(c *Character) {
	for {
		fmt.Println("\n🔥🏯 Bienvenue au Forgeron Japonais 🏯🔥")
		fmt.Println("1. 🗡️ Kabuto du guerrier (兜) - +10 PV max (5 pièces) 🐗🪶")
		fmt.Println("2. 🛡️ Do-Maru léger (胴丸) - +25 PV max (5 pièces) 🐺🛡️")
		fmt.Println("3. 🥾 Tabi du vent (足袋) - +15 PV max (5 pièces) 🐺🐗")
		fmt.Println("4. ⚔️ Katana enchanté (刀) - +5 attaque (20 pièces) 🔥🪨")
		fmt.Println("5. 🏹 Arc du Kitsune (狐の弓) - +10 attaque distance (15 pièces) 🦊🪶")
		fmt.Println("6. 🎐 Amulette Shinto (神のお守り) - Protection malus (10 pièces) ⛩️💎")
		fmt.Println("0. 🔙 Retour")
		fmt.Print("Choix : ")
		var choix int
		fmt.Scanln(&choix)
		if choix == 0 {
			fmt.Println("↩️ Retour au menu principal...")
			return
		}
		// Coût et nom des items
		cost := map[int]int{1: 5, 2: 5, 3: 5, 4: 20, 5: 15, 6: 10}
		items := map[int]string{
			1: "Kabuto du guerrier 🗡️", 2: "Do-Maru léger 🛡️", 3: "Tabi du vent 🥾",
			4: "Katana enchanté ⚔️", 5: "Arc du Kitsune 🏹", 6: "Amulette Shinto 🎐",
		}
		// Vérifier l'or
		if c.or < cost[choix] {
			fmt.Println("💰 Pas assez d’or ! Retournez gagner plus d'or avant d'acheter.")
			continue
		}
		// Vérifier inventaire max
		if len(c.inventaire) >= 10 {
			fmt.Println("🎒 Inventaire plein ! Videz quelques objets avant d'acheter.")
			continue
		}
		// Achat réussi
		c.or -= cost[choix]
		c.inventaire = append(c.inventaire, items[choix])
		fmt.Printf("✅ %s ajouté à ton inventaire ! Il te reste %d pièces d'or.\n", items[choix], c.or)
	}
}
