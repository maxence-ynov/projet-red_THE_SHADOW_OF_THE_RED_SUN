package monster

import "YAMA/src/character"

func InitRokurokubi() Monster {
	return Monster{
		Nom:        "Rokurokubi (轆轤首)",
		HpMax:      60,
		HpActuels:  60,
		Atk:        8,
		Pattern:    []string{"attack", "neckStrike"},
		Special:    func(j *character.Character) {},
		RewardOr:   32,
		RewardItem: "Amulette étrange",
	}
}
