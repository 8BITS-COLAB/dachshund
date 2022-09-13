package credential

import (
	"log"

	"gorm.io/gorm"
)

type CredentialRepo struct {
	Conn *gorm.DB
}

func NewRepo(conn *gorm.DB) *CredentialRepo {
	return &CredentialRepo{
		Conn: conn,
	}
}

func (repo *CredentialRepo) GetByApiKey(apiKey string) *Credential {
	var c Credential

	apiKey = HashApiKey(apiKey)

	repo.Conn.Find(&c, "api_key = ?", apiKey)

	if c.ApiKey == "" {
		return nil
	}

	return &c
}

func (repo *CredentialRepo) Create(c Credential) error {
	log.Printf("api_key: %s", c.ApiKey)

	c.ApiKey = HashApiKey(c.ApiKey)

	if err := repo.Conn.Create(&c).Error; err != nil {
		return err
	}

	log.Printf("api_key_hash: %s", c.ApiKey)

	return nil
}
