// Code generated by protoc-gen-go.
// source: ml.proto
// DO NOT EDIT!

package bmeg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SignatureTraining struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Gid  string `protobuf:"bytes,2,opt,name=gid" json:"gid,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	// Specify what the model predicts, i.e. 'drug response'.
	Predicts string `protobuf:"bytes,4,opt,name=predicts" json:"predicts,omitempty"`
	// Specify what phenotype it predicts, i.e. 'AZD6482'.
	Phenotype string `protobuf:"bytes,5,opt,name=phenotype" json:"phenotype,omitempty"`
	// Array for quantile normalization.
	Quantile         []float64          `protobuf:"fixed64,6,rep,packed,name=quantile" json:"quantile,omitempty"`
	Intercept        float64            `protobuf:"fixed64,7,opt,name=intercept" json:"intercept,omitempty"`
	Coefficients     map[string]float64 `protobuf:"bytes,8,rep,name=coefficients" json:"coefficients,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	Background       map[string]float64 `protobuf:"bytes,9,rep,name=background" json:"background,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	BackgroundCohort []string           `protobuf:"bytes,10,rep,name=backgroundCohort" json:"backgroundCohort,omitempty"`
	// Target: Drug
	SignatureFor []string `protobuf:"bytes,11,rep,name=signatureFor" json:"signatureFor,omitempty"`
}

func (m *SignatureTraining) Reset()                    { *m = SignatureTraining{} }
func (m *SignatureTraining) String() string            { return proto.CompactTextString(m) }
func (*SignatureTraining) ProtoMessage()               {}
func (*SignatureTraining) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *SignatureTraining) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SignatureTraining) GetGid() string {
	if m != nil {
		return m.Gid
	}
	return ""
}

func (m *SignatureTraining) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SignatureTraining) GetPredicts() string {
	if m != nil {
		return m.Predicts
	}
	return ""
}

func (m *SignatureTraining) GetPhenotype() string {
	if m != nil {
		return m.Phenotype
	}
	return ""
}

func (m *SignatureTraining) GetQuantile() []float64 {
	if m != nil {
		return m.Quantile
	}
	return nil
}

func (m *SignatureTraining) GetIntercept() float64 {
	if m != nil {
		return m.Intercept
	}
	return 0
}

func (m *SignatureTraining) GetCoefficients() map[string]float64 {
	if m != nil {
		return m.Coefficients
	}
	return nil
}

func (m *SignatureTraining) GetBackground() map[string]float64 {
	if m != nil {
		return m.Background
	}
	return nil
}

func (m *SignatureTraining) GetBackgroundCohort() []string {
	if m != nil {
		return m.BackgroundCohort
	}
	return nil
}

func (m *SignatureTraining) GetSignatureFor() []string {
	if m != nil {
		return m.SignatureFor
	}
	return nil
}

type SignatureExpressionEdge struct {
	Type  string  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	In    string  `protobuf:"bytes,2,opt,name=in" json:"in,omitempty"`
	Out   string  `protobuf:"bytes,3,opt,name=out" json:"out,omitempty"`
	Level float64 `protobuf:"fixed64,4,opt,name=level" json:"level,omitempty"`
}

func (m *SignatureExpressionEdge) Reset()                    { *m = SignatureExpressionEdge{} }
func (m *SignatureExpressionEdge) String() string            { return proto.CompactTextString(m) }
func (*SignatureExpressionEdge) ProtoMessage()               {}
func (*SignatureExpressionEdge) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *SignatureExpressionEdge) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SignatureExpressionEdge) GetIn() string {
	if m != nil {
		return m.In
	}
	return ""
}

func (m *SignatureExpressionEdge) GetOut() string {
	if m != nil {
		return m.Out
	}
	return ""
}

func (m *SignatureExpressionEdge) GetLevel() float64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*SignatureTraining)(nil), "bmeg.SignatureTraining")
	proto.RegisterType((*SignatureExpressionEdge)(nil), "bmeg.SignatureExpressionEdge")
}

func init() { proto.RegisterFile("ml.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0xb3, 0x2c, 0x20, 0x9d, 0x42, 0x84, 0x8d, 0x89, 0x1b, 0x4e, 0x0d, 0x17, 0xeb, 0xa5,
	0x89, 0x72, 0x31, 0x7a, 0x30, 0x91, 0xe0, 0xd5, 0x83, 0xbe, 0x40, 0x69, 0x87, 0x65, 0x43, 0xd9,
	0x5d, 0xb7, 0x5b, 0x22, 0xaf, 0xe2, 0xd3, 0x9a, 0x2e, 0x50, 0x8d, 0xc4, 0x78, 0x9c, 0xc9, 0xff,
	0x65, 0xfe, 0xcc, 0x07, 0xbd, 0x4d, 0x91, 0x18, 0xab, 0x9d, 0x66, 0xed, 0xc5, 0x06, 0xc5, 0xe4,
	0x93, 0xc2, 0xe8, 0x55, 0x0a, 0x95, 0xba, 0xca, 0xe2, 0x9b, 0x4d, 0xa5, 0x92, 0x4a, 0x30, 0x80,
	0x96, 0xcc, 0x39, 0x89, 0x48, 0x1c, 0xb0, 0x10, 0xa8, 0x90, 0x39, 0x6f, 0xf9, 0xa1, 0x0f, 0x6d,
	0xb7, 0x33, 0xc8, 0xa9, 0x9f, 0x86, 0xd0, 0x33, 0x16, 0x73, 0x99, 0xb9, 0x92, 0xb7, 0xfd, 0x66,
	0x04, 0x81, 0x59, 0xa1, 0xd2, 0x3e, 0xd4, 0x39, 0x86, 0xde, 0xab, 0x54, 0x39, 0x59, 0x20, 0xef,
	0x46, 0x34, 0x26, 0x75, 0x48, 0x2a, 0x87, 0x36, 0x43, 0xe3, 0xf8, 0x59, 0x44, 0x62, 0xc2, 0x1e,
	0xa1, 0x9f, 0x69, 0x5c, 0x2e, 0x65, 0x26, 0x51, 0xb9, 0x92, 0xf7, 0x22, 0x1a, 0x87, 0xb7, 0xd7,
	0x49, 0xdd, 0x31, 0x39, 0xe9, 0x97, 0xcc, 0x7e, 0x64, 0xe7, 0xca, 0xd9, 0x1d, 0x7b, 0x00, 0x58,
	0xa4, 0xd9, 0x5a, 0x58, 0x5d, 0xa9, 0x9c, 0x07, 0x1e, 0xbf, 0xfa, 0x0b, 0x7f, 0x6a, 0x92, 0x7b,
	0x98, 0xc3, 0xf0, 0x1b, 0x9e, 0xe9, 0x95, 0xb6, 0x8e, 0x43, 0x44, 0xe3, 0x80, 0x5d, 0x40, 0xbf,
	0x3c, 0xe2, 0xcf, 0xda, 0xf2, 0xb0, 0xde, 0x8e, 0xa7, 0x30, 0x3a, 0x6d, 0x10, 0x02, 0x5d, 0xe3,
	0xee, 0xf0, 0xb4, 0x01, 0x74, 0xb6, 0x69, 0x51, 0xa1, 0x7f, 0x1b, 0xb9, 0x6f, 0xdd, 0x91, 0xf1,
	0x0d, 0x9c, 0xff, 0xbe, 0xfb, 0x0f, 0x32, 0x79, 0x81, 0xcb, 0xa6, 0xfc, 0xfc, 0xc3, 0x58, 0x2c,
	0x4b, 0xa9, 0xd5, 0x3c, 0x17, 0xd8, 0x88, 0xd8, 0xb3, 0xb5, 0x2f, 0x75, 0x50, 0x14, 0x02, 0xd5,
	0x95, 0x3b, 0x18, 0x1a, 0x40, 0xa7, 0xc0, 0x2d, 0x16, 0x5e, 0x0f, 0x59, 0x74, 0xbd, 0xfa, 0xe9,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x0d, 0x50, 0xa3, 0x06, 0x02, 0x00, 0x00,
}
