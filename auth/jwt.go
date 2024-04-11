package auth

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/ko44d/go-api-sample/clock"
	"github.com/ko44d/go-api-sample/entity"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public,pem
var rawPubKey []byte

type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	Store                 Store
	Clocker               clock.Clocker
}

//go:generate go run github.com/matryer/moq -out moq_test.go . Store
type Store interface {
	Save(ctx context.Context, key string, userID entity.UserID) error
	Load(ctx context.Context, key string) (entity.UserID, error)
}

func NewJWTer(s Store) (*JWTer, error) {
	j := &JWTer{Store: s}
	privkey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)

	}
	pubkey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: public key: %w", err)
	}
	j.PrivateKey = privkey
	j.PublicKey = pubkey
	j.Clocker = clock.RealClocker{}
	return j, nil
}

func parse(rawkey []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(rawkey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	return key, nil
}
