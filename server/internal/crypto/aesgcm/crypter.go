package aesgcm

import (
	"github.com/iurnickita/gophkeeper/server/internal/model"
)

/*

type Crypter interface {
	UnitEncrypt(unit model.Unit) (model.Unit, error)
	UnitDecrypt(unit model.Unit) (model.Unit, error)
}

type crypter struct {
	cfg       config.Config
	encryptSK encryptSK
}

func (c crypter) UnitEncrypt(unit model.Unit) (model.Unit, error)

func NewCrypter(cfg config.Config, store store.Store) (Crypter, error) {

	// Получение промежуточных ключей из БД
	encrStrings, err := store.GetEncryptSK()
	if err != nil {
		return nil, err
	}

	// Дешифрование промежуточных ключей мастер-ключом
	var decrStrings []string
	for _, encrString := range encrStrings {
		decrString, err := decrypt(encrString, cfg.MasterSK)
		if err != nil {
			return nil, err
		}
		decrStrings = append(decrStrings, decrString)
	}

	// Создание объекта промежуточных ключей
	encryptSK, err := NewEncryptSK(decrStrings)
	if err != nil {
		return nil, err
	}

	// Создание нового ключа
	key, err := encryptSK.GetActual()
	var newKey string
	if err != nil {
		switch err {
		// Ключи отсутствуют
		case ErrNoEncryptKeys:
			newKey = createNewKey()
		default:
			return nil, err
		}
	} else {
		// Ключ устарел
		if (time.Now().Sub(key.Begin).Hours() / 24) > float64(cfg.NewSKIntervalD) {
			newKey = createNewKey()
		}
	}
	if newKey != "" {
		// Шифрование мастер-ключом
		encrNewKey, err := encrypt(newKey, cfg.MasterSK)
		if err != nil {
			return nil, err
		}
		// Запись в БД
		store.SetEncryptSK(encrNewKey)
	}

	var crypter crypter
	crypter.cfg = cfg
	crypter.encryptSK = encryptSK
	return *crypter, nil
} */

// UnitEncrypt шифрует единицу данных
func UnitEncrypt(unit model.Unit) (model.Unit, error) {
	unit.Meta.DataSK = createNewKey()
	encrString, err := encrypt(string(unit.Data), unit.Meta.DataSK)
	if err != nil {
		return model.Unit{}, err
	}
	unit.Data = []byte(encrString)

	return unit, nil
}

// UnitDecrypt дешифрует единицу данных
func UnitDecrypt(unit model.Unit) (model.Unit, error) {
	decrString, err := decrypt(string(unit.Data), unit.Meta.DataSK)
	if err != nil {
		return model.Unit{}, err
	}
	unit.Meta.DataSK = ""
	unit.Data = []byte(decrString)

	return unit, nil
}
