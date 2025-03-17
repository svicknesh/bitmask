package bitmask

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/bits"
	"strconv"
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
	b.largest = Bit(1) << uint(size)
	//b.bv = 1 << (size - 1)

	return
}

// NewFromStr - creates a new instance of bitmask from a bit string
func NewFromStr(bitstr string) (b *Bitmask, err error) {

	if len(bitstr) == 0 {
		return nil, errors.New("newfromstr: bit string cannot be empty")
	}

	b = new(Bitmask)
	b.size = len(bitstr)

	bitvalue, err := strconv.ParseUint(bitstr, 2, 64)
	if nil != err {
		return nil, fmt.Errorf("newfromstr: error parsing bit string -> %w", err)
	}

	b.largest = Bit(1) << b.size
	b.bv = Bit(bitvalue)
	return
}

// SetLength - sets the length of the bit string
func (b *Bitmask) SetLength(newSize int) {
	b.largest = Bit(1) << uint(newSize)
	b.size = newSize
}

// Set - sets the given bits to 1, expanding size as needed for string padding and clearing
func (b *Bitmask) Set(bitsToSet ...Bit) {

	for _, bit := range bitsToSet {
		b.bv |= bit

		if bit >= b.largest {
			// the size can't accomodate the new largest bit, expand it
			b.size = bits.Len64(uint64(bit))
			b.largest = Bit(1) << uint(b.size)
		}

	}

}

// SetAll - sets all the bits to 1
func (b *Bitmask) SetAll() {
	b.bv |= Bit((1 << uint(b.size)) - 1)
}

// Remove - sets the given bits to 0
func (b *Bitmask) Remove(bits ...Bit) {
	for _, bit := range bits {
		b.bv &^= bit
	}
}

// Clear - clears all the bits to 0, essentially a blank instance
func (b *Bitmask) Clear() {
	b.bv = 0
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

// Uint64 - returns the interger value of the bit
func (b *Bitmask) Uint64() (i uint64) {
	return uint64(b.bv)
}

// MarshalJSON - returnsa JSON marshal bytes for the bitmask
func (b Bitmask) MarshalJSON() ([]byte, error) {
	// do not change this to `[]byte(b.String())`, calling json.Marshal will return it as a proper string
	return json.Marshal(b.String())
}

// UnmarshalJSON - returns an unmarshaled JSON of the bitmask
func (b *Bitmask) UnmarshalJSON(bitbytes []byte) (err error) {

	unquoted, err := strconv.Unquote(string(bitbytes))
	if err != nil {
		return fmt.Errorf("unmarshaljson: error unquoting bit string -> %w", err)
	}

	temp, err := NewFromStr(unquoted)
	if nil != err {
		return fmt.Errorf("unmarshaljson: error decoding bit string -> %w", err)
	}

	*b = *temp // assign the temp variable to the real bitmask

	return
}
