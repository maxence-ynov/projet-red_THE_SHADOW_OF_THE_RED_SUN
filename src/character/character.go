package character

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	HpMax      int
	HpActuels  int
	Inventaire []string
	Or         int
}

func InitCharacter(nom, classe string, niveau, hpMax, hpActuels, or int, inventaire []string) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		HpMax:      hpMax,
		HpActuels:  hpActuels,
		Inventaire: inventaire,
		Or:         or,
	}
}
