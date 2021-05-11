package account

import (
	"fmt"
	"strings"
)

type KeyType uint16

var nameMap = map[string]string{}
var keyTypes = map[string]KeyType{}
var keyNames = map[KeyType]string{}
var keyFactory = map[KeyType]func(t KeyType) Key{}

func (t KeyType) Register(name string, f func(t KeyType) Key) {
	n := strings.ToUpper(name)
	nameMap[n] = name

	keyTypes[n] = t
	keyNames[t] = n
	keyFactory[t] = f
}

func (t KeyType) String() string {
	n := keyNames[t]
	return nameMap[n]
}

func (t KeyType) Create() Key {
	f := keyFactory[t]
	if f != nil {
		return f(t)
	}
	return nil
}

func (t KeyType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *KeyType) UnmarshalText(b []byte) error {
	name := string(b)
	n := strings.ToUpper(name)

	if keyType, ok := keyTypes[n]; ok {
		*t = keyType
		return nil
	}
	return fmt.Errorf("Unknown KeyType: %s", string(b))
}

func GetKeyTypeByName(name string) (KeyType, error) {
	n := strings.ToUpper(name)

	if keyType, ok := keyTypes[n]; ok {
		return keyType, nil
	}
	return 0, fmt.Errorf("Unknown KeyType: %s", name)
}

func GetKeyTypes() []KeyType {
	list := make([]KeyType, 0)
	for _, t := range keyTypes {
		list = append(list, t)
	}
	return list
}
