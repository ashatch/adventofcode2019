package orbit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDeclaration(t *testing.T) {
	declaration := ParseDeclaration("A)B")

	assert.Equal(t, declaration.body, "A")
	assert.Equal(t, declaration.satellite, "B")
}

func TestAddingToSystem(t *testing.T) {
	system := newSystem("COM")
	planetA := ParseDeclaration("COM)A")
	planetB := ParseDeclaration("COM)B")

	AddDeclaration(planetA, system)
	AddDeclaration(planetB, system)

	assert.Equal(t, len(system.bodies), 3)

	comBody := system.bodies["COM"]
	assert.Equal(t, system.root, comBody)
	assert.Equal(t, len(comBody.Satellites), 2)
	assert.Equal(t, comBody.Satellites[0].Orbiting, comBody)
	assert.Equal(t, comBody.Satellites[1].Orbiting, comBody)
}
