package transform

import (
	"encoding/json"

	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
)

func MarshalAllTransactionInfo(info *model.AllTransactionInfo) (*[][]byte, error) {
	var result [][]byte

	if info.BuyedItem != nil {
		data, err := json.Marshal(info.BuyedItem)
		if err != nil {
			return nil, errors.ErrJsonMarshall
		}
		result = append(result, data)
	}

	if info.SendedMoney != nil {
		data, err := json.Marshal(info.SendedMoney)
		if err != nil {
			return nil, errors.ErrJsonMarshall
		}
		result = append(result, data)
	}

	if info.RecievedMoney != nil {
		data, err := json.Marshal(info.RecievedMoney)
		if err != nil {
			return nil, errors.ErrJsonMarshall
		}
		result = append(result, data)
	}

	return &result, nil
}
