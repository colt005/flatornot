package repo

import (
	"fmt"
	"os"

	"github.com/colt005/flatornot/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New() (*Repo, error) {

	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	dsn = fmt.Sprintf(dsn, os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Repo{
		db: db,
	}, nil
}

func (r *Repo) GetTotalVotes() ([]models.VoteTotalResult, error) {

	var results []models.VoteTotalResult

	err := r.db.Raw("SELECT vote, COUNT(vote) AS total FROM votes GROUP BY vote").Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil

}

func (r *Repo) GetRecent5Votes() ([]models.VoteRow, error) {

	var results []models.VoteRow

	err := r.db.Raw("SELECT * from votes order by created_at desc limit 5").Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil

}
