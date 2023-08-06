package types

type ConfigSet struct {
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
}

type ErrMsg struct{
	Msg string 	`json:"message"`
}

type UserDetail struct{
	Id int
	UserName string 
	Hash string
	Admin int
}

type ValidateCookie struct {
	SessionID string 
	UserId int
	Admin int
	Username string
}

type Data struct {
	UserName string
	Books []map[string]interface{}
	ReqBook []map[string]interface{}
	AdminReq []map[string]interface{}		
}