package demo

#Proto3MessageCUE: #Proto3Message & {
	int32Value?: int32 & >=0 & <20
	mapValue2: {
		more?: int32 & >=-128 & <128
	}
}

validate: {
	for k, v in examples {
		"\(k)": {
			if v.int32Value != _|_ {
				assert: v.int32Value <= 100
			}
			if v.nested != _|_ && v.nested.int32Value != _|_ && v.int32Value != _|_ {
				assert: v.nested.int32Value != v.int32Value
			}
		}
	}
}

examples: {
	"example1": #Proto3MessageCUE & {
		int32Value:  0
		floatValue:  3.14
		doubleValue: 6.28
	}
	"example2": #Proto3MessageCUE & {
		int32Value: 11
		int64Value: 9999
		nested: {
			int32Value: 12
		}
		mapValue2: {
			more: 127
		}
	}
}
