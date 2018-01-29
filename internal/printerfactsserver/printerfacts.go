package printerfactsserver

import (
	"context"
	"math/rand"

	"github.com/Xe/printerfacts/rpc/printerfacts"
)

// Impl implements printerfacts.Printerfacts.
type Impl struct {
	Facts []string
}

// Fact grabs a random set of printer facts and returns them to the user.
func (i *Impl) Fact(ctx context.Context, prm *printerfacts.FactParams) (*printerfacts.Facts, error) {
	result := &printerfacts.Facts{}

	if prm.Count == 0 {
		prm.Count = 1
	}

	for range make([]struct{}, prm.Count) {
		result.Facts = append(result.Facts, i.Facts[rand.Intn(len(i.Facts))])
	}

	return result, nil
}
