package internal

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
)

var ErrSignatureNotMatch = errors.New("signature not match")

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Signer

type Signer interface {
	Sign(string) (string, error)
	Compare(data, signature string) error
}

type signer struct {
	key []byte
}

func NewSigner(secret string) Signer {
	return &signer{
		key: []byte(secret),
	}
}

func (s *signer) sign(data string) ([]byte, error) {
	h := hmac.New(sha256.New, s.key)
	if _, err := h.Write([]byte(data)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func (s *signer) Sign(data string) (string, error) {
	signature, err := s.sign(data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signature), nil
}

func (s *signer) Compare(data, signature string) error {
	expected, err := s.sign(data)
	if err != nil {
		return err
	}
	actual, err := hex.DecodeString(signature)
	if err != nil {
		return err
	}
	if subtle.ConstantTimeCompare(expected, actual) == 0 {
		return ErrSignatureNotMatch
	}
	return nil
}
