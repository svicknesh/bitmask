package bitmask

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Bit - new type for bitmask
type Bit uint64

// Bitmask - structure for storing bitmask information
type Bitmask struct {
	bv      Bit
	largest Bit
	size    int
}

// New - creates new instance of bitmask, size is meant for printing the output string with padding. if the bits exceed the size, the new size will be printed
func New(size int) (b *Bitmask) {

	b = new(Bitmask)
	b.size = size
	b.largest = Bit(math.Exp2(float64(size)))
	//b.bv = 1 << (size - 1)

	return
}

// NewFromStr - creates a new instance of bitmask from a bit string
func NewFromStr(bitstr string) (b *Bitmask, err error) {
	b = new(Bitmask)
	b.size = len(bitstr)

	bitvalue, err := strconv.ParseUint(bitstr, 2, 64)
	if nil != err {
		return nil, fmt.Errorf("newfromstr: error parsing bit string -> %w", err)
	}

	b.largest = 1 << b.size
	b.bv = Bit(bitvalue)
	return
}

// SetLength - sets the length of the bit string
func (b *Bitmask) SetLength(len int) {
	b.largest = 1 << len
	b.size = len
}

// Set - sets the given bits to 1, expanding size as needed for string padding and clearing
func (b *Bitmask) Set(bits ...Bit) {

	for _, bit := range bits {
		b.bv |= bit

		if bit >= b.largest {
			// the size can't accomodate the new largest bit, expand it
			b.largest = bit
			v := int(math.Log2(float64(bit)))
			b.size = v + 1
		}

	}

}

// SetAll - sets all the bits to 1
func (b *Bitmask) SetAll() {

	var t uint64

	for i := 0; i < b.size; i++ {
		t += uint64(math.Exp2(float64(i)))
	}

	b.bv |= Bit(t)

}

// Remove - sets the given bits to 0
func (b *Bitmask) Remove(bits ...Bit) {
	for _, bit := range bits {
		// //b.bv &= ^bit
		b.bv &^= bit
	}
}

// Clear - clears all the bits to 0, essentially a blank instance
func (b *Bitmask) Clear() {
	b.bv = b.bv >> Bit(63) // max is 64 bits, so we just shift it, there is no real benefit for shifting by the size
}

// Toggle - flip the given bits
func (b *Bitmask) Toggle(bits ...Bit) {
	for _, bit := range bits {
		b.bv ^= bit
	}
}

// Has - checks if a given bit is set
func (b *Bitmask) Has(bit Bit) (ok bool) {
	return (b.bv & bit) != 0
}

// String - returns a bit string with padding of size
func (b *Bitmask) String() (bitStr string) {
	return fmt.Sprintf("%0*b", b.size, b.bv)
}

// MarshalJSON - returnsa JSON marshal bytes for the bitmask
func (b Bitmask) MarshalJSON() ([]byte, error) {
	// do not change this to `[]byte(b.String())`, calling json.Marshal will return it as a proper string
	return json.Marshal(b.String())
}

// UnmarshalJSON - returnsa unmarshaled JSON of the bitmask
func (b *Bitmask) UnmarshalJSON(bitbytes []byte) (err error) {

	bitstr := string(bitbytes)
	bitstr = strings.ReplaceAll(bitstr, "\"", "") // strip the quotes (") before using

	b1, err := NewFromStr(bitstr)
	if nil != err {
		return fmt.Errorf("unmarshaljson: error decoding bit string -> %w", err)
	}

	*b = *b1 // assign the temp variable to the real bitmask

	return
}
