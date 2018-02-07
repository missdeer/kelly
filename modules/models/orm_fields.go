package models

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/orm"
)

// A true/false field.
type SliceStringField []string

func (e SliceStringField) Value() []string {
	return []string(e)
}

func (e *SliceStringField) Set(d []string) {
	*e = SliceStringField(d)
}

func (e *SliceStringField) Add(v string) {
	*e = append(*e, v)
}

func (e *SliceStringField) String() string {
	return strings.Join(e.Value(), ",")
}

func (e *SliceStringField) FieldType() int {
	return orm.TypeCharField
}

func (e *SliceStringField) SetRaw(value interface{}) error {
	switch d := value.(type) {
	case []string:
		e.Set(d)
	case string:
		if len(d) > 0 {
			parts := strings.Split(d, ",")
			v := make([]string, 0, len(parts))
			for _, p := range parts {
				v = append(v, strings.TrimSpace(p))
			}
			e.Set(v)
		}
	default:
		return fmt.Errorf("<SliceStringField.SetRaw> unknown value `%v`", value)
	}
	return nil
}

func (e *SliceStringField) RawValue() interface{} {
	return e.String()
}

func (e *SliceStringField) Clean() error {
	return nil
}

var _ orm.Fielder = new(SliceStringField)
