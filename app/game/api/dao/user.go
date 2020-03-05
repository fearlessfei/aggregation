package dao

import (
	"context"
	"solo/app/game/api/models"
)

func (d *Dao) User(c context.Context) []models.User {
	var users2 []models.User

	//d.db.Create(&models.User{Name: "feifei123", Age: 20})

	d.db.Find(&users2)
	//fmt.Println(users2)

	d.db.AutoMigrate(&models.User{})

	return users2
}
