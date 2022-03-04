package authv1impl

import (
	"errors"

	"github.com/harmony-development/legato/api/auth"
	. "github.com/harmony-development/legato/api/auth"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
)

type AuthStepType int

const (
	InitialStep AuthStepType = iota
	LoginStep
	RegisterStep
)

var ErrorNoSuchStep = errors.New("no step in enum for given string")

func fromString(s string) (AuthStepType, error) {
	switch s {
	case "initial":
		return InitialStep, nil
	case "login":
		return LoginStep, nil
	case "register":
		return RegisterStep, nil
	default:
		return 0, ErrorNoSuchStep
	}
}

var steps = map[AuthStepType]interface {
	ToProtoV1() *authv1.AuthStep
	StepInfo() auth.AuthStep
}{
	InitialStep: NewChoice("initial", []string{
		"login",
		"register",
	}, ""),
	LoginStep: NewForm("login", []FormField{
		NewField("email", "email"),
		NewField("password", "password"),
	}, "initial"),
	RegisterStep: NewForm("register", []FormField{
		NewField("email", "email"),
		NewField("username", "username"),
		NewField("password", "new-password"),
	}, "initial"),
}
