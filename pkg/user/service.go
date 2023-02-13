package user

import (
	"context"
	"regexp"
	"sync"

	"gitlab.com/mlc-d/table/pkg/dto"
	"gitlab.com/mlc-d/table/pkg/user/internal"
	"golang.org/x/crypto/bcrypt"
)

var (
	s         *service
	singleton sync.Once
)

const (
	DevRole = iota
	AdminRole
	ModRole
	AnonRole

	allowedNameCharacters = `[\p{L}\p{N}]`
	MinimumNameLength     = 3
	MaximumNameLength     = 25
)

type service struct {
	repo internal.Repo
}

type Service interface {
	Signup(context.Context, *dto.User) (int64, error)
	Login(ctx context.Context, user *dto.User) (*dto.User, error)
}

func GetService() Service {
	singleton.Do(func() {
		s = &service{repo: internal.GetRepo()}
	})
	return s
}

func (s *service) Signup(ctx context.Context, user *dto.User) (int64, error) {
	var err error

	if ok := s.validateName(user.Name); !ok {
		return 0, err
	}

	user.Password, err = s.saltPassword(user.Password)
	if err != nil {
		return 0, err
	}

	u := &internal.User{
		Name:     user.Name,
		Password: user.Password,
		RoleID:   AnonRole,
	}
	return s.repo.Insert(ctx, u)
}

func (s *service) Login(ctx context.Context, user *dto.User) (*dto.User, error) {
	userFromDB, err := s.repo.GetPasswordByName(ctx, &internal.User{Name: user.Name})
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}
	user.ID = userFromDB.ID
	user.Password = ""
	user.RoleID = userFromDB.RoleID
	return user, nil
}

func (s *service) validateName(name string) bool {
	runeArray := []rune(name)
	nameLength := len(runeArray)

	for i := 0; i < len(runeArray); i++ {
		if ok, _ := regexp.MatchString(allowedNameCharacters, string(runeArray[i])); !ok {
			return false
		}
	}

	if nameLength < MinimumNameLength || nameLength > MaximumNameLength {
		return false
	}
	return true
}

func (s *service) saltPassword(password string) (string, error) {
	salted, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return "", err
	}
	return string(salted), nil
}
