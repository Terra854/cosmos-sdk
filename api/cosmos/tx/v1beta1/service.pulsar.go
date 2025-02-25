package txv1beta1

import (
	context "context"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta11 "github.com/cosmos/cosmos-sdk/api/cosmos/base/abci/v1beta1"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/query/v1beta1"
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

var _ protoreflect.List = (*_GetTxsEventRequest_1_list)(nil)

type _GetTxsEventRequest_1_list struct {
	list *[]string
}

func (x *_GetTxsEventRequest_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GetTxsEventRequest_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_GetTxsEventRequest_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_GetTxsEventRequest_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_GetTxsEventRequest_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message GetTxsEventRequest at list field Events as it is not of Message kind"))
}

func (x *_GetTxsEventRequest_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_GetTxsEventRequest_1_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_GetTxsEventRequest_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_GetTxsEventRequest            protoreflect.MessageDescriptor
	fd_GetTxsEventRequest_events     protoreflect.FieldDescriptor
	fd_GetTxsEventRequest_pagination protoreflect.FieldDescriptor
	fd_GetTxsEventRequest_order_by   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_service_proto_init()
	md_GetTxsEventRequest = File_cosmos_tx_v1beta1_service_proto.Messages().ByName("GetTxsEventRequest")
	fd_GetTxsEventRequest_events = md_GetTxsEventRequest.Fields().ByName("events")
	fd_GetTxsEventRequest_pagination = md_GetTxsEventRequest.Fields().ByName("pagination")
	fd_GetTxsEventRequest_order_by = md_GetTxsEventRequest.Fields().ByName("order_by")
}

var _ protoreflect.Message = (*fastReflection_GetTxsEventRequest)(nil)

type fastReflection_GetTxsEventRequest GetTxsEventRequest

func (x *GetTxsEventRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetTxsEventRequest)(x)
}

func (x *GetTxsEventRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetTxsEventRequest_messageType fastReflection_GetTxsEventRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetTxsEventRequest_messageType{}

type fastReflection_GetTxsEventRequest_messageType struct{}

func (x fastReflection_GetTxsEventRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetTxsEventRequest)(nil)
}
func (x fastReflection_GetTxsEventRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetTxsEventRequest)
}
func (x fastReflection_GetTxsEventRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetTxsEventRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetTxsEventRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetTxsEventRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetTxsEventRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetTxsEventRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetTxsEventRequest) New() protoreflect.Message {
	return new(fastReflection_GetTxsEventRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetTxsEventRequest) Interface() protoreflect.ProtoMessage {
	return (*GetTxsEventRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetTxsEventRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Events) != 0 {
		value := protoreflect.ValueOfList(&_GetTxsEventRequest_1_list{list: &x.Events})
		if !f(fd_GetTxsEventRequest_events, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_GetTxsEventRequest_pagination, value) {
			return
		}
	}
	if x.OrderBy != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.OrderBy))
		if !f(fd_GetTxsEventRequest_order_by, value) {
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
func (x *fastReflection_GetTxsEventRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxsEventRequest.events":
		return len(x.Events) != 0
	case "cosmos.tx.v1beta1.GetTxsEventRequest.pagination":
		return x.Pagination != nil
	case "cosmos.tx.v1beta1.GetTxsEventRequest.order_by":
		return x.OrderBy != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxsEventRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxsEventRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetTxsEventRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxsEventRequest.events":
		x.Events = nil
	case "cosmos.tx.v1beta1.GetTxsEventRequest.pagination":
		x.Pagination = nil
	case "cosmos.tx.v1beta1.GetTxsEventRequest.order_by":
		x.OrderBy = 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxsEventRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxsEventRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetTxsEventRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.GetTxsEventRequest.events":
		if len(x.Events) == 0 {
			return protoreflect.ValueOfList(&_GetTxsEventRequest_1_list{})
		}
		listValue := &_GetTxsEventRequest_1_list{list: &x.Events}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.tx.v1beta1.GetTxsEventRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.GetTxsEventRequest.order_by":
		value := x.OrderBy
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxsEventRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxsEventRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetTxsEventRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxsEventRequest.events":
		lv := value.List()
		clv := lv.(*_GetTxsEventRequest_1_list)
		x.Events = *clv.list
	case "cosmos.tx.v1beta1.GetTxsEventRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	case "cosmos.tx.v1beta1.GetTxsEventRequest.order_by":
		x.OrderBy = (OrderBy)(value.Enum())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxsEventRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxsEventRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetTxsEventRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxsEventRequest.events":
		if x.Events == nil {
			x.Events = []string{}
		}
		value := &_GetTxsEventRequest_1_list{list: &x.Events}
		return protoreflect.ValueOfList(value)
	case "cosmos.tx.v1beta1.GetTxsEventRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.tx.v1beta1.GetTxsEventRequest.order_by":
		panic(fmt.Errorf("field order_by of message cosmos.tx.v1beta1.GetTxsEventRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxsEventRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxsEventRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetTxsEventRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxsEventRequest.events":
		list := []string{}
		return protoreflect.ValueOfList(&_GetTxsEventRequest_1_list{list: &list})
	case "cosmos.tx.v1beta1.GetTxsEventRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.v1beta1.GetTxsEventRequest.order_by":
		return protoreflect.ValueOfEnum(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxsEventRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxsEventRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetTxsEventRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.GetTxsEventRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetTxsEventRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetTxsEventRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetTxsEventRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetTxsEventRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetTxsEventRequest)
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
		if len(x.Events) > 0 {
			for _, s := range x.Events {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.OrderBy != 0 {
			n += 1 + runtime.Sov(uint64(x.OrderBy))
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
		x := input.Message.Interface().(*GetTxsEventRequest)
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
		if x.OrderBy != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.OrderBy))
			i--
			dAtA[i] = 0x18
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
		if len(x.Events) > 0 {
			for iNdEx := len(x.Events) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Events[iNdEx])
				copy(dAtA[i:], x.Events[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Events[iNdEx])))
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
		x := input.Message.Interface().(*GetTxsEventRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetTxsEventRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetTxsEventRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
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
				x.Events = append(x.Events, string(dAtA[iNdEx:postIndex]))
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
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OrderBy", wireType)
				}
				x.OrderBy = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.OrderBy |= OrderBy(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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

var _ protoreflect.List = (*_GetTxsEventResponse_1_list)(nil)

type _GetTxsEventResponse_1_list struct {
	list *[]*Tx
}

func (x *_GetTxsEventResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GetTxsEventResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GetTxsEventResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Tx)
	(*x.list)[i] = concreteValue
}

func (x *_GetTxsEventResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Tx)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GetTxsEventResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(Tx)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GetTxsEventResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GetTxsEventResponse_1_list) NewElement() protoreflect.Value {
	v := new(Tx)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GetTxsEventResponse_1_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GetTxsEventResponse_2_list)(nil)

type _GetTxsEventResponse_2_list struct {
	list *[]*v1beta11.TxResponse
}

func (x *_GetTxsEventResponse_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GetTxsEventResponse_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GetTxsEventResponse_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta11.TxResponse)
	(*x.list)[i] = concreteValue
}

func (x *_GetTxsEventResponse_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta11.TxResponse)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GetTxsEventResponse_2_list) AppendMutable() protoreflect.Value {
	v := new(v1beta11.TxResponse)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GetTxsEventResponse_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GetTxsEventResponse_2_list) NewElement() protoreflect.Value {
	v := new(v1beta11.TxResponse)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GetTxsEventResponse_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_GetTxsEventResponse              protoreflect.MessageDescriptor
	fd_GetTxsEventResponse_txs          protoreflect.FieldDescriptor
	fd_GetTxsEventResponse_tx_responses protoreflect.FieldDescriptor
	fd_GetTxsEventResponse_pagination   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_service_proto_init()
	md_GetTxsEventResponse = File_cosmos_tx_v1beta1_service_proto.Messages().ByName("GetTxsEventResponse")
	fd_GetTxsEventResponse_txs = md_GetTxsEventResponse.Fields().ByName("txs")
	fd_GetTxsEventResponse_tx_responses = md_GetTxsEventResponse.Fields().ByName("tx_responses")
	fd_GetTxsEventResponse_pagination = md_GetTxsEventResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_GetTxsEventResponse)(nil)

type fastReflection_GetTxsEventResponse GetTxsEventResponse

func (x *GetTxsEventResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetTxsEventResponse)(x)
}

func (x *GetTxsEventResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetTxsEventResponse_messageType fastReflection_GetTxsEventResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetTxsEventResponse_messageType{}

type fastReflection_GetTxsEventResponse_messageType struct{}

func (x fastReflection_GetTxsEventResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetTxsEventResponse)(nil)
}
func (x fastReflection_GetTxsEventResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetTxsEventResponse)
}
func (x fastReflection_GetTxsEventResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetTxsEventResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetTxsEventResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetTxsEventResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetTxsEventResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetTxsEventResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetTxsEventResponse) New() protoreflect.Message {
	return new(fastReflection_GetTxsEventResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetTxsEventResponse) Interface() protoreflect.ProtoMessage {
	return (*GetTxsEventResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetTxsEventResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Txs) != 0 {
		value := protoreflect.ValueOfList(&_GetTxsEventResponse_1_list{list: &x.Txs})
		if !f(fd_GetTxsEventResponse_txs, value) {
			return
		}
	}
	if len(x.TxResponses) != 0 {
		value := protoreflect.ValueOfList(&_GetTxsEventResponse_2_list{list: &x.TxResponses})
		if !f(fd_GetTxsEventResponse_tx_responses, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_GetTxsEventResponse_pagination, value) {
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
func (x *fastReflection_GetTxsEventResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxsEventResponse.txs":
		return len(x.Txs) != 0
	case "cosmos.tx.v1beta1.GetTxsEventResponse.tx_responses":
		return len(x.TxResponses) != 0
	case "cosmos.tx.v1beta1.GetTxsEventResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxsEventResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxsEventResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetTxsEventResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxsEventResponse.txs":
		x.Txs = nil
	case "cosmos.tx.v1beta1.GetTxsEventResponse.tx_responses":
		x.TxResponses = nil
	case "cosmos.tx.v1beta1.GetTxsEventResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxsEventResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxsEventResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetTxsEventResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.GetTxsEventResponse.txs":
		if len(x.Txs) == 0 {
			return protoreflect.ValueOfList(&_GetTxsEventResponse_1_list{})
		}
		listValue := &_GetTxsEventResponse_1_list{list: &x.Txs}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.tx.v1beta1.GetTxsEventResponse.tx_responses":
		if len(x.TxResponses) == 0 {
			return protoreflect.ValueOfList(&_GetTxsEventResponse_2_list{})
		}
		listValue := &_GetTxsEventResponse_2_list{list: &x.TxResponses}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.tx.v1beta1.GetTxsEventResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxsEventResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxsEventResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetTxsEventResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxsEventResponse.txs":
		lv := value.List()
		clv := lv.(*_GetTxsEventResponse_1_list)
		x.Txs = *clv.list
	case "cosmos.tx.v1beta1.GetTxsEventResponse.tx_responses":
		lv := value.List()
		clv := lv.(*_GetTxsEventResponse_2_list)
		x.TxResponses = *clv.list
	case "cosmos.tx.v1beta1.GetTxsEventResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxsEventResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxsEventResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetTxsEventResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxsEventResponse.txs":
		if x.Txs == nil {
			x.Txs = []*Tx{}
		}
		value := &_GetTxsEventResponse_1_list{list: &x.Txs}
		return protoreflect.ValueOfList(value)
	case "cosmos.tx.v1beta1.GetTxsEventResponse.tx_responses":
		if x.TxResponses == nil {
			x.TxResponses = []*v1beta11.TxResponse{}
		}
		value := &_GetTxsEventResponse_2_list{list: &x.TxResponses}
		return protoreflect.ValueOfList(value)
	case "cosmos.tx.v1beta1.GetTxsEventResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxsEventResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxsEventResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetTxsEventResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxsEventResponse.txs":
		list := []*Tx{}
		return protoreflect.ValueOfList(&_GetTxsEventResponse_1_list{list: &list})
	case "cosmos.tx.v1beta1.GetTxsEventResponse.tx_responses":
		list := []*v1beta11.TxResponse{}
		return protoreflect.ValueOfList(&_GetTxsEventResponse_2_list{list: &list})
	case "cosmos.tx.v1beta1.GetTxsEventResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxsEventResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxsEventResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetTxsEventResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.GetTxsEventResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetTxsEventResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetTxsEventResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetTxsEventResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetTxsEventResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetTxsEventResponse)
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
		if len(x.Txs) > 0 {
			for _, e := range x.Txs {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.TxResponses) > 0 {
			for _, e := range x.TxResponses {
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
		x := input.Message.Interface().(*GetTxsEventResponse)
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
			dAtA[i] = 0x1a
		}
		if len(x.TxResponses) > 0 {
			for iNdEx := len(x.TxResponses) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.TxResponses[iNdEx])
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
		if len(x.Txs) > 0 {
			for iNdEx := len(x.Txs) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Txs[iNdEx])
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
		x := input.Message.Interface().(*GetTxsEventResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetTxsEventResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetTxsEventResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
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
				x.Txs = append(x.Txs, &Tx{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Txs[len(x.Txs)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TxResponses", wireType)
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
				x.TxResponses = append(x.TxResponses, &v1beta11.TxResponse{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.TxResponses[len(x.TxResponses)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
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
	md_BroadcastTxRequest          protoreflect.MessageDescriptor
	fd_BroadcastTxRequest_tx_bytes protoreflect.FieldDescriptor
	fd_BroadcastTxRequest_mode     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_service_proto_init()
	md_BroadcastTxRequest = File_cosmos_tx_v1beta1_service_proto.Messages().ByName("BroadcastTxRequest")
	fd_BroadcastTxRequest_tx_bytes = md_BroadcastTxRequest.Fields().ByName("tx_bytes")
	fd_BroadcastTxRequest_mode = md_BroadcastTxRequest.Fields().ByName("mode")
}

var _ protoreflect.Message = (*fastReflection_BroadcastTxRequest)(nil)

type fastReflection_BroadcastTxRequest BroadcastTxRequest

func (x *BroadcastTxRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BroadcastTxRequest)(x)
}

func (x *BroadcastTxRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BroadcastTxRequest_messageType fastReflection_BroadcastTxRequest_messageType
var _ protoreflect.MessageType = fastReflection_BroadcastTxRequest_messageType{}

type fastReflection_BroadcastTxRequest_messageType struct{}

func (x fastReflection_BroadcastTxRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BroadcastTxRequest)(nil)
}
func (x fastReflection_BroadcastTxRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_BroadcastTxRequest)
}
func (x fastReflection_BroadcastTxRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BroadcastTxRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BroadcastTxRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_BroadcastTxRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BroadcastTxRequest) Type() protoreflect.MessageType {
	return _fastReflection_BroadcastTxRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BroadcastTxRequest) New() protoreflect.Message {
	return new(fastReflection_BroadcastTxRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BroadcastTxRequest) Interface() protoreflect.ProtoMessage {
	return (*BroadcastTxRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BroadcastTxRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.TxBytes) != 0 {
		value := protoreflect.ValueOfBytes(x.TxBytes)
		if !f(fd_BroadcastTxRequest_tx_bytes, value) {
			return
		}
	}
	if x.Mode != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Mode))
		if !f(fd_BroadcastTxRequest_mode, value) {
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
func (x *fastReflection_BroadcastTxRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.BroadcastTxRequest.tx_bytes":
		return len(x.TxBytes) != 0
	case "cosmos.tx.v1beta1.BroadcastTxRequest.mode":
		return x.Mode != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.BroadcastTxRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.BroadcastTxRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BroadcastTxRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.BroadcastTxRequest.tx_bytes":
		x.TxBytes = nil
	case "cosmos.tx.v1beta1.BroadcastTxRequest.mode":
		x.Mode = 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.BroadcastTxRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.BroadcastTxRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BroadcastTxRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.BroadcastTxRequest.tx_bytes":
		value := x.TxBytes
		return protoreflect.ValueOfBytes(value)
	case "cosmos.tx.v1beta1.BroadcastTxRequest.mode":
		value := x.Mode
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.BroadcastTxRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.BroadcastTxRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BroadcastTxRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.BroadcastTxRequest.tx_bytes":
		x.TxBytes = value.Bytes()
	case "cosmos.tx.v1beta1.BroadcastTxRequest.mode":
		x.Mode = (BroadcastMode)(value.Enum())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.BroadcastTxRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.BroadcastTxRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BroadcastTxRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.BroadcastTxRequest.tx_bytes":
		panic(fmt.Errorf("field tx_bytes of message cosmos.tx.v1beta1.BroadcastTxRequest is not mutable"))
	case "cosmos.tx.v1beta1.BroadcastTxRequest.mode":
		panic(fmt.Errorf("field mode of message cosmos.tx.v1beta1.BroadcastTxRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.BroadcastTxRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.BroadcastTxRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BroadcastTxRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.BroadcastTxRequest.tx_bytes":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.tx.v1beta1.BroadcastTxRequest.mode":
		return protoreflect.ValueOfEnum(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.BroadcastTxRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.BroadcastTxRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BroadcastTxRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.BroadcastTxRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BroadcastTxRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BroadcastTxRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BroadcastTxRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BroadcastTxRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BroadcastTxRequest)
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
		l = len(x.TxBytes)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Mode != 0 {
			n += 1 + runtime.Sov(uint64(x.Mode))
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
		x := input.Message.Interface().(*BroadcastTxRequest)
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
		if x.Mode != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Mode))
			i--
			dAtA[i] = 0x10
		}
		if len(x.TxBytes) > 0 {
			i -= len(x.TxBytes)
			copy(dAtA[i:], x.TxBytes)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TxBytes)))
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
		x := input.Message.Interface().(*BroadcastTxRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BroadcastTxRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BroadcastTxRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TxBytes", wireType)
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
				x.TxBytes = append(x.TxBytes[:0], dAtA[iNdEx:postIndex]...)
				if x.TxBytes == nil {
					x.TxBytes = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
				}
				x.Mode = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Mode |= BroadcastMode(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
	md_BroadcastTxResponse             protoreflect.MessageDescriptor
	fd_BroadcastTxResponse_tx_response protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_service_proto_init()
	md_BroadcastTxResponse = File_cosmos_tx_v1beta1_service_proto.Messages().ByName("BroadcastTxResponse")
	fd_BroadcastTxResponse_tx_response = md_BroadcastTxResponse.Fields().ByName("tx_response")
}

var _ protoreflect.Message = (*fastReflection_BroadcastTxResponse)(nil)

type fastReflection_BroadcastTxResponse BroadcastTxResponse

func (x *BroadcastTxResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BroadcastTxResponse)(x)
}

func (x *BroadcastTxResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BroadcastTxResponse_messageType fastReflection_BroadcastTxResponse_messageType
var _ protoreflect.MessageType = fastReflection_BroadcastTxResponse_messageType{}

type fastReflection_BroadcastTxResponse_messageType struct{}

func (x fastReflection_BroadcastTxResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BroadcastTxResponse)(nil)
}
func (x fastReflection_BroadcastTxResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_BroadcastTxResponse)
}
func (x fastReflection_BroadcastTxResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BroadcastTxResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BroadcastTxResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_BroadcastTxResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BroadcastTxResponse) Type() protoreflect.MessageType {
	return _fastReflection_BroadcastTxResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BroadcastTxResponse) New() protoreflect.Message {
	return new(fastReflection_BroadcastTxResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BroadcastTxResponse) Interface() protoreflect.ProtoMessage {
	return (*BroadcastTxResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BroadcastTxResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.TxResponse != nil {
		value := protoreflect.ValueOfMessage(x.TxResponse.ProtoReflect())
		if !f(fd_BroadcastTxResponse_tx_response, value) {
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
func (x *fastReflection_BroadcastTxResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.BroadcastTxResponse.tx_response":
		return x.TxResponse != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.BroadcastTxResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.BroadcastTxResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BroadcastTxResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.BroadcastTxResponse.tx_response":
		x.TxResponse = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.BroadcastTxResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.BroadcastTxResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BroadcastTxResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.BroadcastTxResponse.tx_response":
		value := x.TxResponse
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.BroadcastTxResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.BroadcastTxResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BroadcastTxResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.BroadcastTxResponse.tx_response":
		x.TxResponse = value.Message().Interface().(*v1beta11.TxResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.BroadcastTxResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.BroadcastTxResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BroadcastTxResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.BroadcastTxResponse.tx_response":
		if x.TxResponse == nil {
			x.TxResponse = new(v1beta11.TxResponse)
		}
		return protoreflect.ValueOfMessage(x.TxResponse.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.BroadcastTxResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.BroadcastTxResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BroadcastTxResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.BroadcastTxResponse.tx_response":
		m := new(v1beta11.TxResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.BroadcastTxResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.BroadcastTxResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BroadcastTxResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.BroadcastTxResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BroadcastTxResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BroadcastTxResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BroadcastTxResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BroadcastTxResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BroadcastTxResponse)
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
		if x.TxResponse != nil {
			l = options.Size(x.TxResponse)
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
		x := input.Message.Interface().(*BroadcastTxResponse)
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
		if x.TxResponse != nil {
			encoded, err := options.Marshal(x.TxResponse)
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
		x := input.Message.Interface().(*BroadcastTxResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BroadcastTxResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BroadcastTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TxResponse", wireType)
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
				if x.TxResponse == nil {
					x.TxResponse = &v1beta11.TxResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.TxResponse); err != nil {
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
	md_SimulateRequest          protoreflect.MessageDescriptor
	fd_SimulateRequest_tx       protoreflect.FieldDescriptor
	fd_SimulateRequest_tx_bytes protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_service_proto_init()
	md_SimulateRequest = File_cosmos_tx_v1beta1_service_proto.Messages().ByName("SimulateRequest")
	fd_SimulateRequest_tx = md_SimulateRequest.Fields().ByName("tx")
	fd_SimulateRequest_tx_bytes = md_SimulateRequest.Fields().ByName("tx_bytes")
}

var _ protoreflect.Message = (*fastReflection_SimulateRequest)(nil)

type fastReflection_SimulateRequest SimulateRequest

func (x *SimulateRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SimulateRequest)(x)
}

func (x *SimulateRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SimulateRequest_messageType fastReflection_SimulateRequest_messageType
var _ protoreflect.MessageType = fastReflection_SimulateRequest_messageType{}

type fastReflection_SimulateRequest_messageType struct{}

func (x fastReflection_SimulateRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SimulateRequest)(nil)
}
func (x fastReflection_SimulateRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_SimulateRequest)
}
func (x fastReflection_SimulateRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SimulateRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SimulateRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_SimulateRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SimulateRequest) Type() protoreflect.MessageType {
	return _fastReflection_SimulateRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SimulateRequest) New() protoreflect.Message {
	return new(fastReflection_SimulateRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SimulateRequest) Interface() protoreflect.ProtoMessage {
	return (*SimulateRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SimulateRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Tx != nil {
		value := protoreflect.ValueOfMessage(x.Tx.ProtoReflect())
		if !f(fd_SimulateRequest_tx, value) {
			return
		}
	}
	if len(x.TxBytes) != 0 {
		value := protoreflect.ValueOfBytes(x.TxBytes)
		if !f(fd_SimulateRequest_tx_bytes, value) {
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
func (x *fastReflection_SimulateRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SimulateRequest.tx":
		return x.Tx != nil
	case "cosmos.tx.v1beta1.SimulateRequest.tx_bytes":
		return len(x.TxBytes) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SimulateRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SimulateRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SimulateRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SimulateRequest.tx":
		x.Tx = nil
	case "cosmos.tx.v1beta1.SimulateRequest.tx_bytes":
		x.TxBytes = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SimulateRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SimulateRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SimulateRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.SimulateRequest.tx":
		value := x.Tx
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.SimulateRequest.tx_bytes":
		value := x.TxBytes
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SimulateRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SimulateRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SimulateRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SimulateRequest.tx":
		x.Tx = value.Message().Interface().(*Tx)
	case "cosmos.tx.v1beta1.SimulateRequest.tx_bytes":
		x.TxBytes = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SimulateRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SimulateRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SimulateRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SimulateRequest.tx":
		if x.Tx == nil {
			x.Tx = new(Tx)
		}
		return protoreflect.ValueOfMessage(x.Tx.ProtoReflect())
	case "cosmos.tx.v1beta1.SimulateRequest.tx_bytes":
		panic(fmt.Errorf("field tx_bytes of message cosmos.tx.v1beta1.SimulateRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SimulateRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SimulateRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SimulateRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SimulateRequest.tx":
		m := new(Tx)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.v1beta1.SimulateRequest.tx_bytes":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SimulateRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SimulateRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SimulateRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.SimulateRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SimulateRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SimulateRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SimulateRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SimulateRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SimulateRequest)
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
		if x.Tx != nil {
			l = options.Size(x.Tx)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TxBytes)
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
		x := input.Message.Interface().(*SimulateRequest)
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
		if len(x.TxBytes) > 0 {
			i -= len(x.TxBytes)
			copy(dAtA[i:], x.TxBytes)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TxBytes)))
			i--
			dAtA[i] = 0x12
		}
		if x.Tx != nil {
			encoded, err := options.Marshal(x.Tx)
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
		x := input.Message.Interface().(*SimulateRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SimulateRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SimulateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
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
				if x.Tx == nil {
					x.Tx = &Tx{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Tx); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TxBytes", wireType)
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
				x.TxBytes = append(x.TxBytes[:0], dAtA[iNdEx:postIndex]...)
				if x.TxBytes == nil {
					x.TxBytes = []byte{}
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
	md_SimulateResponse          protoreflect.MessageDescriptor
	fd_SimulateResponse_gas_info protoreflect.FieldDescriptor
	fd_SimulateResponse_result   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_service_proto_init()
	md_SimulateResponse = File_cosmos_tx_v1beta1_service_proto.Messages().ByName("SimulateResponse")
	fd_SimulateResponse_gas_info = md_SimulateResponse.Fields().ByName("gas_info")
	fd_SimulateResponse_result = md_SimulateResponse.Fields().ByName("result")
}

var _ protoreflect.Message = (*fastReflection_SimulateResponse)(nil)

type fastReflection_SimulateResponse SimulateResponse

func (x *SimulateResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SimulateResponse)(x)
}

func (x *SimulateResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SimulateResponse_messageType fastReflection_SimulateResponse_messageType
var _ protoreflect.MessageType = fastReflection_SimulateResponse_messageType{}

type fastReflection_SimulateResponse_messageType struct{}

func (x fastReflection_SimulateResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SimulateResponse)(nil)
}
func (x fastReflection_SimulateResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_SimulateResponse)
}
func (x fastReflection_SimulateResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SimulateResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SimulateResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_SimulateResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SimulateResponse) Type() protoreflect.MessageType {
	return _fastReflection_SimulateResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SimulateResponse) New() protoreflect.Message {
	return new(fastReflection_SimulateResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SimulateResponse) Interface() protoreflect.ProtoMessage {
	return (*SimulateResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SimulateResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.GasInfo != nil {
		value := protoreflect.ValueOfMessage(x.GasInfo.ProtoReflect())
		if !f(fd_SimulateResponse_gas_info, value) {
			return
		}
	}
	if x.Result != nil {
		value := protoreflect.ValueOfMessage(x.Result.ProtoReflect())
		if !f(fd_SimulateResponse_result, value) {
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
func (x *fastReflection_SimulateResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SimulateResponse.gas_info":
		return x.GasInfo != nil
	case "cosmos.tx.v1beta1.SimulateResponse.result":
		return x.Result != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SimulateResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SimulateResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SimulateResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SimulateResponse.gas_info":
		x.GasInfo = nil
	case "cosmos.tx.v1beta1.SimulateResponse.result":
		x.Result = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SimulateResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SimulateResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SimulateResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.SimulateResponse.gas_info":
		value := x.GasInfo
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.SimulateResponse.result":
		value := x.Result
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SimulateResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SimulateResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SimulateResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SimulateResponse.gas_info":
		x.GasInfo = value.Message().Interface().(*v1beta11.GasInfo)
	case "cosmos.tx.v1beta1.SimulateResponse.result":
		x.Result = value.Message().Interface().(*v1beta11.Result)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SimulateResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SimulateResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SimulateResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SimulateResponse.gas_info":
		if x.GasInfo == nil {
			x.GasInfo = new(v1beta11.GasInfo)
		}
		return protoreflect.ValueOfMessage(x.GasInfo.ProtoReflect())
	case "cosmos.tx.v1beta1.SimulateResponse.result":
		if x.Result == nil {
			x.Result = new(v1beta11.Result)
		}
		return protoreflect.ValueOfMessage(x.Result.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SimulateResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SimulateResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SimulateResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.SimulateResponse.gas_info":
		m := new(v1beta11.GasInfo)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.v1beta1.SimulateResponse.result":
		m := new(v1beta11.Result)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.SimulateResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.SimulateResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SimulateResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.SimulateResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SimulateResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SimulateResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SimulateResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SimulateResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SimulateResponse)
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
		if x.GasInfo != nil {
			l = options.Size(x.GasInfo)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Result != nil {
			l = options.Size(x.Result)
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
		x := input.Message.Interface().(*SimulateResponse)
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
		if x.Result != nil {
			encoded, err := options.Marshal(x.Result)
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
		if x.GasInfo != nil {
			encoded, err := options.Marshal(x.GasInfo)
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
		x := input.Message.Interface().(*SimulateResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SimulateResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SimulateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GasInfo", wireType)
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
				if x.GasInfo == nil {
					x.GasInfo = &v1beta11.GasInfo{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.GasInfo); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
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
				if x.Result == nil {
					x.Result = &v1beta11.Result{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Result); err != nil {
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
	md_GetTxRequest      protoreflect.MessageDescriptor
	fd_GetTxRequest_hash protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_service_proto_init()
	md_GetTxRequest = File_cosmos_tx_v1beta1_service_proto.Messages().ByName("GetTxRequest")
	fd_GetTxRequest_hash = md_GetTxRequest.Fields().ByName("hash")
}

var _ protoreflect.Message = (*fastReflection_GetTxRequest)(nil)

type fastReflection_GetTxRequest GetTxRequest

func (x *GetTxRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetTxRequest)(x)
}

func (x *GetTxRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetTxRequest_messageType fastReflection_GetTxRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetTxRequest_messageType{}

type fastReflection_GetTxRequest_messageType struct{}

func (x fastReflection_GetTxRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetTxRequest)(nil)
}
func (x fastReflection_GetTxRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetTxRequest)
}
func (x fastReflection_GetTxRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetTxRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetTxRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetTxRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetTxRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetTxRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetTxRequest) New() protoreflect.Message {
	return new(fastReflection_GetTxRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetTxRequest) Interface() protoreflect.ProtoMessage {
	return (*GetTxRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetTxRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Hash != "" {
		value := protoreflect.ValueOfString(x.Hash)
		if !f(fd_GetTxRequest_hash, value) {
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
func (x *fastReflection_GetTxRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxRequest.hash":
		return x.Hash != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetTxRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxRequest.hash":
		x.Hash = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetTxRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.GetTxRequest.hash":
		value := x.Hash
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetTxRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxRequest.hash":
		x.Hash = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetTxRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxRequest.hash":
		panic(fmt.Errorf("field hash of message cosmos.tx.v1beta1.GetTxRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetTxRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxRequest.hash":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxRequest"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetTxRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.GetTxRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetTxRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetTxRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetTxRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetTxRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetTxRequest)
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
		x := input.Message.Interface().(*GetTxRequest)
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
		x := input.Message.Interface().(*GetTxRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetTxRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetTxRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
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
				x.Hash = string(dAtA[iNdEx:postIndex])
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
	md_GetTxResponse             protoreflect.MessageDescriptor
	fd_GetTxResponse_tx          protoreflect.FieldDescriptor
	fd_GetTxResponse_tx_response protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_v1beta1_service_proto_init()
	md_GetTxResponse = File_cosmos_tx_v1beta1_service_proto.Messages().ByName("GetTxResponse")
	fd_GetTxResponse_tx = md_GetTxResponse.Fields().ByName("tx")
	fd_GetTxResponse_tx_response = md_GetTxResponse.Fields().ByName("tx_response")
}

var _ protoreflect.Message = (*fastReflection_GetTxResponse)(nil)

type fastReflection_GetTxResponse GetTxResponse

func (x *GetTxResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetTxResponse)(x)
}

func (x *GetTxResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetTxResponse_messageType fastReflection_GetTxResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetTxResponse_messageType{}

type fastReflection_GetTxResponse_messageType struct{}

func (x fastReflection_GetTxResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetTxResponse)(nil)
}
func (x fastReflection_GetTxResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetTxResponse)
}
func (x fastReflection_GetTxResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetTxResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetTxResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetTxResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetTxResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetTxResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetTxResponse) New() protoreflect.Message {
	return new(fastReflection_GetTxResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetTxResponse) Interface() protoreflect.ProtoMessage {
	return (*GetTxResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetTxResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Tx != nil {
		value := protoreflect.ValueOfMessage(x.Tx.ProtoReflect())
		if !f(fd_GetTxResponse_tx, value) {
			return
		}
	}
	if x.TxResponse != nil {
		value := protoreflect.ValueOfMessage(x.TxResponse.ProtoReflect())
		if !f(fd_GetTxResponse_tx_response, value) {
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
func (x *fastReflection_GetTxResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxResponse.tx":
		return x.Tx != nil
	case "cosmos.tx.v1beta1.GetTxResponse.tx_response":
		return x.TxResponse != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetTxResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxResponse.tx":
		x.Tx = nil
	case "cosmos.tx.v1beta1.GetTxResponse.tx_response":
		x.TxResponse = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetTxResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.v1beta1.GetTxResponse.tx":
		value := x.Tx
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.v1beta1.GetTxResponse.tx_response":
		value := x.TxResponse
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetTxResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxResponse.tx":
		x.Tx = value.Message().Interface().(*Tx)
	case "cosmos.tx.v1beta1.GetTxResponse.tx_response":
		x.TxResponse = value.Message().Interface().(*v1beta11.TxResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetTxResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxResponse.tx":
		if x.Tx == nil {
			x.Tx = new(Tx)
		}
		return protoreflect.ValueOfMessage(x.Tx.ProtoReflect())
	case "cosmos.tx.v1beta1.GetTxResponse.tx_response":
		if x.TxResponse == nil {
			x.TxResponse = new(v1beta11.TxResponse)
		}
		return protoreflect.ValueOfMessage(x.TxResponse.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetTxResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.v1beta1.GetTxResponse.tx":
		m := new(Tx)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.v1beta1.GetTxResponse.tx_response":
		m := new(v1beta11.TxResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.v1beta1.GetTxResponse"))
		}
		panic(fmt.Errorf("message cosmos.tx.v1beta1.GetTxResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetTxResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.v1beta1.GetTxResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetTxResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetTxResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetTxResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetTxResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetTxResponse)
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
		if x.Tx != nil {
			l = options.Size(x.Tx)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.TxResponse != nil {
			l = options.Size(x.TxResponse)
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
		x := input.Message.Interface().(*GetTxResponse)
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
		if x.TxResponse != nil {
			encoded, err := options.Marshal(x.TxResponse)
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
		if x.Tx != nil {
			encoded, err := options.Marshal(x.Tx)
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
		x := input.Message.Interface().(*GetTxResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetTxResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
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
				if x.Tx == nil {
					x.Tx = &Tx{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Tx); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TxResponse", wireType)
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
				if x.TxResponse == nil {
					x.TxResponse = &v1beta11.TxResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.TxResponse); err != nil {
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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// Simulate simulates executing a transaction for estimating gas usage.
	Simulate(ctx context.Context, in *SimulateRequest, opts ...grpc.CallOption) (*SimulateResponse, error)
	// GetTx fetches a tx by hash.
	GetTx(ctx context.Context, in *GetTxRequest, opts ...grpc.CallOption) (*GetTxResponse, error)
	// BroadcastTx broadcast transaction.
	BroadcastTx(ctx context.Context, in *BroadcastTxRequest, opts ...grpc.CallOption) (*BroadcastTxResponse, error)
	// GetTxsEvent fetches txs by event.
	GetTxsEvent(ctx context.Context, in *GetTxsEventRequest, opts ...grpc.CallOption) (*GetTxsEventResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Simulate(ctx context.Context, in *SimulateRequest, opts ...grpc.CallOption) (*SimulateResponse, error) {
	out := new(SimulateResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.Service/Simulate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetTx(ctx context.Context, in *GetTxRequest, opts ...grpc.CallOption) (*GetTxResponse, error) {
	out := new(GetTxResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.Service/GetTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BroadcastTx(ctx context.Context, in *BroadcastTxRequest, opts ...grpc.CallOption) (*BroadcastTxResponse, error) {
	out := new(BroadcastTxResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.Service/BroadcastTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetTxsEvent(ctx context.Context, in *GetTxsEventRequest, opts ...grpc.CallOption) (*GetTxsEventResponse, error) {
	out := new(GetTxsEventResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.Service/GetTxsEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// Simulate simulates executing a transaction for estimating gas usage.
	Simulate(context.Context, *SimulateRequest) (*SimulateResponse, error)
	// GetTx fetches a tx by hash.
	GetTx(context.Context, *GetTxRequest) (*GetTxResponse, error)
	// BroadcastTx broadcast transaction.
	BroadcastTx(context.Context, *BroadcastTxRequest) (*BroadcastTxResponse, error)
	// GetTxsEvent fetches txs by event.
	GetTxsEvent(context.Context, *GetTxsEventRequest) (*GetTxsEventResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Simulate(context.Context, *SimulateRequest) (*SimulateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Simulate not implemented")
}
func (UnimplementedServiceServer) GetTx(context.Context, *GetTxRequest) (*GetTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTx not implemented")
}
func (UnimplementedServiceServer) BroadcastTx(context.Context, *BroadcastTxRequest) (*BroadcastTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastTx not implemented")
}
func (UnimplementedServiceServer) GetTxsEvent(context.Context, *GetTxsEventRequest) (*GetTxsEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxsEvent not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Simulate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimulateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Simulate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.Service/Simulate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Simulate(ctx, req.(*SimulateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.Service/GetTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTx(ctx, req.(*GetTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_BroadcastTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).BroadcastTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.Service/BroadcastTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).BroadcastTx(ctx, req.(*BroadcastTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetTxsEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxsEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTxsEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.Service/GetTxsEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTxsEvent(ctx, req.(*GetTxsEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.tx.v1beta1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Simulate",
			Handler:    _Service_Simulate_Handler,
		},
		{
			MethodName: "GetTx",
			Handler:    _Service_GetTx_Handler,
		},
		{
			MethodName: "BroadcastTx",
			Handler:    _Service_BroadcastTx_Handler,
		},
		{
			MethodName: "GetTxsEvent",
			Handler:    _Service_GetTxsEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/tx/v1beta1/service.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/tx/v1beta1/service.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// OrderBy defines the sorting order
type OrderBy int32

const (
	// ORDER_BY_UNSPECIFIED specifies an unknown sorting order. OrderBy defaults to ASC in this case.
	OrderBy_ORDER_BY_UNSPECIFIED OrderBy = 0
	// ORDER_BY_ASC defines ascending order
	OrderBy_ORDER_BY_ASC OrderBy = 1
	// ORDER_BY_DESC defines descending order
	OrderBy_ORDER_BY_DESC OrderBy = 2
)

// Enum value maps for OrderBy.
var (
	OrderBy_name = map[int32]string{
		0: "ORDER_BY_UNSPECIFIED",
		1: "ORDER_BY_ASC",
		2: "ORDER_BY_DESC",
	}
	OrderBy_value = map[string]int32{
		"ORDER_BY_UNSPECIFIED": 0,
		"ORDER_BY_ASC":         1,
		"ORDER_BY_DESC":        2,
	}
)

func (x OrderBy) Enum() *OrderBy {
	p := new(OrderBy)
	*p = x
	return p
}

func (x OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_cosmos_tx_v1beta1_service_proto_enumTypes[0].Descriptor()
}

func (OrderBy) Type() protoreflect.EnumType {
	return &file_cosmos_tx_v1beta1_service_proto_enumTypes[0]
}

func (x OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderBy.Descriptor instead.
func (OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_service_proto_rawDescGZIP(), []int{0}
}

// BroadcastMode specifies the broadcast mode for the TxService.Broadcast RPC method.
type BroadcastMode int32

const (
	// zero-value for mode ordering
	BroadcastMode_BROADCAST_MODE_UNSPECIFIED BroadcastMode = 0
	// BROADCAST_MODE_BLOCK defines a tx broadcasting mode where the client waits for
	// the tx to be committed in a block.
	BroadcastMode_BROADCAST_MODE_BLOCK BroadcastMode = 1
	// BROADCAST_MODE_SYNC defines a tx broadcasting mode where the client waits for
	// a CheckTx execution response only.
	BroadcastMode_BROADCAST_MODE_SYNC BroadcastMode = 2
	// BROADCAST_MODE_ASYNC defines a tx broadcasting mode where the client returns
	// immediately.
	BroadcastMode_BROADCAST_MODE_ASYNC BroadcastMode = 3
)

// Enum value maps for BroadcastMode.
var (
	BroadcastMode_name = map[int32]string{
		0: "BROADCAST_MODE_UNSPECIFIED",
		1: "BROADCAST_MODE_BLOCK",
		2: "BROADCAST_MODE_SYNC",
		3: "BROADCAST_MODE_ASYNC",
	}
	BroadcastMode_value = map[string]int32{
		"BROADCAST_MODE_UNSPECIFIED": 0,
		"BROADCAST_MODE_BLOCK":       1,
		"BROADCAST_MODE_SYNC":        2,
		"BROADCAST_MODE_ASYNC":       3,
	}
)

func (x BroadcastMode) Enum() *BroadcastMode {
	p := new(BroadcastMode)
	*p = x
	return p
}

func (x BroadcastMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BroadcastMode) Descriptor() protoreflect.EnumDescriptor {
	return file_cosmos_tx_v1beta1_service_proto_enumTypes[1].Descriptor()
}

func (BroadcastMode) Type() protoreflect.EnumType {
	return &file_cosmos_tx_v1beta1_service_proto_enumTypes[1]
}

func (x BroadcastMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BroadcastMode.Descriptor instead.
func (BroadcastMode) EnumDescriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_service_proto_rawDescGZIP(), []int{1}
}

// GetTxsEventRequest is the request type for the Service.TxsByEvents
// RPC method.
type GetTxsEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// events is the list of transaction event type.
	Events []string `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	// pagination defines an pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	OrderBy    OrderBy              `protobuf:"varint,3,opt,name=order_by,json=orderBy,proto3,enum=cosmos.tx.v1beta1.OrderBy" json:"order_by,omitempty"`
}

func (x *GetTxsEventRequest) Reset() {
	*x = GetTxsEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxsEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxsEventRequest) ProtoMessage() {}

// Deprecated: Use GetTxsEventRequest.ProtoReflect.Descriptor instead.
func (*GetTxsEventRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetTxsEventRequest) GetEvents() []string {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *GetTxsEventRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *GetTxsEventRequest) GetOrderBy() OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return OrderBy_ORDER_BY_UNSPECIFIED
}

// GetTxsEventResponse is the response type for the Service.TxsByEvents
// RPC method.
type GetTxsEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// txs is the list of queried transactions.
	Txs []*Tx `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	// tx_responses is the list of queried TxResponses.
	TxResponses []*v1beta11.TxResponse `protobuf:"bytes,2,rep,name=tx_responses,json=txResponses,proto3" json:"tx_responses,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetTxsEventResponse) Reset() {
	*x = GetTxsEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxsEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxsEventResponse) ProtoMessage() {}

// Deprecated: Use GetTxsEventResponse.ProtoReflect.Descriptor instead.
func (*GetTxsEventResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTxsEventResponse) GetTxs() []*Tx {
	if x != nil {
		return x.Txs
	}
	return nil
}

func (x *GetTxsEventResponse) GetTxResponses() []*v1beta11.TxResponse {
	if x != nil {
		return x.TxResponses
	}
	return nil
}

func (x *GetTxsEventResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// BroadcastTxRequest is the request type for the Service.BroadcastTxRequest
// RPC method.
type BroadcastTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tx_bytes is the raw transaction.
	TxBytes []byte        `protobuf:"bytes,1,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
	Mode    BroadcastMode `protobuf:"varint,2,opt,name=mode,proto3,enum=cosmos.tx.v1beta1.BroadcastMode" json:"mode,omitempty"`
}

func (x *BroadcastTxRequest) Reset() {
	*x = BroadcastTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastTxRequest) ProtoMessage() {}

// Deprecated: Use BroadcastTxRequest.ProtoReflect.Descriptor instead.
func (*BroadcastTxRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_service_proto_rawDescGZIP(), []int{2}
}

func (x *BroadcastTxRequest) GetTxBytes() []byte {
	if x != nil {
		return x.TxBytes
	}
	return nil
}

func (x *BroadcastTxRequest) GetMode() BroadcastMode {
	if x != nil {
		return x.Mode
	}
	return BroadcastMode_BROADCAST_MODE_UNSPECIFIED
}

// BroadcastTxResponse is the response type for the
// Service.BroadcastTx method.
type BroadcastTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tx_response is the queried TxResponses.
	TxResponse *v1beta11.TxResponse `protobuf:"bytes,1,opt,name=tx_response,json=txResponse,proto3" json:"tx_response,omitempty"`
}

func (x *BroadcastTxResponse) Reset() {
	*x = BroadcastTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastTxResponse) ProtoMessage() {}

// Deprecated: Use BroadcastTxResponse.ProtoReflect.Descriptor instead.
func (*BroadcastTxResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_service_proto_rawDescGZIP(), []int{3}
}

func (x *BroadcastTxResponse) GetTxResponse() *v1beta11.TxResponse {
	if x != nil {
		return x.TxResponse
	}
	return nil
}

// SimulateRequest is the request type for the Service.Simulate
// RPC method.
type SimulateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tx is the transaction to simulate.
	// Deprecated. Send raw tx bytes instead.
	//
	// Deprecated: Do not use.
	Tx *Tx `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	// tx_bytes is the raw transaction.
	//
	// Since: cosmos-sdk 0.43
	TxBytes []byte `protobuf:"bytes,2,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
}

func (x *SimulateRequest) Reset() {
	*x = SimulateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulateRequest) ProtoMessage() {}

// Deprecated: Use SimulateRequest.ProtoReflect.Descriptor instead.
func (*SimulateRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_service_proto_rawDescGZIP(), []int{4}
}

// Deprecated: Do not use.
func (x *SimulateRequest) GetTx() *Tx {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *SimulateRequest) GetTxBytes() []byte {
	if x != nil {
		return x.TxBytes
	}
	return nil
}

// SimulateResponse is the response type for the
// Service.SimulateRPC method.
type SimulateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// gas_info is the information about gas used in the simulation.
	GasInfo *v1beta11.GasInfo `protobuf:"bytes,1,opt,name=gas_info,json=gasInfo,proto3" json:"gas_info,omitempty"`
	// result is the result of the simulation.
	Result *v1beta11.Result `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SimulateResponse) Reset() {
	*x = SimulateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulateResponse) ProtoMessage() {}

// Deprecated: Use SimulateResponse.ProtoReflect.Descriptor instead.
func (*SimulateResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_service_proto_rawDescGZIP(), []int{5}
}

func (x *SimulateResponse) GetGasInfo() *v1beta11.GasInfo {
	if x != nil {
		return x.GasInfo
	}
	return nil
}

func (x *SimulateResponse) GetResult() *v1beta11.Result {
	if x != nil {
		return x.Result
	}
	return nil
}

// GetTxRequest is the request type for the Service.GetTx
// RPC method.
type GetTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hash is the tx hash to query, encoded as a hex string.
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *GetTxRequest) Reset() {
	*x = GetTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxRequest) ProtoMessage() {}

// Deprecated: Use GetTxRequest.ProtoReflect.Descriptor instead.
func (*GetTxRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetTxRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// GetTxResponse is the response type for the Service.GetTx method.
type GetTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tx is the queried transaction.
	Tx *Tx `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	// tx_response is the queried TxResponses.
	TxResponse *v1beta11.TxResponse `protobuf:"bytes,2,opt,name=tx_response,json=txResponse,proto3" json:"tx_response,omitempty"`
}

func (x *GetTxResponse) Reset() {
	*x = GetTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_v1beta1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxResponse) ProtoMessage() {}

// Deprecated: Use GetTxResponse.ProtoReflect.Descriptor instead.
func (*GetTxResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_v1beta1_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetTxResponse) GetTx() *Tx {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *GetTxResponse) GetTxResponse() *v1beta11.TxResponse {
	if x != nil {
		return x.TxResponse
	}
	return nil
}

var File_cosmos_tx_v1beta1_service_proto protoreflect.FileDescriptor

var file_cosmos_tx_v1beta1_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x78, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x61, 0x62, 0x63, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x62, 0x63,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x74, 0x78, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x22, 0xd0, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x74,
	0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x78, 0x52,
	0x03, 0x74, 0x78, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x74, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0b, 0x74, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x47, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x65, 0x0a, 0x12, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x74, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74,
	0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x5c, 0x0a,
	0x13, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x74, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x0a, 0x74, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x0a, 0x0f, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54,
	0x78, 0x42, 0x02, 0x18, 0x01, 0x52, 0x02, 0x74, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x78, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x67, 0x61, 0x73,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x61, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07,
	0x67, 0x61, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x22, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x7d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x78, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x02, 0x74, 0x78, 0x12, 0x45, 0x0a,
	0x0b, 0x74, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x78,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x74, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x48, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12,
	0x18, 0x0a, 0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4f,
	0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x02, 0x2a, 0x7c,
	0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x1e, 0x0a, 0x1a, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x18, 0x0a, 0x14, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x52, 0x4f,
	0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x59, 0x4e, 0x43,
	0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x5f,
	0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x03, 0x32, 0xf8, 0x03, 0x0a,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x08, 0x53, 0x69, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74,
	0x78, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x54, 0x78, 0x12, 0x1f,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x74, 0x78, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78,
	0x73, 0x2f, 0x7b, 0x68, 0x61, 0x73, 0x68, 0x7d, 0x12, 0x7f, 0x0a, 0x0b, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x78, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x78, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x78, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x74, 0x78, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7c, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x54, 0x78, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x78, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12,
	0x16, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x78, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x73, 0x42, 0xcd, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x78, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x74, 0x78, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x54, 0x58, 0xaa, 0x02, 0x11, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x54,
	0x78, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x11, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x5c, 0x54, 0x78, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1d,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x54, 0x78, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x54, 0x78, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xc0, 0xe3, 0x1e, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_tx_v1beta1_service_proto_rawDescOnce sync.Once
	file_cosmos_tx_v1beta1_service_proto_rawDescData = file_cosmos_tx_v1beta1_service_proto_rawDesc
)

func file_cosmos_tx_v1beta1_service_proto_rawDescGZIP() []byte {
	file_cosmos_tx_v1beta1_service_proto_rawDescOnce.Do(func() {
		file_cosmos_tx_v1beta1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_tx_v1beta1_service_proto_rawDescData)
	})
	return file_cosmos_tx_v1beta1_service_proto_rawDescData
}

var file_cosmos_tx_v1beta1_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cosmos_tx_v1beta1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cosmos_tx_v1beta1_service_proto_goTypes = []interface{}{
	(OrderBy)(0),                 // 0: cosmos.tx.v1beta1.OrderBy
	(BroadcastMode)(0),           // 1: cosmos.tx.v1beta1.BroadcastMode
	(*GetTxsEventRequest)(nil),   // 2: cosmos.tx.v1beta1.GetTxsEventRequest
	(*GetTxsEventResponse)(nil),  // 3: cosmos.tx.v1beta1.GetTxsEventResponse
	(*BroadcastTxRequest)(nil),   // 4: cosmos.tx.v1beta1.BroadcastTxRequest
	(*BroadcastTxResponse)(nil),  // 5: cosmos.tx.v1beta1.BroadcastTxResponse
	(*SimulateRequest)(nil),      // 6: cosmos.tx.v1beta1.SimulateRequest
	(*SimulateResponse)(nil),     // 7: cosmos.tx.v1beta1.SimulateResponse
	(*GetTxRequest)(nil),         // 8: cosmos.tx.v1beta1.GetTxRequest
	(*GetTxResponse)(nil),        // 9: cosmos.tx.v1beta1.GetTxResponse
	(*v1beta1.PageRequest)(nil),  // 10: cosmos.base.query.v1beta1.PageRequest
	(*Tx)(nil),                   // 11: cosmos.tx.v1beta1.Tx
	(*v1beta11.TxResponse)(nil),  // 12: cosmos.base.abci.v1beta1.TxResponse
	(*v1beta1.PageResponse)(nil), // 13: cosmos.base.query.v1beta1.PageResponse
	(*v1beta11.GasInfo)(nil),     // 14: cosmos.base.abci.v1beta1.GasInfo
	(*v1beta11.Result)(nil),      // 15: cosmos.base.abci.v1beta1.Result
}
var file_cosmos_tx_v1beta1_service_proto_depIdxs = []int32{
	10, // 0: cosmos.tx.v1beta1.GetTxsEventRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	0,  // 1: cosmos.tx.v1beta1.GetTxsEventRequest.order_by:type_name -> cosmos.tx.v1beta1.OrderBy
	11, // 2: cosmos.tx.v1beta1.GetTxsEventResponse.txs:type_name -> cosmos.tx.v1beta1.Tx
	12, // 3: cosmos.tx.v1beta1.GetTxsEventResponse.tx_responses:type_name -> cosmos.base.abci.v1beta1.TxResponse
	13, // 4: cosmos.tx.v1beta1.GetTxsEventResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	1,  // 5: cosmos.tx.v1beta1.BroadcastTxRequest.mode:type_name -> cosmos.tx.v1beta1.BroadcastMode
	12, // 6: cosmos.tx.v1beta1.BroadcastTxResponse.tx_response:type_name -> cosmos.base.abci.v1beta1.TxResponse
	11, // 7: cosmos.tx.v1beta1.SimulateRequest.tx:type_name -> cosmos.tx.v1beta1.Tx
	14, // 8: cosmos.tx.v1beta1.SimulateResponse.gas_info:type_name -> cosmos.base.abci.v1beta1.GasInfo
	15, // 9: cosmos.tx.v1beta1.SimulateResponse.result:type_name -> cosmos.base.abci.v1beta1.Result
	11, // 10: cosmos.tx.v1beta1.GetTxResponse.tx:type_name -> cosmos.tx.v1beta1.Tx
	12, // 11: cosmos.tx.v1beta1.GetTxResponse.tx_response:type_name -> cosmos.base.abci.v1beta1.TxResponse
	6,  // 12: cosmos.tx.v1beta1.Service.Simulate:input_type -> cosmos.tx.v1beta1.SimulateRequest
	8,  // 13: cosmos.tx.v1beta1.Service.GetTx:input_type -> cosmos.tx.v1beta1.GetTxRequest
	4,  // 14: cosmos.tx.v1beta1.Service.BroadcastTx:input_type -> cosmos.tx.v1beta1.BroadcastTxRequest
	2,  // 15: cosmos.tx.v1beta1.Service.GetTxsEvent:input_type -> cosmos.tx.v1beta1.GetTxsEventRequest
	7,  // 16: cosmos.tx.v1beta1.Service.Simulate:output_type -> cosmos.tx.v1beta1.SimulateResponse
	9,  // 17: cosmos.tx.v1beta1.Service.GetTx:output_type -> cosmos.tx.v1beta1.GetTxResponse
	5,  // 18: cosmos.tx.v1beta1.Service.BroadcastTx:output_type -> cosmos.tx.v1beta1.BroadcastTxResponse
	3,  // 19: cosmos.tx.v1beta1.Service.GetTxsEvent:output_type -> cosmos.tx.v1beta1.GetTxsEventResponse
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_cosmos_tx_v1beta1_service_proto_init() }
func file_cosmos_tx_v1beta1_service_proto_init() {
	if File_cosmos_tx_v1beta1_service_proto != nil {
		return
	}
	file_cosmos_tx_v1beta1_tx_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_tx_v1beta1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxsEventRequest); i {
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
		file_cosmos_tx_v1beta1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxsEventResponse); i {
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
		file_cosmos_tx_v1beta1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastTxRequest); i {
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
		file_cosmos_tx_v1beta1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastTxResponse); i {
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
		file_cosmos_tx_v1beta1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulateRequest); i {
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
		file_cosmos_tx_v1beta1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulateResponse); i {
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
		file_cosmos_tx_v1beta1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxRequest); i {
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
		file_cosmos_tx_v1beta1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxResponse); i {
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
			RawDescriptor: file_cosmos_tx_v1beta1_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_tx_v1beta1_service_proto_goTypes,
		DependencyIndexes: file_cosmos_tx_v1beta1_service_proto_depIdxs,
		EnumInfos:         file_cosmos_tx_v1beta1_service_proto_enumTypes,
		MessageInfos:      file_cosmos_tx_v1beta1_service_proto_msgTypes,
	}.Build()
	File_cosmos_tx_v1beta1_service_proto = out.File
	file_cosmos_tx_v1beta1_service_proto_rawDesc = nil
	file_cosmos_tx_v1beta1_service_proto_goTypes = nil
	file_cosmos_tx_v1beta1_service_proto_depIdxs = nil
}
