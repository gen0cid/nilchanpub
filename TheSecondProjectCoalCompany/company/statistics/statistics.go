package statistics

import (
	"CoalCompany/company/miner"
	"sync/atomic"
	"time"
)

type StatisticsCompany struct {
	balance      atomic.Int64 // текущий баланс
	totalBalance atomic.Int64 // баланс за все время

	minerStatistics map[miner.MinerType]int // количество всех шахтеров которых нанимал пользователь

	startGameTime  time.Time
	finishGameTime *time.Time
}
