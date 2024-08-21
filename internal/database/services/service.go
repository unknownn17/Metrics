package services

import (
	interface17 "conn/internal/interface"
	"conn/internal/models"
)

type Services struct {
	I interface17.Books
}

func (u *Services) Create(req *models.Create) (int, error) {
	return u.I.Create(req)
}

func (u *Services) GetbyId(req int) (*models.Update, error) {
	return u.I.GetbyId(req)
}

func (u *Services) Get() ([]*models.Update, error) {
	return u.I.Get()
}

func (u *Services) Update(req *models.Update) (int, error) {
	return u.I.Update(req)
}

func (u *Services) Delete(req int) error {
	return u.I.Delete(req)
}
