package crisisv1beta1

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
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
	md_MsgVerifyInvariant                       protoreflect.MessageDescriptor
	fd_MsgVerifyInvariant_sender                protoreflect.FieldDescriptor
	fd_MsgVerifyInvariant_invariant_module_name protoreflect.FieldDescriptor
	fd_MsgVerifyInvariant_invariant_route       protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_crisis_v1beta1_tx_proto_init()
	md_MsgVerifyInvariant = File_cosmos_crisis_v1beta1_tx_proto.Messages().ByName("MsgVerifyInvariant")
	fd_MsgVerifyInvariant_sender = md_MsgVerifyInvariant.Fields().ByName("sender")
	fd_MsgVerifyInvariant_invariant_module_name = md_MsgVerifyInvariant.Fields().ByName("invariant_module_name")
	fd_MsgVerifyInvariant_invariant_route = md_MsgVerifyInvariant.Fields().ByName("invariant_route")
}

var _ protoreflect.Message = (*fastReflection_MsgVerifyInvariant)(nil)

type fastReflection_MsgVerifyInvariant MsgVerifyInvariant

func (x *MsgVerifyInvariant) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgVerifyInvariant)(x)
}

func (x *MsgVerifyInvariant) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crisis_v1beta1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgVerifyInvariant_messageType fastReflection_MsgVerifyInvariant_messageType
var _ protoreflect.MessageType = fastReflection_MsgVerifyInvariant_messageType{}

type fastReflection_MsgVerifyInvariant_messageType struct{}

func (x fastReflection_MsgVerifyInvariant_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgVerifyInvariant)(nil)
}
func (x fastReflection_MsgVerifyInvariant_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgVerifyInvariant)
}
func (x fastReflection_MsgVerifyInvariant_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgVerifyInvariant
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgVerifyInvariant) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgVerifyInvariant
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgVerifyInvariant) Type() protoreflect.MessageType {
	return _fastReflection_MsgVerifyInvariant_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgVerifyInvariant) New() protoreflect.Message {
	return new(fastReflection_MsgVerifyInvariant)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgVerifyInvariant) Interface() protoreflect.ProtoMessage {
	return (*MsgVerifyInvariant)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgVerifyInvariant) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sender != "" {
		value := protoreflect.ValueOfString(x.Sender)
		if !f(fd_MsgVerifyInvariant_sender, value) {
			return
		}
	}
	if x.InvariantModuleName != "" {
		value := protoreflect.ValueOfString(x.InvariantModuleName)
		if !f(fd_MsgVerifyInvariant_invariant_module_name, value) {
			return
		}
	}
	if x.InvariantRoute != "" {
		value := protoreflect.ValueOfString(x.InvariantRoute)
		if !f(fd_MsgVerifyInvariant_invariant_route, value) {
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
func (x *fastReflection_MsgVerifyInvariant) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.sender":
		return x.Sender != ""
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.invariant_module_name":
		return x.InvariantModuleName != ""
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.invariant_route":
		return x.InvariantRoute != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crisis.v1beta1.MsgVerifyInvariant"))
		}
		panic(fmt.Errorf("message cosmos.crisis.v1beta1.MsgVerifyInvariant does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVerifyInvariant) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.sender":
		x.Sender = ""
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.invariant_module_name":
		x.InvariantModuleName = ""
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.invariant_route":
		x.InvariantRoute = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crisis.v1beta1.MsgVerifyInvariant"))
		}
		panic(fmt.Errorf("message cosmos.crisis.v1beta1.MsgVerifyInvariant does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgVerifyInvariant) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.sender":
		value := x.Sender
		return protoreflect.ValueOfString(value)
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.invariant_module_name":
		value := x.InvariantModuleName
		return protoreflect.ValueOfString(value)
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.invariant_route":
		value := x.InvariantRoute
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crisis.v1beta1.MsgVerifyInvariant"))
		}
		panic(fmt.Errorf("message cosmos.crisis.v1beta1.MsgVerifyInvariant does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgVerifyInvariant) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.sender":
		x.Sender = value.Interface().(string)
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.invariant_module_name":
		x.InvariantModuleName = value.Interface().(string)
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.invariant_route":
		x.InvariantRoute = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crisis.v1beta1.MsgVerifyInvariant"))
		}
		panic(fmt.Errorf("message cosmos.crisis.v1beta1.MsgVerifyInvariant does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgVerifyInvariant) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.sender":
		panic(fmt.Errorf("field sender of message cosmos.crisis.v1beta1.MsgVerifyInvariant is not mutable"))
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.invariant_module_name":
		panic(fmt.Errorf("field invariant_module_name of message cosmos.crisis.v1beta1.MsgVerifyInvariant is not mutable"))
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.invariant_route":
		panic(fmt.Errorf("field invariant_route of message cosmos.crisis.v1beta1.MsgVerifyInvariant is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crisis.v1beta1.MsgVerifyInvariant"))
		}
		panic(fmt.Errorf("message cosmos.crisis.v1beta1.MsgVerifyInvariant does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgVerifyInvariant) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.sender":
		return protoreflect.ValueOfString("")
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.invariant_module_name":
		return protoreflect.ValueOfString("")
	case "cosmos.crisis.v1beta1.MsgVerifyInvariant.invariant_route":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crisis.v1beta1.MsgVerifyInvariant"))
		}
		panic(fmt.Errorf("message cosmos.crisis.v1beta1.MsgVerifyInvariant does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgVerifyInvariant) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crisis.v1beta1.MsgVerifyInvariant", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgVerifyInvariant) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVerifyInvariant) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgVerifyInvariant) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgVerifyInvariant) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgVerifyInvariant)
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
		l = len(x.Sender)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.InvariantModuleName)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.InvariantRoute)
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
		x := input.Message.Interface().(*MsgVerifyInvariant)
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
		if len(x.InvariantRoute) > 0 {
			i -= len(x.InvariantRoute)
			copy(dAtA[i:], x.InvariantRoute)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.InvariantRoute)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.InvariantModuleName) > 0 {
			i -= len(x.InvariantModuleName)
			copy(dAtA[i:], x.InvariantModuleName)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.InvariantModuleName)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Sender) > 0 {
			i -= len(x.Sender)
			copy(dAtA[i:], x.Sender)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Sender)))
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
		x := input.Message.Interface().(*MsgVerifyInvariant)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgVerifyInvariant: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgVerifyInvariant: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
				x.Sender = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field InvariantModuleName", wireType)
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
				x.InvariantModuleName = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field InvariantRoute", wireType)
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
				x.InvariantRoute = string(dAtA[iNdEx:postIndex])
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
	md_MsgVerifyInvariantResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_crisis_v1beta1_tx_proto_init()
	md_MsgVerifyInvariantResponse = File_cosmos_crisis_v1beta1_tx_proto.Messages().ByName("MsgVerifyInvariantResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgVerifyInvariantResponse)(nil)

type fastReflection_MsgVerifyInvariantResponse MsgVerifyInvariantResponse

func (x *MsgVerifyInvariantResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgVerifyInvariantResponse)(x)
}

func (x *MsgVerifyInvariantResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crisis_v1beta1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgVerifyInvariantResponse_messageType fastReflection_MsgVerifyInvariantResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgVerifyInvariantResponse_messageType{}

type fastReflection_MsgVerifyInvariantResponse_messageType struct{}

func (x fastReflection_MsgVerifyInvariantResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgVerifyInvariantResponse)(nil)
}
func (x fastReflection_MsgVerifyInvariantResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgVerifyInvariantResponse)
}
func (x fastReflection_MsgVerifyInvariantResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgVerifyInvariantResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgVerifyInvariantResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgVerifyInvariantResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgVerifyInvariantResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgVerifyInvariantResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgVerifyInvariantResponse) New() protoreflect.Message {
	return new(fastReflection_MsgVerifyInvariantResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgVerifyInvariantResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgVerifyInvariantResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgVerifyInvariantResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgVerifyInvariantResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crisis.v1beta1.MsgVerifyInvariantResponse"))
		}
		panic(fmt.Errorf("message cosmos.crisis.v1beta1.MsgVerifyInvariantResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVerifyInvariantResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crisis.v1beta1.MsgVerifyInvariantResponse"))
		}
		panic(fmt.Errorf("message cosmos.crisis.v1beta1.MsgVerifyInvariantResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgVerifyInvariantResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crisis.v1beta1.MsgVerifyInvariantResponse"))
		}
		panic(fmt.Errorf("message cosmos.crisis.v1beta1.MsgVerifyInvariantResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgVerifyInvariantResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crisis.v1beta1.MsgVerifyInvariantResponse"))
		}
		panic(fmt.Errorf("message cosmos.crisis.v1beta1.MsgVerifyInvariantResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgVerifyInvariantResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crisis.v1beta1.MsgVerifyInvariantResponse"))
		}
		panic(fmt.Errorf("message cosmos.crisis.v1beta1.MsgVerifyInvariantResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgVerifyInvariantResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crisis.v1beta1.MsgVerifyInvariantResponse"))
		}
		panic(fmt.Errorf("message cosmos.crisis.v1beta1.MsgVerifyInvariantResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgVerifyInvariantResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crisis.v1beta1.MsgVerifyInvariantResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgVerifyInvariantResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVerifyInvariantResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgVerifyInvariantResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgVerifyInvariantResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgVerifyInvariantResponse)
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
		x := input.Message.Interface().(*MsgVerifyInvariantResponse)
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
		x := input.Message.Interface().(*MsgVerifyInvariantResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgVerifyInvariantResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgVerifyInvariantResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// VerifyInvariant defines a method to verify a particular invariance.
	VerifyInvariant(ctx context.Context, in *MsgVerifyInvariant, opts ...grpc.CallOption) (*MsgVerifyInvariantResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) VerifyInvariant(ctx context.Context, in *MsgVerifyInvariant, opts ...grpc.CallOption) (*MsgVerifyInvariantResponse, error) {
	out := new(MsgVerifyInvariantResponse)
	err := c.cc.Invoke(ctx, "/cosmos.crisis.v1beta1.Msg/VerifyInvariant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// VerifyInvariant defines a method to verify a particular invariance.
	VerifyInvariant(context.Context, *MsgVerifyInvariant) (*MsgVerifyInvariantResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) VerifyInvariant(context.Context, *MsgVerifyInvariant) (*MsgVerifyInvariantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyInvariant not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_VerifyInvariant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVerifyInvariant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VerifyInvariant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.crisis.v1beta1.Msg/VerifyInvariant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VerifyInvariant(ctx, req.(*MsgVerifyInvariant))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.crisis.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyInvariant",
			Handler:    _Msg_VerifyInvariant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/crisis/v1beta1/tx.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/crisis/v1beta1/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MsgVerifyInvariant represents a message to verify a particular invariance.
type MsgVerifyInvariant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender              string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	InvariantModuleName string `protobuf:"bytes,2,opt,name=invariant_module_name,json=invariantModuleName,proto3" json:"invariant_module_name,omitempty"`
	InvariantRoute      string `protobuf:"bytes,3,opt,name=invariant_route,json=invariantRoute,proto3" json:"invariant_route,omitempty"`
}

func (x *MsgVerifyInvariant) Reset() {
	*x = MsgVerifyInvariant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crisis_v1beta1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgVerifyInvariant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgVerifyInvariant) ProtoMessage() {}

// Deprecated: Use MsgVerifyInvariant.ProtoReflect.Descriptor instead.
func (*MsgVerifyInvariant) Descriptor() ([]byte, []int) {
	return file_cosmos_crisis_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgVerifyInvariant) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MsgVerifyInvariant) GetInvariantModuleName() string {
	if x != nil {
		return x.InvariantModuleName
	}
	return ""
}

func (x *MsgVerifyInvariant) GetInvariantRoute() string {
	if x != nil {
		return x.InvariantRoute
	}
	return ""
}

// MsgVerifyInvariantResponse defines the Msg/VerifyInvariant response type.
type MsgVerifyInvariantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgVerifyInvariantResponse) Reset() {
	*x = MsgVerifyInvariantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crisis_v1beta1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgVerifyInvariantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgVerifyInvariantResponse) ProtoMessage() {}

// Deprecated: Use MsgVerifyInvariantResponse.ProtoReflect.Descriptor instead.
func (*MsgVerifyInvariantResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_crisis_v1beta1_tx_proto_rawDescGZIP(), []int{1}
}

var File_cosmos_crisis_v1beta1_tx_proto protoreflect.FileDescriptor

var file_cosmos_crisis_v1beta1_tx_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x69, 0x73, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x69, 0x73, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x12, 0x4d, 0x73, 0x67,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12,
	0x30, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x3a, 0x08,
	0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x1c, 0x0a, 0x1a, 0x4d, 0x73, 0x67, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x76, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x6f, 0x0a,
	0x0f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x12, 0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x69, 0x73, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x49, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x69, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xe0,
	0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72,
	0x69, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x07, 0x54, 0x78,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x63, 0x72, 0x69, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b,
	0x63, 0x72, 0x69, 0x73, 0x69, 0x73, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x43, 0x58, 0xaa, 0x02, 0x15, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x69,
	0x73, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x15, 0x43, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x43, 0x72, 0x69, 0x73, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xe2, 0x02, 0x21, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x43, 0x72, 0x69,
	0x73, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x3a, 0x3a, 0x43, 0x72, 0x69, 0x73, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_crisis_v1beta1_tx_proto_rawDescOnce sync.Once
	file_cosmos_crisis_v1beta1_tx_proto_rawDescData = file_cosmos_crisis_v1beta1_tx_proto_rawDesc
)

func file_cosmos_crisis_v1beta1_tx_proto_rawDescGZIP() []byte {
	file_cosmos_crisis_v1beta1_tx_proto_rawDescOnce.Do(func() {
		file_cosmos_crisis_v1beta1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_crisis_v1beta1_tx_proto_rawDescData)
	})
	return file_cosmos_crisis_v1beta1_tx_proto_rawDescData
}

var file_cosmos_crisis_v1beta1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cosmos_crisis_v1beta1_tx_proto_goTypes = []interface{}{
	(*MsgVerifyInvariant)(nil),         // 0: cosmos.crisis.v1beta1.MsgVerifyInvariant
	(*MsgVerifyInvariantResponse)(nil), // 1: cosmos.crisis.v1beta1.MsgVerifyInvariantResponse
}
var file_cosmos_crisis_v1beta1_tx_proto_depIdxs = []int32{
	0, // 0: cosmos.crisis.v1beta1.Msg.VerifyInvariant:input_type -> cosmos.crisis.v1beta1.MsgVerifyInvariant
	1, // 1: cosmos.crisis.v1beta1.Msg.VerifyInvariant:output_type -> cosmos.crisis.v1beta1.MsgVerifyInvariantResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_crisis_v1beta1_tx_proto_init() }
func file_cosmos_crisis_v1beta1_tx_proto_init() {
	if File_cosmos_crisis_v1beta1_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_crisis_v1beta1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgVerifyInvariant); i {
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
		file_cosmos_crisis_v1beta1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgVerifyInvariantResponse); i {
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
			RawDescriptor: file_cosmos_crisis_v1beta1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_crisis_v1beta1_tx_proto_goTypes,
		DependencyIndexes: file_cosmos_crisis_v1beta1_tx_proto_depIdxs,
		MessageInfos:      file_cosmos_crisis_v1beta1_tx_proto_msgTypes,
	}.Build()
	File_cosmos_crisis_v1beta1_tx_proto = out.File
	file_cosmos_crisis_v1beta1_tx_proto_rawDesc = nil
	file_cosmos_crisis_v1beta1_tx_proto_goTypes = nil
	file_cosmos_crisis_v1beta1_tx_proto_depIdxs = nil
}
