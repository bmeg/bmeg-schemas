// Code generated by protoc-gen-go.
// source: nlp.proto
// DO NOT EDIT!

package bmeg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Pubmed struct {
	Pmid     string   `protobuf:"bytes,1,opt,name=pmid" json:"pmid,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Abstract string   `protobuf:"bytes,3,opt,name=abstract" json:"abstract,omitempty"`
	Text     string   `protobuf:"bytes,4,opt,name=text" json:"text,omitempty"`
	Date     string   `protobuf:"bytes,5,opt,name=date" json:"date,omitempty"`
	Author   []string `protobuf:"bytes,6,rep,name=author" json:"author,omitempty"`
	Citation []string `protobuf:"bytes,7,rep,name=citation" json:"citation,omitempty"`
}

func (m *Pubmed) Reset()                    { *m = Pubmed{} }
func (m *Pubmed) String() string            { return proto.CompactTextString(m) }
func (*Pubmed) ProtoMessage()               {}
func (*Pubmed) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Pubmed) GetPmid() string {
	if m != nil {
		return m.Pmid
	}
	return ""
}

func (m *Pubmed) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Pubmed) GetAbstract() string {
	if m != nil {
		return m.Abstract
	}
	return ""
}

func (m *Pubmed) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Pubmed) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Pubmed) GetAuthor() []string {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Pubmed) GetCitation() []string {
	if m != nil {
		return m.Citation
	}
	return nil
}

func init() {
	proto.RegisterType((*Pubmed)(nil), "bmeg.Pubmed")
}

func init() { proto.RegisterFile("nlp.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xce, 0x41, 0xaa, 0xc2, 0x30,
	0x14, 0x85, 0x61, 0xfa, 0x9a, 0xe6, 0xd9, 0x4b, 0x15, 0xc9, 0xe8, 0x0e, 0x8b, 0x23, 0x47, 0x4e,
	0xdc, 0x88, 0x5b, 0x48, 0x9a, 0x8b, 0x06, 0x9a, 0x26, 0xc4, 0x13, 0x70, 0xf9, 0x62, 0xc4, 0xe1,
	0xf7, 0xc3, 0x81, 0x43, 0xe3, 0xb6, 0xe6, 0x4b, 0x2e, 0x09, 0xc9, 0x28, 0x17, 0xe5, 0x7e, 0xaa,
	0xa4, 0x6f, 0xd5, 0x45, 0xf1, 0x66, 0x22, 0x95, 0x63, 0xf0, 0xdc, 0xcd, 0xdd, 0x79, 0x34, 0x7b,
	0x1a, 0x10, 0xb0, 0x0a, 0xff, 0x35, 0x1e, 0x69, 0x67, 0xdd, 0x13, 0xc5, 0x2e, 0xe0, 0xbe, 0x95,
	0x89, 0x14, 0xe4, 0x05, 0x56, 0x3f, 0x79, 0x0b, 0xe1, 0xa1, 0xe9, 0x40, 0xda, 0x56, 0x3c, 0x52,
	0x61, 0x3d, 0xf7, 0xdf, 0xf5, 0x12, 0x60, 0x11, 0xd2, 0xc6, 0xff, 0x9f, 0xe2, 0x74, 0xfb, 0x70,
	0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x5f, 0xc3, 0x89, 0x90, 0x00, 0x00, 0x00,
}
