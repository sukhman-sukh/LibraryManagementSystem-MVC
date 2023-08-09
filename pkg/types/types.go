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
	Books  []Books
	ReqBook []ReqBooks
	AdminReq []AdminReq	
}

type Books struct {
	BookId string
	BookName string
	Author string
	Copies int
}

type ReqBooks struct {
	ReqId string
	Date string
	BookId string
	UserId string
	Status string
}


type AdminReq struct {
	ReqId string
	UserId string
	Status string
}