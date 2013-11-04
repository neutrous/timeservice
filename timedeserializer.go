// -*- mode: go -*-

// timedeserializer defines the implementation to deserialize the data
// from the raw bytes.

package tu

import (
	"log"
	
	"code.google.com/p/goprotobuf/proto"
	"github.com/neutrous/TimeService/model"
	"github.com/neutrous/notify/entity"
)

type TimeDeserializer struct {
}

func NewTimeDeserializer() entity.Deserializer {
	return &TimeDeserializer{}
}

func (obj *TimeDeserializer) Name() string {
	return "model.TimeStruct"
}

func (obj *TimeDeserializer) Deserialize(data []byte) (inst interface{},
	ok bool) {
	timeobj := &model.TimeStruct{}
	inst, ok = timeobj, true
	if proto.Unmarshal(data, timeobj) != nil {
		ok = false
	}
	return
}

func (obj *TimeDeserializer) DataReceived(inst interface{},
	ok bool) {
	if !ok {
		log.Println("TimeDeserializer::DataReceived " +
			"unmarshaled failure.")
		return
	}
	if obj, is := inst.(*model.TimeStruct); is {
		log.Println(obj.GetFormat())
	}
}










