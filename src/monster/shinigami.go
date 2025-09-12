package monster

import "YAMA/src/character"

func InitShinigami() Monster {
	return Monster{
		Nom:        "Shinigami (死神)",
		HpMax:      130,
		HpActuels:  130,
		Atk:        16,
		Pattern:    []string{"attack", "soulReap"},
		Special:    func(j *character.Character) {},
		RewardOr:   70,
		RewardItem: "Faux spectrale",
	}
}
