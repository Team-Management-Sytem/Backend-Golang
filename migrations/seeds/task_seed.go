package seeds

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"gorm.io/gorm"
)

func ListTaskSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/tasks.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listTask []entity.Task
	if err := json.Unmarshal(jsonData, &listTask); err != nil {
		return err
	}

	if !db.Migrator().HasTable(&entity.Task{}) {
		if err := db.Migrator().CreateTable(&entity.Task{}); err != nil {
			return err
		}
	}

	for _, task := range listTask {
		if err := db.Create(&task).Error; err != nil {
			return err
		}
	}

	return nil
}
