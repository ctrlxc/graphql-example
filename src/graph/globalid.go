package graph

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

func toGlobalID(typ string, id string) string {
	gid := fmt.Sprintf("%s/%s", typ, id)

	return base64.StdEncoding.EncodeToString([]byte(gid))
}

func toGlobalIDs(typ string, ids []string) []string {
	gids := make([]string, len(ids))

	for i, id := range ids {
		gids[i] = toGlobalID(typ, id)
	}

	return gids
}

func toGlobalIDInt64(typ string, id int64) string {
	gid := fmt.Sprintf("%s/%v", typ, id)

	return base64.StdEncoding.EncodeToString([]byte(gid))
}

func toGlobalIDInt64s(typ string, ids []int64) []string {
	gids := make([]string, len(ids))

	for i, id := range ids {
		gids[i] = toGlobalIDInt64(typ, id)
	}

	return gids
}

func fromGlobalID(id string) (*GlobalID, error) {
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

func fromGlobalIDInt64(gid string, typ string) (int64, error) {
	id, err := fromGlobalID(gid)

	if err != nil {
		return 0, err
	}

	if id.Type != typ {
		return 0, fmt.Errorf("invalid global id: %s", id)
	}

	return strconv.ParseInt(id.ID, 10, 64)
}

func fromGlobalIDInt64s(gids []string, typ string) ([]int64, error) {
	ids := make([]int64, len(gids))

	for i, gid := range gids {
		id, err := fromGlobalIDInt64(gid, typ)

		if err != nil {
			return nil, err
		}

		ids[i] = id
	}

	return ids, nil
}
