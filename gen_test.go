package GenName

import "testing"

func TestGenName(t *testing.T) {
    gen := NewNameGenerator()
    t.Logf("%v", gen.Generate())
    t.Logf("%v", gen.Generate())
}