package monster

import "YAMA/src/character"

func InitKappa() Monster {
	return Monster{
		Nom:        "Kappa (河童)",
		HpMax:      40,
		HpActuels:  40,
		Atk:        5,
		Pattern:    []string{"attack", "special"},
		Special:    func(j *character.Character) {},
		RewardOr:   20,
		RewardItem: "Potion de vie",
	}
}
