package combat

import (
	"YAMA/src/character"
	"YAMA/src/monster"
	"fmt"
)

// Gère les actions du joueur pendant son tour
func PlayerTurn(j *character.Character, m *monster.Monster) {
	fmt.Println("\n=== Tour du joueur ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Utiliser un objet")
	fmt.Print("Choix : ")

	var choix int
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		dmg := 5 + j.Niveau // dégâts augmentent avec le niveau
		m.HpActuels -= dmg
		if m.HpActuels < 0 {
			m.HpActuels = 0
		}
		fmt.Printf("⚔️ %s inflige %d dégâts à %s ! PV du monstre : %d/%d\n",
			j.Nom, dmg, m.Nom, m.HpActuels, m.HpMax)
	case 2:
		character.UseItem(j)
	default:
		fmt.Println("Choix invalide !")
	}
}
