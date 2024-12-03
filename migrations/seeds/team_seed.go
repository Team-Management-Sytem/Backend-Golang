package seeds

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"gorm.io/gorm"
)

func ListTeamSeeder(db *gorm.DB) error {
	// Open JSON file
	jsonFile, err := os.Open("./migrations/json/teams.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// Read JSON file
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	// Unmarshal JSON data into slice of Team
	var listTeam []entity.Team
	if err := json.Unmarshal(jsonData, &listTeam); err != nil {
		return err
	}

	// Check if table exists
	if !db.Migrator().HasTable(&entity.Team{}) {
		if err := db.Migrator().CreateTable(&entity.Team{}); err != nil {
			return err
		}
	}

	// Insert data into the database
	for _, team := range listTeam {
		if err := db.Create(&team).Error; err != nil {
			return err
		}
	}

	return nil
}
