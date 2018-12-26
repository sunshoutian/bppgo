package bppgo

import (
	"fmt"
	"math"
	"sort"
)

// Bin represents a container in which items will be put into.
type Bin struct {
	Name      string
	Width     float64
	Height    float64
	Depth     float64
	MaxWeight float64

	Items []*Item // Items that packed in this bin
}
