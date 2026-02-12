package miner

import "github.com/google/uuid"

const (
	typeNormalMiner MinerType = "normal"
	normalMinerCost coalType  = 50
)

type normalMiner struct {
	id uuid.UUID
}
