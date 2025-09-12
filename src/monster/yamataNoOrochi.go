package monster

import "YAMA/src/character"

func InitYamataNoOrochi() Monster {
	return Monster{
		Nom:        "Yamata-no-Orochi (八岐大蛇)",
		HpMax:      150,
		HpActuels:  150,
		Atk:        18,
		Pattern:    []string{"attack", "multiStrike"},
		Special:    func(j *character.Character) {},
		RewardOr:   80,
		RewardItem: "Écaille de dragon",
	}
}
