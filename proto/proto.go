package sqlpx

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type pb struct {
	proto.Message
}

var pbPool = sync.Pool{
	New: func() interface{} {
		return new(proto.Buffer)
	},
}

// Proto is convenient wrapper to auto marshal and unmarshal Proto messages as []byte
// It works similar to proto Unmarshal and Marshal
func Proto(m proto.Message) interface {
	sql.Scanner
	driver.Valuer
} {
	return &pb{m}
}

func (p *pb) Scan(src interface{}) error {
	d, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("sqlpx: error while scanning proto: expected []byte go %T", src)
	}
	buf := pbPool.Get().(*proto.Buffer)

	buf.SetBuf(d)
	err := buf.Unmarshal(p.Message)
	pbPool.Put(buf)

	return err
}

func (p *pb) Value() (driver.Value, error) {
	buf := pbPool.Get().(*proto.Buffer)
	buf.Reset()
	if err := buf.Marshal(p.Message); err != nil {
		pbPool.Put(buf)
		return nil, err
	}

	b := buf.Bytes()
	dat := make([]byte, len(b))
	copy(dat, b)
	pbPool.Put(buf)

	return dat, nil
}

type pTime struct {
	*timestamp.Timestamp
}

// Timestamp is convenient wrapper to auto marshal and unmarshal proto Timestamp to time.Time
func Timestamp(t *timestamp.Timestamp) interface {
	sql.Scanner
	driver.Valuer
} {
	return &pTime{t}
}

func (t *pTime) Scan(src interface{}) error {
	v, ok := src.(time.Time)
	if !ok {
		return fmt.Errorf("sqlpx: error while scanning timestmap: expected time.Time go %T", src)
	}

	tsp, err := ptypes.TimestampProto(v)
	if err != nil {
		return err
	}

	t.Timestamp.Seconds = tsp.Seconds
	t.Timestamp.Nanos = tsp.Nanos
	return nil
}

func (t *pTime) Value() (driver.Value, error) {
	return ptypes.Timestamp(t.Timestamp)
}
