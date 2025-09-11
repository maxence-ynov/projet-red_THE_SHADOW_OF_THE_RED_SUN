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

// Fonction du marchand
func merchant(c *Character) {
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

func main() {
	// Création du personnage
	joueur := initCharacter(
		"Aiko (愛子)",
		"Samouraï (侍)",
		1,
		120,
		30,
		[]string{"Katana (刀)", "Yoroi no Tate (鎧の盾)"},
		100, // Or de départ
	)
	// Boucle du menu principal
	for {
		fmt.Println(`⠀⠀⠘⡇⠀⠀⠀⢠⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢷⠀⢠⢣⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢘⣷⢸⣾⣇⣶⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣿⣿⣿⣹⣿⣿⣷⣿⣆⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢼⡇⣿⣿⣽⣶⣶⣯⣭⣷⣶⣿⣿⣶⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠸⠣⢿⣿⣿⣿⣿⡿⣛⣭⣭⣭⡙⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠸⣿⣿⣿⣿⣿⠿⠿⠿⢯⡛⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣾⣿⡿⡷⢿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⡔⣺⣿⣿⣽⡿⣿⣿⣿⣟⡳⠦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢠⣭⣾⣿⠃⣿⡇⣿⣿⡷⢾⣭⡓⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣾⣿⡿⠷⣿⣿⡇⣿⣿⣟⣻⠶⣭⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣋⣵⣞⣭⣮⢿⣧⣝⣛⡛⠿⢿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⣀⣀⣠⣶⣿⣿⣿⣿⡿⠟⣼⣿⡿⣟⣿⡇⠀⠙⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⡼⣿⣿⣿⢟⣿⣿⣿⣷⡿⠿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠉⠁⠀⢉⣭⣭⣽⣯⣿⣿⢿⣫⣮⣅⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢀⣿⣟⣽⣿⣿⣿⣿⣾⣿⣿⣯⡛⠻⢷⣶⣤⣄⡀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢀⡞⣾⣿⣿⣿⣿⡟⣿⣿⣽⣿⣿⡿⠀⠀⠀⠈⠙⠛⠿⣶⣤⣄⡀⠀⠀
⠀⠀⠀⣾⣸⣿⣿⣷⣿⣿⢧⣿⣿⣿⣿⣿⣷⠁⠀⠀⠀⠀⠀⠀⠀⠈⠙⠻⢷⣦
⠀⠀⠀⡿⣛⣛⣛⣛⣿⣿⣸⣿⣿⣿⣻⣿⣿⠆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢸⡇⣿⣿⣿⣿⣿⡏⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢰⣶⣶⣶⣾⣿⢃⣿⣿⣿⣿⣯⣿⣭⠁⠀⠀`)
		fmt.Println("\n===== MENU PRINCIPAL =====")
		fmt.Println("1. Infos personnage")
		fmt.Println("2. Mode Histoire")
		fmt.Println("3. Voir l'inventaire")
		fmt.Println("4. Utiliser un objet")
		fmt.Println("5. Aller au marchand")
		fmt.Println("6. Entrainement")
		fmt.Println("7. Quitter")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			displayInfo(joueur)
		case 2:
			accessInventory(joueur)
		case 3:
			useItem(&joueur)
		case 4:
			merchant(&joueur)
		case 5:
			fmt.Println("👋 Au revoir !")
			return
		case 6:
			trainingFight(&joueur)
		case 7:
			fmt.Println("👋 Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

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
func initKappa() Monster {
	return Monster{
		nom:       "Kappa (河童) d'entraînement",
		hpMax:     40,
		hpActuels: 40,
		atk:       5,
		pattern:   []string{"attack", "special"},
		special: func(j *Character) {
			fmt.Println("💦 Éclaboussure d’eau !")
			// Chance de réduire l'attaque du joueur, par ex :
			// j.atk -= 2
		},
		rewardOr:   20,
		rewardItem: "Potion de vie",
	}
}

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
		useItem(j)
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
	monstre := initKappa()

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

// ======== TOUS LES MONSTRES/BOSS ========
// Fonction pour créer un Oni mineur
func initOni() Monster {
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
func initKitsune() Monster {
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
func initTengu() Monster {
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
func initYukiOnna() Monster {
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
func initNureOnna() Monster {
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
func initRokurokubi() Monster {
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
func initGashadokuro() Monster {
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
func initEnma() Monster {
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
func initYamataNoOrochi() Monster {
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
func initShinigami() Monster {
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
