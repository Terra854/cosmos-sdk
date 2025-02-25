package evidencev1beta1

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
	anypb "google.golang.org/protobuf/types/known/anypb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_MsgSubmitEvidence           protoreflect.MessageDescriptor
	fd_MsgSubmitEvidence_submitter protoreflect.FieldDescriptor
	fd_MsgSubmitEvidence_evidence  protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_evidence_v1beta1_tx_proto_init()
	md_MsgSubmitEvidence = File_cosmos_evidence_v1beta1_tx_proto.Messages().ByName("MsgSubmitEvidence")
	fd_MsgSubmitEvidence_submitter = md_MsgSubmitEvidence.Fields().ByName("submitter")
	fd_MsgSubmitEvidence_evidence = md_MsgSubmitEvidence.Fields().ByName("evidence")
}

var _ protoreflect.Message = (*fastReflection_MsgSubmitEvidence)(nil)

type fastReflection_MsgSubmitEvidence MsgSubmitEvidence

func (x *MsgSubmitEvidence) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSubmitEvidence)(x)
}

func (x *MsgSubmitEvidence) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_evidence_v1beta1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSubmitEvidence_messageType fastReflection_MsgSubmitEvidence_messageType
var _ protoreflect.MessageType = fastReflection_MsgSubmitEvidence_messageType{}

type fastReflection_MsgSubmitEvidence_messageType struct{}

func (x fastReflection_MsgSubmitEvidence_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSubmitEvidence)(nil)
}
func (x fastReflection_MsgSubmitEvidence_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSubmitEvidence)
}
func (x fastReflection_MsgSubmitEvidence_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSubmitEvidence
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSubmitEvidence) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSubmitEvidence
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSubmitEvidence) Type() protoreflect.MessageType {
	return _fastReflection_MsgSubmitEvidence_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSubmitEvidence) New() protoreflect.Message {
	return new(fastReflection_MsgSubmitEvidence)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSubmitEvidence) Interface() protoreflect.ProtoMessage {
	return (*MsgSubmitEvidence)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSubmitEvidence) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Submitter != "" {
		value := protoreflect.ValueOfString(x.Submitter)
		if !f(fd_MsgSubmitEvidence_submitter, value) {
			return
		}
	}
	if x.Evidence != nil {
		value := protoreflect.ValueOfMessage(x.Evidence.ProtoReflect())
		if !f(fd_MsgSubmitEvidence_evidence, value) {
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
func (x *fastReflection_MsgSubmitEvidence) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.MsgSubmitEvidence.submitter":
		return x.Submitter != ""
	case "cosmos.evidence.v1beta1.MsgSubmitEvidence.evidence":
		return x.Evidence != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.MsgSubmitEvidence"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.MsgSubmitEvidence does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSubmitEvidence) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.MsgSubmitEvidence.submitter":
		x.Submitter = ""
	case "cosmos.evidence.v1beta1.MsgSubmitEvidence.evidence":
		x.Evidence = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.MsgSubmitEvidence"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.MsgSubmitEvidence does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSubmitEvidence) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.evidence.v1beta1.MsgSubmitEvidence.submitter":
		value := x.Submitter
		return protoreflect.ValueOfString(value)
	case "cosmos.evidence.v1beta1.MsgSubmitEvidence.evidence":
		value := x.Evidence
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.MsgSubmitEvidence"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.MsgSubmitEvidence does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSubmitEvidence) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.MsgSubmitEvidence.submitter":
		x.Submitter = value.Interface().(string)
	case "cosmos.evidence.v1beta1.MsgSubmitEvidence.evidence":
		x.Evidence = value.Message().Interface().(*anypb.Any)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.MsgSubmitEvidence"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.MsgSubmitEvidence does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSubmitEvidence) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.MsgSubmitEvidence.evidence":
		if x.Evidence == nil {
			x.Evidence = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.Evidence.ProtoReflect())
	case "cosmos.evidence.v1beta1.MsgSubmitEvidence.submitter":
		panic(fmt.Errorf("field submitter of message cosmos.evidence.v1beta1.MsgSubmitEvidence is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.MsgSubmitEvidence"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.MsgSubmitEvidence does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSubmitEvidence) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.MsgSubmitEvidence.submitter":
		return protoreflect.ValueOfString("")
	case "cosmos.evidence.v1beta1.MsgSubmitEvidence.evidence":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.MsgSubmitEvidence"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.MsgSubmitEvidence does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSubmitEvidence) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.evidence.v1beta1.MsgSubmitEvidence", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSubmitEvidence) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSubmitEvidence) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSubmitEvidence) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSubmitEvidence) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSubmitEvidence)
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
		l = len(x.Submitter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Evidence != nil {
			l = options.Size(x.Evidence)
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
		x := input.Message.Interface().(*MsgSubmitEvidence)
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
		if x.Evidence != nil {
			encoded, err := options.Marshal(x.Evidence)
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
		if len(x.Submitter) > 0 {
			i -= len(x.Submitter)
			copy(dAtA[i:], x.Submitter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Submitter)))
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
		x := input.Message.Interface().(*MsgSubmitEvidence)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSubmitEvidence: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSubmitEvidence: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Submitter", wireType)
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
				x.Submitter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
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
				if x.Evidence == nil {
					x.Evidence = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Evidence); err != nil {
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
	md_MsgSubmitEvidenceResponse      protoreflect.MessageDescriptor
	fd_MsgSubmitEvidenceResponse_hash protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_evidence_v1beta1_tx_proto_init()
	md_MsgSubmitEvidenceResponse = File_cosmos_evidence_v1beta1_tx_proto.Messages().ByName("MsgSubmitEvidenceResponse")
	fd_MsgSubmitEvidenceResponse_hash = md_MsgSubmitEvidenceResponse.Fields().ByName("hash")
}

var _ protoreflect.Message = (*fastReflection_MsgSubmitEvidenceResponse)(nil)

type fastReflection_MsgSubmitEvidenceResponse MsgSubmitEvidenceResponse

func (x *MsgSubmitEvidenceResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSubmitEvidenceResponse)(x)
}

func (x *MsgSubmitEvidenceResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_evidence_v1beta1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSubmitEvidenceResponse_messageType fastReflection_MsgSubmitEvidenceResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgSubmitEvidenceResponse_messageType{}

type fastReflection_MsgSubmitEvidenceResponse_messageType struct{}

func (x fastReflection_MsgSubmitEvidenceResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSubmitEvidenceResponse)(nil)
}
func (x fastReflection_MsgSubmitEvidenceResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSubmitEvidenceResponse)
}
func (x fastReflection_MsgSubmitEvidenceResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSubmitEvidenceResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSubmitEvidenceResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSubmitEvidenceResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSubmitEvidenceResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgSubmitEvidenceResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSubmitEvidenceResponse) New() protoreflect.Message {
	return new(fastReflection_MsgSubmitEvidenceResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSubmitEvidenceResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgSubmitEvidenceResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSubmitEvidenceResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Hash) != 0 {
		value := protoreflect.ValueOfBytes(x.Hash)
		if !f(fd_MsgSubmitEvidenceResponse_hash, value) {
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
func (x *fastReflection_MsgSubmitEvidenceResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.hash":
		return len(x.Hash) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSubmitEvidenceResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.hash":
		x.Hash = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSubmitEvidenceResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.hash":
		value := x.Hash
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSubmitEvidenceResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.hash":
		x.Hash = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSubmitEvidenceResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.hash":
		panic(fmt.Errorf("field hash of message cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSubmitEvidenceResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.hash":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse"))
		}
		panic(fmt.Errorf("message cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSubmitEvidenceResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSubmitEvidenceResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSubmitEvidenceResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSubmitEvidenceResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSubmitEvidenceResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSubmitEvidenceResponse)
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
		l = len(x.Hash)
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
		x := input.Message.Interface().(*MsgSubmitEvidenceResponse)
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
		if len(x.Hash) > 0 {
			i -= len(x.Hash)
			copy(dAtA[i:], x.Hash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Hash)))
			i--
			dAtA[i] = 0x22
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
		x := input.Message.Interface().(*MsgSubmitEvidenceResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSubmitEvidenceResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSubmitEvidenceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Hash = append(x.Hash[:0], dAtA[iNdEx:postIndex]...)
				if x.Hash == nil {
					x.Hash = []byte{}
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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// SubmitEvidence submits an arbitrary Evidence of misbehavior such as equivocation or
	// counterfactual signing.
	SubmitEvidence(ctx context.Context, in *MsgSubmitEvidence, opts ...grpc.CallOption) (*MsgSubmitEvidenceResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitEvidence(ctx context.Context, in *MsgSubmitEvidence, opts ...grpc.CallOption) (*MsgSubmitEvidenceResponse, error) {
	out := new(MsgSubmitEvidenceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.evidence.v1beta1.Msg/SubmitEvidence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// SubmitEvidence submits an arbitrary Evidence of misbehavior such as equivocation or
	// counterfactual signing.
	SubmitEvidence(context.Context, *MsgSubmitEvidence) (*MsgSubmitEvidenceResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) SubmitEvidence(context.Context, *MsgSubmitEvidence) (*MsgSubmitEvidenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitEvidence not implemented")
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

func _Msg_SubmitEvidence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitEvidence)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitEvidence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.evidence.v1beta1.Msg/SubmitEvidence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitEvidence(ctx, req.(*MsgSubmitEvidence))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.evidence.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitEvidence",
			Handler:    _Msg_SubmitEvidence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/evidence/v1beta1/tx.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/evidence/v1beta1/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MsgSubmitEvidence represents a message that supports submitting arbitrary
// Evidence of misbehavior such as equivocation or counterfactual signing.
type MsgSubmitEvidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submitter string     `protobuf:"bytes,1,opt,name=submitter,proto3" json:"submitter,omitempty"`
	Evidence  *anypb.Any `protobuf:"bytes,2,opt,name=evidence,proto3" json:"evidence,omitempty"`
}

func (x *MsgSubmitEvidence) Reset() {
	*x = MsgSubmitEvidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_evidence_v1beta1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSubmitEvidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSubmitEvidence) ProtoMessage() {}

// Deprecated: Use MsgSubmitEvidence.ProtoReflect.Descriptor instead.
func (*MsgSubmitEvidence) Descriptor() ([]byte, []int) {
	return file_cosmos_evidence_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgSubmitEvidence) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

func (x *MsgSubmitEvidence) GetEvidence() *anypb.Any {
	if x != nil {
		return x.Evidence
	}
	return nil
}

// MsgSubmitEvidenceResponse defines the Msg/SubmitEvidence response type.
type MsgSubmitEvidenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hash defines the hash of the evidence.
	Hash []byte `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *MsgSubmitEvidenceResponse) Reset() {
	*x = MsgSubmitEvidenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_evidence_v1beta1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSubmitEvidenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSubmitEvidenceResponse) ProtoMessage() {}

// Deprecated: Use MsgSubmitEvidenceResponse.ProtoReflect.Descriptor instead.
func (*MsgSubmitEvidenceResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_evidence_v1beta1_tx_proto_rawDescGZIP(), []int{1}
}

func (x *MsgSubmitEvidenceResponse) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

var File_cosmos_evidence_v1beta1_tx_proto protoreflect.FileDescriptor

var file_cosmos_evidence_v1beta1_tx_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x65, 0x76, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x36, 0x0a,
	0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x0c, 0xca,
	0xb4, 0x2d, 0x08, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x65, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22,
	0x2f, 0x0a, 0x19, 0x4d, 0x73, 0x67, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x32, 0x77, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x70, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x65,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x4d, 0x73, 0x67, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xf2, 0x01, 0x0a, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x07, 0x54, 0x78, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x65, 0x76,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x65,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x45, 0x58, 0xaa, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x45, 0x76,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02,
	0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x23, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa8, 0xe2, 0x1e, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_evidence_v1beta1_tx_proto_rawDescOnce sync.Once
	file_cosmos_evidence_v1beta1_tx_proto_rawDescData = file_cosmos_evidence_v1beta1_tx_proto_rawDesc
)

func file_cosmos_evidence_v1beta1_tx_proto_rawDescGZIP() []byte {
	file_cosmos_evidence_v1beta1_tx_proto_rawDescOnce.Do(func() {
		file_cosmos_evidence_v1beta1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_evidence_v1beta1_tx_proto_rawDescData)
	})
	return file_cosmos_evidence_v1beta1_tx_proto_rawDescData
}

var file_cosmos_evidence_v1beta1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cosmos_evidence_v1beta1_tx_proto_goTypes = []interface{}{
	(*MsgSubmitEvidence)(nil),         // 0: cosmos.evidence.v1beta1.MsgSubmitEvidence
	(*MsgSubmitEvidenceResponse)(nil), // 1: cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse
	(*anypb.Any)(nil),                 // 2: google.protobuf.Any
}
var file_cosmos_evidence_v1beta1_tx_proto_depIdxs = []int32{
	2, // 0: cosmos.evidence.v1beta1.MsgSubmitEvidence.evidence:type_name -> google.protobuf.Any
	0, // 1: cosmos.evidence.v1beta1.Msg.SubmitEvidence:input_type -> cosmos.evidence.v1beta1.MsgSubmitEvidence
	1, // 2: cosmos.evidence.v1beta1.Msg.SubmitEvidence:output_type -> cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cosmos_evidence_v1beta1_tx_proto_init() }
func file_cosmos_evidence_v1beta1_tx_proto_init() {
	if File_cosmos_evidence_v1beta1_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_evidence_v1beta1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSubmitEvidence); i {
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
		file_cosmos_evidence_v1beta1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSubmitEvidenceResponse); i {
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
			RawDescriptor: file_cosmos_evidence_v1beta1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_evidence_v1beta1_tx_proto_goTypes,
		DependencyIndexes: file_cosmos_evidence_v1beta1_tx_proto_depIdxs,
		MessageInfos:      file_cosmos_evidence_v1beta1_tx_proto_msgTypes,
	}.Build()
	File_cosmos_evidence_v1beta1_tx_proto = out.File
	file_cosmos_evidence_v1beta1_tx_proto_rawDesc = nil
	file_cosmos_evidence_v1beta1_tx_proto_goTypes = nil
	file_cosmos_evidence_v1beta1_tx_proto_depIdxs = nil
}
