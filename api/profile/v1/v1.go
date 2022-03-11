// Package v1 contains the implemention for profile v1 api
package v1

import (
	"github.com/harmony-development/legato/api"
	"github.com/harmony-development/legato/db"
	profilev1 "github.com/harmony-development/legato/gen/profile/v1"
)

type ProfileV1 struct {
	db *db.DB
}

func New(db *db.DB) *ProfileV1 {
	return &ProfileV1{
		db: db,
	}
}

func (v1 *ProfileV1) GetProfile(c *api.LegatoContext, req *profilev1.GetProfileRequest) (*profilev1.GetProfileResponse, error) {
	res := map[uint64]*profilev1.Profile{}
	for _, id := range req.UserId {
		if allowed, err := v1.db.HasSharedGuilds(c, c.UserID, id); err != nil {
			return nil, err
		} else if !allowed {
			continue
		}

		user, err := v1.db.GetUserByID(c, id)
		if err != nil {
			return nil, err
		}

		res[id] = &profilev1.Profile{
			UserName:    user.Username,
			UserAvatar:  user.Avatar,
			AccountKind: profilev1.AccountKind_ACCOUNT_KIND_FULL_UNSPECIFIED,
			UserStatus:  nil,
		}
	}

	return &profilev1.GetProfileResponse{
		Profile: res,
	}, nil
}

func (v1 *ProfileV1) UpdateProfile(_ *api.LegatoContext, _ *profilev1.UpdateProfileRequest) (*profilev1.UpdateProfileResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ProfileV1) UpdateStatus(_ *api.LegatoContext, _ *profilev1.UpdateStatusRequest) (*profilev1.UpdateStatusResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ProfileV1) GetAppData(_ *api.LegatoContext, _ *profilev1.GetAppDataRequest) (*profilev1.GetAppDataResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ProfileV1) SetAppData(_ *api.LegatoContext, _ *profilev1.SetAppDataRequest) (*profilev1.SetAppDataResponse, error) {
	panic("not implemented") // TODO: Implement
}
