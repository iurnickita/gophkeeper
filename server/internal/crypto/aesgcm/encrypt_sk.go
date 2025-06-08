package aesgcm

/* import (
	"encoding/json"
	"errors"
	"sort"
	"time"
)

var (
	ErrNoEncryptKeys = errors.New("no encryption keys")
)

// Промежуточный ключ (хранится в БД)
type encryptSK struct {
	keys []Key
}

type Key struct {
	EncryptSK string    `json:"encrypt_sk"`
	Begin     time.Time `json:"begin"`
}

// GetActual возвращает последний ключ
func (sk encryptSK) GetActual() (Key, error) {
	if len(sk.keys) < 1 {
		return Key{}, ErrNoEncryptKeys
	}
	return sk.keys[0], nil
}

// GetOld возвращает архивный ключ для чтения старых записей
func (sk encryptSK) GetOld(uploadedat time.Time) (Key, error) {
	// Поиск ключа по дате
	for _, key := range sk.keys {
		if key.Begin.Before(uploadedat) {
			return key, nil
		}
	}
	return Key{}, ErrNoEncryptKeys
}

// NewEncryptSK принимает набор ключей в формате json
func NewEncryptSK(strings []string) (encryptSK, error) {
	var sk encryptSK

	// Десериализация
	for _, string := range strings {
		var key Key
		err := json.Unmarshal([]byte(string), key)
		if err != nil {
			return encryptSK{}, err
		}
		sk.keys = append(sk.keys, key)
	}

	// Сортировка по дате (по убыванию)
	sort.Slice(sk.keys, func(i, j int) bool {
		return sk.keys[i].Begin.After(sk.keys[j].Begin)
	})

	return sk, nil
}
*/
