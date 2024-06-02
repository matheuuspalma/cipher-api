package adapters

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Adapter struct {
	Body   any
	Header any
	data   []byte
	Schema string
}

func (a Adapter) RequestAdapt(r *http.Request) error {

	if r.Method != "POST" {
		return fmt.Errorf("wrong method")
	}

	if err := a.validateBody(r); err != nil {
		return err
	}

	return nil
}

func (a Adapter) validateBody(r *http.Request) error {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bodyByte, &a.Body)
}

func (a Adapter) GetSchema() string {
	return a.Schema
}
