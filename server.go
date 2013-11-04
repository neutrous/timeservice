// -*- mode: go -*-

package tu 						// for time uniform

import (
	"time"
	
	"github.com/neutrous/notify/entity"
	"github.com/neutrous/TimeService/model"
	"code.google.com/p/goprotobuf/proto"
)

type TimeGenerator func() *model.TimeStruct 

type Server struct {
	env entity.CommEnv
	pub entity.Publisher
	// indicates the time interval to send
	// unified time. Unit: seconds
	Interval int32				
	data chan *model.TimeStruct
	Generator TimeGenerator
}

// NewServer creates a new Server instance.
func NewServer(address []string) *Server {
	server := Server{Interval: 1,
		data: make(chan *model.TimeStruct), 
		Generator: newTimeFormatData,
	}
	for _, val := range address {
		server.pub.AppendAddress(entity.Address(val))
	}
	env, err := entity.CreateZMQCommEnv(true)
	if err != nil {
		return nil
	}
	if server.pub.InitialBinding(env) != nil {
		return nil
	}
	return &server
}

// Close release the instance's resources. Calling this method would
// cause the Publish method to quit.
func (obj *Server)Close() {
	if (obj != nil) {
		obj.env.Close()
	}
}

// Publish pubs the current server timetamp to all of the subscribers.
// This method would cause the goroutine blocked.
func (obj* Server)Publish() {
	// initialize the serializer, uses the data channel to communicate
	// with the sender.
	ser := NewSerializer(obj.data)
	go func (loser entity.Serializer) {
		for {
			if err := obj.pub.Send(loser); err != nil {
				close(obj.data)
				return
			}
		}
	}(ser)
	
	for {
		select {
		case <-time.After(time.Duration(obj.Interval) * time.Second):
			if obj.data != nil {
				obj.data <- obj.Generator()
			}
			// if err := obj.pub.Send(ser); err != nil {
			// 	return
			// }
		}
	}
}

func newTimeFormatData() *model.TimeStruct {
	current := time.Now()
	serialData := model.TimeStruct{}
	serialData.Year = proto.Int(current.Year())
	serialData.Month = proto.Int(int(current.Month()))
	serialData.Day = proto.Int(current.Day())
	serialData.Hour = proto.Int(current.Hour())
	serialData.Minitue = proto.Int(current.Minute())
	serialData.Second = proto.Int(current.Second())

	serialData.Nsec = proto.Int(current.Nanosecond())
	serialData.Format = proto.String(current.Format(time.RFC1123Z))
	
	return &serialData
}

















