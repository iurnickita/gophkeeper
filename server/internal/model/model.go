// Пакет model. Модели данных
package model

import (
	"time"
)

// Unit - модель единицы данных
type Unit struct {
	Key  UnitKey
	Meta UnitMeta
	Data []byte
}

// UnitKey - ключ единицы данных
type UnitKey struct {
	UserID   string
	UnitName string
}

// UnitMeta - метаданные единицы данных
type UnitMeta struct {
	Type       string
	DataSK     string
	UploadedAt time.Time
}
