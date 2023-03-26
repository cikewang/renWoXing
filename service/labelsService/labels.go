package labelsService

import "renWoXing/models"

type Labels struct {
	lab *models.Labels
}

func (l *Labels) Get() ([]models.Labels, error) {

	admin, err := models.GetLabels()
	if err != nil {
		return nil, err
	}

	return admin, nil
}
