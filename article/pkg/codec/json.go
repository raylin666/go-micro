package codec

import (
	"encoding/json"
	"reflect"

	"github.com/go-kratos/kratos/v2/encoding"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// JSON Name is the name registered for the json codec.
const JSON = "json"

var (
	// JsonMarshalOptions is a configurable JSON format marshaller.
	JsonMarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true, // 不将 ProtoBuf 定义的下划线字段转换为驼峰写法
	}
	// JsonUnmarshalOptions is a configurable JSON format parser.
	JsonUnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)

func init() {
	encoding.RegisterCodec(jsonCodec{})
}

// jsonCodec is a Codec implementation with json.
type jsonCodec struct{}

func (jsonCodec) Marshal(v interface{}) ([]byte, error) {
	switch m := v.(type) {
	case json.Marshaler:
		return m.MarshalJSON()
	case proto.Message:
		return JsonMarshalOptions.Marshal(m)
	default:
		return json.Marshal(m)
	}
}

func (jsonCodec) Unmarshal(data []byte, v interface{}) error {
	switch m := v.(type) {
	case json.Unmarshaler:
		return m.UnmarshalJSON(data)
	case proto.Message:
		return JsonUnmarshalOptions.Unmarshal(data, m)
	default:
		rv := reflect.ValueOf(v)
		for rv := rv; rv.Kind() == reflect.Ptr; {
			if rv.IsNil() {
				rv.Set(reflect.New(rv.Type().Elem()))
			}
			rv = rv.Elem()
		}
		if m, ok := reflect.Indirect(rv).Interface().(proto.Message); ok {
			return JsonUnmarshalOptions.Unmarshal(data, m)
		}
		return json.Unmarshal(data, m)
	}
}

func (jsonCodec) Name() string {
	return JSON
}
