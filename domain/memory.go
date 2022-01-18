package domain

type Memory struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewMemory(Key, Value string) Memory {
	return Memory{
		Key:   Key,
		Value: Value,
	}
}
