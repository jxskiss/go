package test

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/stretchr/testify/require"
	"github.com/thrift-iterator/go/general"
	"github.com/thrift-iterator/go/protocol"
	"github.com/thrift-iterator/go/raw"
	"github.com/thrift-iterator/go/test"
	"testing"
)

func Test_decode_set_by_iterator(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteSetBegin(thrift.STRING, 3)
		proto.WriteString("e1")
		proto.WriteString("e2")
		proto.WriteString("e3")
		proto.WriteSetEnd()
		iter := c.CreateIterator(buf.Bytes())
		elemType, length := iter.ReadSetHeader()
		should.Equal(protocol.TypeString, elemType)
		should.Equal(3, length)
		should.Equal("e1", iter.ReadString())
		should.Equal("e2", iter.ReadString())
		should.Equal("e3", iter.ReadString())
	}
}

func Test_encode_set_by_stream(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteSetHeader(protocol.TypeString, 3)
		stream.WriteString("e1")
		stream.WriteString("e2")
		stream.WriteString("e3")
		iter := c.CreateIterator(stream.Buffer())
		elemType, length := iter.ReadSetHeader()
		should.Equal(protocol.TypeString, elemType)
		should.Equal(3, length)
		should.Equal("e1", iter.ReadString())
		should.Equal("e2", iter.ReadString())
		should.Equal("e3", iter.ReadString())
	}
}

func Test_skip_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteSetBegin(thrift.I64, 3)
		proto.WriteI64(1)
		proto.WriteI64(2)
		proto.WriteI64(3)
		proto.WriteSetEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipSet(nil))
	}
}

func Test_unmarshal_general_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteSetBegin(thrift.I64, 3)
		proto.WriteI64(1)
		proto.WriteI64(2)
		proto.WriteI64(3)
		var val general.Set
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.Set{
			int64(1): true,
			int64(2): true,
			int64(3): true,
		}, val)
	}
}

func Test_unmarshal_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteSetBegin(thrift.I64, 3)
		proto.WriteI64(1)
		proto.WriteI64(2)
		proto.WriteI64(3)
		proto.WriteSetEnd()
		val := map[int64]bool{}
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(map[int64]bool{
			int64(1): true,
			int64(2): true,
			int64(3): true,
		}, val)
	}
}

func Test_marshal_general_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		m := general.Set{
			int32(1): true,
			int32(2): true,
			int32(3): true,
		}

		output, err := c.Marshal(m)
		should.NoError(err)
		output1, err := c.Marshal(&m)
		should.NoError(err)
		var val, val1 general.Set
		should.NoError(c.Unmarshal(output, &val))
		should.NoError(c.Unmarshal(output1, &val1))
		should.Equal(val, val1)
		should.Equal(general.Set{
			int32(1): true,
			int32(2): true,
			int32(3): true,
		}, val)
	}
}

func Test_marshal_raw_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteSetBegin(thrift.I32, 3)
		proto.WriteI32(1)
		proto.WriteI32(2)
		proto.WriteI32(3)
		proto.WriteSetEnd()
		var val raw.Set
		should.NoError(c.Unmarshal(buf.Bytes(), &val))

		output, err := c.Marshal(val)
		should.NoError(err)
		output1, err := c.Marshal(&val)
		should.NoError(err)
		var rawVal, rawVal1 general.Set
		should.NoError(c.Unmarshal(output, &rawVal))
		should.NoError(c.Unmarshal(output1, &rawVal1))
		should.Equal(rawVal, rawVal1)
		should.Equal(general.Set{
			int32(1): true,
			int32(2): true,
			int32(3): true,
		}, rawVal)
	}
}

func Test_marshal_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		m := map[string]bool{
			"e1": true,
			"e2": true,
			"e3": true,
		}

		output, err := c.Marshal(m)
		should.NoError(err)
		output1, err := c.Marshal(&m)
		should.NoError(err)
		var val, val1 general.Set
		should.NoError(c.Unmarshal(output, &val))
		should.NoError(c.Unmarshal(output1, &val1))
		should.Equal(val, val1)
		should.Equal(general.Set{
			"e1": true,
			"e2": true,
			"e3": true,
		}, val)
	}
}

func Test_marshal_empty_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(map[string]bool{})
		should.NoError(err)
		var val general.Set
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.Set{}, val)
	}
}
