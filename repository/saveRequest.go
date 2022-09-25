package repository

import "github.com/SBEKzy/testTask/model"

func (r *repo) SaveRequest(request, response string) error {
	var saveModel model.Save

	saveModel.Request = request
	saveModel.Response = response

	err := r.DB.Create(&saveModel).Error

	return err

}
