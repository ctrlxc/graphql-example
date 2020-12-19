package pagination

import (
	"encoding/base64"
	"encoding/json"
	"reflect"

	"github.com/volatiletech/strmangle"
)

type Cursor struct {
	Items []*CursorItem `json:"items"`
}

type CursorItem struct {
	Field     string      `json:"field"`
	Direction Direction   `json:"direction"`
	Value     interface{} `json:"value"`
	ValueType string      `json:"value_type"`
}

func (p *Paginator) CursorEncode(cursor *Cursor) (string, error) {
	b, err := json.Marshal(cursor)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func (p *Paginator) CursorDecode(cursorstr string) (*Cursor, error) {
	b, err := base64.StdEncoding.DecodeString(cursorstr)

	if err != nil {
		return nil, err
	}

	cursor := Cursor{}

	err = json.Unmarshal(b, &cursor)

	if err != nil {
		return nil, err
	}

	// FIXME: date or time convert?

	// for _, item := range cursor.Items {
	// 	if item.ValueType == "Time" {
	// 		if v, ok := (item.Value.(string)); ok {
	// 			item.Value, _ = time.Parse("", v)
	// 		}
	// 	}
	// }

	return &cursor, nil
}

func (p *Paginator) CreateCursor(v interface{}) *Cursor {
	rv := reflectValue(v)

	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	items := make([]*CursorItem, len(p.Orders))

	for i, o := range p.Orders {
		camelField := strmangle.TitleCase(o.Field)

		f := rv.FieldByName(camelField)

		items[i] = &CursorItem{
			Field:     o.Field,
			Direction: o.Direction,
			Value:     f.Interface(), // FIXME: date or time convert?
			ValueType: f.Type().Name(),
		}
	}

	return &Cursor{Items: items}
}

func (p *Paginator) CreateEncodedCursor(v interface{}) (string, error) {
	c := p.CreateCursor(v)
	return p.CursorEncode(c)
}

func reflectValue(v interface{}) reflect.Value {
	rv, ok := v.(reflect.Value)

	if !ok {
		return reflect.ValueOf(v)
	}

	return rv
}
