package monster

import "YAMA/src/character"

func InitTengu() Monster {
	return Monster{
		Nom:        "Tengu (天狗)",
		HpMax:      70,
		HpActuels:  70,
		Atk:        9,
		Pattern:    []string{"attack", "windSlash"},
		Special:    func(j *character.Character) {},
		RewardOr:   35,
		RewardItem: "Plume de Tengu",
	}
}
