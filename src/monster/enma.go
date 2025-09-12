package monster

import (
	"YAMA/src/character"
)

func InitEnma() Monster {
	return Monster{
		Nom:        "Enma (閻魔)",
		HpMax:      120,
		HpActuels:  120,
		Atk:        15,
		Pattern:    []string{"attack", "judgment"},
		Special:    func(j *character.Character) {},
		RewardOr:   60,
		RewardItem: "Marque d’Enma",
	}
}
