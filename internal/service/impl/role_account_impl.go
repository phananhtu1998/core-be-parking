package impl

import "go-backend-api/internal/database"

type sRoleAccount struct {
	r *database.Queries
}

func NewRoleAccountImpl(r *database.Queries) *sRoleAccount {
	return &sRoleAccount{r: r}
}
