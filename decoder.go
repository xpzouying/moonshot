package moonshot

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// decodeResponse 解析响应体，根据内容填充v或返回错误
func decodeResponse(r io.Reader, v interface{}) error {
	var raw json.RawMessage
	if err := json.NewDecoder(r).Decode(&raw); err != nil {
		return err
	}

	var errResp struct {
		Error struct {
			Message string `json:"message"`
			Type    string `json:"type"`
		} `json:"error"`
	}
	if err := json.Unmarshal(raw, &errResp); err == nil && errResp.Error.Message != "" {
		return errors.New(errResp.Error.Message)
	}

	if err := json.Unmarshal(raw, v); err != nil {
		return err
	}

	return nil
}
