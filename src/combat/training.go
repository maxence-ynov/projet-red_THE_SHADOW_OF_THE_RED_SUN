package combat

import (
	monsters "YAMA/src/monster"
	"fmt"
	"src/character"
)

func TrainingFight(j *character.Character) {
	monstre := monsters.InitKappa()
	fmt.Println("\n Combat d'entraînement ! Affrontez le Kappa d'entraînement !")

	for j.HpActuels > 0 && monstre.HpActuels > 0 {
		PlayerTurn(j, &monstre)
		if monstre.HpActuels <= 0 {
			fmt.Println("Vous avez vaincu le Kappa !")
			break
		}
		MonsterTurn(j, &monstre)
		if j.HpActuels <= 0 {
			fmt.Println("Vous êtes mort !")
			break
		}
	}
}
