package monster

import "YAMA/src/character"

func InitGashadokuro() Monster {
	return Monster{
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
