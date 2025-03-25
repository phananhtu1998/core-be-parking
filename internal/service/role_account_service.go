package service

type (
	IRoleAccount interface{}
)

var (
	localRoleAccount IRoleAccount
)

func RoleAccountItem() IRoleAccount {
	if localRoleAccount == nil {
		panic("implement localRoleAccount not found for interface IRoleAccount")
	}
	return localRoleAccount
}

func InitRoleAccountItem(i IRoleAccount) {
	localRoleAccount = i
}
