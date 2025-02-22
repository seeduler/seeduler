package repositories

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/seeduler/seeduler/models"
)

type EventRepository struct {
	FilePath string
}

func NewEventRepository(filePath string) *EventRepository {
	return &EventRepository{FilePath: filePath}
}

func (repo *EventRepository) readJSONFile() ([]models.Event, error) {
	log.Println("Reading JSON file")
	var events []models.Event
	file, err := os.Open(repo.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (repo *EventRepository) GetEvents() (resp []models.Event, err error) {
	log.Println("Getting all events (in repository)")
	events, err := repo.readJSONFile()
	if err != nil {
		return nil, err
	}
	return events, nil
}
