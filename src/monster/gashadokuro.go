package monsters

import (
	"YAMA/src/character"
	"YAMA/src/combat"
)

func InitGashadokuro() combat.Monster {
	return combat.Monster{
		Nom:        "Gashadokuro (餓者髑髏)",
		HpMax:      100,
		HpActuels:  100,
		Atk:        12,
		Pattern:    []string{"attack", "boneCrush"},
		Special:    func(j *character.Character) {},
		RewardOr:   50,
		RewardItem: "Os maudit",
	}
}
