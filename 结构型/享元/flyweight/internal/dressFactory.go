package internal

import "fmt"

//   dress 衣服
const (
	//Terrorist Dress Type terrorist dress type 恐怖分子服装类型
	TerroristDressType = "tDress"
	//Counter Terrrorist DressType terrorist dress type 反恐服装类型
	CounterTerrroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		DressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	DressMap map[string]Dress
}

func (d *DressFactory) GetDressByType(dressType string) (Dress, error) {
	if d.DressMap[dressType] != nil {
		return d.DressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.DressMap[dressType] = NewTerroristDress()
		return d.DressMap[dressType], nil
	}
	if dressType == CounterTerrroristDressType {
		d.DressMap[dressType] = NewCounterTerroristDress()
		return d.DressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func GetDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
