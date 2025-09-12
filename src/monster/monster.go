package monster

import "YAMA/src/character"

type Monster struct {
	Nom        string
	HpMax      int
	HpActuels  int
	Atk        int
	Pattern    []string
	Special    func(*character.Character)
	RewardOr   int
	RewardItem string
}
