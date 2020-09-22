package internals

import (
	"time"
	"user/users/models"

	"github.com/go-pg/pg"
)

func UpdateUser(db *pg.DB, user *models.Users) error {

	user.CreatedDate = time.Now().UTC()
	if _, err := db.Model(user).Column("bankruptcy_indicator_flag", "company_name", "created_date", "date_of_birth", "first_name", "last_name", "legal_entity_id", "legal_entity_stage", "legal_entity_type").Where("users.id=?", user.LegalEntityID).Update(); err != nil {
		return err
	}

	return nil
}
