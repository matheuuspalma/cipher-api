package adapters

import (
	"fmt"
	"net/http"
)

type Adapter struct {
	Body   any
	Header any
}

func (Adapter) RequestAdapt(r *http.Request) error {

	if r.Method != "POST" {
		return fmt.Errorf("Wrong method")
	}

	return nil
}
