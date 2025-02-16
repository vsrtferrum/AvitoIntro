package transform

import (
	"encoding/json"

	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
)

func MarshalAllTransactionInfo(info *model.InfoResponse, name string) (*[]byte, error) {
	if info == nil {
		return nil, errors.ErrNoUserFound
	}

	allTransactions := make([]model.CoinTransaction, 0)

	if info.CoinHistory.Received != nil {
		for _, received := range *info.CoinHistory.Received {
			transaction := model.CoinTransaction{
				Amount:   received.Amount,
				FromUser: received.FromUser,
				ToUser:   name,
			}
			allTransactions = append(allTransactions, transaction)

		}
	}

	if info.CoinHistory.Sent != nil {
		for _, sent := range *info.CoinHistory.Sent {
			transaction := model.CoinTransaction{
				Amount:   sent.Amount,
				FromUser: name,
				ToUser:   sent.ToUser,
			}
			allTransactions = append(allTransactions, transaction)

		}
	}

	if len(allTransactions) == 0 {
		return nil, errors.ErrNoUserFound
	}

	data, err := json.Marshal(allTransactions)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func UnmarshalInfoResponse(data *[]byte) (*model.InfoResponse, error) {
	if len(*data) == 0 {
		return nil, errors.ErrJsonUnMarshall
	}

	var infoResponse model.InfoResponse
	err := json.Unmarshal(*data, &infoResponse)
	if err != nil {
		return nil, err
	}

	return &infoResponse, nil
}
