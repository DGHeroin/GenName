package GenName

import (
    "fmt"
    "math/rand"
    "time"
)

type Generator interface {
    Generate() string
}

type NameGenerator struct{
    random *rand.Rand
}

func NewNameGenerator() Generator {
    gen := &NameGenerator{
        random:rand.New(rand.NewSource(time.Now().Unix())),
    }
    return gen
}

func (r *NameGenerator) Generate() string {
    randomAdjective := adjectives[r.random.Intn(len(adjectives))]
    randomNoun      := nouns[r.random.Intn(len(nouns))]

    return fmt.Sprintf("%v-%v", randomAdjective, randomNoun)
}