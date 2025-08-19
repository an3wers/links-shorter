package req

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var data T
	err := json.NewDecoder(body).Decode(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}
