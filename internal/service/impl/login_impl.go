package impl

import (
	"context"
	"fmt"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/utils/crypto"
	"go-backend-api/pkg/response"
)

type sLogin struct {
	r *database.Queries
}

func NewLoginImpl(r *database.Queries) *sLogin {
	return &sLogin{r: r}
}

func (s *sLogin) Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error) {
	accountBase, err := s.r.GetOneAccountInfoAdmin(ctx, in.Email)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	if !crypto.MatchingPassword(accountBase.Password, in.Password, accountBase.Salt) {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("does not match password")
	}
	//subToken := utils.GenerateCliTokenUUID(int(accountBase.ID))
	return 200, out, err
}
