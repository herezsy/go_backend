package authparams

type AuthSecret struct {
	Account     string
	AccountType string
	CodeType    string
	Code        string
}

type ResWithToken struct {
	Uid            int64
	Nickname       string
	PrivilegeType  string
	PrivilegeLevel int64
	Token          string
}

type ResWithoutToken struct {
	Uid            int64
	Nickname       string
	PrivilegeType  string
	PrivilegeLevel int64
}
