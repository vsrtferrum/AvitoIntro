package module

type ModuleActions interface {
	GetInfo(auth string) (*[][]byte, error)
	SendMoney(auth string, toUser string, amount uint64) error
	Buy(auth string, itemName string) error
	Auth(name, passwordHash string) (string, error)
}
