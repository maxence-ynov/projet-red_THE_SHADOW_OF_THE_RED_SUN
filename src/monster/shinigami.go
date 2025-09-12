package monsters

import "src/combat"

func InitShinigami() combat.Monster {
	return combat.Monster{
		Nom:        "Shinigami (死神)",
		HpMax:      130,
		HpActuels:  130,
		Atk:        16,
		Pattern:    []string{"attack", "soulReap"},
		Special:    func(j *combat.Character) {},
		RewardOr:   70,
		RewardItem: "Faux spectrale",
	}
}
