package miner

import "github.com/google/uuid"

const (
	typelittleMiner MinerType = "little"
	littleMinerCost coalType  = 5
)

type littleMiner struct {
	id uuid.UUID
}
