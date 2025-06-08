// Пакет service. Логика сервиса
package service

import (
	"context"

	"github.com/iurnickita/gophkeeper/server/internal/crypto/aesgcm"
	"github.com/iurnickita/gophkeeper/server/internal/model"
	"github.com/iurnickita/gophkeeper/server/internal/service/config"
	"github.com/iurnickita/gophkeeper/server/internal/store"
	"go.uber.org/zap"
)

// Service интерфейс сервиса
type Service interface {
	List(ctx context.Context, userID int) ([]string, error)
	Read(ctx context.Context, userID int, unitName string) (model.Unit, error)
	Write(ctx context.Context, unit model.Unit) error
	Delete(ctx context.Context, userID int, unitName string) error
}

// service реализация сервиса
type service struct {
	cfg    config.Config
	store  store.Store
	zaplog *zap.Logger
}

// List implements Service.
func (s service) List(ctx context.Context, userID int) ([]string, error) {
	return s.store.List(ctx, userID)
}

// Read читает единицу данных
func (s service) Read(ctx context.Context, userID int, unitName string) (model.Unit, error) {
	s.zaplog.Sugar().Debug("inbound unitname")
	s.zaplog.Sugar().Debug(unitName)

	// Чтение
	unit, err := s.store.Read(ctx, userID, unitName)
	if err != nil {
		s.zaplog.Error(err.Error())
		return model.Unit{}, err
	}
	s.zaplog.Sugar().Debug("read unit")
	s.zaplog.Sugar().Debug(unit)

	// Дешифрование
	decrUnit, err := aesgcm.UnitDecrypt(unit)
	if err != nil {
		return model.Unit{}, err
	}
	s.zaplog.Sugar().Debug("decrypted unit")
	s.zaplog.Sugar().Debug(decrUnit)

	return decrUnit, nil
}

// Write записывает новую единицу данных
func (s service) Write(ctx context.Context, unit model.Unit) error {
	s.zaplog.Sugar().Debug("inbound unit")
	s.zaplog.Sugar().Debug(unit)

	// Шифрование
	encrUnit, err := aesgcm.UnitEncrypt(unit)
	if err != nil {
		return err
	}
	s.zaplog.Sugar().Debug("encrypted unit")
	s.zaplog.Sugar().Debug(encrUnit)

	// Запись
	err = s.store.Write(ctx, encrUnit)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements Service.
func (s service) Delete(ctx context.Context, userID int, unitName string) error {
	panic("unimplemented")
}

// NewService создает объект сервиса
func NewService(cfg config.Config, store store.Store, zaplog *zap.Logger) (Service, error) {
	service := service{
		cfg:    cfg,
		store:  store,
		zaplog: zaplog}

	return &service, nil
}
