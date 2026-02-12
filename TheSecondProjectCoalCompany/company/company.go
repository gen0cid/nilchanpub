package company

import (
	"CoalCompany/company/equipment"
	"CoalCompany/company/miner"
	"CoalCompany/company/statistics"

	"github.com/google/uuid"
)

type Company struct {
	miners    map[miner.MinerType]map[uuid.UUID]miner.Miner
	equipment *equipment.Equipment

	statistics *statistics.StatisticsCompany
}

func newCompany() *Company {
	c := &Company{
		miners:     make(map[miner.MinerType]map[uuid.UUID]miner.Miner),
		equipment:  equipment.NewEquipment(),
		statistics: statistics.NewStatistics(),
	}

	return c
}
func baseIncome() {
for 
}
