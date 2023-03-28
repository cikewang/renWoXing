package actionService

import "renWoXing/models"

type Action struct {
	act *models.Action
}

func (a *Action) GetUpdateTime() (models.Action, error) {

	act, err := models.GetUpDateTime()
	if err != nil {
		return act, err
	}

	return act, nil
}
