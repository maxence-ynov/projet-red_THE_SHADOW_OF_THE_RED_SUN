package monster

import "YAMA/src/character"

func InitNureOnna() Monster {
	return Monster{
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
