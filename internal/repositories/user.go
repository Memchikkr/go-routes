package repositories

import (
	"github.com/jmoiron/sqlx"

	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) interfaces.UserRepository {
	return &UserRepository{db: db}
}

func (rep *UserRepository) GetUserByID(id int) (*models.UserResponse, error) {
	var user models.UserResponse
	query := `SELECT id, name, surname, phone_number, description FROM "user" WHERE id = $1`
	err := rep.db.Get(&user, query, id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (rep *UserRepository) PutUserData(id int, data *models.UserPutRequest) error {
	query := `UPDATE "user" SET phone_number = $1, name = $2, surname = $3, username = $4, description = $5 WHERE id = $6 RETURNING *`
	_, err := rep.db.Exec(query, data.PhoneNumber, data.Name, data.SurName, data.UserName, data.Description, id)
	return err
}
