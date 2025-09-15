package main

import "fmt"

// Tour du joueur pendant le combat
func playerTurn(j *Character, m *Monster) {
	fmt.Println("\n=== Tour du joueur ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Utiliser un objet")
	fmt.Print("Choix : ")
	var choix int
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		dmg := 5
		m.hpActuels -= dmg
		if m.hpActuels < 0 {
			m.hpActuels = 0
		}
		fmt.Printf("⚔️ %s inflige %d dégâts à %s ! PV restant du monstre : %d/%d\n", j.nom, dmg, m.nom, m.hpActuels, m.hpMax)
	case 2:
		UseItem(j)
	default:
		fmt.Println("Choix invalide !")
	}
}

// Tour du monstre pendant le combat
func monsterTurn(j *Character, m *Monster) {
	fmt.Println("\n=== Tour du monstre ===")
	dmg := m.atk
	j.hpActuels -= dmg
	if j.hpActuels < 0 {
		j.hpActuels = 0
	}
	fmt.Printf("👹 %s inflige %d dégâts à %s ! PV restant du joueur : %d/%d\n", m.nom, dmg, j.nom, j.hpActuels, j.hpMax)
}

// Combat d'entraînement complet
func trainingFight(j *Character) {
	monstre := InitKappa()
	fmt.Println("\n Combat d'entraînement ! Affrontez le Kappa d'entraînement !")
	for j.hpActuels > 0 && monstre.hpActuels > 0 {
		playerTurn(j, &monstre)
		if monstre.hpActuels <= 0 {
			fmt.Println(" Vous avez vaincu le Kappa !")
			break
		}
		monsterTurn(j, &monstre)
		if j.hpActuels <= 0 {
			fmt.Println(" Vous êtes mort ! t nul ")
			break
		}
	}
}
