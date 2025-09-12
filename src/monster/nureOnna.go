package monsters

import (
	"YAMA/src/character"
	"YAMA/src/combat"
)

func InitNureOnna() combat.Monster {
	return combat.Monster{
		Nom:        "Nure-Onna (濡女)",
		HpMax:      55,
		HpActuels:  55,
		Atk:        7,
		Pattern:    []string{"attack", "waterWhip"},
		Special:    func(j *character.Character) {},
		RewardOr:   28,
		RewardItem: "Perle aquatique",
	}
}
