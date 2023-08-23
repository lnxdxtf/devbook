package models_user

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Stage int

const (
	Signup Stage = iota
	Login
	Reset
	Update
)

type User struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Pswrd     string    `json:"pswrd,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) validate(stage Stage) []error {
	var errorsArr []error

	if user.Name == "" {
		errorsArr = append(errorsArr, errors.New("name is required"))
	}

	if user.Nick == "" {
		errorsArr = append(errorsArr, errors.New("nick is required"))
	}

	if user.Email == "" {
		errorsArr = append(errorsArr, errors.New("email is required"))
	} else {
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			errorsArr = append(errorsArr, errors.New("email is invalid"))
		}
	}

	if stage == Signup && user.Pswrd == "" {
		errorsArr = append(errorsArr, errors.New("pswrd is required"))
	}

	if len(errorsArr) == 0 {
		return nil
	}

	return errorsArr
}

func (user *User) format(stage Stage) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	if stage == Signup {
		hash_pwrd, err := security.HashGen(user.Pswrd)
		if err != nil {
			return err
		}
		user.Pswrd = string(hash_pwrd)
	}
	return nil
}

func (user *User) Prepare(stage Stage) []error {
	errorsArr := []error{}
	if errs := user.validate(stage); errs != nil {
		errorsArr = errs
	}
	err := user.format(stage)
	if err != nil {
		errorsArr = append(errorsArr, err)
	}

	if len(errorsArr) == 0 {
		return nil
	}
	return errorsArr
}
