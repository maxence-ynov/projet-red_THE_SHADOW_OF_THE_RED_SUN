package monster

import "YAMA/src/character"

func InitOni() Monster {
	return Monster{
		Nom:        "Oni (鬼)",
		HpMax:      60,
		HpActuels:  60,
		Atk:        8,
		Pattern:    []string{"attack", "roar"},
		Special:    func(j *character.Character) {},
		RewardOr:   30,
		RewardItem: "Potion de vie",
	}
}
