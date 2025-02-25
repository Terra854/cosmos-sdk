package distributionv1beta1

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/query/v1beta1"
	v1beta11 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_QueryParamsRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryParamsRequest = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryParamsRequest")
}

var _ protoreflect.Message = (*fastReflection_QueryParamsRequest)(nil)

type fastReflection_QueryParamsRequest QueryParamsRequest

func (x *QueryParamsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryParamsRequest)(x)
}

func (x *QueryParamsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryParamsRequest_messageType fastReflection_QueryParamsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryParamsRequest_messageType{}

type fastReflection_QueryParamsRequest_messageType struct{}

func (x fastReflection_QueryParamsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryParamsRequest)(nil)
}
func (x fastReflection_QueryParamsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryParamsRequest)
}
func (x fastReflection_QueryParamsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryParamsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryParamsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryParamsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryParamsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryParamsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryParamsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryParamsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryParamsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryParamsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryParamsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryParamsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryParamsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryParamsRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryParamsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryParamsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryParamsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryParamsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryParamsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryParamsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryParamsRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryParamsRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryParamsRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryParamsResponse        protoreflect.MessageDescriptor
	fd_QueryParamsResponse_params protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryParamsResponse = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryParamsResponse")
	fd_QueryParamsResponse_params = md_QueryParamsResponse.Fields().ByName("params")
}

var _ protoreflect.Message = (*fastReflection_QueryParamsResponse)(nil)

type fastReflection_QueryParamsResponse QueryParamsResponse

func (x *QueryParamsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryParamsResponse)(x)
}

func (x *QueryParamsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryParamsResponse_messageType fastReflection_QueryParamsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryParamsResponse_messageType{}

type fastReflection_QueryParamsResponse_messageType struct{}

func (x fastReflection_QueryParamsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryParamsResponse)(nil)
}
func (x fastReflection_QueryParamsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryParamsResponse)
}
func (x fastReflection_QueryParamsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryParamsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryParamsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryParamsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryParamsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryParamsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryParamsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryParamsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryParamsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryParamsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryParamsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Params != nil {
		value := protoreflect.ValueOfMessage(x.Params.ProtoReflect())
		if !f(fd_QueryParamsResponse_params, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryParamsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryParamsResponse.params":
		return x.Params != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryParamsResponse.params":
		x.Params = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryParamsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryParamsResponse.params":
		value := x.Params
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryParamsResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryParamsResponse.params":
		x.Params = value.Message().Interface().(*Params)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryParamsResponse.params":
		if x.Params == nil {
			x.Params = new(Params)
		}
		return protoreflect.ValueOfMessage(x.Params.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryParamsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryParamsResponse.params":
		m := new(Params)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryParamsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryParamsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryParamsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryParamsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryParamsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryParamsResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Params != nil {
			l = options.Size(x.Params)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryParamsResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Params != nil {
			encoded, err := options.Marshal(x.Params)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryParamsResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Params == nil {
					x.Params = &Params{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Params); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryValidatorOutstandingRewardsRequest                   protoreflect.MessageDescriptor
	fd_QueryValidatorOutstandingRewardsRequest_validator_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryValidatorOutstandingRewardsRequest = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryValidatorOutstandingRewardsRequest")
	fd_QueryValidatorOutstandingRewardsRequest_validator_address = md_QueryValidatorOutstandingRewardsRequest.Fields().ByName("validator_address")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorOutstandingRewardsRequest)(nil)

type fastReflection_QueryValidatorOutstandingRewardsRequest QueryValidatorOutstandingRewardsRequest

func (x *QueryValidatorOutstandingRewardsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorOutstandingRewardsRequest)(x)
}

func (x *QueryValidatorOutstandingRewardsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorOutstandingRewardsRequest_messageType fastReflection_QueryValidatorOutstandingRewardsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorOutstandingRewardsRequest_messageType{}

type fastReflection_QueryValidatorOutstandingRewardsRequest_messageType struct{}

func (x fastReflection_QueryValidatorOutstandingRewardsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorOutstandingRewardsRequest)(nil)
}
func (x fastReflection_QueryValidatorOutstandingRewardsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorOutstandingRewardsRequest)
}
func (x fastReflection_QueryValidatorOutstandingRewardsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorOutstandingRewardsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorOutstandingRewardsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorOutstandingRewardsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorOutstandingRewardsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorOutstandingRewardsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_QueryValidatorOutstandingRewardsRequest_validator_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest.validator_address":
		return x.ValidatorAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest.validator_address":
		x.ValidatorAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest.validator_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorOutstandingRewardsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorOutstandingRewardsRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ValidatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryValidatorOutstandingRewardsRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.ValidatorAddress) > 0 {
			i -= len(x.ValidatorAddress)
			copy(dAtA[i:], x.ValidatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryValidatorOutstandingRewardsRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorOutstandingRewardsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorOutstandingRewardsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryValidatorOutstandingRewardsResponse         protoreflect.MessageDescriptor
	fd_QueryValidatorOutstandingRewardsResponse_rewards protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryValidatorOutstandingRewardsResponse = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryValidatorOutstandingRewardsResponse")
	fd_QueryValidatorOutstandingRewardsResponse_rewards = md_QueryValidatorOutstandingRewardsResponse.Fields().ByName("rewards")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorOutstandingRewardsResponse)(nil)

type fastReflection_QueryValidatorOutstandingRewardsResponse QueryValidatorOutstandingRewardsResponse

func (x *QueryValidatorOutstandingRewardsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorOutstandingRewardsResponse)(x)
}

func (x *QueryValidatorOutstandingRewardsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorOutstandingRewardsResponse_messageType fastReflection_QueryValidatorOutstandingRewardsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorOutstandingRewardsResponse_messageType{}

type fastReflection_QueryValidatorOutstandingRewardsResponse_messageType struct{}

func (x fastReflection_QueryValidatorOutstandingRewardsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorOutstandingRewardsResponse)(nil)
}
func (x fastReflection_QueryValidatorOutstandingRewardsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorOutstandingRewardsResponse)
}
func (x fastReflection_QueryValidatorOutstandingRewardsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorOutstandingRewardsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorOutstandingRewardsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorOutstandingRewardsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorOutstandingRewardsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorOutstandingRewardsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Rewards != nil {
		value := protoreflect.ValueOfMessage(x.Rewards.ProtoReflect())
		if !f(fd_QueryValidatorOutstandingRewardsResponse_rewards, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse.rewards":
		return x.Rewards != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse.rewards":
		x.Rewards = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse.rewards":
		value := x.Rewards
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse.rewards":
		x.Rewards = value.Message().Interface().(*ValidatorOutstandingRewards)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse.rewards":
		if x.Rewards == nil {
			x.Rewards = new(ValidatorOutstandingRewards)
		}
		return protoreflect.ValueOfMessage(x.Rewards.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse.rewards":
		m := new(ValidatorOutstandingRewards)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorOutstandingRewardsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorOutstandingRewardsResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Rewards != nil {
			l = options.Size(x.Rewards)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryValidatorOutstandingRewardsResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Rewards != nil {
			encoded, err := options.Marshal(x.Rewards)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryValidatorOutstandingRewardsResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorOutstandingRewardsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorOutstandingRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Rewards == nil {
					x.Rewards = &ValidatorOutstandingRewards{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Rewards); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryValidatorCommissionRequest                   protoreflect.MessageDescriptor
	fd_QueryValidatorCommissionRequest_validator_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryValidatorCommissionRequest = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryValidatorCommissionRequest")
	fd_QueryValidatorCommissionRequest_validator_address = md_QueryValidatorCommissionRequest.Fields().ByName("validator_address")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorCommissionRequest)(nil)

type fastReflection_QueryValidatorCommissionRequest QueryValidatorCommissionRequest

func (x *QueryValidatorCommissionRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorCommissionRequest)(x)
}

func (x *QueryValidatorCommissionRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorCommissionRequest_messageType fastReflection_QueryValidatorCommissionRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorCommissionRequest_messageType{}

type fastReflection_QueryValidatorCommissionRequest_messageType struct{}

func (x fastReflection_QueryValidatorCommissionRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorCommissionRequest)(nil)
}
func (x fastReflection_QueryValidatorCommissionRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorCommissionRequest)
}
func (x fastReflection_QueryValidatorCommissionRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorCommissionRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorCommissionRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorCommissionRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorCommissionRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorCommissionRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorCommissionRequest) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorCommissionRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorCommissionRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorCommissionRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorCommissionRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_QueryValidatorCommissionRequest_validator_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryValidatorCommissionRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorCommissionRequest.validator_address":
		return x.ValidatorAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorCommissionRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorCommissionRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorCommissionRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorCommissionRequest.validator_address":
		x.ValidatorAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorCommissionRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorCommissionRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorCommissionRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorCommissionRequest.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorCommissionRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorCommissionRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorCommissionRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorCommissionRequest.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorCommissionRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorCommissionRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorCommissionRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorCommissionRequest.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.distribution.v1beta1.QueryValidatorCommissionRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorCommissionRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorCommissionRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorCommissionRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorCommissionRequest.validator_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorCommissionRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorCommissionRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorCommissionRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryValidatorCommissionRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorCommissionRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorCommissionRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryValidatorCommissionRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorCommissionRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorCommissionRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ValidatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryValidatorCommissionRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.ValidatorAddress) > 0 {
			i -= len(x.ValidatorAddress)
			copy(dAtA[i:], x.ValidatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryValidatorCommissionRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorCommissionRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorCommissionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryValidatorCommissionResponse            protoreflect.MessageDescriptor
	fd_QueryValidatorCommissionResponse_commission protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryValidatorCommissionResponse = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryValidatorCommissionResponse")
	fd_QueryValidatorCommissionResponse_commission = md_QueryValidatorCommissionResponse.Fields().ByName("commission")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorCommissionResponse)(nil)

type fastReflection_QueryValidatorCommissionResponse QueryValidatorCommissionResponse

func (x *QueryValidatorCommissionResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorCommissionResponse)(x)
}

func (x *QueryValidatorCommissionResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorCommissionResponse_messageType fastReflection_QueryValidatorCommissionResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorCommissionResponse_messageType{}

type fastReflection_QueryValidatorCommissionResponse_messageType struct{}

func (x fastReflection_QueryValidatorCommissionResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorCommissionResponse)(nil)
}
func (x fastReflection_QueryValidatorCommissionResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorCommissionResponse)
}
func (x fastReflection_QueryValidatorCommissionResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorCommissionResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorCommissionResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorCommissionResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorCommissionResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorCommissionResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorCommissionResponse) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorCommissionResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorCommissionResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorCommissionResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorCommissionResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Commission != nil {
		value := protoreflect.ValueOfMessage(x.Commission.ProtoReflect())
		if !f(fd_QueryValidatorCommissionResponse_commission, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryValidatorCommissionResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorCommissionResponse.commission":
		return x.Commission != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorCommissionResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorCommissionResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorCommissionResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorCommissionResponse.commission":
		x.Commission = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorCommissionResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorCommissionResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorCommissionResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorCommissionResponse.commission":
		value := x.Commission
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorCommissionResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorCommissionResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorCommissionResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorCommissionResponse.commission":
		x.Commission = value.Message().Interface().(*ValidatorAccumulatedCommission)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorCommissionResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorCommissionResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorCommissionResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorCommissionResponse.commission":
		if x.Commission == nil {
			x.Commission = new(ValidatorAccumulatedCommission)
		}
		return protoreflect.ValueOfMessage(x.Commission.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorCommissionResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorCommissionResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorCommissionResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorCommissionResponse.commission":
		m := new(ValidatorAccumulatedCommission)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorCommissionResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorCommissionResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorCommissionResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryValidatorCommissionResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorCommissionResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorCommissionResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryValidatorCommissionResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorCommissionResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorCommissionResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Commission != nil {
			l = options.Size(x.Commission)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryValidatorCommissionResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Commission != nil {
			encoded, err := options.Marshal(x.Commission)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryValidatorCommissionResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorCommissionResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorCommissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Commission == nil {
					x.Commission = &ValidatorAccumulatedCommission{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Commission); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryValidatorSlashesRequest                   protoreflect.MessageDescriptor
	fd_QueryValidatorSlashesRequest_validator_address protoreflect.FieldDescriptor
	fd_QueryValidatorSlashesRequest_starting_height   protoreflect.FieldDescriptor
	fd_QueryValidatorSlashesRequest_ending_height     protoreflect.FieldDescriptor
	fd_QueryValidatorSlashesRequest_pagination        protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryValidatorSlashesRequest = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryValidatorSlashesRequest")
	fd_QueryValidatorSlashesRequest_validator_address = md_QueryValidatorSlashesRequest.Fields().ByName("validator_address")
	fd_QueryValidatorSlashesRequest_starting_height = md_QueryValidatorSlashesRequest.Fields().ByName("starting_height")
	fd_QueryValidatorSlashesRequest_ending_height = md_QueryValidatorSlashesRequest.Fields().ByName("ending_height")
	fd_QueryValidatorSlashesRequest_pagination = md_QueryValidatorSlashesRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorSlashesRequest)(nil)

type fastReflection_QueryValidatorSlashesRequest QueryValidatorSlashesRequest

func (x *QueryValidatorSlashesRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorSlashesRequest)(x)
}

func (x *QueryValidatorSlashesRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorSlashesRequest_messageType fastReflection_QueryValidatorSlashesRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorSlashesRequest_messageType{}

type fastReflection_QueryValidatorSlashesRequest_messageType struct{}

func (x fastReflection_QueryValidatorSlashesRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorSlashesRequest)(nil)
}
func (x fastReflection_QueryValidatorSlashesRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorSlashesRequest)
}
func (x fastReflection_QueryValidatorSlashesRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorSlashesRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorSlashesRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorSlashesRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorSlashesRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorSlashesRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorSlashesRequest) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorSlashesRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorSlashesRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorSlashesRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorSlashesRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_QueryValidatorSlashesRequest_validator_address, value) {
			return
		}
	}
	if x.StartingHeight != uint64(0) {
		value := protoreflect.ValueOfUint64(x.StartingHeight)
		if !f(fd_QueryValidatorSlashesRequest_starting_height, value) {
			return
		}
	}
	if x.EndingHeight != uint64(0) {
		value := protoreflect.ValueOfUint64(x.EndingHeight)
		if !f(fd_QueryValidatorSlashesRequest_ending_height, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryValidatorSlashesRequest_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryValidatorSlashesRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.validator_address":
		return x.ValidatorAddress != ""
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.starting_height":
		return x.StartingHeight != uint64(0)
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.ending_height":
		return x.EndingHeight != uint64(0)
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorSlashesRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorSlashesRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorSlashesRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.validator_address":
		x.ValidatorAddress = ""
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.starting_height":
		x.StartingHeight = uint64(0)
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.ending_height":
		x.EndingHeight = uint64(0)
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorSlashesRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorSlashesRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorSlashesRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.starting_height":
		value := x.StartingHeight
		return protoreflect.ValueOfUint64(value)
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.ending_height":
		value := x.EndingHeight
		return protoreflect.ValueOfUint64(value)
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorSlashesRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorSlashesRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorSlashesRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.starting_height":
		x.StartingHeight = value.Uint()
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.ending_height":
		x.EndingHeight = value.Uint()
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorSlashesRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorSlashesRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorSlashesRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.distribution.v1beta1.QueryValidatorSlashesRequest is not mutable"))
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.starting_height":
		panic(fmt.Errorf("field starting_height of message cosmos.distribution.v1beta1.QueryValidatorSlashesRequest is not mutable"))
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.ending_height":
		panic(fmt.Errorf("field ending_height of message cosmos.distribution.v1beta1.QueryValidatorSlashesRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorSlashesRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorSlashesRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorSlashesRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.validator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.starting_height":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.ending_height":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorSlashesRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorSlashesRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorSlashesRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryValidatorSlashesRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorSlashesRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorSlashesRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryValidatorSlashesRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorSlashesRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorSlashesRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ValidatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.StartingHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.StartingHeight))
		}
		if x.EndingHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.EndingHeight))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryValidatorSlashesRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if x.EndingHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.EndingHeight))
			i--
			dAtA[i] = 0x18
		}
		if x.StartingHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.StartingHeight))
			i--
			dAtA[i] = 0x10
		}
		if len(x.ValidatorAddress) > 0 {
			i -= len(x.ValidatorAddress)
			copy(dAtA[i:], x.ValidatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryValidatorSlashesRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorSlashesRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorSlashesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartingHeight", wireType)
				}
				x.StartingHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.StartingHeight |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EndingHeight", wireType)
				}
				x.EndingHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.EndingHeight |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_QueryValidatorSlashesResponse_1_list)(nil)

type _QueryValidatorSlashesResponse_1_list struct {
	list *[]*ValidatorSlashEvent
}

func (x *_QueryValidatorSlashesResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryValidatorSlashesResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryValidatorSlashesResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorSlashEvent)
	(*x.list)[i] = concreteValue
}

func (x *_QueryValidatorSlashesResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorSlashEvent)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryValidatorSlashesResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(ValidatorSlashEvent)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryValidatorSlashesResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryValidatorSlashesResponse_1_list) NewElement() protoreflect.Value {
	v := new(ValidatorSlashEvent)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryValidatorSlashesResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryValidatorSlashesResponse            protoreflect.MessageDescriptor
	fd_QueryValidatorSlashesResponse_slashes    protoreflect.FieldDescriptor
	fd_QueryValidatorSlashesResponse_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryValidatorSlashesResponse = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryValidatorSlashesResponse")
	fd_QueryValidatorSlashesResponse_slashes = md_QueryValidatorSlashesResponse.Fields().ByName("slashes")
	fd_QueryValidatorSlashesResponse_pagination = md_QueryValidatorSlashesResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorSlashesResponse)(nil)

type fastReflection_QueryValidatorSlashesResponse QueryValidatorSlashesResponse

func (x *QueryValidatorSlashesResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorSlashesResponse)(x)
}

func (x *QueryValidatorSlashesResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorSlashesResponse_messageType fastReflection_QueryValidatorSlashesResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorSlashesResponse_messageType{}

type fastReflection_QueryValidatorSlashesResponse_messageType struct{}

func (x fastReflection_QueryValidatorSlashesResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorSlashesResponse)(nil)
}
func (x fastReflection_QueryValidatorSlashesResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorSlashesResponse)
}
func (x fastReflection_QueryValidatorSlashesResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorSlashesResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorSlashesResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorSlashesResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorSlashesResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorSlashesResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorSlashesResponse) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorSlashesResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorSlashesResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorSlashesResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorSlashesResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Slashes) != 0 {
		value := protoreflect.ValueOfList(&_QueryValidatorSlashesResponse_1_list{list: &x.Slashes})
		if !f(fd_QueryValidatorSlashesResponse_slashes, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryValidatorSlashesResponse_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryValidatorSlashesResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.slashes":
		return len(x.Slashes) != 0
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorSlashesResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorSlashesResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorSlashesResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.slashes":
		x.Slashes = nil
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorSlashesResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorSlashesResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorSlashesResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.slashes":
		if len(x.Slashes) == 0 {
			return protoreflect.ValueOfList(&_QueryValidatorSlashesResponse_1_list{})
		}
		listValue := &_QueryValidatorSlashesResponse_1_list{list: &x.Slashes}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorSlashesResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorSlashesResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorSlashesResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.slashes":
		lv := value.List()
		clv := lv.(*_QueryValidatorSlashesResponse_1_list)
		x.Slashes = *clv.list
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorSlashesResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorSlashesResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorSlashesResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.slashes":
		if x.Slashes == nil {
			x.Slashes = []*ValidatorSlashEvent{}
		}
		value := &_QueryValidatorSlashesResponse_1_list{list: &x.Slashes}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorSlashesResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorSlashesResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorSlashesResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.slashes":
		list := []*ValidatorSlashEvent{}
		return protoreflect.ValueOfList(&_QueryValidatorSlashesResponse_1_list{list: &list})
	case "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryValidatorSlashesResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryValidatorSlashesResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorSlashesResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryValidatorSlashesResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorSlashesResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorSlashesResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryValidatorSlashesResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorSlashesResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorSlashesResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Slashes) > 0 {
			for _, e := range x.Slashes {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryValidatorSlashesResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Slashes) > 0 {
			for iNdEx := len(x.Slashes) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Slashes[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryValidatorSlashesResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorSlashesResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorSlashesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Slashes", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Slashes = append(x.Slashes, &ValidatorSlashEvent{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Slashes[len(x.Slashes)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryDelegationRewardsRequest                   protoreflect.MessageDescriptor
	fd_QueryDelegationRewardsRequest_delegator_address protoreflect.FieldDescriptor
	fd_QueryDelegationRewardsRequest_validator_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryDelegationRewardsRequest = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryDelegationRewardsRequest")
	fd_QueryDelegationRewardsRequest_delegator_address = md_QueryDelegationRewardsRequest.Fields().ByName("delegator_address")
	fd_QueryDelegationRewardsRequest_validator_address = md_QueryDelegationRewardsRequest.Fields().ByName("validator_address")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegationRewardsRequest)(nil)

type fastReflection_QueryDelegationRewardsRequest QueryDelegationRewardsRequest

func (x *QueryDelegationRewardsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegationRewardsRequest)(x)
}

func (x *QueryDelegationRewardsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegationRewardsRequest_messageType fastReflection_QueryDelegationRewardsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegationRewardsRequest_messageType{}

type fastReflection_QueryDelegationRewardsRequest_messageType struct{}

func (x fastReflection_QueryDelegationRewardsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegationRewardsRequest)(nil)
}
func (x fastReflection_QueryDelegationRewardsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegationRewardsRequest)
}
func (x fastReflection_QueryDelegationRewardsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegationRewardsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegationRewardsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegationRewardsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegationRewardsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegationRewardsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegationRewardsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryDelegationRewardsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegationRewardsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegationRewardsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegationRewardsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddress != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddress)
		if !f(fd_QueryDelegationRewardsRequest_delegator_address, value) {
			return
		}
	}
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_QueryDelegationRewardsRequest_validator_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryDelegationRewardsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest.delegator_address":
		return x.DelegatorAddress != ""
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest.validator_address":
		return x.ValidatorAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationRewardsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest.delegator_address":
		x.DelegatorAddress = ""
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest.validator_address":
		x.ValidatorAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegationRewardsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest.delegator_address":
		value := x.DelegatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationRewardsRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationRewardsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest.delegator_address":
		x.DelegatorAddress = value.Interface().(string)
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationRewardsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest.delegator_address":
		panic(fmt.Errorf("field delegator_address of message cosmos.distribution.v1beta1.QueryDelegationRewardsRequest is not mutable"))
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.distribution.v1beta1.QueryDelegationRewardsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegationRewardsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest.delegator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest.validator_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegationRewardsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryDelegationRewardsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegationRewardsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationRewardsRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryDelegationRewardsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegationRewardsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegationRewardsRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.DelegatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegationRewardsRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.ValidatorAddress) > 0 {
			i -= len(x.ValidatorAddress)
			copy(dAtA[i:], x.ValidatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddress) > 0 {
			i -= len(x.DelegatorAddress)
			copy(dAtA[i:], x.DelegatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegationRewardsRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegationRewardsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegationRewardsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_QueryDelegationRewardsResponse_1_list)(nil)

type _QueryDelegationRewardsResponse_1_list struct {
	list *[]*v1beta11.DecCoin
}

func (x *_QueryDelegationRewardsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryDelegationRewardsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryDelegationRewardsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta11.DecCoin)
	(*x.list)[i] = concreteValue
}

func (x *_QueryDelegationRewardsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta11.DecCoin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryDelegationRewardsResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta11.DecCoin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryDelegationRewardsResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryDelegationRewardsResponse_1_list) NewElement() protoreflect.Value {
	v := new(v1beta11.DecCoin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryDelegationRewardsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryDelegationRewardsResponse         protoreflect.MessageDescriptor
	fd_QueryDelegationRewardsResponse_rewards protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryDelegationRewardsResponse = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryDelegationRewardsResponse")
	fd_QueryDelegationRewardsResponse_rewards = md_QueryDelegationRewardsResponse.Fields().ByName("rewards")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegationRewardsResponse)(nil)

type fastReflection_QueryDelegationRewardsResponse QueryDelegationRewardsResponse

func (x *QueryDelegationRewardsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegationRewardsResponse)(x)
}

func (x *QueryDelegationRewardsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegationRewardsResponse_messageType fastReflection_QueryDelegationRewardsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegationRewardsResponse_messageType{}

type fastReflection_QueryDelegationRewardsResponse_messageType struct{}

func (x fastReflection_QueryDelegationRewardsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegationRewardsResponse)(nil)
}
func (x fastReflection_QueryDelegationRewardsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegationRewardsResponse)
}
func (x fastReflection_QueryDelegationRewardsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegationRewardsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegationRewardsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegationRewardsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegationRewardsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegationRewardsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegationRewardsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryDelegationRewardsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegationRewardsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegationRewardsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegationRewardsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Rewards) != 0 {
		value := protoreflect.ValueOfList(&_QueryDelegationRewardsResponse_1_list{list: &x.Rewards})
		if !f(fd_QueryDelegationRewardsResponse_rewards, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryDelegationRewardsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsResponse.rewards":
		return len(x.Rewards) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationRewardsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsResponse.rewards":
		x.Rewards = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegationRewardsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsResponse.rewards":
		if len(x.Rewards) == 0 {
			return protoreflect.ValueOfList(&_QueryDelegationRewardsResponse_1_list{})
		}
		listValue := &_QueryDelegationRewardsResponse_1_list{list: &x.Rewards}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationRewardsResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationRewardsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsResponse.rewards":
		lv := value.List()
		clv := lv.(*_QueryDelegationRewardsResponse_1_list)
		x.Rewards = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationRewardsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsResponse.rewards":
		if x.Rewards == nil {
			x.Rewards = []*v1beta11.DecCoin{}
		}
		value := &_QueryDelegationRewardsResponse_1_list{list: &x.Rewards}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegationRewardsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationRewardsResponse.rewards":
		list := []*v1beta11.DecCoin{}
		return protoreflect.ValueOfList(&_QueryDelegationRewardsResponse_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegationRewardsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryDelegationRewardsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegationRewardsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationRewardsResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryDelegationRewardsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegationRewardsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegationRewardsResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Rewards) > 0 {
			for _, e := range x.Rewards {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegationRewardsResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Rewards) > 0 {
			for iNdEx := len(x.Rewards) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Rewards[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegationRewardsResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegationRewardsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegationRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Rewards = append(x.Rewards, &v1beta11.DecCoin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Rewards[len(x.Rewards)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryDelegationTotalRewardsRequest                   protoreflect.MessageDescriptor
	fd_QueryDelegationTotalRewardsRequest_delegator_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryDelegationTotalRewardsRequest = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryDelegationTotalRewardsRequest")
	fd_QueryDelegationTotalRewardsRequest_delegator_address = md_QueryDelegationTotalRewardsRequest.Fields().ByName("delegator_address")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegationTotalRewardsRequest)(nil)

type fastReflection_QueryDelegationTotalRewardsRequest QueryDelegationTotalRewardsRequest

func (x *QueryDelegationTotalRewardsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegationTotalRewardsRequest)(x)
}

func (x *QueryDelegationTotalRewardsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegationTotalRewardsRequest_messageType fastReflection_QueryDelegationTotalRewardsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegationTotalRewardsRequest_messageType{}

type fastReflection_QueryDelegationTotalRewardsRequest_messageType struct{}

func (x fastReflection_QueryDelegationTotalRewardsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegationTotalRewardsRequest)(nil)
}
func (x fastReflection_QueryDelegationTotalRewardsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegationTotalRewardsRequest)
}
func (x fastReflection_QueryDelegationTotalRewardsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegationTotalRewardsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegationTotalRewardsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegationTotalRewardsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryDelegationTotalRewardsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegationTotalRewardsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddress != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddress)
		if !f(fd_QueryDelegationTotalRewardsRequest_delegator_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest.delegator_address":
		return x.DelegatorAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest.delegator_address":
		x.DelegatorAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest.delegator_address":
		value := x.DelegatorAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest.delegator_address":
		x.DelegatorAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest.delegator_address":
		panic(fmt.Errorf("field delegator_address of message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest.delegator_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegationTotalRewardsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegationTotalRewardsRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.DelegatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegationTotalRewardsRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.DelegatorAddress) > 0 {
			i -= len(x.DelegatorAddress)
			copy(dAtA[i:], x.DelegatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegationTotalRewardsRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegationTotalRewardsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegationTotalRewardsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_QueryDelegationTotalRewardsResponse_1_list)(nil)

type _QueryDelegationTotalRewardsResponse_1_list struct {
	list *[]*DelegationDelegatorReward
}

func (x *_QueryDelegationTotalRewardsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryDelegationTotalRewardsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryDelegationTotalRewardsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DelegationDelegatorReward)
	(*x.list)[i] = concreteValue
}

func (x *_QueryDelegationTotalRewardsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DelegationDelegatorReward)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryDelegationTotalRewardsResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(DelegationDelegatorReward)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryDelegationTotalRewardsResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryDelegationTotalRewardsResponse_1_list) NewElement() protoreflect.Value {
	v := new(DelegationDelegatorReward)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryDelegationTotalRewardsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_QueryDelegationTotalRewardsResponse_2_list)(nil)

type _QueryDelegationTotalRewardsResponse_2_list struct {
	list *[]*v1beta11.DecCoin
}

func (x *_QueryDelegationTotalRewardsResponse_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryDelegationTotalRewardsResponse_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryDelegationTotalRewardsResponse_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta11.DecCoin)
	(*x.list)[i] = concreteValue
}

func (x *_QueryDelegationTotalRewardsResponse_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta11.DecCoin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryDelegationTotalRewardsResponse_2_list) AppendMutable() protoreflect.Value {
	v := new(v1beta11.DecCoin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryDelegationTotalRewardsResponse_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryDelegationTotalRewardsResponse_2_list) NewElement() protoreflect.Value {
	v := new(v1beta11.DecCoin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryDelegationTotalRewardsResponse_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryDelegationTotalRewardsResponse         protoreflect.MessageDescriptor
	fd_QueryDelegationTotalRewardsResponse_rewards protoreflect.FieldDescriptor
	fd_QueryDelegationTotalRewardsResponse_total   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryDelegationTotalRewardsResponse = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryDelegationTotalRewardsResponse")
	fd_QueryDelegationTotalRewardsResponse_rewards = md_QueryDelegationTotalRewardsResponse.Fields().ByName("rewards")
	fd_QueryDelegationTotalRewardsResponse_total = md_QueryDelegationTotalRewardsResponse.Fields().ByName("total")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegationTotalRewardsResponse)(nil)

type fastReflection_QueryDelegationTotalRewardsResponse QueryDelegationTotalRewardsResponse

func (x *QueryDelegationTotalRewardsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegationTotalRewardsResponse)(x)
}

func (x *QueryDelegationTotalRewardsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegationTotalRewardsResponse_messageType fastReflection_QueryDelegationTotalRewardsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegationTotalRewardsResponse_messageType{}

type fastReflection_QueryDelegationTotalRewardsResponse_messageType struct{}

func (x fastReflection_QueryDelegationTotalRewardsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegationTotalRewardsResponse)(nil)
}
func (x fastReflection_QueryDelegationTotalRewardsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegationTotalRewardsResponse)
}
func (x fastReflection_QueryDelegationTotalRewardsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegationTotalRewardsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegationTotalRewardsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegationTotalRewardsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryDelegationTotalRewardsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegationTotalRewardsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Rewards) != 0 {
		value := protoreflect.ValueOfList(&_QueryDelegationTotalRewardsResponse_1_list{list: &x.Rewards})
		if !f(fd_QueryDelegationTotalRewardsResponse_rewards, value) {
			return
		}
	}
	if len(x.Total) != 0 {
		value := protoreflect.ValueOfList(&_QueryDelegationTotalRewardsResponse_2_list{list: &x.Total})
		if !f(fd_QueryDelegationTotalRewardsResponse_total, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.rewards":
		return len(x.Rewards) != 0
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.total":
		return len(x.Total) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.rewards":
		x.Rewards = nil
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.total":
		x.Total = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.rewards":
		if len(x.Rewards) == 0 {
			return protoreflect.ValueOfList(&_QueryDelegationTotalRewardsResponse_1_list{})
		}
		listValue := &_QueryDelegationTotalRewardsResponse_1_list{list: &x.Rewards}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.total":
		if len(x.Total) == 0 {
			return protoreflect.ValueOfList(&_QueryDelegationTotalRewardsResponse_2_list{})
		}
		listValue := &_QueryDelegationTotalRewardsResponse_2_list{list: &x.Total}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.rewards":
		lv := value.List()
		clv := lv.(*_QueryDelegationTotalRewardsResponse_1_list)
		x.Rewards = *clv.list
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.total":
		lv := value.List()
		clv := lv.(*_QueryDelegationTotalRewardsResponse_2_list)
		x.Total = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.rewards":
		if x.Rewards == nil {
			x.Rewards = []*DelegationDelegatorReward{}
		}
		value := &_QueryDelegationTotalRewardsResponse_1_list{list: &x.Rewards}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.total":
		if x.Total == nil {
			x.Total = []*v1beta11.DecCoin{}
		}
		value := &_QueryDelegationTotalRewardsResponse_2_list{list: &x.Total}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.rewards":
		list := []*DelegationDelegatorReward{}
		return protoreflect.ValueOfList(&_QueryDelegationTotalRewardsResponse_1_list{list: &list})
	case "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.total":
		list := []*v1beta11.DecCoin{}
		return protoreflect.ValueOfList(&_QueryDelegationTotalRewardsResponse_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegationTotalRewardsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegationTotalRewardsResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Rewards) > 0 {
			for _, e := range x.Rewards {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.Total) > 0 {
			for _, e := range x.Total {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegationTotalRewardsResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Total) > 0 {
			for iNdEx := len(x.Total) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Total[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if len(x.Rewards) > 0 {
			for iNdEx := len(x.Rewards) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Rewards[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegationTotalRewardsResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegationTotalRewardsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegationTotalRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Rewards = append(x.Rewards, &DelegationDelegatorReward{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Rewards[len(x.Rewards)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Total = append(x.Total, &v1beta11.DecCoin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Total[len(x.Total)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryDelegatorValidatorsRequest                   protoreflect.MessageDescriptor
	fd_QueryDelegatorValidatorsRequest_delegator_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryDelegatorValidatorsRequest = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryDelegatorValidatorsRequest")
	fd_QueryDelegatorValidatorsRequest_delegator_address = md_QueryDelegatorValidatorsRequest.Fields().ByName("delegator_address")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegatorValidatorsRequest)(nil)

type fastReflection_QueryDelegatorValidatorsRequest QueryDelegatorValidatorsRequest

func (x *QueryDelegatorValidatorsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegatorValidatorsRequest)(x)
}

func (x *QueryDelegatorValidatorsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegatorValidatorsRequest_messageType fastReflection_QueryDelegatorValidatorsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegatorValidatorsRequest_messageType{}

type fastReflection_QueryDelegatorValidatorsRequest_messageType struct{}

func (x fastReflection_QueryDelegatorValidatorsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegatorValidatorsRequest)(nil)
}
func (x fastReflection_QueryDelegatorValidatorsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorValidatorsRequest)
}
func (x fastReflection_QueryDelegatorValidatorsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorValidatorsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorValidatorsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegatorValidatorsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegatorValidatorsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorValidatorsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegatorValidatorsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddress != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddress)
		if !f(fd_QueryDelegatorValidatorsRequest_delegator_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest.delegator_address":
		return x.DelegatorAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest.delegator_address":
		x.DelegatorAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest.delegator_address":
		value := x.DelegatorAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest.delegator_address":
		x.DelegatorAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest.delegator_address":
		panic(fmt.Errorf("field delegator_address of message cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegatorValidatorsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest.delegator_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegatorValidatorsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegatorValidatorsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorsRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryDelegatorValidatorsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegatorValidatorsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegatorValidatorsRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.DelegatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegatorValidatorsRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.DelegatorAddress) > 0 {
			i -= len(x.DelegatorAddress)
			copy(dAtA[i:], x.DelegatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegatorValidatorsRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorValidatorsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_QueryDelegatorValidatorsResponse_1_list)(nil)

type _QueryDelegatorValidatorsResponse_1_list struct {
	list *[]string
}

func (x *_QueryDelegatorValidatorsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryDelegatorValidatorsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_QueryDelegatorValidatorsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_QueryDelegatorValidatorsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryDelegatorValidatorsResponse_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message QueryDelegatorValidatorsResponse at list field Validators as it is not of Message kind"))
}

func (x *_QueryDelegatorValidatorsResponse_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_QueryDelegatorValidatorsResponse_1_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_QueryDelegatorValidatorsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryDelegatorValidatorsResponse            protoreflect.MessageDescriptor
	fd_QueryDelegatorValidatorsResponse_validators protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryDelegatorValidatorsResponse = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryDelegatorValidatorsResponse")
	fd_QueryDelegatorValidatorsResponse_validators = md_QueryDelegatorValidatorsResponse.Fields().ByName("validators")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegatorValidatorsResponse)(nil)

type fastReflection_QueryDelegatorValidatorsResponse QueryDelegatorValidatorsResponse

func (x *QueryDelegatorValidatorsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegatorValidatorsResponse)(x)
}

func (x *QueryDelegatorValidatorsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegatorValidatorsResponse_messageType fastReflection_QueryDelegatorValidatorsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegatorValidatorsResponse_messageType{}

type fastReflection_QueryDelegatorValidatorsResponse_messageType struct{}

func (x fastReflection_QueryDelegatorValidatorsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegatorValidatorsResponse)(nil)
}
func (x fastReflection_QueryDelegatorValidatorsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorValidatorsResponse)
}
func (x fastReflection_QueryDelegatorValidatorsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorValidatorsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorValidatorsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegatorValidatorsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegatorValidatorsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorValidatorsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegatorValidatorsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Validators) != 0 {
		value := protoreflect.ValueOfList(&_QueryDelegatorValidatorsResponse_1_list{list: &x.Validators})
		if !f(fd_QueryDelegatorValidatorsResponse_validators, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse.validators":
		return len(x.Validators) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse.validators":
		x.Validators = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse.validators":
		if len(x.Validators) == 0 {
			return protoreflect.ValueOfList(&_QueryDelegatorValidatorsResponse_1_list{})
		}
		listValue := &_QueryDelegatorValidatorsResponse_1_list{list: &x.Validators}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse.validators":
		lv := value.List()
		clv := lv.(*_QueryDelegatorValidatorsResponse_1_list)
		x.Validators = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse.validators":
		if x.Validators == nil {
			x.Validators = []string{}
		}
		value := &_QueryDelegatorValidatorsResponse_1_list{list: &x.Validators}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegatorValidatorsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse.validators":
		list := []string{}
		return protoreflect.ValueOfList(&_QueryDelegatorValidatorsResponse_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegatorValidatorsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegatorValidatorsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorsResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryDelegatorValidatorsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegatorValidatorsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegatorValidatorsResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Validators) > 0 {
			for _, s := range x.Validators {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegatorValidatorsResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Validators) > 0 {
			for iNdEx := len(x.Validators) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Validators[iNdEx])
				copy(dAtA[i:], x.Validators[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Validators[iNdEx])))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegatorValidatorsResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorValidatorsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Validators = append(x.Validators, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryDelegatorWithdrawAddressRequest                   protoreflect.MessageDescriptor
	fd_QueryDelegatorWithdrawAddressRequest_delegator_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryDelegatorWithdrawAddressRequest = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryDelegatorWithdrawAddressRequest")
	fd_QueryDelegatorWithdrawAddressRequest_delegator_address = md_QueryDelegatorWithdrawAddressRequest.Fields().ByName("delegator_address")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegatorWithdrawAddressRequest)(nil)

type fastReflection_QueryDelegatorWithdrawAddressRequest QueryDelegatorWithdrawAddressRequest

func (x *QueryDelegatorWithdrawAddressRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegatorWithdrawAddressRequest)(x)
}

func (x *QueryDelegatorWithdrawAddressRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegatorWithdrawAddressRequest_messageType fastReflection_QueryDelegatorWithdrawAddressRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegatorWithdrawAddressRequest_messageType{}

type fastReflection_QueryDelegatorWithdrawAddressRequest_messageType struct{}

func (x fastReflection_QueryDelegatorWithdrawAddressRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegatorWithdrawAddressRequest)(nil)
}
func (x fastReflection_QueryDelegatorWithdrawAddressRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorWithdrawAddressRequest)
}
func (x fastReflection_QueryDelegatorWithdrawAddressRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorWithdrawAddressRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorWithdrawAddressRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegatorWithdrawAddressRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorWithdrawAddressRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegatorWithdrawAddressRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddress != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddress)
		if !f(fd_QueryDelegatorWithdrawAddressRequest_delegator_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest.delegator_address":
		return x.DelegatorAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest.delegator_address":
		x.DelegatorAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest.delegator_address":
		value := x.DelegatorAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest.delegator_address":
		x.DelegatorAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest.delegator_address":
		panic(fmt.Errorf("field delegator_address of message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest.delegator_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegatorWithdrawAddressRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegatorWithdrawAddressRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.DelegatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegatorWithdrawAddressRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.DelegatorAddress) > 0 {
			i -= len(x.DelegatorAddress)
			copy(dAtA[i:], x.DelegatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegatorWithdrawAddressRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorWithdrawAddressRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorWithdrawAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryDelegatorWithdrawAddressResponse                  protoreflect.MessageDescriptor
	fd_QueryDelegatorWithdrawAddressResponse_withdraw_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryDelegatorWithdrawAddressResponse = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryDelegatorWithdrawAddressResponse")
	fd_QueryDelegatorWithdrawAddressResponse_withdraw_address = md_QueryDelegatorWithdrawAddressResponse.Fields().ByName("withdraw_address")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegatorWithdrawAddressResponse)(nil)

type fastReflection_QueryDelegatorWithdrawAddressResponse QueryDelegatorWithdrawAddressResponse

func (x *QueryDelegatorWithdrawAddressResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegatorWithdrawAddressResponse)(x)
}

func (x *QueryDelegatorWithdrawAddressResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegatorWithdrawAddressResponse_messageType fastReflection_QueryDelegatorWithdrawAddressResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegatorWithdrawAddressResponse_messageType{}

type fastReflection_QueryDelegatorWithdrawAddressResponse_messageType struct{}

func (x fastReflection_QueryDelegatorWithdrawAddressResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegatorWithdrawAddressResponse)(nil)
}
func (x fastReflection_QueryDelegatorWithdrawAddressResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorWithdrawAddressResponse)
}
func (x fastReflection_QueryDelegatorWithdrawAddressResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorWithdrawAddressResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorWithdrawAddressResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegatorWithdrawAddressResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorWithdrawAddressResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegatorWithdrawAddressResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.WithdrawAddress != "" {
		value := protoreflect.ValueOfString(x.WithdrawAddress)
		if !f(fd_QueryDelegatorWithdrawAddressResponse_withdraw_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse.withdraw_address":
		return x.WithdrawAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse.withdraw_address":
		x.WithdrawAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse.withdraw_address":
		value := x.WithdrawAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse.withdraw_address":
		x.WithdrawAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse.withdraw_address":
		panic(fmt.Errorf("field withdraw_address of message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse.withdraw_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegatorWithdrawAddressResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegatorWithdrawAddressResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.WithdrawAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegatorWithdrawAddressResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.WithdrawAddress) > 0 {
			i -= len(x.WithdrawAddress)
			copy(dAtA[i:], x.WithdrawAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.WithdrawAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryDelegatorWithdrawAddressResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorWithdrawAddressResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorWithdrawAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.WithdrawAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryCommunityPoolRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryCommunityPoolRequest = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryCommunityPoolRequest")
}

var _ protoreflect.Message = (*fastReflection_QueryCommunityPoolRequest)(nil)

type fastReflection_QueryCommunityPoolRequest QueryCommunityPoolRequest

func (x *QueryCommunityPoolRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryCommunityPoolRequest)(x)
}

func (x *QueryCommunityPoolRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryCommunityPoolRequest_messageType fastReflection_QueryCommunityPoolRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryCommunityPoolRequest_messageType{}

type fastReflection_QueryCommunityPoolRequest_messageType struct{}

func (x fastReflection_QueryCommunityPoolRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryCommunityPoolRequest)(nil)
}
func (x fastReflection_QueryCommunityPoolRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryCommunityPoolRequest)
}
func (x fastReflection_QueryCommunityPoolRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryCommunityPoolRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryCommunityPoolRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryCommunityPoolRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryCommunityPoolRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryCommunityPoolRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryCommunityPoolRequest) New() protoreflect.Message {
	return new(fastReflection_QueryCommunityPoolRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryCommunityPoolRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryCommunityPoolRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryCommunityPoolRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryCommunityPoolRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryCommunityPoolRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryCommunityPoolRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryCommunityPoolRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryCommunityPoolRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryCommunityPoolRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryCommunityPoolRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryCommunityPoolRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryCommunityPoolRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryCommunityPoolRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryCommunityPoolRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryCommunityPoolRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryCommunityPoolRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryCommunityPoolRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryCommunityPoolRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryCommunityPoolRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryCommunityPoolRequest"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryCommunityPoolRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryCommunityPoolRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryCommunityPoolRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryCommunityPoolRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryCommunityPoolRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryCommunityPoolRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryCommunityPoolRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryCommunityPoolRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryCommunityPoolRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryCommunityPoolRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryCommunityPoolRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryCommunityPoolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_QueryCommunityPoolResponse_1_list)(nil)

type _QueryCommunityPoolResponse_1_list struct {
	list *[]*v1beta11.DecCoin
}

func (x *_QueryCommunityPoolResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryCommunityPoolResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryCommunityPoolResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta11.DecCoin)
	(*x.list)[i] = concreteValue
}

func (x *_QueryCommunityPoolResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta11.DecCoin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryCommunityPoolResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta11.DecCoin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryCommunityPoolResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryCommunityPoolResponse_1_list) NewElement() protoreflect.Value {
	v := new(v1beta11.DecCoin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryCommunityPoolResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryCommunityPoolResponse      protoreflect.MessageDescriptor
	fd_QueryCommunityPoolResponse_pool protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_query_proto_init()
	md_QueryCommunityPoolResponse = File_cosmos_distribution_v1beta1_query_proto.Messages().ByName("QueryCommunityPoolResponse")
	fd_QueryCommunityPoolResponse_pool = md_QueryCommunityPoolResponse.Fields().ByName("pool")
}

var _ protoreflect.Message = (*fastReflection_QueryCommunityPoolResponse)(nil)

type fastReflection_QueryCommunityPoolResponse QueryCommunityPoolResponse

func (x *QueryCommunityPoolResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryCommunityPoolResponse)(x)
}

func (x *QueryCommunityPoolResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryCommunityPoolResponse_messageType fastReflection_QueryCommunityPoolResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryCommunityPoolResponse_messageType{}

type fastReflection_QueryCommunityPoolResponse_messageType struct{}

func (x fastReflection_QueryCommunityPoolResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryCommunityPoolResponse)(nil)
}
func (x fastReflection_QueryCommunityPoolResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryCommunityPoolResponse)
}
func (x fastReflection_QueryCommunityPoolResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryCommunityPoolResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryCommunityPoolResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryCommunityPoolResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryCommunityPoolResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryCommunityPoolResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryCommunityPoolResponse) New() protoreflect.Message {
	return new(fastReflection_QueryCommunityPoolResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryCommunityPoolResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryCommunityPoolResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryCommunityPoolResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Pool) != 0 {
		value := protoreflect.ValueOfList(&_QueryCommunityPoolResponse_1_list{list: &x.Pool})
		if !f(fd_QueryCommunityPoolResponse_pool, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryCommunityPoolResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryCommunityPoolResponse.pool":
		return len(x.Pool) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryCommunityPoolResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryCommunityPoolResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryCommunityPoolResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryCommunityPoolResponse.pool":
		x.Pool = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryCommunityPoolResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryCommunityPoolResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryCommunityPoolResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.QueryCommunityPoolResponse.pool":
		if len(x.Pool) == 0 {
			return protoreflect.ValueOfList(&_QueryCommunityPoolResponse_1_list{})
		}
		listValue := &_QueryCommunityPoolResponse_1_list{list: &x.Pool}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryCommunityPoolResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryCommunityPoolResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryCommunityPoolResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryCommunityPoolResponse.pool":
		lv := value.List()
		clv := lv.(*_QueryCommunityPoolResponse_1_list)
		x.Pool = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryCommunityPoolResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryCommunityPoolResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryCommunityPoolResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryCommunityPoolResponse.pool":
		if x.Pool == nil {
			x.Pool = []*v1beta11.DecCoin{}
		}
		value := &_QueryCommunityPoolResponse_1_list{list: &x.Pool}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryCommunityPoolResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryCommunityPoolResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryCommunityPoolResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.QueryCommunityPoolResponse.pool":
		list := []*v1beta11.DecCoin{}
		return protoreflect.ValueOfList(&_QueryCommunityPoolResponse_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.QueryCommunityPoolResponse"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.QueryCommunityPoolResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryCommunityPoolResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.QueryCommunityPoolResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryCommunityPoolResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryCommunityPoolResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryCommunityPoolResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryCommunityPoolResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryCommunityPoolResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Pool) > 0 {
			for _, e := range x.Pool {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryCommunityPoolResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Pool) > 0 {
			for iNdEx := len(x.Pool) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Pool[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryCommunityPoolResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryCommunityPoolResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryCommunityPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Pool = append(x.Pool, &v1beta11.DecCoin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pool[len(x.Pool)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries params of the distribution module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// ValidatorOutstandingRewards queries rewards of a validator address.
	ValidatorOutstandingRewards(ctx context.Context, in *QueryValidatorOutstandingRewardsRequest, opts ...grpc.CallOption) (*QueryValidatorOutstandingRewardsResponse, error)
	// ValidatorCommission queries accumulated commission for a validator.
	ValidatorCommission(ctx context.Context, in *QueryValidatorCommissionRequest, opts ...grpc.CallOption) (*QueryValidatorCommissionResponse, error)
	// ValidatorSlashes queries slash events of a validator.
	ValidatorSlashes(ctx context.Context, in *QueryValidatorSlashesRequest, opts ...grpc.CallOption) (*QueryValidatorSlashesResponse, error)
	// DelegationRewards queries the total rewards accrued by a delegation.
	DelegationRewards(ctx context.Context, in *QueryDelegationRewardsRequest, opts ...grpc.CallOption) (*QueryDelegationRewardsResponse, error)
	// DelegationTotalRewards queries the total rewards accrued by a each
	// validator.
	DelegationTotalRewards(ctx context.Context, in *QueryDelegationTotalRewardsRequest, opts ...grpc.CallOption) (*QueryDelegationTotalRewardsResponse, error)
	// DelegatorValidators queries the validators of a delegator.
	DelegatorValidators(ctx context.Context, in *QueryDelegatorValidatorsRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorsResponse, error)
	// DelegatorWithdrawAddress queries withdraw address of a delegator.
	DelegatorWithdrawAddress(ctx context.Context, in *QueryDelegatorWithdrawAddressRequest, opts ...grpc.CallOption) (*QueryDelegatorWithdrawAddressResponse, error)
	// CommunityPool queries the community pool coins.
	CommunityPool(ctx context.Context, in *QueryCommunityPoolRequest, opts ...grpc.CallOption) (*QueryCommunityPoolResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorOutstandingRewards(ctx context.Context, in *QueryValidatorOutstandingRewardsRequest, opts ...grpc.CallOption) (*QueryValidatorOutstandingRewardsResponse, error) {
	out := new(QueryValidatorOutstandingRewardsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/ValidatorOutstandingRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorCommission(ctx context.Context, in *QueryValidatorCommissionRequest, opts ...grpc.CallOption) (*QueryValidatorCommissionResponse, error) {
	out := new(QueryValidatorCommissionResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/ValidatorCommission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorSlashes(ctx context.Context, in *QueryValidatorSlashesRequest, opts ...grpc.CallOption) (*QueryValidatorSlashesResponse, error) {
	out := new(QueryValidatorSlashesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/ValidatorSlashes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegationRewards(ctx context.Context, in *QueryDelegationRewardsRequest, opts ...grpc.CallOption) (*QueryDelegationRewardsResponse, error) {
	out := new(QueryDelegationRewardsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/DelegationRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegationTotalRewards(ctx context.Context, in *QueryDelegationTotalRewardsRequest, opts ...grpc.CallOption) (*QueryDelegationTotalRewardsResponse, error) {
	out := new(QueryDelegationTotalRewardsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/DelegationTotalRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorValidators(ctx context.Context, in *QueryDelegatorValidatorsRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorsResponse, error) {
	out := new(QueryDelegatorValidatorsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/DelegatorValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorWithdrawAddress(ctx context.Context, in *QueryDelegatorWithdrawAddressRequest, opts ...grpc.CallOption) (*QueryDelegatorWithdrawAddressResponse, error) {
	out := new(QueryDelegatorWithdrawAddressResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/DelegatorWithdrawAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CommunityPool(ctx context.Context, in *QueryCommunityPoolRequest, opts ...grpc.CallOption) (*QueryCommunityPoolResponse, error) {
	out := new(QueryCommunityPoolResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/CommunityPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Params queries params of the distribution module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// ValidatorOutstandingRewards queries rewards of a validator address.
	ValidatorOutstandingRewards(context.Context, *QueryValidatorOutstandingRewardsRequest) (*QueryValidatorOutstandingRewardsResponse, error)
	// ValidatorCommission queries accumulated commission for a validator.
	ValidatorCommission(context.Context, *QueryValidatorCommissionRequest) (*QueryValidatorCommissionResponse, error)
	// ValidatorSlashes queries slash events of a validator.
	ValidatorSlashes(context.Context, *QueryValidatorSlashesRequest) (*QueryValidatorSlashesResponse, error)
	// DelegationRewards queries the total rewards accrued by a delegation.
	DelegationRewards(context.Context, *QueryDelegationRewardsRequest) (*QueryDelegationRewardsResponse, error)
	// DelegationTotalRewards queries the total rewards accrued by a each
	// validator.
	DelegationTotalRewards(context.Context, *QueryDelegationTotalRewardsRequest) (*QueryDelegationTotalRewardsResponse, error)
	// DelegatorValidators queries the validators of a delegator.
	DelegatorValidators(context.Context, *QueryDelegatorValidatorsRequest) (*QueryDelegatorValidatorsResponse, error)
	// DelegatorWithdrawAddress queries withdraw address of a delegator.
	DelegatorWithdrawAddress(context.Context, *QueryDelegatorWithdrawAddressRequest) (*QueryDelegatorWithdrawAddressResponse, error)
	// CommunityPool queries the community pool coins.
	CommunityPool(context.Context, *QueryCommunityPoolRequest) (*QueryCommunityPoolResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) ValidatorOutstandingRewards(context.Context, *QueryValidatorOutstandingRewardsRequest) (*QueryValidatorOutstandingRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorOutstandingRewards not implemented")
}
func (UnimplementedQueryServer) ValidatorCommission(context.Context, *QueryValidatorCommissionRequest) (*QueryValidatorCommissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorCommission not implemented")
}
func (UnimplementedQueryServer) ValidatorSlashes(context.Context, *QueryValidatorSlashesRequest) (*QueryValidatorSlashesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorSlashes not implemented")
}
func (UnimplementedQueryServer) DelegationRewards(context.Context, *QueryDelegationRewardsRequest) (*QueryDelegationRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegationRewards not implemented")
}
func (UnimplementedQueryServer) DelegationTotalRewards(context.Context, *QueryDelegationTotalRewardsRequest) (*QueryDelegationTotalRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegationTotalRewards not implemented")
}
func (UnimplementedQueryServer) DelegatorValidators(context.Context, *QueryDelegatorValidatorsRequest) (*QueryDelegatorValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatorValidators not implemented")
}
func (UnimplementedQueryServer) DelegatorWithdrawAddress(context.Context, *QueryDelegatorWithdrawAddressRequest) (*QueryDelegatorWithdrawAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatorWithdrawAddress not implemented")
}
func (UnimplementedQueryServer) CommunityPool(context.Context, *QueryCommunityPoolRequest) (*QueryCommunityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommunityPool not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorOutstandingRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorOutstandingRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorOutstandingRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/ValidatorOutstandingRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorOutstandingRewards(ctx, req.(*QueryValidatorOutstandingRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorCommission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorCommissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorCommission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/ValidatorCommission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorCommission(ctx, req.(*QueryValidatorCommissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorSlashes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorSlashesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorSlashes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/ValidatorSlashes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorSlashes(ctx, req.(*QueryValidatorSlashesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegationRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegationRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegationRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegationRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegationRewards(ctx, req.(*QueryDelegationRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegationTotalRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegationTotalRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegationTotalRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegationTotalRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegationTotalRewards(ctx, req.(*QueryDelegationTotalRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegatorValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorValidators(ctx, req.(*QueryDelegatorValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorWithdrawAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorWithdrawAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorWithdrawAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegatorWithdrawAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorWithdrawAddress(ctx, req.(*QueryDelegatorWithdrawAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommunityPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/CommunityPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CommunityPool(ctx, req.(*QueryCommunityPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.distribution.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ValidatorOutstandingRewards",
			Handler:    _Query_ValidatorOutstandingRewards_Handler,
		},
		{
			MethodName: "ValidatorCommission",
			Handler:    _Query_ValidatorCommission_Handler,
		},
		{
			MethodName: "ValidatorSlashes",
			Handler:    _Query_ValidatorSlashes_Handler,
		},
		{
			MethodName: "DelegationRewards",
			Handler:    _Query_DelegationRewards_Handler,
		},
		{
			MethodName: "DelegationTotalRewards",
			Handler:    _Query_DelegationTotalRewards_Handler,
		},
		{
			MethodName: "DelegatorValidators",
			Handler:    _Query_DelegatorValidators_Handler,
		},
		{
			MethodName: "DelegatorWithdrawAddress",
			Handler:    _Query_DelegatorWithdrawAddress_Handler,
		},
		{
			MethodName: "CommunityPool",
			Handler:    _Query_CommunityPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/distribution/v1beta1/query.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/distribution/v1beta1/query.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryParamsRequest) Reset() {
	*x = QueryParamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsRequest) ProtoMessage() {}

// Deprecated: Use QueryParamsRequest.ProtoReflect.Descriptor instead.
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{0}
}

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params defines the parameters of the module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *QueryParamsResponse) Reset() {
	*x = QueryParamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsResponse) ProtoMessage() {}

// Deprecated: Use QueryParamsResponse.ProtoReflect.Descriptor instead.
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryParamsResponse) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

// QueryValidatorOutstandingRewardsRequest is the request type for the
// Query/ValidatorOutstandingRewards RPC method.
type QueryValidatorOutstandingRewardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator_address defines the validator address to query for.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (x *QueryValidatorOutstandingRewardsRequest) Reset() {
	*x = QueryValidatorOutstandingRewardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorOutstandingRewardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorOutstandingRewardsRequest) ProtoMessage() {}

// Deprecated: Use QueryValidatorOutstandingRewardsRequest.ProtoReflect.Descriptor instead.
func (*QueryValidatorOutstandingRewardsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryValidatorOutstandingRewardsRequest) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

// QueryValidatorOutstandingRewardsResponse is the response type for the
// Query/ValidatorOutstandingRewards RPC method.
type QueryValidatorOutstandingRewardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rewards *ValidatorOutstandingRewards `protobuf:"bytes,1,opt,name=rewards,proto3" json:"rewards,omitempty"`
}

func (x *QueryValidatorOutstandingRewardsResponse) Reset() {
	*x = QueryValidatorOutstandingRewardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorOutstandingRewardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorOutstandingRewardsResponse) ProtoMessage() {}

// Deprecated: Use QueryValidatorOutstandingRewardsResponse.ProtoReflect.Descriptor instead.
func (*QueryValidatorOutstandingRewardsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{3}
}

func (x *QueryValidatorOutstandingRewardsResponse) GetRewards() *ValidatorOutstandingRewards {
	if x != nil {
		return x.Rewards
	}
	return nil
}

// QueryValidatorCommissionRequest is the request type for the
// Query/ValidatorCommission RPC method
type QueryValidatorCommissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator_address defines the validator address to query for.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (x *QueryValidatorCommissionRequest) Reset() {
	*x = QueryValidatorCommissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorCommissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorCommissionRequest) ProtoMessage() {}

// Deprecated: Use QueryValidatorCommissionRequest.ProtoReflect.Descriptor instead.
func (*QueryValidatorCommissionRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{4}
}

func (x *QueryValidatorCommissionRequest) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

// QueryValidatorCommissionResponse is the response type for the
// Query/ValidatorCommission RPC method
type QueryValidatorCommissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// commission defines the commision the validator received.
	Commission *ValidatorAccumulatedCommission `protobuf:"bytes,1,opt,name=commission,proto3" json:"commission,omitempty"`
}

func (x *QueryValidatorCommissionResponse) Reset() {
	*x = QueryValidatorCommissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorCommissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorCommissionResponse) ProtoMessage() {}

// Deprecated: Use QueryValidatorCommissionResponse.ProtoReflect.Descriptor instead.
func (*QueryValidatorCommissionResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{5}
}

func (x *QueryValidatorCommissionResponse) GetCommission() *ValidatorAccumulatedCommission {
	if x != nil {
		return x.Commission
	}
	return nil
}

// QueryValidatorSlashesRequest is the request type for the
// Query/ValidatorSlashes RPC method
type QueryValidatorSlashesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator_address defines the validator address to query for.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// starting_height defines the optional starting height to query the slashes.
	StartingHeight uint64 `protobuf:"varint,2,opt,name=starting_height,json=startingHeight,proto3" json:"starting_height,omitempty"`
	// starting_height defines the optional ending height to query the slashes.
	EndingHeight uint64 `protobuf:"varint,3,opt,name=ending_height,json=endingHeight,proto3" json:"ending_height,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryValidatorSlashesRequest) Reset() {
	*x = QueryValidatorSlashesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorSlashesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorSlashesRequest) ProtoMessage() {}

// Deprecated: Use QueryValidatorSlashesRequest.ProtoReflect.Descriptor instead.
func (*QueryValidatorSlashesRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{6}
}

func (x *QueryValidatorSlashesRequest) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

func (x *QueryValidatorSlashesRequest) GetStartingHeight() uint64 {
	if x != nil {
		return x.StartingHeight
	}
	return 0
}

func (x *QueryValidatorSlashesRequest) GetEndingHeight() uint64 {
	if x != nil {
		return x.EndingHeight
	}
	return 0
}

func (x *QueryValidatorSlashesRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryValidatorSlashesResponse is the response type for the
// Query/ValidatorSlashes RPC method.
type QueryValidatorSlashesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// slashes defines the slashes the validator received.
	Slashes []*ValidatorSlashEvent `protobuf:"bytes,1,rep,name=slashes,proto3" json:"slashes,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryValidatorSlashesResponse) Reset() {
	*x = QueryValidatorSlashesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorSlashesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorSlashesResponse) ProtoMessage() {}

// Deprecated: Use QueryValidatorSlashesResponse.ProtoReflect.Descriptor instead.
func (*QueryValidatorSlashesResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{7}
}

func (x *QueryValidatorSlashesResponse) GetSlashes() []*ValidatorSlashEvent {
	if x != nil {
		return x.Slashes
	}
	return nil
}

func (x *QueryValidatorSlashesResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryDelegationRewardsRequest is the request type for the
// Query/DelegationRewards RPC method.
type QueryDelegationRewardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_address defines the delegator address to query for.
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	// validator_address defines the validator address to query for.
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (x *QueryDelegationRewardsRequest) Reset() {
	*x = QueryDelegationRewardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegationRewardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegationRewardsRequest) ProtoMessage() {}

// Deprecated: Use QueryDelegationRewardsRequest.ProtoReflect.Descriptor instead.
func (*QueryDelegationRewardsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{8}
}

func (x *QueryDelegationRewardsRequest) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

func (x *QueryDelegationRewardsRequest) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

// QueryDelegationRewardsResponse is the response type for the
// Query/DelegationRewards RPC method.
type QueryDelegationRewardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// rewards defines the rewards accrued by a delegation.
	Rewards []*v1beta11.DecCoin `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards,omitempty"`
}

func (x *QueryDelegationRewardsResponse) Reset() {
	*x = QueryDelegationRewardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegationRewardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegationRewardsResponse) ProtoMessage() {}

// Deprecated: Use QueryDelegationRewardsResponse.ProtoReflect.Descriptor instead.
func (*QueryDelegationRewardsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{9}
}

func (x *QueryDelegationRewardsResponse) GetRewards() []*v1beta11.DecCoin {
	if x != nil {
		return x.Rewards
	}
	return nil
}

// QueryDelegationTotalRewardsRequest is the request type for the
// Query/DelegationTotalRewards RPC method.
type QueryDelegationTotalRewardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_address defines the delegator address to query for.
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
}

func (x *QueryDelegationTotalRewardsRequest) Reset() {
	*x = QueryDelegationTotalRewardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegationTotalRewardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegationTotalRewardsRequest) ProtoMessage() {}

// Deprecated: Use QueryDelegationTotalRewardsRequest.ProtoReflect.Descriptor instead.
func (*QueryDelegationTotalRewardsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{10}
}

func (x *QueryDelegationTotalRewardsRequest) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

// QueryDelegationTotalRewardsResponse is the response type for the
// Query/DelegationTotalRewards RPC method.
type QueryDelegationTotalRewardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// rewards defines all the rewards accrued by a delegator.
	Rewards []*DelegationDelegatorReward `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards,omitempty"`
	// total defines the sum of all the rewards.
	Total []*v1beta11.DecCoin `protobuf:"bytes,2,rep,name=total,proto3" json:"total,omitempty"`
}

func (x *QueryDelegationTotalRewardsResponse) Reset() {
	*x = QueryDelegationTotalRewardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegationTotalRewardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegationTotalRewardsResponse) ProtoMessage() {}

// Deprecated: Use QueryDelegationTotalRewardsResponse.ProtoReflect.Descriptor instead.
func (*QueryDelegationTotalRewardsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{11}
}

func (x *QueryDelegationTotalRewardsResponse) GetRewards() []*DelegationDelegatorReward {
	if x != nil {
		return x.Rewards
	}
	return nil
}

func (x *QueryDelegationTotalRewardsResponse) GetTotal() []*v1beta11.DecCoin {
	if x != nil {
		return x.Total
	}
	return nil
}

// QueryDelegatorValidatorsRequest is the request type for the
// Query/DelegatorValidators RPC method.
type QueryDelegatorValidatorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_address defines the delegator address to query for.
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
}

func (x *QueryDelegatorValidatorsRequest) Reset() {
	*x = QueryDelegatorValidatorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatorValidatorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatorValidatorsRequest) ProtoMessage() {}

// Deprecated: Use QueryDelegatorValidatorsRequest.ProtoReflect.Descriptor instead.
func (*QueryDelegatorValidatorsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{12}
}

func (x *QueryDelegatorValidatorsRequest) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

// QueryDelegatorValidatorsResponse is the response type for the
// Query/DelegatorValidators RPC method.
type QueryDelegatorValidatorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validators defines the validators a delegator is delegating for.
	Validators []string `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (x *QueryDelegatorValidatorsResponse) Reset() {
	*x = QueryDelegatorValidatorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatorValidatorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatorValidatorsResponse) ProtoMessage() {}

// Deprecated: Use QueryDelegatorValidatorsResponse.ProtoReflect.Descriptor instead.
func (*QueryDelegatorValidatorsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{13}
}

func (x *QueryDelegatorValidatorsResponse) GetValidators() []string {
	if x != nil {
		return x.Validators
	}
	return nil
}

// QueryDelegatorWithdrawAddressRequest is the request type for the
// Query/DelegatorWithdrawAddress RPC method.
type QueryDelegatorWithdrawAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_address defines the delegator address to query for.
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
}

func (x *QueryDelegatorWithdrawAddressRequest) Reset() {
	*x = QueryDelegatorWithdrawAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatorWithdrawAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatorWithdrawAddressRequest) ProtoMessage() {}

// Deprecated: Use QueryDelegatorWithdrawAddressRequest.ProtoReflect.Descriptor instead.
func (*QueryDelegatorWithdrawAddressRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{14}
}

func (x *QueryDelegatorWithdrawAddressRequest) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

// QueryDelegatorWithdrawAddressResponse is the response type for the
// Query/DelegatorWithdrawAddress RPC method.
type QueryDelegatorWithdrawAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// withdraw_address defines the delegator address to query for.
	WithdrawAddress string `protobuf:"bytes,1,opt,name=withdraw_address,json=withdrawAddress,proto3" json:"withdraw_address,omitempty"`
}

func (x *QueryDelegatorWithdrawAddressResponse) Reset() {
	*x = QueryDelegatorWithdrawAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatorWithdrawAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatorWithdrawAddressResponse) ProtoMessage() {}

// Deprecated: Use QueryDelegatorWithdrawAddressResponse.ProtoReflect.Descriptor instead.
func (*QueryDelegatorWithdrawAddressResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{15}
}

func (x *QueryDelegatorWithdrawAddressResponse) GetWithdrawAddress() string {
	if x != nil {
		return x.WithdrawAddress
	}
	return ""
}

// QueryCommunityPoolRequest is the request type for the Query/CommunityPool RPC
// method.
type QueryCommunityPoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryCommunityPoolRequest) Reset() {
	*x = QueryCommunityPoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCommunityPoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCommunityPoolRequest) ProtoMessage() {}

// Deprecated: Use QueryCommunityPoolRequest.ProtoReflect.Descriptor instead.
func (*QueryCommunityPoolRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{16}
}

// QueryCommunityPoolResponse is the response type for the Query/CommunityPool
// RPC method.
type QueryCommunityPoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pool defines community pool's coins.
	Pool []*v1beta11.DecCoin `protobuf:"bytes,1,rep,name=pool,proto3" json:"pool,omitempty"`
}

func (x *QueryCommunityPoolResponse) Reset() {
	*x = QueryCommunityPoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_query_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCommunityPoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCommunityPoolResponse) ProtoMessage() {}

// Deprecated: Use QueryCommunityPoolResponse.ProtoReflect.Descriptor instead.
func (*QueryCommunityPoolResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP(), []int{17}
}

func (x *QueryCommunityPoolResponse) GetPool() []*v1beta11.DecCoin {
	if x != nil {
		return x.Pool
	}
	return nil
}

var File_cosmos_distribution_v1beta1_query_proto protoreflect.FileDescriptor

var file_cosmos_distribution_v1beta1_query_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x14, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x22, 0x70, 0x0a, 0x27, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x11,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x28, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x58, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f,
	0x00, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x22, 0x68, 0x0a, 0x1f, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a,
	0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00,
	0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x85, 0x02, 0x0a,
	0x1c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a,
	0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00,
	0x98, 0xa0, 0x1f, 0x01, 0x22, 0xba, 0x01, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x6c, 0x61, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52,
	0x07, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xb7, 0x01, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18,
	0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x45, 0x0a, 0x11, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x8d, 0x01, 0x0a, 0x1e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b,
	0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x33, 0xc8,
	0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f, 0x69,
	0x6e, 0x73, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x22, 0x75, 0x0a, 0x22, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x45, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4,
	0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0,
	0x1f, 0x00, 0x22, 0xe6, 0x01, 0x0a, 0x23, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x07, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x12, 0x67, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x42,
	0x33, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x43,
	0x6f, 0x69, 0x6e, 0x73, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x72, 0x0a, 0x1f, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45,
	0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22,
	0x4c, 0x0a, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x77, 0x0a,
	0x24, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x08, 0x88, 0xa0,
	0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x76, 0x0a, 0x25, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x43, 0x0a, 0x10, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x1b,
	0x0a, 0x19, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x1a,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x50, 0x6f,
	0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x04, 0x70, 0x6f,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44,
	0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x33, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x04, 0x70, 0x6f, 0x6f,
	0x6c, 0x32, 0xd8, 0x0f, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x98, 0x01, 0x0a, 0x06,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x25, 0x12, 0x23, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x83, 0x02, 0x0a, 0x1b, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x44, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x57, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x51, 0x12, 0x4f, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0xe2, 0x01, 0x0a,
	0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x48, 0x12, 0x46, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0xd6, 0x01, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x39, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x6c,
	0x61, 0x73, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x45, 0x12, 0x43, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x7d, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0xed, 0x01, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x59, 0x12, 0x57, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x12, 0xe8, 0x01, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x3f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x45,
	0x12, 0x43, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0xe2, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x3c, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x48, 0x12, 0x46, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0xf7, 0x01, 0x0a, 0x18, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x41, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4e, 0x12, 0x4c, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x7b, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x7d, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0xb5, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12,
	0x2b, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x42, 0x8d, 0x02, 0x0a,
	0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x43, 0x44, 0x58, 0xaa, 0x02, 0x1b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xe2, 0x02, 0x27, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_distribution_v1beta1_query_proto_rawDescOnce sync.Once
	file_cosmos_distribution_v1beta1_query_proto_rawDescData = file_cosmos_distribution_v1beta1_query_proto_rawDesc
)

func file_cosmos_distribution_v1beta1_query_proto_rawDescGZIP() []byte {
	file_cosmos_distribution_v1beta1_query_proto_rawDescOnce.Do(func() {
		file_cosmos_distribution_v1beta1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_distribution_v1beta1_query_proto_rawDescData)
	})
	return file_cosmos_distribution_v1beta1_query_proto_rawDescData
}

var file_cosmos_distribution_v1beta1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_cosmos_distribution_v1beta1_query_proto_goTypes = []interface{}{
	(*QueryParamsRequest)(nil),                       // 0: cosmos.distribution.v1beta1.QueryParamsRequest
	(*QueryParamsResponse)(nil),                      // 1: cosmos.distribution.v1beta1.QueryParamsResponse
	(*QueryValidatorOutstandingRewardsRequest)(nil),  // 2: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest
	(*QueryValidatorOutstandingRewardsResponse)(nil), // 3: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse
	(*QueryValidatorCommissionRequest)(nil),          // 4: cosmos.distribution.v1beta1.QueryValidatorCommissionRequest
	(*QueryValidatorCommissionResponse)(nil),         // 5: cosmos.distribution.v1beta1.QueryValidatorCommissionResponse
	(*QueryValidatorSlashesRequest)(nil),             // 6: cosmos.distribution.v1beta1.QueryValidatorSlashesRequest
	(*QueryValidatorSlashesResponse)(nil),            // 7: cosmos.distribution.v1beta1.QueryValidatorSlashesResponse
	(*QueryDelegationRewardsRequest)(nil),            // 8: cosmos.distribution.v1beta1.QueryDelegationRewardsRequest
	(*QueryDelegationRewardsResponse)(nil),           // 9: cosmos.distribution.v1beta1.QueryDelegationRewardsResponse
	(*QueryDelegationTotalRewardsRequest)(nil),       // 10: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest
	(*QueryDelegationTotalRewardsResponse)(nil),      // 11: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse
	(*QueryDelegatorValidatorsRequest)(nil),          // 12: cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest
	(*QueryDelegatorValidatorsResponse)(nil),         // 13: cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse
	(*QueryDelegatorWithdrawAddressRequest)(nil),     // 14: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest
	(*QueryDelegatorWithdrawAddressResponse)(nil),    // 15: cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse
	(*QueryCommunityPoolRequest)(nil),                // 16: cosmos.distribution.v1beta1.QueryCommunityPoolRequest
	(*QueryCommunityPoolResponse)(nil),               // 17: cosmos.distribution.v1beta1.QueryCommunityPoolResponse
	(*Params)(nil),                                   // 18: cosmos.distribution.v1beta1.Params
	(*ValidatorOutstandingRewards)(nil),              // 19: cosmos.distribution.v1beta1.ValidatorOutstandingRewards
	(*ValidatorAccumulatedCommission)(nil),           // 20: cosmos.distribution.v1beta1.ValidatorAccumulatedCommission
	(*v1beta1.PageRequest)(nil),                      // 21: cosmos.base.query.v1beta1.PageRequest
	(*ValidatorSlashEvent)(nil),                      // 22: cosmos.distribution.v1beta1.ValidatorSlashEvent
	(*v1beta1.PageResponse)(nil),                     // 23: cosmos.base.query.v1beta1.PageResponse
	(*v1beta11.DecCoin)(nil),                         // 24: cosmos.base.v1beta1.DecCoin
	(*DelegationDelegatorReward)(nil),                // 25: cosmos.distribution.v1beta1.DelegationDelegatorReward
}
var file_cosmos_distribution_v1beta1_query_proto_depIdxs = []int32{
	18, // 0: cosmos.distribution.v1beta1.QueryParamsResponse.params:type_name -> cosmos.distribution.v1beta1.Params
	19, // 1: cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse.rewards:type_name -> cosmos.distribution.v1beta1.ValidatorOutstandingRewards
	20, // 2: cosmos.distribution.v1beta1.QueryValidatorCommissionResponse.commission:type_name -> cosmos.distribution.v1beta1.ValidatorAccumulatedCommission
	21, // 3: cosmos.distribution.v1beta1.QueryValidatorSlashesRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	22, // 4: cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.slashes:type_name -> cosmos.distribution.v1beta1.ValidatorSlashEvent
	23, // 5: cosmos.distribution.v1beta1.QueryValidatorSlashesResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	24, // 6: cosmos.distribution.v1beta1.QueryDelegationRewardsResponse.rewards:type_name -> cosmos.base.v1beta1.DecCoin
	25, // 7: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.rewards:type_name -> cosmos.distribution.v1beta1.DelegationDelegatorReward
	24, // 8: cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse.total:type_name -> cosmos.base.v1beta1.DecCoin
	24, // 9: cosmos.distribution.v1beta1.QueryCommunityPoolResponse.pool:type_name -> cosmos.base.v1beta1.DecCoin
	0,  // 10: cosmos.distribution.v1beta1.Query.Params:input_type -> cosmos.distribution.v1beta1.QueryParamsRequest
	2,  // 11: cosmos.distribution.v1beta1.Query.ValidatorOutstandingRewards:input_type -> cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest
	4,  // 12: cosmos.distribution.v1beta1.Query.ValidatorCommission:input_type -> cosmos.distribution.v1beta1.QueryValidatorCommissionRequest
	6,  // 13: cosmos.distribution.v1beta1.Query.ValidatorSlashes:input_type -> cosmos.distribution.v1beta1.QueryValidatorSlashesRequest
	8,  // 14: cosmos.distribution.v1beta1.Query.DelegationRewards:input_type -> cosmos.distribution.v1beta1.QueryDelegationRewardsRequest
	10, // 15: cosmos.distribution.v1beta1.Query.DelegationTotalRewards:input_type -> cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest
	12, // 16: cosmos.distribution.v1beta1.Query.DelegatorValidators:input_type -> cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest
	14, // 17: cosmos.distribution.v1beta1.Query.DelegatorWithdrawAddress:input_type -> cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest
	16, // 18: cosmos.distribution.v1beta1.Query.CommunityPool:input_type -> cosmos.distribution.v1beta1.QueryCommunityPoolRequest
	1,  // 19: cosmos.distribution.v1beta1.Query.Params:output_type -> cosmos.distribution.v1beta1.QueryParamsResponse
	3,  // 20: cosmos.distribution.v1beta1.Query.ValidatorOutstandingRewards:output_type -> cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse
	5,  // 21: cosmos.distribution.v1beta1.Query.ValidatorCommission:output_type -> cosmos.distribution.v1beta1.QueryValidatorCommissionResponse
	7,  // 22: cosmos.distribution.v1beta1.Query.ValidatorSlashes:output_type -> cosmos.distribution.v1beta1.QueryValidatorSlashesResponse
	9,  // 23: cosmos.distribution.v1beta1.Query.DelegationRewards:output_type -> cosmos.distribution.v1beta1.QueryDelegationRewardsResponse
	11, // 24: cosmos.distribution.v1beta1.Query.DelegationTotalRewards:output_type -> cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse
	13, // 25: cosmos.distribution.v1beta1.Query.DelegatorValidators:output_type -> cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse
	15, // 26: cosmos.distribution.v1beta1.Query.DelegatorWithdrawAddress:output_type -> cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse
	17, // 27: cosmos.distribution.v1beta1.Query.CommunityPool:output_type -> cosmos.distribution.v1beta1.QueryCommunityPoolResponse
	19, // [19:28] is the sub-list for method output_type
	10, // [10:19] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_cosmos_distribution_v1beta1_query_proto_init() }
func file_cosmos_distribution_v1beta1_query_proto_init() {
	if File_cosmos_distribution_v1beta1_query_proto != nil {
		return
	}
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorOutstandingRewardsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorOutstandingRewardsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorCommissionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorCommissionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorSlashesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorSlashesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegationRewardsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegationRewardsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegationTotalRewardsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegationTotalRewardsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatorValidatorsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatorValidatorsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatorWithdrawAddressRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatorWithdrawAddressResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCommunityPoolRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_query_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCommunityPoolResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_distribution_v1beta1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_distribution_v1beta1_query_proto_goTypes,
		DependencyIndexes: file_cosmos_distribution_v1beta1_query_proto_depIdxs,
		MessageInfos:      file_cosmos_distribution_v1beta1_query_proto_msgTypes,
	}.Build()
	File_cosmos_distribution_v1beta1_query_proto = out.File
	file_cosmos_distribution_v1beta1_query_proto_rawDesc = nil
	file_cosmos_distribution_v1beta1_query_proto_goTypes = nil
	file_cosmos_distribution_v1beta1_query_proto_depIdxs = nil
}
