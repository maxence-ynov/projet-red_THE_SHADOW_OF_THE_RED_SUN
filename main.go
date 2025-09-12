package main

import (
	"YAMA/src/character"
	"YAMA/src/combat"
	"fmt"
)

func main() {
	// Création du personnage
	joueur := character.InitCharacter(
		Nom:"Aiko (愛子)",
		"Samouraï (侍)",
		1,
		120,
		30,
		[]string{"Katana (刀)", "Yoroi no Tate (鎧の盾)"},
		100,
	)

	for {
		fmt.Println("\n===== MENU PRINCIPAL =====")
		fmt.Println("1. Infos personnage")
		fmt.Println("2. Voir l'inventaire")
		fmt.Println("3. Utiliser un objet")
		fmt.Println("4. Aller au marchand")
		fmt.Println("5. Entrainement")
		fmt.Println("6. Forgeron")
		fmt.Println("7. Quitter")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			character.DisplayInfo(joueur)
		case 2:
			character.AccessInventory(joueur)
		case 3:
			character.UseItem(&joueur)
		case 4:
			character.Merchant(&joueur)
		case 5:
			combat.TrainingFight(&joueur)
		case 6:
			combat.BlacksmithJP(&joueur)
		case 7:
			fmt.Println("👋 Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
