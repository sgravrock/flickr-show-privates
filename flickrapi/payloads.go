package flickrapi

import (
	"errors"
)

type PhotoListEntry struct {
	Data map[string]interface{}
}

func (e *PhotoListEntry) Id() (string, error) {
	id, ok := e.Data["id"].(string)
	if !ok {
		return "", errors.New("Unexpected API result format (no id in photo list entry)")
	}

	return id, nil
}

func (e *PhotoListEntry) IsPublic() (bool, error) {
	return e.boolField("ispublic")
}

func (e *PhotoListEntry) IsFriend() (bool, error) {
	return e.boolField("isfriend")
}

func (e *PhotoListEntry) IsFamily() (bool, error) {
	return e.boolField("isfamily")
}

func (e *PhotoListEntry) boolField(name string) (bool, error) {
	n, err := requireNum(e.Data, []string{name})
	if err != nil {
		return false, err
	}

	return n != 0, nil
}