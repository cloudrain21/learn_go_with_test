package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary)Search(key string) (string, error) {
	v, ok := d[key]
	if ok {
		return v, nil
	} else {
		return "", errors.New("not found")
	}
}