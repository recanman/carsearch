// made by recanman
package models

import (
	"encoding/json"
	"os"
)

var Makes map[string]string

type DataJson struct {
	Makes map[string]string `json:"makes"`
}

func Initialize() error {
	file, err := os.Open("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	data := DataJson{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return err
	}

	Makes = data.Makes
	return nil
}
