// -*- mode: go -*-

// This file defines the serializer implementation for time datas.

package tu

import (
	"errors"
	
	"code.google.com/p/goprotobuf/proto"
	"github.com/neutrous/TimeService/model"
	"github.com/neutrous/notify/entity"
)

type TimeSerializer struct {
	data chan *model.TimeStruct
}

func NewSerializer(datach chan *model.TimeStruct) entity.Serializer {
	return &TimeSerializer{data: datach}
}

// Name indicates the name of the serializing data
func (obj *TimeSerializer)Name() string {
	return "model.TimeStruct"
}

// Serialize does the concrete serialize action on the data.
func (obj *TimeSerializer)Serialize() ([]byte, error) {
	if (obj.data != nil) {
		return proto.Marshal(<-obj.data)
	}
	return nil, errors.New("Channel model.TimeStruct has been close.")
}
