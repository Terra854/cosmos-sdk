package vestingv1beta1

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
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

var _ protoreflect.List = (*_MsgCreateVestingAccount_3_list)(nil)

type _MsgCreateVestingAccount_3_list struct {
	list *[]*v1beta1.Coin
}

func (x *_MsgCreateVestingAccount_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreateVestingAccount_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgCreateVestingAccount_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreateVestingAccount_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreateVestingAccount_3_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateVestingAccount_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreateVestingAccount_3_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateVestingAccount_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCreateVestingAccount              protoreflect.MessageDescriptor
	fd_MsgCreateVestingAccount_from_address protoreflect.FieldDescriptor
	fd_MsgCreateVestingAccount_to_address   protoreflect.FieldDescriptor
	fd_MsgCreateVestingAccount_amount       protoreflect.FieldDescriptor
	fd_MsgCreateVestingAccount_end_time     protoreflect.FieldDescriptor
	fd_MsgCreateVestingAccount_delayed      protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_vesting_v1beta1_tx_proto_init()
	md_MsgCreateVestingAccount = File_cosmos_vesting_v1beta1_tx_proto.Messages().ByName("MsgCreateVestingAccount")
	fd_MsgCreateVestingAccount_from_address = md_MsgCreateVestingAccount.Fields().ByName("from_address")
	fd_MsgCreateVestingAccount_to_address = md_MsgCreateVestingAccount.Fields().ByName("to_address")
	fd_MsgCreateVestingAccount_amount = md_MsgCreateVestingAccount.Fields().ByName("amount")
	fd_MsgCreateVestingAccount_end_time = md_MsgCreateVestingAccount.Fields().ByName("end_time")
	fd_MsgCreateVestingAccount_delayed = md_MsgCreateVestingAccount.Fields().ByName("delayed")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateVestingAccount)(nil)

type fastReflection_MsgCreateVestingAccount MsgCreateVestingAccount

func (x *MsgCreateVestingAccount) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateVestingAccount)(x)
}

func (x *MsgCreateVestingAccount) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateVestingAccount_messageType fastReflection_MsgCreateVestingAccount_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateVestingAccount_messageType{}

type fastReflection_MsgCreateVestingAccount_messageType struct{}

func (x fastReflection_MsgCreateVestingAccount_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateVestingAccount)(nil)
}
func (x fastReflection_MsgCreateVestingAccount_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateVestingAccount)
}
func (x fastReflection_MsgCreateVestingAccount_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateVestingAccount
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateVestingAccount) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateVestingAccount
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateVestingAccount) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateVestingAccount_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateVestingAccount) New() protoreflect.Message {
	return new(fastReflection_MsgCreateVestingAccount)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateVestingAccount) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateVestingAccount)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateVestingAccount) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.FromAddress != "" {
		value := protoreflect.ValueOfString(x.FromAddress)
		if !f(fd_MsgCreateVestingAccount_from_address, value) {
			return
		}
	}
	if x.ToAddress != "" {
		value := protoreflect.ValueOfString(x.ToAddress)
		if !f(fd_MsgCreateVestingAccount_to_address, value) {
			return
		}
	}
	if len(x.Amount) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreateVestingAccount_3_list{list: &x.Amount})
		if !f(fd_MsgCreateVestingAccount_amount, value) {
			return
		}
	}
	if x.EndTime != int64(0) {
		value := protoreflect.ValueOfInt64(x.EndTime)
		if !f(fd_MsgCreateVestingAccount_end_time, value) {
			return
		}
	}
	if x.Delayed != false {
		value := protoreflect.ValueOfBool(x.Delayed)
		if !f(fd_MsgCreateVestingAccount_delayed, value) {
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
func (x *fastReflection_MsgCreateVestingAccount) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.from_address":
		return x.FromAddress != ""
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.to_address":
		return x.ToAddress != ""
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.amount":
		return len(x.Amount) != 0
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.end_time":
		return x.EndTime != int64(0)
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.delayed":
		return x.Delayed != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreateVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreateVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateVestingAccount) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.from_address":
		x.FromAddress = ""
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.to_address":
		x.ToAddress = ""
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.amount":
		x.Amount = nil
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.end_time":
		x.EndTime = int64(0)
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.delayed":
		x.Delayed = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreateVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreateVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateVestingAccount) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.from_address":
		value := x.FromAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.to_address":
		value := x.ToAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.amount":
		if len(x.Amount) == 0 {
			return protoreflect.ValueOfList(&_MsgCreateVestingAccount_3_list{})
		}
		listValue := &_MsgCreateVestingAccount_3_list{list: &x.Amount}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.end_time":
		value := x.EndTime
		return protoreflect.ValueOfInt64(value)
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.delayed":
		value := x.Delayed
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreateVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreateVestingAccount does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateVestingAccount) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.from_address":
		x.FromAddress = value.Interface().(string)
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.to_address":
		x.ToAddress = value.Interface().(string)
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.amount":
		lv := value.List()
		clv := lv.(*_MsgCreateVestingAccount_3_list)
		x.Amount = *clv.list
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.end_time":
		x.EndTime = value.Int()
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.delayed":
		x.Delayed = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreateVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreateVestingAccount does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateVestingAccount) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.amount":
		if x.Amount == nil {
			x.Amount = []*v1beta1.Coin{}
		}
		value := &_MsgCreateVestingAccount_3_list{list: &x.Amount}
		return protoreflect.ValueOfList(value)
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.from_address":
		panic(fmt.Errorf("field from_address of message cosmos.vesting.v1beta1.MsgCreateVestingAccount is not mutable"))
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.to_address":
		panic(fmt.Errorf("field to_address of message cosmos.vesting.v1beta1.MsgCreateVestingAccount is not mutable"))
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.end_time":
		panic(fmt.Errorf("field end_time of message cosmos.vesting.v1beta1.MsgCreateVestingAccount is not mutable"))
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.delayed":
		panic(fmt.Errorf("field delayed of message cosmos.vesting.v1beta1.MsgCreateVestingAccount is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreateVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreateVestingAccount does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateVestingAccount) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.from_address":
		return protoreflect.ValueOfString("")
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.to_address":
		return protoreflect.ValueOfString("")
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.amount":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_MsgCreateVestingAccount_3_list{list: &list})
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.end_time":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.vesting.v1beta1.MsgCreateVestingAccount.delayed":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreateVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreateVestingAccount does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateVestingAccount) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.vesting.v1beta1.MsgCreateVestingAccount", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateVestingAccount) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateVestingAccount) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateVestingAccount) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateVestingAccount) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateVestingAccount)
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
		l = len(x.FromAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ToAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Amount) > 0 {
			for _, e := range x.Amount {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.EndTime != 0 {
			n += 1 + runtime.Sov(uint64(x.EndTime))
		}
		if x.Delayed {
			n += 2
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
		x := input.Message.Interface().(*MsgCreateVestingAccount)
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
		if x.Delayed {
			i--
			if x.Delayed {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x28
		}
		if x.EndTime != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.EndTime))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Amount) > 0 {
			for iNdEx := len(x.Amount) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Amount[iNdEx])
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
		}
		if len(x.ToAddress) > 0 {
			i -= len(x.ToAddress)
			copy(dAtA[i:], x.ToAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ToAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.FromAddress) > 0 {
			i -= len(x.FromAddress)
			copy(dAtA[i:], x.FromAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.FromAddress)))
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
		x := input.Message.Interface().(*MsgCreateVestingAccount)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateVestingAccount: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
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
				x.FromAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
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
				x.ToAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
				x.Amount = append(x.Amount, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Amount[len(x.Amount)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
				}
				x.EndTime = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.EndTime |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Delayed", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Delayed = bool(v != 0)
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
	md_MsgCreateVestingAccountResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_vesting_v1beta1_tx_proto_init()
	md_MsgCreateVestingAccountResponse = File_cosmos_vesting_v1beta1_tx_proto.Messages().ByName("MsgCreateVestingAccountResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateVestingAccountResponse)(nil)

type fastReflection_MsgCreateVestingAccountResponse MsgCreateVestingAccountResponse

func (x *MsgCreateVestingAccountResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateVestingAccountResponse)(x)
}

func (x *MsgCreateVestingAccountResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateVestingAccountResponse_messageType fastReflection_MsgCreateVestingAccountResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateVestingAccountResponse_messageType{}

type fastReflection_MsgCreateVestingAccountResponse_messageType struct{}

func (x fastReflection_MsgCreateVestingAccountResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateVestingAccountResponse)(nil)
}
func (x fastReflection_MsgCreateVestingAccountResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateVestingAccountResponse)
}
func (x fastReflection_MsgCreateVestingAccountResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateVestingAccountResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateVestingAccountResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateVestingAccountResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateVestingAccountResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateVestingAccountResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateVestingAccountResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateVestingAccountResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateVestingAccountResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateVestingAccountResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateVestingAccountResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgCreateVestingAccountResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateVestingAccountResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateVestingAccountResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateVestingAccountResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateVestingAccountResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateVestingAccountResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateVestingAccountResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateVestingAccountResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateVestingAccountResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateVestingAccountResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateVestingAccountResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateVestingAccountResponse)
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
		x := input.Message.Interface().(*MsgCreateVestingAccountResponse)
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
		x := input.Message.Interface().(*MsgCreateVestingAccountResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateVestingAccountResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateVestingAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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

var _ protoreflect.List = (*_MsgCreatePeriodicVestingAccount_4_list)(nil)

type _MsgCreatePeriodicVestingAccount_4_list struct {
	list *[]*Period
}

func (x *_MsgCreatePeriodicVestingAccount_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreatePeriodicVestingAccount_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgCreatePeriodicVestingAccount_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Period)
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreatePeriodicVestingAccount_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Period)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreatePeriodicVestingAccount_4_list) AppendMutable() protoreflect.Value {
	v := new(Period)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreatePeriodicVestingAccount_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreatePeriodicVestingAccount_4_list) NewElement() protoreflect.Value {
	v := new(Period)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreatePeriodicVestingAccount_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCreatePeriodicVestingAccount                 protoreflect.MessageDescriptor
	fd_MsgCreatePeriodicVestingAccount_from_address    protoreflect.FieldDescriptor
	fd_MsgCreatePeriodicVestingAccount_to_address      protoreflect.FieldDescriptor
	fd_MsgCreatePeriodicVestingAccount_start_time      protoreflect.FieldDescriptor
	fd_MsgCreatePeriodicVestingAccount_vesting_periods protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_vesting_v1beta1_tx_proto_init()
	md_MsgCreatePeriodicVestingAccount = File_cosmos_vesting_v1beta1_tx_proto.Messages().ByName("MsgCreatePeriodicVestingAccount")
	fd_MsgCreatePeriodicVestingAccount_from_address = md_MsgCreatePeriodicVestingAccount.Fields().ByName("from_address")
	fd_MsgCreatePeriodicVestingAccount_to_address = md_MsgCreatePeriodicVestingAccount.Fields().ByName("to_address")
	fd_MsgCreatePeriodicVestingAccount_start_time = md_MsgCreatePeriodicVestingAccount.Fields().ByName("start_time")
	fd_MsgCreatePeriodicVestingAccount_vesting_periods = md_MsgCreatePeriodicVestingAccount.Fields().ByName("vesting_periods")
}

var _ protoreflect.Message = (*fastReflection_MsgCreatePeriodicVestingAccount)(nil)

type fastReflection_MsgCreatePeriodicVestingAccount MsgCreatePeriodicVestingAccount

func (x *MsgCreatePeriodicVestingAccount) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreatePeriodicVestingAccount)(x)
}

func (x *MsgCreatePeriodicVestingAccount) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreatePeriodicVestingAccount_messageType fastReflection_MsgCreatePeriodicVestingAccount_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreatePeriodicVestingAccount_messageType{}

type fastReflection_MsgCreatePeriodicVestingAccount_messageType struct{}

func (x fastReflection_MsgCreatePeriodicVestingAccount_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreatePeriodicVestingAccount)(nil)
}
func (x fastReflection_MsgCreatePeriodicVestingAccount_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreatePeriodicVestingAccount)
}
func (x fastReflection_MsgCreatePeriodicVestingAccount_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreatePeriodicVestingAccount
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreatePeriodicVestingAccount) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreatePeriodicVestingAccount
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreatePeriodicVestingAccount) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreatePeriodicVestingAccount_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreatePeriodicVestingAccount) New() protoreflect.Message {
	return new(fastReflection_MsgCreatePeriodicVestingAccount)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreatePeriodicVestingAccount) Interface() protoreflect.ProtoMessage {
	return (*MsgCreatePeriodicVestingAccount)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreatePeriodicVestingAccount) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.FromAddress != "" {
		value := protoreflect.ValueOfString(x.FromAddress)
		if !f(fd_MsgCreatePeriodicVestingAccount_from_address, value) {
			return
		}
	}
	if x.ToAddress != "" {
		value := protoreflect.ValueOfString(x.ToAddress)
		if !f(fd_MsgCreatePeriodicVestingAccount_to_address, value) {
			return
		}
	}
	if x.StartTime != int64(0) {
		value := protoreflect.ValueOfInt64(x.StartTime)
		if !f(fd_MsgCreatePeriodicVestingAccount_start_time, value) {
			return
		}
	}
	if len(x.VestingPeriods) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreatePeriodicVestingAccount_4_list{list: &x.VestingPeriods})
		if !f(fd_MsgCreatePeriodicVestingAccount_vesting_periods, value) {
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
func (x *fastReflection_MsgCreatePeriodicVestingAccount) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.from_address":
		return x.FromAddress != ""
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.to_address":
		return x.ToAddress != ""
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.start_time":
		return x.StartTime != int64(0)
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.vesting_periods":
		return len(x.VestingPeriods) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreatePeriodicVestingAccount) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.from_address":
		x.FromAddress = ""
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.to_address":
		x.ToAddress = ""
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.start_time":
		x.StartTime = int64(0)
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.vesting_periods":
		x.VestingPeriods = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreatePeriodicVestingAccount) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.from_address":
		value := x.FromAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.to_address":
		value := x.ToAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.start_time":
		value := x.StartTime
		return protoreflect.ValueOfInt64(value)
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.vesting_periods":
		if len(x.VestingPeriods) == 0 {
			return protoreflect.ValueOfList(&_MsgCreatePeriodicVestingAccount_4_list{})
		}
		listValue := &_MsgCreatePeriodicVestingAccount_4_list{list: &x.VestingPeriods}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreatePeriodicVestingAccount) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.from_address":
		x.FromAddress = value.Interface().(string)
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.to_address":
		x.ToAddress = value.Interface().(string)
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.start_time":
		x.StartTime = value.Int()
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.vesting_periods":
		lv := value.List()
		clv := lv.(*_MsgCreatePeriodicVestingAccount_4_list)
		x.VestingPeriods = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreatePeriodicVestingAccount) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.vesting_periods":
		if x.VestingPeriods == nil {
			x.VestingPeriods = []*Period{}
		}
		value := &_MsgCreatePeriodicVestingAccount_4_list{list: &x.VestingPeriods}
		return protoreflect.ValueOfList(value)
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.from_address":
		panic(fmt.Errorf("field from_address of message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount is not mutable"))
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.to_address":
		panic(fmt.Errorf("field to_address of message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount is not mutable"))
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.start_time":
		panic(fmt.Errorf("field start_time of message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreatePeriodicVestingAccount) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.from_address":
		return protoreflect.ValueOfString("")
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.to_address":
		return protoreflect.ValueOfString("")
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.start_time":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.vesting_periods":
		list := []*Period{}
		return protoreflect.ValueOfList(&_MsgCreatePeriodicVestingAccount_4_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreatePeriodicVestingAccount) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreatePeriodicVestingAccount) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreatePeriodicVestingAccount) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreatePeriodicVestingAccount) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreatePeriodicVestingAccount) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreatePeriodicVestingAccount)
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
		l = len(x.FromAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ToAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.StartTime != 0 {
			n += 1 + runtime.Sov(uint64(x.StartTime))
		}
		if len(x.VestingPeriods) > 0 {
			for _, e := range x.VestingPeriods {
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
		x := input.Message.Interface().(*MsgCreatePeriodicVestingAccount)
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
		if len(x.VestingPeriods) > 0 {
			for iNdEx := len(x.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.VestingPeriods[iNdEx])
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
		}
		if x.StartTime != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.StartTime))
			i--
			dAtA[i] = 0x18
		}
		if len(x.ToAddress) > 0 {
			i -= len(x.ToAddress)
			copy(dAtA[i:], x.ToAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ToAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.FromAddress) > 0 {
			i -= len(x.FromAddress)
			copy(dAtA[i:], x.FromAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.FromAddress)))
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
		x := input.Message.Interface().(*MsgCreatePeriodicVestingAccount)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreatePeriodicVestingAccount: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreatePeriodicVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
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
				x.FromAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
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
				x.ToAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
				}
				x.StartTime = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.StartTime |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
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
				x.VestingPeriods = append(x.VestingPeriods, &Period{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.VestingPeriods[len(x.VestingPeriods)-1]); err != nil {
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
	md_MsgCreatePeriodicVestingAccountResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_vesting_v1beta1_tx_proto_init()
	md_MsgCreatePeriodicVestingAccountResponse = File_cosmos_vesting_v1beta1_tx_proto.Messages().ByName("MsgCreatePeriodicVestingAccountResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgCreatePeriodicVestingAccountResponse)(nil)

type fastReflection_MsgCreatePeriodicVestingAccountResponse MsgCreatePeriodicVestingAccountResponse

func (x *MsgCreatePeriodicVestingAccountResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreatePeriodicVestingAccountResponse)(x)
}

func (x *MsgCreatePeriodicVestingAccountResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreatePeriodicVestingAccountResponse_messageType fastReflection_MsgCreatePeriodicVestingAccountResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreatePeriodicVestingAccountResponse_messageType{}

type fastReflection_MsgCreatePeriodicVestingAccountResponse_messageType struct{}

func (x fastReflection_MsgCreatePeriodicVestingAccountResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreatePeriodicVestingAccountResponse)(nil)
}
func (x fastReflection_MsgCreatePeriodicVestingAccountResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreatePeriodicVestingAccountResponse)
}
func (x fastReflection_MsgCreatePeriodicVestingAccountResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreatePeriodicVestingAccountResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreatePeriodicVestingAccountResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreatePeriodicVestingAccountResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreatePeriodicVestingAccountResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreatePeriodicVestingAccountResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreatePeriodicVestingAccountResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreatePeriodicVestingAccountResponse)
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
		x := input.Message.Interface().(*MsgCreatePeriodicVestingAccountResponse)
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
		x := input.Message.Interface().(*MsgCreatePeriodicVestingAccountResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreatePeriodicVestingAccountResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreatePeriodicVestingAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
	// CreateVestingAccount defines a method that enables creating a vesting
	// account.
	CreateVestingAccount(ctx context.Context, in *MsgCreateVestingAccount, opts ...grpc.CallOption) (*MsgCreateVestingAccountResponse, error)
	// CreatePeriodicVestingAccount defines a method that enables creating a
	// periodic vesting account.
	CreatePeriodicVestingAccount(ctx context.Context, in *MsgCreatePeriodicVestingAccount, opts ...grpc.CallOption) (*MsgCreatePeriodicVestingAccountResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateVestingAccount(ctx context.Context, in *MsgCreateVestingAccount, opts ...grpc.CallOption) (*MsgCreateVestingAccountResponse, error) {
	out := new(MsgCreateVestingAccountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.vesting.v1beta1.Msg/CreateVestingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePeriodicVestingAccount(ctx context.Context, in *MsgCreatePeriodicVestingAccount, opts ...grpc.CallOption) (*MsgCreatePeriodicVestingAccountResponse, error) {
	out := new(MsgCreatePeriodicVestingAccountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.vesting.v1beta1.Msg/CreatePeriodicVestingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateVestingAccount defines a method that enables creating a vesting
	// account.
	CreateVestingAccount(context.Context, *MsgCreateVestingAccount) (*MsgCreateVestingAccountResponse, error)
	// CreatePeriodicVestingAccount defines a method that enables creating a
	// periodic vesting account.
	CreatePeriodicVestingAccount(context.Context, *MsgCreatePeriodicVestingAccount) (*MsgCreatePeriodicVestingAccountResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateVestingAccount(context.Context, *MsgCreateVestingAccount) (*MsgCreateVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVestingAccount not implemented")
}
func (UnimplementedMsgServer) CreatePeriodicVestingAccount(context.Context, *MsgCreatePeriodicVestingAccount) (*MsgCreatePeriodicVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePeriodicVestingAccount not implemented")
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

func _Msg_CreateVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.vesting.v1beta1.Msg/CreateVestingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateVestingAccount(ctx, req.(*MsgCreateVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePeriodicVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePeriodicVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePeriodicVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.vesting.v1beta1.Msg/CreatePeriodicVestingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePeriodicVestingAccount(ctx, req.(*MsgCreatePeriodicVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.vesting.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVestingAccount",
			Handler:    _Msg_CreateVestingAccount_Handler,
		},
		{
			MethodName: "CreatePeriodicVestingAccount",
			Handler:    _Msg_CreatePeriodicVestingAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/vesting/v1beta1/tx.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/vesting/v1beta1/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MsgCreateVestingAccount defines a message that enables creating a vesting
// account.
type MsgCreateVestingAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress string          `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress   string          `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount      []*v1beta1.Coin `protobuf:"bytes,3,rep,name=amount,proto3" json:"amount,omitempty"`
	EndTime     int64           `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Delayed     bool            `protobuf:"varint,5,opt,name=delayed,proto3" json:"delayed,omitempty"`
}

func (x *MsgCreateVestingAccount) Reset() {
	*x = MsgCreateVestingAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateVestingAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateVestingAccount) ProtoMessage() {}

// Deprecated: Use MsgCreateVestingAccount.ProtoReflect.Descriptor instead.
func (*MsgCreateVestingAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgCreateVestingAccount) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *MsgCreateVestingAccount) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *MsgCreateVestingAccount) GetAmount() []*v1beta1.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *MsgCreateVestingAccount) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *MsgCreateVestingAccount) GetDelayed() bool {
	if x != nil {
		return x.Delayed
	}
	return false
}

// MsgCreateVestingAccountResponse defines the Msg/CreateVestingAccount response type.
type MsgCreateVestingAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgCreateVestingAccountResponse) Reset() {
	*x = MsgCreateVestingAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateVestingAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateVestingAccountResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateVestingAccountResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateVestingAccountResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_tx_proto_rawDescGZIP(), []int{1}
}

// MsgCreateVestingAccount defines a message that enables creating a vesting
// account.
type MsgCreatePeriodicVestingAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress    string    `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress      string    `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	StartTime      int64     `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	VestingPeriods []*Period `protobuf:"bytes,4,rep,name=vesting_periods,json=vestingPeriods,proto3" json:"vesting_periods,omitempty"`
}

func (x *MsgCreatePeriodicVestingAccount) Reset() {
	*x = MsgCreatePeriodicVestingAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreatePeriodicVestingAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreatePeriodicVestingAccount) ProtoMessage() {}

// Deprecated: Use MsgCreatePeriodicVestingAccount.ProtoReflect.Descriptor instead.
func (*MsgCreatePeriodicVestingAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgCreatePeriodicVestingAccount) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *MsgCreatePeriodicVestingAccount) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *MsgCreatePeriodicVestingAccount) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *MsgCreatePeriodicVestingAccount) GetVestingPeriods() []*Period {
	if x != nil {
		return x.VestingPeriods
	}
	return nil
}

// MsgCreateVestingAccountResponse defines the Msg/CreatePeriodicVestingAccount
// response type.
type MsgCreatePeriodicVestingAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgCreatePeriodicVestingAccountResponse) Reset() {
	*x = MsgCreatePeriodicVestingAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreatePeriodicVestingAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreatePeriodicVestingAccountResponse) ProtoMessage() {}

// Deprecated: Use MsgCreatePeriodicVestingAccountResponse.ProtoReflect.Descriptor instead.
func (*MsgCreatePeriodicVestingAccountResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_tx_proto_rawDescGZIP(), []int{3}
}

var File_cosmos_vesting_v1beta1_tx_proto protoreflect.FileDescriptor

var file_cosmos_vesting_v1beta1_tx_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xaf, 0x02, 0x0a, 0x17, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0c,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x66, 0x72,
	0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x74, 0x6f, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2,
	0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x63, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x30, 0xc8,
	0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x3a, 0x04, 0xe8, 0xa0,
	0x1f, 0x01, 0x22, 0x21, 0x0a, 0x1f, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x1f, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x56, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x0f, 0x76, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0e, 0x76, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x3a, 0x04, 0xe8, 0xa0, 0x1f, 0x00, 0x22,
	0x29, 0x0a, 0x27, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x69, 0x63, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa3, 0x02, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x80, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x37, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69,
	0x63, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a,
	0x3f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xe7, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42,
	0x07, 0x54, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x3b, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x56, 0x58, 0xaa, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xca, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x56, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x22, 0x43, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x5c, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x18, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cosmos_vesting_v1beta1_tx_proto_rawDescOnce sync.Once
	file_cosmos_vesting_v1beta1_tx_proto_rawDescData = file_cosmos_vesting_v1beta1_tx_proto_rawDesc
)

func file_cosmos_vesting_v1beta1_tx_proto_rawDescGZIP() []byte {
	file_cosmos_vesting_v1beta1_tx_proto_rawDescOnce.Do(func() {
		file_cosmos_vesting_v1beta1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_vesting_v1beta1_tx_proto_rawDescData)
	})
	return file_cosmos_vesting_v1beta1_tx_proto_rawDescData
}

var file_cosmos_vesting_v1beta1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cosmos_vesting_v1beta1_tx_proto_goTypes = []interface{}{
	(*MsgCreateVestingAccount)(nil),                 // 0: cosmos.vesting.v1beta1.MsgCreateVestingAccount
	(*MsgCreateVestingAccountResponse)(nil),         // 1: cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse
	(*MsgCreatePeriodicVestingAccount)(nil),         // 2: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount
	(*MsgCreatePeriodicVestingAccountResponse)(nil), // 3: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse
	(*v1beta1.Coin)(nil),                            // 4: cosmos.base.v1beta1.Coin
	(*Period)(nil),                                  // 5: cosmos.vesting.v1beta1.Period
}
var file_cosmos_vesting_v1beta1_tx_proto_depIdxs = []int32{
	4, // 0: cosmos.vesting.v1beta1.MsgCreateVestingAccount.amount:type_name -> cosmos.base.v1beta1.Coin
	5, // 1: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.vesting_periods:type_name -> cosmos.vesting.v1beta1.Period
	0, // 2: cosmos.vesting.v1beta1.Msg.CreateVestingAccount:input_type -> cosmos.vesting.v1beta1.MsgCreateVestingAccount
	2, // 3: cosmos.vesting.v1beta1.Msg.CreatePeriodicVestingAccount:input_type -> cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount
	1, // 4: cosmos.vesting.v1beta1.Msg.CreateVestingAccount:output_type -> cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse
	3, // 5: cosmos.vesting.v1beta1.Msg.CreatePeriodicVestingAccount:output_type -> cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cosmos_vesting_v1beta1_tx_proto_init() }
func file_cosmos_vesting_v1beta1_tx_proto_init() {
	if File_cosmos_vesting_v1beta1_tx_proto != nil {
		return
	}
	file_cosmos_vesting_v1beta1_vesting_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_vesting_v1beta1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateVestingAccount); i {
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
		file_cosmos_vesting_v1beta1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateVestingAccountResponse); i {
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
		file_cosmos_vesting_v1beta1_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreatePeriodicVestingAccount); i {
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
		file_cosmos_vesting_v1beta1_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreatePeriodicVestingAccountResponse); i {
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
			RawDescriptor: file_cosmos_vesting_v1beta1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_vesting_v1beta1_tx_proto_goTypes,
		DependencyIndexes: file_cosmos_vesting_v1beta1_tx_proto_depIdxs,
		MessageInfos:      file_cosmos_vesting_v1beta1_tx_proto_msgTypes,
	}.Build()
	File_cosmos_vesting_v1beta1_tx_proto = out.File
	file_cosmos_vesting_v1beta1_tx_proto_rawDesc = nil
	file_cosmos_vesting_v1beta1_tx_proto_goTypes = nil
	file_cosmos_vesting_v1beta1_tx_proto_depIdxs = nil
}
