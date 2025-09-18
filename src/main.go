package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// Efface la console
func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Impossible de nettoyer la console pour cette OS.")
	}
}

// Définition du personnage
type Character struct {
	nom        string
	classe     string
	niveau     int
	hpMax      int
	hpActuels  int
	atk        int
	inventaire []string
	or         int
	experience int
	levelUpExp int
	armure     int // Réduction des dégâts reçus
}

// Définition du monstre
type Monster struct {
	nom        string
	hpMax      int
	hpActuels  int
	atk        int
	pattern    []string
	special    func(*Character)
	rewardOr   int
	rewardItem string
}

// Initialisation du personnage
func initCharacter(nom string, classe string, niveau int, hpMax int, hpActuels int, atk int, inventaire []string, or int) Character {
	return Character{
		nom:        nom,
		classe:     classe,
		niveau:     niveau,
		hpMax:      hpMax,
		hpActuels:  hpActuels,
		atk:        atk,
		inventaire: inventaire,
		or:         or,
		experience: 0,
		levelUpExp: 100,
		armure:     0,
	}
}

// Afficher les infos du personnage
func displayInfo(c Character) {
	clearScreen()
	fmt.Println("\n=== 🧮 Statistiques du Joueur ===")
	fmt.Printf("Niveau : %d\n", c.niveau)
	fmt.Printf("PV : %d/%d\n", c.hpActuels, c.hpMax)
	fmt.Printf("ATK : %d\n", c.atk)
	fmt.Printf("Armure : %d\n", c.armure)
	fmt.Printf("Or : %d\n", c.or)
	fmt.Printf("Expérience : %d/%d\n", c.experience, c.levelUpExp)
	fmt.Println("Inventaire :")
	for _, item := range c.inventaire {
		fmt.Printf("- %s\n", item)
	}
	fmt.Println("=============================")
	fmt.Print("\nAppuyez sur Entrée pour continuer...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

// Afficher le lore lettre par lettre
func printLore(text string) {
	clearScreen()
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()
}

// Afficher le lore du boss lettre par lettre
func printBossLore(level int) {
	lore := ""
	switch level {
	case 1:
		lore = "Sous les arches effondrées d’un ancien village, un petit oni rôde, brandissant une arme grossière faite d’os et de fer rouillé. Ses yeux luisent d’une malice enfantine, mais chaque coup qu’il porte révèle une force brutale et imprévisible."
	case 2:
		lore = "Dans la brume des lanternes éteintes, un renard rusé vous invite à le suivre par des illusions changeantes. Ses tours semblent inoffensifs, mais ses rires masquent une cruauté subtile qui peut vous perdre à jamais."
	case 3:
		lore = "Perché au sommet d’un pont suspendu brisé, un Tengu aux ailes effilochées jauge votre honneur. Son bec acéré et son éventail divin manipulent le vent, testant votre détermination à gravir la Tour."
	case 4:
		lore = "Dans un couloir gelé de la Tour, une femme spectrale aux yeux d’hiver avance sans bruit. Son souffle glacé fige le sang, et sa beauté tragique murmure les souvenirs d’un monde désormais perdu sous la neige éternelle."
	case 5:
		lore = "Au bord d’un bassin noir et stagnant, une femme-serpent aux cheveux mouillés vous observe. Son regard mielleux vous attire, mais sous la surface, son corps serpentin se tend, prêt à broyer les imprudents."
	case 6:
		lore = "Dans une salle plongée dans l’ombre, une femme au sourire doux incline la tête… avant que son cou s’étire à l’infini. Sa voix résonne depuis mille directions, et chaque mot siffle comme une menace prête à se resserrer autour de vous."
	case 7:
		lore = "Parmi les ruines d’une ville engloutie, un squelette colossal émerge des ténèbres, formé des os des affamés. Chaque pas ébranle le sol, et ses mâchoires claquent dans un silence oppressant qui précède la mort."
	case 8:
		lore = "Au seuil d’un tribunal brisé, Enma, le juge des enfers, vous attend. Sa voix profonde énumère vos fautes, et sa balance fissurée menace de vous condamner au néant avant même que vous ne puissiez lever votre lame."
	case 9:
		lore = "Dans une caverne où coulent des rivières de lave et de sang, surgit un serpent à huit têtes. Chaque gueule crache une colère ancienne, et ses yeux multiples reflètent les flammes d’un monde sur le point de disparaître."
	case 10:
		lore = "Au sommet du Soleil Rouge, le Shinigami vous accueille dans un silence glacé. Sa faux scintille comme une étoile mourante, et ses orbites vides vous jugent : êtes-vous le sauveur du monde ou son dernier fossoyeur ?"
	}
	printLore(lore)
}

// Afficher l’inventaire
func accessInventory(c Character) {
	clearScreen()
	fmt.Println("\n=== 🎒 Inventaire ===")
	if len(c.inventaire) == 0 {
		fmt.Println("Inventaire vide !")
	} else {
		for i, item := range c.inventaire {
			fmt.Printf("%d. %s\n", i+1, item)
		}
	}
	fmt.Println("=================")
	fmt.Print("\nAppuyez sur Entrée pour continuer...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

// Utiliser un objet
func useItem(c *Character) {
	clearScreen()
	if len(c.inventaire) == 0 {
		fmt.Println("Votre inventaire est vide !")
		fmt.Print("\nAppuyez sur Entrée pour continuer...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}
	fmt.Println("\nQuel objet voulez-vous utiliser ?")
	for i, item := range c.inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	var choix int
	fmt.Print("Choix : ")
	fmt.Scanln(&choix)
	if choix < 1 || choix > len(c.inventaire) {
		fmt.Println("Choix invalide.")
		fmt.Print("\nAppuyez sur Entrée pour continuer...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}
	objet := c.inventaire[choix-1]
	switch objet {
	case "Potion de vie":
		c.hpActuels += 20
		if c.hpActuels > c.hpMax {
			c.hpActuels = c.hpMax
		}
		fmt.Println("Vous avez utilisé une Potion de vie ! +20 PV.")
	case "Thé du Kitsune (狐茶)":
		c.hpActuels += 8
		c.armure += 3
		fmt.Println("Vous avez utilisé Thé du Kitsune ! +8 PV, -3 dégâts reçus.")
	case "Amulette protectrice":
		c.armure += 5
		fmt.Println("Vous avez utilisé une Amulette protectrice ! -5 dégâts reçus.")
	case "Plume de Tengu":
		c.armure += 3
		c.atk += 2
		fmt.Println("Vous avez utilisé une Plume de Tengu ! -3 dégâts reçus, +2 ATK.")
	case "Potion de force":
		c.atk += 5
		c.armure += 2
		c.hpActuels += 7
		if c.hpActuels > c.hpMax {
			c.hpActuels = c.hpMax
		}
		fmt.Println("Vous avez utilisé une Potion de force ! +5 ATK, -2 dégâts reçus, +7 PV.")
	case "Armure sacrée":
		c.atk += 15
		c.armure += 15
		fmt.Println("Vous avez utilisé une Armure sacrée ! +15 ATK, -15 dégâts reçus.")
	default:
		fmt.Println("Cet objet ne peut pas être utilisé pour l’instant.")
		fmt.Print("\nAppuyez sur Entrée pour continuer...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}
	// Retirer l’objet après utilisation
	c.inventaire = append(c.inventaire[:choix-1], c.inventaire[choix:]...)
	fmt.Print("\nAppuyez sur Entrée pour continuer...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

// Marchand itinérant
func visitMerchant(c *Character) {
	for {
		clearScreen()
		fmt.Println("\n=== 🏪 Marchand Itinérant ===")
		fmt.Println("1. 💊 Potion de vie (+20 PV) — 3 pièces")
		fmt.Println("2. 🛡️ Amulette protectrice (-5 dgt reçu) — 5 pièces")
		fmt.Println("3. ☕ Thé du Kitsune (+8 PV, -3 dgt reçu) — 7 pièces")
		fmt.Println("4. 🪶 Plume de Tengu (-3 dgt reçu, +2 ATK) — 9 pièces")
		fmt.Println("5. 🔨 Katana améliorée (+5 ATK) — 10 pièces")
		fmt.Println("6. 🏺 Amulette de longévité (-15 dgt reçu) — 15 pièces")
		fmt.Println("7. 💪 Potion de force (+5 ATK, -2 dgt reçu, +7 PV) — 18 pièces")
		fmt.Println("8. 🛡️ Armure sacrée (+15 ATK, -15 dgt reçu) — 25 pièces")
		fmt.Println("0. Retour")
		fmt.Print("Choix : ")
		var choix int
		fmt.Scanln(&choix)
		if choix == 0 {
			return
		}
		// Logique d'achat
		items := map[int]struct {
			name string
			cost int
			use  func(*Character)
		}{
			1: {"Potion de vie", 3, func(c *Character) { c.inventaire = append(c.inventaire, "Potion de vie") }},
			2: {"Amulette protectrice", 5, func(c *Character) { c.inventaire = append(c.inventaire, "Amulette protectrice") }},
			3: {"Thé du Kitsune", 7, func(c *Character) { c.inventaire = append(c.inventaire, "Thé du Kitsune (狐茶)") }},
			4: {"Plume de Tengu", 9, func(c *Character) { c.inventaire = append(c.inventaire, "Plume de Tengu") }},
			5: {"Katana améliorée", 10, func(c *Character) { c.atk += 5; fmt.Println("Votre ATK augmente de 5 !") }},
			6: {"Amulette de longévité", 15, func(c *Character) { c.armure += 15; fmt.Println("Votre armure augmente de 15 !") }},
			7: {"Potion de force", 18, func(c *Character) { c.inventaire = append(c.inventaire, "Potion de force") }},
			8: {"Armure sacrée", 25, func(c *Character) { c.inventaire = append(c.inventaire, "Armure sacrée") }},
		}
		item, exists := items[choix]
		if !exists {
			fmt.Println("Choix invalide.")
			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			bufio.NewReader(os.Stdin).ReadString('\n')
			continue
		}
		if c.or < item.cost {
			fmt.Println("Pas assez d’or !")
			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			bufio.NewReader(os.Stdin).ReadString('\n')
			continue
		}
		c.or -= item.cost
		item.use(c)
		fmt.Printf("%s acheté ! Il vous reste %d pièces.\n", item.name, c.or)
		fmt.Print("\nAppuyez sur Entrée pour continuer...")
		bufio.NewReader(os.Stdin).ReadString('\n')
	}
}

// Tour du joueur pendant le combat
func playerTurn(j *Character, m *Monster) bool {
	clearScreen()
	fmt.Println("\n=== ⚔️ Tour du joueur ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Utiliser un objet")
	fmt.Println("3. Visiter le marchand")
	fmt.Print("Choix : ")
	var choix int
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		dmg := j.atk
		m.hpActuels -= dmg
		fmt.Printf("⚔️ %s inflige %d dégâts à %s !\n", j.nom, dmg, m.nom)
	case 2:
		useItem(j)
		return false
	case 3:
		visitMerchant(j)
		return false // Le joueur a choisi de visiter le marchand, on ne continue pas le tour
	default:
		fmt.Println("Choix invalide !")
		fmt.Print("\nAppuyez sur Entrée pour continuer...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return false
	}
	return true
}

// Tour du monstre pendant le combat
func monsterTurn(j *Character, m *Monster) {
	fmt.Println("\n=== Tour du monstre ===")
	dmg := m.atk - j.armure
	if dmg < 1 {
		dmg = 1
	}
	j.hpActuels -= dmg
	fmt.Printf("👹 %s inflige %d dégâts à %s ! PV restant : %d/%d\n", m.nom, dmg, j.nom, j.hpActuels, j.hpMax)
	fmt.Print("\nAppuyez sur Entrée pour continuer...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

// Combat contre un boss
func fightBoss(j *Character, m *Monster) bool {
	for j.hpActuels > 0 && m.hpActuels > 0 {
		if !playerTurn(j, m) {
			continue // Le joueur a visité le marchand, on recommence le tour
		}
		if m.hpActuels <= 0 {
			return true
		}
		monsterTurn(j, m)
		if j.hpActuels <= 0 {
			return false
		}
	}
	return true
}

// Liste des boss
var bossList = []Monster{
	initOni(),
	initKitsune(),
	initTengu(),
	initYukiOnna(),
	initNureOnna(),
	initRokurokubi(),
	initGashadokuro(),
	initEnma(),
	initYamataNoOrochi(),
	initShinigami(),
}

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
		},
		rewardOr:   20,
		rewardItem: "Potion de vie",
	}
}

// Combat d'entraînement complet
func trainingFight(j *Character) {
	monstre := initKappa()
	clearScreen()
	fmt.Println("\n=== ⚔ Combat d'entraînement ===")
	fmt.Printf("Affrontez le %s !\n", monstre.nom)
	for j.hpActuels > 0 && monstre.hpActuels > 0 {
		playerTurn(j, &monstre)
		if monstre.hpActuels <= 0 {
			fmt.Println("Vous avez vaincu le Kappa !")
			j.or += monstre.rewardOr
			j.inventaire = append(j.inventaire, monstre.rewardItem)
			fmt.Printf("Récompense : %d pièces et %s\n", monstre.rewardOr, monstre.rewardItem)
			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			bufio.NewReader(os.Stdin).ReadString('\n')
			break
		}
		monsterTurn(j, &monstre)
		if j.hpActuels <= 0 {
			fmt.Println("Vous êtes mort !")
			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			bufio.NewReader(os.Stdin).ReadString('\n')
			break
		}
	}
}

// Fonction pour jouer le mode histoire
func playGame(j *Character) {
	printLore("\nSous le ciel rouge sang, les ruines du Japon s’étendent à perte de vue...\nUn samouraï solitaire avance, son bras mécanique brillant dans la poussière du crépuscule...\nLa Tour du Soleil Rouge se dresse, promesse de réponses et de dangers…")
	for i := 0; i < 10; i++ {
		boss := bossList[i]
		printBossLore(i + 1) // Affiche l'histoire du boss
		clearScreen()
		fmt.Printf("\n=== Niveau %d : Affronte %s ! ===\n", i+1, boss.nom)
		if !fightBoss(j, &boss) {
			fmt.Println("Game Over !")
			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			bufio.NewReader(os.Stdin).ReadString('\n')
			return
		}
		fmt.Println("Vous avez vaincu le boss !")
		j.or += boss.rewardOr
		j.inventaire = append(j.inventaire, boss.rewardItem)
		fmt.Printf("Récompense : %d pièces et %s\n", boss.rewardOr, boss.rewardItem)
		j.experience += 50
		if j.experience >= j.levelUpExp {
			j.niveau++
			j.hpMax += 20
			j.hpActuels = j.hpMax
			j.atk += 5
			j.levelUpExp += 100
			fmt.Printf("🌟 Niveau %d atteint ! Vos statistiques ont augmenté !\n", j.niveau)
			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			bufio.NewReader(os.Stdin).ReadString('\n')
		}
		if i < 9 {
			visitMerchant(j)
		}
	}
	fmt.Println("Félicitations, vous avez terminé le jeu !")
	fmt.Print("\nAppuyez sur Entrée pour continuer...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func main() {
	// Création du personnage
	joueur := initCharacter(
		"Samouraï (侍)",
		"Guerrier",
		1,
		120,
		120,
		10,
		[]string{"Katana (刀)", "Potion de vie"},
		100,
	)
	// Boucle du menu principal
	for {
		clearScreen()
		fmt.Println(`⠀⠀⠘⡇⠀⠀⠀⢠⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢷⠀⢠⢣⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢘⣷⢸⣾⣇⣶⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣿⣿⣿⣹⣿⣿⣷⣿⣆⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢼⡇⣿⣿⣽⣶⣶⣯⣭⣷⣶⣿⣿⣶⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠸⠣⢿⣿⣿⣿⣿⡿⣛⣭⣭⣭⡙⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠸⣿⣿⣿⣿⣿⠿⠿⠿⢯⡛⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣾⣿⡿⡷⢿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⡔⣺⣿⣿⣽⡿⣿⣿⣿⣟⡳⠦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢠⣭⣾⣿⠃⣿⡇⣿⣿⡷⢾⣭⡓⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣾⣿⡿⠷⣿⣿⡇⣿⣿⣟⣻⠶⣭⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣋⣵⣞⣭⣮⢿⣧⣝⣛⡛⠿⢿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⣀⣀⣠⣶⣿⣿⣿⣿⡿⠟⣼⣿⡿⣟⣿⡇⠀⠙⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⡼⣿⣿⣿⢟⣿⣿⣿⣷⡿⠿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠉⠁⠀⢉⣭⣭⣽⣯⣿⣿⢿⣫⣮⣅⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢀⣿⣟⣽⣿⣿⣿⣿⣾⣿⣿⣯⡛⠻⢷⣶⣤⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢀⡞⣾⣿⣿⣿⣿⡟⣿⣿⣽⣿⣿⡿⠀⠀⠀⠈⠙⠛⠿⣶⣤⣄⡀⠀⠀⠀
⠀⠀⠀⣾⣸⣿⣿⣷⣿⣿⢧⣿⣿⣿⣿⣿⣷⠁⠀⠀⠀⠀⠀⠀⠀⠈⠙⠻⢷⣦
⠀⠀⠀⡿⣛⣛⣛⣛⣿⣿⣸⣿⣿⣿⣻⣿⣿⠆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢸⡇⣿⣿⣿⣿⣿⡏⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢰⣶⣶⣶⣾⣿⢃⣿⣿⣿⣿⣯⣿⣭⠁⠀⠀`)
		fmt.Println("\n===== THE SHADOW OF THE RED SUN =====")
		fmt.Println("🎮 1. Jouer")
		fmt.Println("⚔  2. Lancer un combat d’entraînement")
		fmt.Println("🧮 3. Voir les statistiques")
		fmt.Println("❌ 4. Quitter le jeu")
		fmt.Print("Choix : ")
		var choix int
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			playGame(&joueur)
		case 2:
			trainingFight(&joueur)
		case 3:
			displayInfo(joueur)
		case 4:
			clearScreen()
			fmt.Println("👋 Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			bufio.NewReader(os.Stdin).ReadString('\n')
		}
	}
}
