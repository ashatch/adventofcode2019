package orbit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	system := NewSystem("COM")

	AddDeclaration(ParseDeclaration("COM)A"), system)
	AddDeclaration(ParseDeclaration("A)B"), system)
	AddDeclaration(ParseDeclaration("A)C"), system)
	AddDeclaration(ParseDeclaration("B)D"), system)

	stats := Count(system)

	assert.Equal(t, 4, stats.DirectOrbitCount)
	assert.Equal(t, 4, stats.IndirectOrbitCount)
}

func TestExampleCount(t *testing.T) {
	system := NewSystem("COM")
	AddDeclaration(ParseDeclaration("COM)B"), system)
	AddDeclaration(ParseDeclaration("B)C"), system)
	AddDeclaration(ParseDeclaration("C)D"), system)
	AddDeclaration(ParseDeclaration("D)E"), system)
	AddDeclaration(ParseDeclaration("E)F"), system)
	AddDeclaration(ParseDeclaration("B)G"), system)
	AddDeclaration(ParseDeclaration("G)H"), system)
	AddDeclaration(ParseDeclaration("D)I"), system)
	AddDeclaration(ParseDeclaration("E)J"), system)
	AddDeclaration(ParseDeclaration("J)K"), system)
	AddDeclaration(ParseDeclaration("K)L"), system)

	stats := Count(system)

	assert.Equal(t, 11, stats.DirectOrbitCount)
	assert.Equal(t, 31, stats.IndirectOrbitCount)
}
