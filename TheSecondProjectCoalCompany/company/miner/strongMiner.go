package miner

import "github.com/google/uuid"

const (
	typeStrongMiner MinerType = "strong"
	strongMinerCost coalType  = 450
)

type strongMiner struct {
	id uuid.UUID
}
