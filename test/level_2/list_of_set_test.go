package test

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/stretchr/testify/require"
	"github.com/thrift-iterator/go/general"
	"github.com/thrift-iterator/go/test"
	"testing"
)

func Test_skip_list_of_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.SET, 2)
		proto.WriteSetBegin(thrift.I32, 1)
		proto.WriteI32(1)
		proto.WriteSetEnd()
		proto.WriteSetBegin(thrift.I32, 1)
		proto.WriteI32(2)
		proto.WriteSetEnd()
		proto.WriteListEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipSet(nil))
	}
}

func Test_unmarshal_general_list_of_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.SET, 2)
		proto.WriteSetBegin(thrift.I32, 1)
		proto.WriteI32(1)
		proto.WriteSetEnd()
		proto.WriteSetBegin(thrift.I32, 1)
		proto.WriteI32(2)
		proto.WriteSetEnd()
		proto.WriteListEnd()
		var val general.List
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.Set{
			int32(1): true,
		}, val[0])
		should.Equal(true, val.Get(0, int32(1)))
		should.Equal(general.Set{
			int32(2): true,
		}, val[1])
		should.Equal(true, val.Get(1, int32(2)))
	}
}

func Test_unmarshal_list_of_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.SET, 2)
		proto.WriteSetBegin(thrift.I32, 1)
		proto.WriteI32(1)
		proto.WriteSetEnd()
		proto.WriteSetBegin(thrift.I32, 1)
		proto.WriteI32(2)
		proto.WriteSetEnd()
		proto.WriteListEnd()
		var val []map[int32]bool
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal([]map[int32]bool{
			{1: true}, {2: true},
		}, val)
	}
}

func Test_marshal_general_list_of_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		lst := general.List{
			general.Set{
				int32(1): true,
			},
			general.Set{
				int32(2): true,
			},
		}

		output, err := c.Marshal(lst)
		should.NoError(err)
		output1, err := c.Marshal(&lst)
		should.NoError(err)
		should.Equal(output, output1)
		var val []map[int32]bool
		should.NoError(c.Unmarshal(output, &val))
		should.Equal([]map[int32]bool{
			{1: true}, {2: true},
		}, val)
	}
}

func Test_marshal_list_of_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		lst := []map[int32]bool{
			{1: true}, {2: true},
		}

		output, err := c.Marshal(lst)
		should.NoError(err)
		output1, err := c.Marshal(&lst)
		should.NoError(err)
		should.Equal(output, output1)
		var val []map[int32]bool
		should.NoError(c.Unmarshal(output, &val))
		should.Equal([]map[int32]bool{
			{1: true}, {2: true},
		}, val)
	}
}
