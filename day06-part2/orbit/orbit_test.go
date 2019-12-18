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

func TestPathOfTwo(t *testing.T) {
	system := NewSystem("COM")
	AddDeclaration(ParseDeclaration("COM)B"), system)

	path := Path(system, "B")

	assert.Equal(t, len(path), 2)
	assert.Equal(t, path[0].Label, "B")
	assert.Equal(t, path[1].Label, "COM")
}

func TestPathInTree(t *testing.T) {
	system := NewSystem("COM")
	AddDeclaration(ParseDeclaration("COM)A"), system)
	AddDeclaration(ParseDeclaration("A)B"), system)
	AddDeclaration(ParseDeclaration("B)C"), system)
	AddDeclaration(ParseDeclaration("B)D"), system)
	AddDeclaration(ParseDeclaration("D)E"), system)

	path := Path(system, "C")

	assert.Equal(t, len(path), 4)
	assert.Equal(t, path[0].Label, "C")
	assert.Equal(t, path[1].Label, "B")
	assert.Equal(t, path[2].Label, "A")
	assert.Equal(t, path[3].Label, "COM")
}

func TestPathBetween(t *testing.T) {
	system := NewSystem("COM")
	AddDeclaration(ParseDeclaration("COM)A"), system)
	AddDeclaration(ParseDeclaration("A)B"), system)
	AddDeclaration(ParseDeclaration("B)C"), system)
	AddDeclaration(ParseDeclaration("C)D"), system)
	AddDeclaration(ParseDeclaration("D)E"), system)

	pathBetween := PathBetween(system, "D", "B")

	assert.Equal(t, 3, len(pathBetween))
	assert.Equal(t, "D", pathBetween[0].Label)
	assert.Equal(t, "C", pathBetween[1].Label)
	assert.Equal(t, "B", pathBetween[2].Label)
}

func TestCommonBody(t *testing.T) {
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
	AddDeclaration(ParseDeclaration("K)YOU"), system)
	AddDeclaration(ParseDeclaration("I)SAN"), system)

	commonBody := CommonBody(system, "YOU", "SAN")

	assert.Equal(t, "D", commonBody.Label)
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

func TestTransferCountExapmle(t *testing.T) {
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
	AddDeclaration(ParseDeclaration("K)YOU"), system)
	AddDeclaration(ParseDeclaration("I)SAN"), system)

	count := TransferCount(system, "YOU", "SAN")

	assert.Equal(t, 4, count)
}
