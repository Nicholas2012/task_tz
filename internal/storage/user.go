package storage

import "encoding/json"

type User struct {
	Login string
	Data  json.RawMessage
}
