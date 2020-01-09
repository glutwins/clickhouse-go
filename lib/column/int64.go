package column

import (
	"github.com/ClickHouse/clickhouse-go/lib/binary"
	"strconv"
)

type Int64 struct{ base }

func (Int64) Read(decoder *binary.Decoder) (interface{}, error) {
	v, err := decoder.Int64()
	if err != nil {
		return int64(0), err
	}
	return v, nil
}

func (i *Int64) Write(encoder *binary.Encoder, v interface{}) error {
	switch v := v.(type) {
	case int:
		return encoder.Int64(int64(v))
	case int32:
		return encoder.Int64(int64(v))
	case uint32:
		return encoder.Int64(int64(v))
	case int16:
		return encoder.Int64(int64(v))
	case uint16:
		return encoder.Int64(int64(v))
	case int8:
		return encoder.Int64(int64(v))
	case uint8:
		return encoder.Int64(int64(v))
	case int64:
		return encoder.Int64(v)
	case []byte:
		if _, err := encoder.Write(v); err != nil {
			return err
		}
		return nil
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return encoder.Int64(0)
		}
		return encoder.Int64(i)

	// this relies on Nullable never sending nil values through
	case *int:
		return encoder.Int64(int64(*v))
	case *int32:
		return encoder.Int64(int64(*v))
	case *uint32:
		return encoder.Int64(int64(*v))
	case *int16:
		return encoder.Int64(int64(*v))
	case *uint16:
		return encoder.Int64(int64(*v))
	case *int8:
		return encoder.Int64(int64(*v))
	case *uint8:
		return encoder.Int64(int64(*v))
	case *int64:
		return encoder.Int64(*v)
	}

	return &ErrUnexpectedType{
		T:      v,
		Column: i,
	}
}
