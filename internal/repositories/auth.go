package repositories

import (
	"github.com/jmoiron/sqlx"

	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) interfaces.AuthRepository {
	return &AuthRepository{db: db}
}

func (rep *AuthRepository) GetUserByTelegramID(tg_id int) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM "user" WHERE tg_id = $1`
	err := rep.db.Get(&user, query, tg_id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (rep *AuthRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO "user" (created_at, pinged_at, tg_username, tg_id, name, surname)
				VALUES (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $1, $2, $3, $4) RETURNING id`
	err := rep.db.QueryRowx(query, &user.TgUsername, &user.TgID, &user.Name, &user.SurName).StructScan(user)

    return err
}
