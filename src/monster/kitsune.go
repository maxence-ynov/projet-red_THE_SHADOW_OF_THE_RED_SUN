package monsters

import "YAMA/src/combat"

func InitKitsune() combat.Monster {
	return combat.Monster{
		Nom:        "Kitsune (狐)",
		HpMax:      50,
		HpActuels:  50,
		Atk:        7,
		Pattern:    []string{"attack", "illusion"},
		Special:    func(j *combat.Character) {},
		RewardOr:   25,
		RewardItem: "Thé du Kitsune (狐茶)",
	}
}
