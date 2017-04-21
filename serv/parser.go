// parser.go
package serv

import (
	"encoding/json"

	"github.com/buger/jsonparser"
)

type DecodeFunc func([]byte) (*TransferRequest, error)
type EncodeFunc func(interface{}) ([]byte, error)

var (
	DecodesList = map[string]DecodeFunc{
		"Standart": DecodeStandart,
		"Buger":    DecodeBuger,
	}

	EncodesList = map[string]EncodeFunc{
		"Standart": EncodeStandart,
		"Buger":    EncodeBuger,
	}
)

// Standart implementation

func DecodeStandart(data []byte) (tr *TransferRequest, err error) {
	tr = &TransferRequest{}
	err = json.Unmarshal(data, tr)
	return tr, err
}

func EncodeStandart(obj interface{}) ([]byte, error) {
	/*
		d, err := json.Marshal(obj)
		if err != nil {
			log.Printf("EncodeStandart error: %s\n", err)
		}
		return d, err
	*/
	return []byte("{}"), nil
}

// Buger implementation

func DecodeBuger(data []byte) (tr *TransferRequest, err error) {
	tr = &TransferRequest{}
	switch {
	case true:
		tr.Sender, err = jsonparser.GetInt(data, "sender")
		if err != nil {
			break
		}
		tr.CreatedAt, err = jsonparser.GetInt(data, "created_at")
		if err != nil {
			break
		}
		tr.PrevHash, err = jsonparser.GetString(data, "prev_hash")
		if err != nil {
			break
		}
		tr.SettingsID, err = jsonparser.GetInt(data, "settings_id")
		if err != nil {
			break
		}
		tr.Sign, err = jsonparser.GetString(data, "sign")
		if err != nil {
			break
		}
		tr.Batch = []Batch{}
		_, err = jsonparser.ArrayEach(data, func(avalue []byte, dataType jsonparser.ValueType, offset int, err error) {
			b := Batch{}
			b.Amount, err = jsonparser.GetInt(avalue, "amount")
			b.Receiver, err = jsonparser.GetInt(avalue, "receiver")
			tr.Batch = append(tr.Batch, b)
		}, "batch")
		if err != nil {
			break
		}
	}
	return tr, err
}

func EncodeBuger(obj interface{}) ([]byte, error) {
	// buger/jsonparser can't encode
	return EncodeStandart(obj)
}
