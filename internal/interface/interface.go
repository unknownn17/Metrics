package interface17

import "conn/internal/models"

type Books interface {
	Create(req *models.Create) (int, error)
	GetbyId(req int) (*models.Update, error)
	Get() ([]*models.Update, error)
	Update(req *models.Update) (int, error)
	Delete(req int) error
}