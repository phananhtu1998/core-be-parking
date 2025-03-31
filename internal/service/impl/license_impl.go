package impl

import (
	"context"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/utils/auth"
	"go-backend-api/pkg/response"
	"time"

	"github.com/google/uuid"
)

type sLicense struct {
	r *database.Queries
}

func NewLicenseImpl(r *database.Queries) *sLicense {
	return &sLicense{r: r}
}

func (s *sLicense) CreateLicense(ctx context.Context, in *model.License) (codeResult int, out model.LicenseOutput, err error) {
	// Tạo token không có thời hạn
	license, err := auth.CreateTokenNoExpiration(in.DateStart, in.DateEnd)
	if err != nil {
		return response.ErrCodeLicenseValid, out, err
	}

	// Parse date strings to time.Time
	dateStart, err := time.Parse("2006-01-02 15:04:05", in.DateStart)
	if err != nil {
		return response.ErrCodeLicenseValid, out, err
	}

	dateEnd, err := time.Parse("2006-01-02 15:04:05", in.DateEnd)
	if err != nil {
		return response.ErrCodeLicenseValid, out, err
	}

	// Tạo license trong database
	_, err = s.r.CreateLicense(ctx, database.CreateLicenseParams{
		ID:        uuid.New().String(),
		License:   license,
		RoleID:    in.RoleId,
		DateStart: dateStart,
		DateEnd:   dateEnd,
	})
	if err != nil {
		return response.ErrCodeLicenseValid, out, err
	}
	out.License = license
	return response.ErrCodeSucces, out, nil
}
