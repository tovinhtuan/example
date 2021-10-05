package repositories

import (
	"context"
	"ex1/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface{
	ReadTokenByToken(ctx context.Context, token string)(*models.AuthToken, error)
	ReadUserByToken(ctx context.Context, token string)(*models.User, error)
}
type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (Repository, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=public port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}
	err1 := db.AutoMigrate(
		&models.User{},
		&models.AuthToken{},
	)
	if err1 != nil {
		return nil, err
	}
	return &dbmanager{db.Debug()}, nil
}
func (m *dbmanager) ReadTokenByToken(ctx context.Context, token string)(*models.AuthToken, error){
	authen := models.AuthToken{}
	if err := m.Where(&models.AuthToken{Token: token}).First(&authen).Error; err != nil {
		return nil, err
	}
	return &authen, nil
}
func (m *dbmanager) ReadUserByToken(ctx context.Context, token string)(*models.User, error){
	authen, err := m.ReadTokenByToken(ctx,token)
	if err != nil {
		return nil, err
	}
	user := models.User{}
	if err := m.Where(&models.User{Id: authen.UserId}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user,nil
}