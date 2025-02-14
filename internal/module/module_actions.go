package module

type ModuleActions interface {
	ListOfBuyedItems(id uint64) (*[][]byte, error)
	SendedMoneyStat(id uint64) (*[][]byte, error)
	RecievedMoneyStat(id uint64) (*[][]byte, error)
}
