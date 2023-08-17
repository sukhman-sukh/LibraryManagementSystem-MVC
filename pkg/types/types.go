package types

type ConfigSet struct {
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
}

type UserDetail struct {
	Id       int
	UserName string
	Hash     string
	Admin    int
}

type ValidateCookie struct {
	SessionID string
	UserId    int
	Admin     int
	Username  string
}

type Data struct {
	UserName     string
	Books        []Books
	RequestBook  []RequestBooks
	AdminRequest []AdminRequest
	IssuedBooks  []IssuedBook
}

type Books struct {
	BookId   string
	BookName string
	Author   string
	Copies   int
}

type RequestBooks struct {
	RequestId string
	BookId    string
	UserId    string
	Status    string
	BookName  string
	UserName  string
}

type AdminRequest struct {
	RequestId string
	UserId    string
	Status    string
	UserName  string
}

type IssuedBook struct {
	RequestId string
	BookId    string
	UserId    string
	Status    string
	BookName  string
	UserName  string
}

type Response struct{
	Message string
}