package impl

import (
	"context"
	"database/sql"
	"fmt"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/utils/auth"
	"go-backend-api/pkg/response"
	"log"
	"time"

	"github.com/google/uuid"
)

type sLicense struct {
	r   *database.Queries
	qTx *sql.Tx
	db  *sql.DB
}

func NewLicenseImpl(r *database.Queries, qTx *sql.Tx, db *sql.DB) *sLicense {
	return &sLicense{
		r:   r,
		qTx: qTx,
		db:  db,
	}
}

func (s *sLicense) CreateLicense(ctx context.Context, in *model.License) (codeResult int, out model.LicenseOutput, err error) {
	// Khởi tạo transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeMenuErrror, out, fmt.Errorf("failed to begin transaction: %w", err)
	}

	var committed bool
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()
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

	var dateEnd string

	// Kiểm tra thời hạn của license cha
	dateend := ctx.Value("dateend")
	log.Println("dateend:::", dateend)
	dateendStr, ok := dateend.(string)
	if !ok {
		return response.ErrCodeLicenseValid, out, fmt.Errorf("dateend is not a valid string")
	}

	if in.DateEnd == "NO_EXPIRATION" {
		if dateendStr == "NO_EXPIRATION" {
			dateEnd = "NO_EXPIRATION"
		} else {
			return response.ErrCodeLicenseValid, out, fmt.Errorf("Lỗi đinh dạng cho ngày kết thúc")
		}
	} else {
		// Kiểm tra DateEnd có đúng định dạng không
		_, err = time.Parse("2006-01-02 15:04:05", in.DateEnd)
		if err != nil {
			return response.ErrCodeLicenseValid, out, fmt.Errorf("Không đúng định dạng cho ngày kết thúc, YYYY-MM-DD HH:mm:ss hoặc NO_EXPIRATION, %s", in.DateEnd)
		}
		if dateendStr != "NO_EXPIRATION" && in.DateEnd > dateendStr {
			return response.ErrCodeLicenseValid, out, fmt.Errorf("Vui lòng chọn ngày kết thúc cho gói này sớm hơn gói cha")
		}
		dateEnd = in.DateEnd // Lưu dưới dạng string vì DB là VARCHAR(255)
	}

	//  Tạo Id cho license
	var licenseId = uuid.New().String()
	// Tạo license trong database
	_, err = s.r.CreateLicense(ctx, database.CreateLicenseParams{
		ID:        licenseId,
		License:   license,
		RoleID:    in.RoleId,
		DateStart: dateStart,
		DateEnd:   dateEnd, // Giữ nguyên kiểu string
	})
	if err != nil {
		return response.ErrCodeLicenseValid, out, err
	}
	// Cập nhật license_id trong bảng role
	err = s.r.UpdateLicenseByRoleId(ctx, database.UpdateLicenseByRoleIdParams{
		LicenseID: licenseId,
		ID:        in.RoleId,
	})
	if err != nil {
		return response.ErrCodeMenuErrror, out, err
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return response.ErrCodeMenuErrror, out, err
	}
	committed = true
	out.License = license
	return response.ErrCodeSucces, out, nil
}
