package characteristic

type ByteCharacteristic struct {
    *Number
}

func NewByteCharacteristic(value byte) *ByteCharacteristic {
    number := NewNumber(value, nil, nil, nil, FormatByte)
    return &ByteCharacteristic{number}
}

func (c *ByteCharacteristic) SetByte(value byte) {
    c.SetNumber(value)
}

func (c *ByteCharacteristic) Byte() byte {
    return c.GetValue().(byte)
}