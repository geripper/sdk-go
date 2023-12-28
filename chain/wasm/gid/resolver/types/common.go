package types

type Config struct {
	IntefaceId      uint64 `json:"interface_id"`
	Owner           string `json:"owner"`
	RegistryAddress string `json:"registry_address"`
}

type TextData struct {
	Keys []string `json:"keys"`
	Node []uint   `json:"node"`
}
