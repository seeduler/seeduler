package services

import (
	"log"

	"github.com/seeduler/seeduler/models"
	"github.com/seeduler/seeduler/repositories"
)

type HallService struct {
	HallRepository *repositories.HallRepository
}

func NewHallService(hallRepository *repositories.HallRepository) *HallService {
	return &HallService{HallRepository: hallRepository}
}

func (s *HallService) GetAllHalls() ([]models.Hall, error) {
	log.Println("Getting all halls (in service)")
	return s.HallRepository.GetHalls()
}

func (s *HallService) SaveHalls(halls []models.Hall) error {
	log.Println("Saving halls (in service)")
	return s.HallRepository.SaveHalls(halls)
}

func (s *HallService) GetHallByID(id int) (models.Hall, error) {
	log.Println("Getting hall by ID (in service)")
	return s.HallRepository.GetHallByID(id)
}
