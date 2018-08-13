package test

import (
	"testing"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/stretchr/testify/require"
	"github.com/thrift-iterator/go/general"
	"github.com/thrift-iterator/go/protocol"
	"github.com/thrift-iterator/go/test"
	"github.com/thrift-iterator/go/test/level_2/struct_of_set_test"
)

func Test_skip_struct_of_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.SET, 1)
		proto.WriteSetBegin(thrift.I32, 1)
		proto.WriteI32(2)
		proto.WriteSetEnd()
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipStruct(nil))
	}
}

func Test_unmarshal_general_struct_of_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.SET, 1)
		proto.WriteSetBegin(thrift.I32, 1)
		proto.WriteI32(2)
		proto.WriteSetEnd()
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		var val general.Struct
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.Set{
			int32(2): true,
		}, val[protocol.FieldId(1)])
	}
}

func Test_unmarshal_struct_of_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.SET, 1)
		proto.WriteSetBegin(thrift.I32, 1)
		proto.WriteI32(2)
		proto.WriteSetEnd()
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		var val struct_of_set_test.TestObject
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(struct_of_set_test.TestObject{
			map[int32]bool{2: true},
		}, val)
	}
}

func Test_marshal_general_struct_of_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		m := general.Struct{
			protocol.FieldId(1): general.Set{
				int32(2): true,
			},
		}

		output, err := c.Marshal(m)
		should.NoError(err)
		output1, err := c.Marshal(&m)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.Struct
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.Set{
			int32(2): true,
		}, val[protocol.FieldId(1)])
	}
}

func Test_marshal_struct_of_set(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		m := struct_of_set_test.TestObject{
			map[int32]bool{2: true},
		}

		output, err := c.Marshal(m)
		should.NoError(err)
		output1, err := c.Marshal(&m)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.Struct
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.Set{
			int32(2): true,
		}, val[protocol.FieldId(1)])
	}
}
