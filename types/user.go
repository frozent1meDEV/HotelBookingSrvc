package types

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 8
)

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"EncryptedPassword" json:"-"`
}

func (p UpdateUserParams) ToBSON() bson.M {
	m := bson.M{}
	if len(p.FirstName) > 0 {
		m["firstName"] = p.FirstName
	}
	if len(p.FirstName) > 0 {
		m["lastName"] = p.LastName
	}
	return m
}

func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstNameLen {
		errors["First Name"] = fmt.Sprintf("First name must be at least %d characters",
			minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		errors["Last Name"] = fmt.Sprintf("Last name must be at least %d characters",
			minLastNameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["Password"] = fmt.Sprintf("Password must be at least %d characters",
			minPasswordLen)
	}
	if !isEmailValid(params.Email) {
		errors["Email"] = fmt.Sprintf("Email is invalid")
	}
	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(e)
}

func NewUserFromParams(params *CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}
