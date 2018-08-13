package reflection

import (
	"github.com/thrift-iterator/go/protocol"
	"github.com/thrift-iterator/go/spi"
	"reflect"
	"unsafe"
)

type mapEncoder struct {
	mapInterface emptyInterface
	keyEncoder   internalEncoder
	elemEncoder  internalEncoder
	tType        protocol.TType
}

func (encoder *mapEncoder) encode(ptr unsafe.Pointer, stream spi.Stream) {
	if encoder.tType == protocol.TypeSet {
		encoder.encodeSet(ptr, stream)
		return
	}
	encoder.encodeMap(ptr, stream)
}

func (encoder *mapEncoder) encodeMap(ptr unsafe.Pointer, stream spi.Stream) {
	mapInterface := encoder.mapInterface
	mapInterface.word = ptr
	realInterface := (*interface{})(unsafe.Pointer(&mapInterface))
	mapVal := reflect.ValueOf(*realInterface)
	keys := mapVal.MapKeys()
	stream.WriteMapHeader(encoder.keyEncoder.thriftType(), encoder.elemEncoder.thriftType(), len(keys))
	for _, key := range keys {
		keyObj := key.Interface()
		keyInf := (*emptyInterface)(unsafe.Pointer(&keyObj))
		encoder.keyEncoder.encode(keyInf.word, stream)
		elem := mapVal.MapIndex(key)
		elemObj := elem.Interface()
		elemInf := (*emptyInterface)(unsafe.Pointer(&elemObj))
		encoder.elemEncoder.encode(elemInf.word, stream)
	}
}

func (encoder *mapEncoder) encodeSet(ptr unsafe.Pointer, stream spi.Stream) {
	mapInterface := encoder.mapInterface
	mapInterface.word = ptr
	realInterface := (*interface{})(unsafe.Pointer(&mapInterface))
	mapVal := reflect.ValueOf(*realInterface)
	keys := mapVal.MapKeys()
	stream.WriteSetHeader(encoder.keyEncoder.thriftType(), len(keys))
	for _, key := range keys {
		keyObj := key.Interface()
		keyInf := (*emptyInterface)(unsafe.Pointer(&keyObj))
		encoder.keyEncoder.encode(keyInf.word, stream)
	}
}

func (encoder *mapEncoder) thriftType() protocol.TType {
	if encoder.tType == protocol.TypeSet {
		return protocol.TypeSet
	}
	return protocol.TypeMap
}
