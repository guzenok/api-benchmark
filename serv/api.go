// api.go
package serv

import (
	"encoding/json"
	"time"
)

type Batch struct {
	Receiver int64 `json:"receiver"`
	Amount   int64 `json:"amount"`
}

type TransferRequest struct {
	Sender     int64   `json:"sender"`
	CreatedAt  int64   `json:"created_at"`
	PrevHash   string  `json:"prev_hash"`
	SettingsID int64   `json:"settings_id"`
	Batch      []Batch `json:"batch"`
	Sign       string  `json:"sign"`
}

func (t *TransferRequest) ToJSON() []byte {
	b, _ := json.Marshal(t)
	return b
}

func NewTransferRequest() *TransferRequest {
	return &TransferRequest{
		Sender:     1000000,
		CreatedAt:  time.Now().Unix(),
		PrevHash:   "0no7aXXXXX5AfJd/LQpw==",
		SettingsID: 12,
		Batch: []Batch{
			Batch{
				Receiver: 20,
				Amount:   100,
			},
			Batch{
				Receiver: 30,
				Amount:   100,
			},
		},
		Sign: "XEQCIC5XXXXXiU/GAnCrD4XXXXXXXXXXXqaavmIAiBQr1DmtvmlUwLXXXXXrRoQR7ZoTXPCxxA==",
	}
}
