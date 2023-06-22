package users

import (
	"SatelitePacheco/database"
	"SatelitePacheco/models"
)
/*GetUsers Trae todos los users de la BD*/
func GetUsers() ([]models.Users, error) {

	var users []models.Users
	result := database.PostgresDB.Db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
/*CreateUser crea un nuevo usuario en la BD*/
func CreateUser(user *models.Users) error {

	result := database.PostgresDB.Db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
