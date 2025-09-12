package monsters

import "src/combat"

func InitYukiOnna() combat.Monster {
	return combat.Monster{
		Nom:        "Yuki-Onna (雪女)",
		HpMax:      65,
		HpActuels:  65,
		Atk:        8,
		Pattern:    []string{"attack", "iceBlast"},
		Special:    func(j *combat.Character) {},
		RewardOr:   30,
		RewardItem: "Cristal de glace",
	}
}
