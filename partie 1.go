package main

import "fmt"

const MaxInventory = 10

type Character struct {
	nom        string
	classe     string
	niveau     int
	hpMax      int
	hpActuels  int
	inventaire []string
	or         int
	skills     []string
}

func initCharacter(nom string, classe string, niveau int, hpMax int, hpActuels int, inventaire []string, or int, skills []string) Character {
	return Character{
		nom:        nom,
		classe:     classe,
		niveau:     niveau,
		hpMax:      hpMax,
		hpActuels:  hpActuels,
		inventaire: inventaire,
		or:         or,
		skills:     []string{"Coup de poing"},
	}
}

func displayInfo(c Character) {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom :", c.nom)
	fmt.Println("Classe :", c.classe)
	fmt.Println("Niveau :", c.niveau)
	fmt.Printf("Points de vie : %d / %d\n", c.hpActuels, c.hpMax)
	fmt.Println("Or :", c.or)
	fmt.Println("Inventaire :", c.inventaire)
	fmt.Println("Compétences :", c.skills)
	fmt.Println("==================================")
}

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
func removeInventoryAt(c *Character, idx int) {
	if idx < 0 || idx >= len(c.inventaire) {
		return
	}
	c.inventaire = append(c.inventaire[:idx], c.inventaire[idx+1:]...)
}

func addInventory(c *Character, item string) bool {
	if len(c.inventaire) >= MaxInventory {
		fmt.Println("Inventaire plein, impossible d'ajouter :", item)
		return false
	}
	c.inventaire = append(c.inventaire, item)
	return true
}

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
	used := false

	switch objet {
	case "Thé du Kitsune (狐茶)":
		c.hpActuels += 50
		if c.hpActuels > c.hpMax {
			c.hpActuels = c.hpMax
		}
		fmt.Println(" Tu as utilisé Thé du Kitsune ! PV restaurés.")
		used = true
	case "Potion de vie":
		c.hpActuels += 50
		if c.hpActuels > c.hpMax {
			c.hpActuels = c.hpMax
		}
		fmt.Println(" Tu as utilisé une Potion de vie !")
		used = true
	case "Potion de poison":
		poisonPot(c)
		used = true
	case "Livre de Sort : Boule de Feu":
		spellBook(c)
		used = true
	default:
		fmt.Println("Cet objet ne peut pas être utilisé pour l’instant.")
	}

	if used {
		removeInventoryAt(c, choix-1)
	}
}

func poisonPot(c *Character) {
	fmt.Println("Vous buvez une Potion de poison !")
	for i := 1; i <= 3; i++ {
		c.hpActuels -= 10
		if c.hpActuels < 0 {
			c.hpActuels = 0
		}
		fmt.Printf("Dégâts de poison (%ds) → PV : %d/%d\n", i, c.hpActuels, c.hpMax)
	}
}

func spellBook(c *Character) {
	for _, s := range c.skills {
		if s == "Boule de feu" {
			fmt.Println("Vous connaissez déjà ce sort.")
			return
		}
	}
	c.skills = append(c.skills, "Boule de feu")
	fmt.Println("🔥 Vous avez appris le sort : Boule de feu !")
}

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

func isDead(c *Character) bool {
	if c.hpActuels <= 0 {
		fmt.Println("\n*** Vous êtes mort... ***")
		c.hpActuels = c.hpMax / 2
		if c.hpActuels <= 0 {
			c.hpActuels = 1
		}
		fmt.Printf("Vous êtes ressuscité avec %d/%d PV.\n", c.hpActuels, c.hpMax)
		return true
	}
	return false
}

func player() {
	joueur := initCharacter(
		"Aiko (愛子)",
		"Samouraï (侍)",
		1,
		120,
		30,
		[]string{"Katana (刀)", "Yoroi no Tate (鎧の盾)"},
		100, // Or de départ
		[]string{},
	)

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
