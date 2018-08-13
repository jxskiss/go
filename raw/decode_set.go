package raw

import (
	"github.com/thrift-iterator/go/spi"
)

type rawSetDecoder struct {
}

func (decoder *rawSetDecoder) Decode(val interface{}, iter spi.Iterator) {
	elemType, length := iter.ReadSetHeader()
	elements := make(map[interface{}][]byte, length)
	generalElemReader := readerOf(elemType)
	elemIter := iter.Spawn()
	for i := 0; i < length; i++ {
		elemBuf := iter.Skip(elemType, nil)
		elem := generalElemReader(elemBuf, elemIter)
		elements[elem] = elemBuf
	}
	obj := val.(*Set)
	obj.ElementType = elemType
	obj.Elements = elements
}
