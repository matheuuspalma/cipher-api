package infrastructure

import b32 "encoding/base32"

type Encode32 struct {
}

func (i Encode32) Encode(data string) (string, error) {

	encoded := b32.StdEncoding.EncodeToString([]byte(data))

	return encoded, nil
}
