package algo

import "github.com/trixky/rubik/models"

func Resolve(set models.Set) models.Set {
	return models.Set{
		// HARD CODE
		Instructions: []models.Instruction{
			{Move: models.MOVE_B, Modifier: models.MODIFIER_NOTHING},
			{Move: models.MOVE_D, Modifier: models.MODIFIER_REVERSE},
			{Move: models.MOVE_F, Modifier: models.MODIFIER_DOUBLE},
		},
	}
}
