// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: team.proto

package proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using Team within kubernetes types, where deepcopy-gen is used.
func (in *Team) DeepCopyInto(out *Team) {
	p := proto.Clone(in).(*Team)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Team. Required by controller-gen.
func (in *Team) DeepCopy() *Team {
	if in == nil {
		return nil
	}
	out := new(Team)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Team. Required by controller-gen.
func (in *Team) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Link within kubernetes types, where deepcopy-gen is used.
func (in *Link) DeepCopyInto(out *Link) {
	p := proto.Clone(in).(*Link)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Link. Required by controller-gen.
func (in *Link) DeepCopy() *Link {
	if in == nil {
		return nil
	}
	out := new(Link)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Link. Required by controller-gen.
func (in *Link) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}