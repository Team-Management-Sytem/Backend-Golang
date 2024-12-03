package seeds

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"gorm.io/gorm"
)

func ListTaskSeeder(db *gorm.DB) error {
	// Open JSON file
	jsonFile, err := os.Open("./migrations/json/tasks.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// Read JSON file
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	// Unmarshal JSON data into slice of Task
	var listTask []entity.Task
	if err := json.Unmarshal(jsonData, &listTask); err != nil {
		return err
	}

	// Check if table exists
	if !db.Migrator().HasTable(&entity.Task{}) {
		if err := db.Migrator().CreateTable(&entity.Task{}); err != nil {
			return err
		}
	}

	// Insert data into the database
	for _, task := range listTask {
		if err := db.Create(&task).Error; err != nil {
			return err
		}
	}

	return nil
}
