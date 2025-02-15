package authHandler

import (
	"errors"
	"fmt"
	"lets-go/libs/bcrypt"
	loglib "lets-go/libs/logLib"
	"lets-go/libs/token"
	userModel "lets-go/models/user"
	localTypes "lets-go/types"
	"reflect"
)

type Model interface {
	Create() error
	Delete() error
	CheckDuplicate() (bool, error)
}

func createModel(model Model) error {
	duplicate, err := model.CheckDuplicate()
	modelName := reflect.ValueOf(model).Type().String()
	var formated string
	if err != nil {
		formated = fmt.Sprintf("an error occured while checking for duplicates of %s", modelName)
		loglib.LogError(formated, err)
		return err
	}

	if duplicate {
		formated = fmt.Sprintf("the %s already exists", modelName)
		err = errors.New(formated)
		loglib.LogError(formated, err)
		return err
	}

	if err := model.Create(); err != nil {
		formated = fmt.Sprintf("an error occured while creating %s", modelName)
		loglib.LogError(formated, err)
		return err
	}

	return nil
}

func validateUser(dataObj *localTypes.LoginRequestData) (*userModel.User, error) {
	user, err := userModel.GetByEmail(dataObj.Email)
	if err != nil {
		loglib.LogError("the e-mail doesn't exist", err)
		return nil, err
	}

	if !bcrypt.CheckPasswordHash(dataObj.Password, user.Password) {
		loglib.LogError("wrong password", nil)
		return nil, errors.New("wrong password")
	}

	return user, nil
}

type Tokens struct {
	accessToken  string
	refreshToken string
}

func generateTokens(userID string, listRoles []string) (*Tokens, error) {
	accessToken, err := token.CreateAccessToken(userID, listRoles)

	if err != nil {
		loglib.LogError("error while creating access token", err)
		return nil, err
	}

	refreshToken, err := token.CreateRefreshToken(userID, listRoles)
	if err != nil {
		loglib.LogError("error while creating refresh token", err)
		return nil, err
	}

	return &Tokens{accessToken, refreshToken}, nil
}
