package encoder

import (
	"github.com/vmihailenco/msgpack/v4"
)

type MsgPackEncoder struct {
	// bytespool bytebufferpool.Pool
}

// Encode
func (m *MsgPackEncoder) Encode(_ string, v interface{}) ([]byte, error) {
	// bb := m.bytespool.Get()
	// err := msgpack.NewEncoder(bb).Encode(v)
	// defer m.bytespool.Put(bb)
	return msgpack.Marshal(v)
}

// Decode
func (m *MsgPackEncoder) Decode(_ string, data []byte, vPtr interface{}) error {
	return msgpack.Unmarshal(data, vPtr)
	// switch arg := vPtr.(type) {
	// case *string:
	// 	// If they want a string and it is a JSON string, strip quotes
	// 	// This allows someone to send a struct but receive as a plain string
	// 	// This cast should be efficient for Go 1.3 and beyond.
	// 	str := string(data)
	// 	if strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`) {
	// 		*arg = str[1 : len(str)-1]
	// 	} else {
	// 		*arg = str
	// 	}
	// case *[]byte:
	// 	*arg = data
	// default:
	// 	err = json.Unmarshal(data, arg)
	// }
}
