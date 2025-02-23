package repositories

import (
    "encoding/json"
    "io"
    "log"
    "os"

    "github.com/seeduler/seeduler/models"
)

type HallRepository struct {
    FilePath string
}

func NewHallRepository(filePath string) *HallRepository {
    return &HallRepository{FilePath: filePath}
}

func (repo *HallRepository) readJSONFile() ([]models.Hall, error) {
    log.Println("Reading JSON file")
    var halls []models.Hall
    file, err := os.Open(repo.FilePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    bytes, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(bytes, &halls)
    if err != nil {
        return nil, err
    }

    return halls, nil
}

func (repo *HallRepository) GetHalls() ([]models.Hall, error) {
    log.Println("Getting all halls (in repository)")
    return repo.readJSONFile()
}

func (repo *HallRepository) SaveHalls(halls []models.Hall) error {
    log.Println("Saving halls to JSON file")
    file, err := os.Create(repo.FilePath)
    if err != nil {
        return err
    }
    defer file.Close()

    bytes, err := json.Marshal(halls)
    if err != nil {
        return err
    }

    _, err = file.Write(bytes)
    return err
}