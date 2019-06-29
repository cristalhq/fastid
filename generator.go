package fastid

import (
	"errors"
	"sync"
	"time"
)

// Generator represents IDs with a given epoch and workedID.
type Generator struct {
	mutex sync.RWMutex

	epoch     uint64
	timestamp uint64
	workerID  uint64
	sequence  uint64

	lastTimestamp uint64
	lastID        uint64
}

// NewGenerator creates a new generator for IDs with a given epoch and workerID.
func NewGenerator(epoch int64, workerID int) (*Generator, error) {
	if workerID > MaxWorkerID {
		return nil, errors.New("workerID is too big")
	}

	g := &Generator{
		epoch:    uint64(epoch),
		workerID: uint64(workerID),
	}
	return g, nil
}

// Next returns a next ID.
func (g *Generator) Next() ID {
	g.mutex.Lock()

	nowNano := time.Now().UnixNano()
	g.timestamp = uint64(nowNano) / uint64(time.Millisecond)

	if g.timestamp <= g.lastTimestamp {
		g.sequence = (g.sequence + 1) & sequenceMask

		if g.sequence == 0 {
			g.timestamp = g.lastTimestamp + 1
		}
	} else {
		g.sequence = 0
	}

	ts := (g.timestamp - g.epoch) << timestampShift
	id := g.workerID << workerIDShift
	seq := g.sequence

	nextID := ts | id | seq

	g.lastID = nextID
	g.lastTimestamp = g.timestamp

	g.mutex.Unlock()
	return ID(nextID)
}

// LastID returns a last generated ID.
func (g *Generator) LastID() ID {
	return ID(g.lastID)
}

// LastSequence returns a last generated sequence.
func (g *Generator) LastSequence() int {
	return int(g.sequence)
}

// LastTimestamp returns a last generated timestamp.
func (g *Generator) LastTimestamp() int64 {
	return int64(g.lastTimestamp)
}
