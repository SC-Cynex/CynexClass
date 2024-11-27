package repository

import (
    "database/sql"
    "github.com/SC-Cynex/cynex-class-service/internal/models"
    "log"
    "errors"
)

type InscriptionRepository struct {
    DB *sql.DB
}

func NewInscriptionRepository(db *sql.DB) *InscriptionRepository {
    return &InscriptionRepository{DB: db}
}

func (r *InscriptionRepository) CreateInscription(inscription *models.Inscription) error {
    query := `INSERT INTO TBLInscription (user_id, id_class) VALUES ($1, $2)`
    _, err := r.DB.Exec(query, inscription.UserID, inscription.ClassID)
    if err != nil {
        log.Printf("Error inserting inscription: %v", err)
        return err
    }
    return nil
}

func (r *InscriptionRepository) GetInscriptionsByUserID(userID int) ([]models.Inscription, error) {
    query := `SELECT user_id, id_class FROM TBLInscription WHERE user_id = $1`
    rows, err := r.DB.Query(query, userID)
    if err != nil {
        log.Printf("Error fetching inscriptions for user_id %d: %v", userID, err)
        return nil, err
    }
    defer rows.Close()

    var inscriptions []models.Inscription
    for rows.Next() {
        var inscription models.Inscription
        if err := rows.Scan(&inscription.UserID, &inscription.ClassID); err != nil {
            log.Printf("Error scanning row: %v", err)
            return nil, err
        }
        inscriptions = append(inscriptions, inscription)
    }

    if err := rows.Err(); err != nil {
        log.Printf("Error iterating rows: %v", err)
        return nil, err
    }

    if len(inscriptions) == 0 {
        return nil, errors.New("no inscriptions found")
    }

    return inscriptions, nil
}

func (r *InscriptionRepository) GetAllInscriptions() ([]models.Inscription, error) {
	var inscriptions []models.Inscription
	query := "SELECT user_id, id_class FROM TBLInscription"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var inscription models.Inscription
		if err := rows.Scan(&inscription.UserID, &inscription.ClassID); err != nil {
			return nil, err
		}
		inscriptions = append(inscriptions, inscription)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return inscriptions, nil
}