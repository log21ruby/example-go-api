package repository

import "example-go-api/model"

// FindOneUserWithUID func
func (repo *Repository) FindOneUserWithUID(uid int, field ...interface{}) (model.User, error) {
	var model model.User
	resp := repo.db.Table("users").Select("id", field...).Limit(1).Find(&model, uid)
	return model, resp.Error
}
