package credential

import (
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"time"
)

type Credential struct {
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	ApiKey    string    `json:"api_key" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCredential(d CreateCredentialDTO) *Credential {
	hasher := sha1.New()

	hasher.Write([]byte(fmt.Sprintf("%s:%d", d.Email, time.Now().Unix())))
	apiKey := base32.StdEncoding.EncodeToString(hasher.Sum(nil))

	return &Credential{
		Name:   d.Name,
		Email:  d.Email,
		ApiKey: apiKey,
	}
}

func HashApiKey(apiKey string) string {
	hasher := sha1.New()

	hasher.Write([]byte(apiKey))

	return base32.StdEncoding.EncodeToString(hasher.Sum(nil))
}
