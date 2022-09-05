package token

import (
	"time"
)

//maker is an interface that manages tokens
type Maker interface {
	//createtoken creates a ner token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//checks if input token is valid or not
	VerifyToken(token string) (*Payload, error)
}
