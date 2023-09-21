package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"time"
)

type Duration struct {
	time.Duration
}

func (d *Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}

func (d *Duration) MarshalYAML() (interface{}, error) {
	return yaml.Marshal(d.String())
}

func (d *Duration) UnmarshalYAML(val *yaml.Node) error {
	var value interface{}
	err := val.Decode(&value)
	if err != nil {
		return err
	}

	switch v := value.(type) {
	case int, int64, float64:
		duration, err := time.ParseDuration(fmt.Sprintf("%v", v))
		if err != nil {
			return err
		}
		d.Duration = duration
	case string:
		duration, err := time.ParseDuration(v)
		if err != nil {
			return err
		}
		d.Duration = duration
	default:
		return fmt.Errorf("invalid duration: %v", v)
	}

	return nil
}
