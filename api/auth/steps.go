package auth

import (
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/samber/lo"
)

type AuthStep struct {
	ID           string
	CanGoBack    bool
	PreviousStep string
}

type FormField struct {
	Name string
	Type string
}

type FormStep struct {
	AuthStep
	Fields []FormField
}

func NewField(name, type_ string) FormField {
	return FormField{
		Name: name,
		Type: type_,
	}
}

type ChoiceStep struct {
	AuthStep
	Choices []string
}

func (s ChoiceStep) StepInfo() AuthStep {
	return s.AuthStep
}

func (s FormStep) StepInfo() AuthStep {
	return s.AuthStep
}

func (s ChoiceStep) ToProtoV1() *authv1.AuthStep {
	return &authv1.AuthStep{
		CanGoBack: s.CanGoBack,
		Step: &authv1.AuthStep_Choice_{
			Choice: &authv1.AuthStep_Choice{
				Title:   s.ID,
				Options: s.Choices,
			},
		},
	}
}

func (s FormStep) ToProtoV1() *authv1.AuthStep {
	return &authv1.AuthStep{
		CanGoBack: s.CanGoBack,
		Step: &authv1.AuthStep_Form_{
			Form: &authv1.AuthStep_Form{
				Title: s.ID,
				Fields: lo.Map(s.Fields, func(field FormField, _ int) *authv1.AuthStep_Form_FormField {
					return &authv1.AuthStep_Form_FormField{
						Name: field.Name,
						Type: field.Type,
					}
				}),
			},
		},
	}
}

func newAuthStep(id string, previousStep string) AuthStep {
	s := AuthStep{ID: id}
	if previousStep != "" {
		s.CanGoBack = true
		s.PreviousStep = previousStep
	}
	return s
}

func NewForm(id string, fields []FormField, previousStep string) FormStep {
	return FormStep{
		AuthStep: newAuthStep(id, previousStep),
		Fields:   fields,
	}
}

func NewChoice(id string, choices []string, previousStep string) ChoiceStep {
	return ChoiceStep{
		AuthStep: newAuthStep(id, previousStep),
		Choices:  choices,
	}
}
