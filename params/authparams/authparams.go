package authparams

type Params struct {
	Account     string
	AccountType string
	CodeType    string
	Code        string

	Uid            int64
	Nickname       string
	PrivilegeType  string
	PrivilegeLevel int64
	Token          string
	Process        map[string]string
	Message        string
}
