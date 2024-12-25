package structTag

import (
	"reflect"
	"strings"
)

type StructTag struct {
	t reflect.Type
}

func New(s interface{}) *StructTag {
	return &StructTag{
		t: reflect.TypeOf(s),
	}
}

func (s StructTag) Get(tag string, fields ...string) (value string) {
	vs, _ := s.get(tag, false, fields...)
	if len(vs) > 0 {
		return vs[0]
	}
	return
}

func (s StructTag) GetFull(tag string, fields ...string) (value []string) {
	value, _ = s.get(tag, true, fields...)
	return
}

// 使用点点模式传输
func (s StructTag) GetPoint(tag, fields string) (value string) {
	values, _ := s.get(tag, true, strings.Split(fields, ".")...)
	value = strings.Join(values, ".")
	return
}

func (s StructTag) get(tag string, full bool, fields ...string) ([]string, bool) {
	var (
		ok bool
		v  string
		sf reflect.StructField
	)

	tags := make([]string, 0, len(fields))
	for i, f := range fields {
		if s.t.Kind() == reflect.Slice {
			s.t = s.t.Elem()
		}
		if s.t.Kind() == reflect.Pointer {
			s.t = s.t.Elem()
		}

		sf, ok = s.t.FieldByName(f)
		if !ok {
			return nil, ok
		}

		if full || len(fields) == i+1 {
			v, ok = sf.Tag.Lookup(tag)
			v = trimLabel(v)
			// 字段Tag 注释 -,不存在，则说明对应 Tag 不存在
			if v == "-" {
				return nil, ok
			}
			// 如果不存在，可以使用field 填值
			if v == "" {
				v = f
			}
			tags = append(tags, v)
		}
		s.t = sf.Type
	}
	return tags, ok
}

// 获取Tag 对应值，去掉描述
func trimLabel(tag string) string {
	tags := strings.Split(tag, ",")
	if len(tags) == 0 {
		return ""
	}
	tag = strings.TrimSpace(tags[0])
	return tag
}
