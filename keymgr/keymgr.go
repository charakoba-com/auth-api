package keymgr

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/SermoDigital/jose/crypto"
	"github.com/pkg/errors"
)

// RSAKeyManager manages rsa Private Key and Public Key
type RSAKeyManager struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var _mgr *RSAKeyManager // Global Key Manager

// Init initialize global keymanager
func Init(private string, public string) error {
	bytes, err := ioutil.ReadFile(private)
	if err != nil {
		return errors.Wrap(err, `loading PrivateKey from file`)
	}
	rsaPrivate, err := crypto.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		return errors.Wrap(err, `loading PrivateKey`)
	}
	bytes, err = ioutil.ReadFile(public)
	if err != nil {
		return errors.Wrap(err, `loading PublicKey from file`)
	}
	rsaPublic, err := crypto.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		return errors.Wrap(err, `loading PublicKey`)
	}
	_mgr = &RSAKeyManager{
		PrivateKey: rsaPrivate,
		PublicKey:  rsaPublic,
	}
	return nil
}

// PrivateKey returns private key
func PrivateKey() (*rsa.PrivateKey, error) {
	if _mgr == nil {
		return nil, errors.New(`keymanager has not been initialized`)
	}
	return _mgr.PrivateKey, nil
}

// PublicKey returns public key
func PublicKey() (*rsa.PublicKey, error) {
	if _mgr == nil {
		return nil, errors.New(`keymanager has not been initialized`)
	}
	return _mgr.PublicKey, nil
}
