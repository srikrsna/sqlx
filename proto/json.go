package sqlpx

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var marshaller jsonpb.Marshaler

type jsonb struct {
	v proto.Message
}

// JSONB is convenient wrapper to auto marshal and unmarshal JSON data type as []byte
// v should be a pointer. It works similar to json.Unmarshal and Marshal
func JSONB(v proto.Message) interface {
	driver.Valuer
	sql.Scanner
} {
	return &jsonb{v}
}

func (j *jsonb) Scan(src interface{}) error {
	dat, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("sqlpx: error while scanning jsonb: expected []byte got %T", src)
	}

	return jsonpb.Unmarshal(bytes.NewReader(dat), j.v)
}

func (j *jsonb) Value() (driver.Value, error) {
	var buf bytes.Buffer

	err := marshaller.Marshal(&buf, j.v)

	return buf.Bytes(), err
}
