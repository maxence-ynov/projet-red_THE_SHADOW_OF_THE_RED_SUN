package main

import "fmt"

// Définition du monstre
type Monster struct {
	nom        string
	hpMax      int
	hpActuels  int
	atk        int
	pattern    []string         // séquence des attaques ou capacités
	special    func(*Character) // capacité spéciale du monstre
	rewardOr   int
	rewardItem string
}

// Fonction pour créer un Kappa (河童) d'entraînement
func InitKappa() Monster {
	return Monster{
		nom:       "Kappa (河童) d'entraînement",
		hpMax:     40,
		hpActuels: 40,
		atk:       5,
		pattern:   []string{"attack", "special"},
		special: func(j *Character) {
			fmt.Println("💦 Éclaboussure d’eau !")
		},
		rewardOr:   20,
		rewardItem: "Potion de vie",
	}
}

// Fonction pour créer un Oni mineur
func InitOni() Monster {
	return Monster{
		nom:       "Oni mineur (小鬼)",
		hpMax:     60,
		hpActuels: 60,
		atk:       10,
		pattern:   []string{"attack", "attack", "special"},
		special: func(j *Character) {
			fmt.Println("💥 Coup brutal ! Dégâts x1.5")
			dmg := 15
			j.hpActuels -= dmg
			if j.hpActuels < 0 {
				j.hpActuels = 0
			}
		},
		rewardOr:   30,
		rewardItem: "Amulette protectrice",
	}
}

// Fonction pour créer un Kitsune farceur
func InitKitsune() Monster {
	return Monster{
		nom:       "Kitsune farceur (狐)",
		hpMax:     75,
		hpActuels: 75,
		atk:       8,
		pattern:   []string{"attack", "special"},
		special: func(j *Character) {
			fmt.Println("🦊 Illusion ! Esquive automatique du prochain coup")
		},
		rewardOr:   40,
		rewardItem: "Thé du Kitsune",
	}
}

// Fonction pour créer un Tengu
func InitTengu() Monster {
	return Monster{
		nom:       "Tengu (天狗)",
		hpMax:     90,
		hpActuels: 90,
		atk:       12,
		pattern:   []string{"attack", "attack", "special"},
		special: func(j *Character) {
			fmt.Println("💨 Coup de vent ! Défense réduite pendant 2 tours")
		},
		rewardOr:   50,
		rewardItem: "Plume de Tengu",
	}
}

// Fonction pour créer une Yuki-onna
func InitYukiOnna() Monster {
	return Monster{
		nom:       "Yuki-onna (雪女)",
		hpMax:     120,
		hpActuels: 120,
		atk:       10,
		pattern:   []string{"attack", "attack", "special"},
		special: func(j *Character) {
			fmt.Println("❄️ Souffle glacé ! 30% de chance de geler le joueur")
		},
		rewardOr:   70,
		rewardItem: "Éventail de glace",
	}
}

// Fonction pour créer une Nure-onna
func InitNureOnna() Monster {
	return Monster{
		nom:       "Nure-onna (濡女)",
		hpMax:     150,
		hpActuels: 150,
		atk:       15,
		pattern:   []string{"attack", "special", "attack"},
		special: func(j *Character) {
			fmt.Println("🐍 Constriction ! Immobilise le joueur et inflige 10 dégâts par tour (x2)")
		},
		rewardOr:   90,
		rewardItem: "Épée améliorée",
	}
}

// Fonction pour créer une Rokurokubi
func InitRokurokubi() Monster {
	return Monster{
		nom:       "Rokurokubi (轆轤首)",
		hpMax:     180,
		hpActuels: 180,
		atk:       12,
		pattern:   []string{"attack", "attack", "special"},
		special: func(j *Character) {
			fmt.Println("🕯️ Cou allongé ! Attaque à distance imparable")
		},
		rewardOr:   120,
		rewardItem: "Amulette de longévité",
	}
}

// Fonction pour créer un Gashadokuro
func InitGashadokuro() Monster {
	return Monster{
		nom:       "Gashadokuro (餓者髑髏)",
		hpMax:     220,
		hpActuels: 220,
		atk:       18,
		pattern:   []string{"attack", "attack", "special"},
		special: func(j *Character) {
			fmt.Println("🦴 Écrasement ! Inflige 25 dégâts fixes")
			j.hpActuels -= 25
			if j.hpActuels < 0 {
				j.hpActuels = 0
			}
		},
		rewardOr:   150,
		rewardItem: "Potion de force",
	}
}

// Fonction pour créer Enma
func InitEnma() Monster {
	return Monster{
		nom:       "Enma (閻魔)",
		hpMax:     300,
		hpActuels: 300,
		atk:       20,
		pattern:   []string{"attack", "attack", "special"},
		special: func(j *Character) {
			fmt.Println("🔥 Jugement ! Inflige 25% des HP max du joueur")
			dmg := j.hpMax / 4
			j.hpActuels -= dmg
			if j.hpActuels < 0 {
				j.hpActuels = 0
			}
		},
		rewardOr:   200,
		rewardItem: "Armure sacrée",
	}
}

// Fonction pour créer Yamata no Orochi
func InitYamataNoOrochi() Monster {
	return Monster{
		nom:       "Yamata no Orochi (八岐大蛇)",
		hpMax:     500,
		hpActuels: 500,
		atk:       25,
		pattern:   []string{"attack", "attack", "special"},
		special: func(j *Character) {
			fmt.Println("🐉 Hydra strike ! 8 attaques de 3 dégâts chacune")
			for i := 0; i < 8; i++ {
				j.hpActuels -= 3
				if j.hpActuels < 0 {
					j.hpActuels = 0
				}
			}
		},
		rewardOr:   500,
		rewardItem: "Katana légendaire",
	}
}

// Fonction pour créer le Shinigami
func InitShinigami() Monster {
	return Monster{
		nom:       "Shinigami (死神, Dieu de la Mort)",
		hpMax:     999,
		hpActuels: 999,
		atk:       30,
		pattern:   []string{"attack", "special", "special2", "special3"},
		special: func(j *Character) {
			fmt.Println("💀 Fauche de l’âme ! Inflige 50 dégâts fixes")
			j.hpActuels -= 50
			if j.hpActuels < 0 {
				j.hpActuels = 0
			}
		},
		rewardOr:   1000,
		rewardItem: "Victoire finale",
	}
}
