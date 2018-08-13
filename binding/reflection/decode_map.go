package reflection

import (
	"github.com/thrift-iterator/go/protocol"
	"github.com/thrift-iterator/go/spi"
	"reflect"
	"unsafe"
)

var reflectTrueValue = reflect.ValueOf(true)

type mapDecoder struct {
	mapType      reflect.Type
	mapInterface emptyInterface
	keyType      reflect.Type
	keyDecoder   internalDecoder
	elemType     reflect.Type
	elemDecoder  internalDecoder
	tType        protocol.TType
}

func (decoder *mapDecoder) decode(ptr unsafe.Pointer, iter spi.Iterator) {
	if decoder.tType == protocol.TypeSet {
		decoder.decodeSet(ptr, iter)
		return
	}
	decoder.decodeMap(ptr, iter)
}

func (decoder *mapDecoder) decodeMap(ptr unsafe.Pointer, iter spi.Iterator) {
	mapInterface := decoder.mapInterface
	mapInterface.word = ptr
	realInterface := (*interface{})(unsafe.Pointer(&mapInterface))
	mapVal := reflect.ValueOf(*realInterface).Elem()
	if mapVal.IsNil() {
		mapVal.Set(reflect.MakeMap(decoder.mapType))
	}
	_, elemType, length := iter.ReadMapHeader()
	if elemType == 0 { // set
		decoder.readSet(mapVal, length, iter)
		return
	}
	decoder.readMap(mapVal, length, iter)
}

func (decoder *mapDecoder) decodeSet(ptr unsafe.Pointer, iter spi.Iterator) {
	mapInterface := decoder.mapInterface
	mapInterface.word = ptr
	realInterface := (*interface{})(unsafe.Pointer(&mapInterface))
	mapVal := reflect.ValueOf(*realInterface).Elem()
	if mapVal.IsNil() {
		mapVal.Set(reflect.MakeMap(decoder.mapType))
	}
	_, length := iter.ReadSetHeader()
	decoder.readSet(mapVal, length, iter)
}

func (decoder *mapDecoder) readMap(mapVal reflect.Value, length int, iter spi.Iterator) {
	for i := 0; i < length; i++ {
		keyVal := reflect.New(decoder.keyType)
		decoder.keyDecoder.decode(unsafe.Pointer(keyVal.Pointer()), iter)
		elemVal := reflect.New(decoder.elemType)
		decoder.elemDecoder.decode(unsafe.Pointer(elemVal.Pointer()), iter)
		mapVal.SetMapIndex(keyVal.Elem(), elemVal.Elem())
	}
}

func (decoder *mapDecoder) readSet(mapVal reflect.Value, length int, iter spi.Iterator) {
	for i := 0; i < length; i++ {
		keyVal := reflect.New(decoder.keyType)
		decoder.keyDecoder.decode(unsafe.Pointer(keyVal.Pointer()), iter)
		mapVal.SetMapIndex(keyVal.Elem(), reflectTrueValue)
	}
}
