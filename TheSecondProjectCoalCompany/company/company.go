package company

import (
	"CoalCompany/company/equipment"
	"CoalCompany/company/miner"
	"CoalCompany/company/statistics"

	"github.com/google/uuid"
)

type company struct {
	miners    map[miner.MinerType]map[uuid.UUID]miner.Miner
	equipment *equipment.Equipment

	statistics *statistics.StatisticsCompany
}
