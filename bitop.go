package bitmask

import (
	"fmt"
	"math"
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
	return (b.bv & bit) == 1
}

// String - returns a bit string with padding of size
func (b *Bitmask) String() (bitStr string) {
	return fmt.Sprintf("%0*b", b.size, b.bv)
}
