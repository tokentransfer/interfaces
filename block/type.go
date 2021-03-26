package block

import (
	"fmt"
)

type TransactionType uint32

var txTypes = map[string]TransactionType{}
var txNames = map[TransactionType]string{}
var txFactory = map[TransactionType]func(t TransactionType) Transaction{}

func (t TransactionType) Register(name string, f func(t TransactionType) Transaction) {
	txTypes[name] = t
	txNames[t] = name
	txFactory[t] = f
}

func (t TransactionType) String() string {
	return txNames[t]
}

func (t TransactionType) Create() Transaction {
	f := txFactory[t]
	return f(t)
}

func (t TransactionType) MarshalText() ([]byte, error) {
	return []byte(txNames[t]), nil
}

func (t *TransactionType) UnmarshalText(b []byte) error {
	if txType, ok := txTypes[string(b)]; ok {
		*t = txType
		return nil
	}
	// If here, add tx type to TxFactory and TxTypes in factory.go
	return fmt.Errorf("Unknown TransactionType: %s", string(b))
}

func GetTransactionTypeByName(name string) TransactionType {
	return txTypes[name]
}

func GetTransactionTypes() []TransactionType {
	list := make([]TransactionType, 0)
	for _, t := range txTypes {
		list = append(list, t)
	}
	return list
}

type StateType uint32

var stateTypes = map[string]StateType{}
var stateNames = map[StateType]string{}
var stateFactory = map[StateType]func(t StateType) State{}

func (t StateType) Register(name string, f func(t StateType) State) {
	stateTypes[name] = t
	stateNames[t] = name
	stateFactory[t] = f
}

func (t StateType) String() string {
	return stateNames[t]
}

func (t StateType) Create() State {
	f := stateFactory[t]
	return f(t)
}

func (t StateType) MarshalText() ([]byte, error) {
	return []byte(stateNames[t]), nil
}

func (t *StateType) UnmarshalText(b []byte) error {
	if stateType, ok := stateTypes[string(b)]; ok {
		*t = stateType
		return nil
	}
	// If here, add tx type to TxFactory and TxTypes in factory.go
	return fmt.Errorf("Unknown StateType: %s", string(b))
}

func GetStateTypeByName(name string) StateType {
	return stateTypes[name]
}

func GetStateTypes() []StateType {
	list := make([]StateType, 0)
	for _, t := range stateTypes {
		list = append(list, t)
	}
	return list
}

type TransactionResult int16

type resultItem struct {
	name string
	desc string
}

var resultNames = map[TransactionResult]resultItem{}
var resultTypes = map[string]TransactionResult{}

func (t TransactionResult) Register(name string, desc string) {
	resultNames[t] = resultItem{
		name: name,
		desc: desc,
	}
	resultTypes[name] = t
}

func (t TransactionResult) String() string {
	return resultNames[t].name
}

func (t TransactionResult) Desc() string {
	return resultNames[t].desc
}

func (r TransactionResult) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *TransactionResult) UnmarshalText(b []byte) error {
	if result, ok := resultTypes[string(b)]; ok {
		*r = result
		return nil
	}
	return fmt.Errorf("Unknown TransactionResult: %s", string(b))
}
