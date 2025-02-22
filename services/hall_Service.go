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