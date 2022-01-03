package user

import (
	"clean-api/model"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	Repository interface {
		Store(user *model.User) (*model.User, error)
		FindEmail(string) (*model.User, error)
		Exist(*model.User) (bool, error)
	}
)

type gormRepository struct {
	db *dynamodb.DynamoDB
}

func NewRepository(database *dynamodb.DynamoDB) Repository {
	return &gormRepository{
		db: database,
	}
}

func (m gormRepository) Exist(u *model.User) (bool, error) {

	return true, nil
}

func (m gormRepository) FindEmail(email string) (*model.User, error) {
	return nil, nil
}

func (m gormRepository) Store(user *model.User) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.ID = uuid.NewString()
	user.Password = string(hash)

	return user, nil
}
