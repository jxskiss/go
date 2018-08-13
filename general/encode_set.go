package general

import (
	"github.com/thrift-iterator/go/protocol"
	"github.com/thrift-iterator/go/spi"
)

type generalSetEncoder struct {
}

func (encoder *generalSetEncoder) Encode(val interface{}, stream spi.Stream) {
	writeSet(val, stream)
}

func (encoder *generalSetEncoder) ThriftType() protocol.TType {
	return protocol.TypeSet
}

func takeSampleFromSet(sample Set) interface{} {
	for key := range sample {
		return key
	}
	panic("should not reach here")
}

func writeSet(val interface{}, stream spi.Stream) {
	obj := val.(Set)
	length := len(obj)
	if length == 0 {
		stream.WriteSetHeader(protocol.TypeI64, 0)
		return
	}
	elemSample := takeSampleFromSet(obj)
	elemType, generalWriter := generalWriterOf(elemSample)
	stream.WriteSetHeader(elemType, length)
	for elem := range obj {
		generalWriter(elem, stream)
	}
}
