package util

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type GlobalID struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func ToGlobalID(typ string, id interface{}) string {
	strid := ""

	switch v := id.(type) {
	case string:
		strid = v
	case int:
		strid = strconv.FormatInt(int64(v), 10)
	case int64:
		strid = strconv.FormatInt(v, 10)
	case uint:
		strid = strconv.FormatUint(uint64(v), 10)
	case uint64:
		strid = strconv.FormatUint(v, 10)
	default:
		strid = fmt.Sprintf("%v", v)
	}

	gid := fmt.Sprintf("%s/%s", typ, strid)

	return base64.StdEncoding.EncodeToString([]byte(gid))
}

func FromGlobalID(id string) (*GlobalID, error) {
	b, err := base64.StdEncoding.DecodeString(id)

	if err != nil {
		return nil, err
	}

	str := string(b)

	tokens := strings.Split(str, "/")

	if len(tokens) < 2 {
		return nil, fmt.Errorf("invalid global id: %s", id)
	}

	return &GlobalID{
		Type: tokens[0],
		ID:   tokens[1],
	}, nil
}

func FromGlobalIDInt64(gid string, typ string) (int64, error) {
	id, err := FromGlobalID(gid)

	if err != nil {
		return 0, err
	}

	if id.Type != typ {
		return 0, fmt.Errorf("invalid global id: %s", id)
	}

	return strconv.ParseInt(id.ID, 10, 64)
}
