package api

import hrpcv1 "github.com/harmony-development/legato/gen/protocol"

var (
	ErrorBadPassword         = NewError("h.bad-password", "invalid credentials")
	ErrorBadAuthId           = NewError("h.bad-auth-id", "invalid auth id")
	ErrorNoStepAction        = NewError("h.no-step-action", "no step action provided")
	ErrorStepMismatch        = NewError("h.step-mismatch", "step mismatch")
	ErrorNoStepBack          = NewError("h.no-step-back", "no step back")
	ErrorMissingForm         = NewError("h.missing-form", "missing form")
	ErrorInvalidForm         = NewError("h.invalid-form", "invalid form")
	ErrorInvalidChoice       = NewError("h.invalid-choice", "invalid choice")
	ErrorEmailAlreadyUsed    = NewError("h.email-already-used", "there is already an account with this email address")
	ErrorUsernameAlreadyUsed = NewError("h.username-already-used", "there is already an account with this username")
	ErrorInvalidAuth         = NewError("h.invalid-auth", "the authentication token was missing or invalid")

	// guilds
	ErrorGuildNotFound     = NewError("h.guild-not-found", "guild not found")
	ErrorUserNotInGuild    = NewError("h.user-not-in-guild", "the requested user is not a member of the guild")
	ErrorUserBanned        = NewError("h.user-banned", "the requested user is banned from this guild")
	ErrorUserAlreadyBanned = NewError("h.user-already-banned", "the requested user is already banned from this guild")
	ErrorUserNotBanned     = NewError("h.user-not-banned", "the requested user is not banned from this guild")
	ErrorAlreadyGuildOwner = NewError("h.already-guild-owner", "the requested user is already an owner of the guild")
	ErrorNotGuildOwner     = NewError("h.not-guild-owner", "you must be an owner of the guild to perform this action")
	ErrorLastGuildOwner    = NewError("h.last-guild-owner", "you may not perform this action because you are the last owner of the guild")

	// channels
	ErrorChannelNotFound = NewError("h.channel-not-found", "channel not found")

	// general errors
	ErrorUserNotFound         = NewError("h.user-not-found", "user not found")
	ErrorInvalidUserForAction = NewError("h.invalid-user-for-action", "cannot perform action on this user")
	ErrorNothingToUpdate      = NewError("h.nothing-to-update", "no values to update")
	ErrorInternalServerError  = NewError("h.internal-error", "internal server error")
)

type Error hrpcv1.Error

func (e *Error) Error() string {
	return e.HumanMessage
}

func NewError(code string, humanMessage string) error {
	return &Error{
		Identifier:   code,
		HumanMessage: humanMessage,
	}
}
