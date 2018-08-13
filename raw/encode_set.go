package raw

import (
	"github.com/thrift-iterator/go/protocol"
	"github.com/thrift-iterator/go/spi"
)

type rawSetEncoder struct {
}

func (encoder *rawSetEncoder) Encode(val interface{}, stream spi.Stream) {
	obj := val.(Set)
	length := len(obj.Elements)
	stream.WriteSetHeader(obj.ElementType, length)
	for _, elem := range obj.Elements {
		stream.Write(elem)
	}
}

func (encoder *rawSetEncoder) ThriftType() protocol.TType {
	return protocol.TypeSet
}
