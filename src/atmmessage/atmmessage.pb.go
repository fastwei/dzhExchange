// Code generated by protoc-gen-go.
// source: atmmessage.proto
// DO NOT EDIT!

/*
Package atmmessage is a generated protocol buffer package.

It is generated from these files:
	atmmessage.proto

It has these top-level messages:
	Header
	Header2
	Did60130
	Did60132
	Did60133
	Did60134
	Did60135
*/
package atmmessage

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Header struct {
	Did              *uint32 `protobuf:"varint,1,req,name=did" json:"did,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}

func (m *Header) GetDid() uint32 {
	if m != nil && m.Did != nil {
		return *m.Did
	}
	return 0
}

type Header2 struct {
	H                *Header `protobuf:"bytes,1,req,name=h" json:"h,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Header2) Reset()         { *m = Header2{} }
func (m *Header2) String() string { return proto.CompactTextString(m) }
func (*Header2) ProtoMessage()    {}

func (m *Header2) GetH() *Header {
	if m != nil {
		return m.H
	}
	return nil
}

type Did60130 struct {
	H                *Header       `protobuf:"bytes,1,req,name=h" json:"h,omitempty"`
	Obj              *string       `protobuf:"bytes,2,opt,name=obj" json:"obj,omitempty"`
	Time             *uint64       `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	Open             *float64      `protobuf:"fixed64,4,opt,name=open" json:"open,omitempty"`
	High             *float64      `protobuf:"fixed64,5,opt,name=high" json:"high,omitempty"`
	Low              *float64      `protobuf:"fixed64,6,opt,name=low" json:"low,omitempty"`
	New              *float64      `protobuf:"fixed64,7,opt,name=new" json:"new,omitempty"`
	Volume           *uint64       `protobuf:"varint,8,opt,name=volume" json:"volume,omitempty"`
	Amout            *uint64       `protobuf:"varint,9,opt,name=amout" json:"amout,omitempty"`
	Innervol         *uint64       `protobuf:"varint,10,opt,name=innervol" json:"innervol,omitempty"`
	Outtervol        *uint64       `protobuf:"varint,11,opt,name=outtervol" json:"outtervol,omitempty"`
	Tickcount        *uint32       `protobuf:"varint,12,opt,name=tickcount" json:"tickcount,omitempty"`
	Avgbuyprice      *float64      `protobuf:"fixed64,13,opt,name=avgbuyprice" json:"avgbuyprice,omitempty"`
	Allbuyvol        *uint64       `protobuf:"varint,14,opt,name=allbuyvol" json:"allbuyvol,omitempty"`
	Avgsellprice     *float64      `protobuf:"fixed64,15,opt,name=avgsellprice" json:"avgsellprice,omitempty"`
	Allsellvol       *uint64       `protobuf:"varint,16,opt,name=allsellvol" json:"allsellvol,omitempty"`
	Buy              []*Did60130_M `protobuf:"bytes,17,rep,name=buy" json:"buy,omitempty"`
	Sell             []*Did60130_M `protobuf:"bytes,18,rep,name=sell" json:"sell,omitempty"`
	Tradestatus      *string       `protobuf:"bytes,19,opt,name=tradestatus" json:"tradestatus,omitempty"`
	Iopv             *float64      `protobuf:"fixed64,20,opt,name=iopv" json:"iopv,omitempty"`
	Accruedint       *float64      `protobuf:"fixed64,21,opt,name=accruedint" json:"accruedint,omitempty"`
	Openinterest     *uint64       `protobuf:"varint,22,opt,name=openinterest" json:"openinterest,omitempty"`
	Settleprice      *float64      `protobuf:"fixed64,23,opt,name=settleprice" json:"settleprice,omitempty"`
	Advstop          *float64      `protobuf:"fixed64,24,opt,name=advstop" json:"advstop,omitempty"`
	Decstop          *float64      `protobuf:"fixed64,25,opt,name=decstop" json:"decstop,omitempty"`
	Lastclose        *float64      `protobuf:"fixed64,26,opt,name=lastclose" json:"lastclose,omitempty"`
	Name             *string       `protobuf:"bytes,27,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Did60130) Reset()         { *m = Did60130{} }
func (m *Did60130) String() string { return proto.CompactTextString(m) }
func (*Did60130) ProtoMessage()    {}

func (m *Did60130) GetH() *Header {
	if m != nil {
		return m.H
	}
	return nil
}

func (m *Did60130) GetObj() string {
	if m != nil && m.Obj != nil {
		return *m.Obj
	}
	return ""
}

func (m *Did60130) GetTime() uint64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Did60130) GetOpen() float64 {
	if m != nil && m.Open != nil {
		return *m.Open
	}
	return 0
}

func (m *Did60130) GetHigh() float64 {
	if m != nil && m.High != nil {
		return *m.High
	}
	return 0
}

func (m *Did60130) GetLow() float64 {
	if m != nil && m.Low != nil {
		return *m.Low
	}
	return 0
}

func (m *Did60130) GetNew() float64 {
	if m != nil && m.New != nil {
		return *m.New
	}
	return 0
}

func (m *Did60130) GetVolume() uint64 {
	if m != nil && m.Volume != nil {
		return *m.Volume
	}
	return 0
}

func (m *Did60130) GetAmout() uint64 {
	if m != nil && m.Amout != nil {
		return *m.Amout
	}
	return 0
}

func (m *Did60130) GetInnervol() uint64 {
	if m != nil && m.Innervol != nil {
		return *m.Innervol
	}
	return 0
}

func (m *Did60130) GetOuttervol() uint64 {
	if m != nil && m.Outtervol != nil {
		return *m.Outtervol
	}
	return 0
}

func (m *Did60130) GetTickcount() uint32 {
	if m != nil && m.Tickcount != nil {
		return *m.Tickcount
	}
	return 0
}

func (m *Did60130) GetAvgbuyprice() float64 {
	if m != nil && m.Avgbuyprice != nil {
		return *m.Avgbuyprice
	}
	return 0
}

func (m *Did60130) GetAllbuyvol() uint64 {
	if m != nil && m.Allbuyvol != nil {
		return *m.Allbuyvol
	}
	return 0
}

func (m *Did60130) GetAvgsellprice() float64 {
	if m != nil && m.Avgsellprice != nil {
		return *m.Avgsellprice
	}
	return 0
}

func (m *Did60130) GetAllsellvol() uint64 {
	if m != nil && m.Allsellvol != nil {
		return *m.Allsellvol
	}
	return 0
}

func (m *Did60130) GetBuy() []*Did60130_M {
	if m != nil {
		return m.Buy
	}
	return nil
}

func (m *Did60130) GetSell() []*Did60130_M {
	if m != nil {
		return m.Sell
	}
	return nil
}

func (m *Did60130) GetTradestatus() string {
	if m != nil && m.Tradestatus != nil {
		return *m.Tradestatus
	}
	return ""
}

func (m *Did60130) GetIopv() float64 {
	if m != nil && m.Iopv != nil {
		return *m.Iopv
	}
	return 0
}

func (m *Did60130) GetAccruedint() float64 {
	if m != nil && m.Accruedint != nil {
		return *m.Accruedint
	}
	return 0
}

func (m *Did60130) GetOpeninterest() uint64 {
	if m != nil && m.Openinterest != nil {
		return *m.Openinterest
	}
	return 0
}

func (m *Did60130) GetSettleprice() float64 {
	if m != nil && m.Settleprice != nil {
		return *m.Settleprice
	}
	return 0
}

func (m *Did60130) GetAdvstop() float64 {
	if m != nil && m.Advstop != nil {
		return *m.Advstop
	}
	return 0
}

func (m *Did60130) GetDecstop() float64 {
	if m != nil && m.Decstop != nil {
		return *m.Decstop
	}
	return 0
}

func (m *Did60130) GetLastclose() float64 {
	if m != nil && m.Lastclose != nil {
		return *m.Lastclose
	}
	return 0
}

func (m *Did60130) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type Did60130_M struct {
	Price            *float64 `protobuf:"fixed64,1,opt,name=price" json:"price,omitempty"`
	Vol              *uint32  `protobuf:"varint,2,opt,name=vol" json:"vol,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Did60130_M) Reset()         { *m = Did60130_M{} }
func (m *Did60130_M) String() string { return proto.CompactTextString(m) }
func (*Did60130_M) ProtoMessage()    {}

func (m *Did60130_M) GetPrice() float64 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *Did60130_M) GetVol() uint32 {
	if m != nil && m.Vol != nil {
		return *m.Vol
	}
	return 0
}

type Did60132 struct {
	H                *Header  `protobuf:"bytes,1,req,name=h" json:"h,omitempty"`
	Obj              *string  `protobuf:"bytes,2,opt,name=obj" json:"obj,omitempty"`
	Time             *uint64  `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	Bidorask         *bool    `protobuf:"varint,4,opt,name=bidorask" json:"bidorask,omitempty"`
	Price            *float64 `protobuf:"fixed64,5,opt,name=price" json:"price,omitempty"`
	Volume           *uint32  `protobuf:"varint,6,opt,name=volume" json:"volume,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Did60132) Reset()         { *m = Did60132{} }
func (m *Did60132) String() string { return proto.CompactTextString(m) }
func (*Did60132) ProtoMessage()    {}

func (m *Did60132) GetH() *Header {
	if m != nil {
		return m.H
	}
	return nil
}

func (m *Did60132) GetObj() string {
	if m != nil && m.Obj != nil {
		return *m.Obj
	}
	return ""
}

func (m *Did60132) GetTime() uint64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Did60132) GetBidorask() bool {
	if m != nil && m.Bidorask != nil {
		return *m.Bidorask
	}
	return false
}

func (m *Did60132) GetPrice() float64 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *Did60132) GetVolume() uint32 {
	if m != nil && m.Volume != nil {
		return *m.Volume
	}
	return 0
}

type Did60133 struct {
	H                *Header       `protobuf:"bytes,1,req,name=h" json:"h,omitempty"`
	Obj              *string       `protobuf:"bytes,2,opt,name=obj" json:"obj,omitempty"`
	Time             *uint64       `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	M                []*Did60133_M `protobuf:"bytes,4,rep,name=m" json:"m,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Did60133) Reset()         { *m = Did60133{} }
func (m *Did60133) String() string { return proto.CompactTextString(m) }
func (*Did60133) ProtoMessage()    {}

func (m *Did60133) GetH() *Header {
	if m != nil {
		return m.H
	}
	return nil
}

func (m *Did60133) GetObj() string {
	if m != nil && m.Obj != nil {
		return *m.Obj
	}
	return ""
}

func (m *Did60133) GetTime() uint64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Did60133) GetM() []*Did60133_M {
	if m != nil {
		return m.M
	}
	return nil
}

type Did60133_M struct {
	Dir              *bool            `protobuf:"varint,1,opt,name=dir" json:"dir,omitempty"`
	Price            *float64         `protobuf:"fixed64,2,opt,name=price" json:"price,omitempty"`
	O                []*Did60133_M_M1 `protobuf:"bytes,3,rep,name=o" json:"o,omitempty"`
	W                []*Did60133_M_M1 `protobuf:"bytes,4,rep,name=w" json:"w,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Did60133_M) Reset()         { *m = Did60133_M{} }
func (m *Did60133_M) String() string { return proto.CompactTextString(m) }
func (*Did60133_M) ProtoMessage()    {}

func (m *Did60133_M) GetDir() bool {
	if m != nil && m.Dir != nil {
		return *m.Dir
	}
	return false
}

func (m *Did60133_M) GetPrice() float64 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *Did60133_M) GetO() []*Did60133_M_M1 {
	if m != nil {
		return m.O
	}
	return nil
}

func (m *Did60133_M) GetW() []*Did60133_M_M1 {
	if m != nil {
		return m.W
	}
	return nil
}

type Did60133_M_M1 struct {
	Vol              *uint32 `protobuf:"varint,1,opt,name=vol" json:"vol,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Did60133_M_M1) Reset()         { *m = Did60133_M_M1{} }
func (m *Did60133_M_M1) String() string { return proto.CompactTextString(m) }
func (*Did60133_M_M1) ProtoMessage()    {}

func (m *Did60133_M_M1) GetVol() uint32 {
	if m != nil && m.Vol != nil {
		return *m.Vol
	}
	return 0
}

type Did60134 struct {
	H                *Header       `protobuf:"bytes,1,req,name=h" json:"h,omitempty"`
	Obj              *string       `protobuf:"bytes,2,opt,name=obj" json:"obj,omitempty"`
	Time             *uint64       `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	Order            []*Did60134_M `protobuf:"bytes,4,rep,name=order" json:"order,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Did60134) Reset()         { *m = Did60134{} }
func (m *Did60134) String() string { return proto.CompactTextString(m) }
func (*Did60134) ProtoMessage()    {}

func (m *Did60134) GetH() *Header {
	if m != nil {
		return m.H
	}
	return nil
}

func (m *Did60134) GetObj() string {
	if m != nil && m.Obj != nil {
		return *m.Obj
	}
	return ""
}

func (m *Did60134) GetTime() uint64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Did60134) GetOrder() []*Did60134_M {
	if m != nil {
		return m.Order
	}
	return nil
}

type Did60134_M struct {
	Dir              *bool   `protobuf:"varint,1,opt,name=dir" json:"dir,omitempty"`
	Vol              *uint32 `protobuf:"varint,2,opt,name=vol" json:"vol,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Did60134_M) Reset()         { *m = Did60134_M{} }
func (m *Did60134_M) String() string { return proto.CompactTextString(m) }
func (*Did60134_M) ProtoMessage()    {}

func (m *Did60134_M) GetDir() bool {
	if m != nil && m.Dir != nil {
		return *m.Dir
	}
	return false
}

func (m *Did60134_M) GetVol() uint32 {
	if m != nil && m.Vol != nil {
		return *m.Vol
	}
	return 0
}

type Did60135 struct {
	H                *Header       `protobuf:"bytes,1,req,name=h" json:"h,omitempty"`
	Obj              *string       `protobuf:"bytes,2,opt,name=obj" json:"obj,omitempty"`
	Time             *uint64       `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	Buy              []*Did60135_M `protobuf:"bytes,4,rep,name=buy" json:"buy,omitempty"`
	Sell             []*Did60135_M `protobuf:"bytes,5,rep,name=sell" json:"sell,omitempty"`
	Order            []*Did60135_M `protobuf:"bytes,6,rep,name=order" json:"order,omitempty"`
	Wd               []*Did60135_M `protobuf:"bytes,7,rep,name=wd" json:"wd,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Did60135) Reset()         { *m = Did60135{} }
func (m *Did60135) String() string { return proto.CompactTextString(m) }
func (*Did60135) ProtoMessage()    {}

func (m *Did60135) GetH() *Header {
	if m != nil {
		return m.H
	}
	return nil
}

func (m *Did60135) GetObj() string {
	if m != nil && m.Obj != nil {
		return *m.Obj
	}
	return ""
}

func (m *Did60135) GetTime() uint64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Did60135) GetBuy() []*Did60135_M {
	if m != nil {
		return m.Buy
	}
	return nil
}

func (m *Did60135) GetSell() []*Did60135_M {
	if m != nil {
		return m.Sell
	}
	return nil
}

func (m *Did60135) GetOrder() []*Did60135_M {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *Did60135) GetWd() []*Did60135_M {
	if m != nil {
		return m.Wd
	}
	return nil
}

type Did60135_M struct {
	Type             *uint32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Vol              *uint64 `protobuf:"varint,2,opt,name=vol" json:"vol,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Did60135_M) Reset()         { *m = Did60135_M{} }
func (m *Did60135_M) String() string { return proto.CompactTextString(m) }
func (*Did60135_M) ProtoMessage()    {}

func (m *Did60135_M) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Did60135_M) GetVol() uint64 {
	if m != nil && m.Vol != nil {
		return *m.Vol
	}
	return 0
}

func init() {
}
