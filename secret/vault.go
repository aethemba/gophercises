package secret

import (
	"errors"
	"secret/encrypt"
)

func Memory(encodingKey string) Vault {
	return Vault{
		encodingKey: encodingKey,
		keyValues:   make(map[string]string),
	}
}

type Vault struct {
	encodingKey string
	keyValues   map[string]string
}

func (v *Vault) Get(key string) (string, error) {
	encryptedValue, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for key")
	}

	ret, err := encrypt.Decrypt(v.encodingKey, encryptedValue)

	if err != nil {
		return "", err
	}

	return ret, nil

}

func (v *Vault) Set(key, value string) error {
	encryptedValue, err := encrypt.Encrypt(v.encodingKey, value)
	if err != nil {
		return err
	}
	v.keyValues[key] = encryptedValue
	return nil
}
