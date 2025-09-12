package monster

import "YAMA/src/character"

func InitKitsune() Monster {
	return Monster{
		Nom:        "Kitsune (狐)",
		HpMax:      50,
		HpActuels:  50,
		Atk:        7,
		Pattern:    []string{"attack", "illusion"},
		Special:    func(j *character.Character) {},
		RewardOr:   25,
		RewardItem: "Thé du Kitsune (狐茶)",
	}
}
