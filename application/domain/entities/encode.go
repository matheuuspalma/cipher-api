package entities

import "encoding/json"

type Encode struct {
	EncodedData string `json:"encoded_data"`
}

func (e Encode) ToByte() []byte {
	dataByte, _ := json.Marshal(e.EncodedData)
	return dataByte
}
