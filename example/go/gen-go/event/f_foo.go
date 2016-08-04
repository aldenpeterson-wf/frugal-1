// Autogenerated by Frugal Compiler (1.12.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package event

import (
	"bytes"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/Workiva/frugal/example/go/gen-go/base"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = base.GoUnusedProtection__

type FooPingArgs struct {
}

func NewFooPingArgs() *FooPingArgs {
	return &FooPingArgs{}
}

func (p *FooPingArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FooPingArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ping_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FooPingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FooPingArgs(%+v)", *p)
}

type FooPingResult struct {
}

func NewFooPingResult() *FooPingResult {
	return &FooPingResult{}
}

func (p *FooPingResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FooPingResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ping_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FooPingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FooPingResult(%+v)", *p)
}

type FooBlahArgs struct {
	Num   int32  `thrift:"num,1" db:"num" json:"num"`
	Str   string `thrift:"Str,2" db:"Str" json:"Str"`
	Event *Event `thrift:"event,3" db:"event" json:"event"`
}

func NewFooBlahArgs() *FooBlahArgs {
	return &FooBlahArgs{}
}

func (p *FooBlahArgs) GetNum() int32 {
	return p.Num
}

func (p *FooBlahArgs) GetStr() string {
	return p.Str
}

var FooBlahArgs_Event_DEFAULT *Event

func (p *FooBlahArgs) IsSetEvent() bool {
	return p.Event != nil
}

func (p *FooBlahArgs) GetEvent() *Event {
	if !p.IsSetEvent() {
		return FooBlahArgs_Event_DEFAULT
	}
	return p.Event
}

func (p *FooBlahArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FooBlahArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Num = v
	}
	return nil
}

func (p *FooBlahArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Str = v
	}
	return nil
}

func (p *FooBlahArgs) ReadField3(iprot thrift.TProtocol) error {
	p.Event = NewEvent()
	if err := p.Event.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Event), err)
	}
	return nil
}

func (p *FooBlahArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("blah_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FooBlahArgs) writeField1(oprot thrift.TProtocol) error {
	if err := oprot.WriteFieldBegin("num", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:num: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Num)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.num (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:num: ", p), err)
	}
	return nil
}

func (p *FooBlahArgs) writeField2(oprot thrift.TProtocol) error {
	if err := oprot.WriteFieldBegin("Str", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:Str: ", p), err)
	}
	if err := oprot.WriteString(string(p.Str)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Str (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:Str: ", p), err)
	}
	return nil
}

func (p *FooBlahArgs) writeField3(oprot thrift.TProtocol) error {
	if err := oprot.WriteFieldBegin("event", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:event: ", p), err)
	}
	if err := p.Event.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Event), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:event: ", p), err)
	}
	return nil
}

func (p *FooBlahArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FooBlahArgs(%+v)", *p)
}

type FooBlahResult struct {
	Success *int64             `thrift:"success,0" db:"success" json:"success,omitempty"`
	Awe     *AwesomeException  `thrift:"awe,1" db:"awe" json:"awe,omitempty"`
	API     *base.APIException `thrift:"api,2" db:"api" json:"api,omitempty"`
}

func NewFooBlahResult() *FooBlahResult {
	return &FooBlahResult{}
}

var FooBlahResult_Success_DEFAULT int64

func (p *FooBlahResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FooBlahResult) GetSuccess() int64 {
	if !p.IsSetSuccess() {
		return FooBlahResult_Success_DEFAULT
	}
	return *p.Success
}

var FooBlahResult_Awe_DEFAULT *AwesomeException

func (p *FooBlahResult) IsSetAwe() bool {
	return p.Awe != nil
}

func (p *FooBlahResult) GetAwe() *AwesomeException {
	if !p.IsSetAwe() {
		return FooBlahResult_Awe_DEFAULT
	}
	return p.Awe
}

var FooBlahResult_API_DEFAULT *base.APIException

func (p *FooBlahResult) IsSetAPI() bool {
	return p.API != nil
}

func (p *FooBlahResult) GetAPI() *base.APIException {
	if !p.IsSetAPI() {
		return FooBlahResult_API_DEFAULT
	}
	return p.API
}

func (p *FooBlahResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FooBlahResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *FooBlahResult) ReadField1(iprot thrift.TProtocol) error {
	p.Awe = NewAwesomeException()
	if err := p.Awe.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Awe), err)
	}
	return nil
}

func (p *FooBlahResult) ReadField2(iprot thrift.TProtocol) error {
	p.API = base.NewAPIException()
	if err := p.API.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.API), err)
	}
	return nil
}

func (p *FooBlahResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("blah_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FooBlahResult) writeField0(oprot thrift.TProtocol) error {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I64, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return nil
}

func (p *FooBlahResult) writeField1(oprot thrift.TProtocol) error {
	if p.IsSetAwe() {
		if err := oprot.WriteFieldBegin("awe", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:awe: ", p), err)
		}
		if err := p.Awe.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Awe), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:awe: ", p), err)
		}
	}
	return nil
}

func (p *FooBlahResult) writeField2(oprot thrift.TProtocol) error {
	if p.IsSetAPI() {
		if err := oprot.WriteFieldBegin("api", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:api: ", p), err)
		}
		if err := p.API.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.API), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:api: ", p), err)
		}
	}
	return nil
}

func (p *FooBlahResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FooBlahResult(%+v)", *p)
}

type FooOneWayArgs struct {
	ID  ID      `thrift:"id,1" db:"id" json:"id"`
	Req Request `thrift:"req,2" db:"req" json:"req"`
}

func NewFooOneWayArgs() *FooOneWayArgs {
	return &FooOneWayArgs{}
}

func (p *FooOneWayArgs) GetID() ID {
	return p.ID
}

func (p *FooOneWayArgs) GetReq() Request {
	return p.Req
}

func (p *FooOneWayArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FooOneWayArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := ID(v)
		p.ID = temp
	}
	return nil
}

func (p *FooOneWayArgs) ReadField2(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	p.Req = make(Request, size)
	for i := 0; i < size; i++ {
		var elem0 Int
		if v, err := iprot.ReadI32(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			temp := Int(v)
			elem0 = temp
		}
		var elem1 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			elem1 = v
		}
		(p.Req)[elem0] = elem1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *FooOneWayArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("oneWay_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FooOneWayArgs) writeField1(oprot thrift.TProtocol) error {
	if err := oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err)
	}
	return nil
}

func (p *FooOneWayArgs) writeField2(oprot thrift.TProtocol) error {
	if err := oprot.WriteFieldBegin("req", thrift.MAP, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:req: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.I32, thrift.STRING, len(p.Req)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Req {
		if err := oprot.WriteI32(int32(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:req: ", p), err)
	}
	return nil
}

func (p *FooOneWayArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FooOneWayArgs(%+v)", *p)
}
