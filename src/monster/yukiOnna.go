package monster

import "YAMA/src/character"

func InitYukiOnna() Monster {
	return Monster{
		Nom:        "Yuki-Onna (雪女)",
		HpMax:      65,
		HpActuels:  65,
		Atk:        8,
		Pattern:    []string{"attack", "iceBlast"},
		Special:    func(j *character.Character) {},
		RewardOr:   30,
		RewardItem: "Cristal de glace",
	}
}
