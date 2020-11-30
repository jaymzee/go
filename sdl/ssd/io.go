package ssd

import (
	"encoding/json"
	"io/ioutil"
)

// Open returns a new Display configured from a file
func Open(name string) (*Display, error) {
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	disp := &Display{}
	err = json.Unmarshal(bytes, disp)
	if err != nil {
		return nil, err
	}
	return disp, nil
}

// Write saves display settings to file
func (d *Display) Write(name string) error {
	bytes, err := json.Marshal(d)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(name, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
