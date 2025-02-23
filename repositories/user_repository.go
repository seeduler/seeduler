package repositories

import (
    "encoding/json"
    "io"
    "log"
    "os"

    "github.com/seeduler/seeduler/models"
)

type UserRepository struct {
    FilePath string
}

func NewUserRepository(filePath string) *UserRepository {
    return &UserRepository{FilePath: filePath}
}

func (repo *UserRepository) readJSONFile() ([]models.User, error) {
    log.Println("Reading JSON file")
    var users []models.User
    file, err := os.Open(repo.FilePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    bytes, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(bytes, &users)
    if err != nil {
        return nil, err
    }

    return users, nil
}

func (repo *UserRepository) GetUsers() ([]models.User, error) {
    log.Println("Getting all users (in repository)")
    return repo.readJSONFile()
}

func (repo *UserRepository) SaveUsers(users []models.User) error {
    log.Println("Saving users to JSON file")
    file, err := os.Create(repo.FilePath)
    if err != nil {
        return err
    }
    defer file.Close()

    bytes, err := json.Marshal(users)
    if err != nil {
        return err
    }

    _, err = file.Write(bytes)
    return err
}