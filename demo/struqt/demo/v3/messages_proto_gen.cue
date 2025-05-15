package demo

import (
	"time"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

#EchoValueRequest: {
	value?: _ @protobuf(1,google.protobuf.Value)
}

#EchoValueResponse: {
	value?:      _              @protobuf(1,google.protobuf.Value)
	timestamp?:  time.Time      @protobuf(2,google.protobuf.Timestamp)
	pb3Message?: #Proto3Message @protobuf(3,Proto3Message,name=pb3_message)
}
#EnumValue:
	#ENUM_VALUE_UNSPECIFIED |
	#ENUM_VALUE_X |
	#ENUM_VALUE_Y |
	#ENUM_VALUE_Z

#ENUM_VALUE_UNSPECIFIED: 0
#ENUM_VALUE_X:           1
#ENUM_VALUE_Y:           2
#ENUM_VALUE_Z:           3

#EnumValue_value: {
	ENUM_VALUE_UNSPECIFIED: 0
	ENUM_VALUE_X:           1
	ENUM_VALUE_Y:           2
	ENUM_VALUE_Z:           3
}

#Proto3Message: {
	nested?:      #Proto3Message @protobuf(41,Proto3Message)
	floatValue?:  float32        @protobuf(42,float,name=float_value)
	doubleValue?: float64        @protobuf(43,double,name=double_value)
	int64Value?:  int64          @protobuf(3,int64,name=int64_value)
	int32Value?:  int32          @protobuf(4,int32,name=int32_value)
	uint64Value?: uint64         @protobuf(5,uint64,name=uint64_value)
	uint32Value?: uint32         @protobuf(6,uint32,name=uint32_value)
	boolValue?:   bool           @protobuf(7,bool,name=bool_value)
	stringValue?: string         @protobuf(8,string,name=string_value)
	bytesValue?:  bytes          @protobuf(9,bytes,name=bytes_value)
	repeatedValue?: [...string] @protobuf(10,string,name=repeated_value)
	repeatedMessage?: [...null | uint64] @protobuf(44,google.protobuf.UInt64Value,name=repeated_message)
	enumValue?: #EnumValue @protobuf(11,EnumValue,name=enum_value)
	repeatedEnum?: [...#EnumValue] @protobuf(12,EnumValue,name=repeated_enum)
	timestampValue?:     time.Time              @protobuf(13,google.protobuf.Timestamp,name=timestamp_value)
	durationValue?:      time.Duration          @protobuf(14,google.protobuf.Duration,name=duration_value)
	fieldMaskValue?:     fieldmaskpb.#FieldMask @protobuf(15,google.protobuf.FieldMask,name=field_mask_value)
	wrapperDoubleValue?: null | float           @protobuf(17,google.protobuf.DoubleValue,name=wrapper_double_value)
	wrapperFloatValue?:  null | float           @protobuf(18,google.protobuf.FloatValue,name=wrapper_float_value)
	wrapperInt64Value?:  null | int64           @protobuf(19,google.protobuf.Int64Value,name=wrapper_int64_value)
	wrapperInt32Value?:  null | int32           @protobuf(20,google.protobuf.Int32Value,name=wrapper_int32_value)
	wrapperUInt64Value?: null | uint64          @protobuf(21,google.protobuf.UInt64Value,name=wrapper_u_int64_value)
	wrapperUInt32Value?: null | uint32          @protobuf(22,google.protobuf.UInt32Value,name=wrapper_u_int32_value)
	wrapperBoolValue?:   null | bool            @protobuf(23,google.protobuf.BoolValue,name=wrapper_bool_value)
	wrapperStringValue?: null | string          @protobuf(24,google.protobuf.StringValue,name=wrapper_string_value)
	wrapperBytesValue?:  null | bytes           @protobuf(25,google.protobuf.BytesValue,name=wrapper_bytes_value)
	mapValue?: {
		[string]: string
	} @protobuf(26,map[string]string,map_value)
	mapValue2?: {
		[string]: int32
	} @protobuf(27,map[string]int32,map_value2)
	mapValue3?: {
		[string]: string
	} @protobuf(28,map[int32]string,map_value3)
	mapValue4?: {
		[string]: int64
	} @protobuf(29,map[string]int64,map_value4)
	mapValue5?: {
		[string]: string
	} @protobuf(30,map[int64]string,map_value5)
	mapValue6?: {
		[string]: uint32
	} @protobuf(31,map[string]uint32,map_value6)
	mapValue7?: {
		[string]: string
	} @protobuf(32,map[uint32]string,map_value7)
	mapValue8?: {
		[string]: uint64
	} @protobuf(33,map[string]uint64,map_value8)
	mapValue9?: {
		[string]: string
	} @protobuf(34,map[uint64]string,map_value9)
	mapValue10?: {
		[string]: float32
	} @protobuf(35,map[string]float,map_value10)
	mapValue12?: {
		[string]: float64
	} @protobuf(37,map[string]double,map_value12)
	mapValue14?: {
		[string]: bool
	} @protobuf(39,map[string]bool,map_value14)
	mapValue15?: {
		[string]: string
	} @protobuf(40,map[bool]string,map_value15)
	mapValue16?: {
		[string]: null | uint64
	} @protobuf(45,map[string]google.protobuf.UInt64Value,map_value16)
	structValueValue?: _ @protobuf(47,google.protobuf.Value,name=struct_value_value)
	structValue?: {} @protobuf(48,google.protobuf.Struct,name=struct_value)
	// Next number: 49
}

#Proto3OneOf: {
	{} | {
		oneofBoolValue: bool @protobuf(11,bool,name=oneof_bool_value)
	} | {
		oneofStringValue: string @protobuf(12,string,name=oneof_string_value)
	}
	{} | {
		nestedOneofValueOne: #Proto3Message @protobuf(21,Proto3Message,name=nested_oneof_value_one)
	}
}
