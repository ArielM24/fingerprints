package kernel

import (
	"math"

	"github.com/alevinval/fingerprints/internal/matrix"
)

type variance struct {
	Base
	phy    *matrix.M
	offset int
}

func NewVariance(directional *matrix.M) *variance {
	k := &variance{phy: directional, offset: 8}
	k.Base = Base{kernel: k}
	return k
}

func (k *variance) Offset() int {
	return k.offset
}

func (k *variance) Apply(in *matrix.M, x, y int) float64 {
	var pos int

	sigSize := float64(k.Offset()*2 + 1)
	signature := make([]float64, int(sigSize))
	angle := k.phy.At(x, y) - math.Pi/2
	for j := y - k.Offset(); j <= y+k.Offset(); j++ {
		for i := x - k.Offset(); i <= x+k.Offset(); i++ {
			xp := int(float64(i-x)*math.Cos(angle)-float64(j-y)*math.Sin(angle)) + x
			yp := int(float64(i-x)*math.Sin(angle)+float64(j-y)*math.Cos(angle)) + y
			if xp >= 0 && xp < in.Bounds().Dx() && yp >= 0 && yp < in.Bounds().Dy() {
				signature[pos] += in.At(xp, yp)
			} else {
				signature[pos] += in.At(x, y)
			}
		}
		pos++
	}

	sum := 0.0
	for _, sig := range signature {
		sum += sig
	}
	mean := sum / sigSize

	variance := 0.0
	for _, sig := range signature {
		d := sig - mean
		variance += d * d
	}
	variance /= sigSize
	return variance
}
