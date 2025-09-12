package combat

import (
	"YAMA/src/character"
	"fmt"
)

type Monster struct {
	Nom        string
	HpMax      int
	HpActuels  int
	Atk        int
	Pattern    []string
	Special    func(*character.Character)
	RewardOr   int
	RewardItem string
}

func PlayerTurn(j *character.Character, m *Monster) {
	fmt.Println("\n=== Tour du joueur ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Utiliser un objet")
	fmt.Print("Choix : ")

	var choix int
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		dmg := 5
		m.HpActuels -= dmg
		if m.HpActuels < 0 {
			m.HpActuels = 0
		}
		fmt.Printf("⚔️ %s inflige %d dégâts à %s ! PV monstre : %d/%d\n", j.Nom, dmg, m.Nom, m.HpActuels, m.HpMax)
	case 2:
		character.UseItem(j)
	default:
		fmt.Println("Choix invalide !")
	}
}

func MonsterTurn(j *character.Character, m *Monster) {
	fmt.Println("\n=== Tour du monstre ===")
	dmg := m.Atk
	j.HpActuels -= dmg
	if j.HpActuels < 0 {
		j.HpActuels = 0
	}
	fmt.Printf("👹 %s inflige %d dégâts à %s ! PV joueur : %d/%d\n", m.Nom, dmg, j.Nom, j.HpActuels, j.HpMax)
}
