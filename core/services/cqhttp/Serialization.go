package cqhttp

import (
	"core/models"
	"encoding/json"
	"fmt"
)

func SerializeRespData[T any](data []byte, v *T) error {
	var response struct {
		models.CQUniversalResp
		Data T `json:"data,omitempty"`
	}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return err
	}
	if response.RetCode != 0 && response.RetCode != 1 {
		return fmt.Errorf("%s", response.Wording)
	}

	*v = response.Data
	return nil
}
