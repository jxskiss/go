package general

import "github.com/thrift-iterator/go/spi"

type generalSetDecoder struct {
}

func (decoder *generalSetDecoder) Decode(val interface{}, iter spi.Iterator) {
	*val.(*Set) = readSet(iter).(Set)
}

func readSet(iter spi.Iterator) interface{} {
	elemType, length := iter.ReadSetHeader()
	generalReader := generalReaderOf(elemType)
	generalSet := Set{}
	if length == 0 {
		return generalSet
	}
	for i := 0; i < length; i++ {
		elem := generalReader(iter)
		generalSet[elem] = true
	}
	return generalSet
}
