package context

import (
	"context"
	"errors"
	"go-backend-api/internal/utils/cache"
)

type InfoUserUUID struct {
	UserId      int64
	UserAccount string
}

func GetSubjectUUID(ctx context.Context) (string, error) {
	sUUID, ok := ctx.Value("subject_uuid").(string)
	if !ok {
		return "", errors.New("subject_uuid not found in context")
	}
	return sUUID, nil
}
func GetUserIdFromUUID(ctx context.Context) (int64, error) {
	sUUID, err := GetSubjectUUID(ctx)
	if err != nil {
		return 0, err
	}
	// get inforuser redis from uuid
	var infoUser InfoUserUUID
	if err := cache.GetCache(ctx, sUUID, &infoUser); err != nil {
		return 0, err
	}
	return infoUser.UserId, nil
}
