package interfaces

import "tdd/models"

type Client interface {
	Get(string) ([]models.Repo, error)
	//Get(string) (models.Repo, error)
}
