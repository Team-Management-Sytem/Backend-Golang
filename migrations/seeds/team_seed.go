package seeds

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"gorm.io/gorm"
)

func ListTeamSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/teams.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listTeam []entity.Team
	if err := json.Unmarshal(jsonData, &listTeam); err != nil {
		return err
	}

	if !db.Migrator().HasTable(&entity.Team{}) {
		if err := db.Migrator().CreateTable(&entity.Team{}); err != nil {
			return err
		}
	}

	for _, team := range listTeam {
		if err := db.Create(&team).Error; err != nil {
			return err
		}
	}

	return nil
}
