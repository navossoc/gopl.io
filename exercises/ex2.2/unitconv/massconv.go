package unitconv

import "fmt"

type Kilogram float64
type Pound float64

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }

// KGToP converts Kilograms weight to Pounds.
func KGToP(k Kilogram) Pound { return Pound(k * 2.2046) }

// PToKG converts Pounds weight to Kilograms.
func PToKG(p Pound) Kilogram { return Kilogram(p / 2.2046) }
