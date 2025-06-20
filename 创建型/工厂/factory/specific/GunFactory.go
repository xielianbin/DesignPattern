package specific

import (
	"factory/abstract"
	"fmt"
)

func GetGun(gunType string) (abstract.IGun, error) {
	g := Gun{}
	switch gunType {
	case "ak47":
		g.SetName("AK-47")
		g.SetPower(7)
		g.SetRange(500)
		return &Ak47{Gun: g}, nil
	case "musket":

		g.SetName("musket")
		g.SetPower(8)
		g.SetRange(300)
		return &Musket{Gun: g}, nil
	default:
		return nil, fmt.Errorf("unknown gun type: %s", gunType)
	}
}
