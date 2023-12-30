package fastid

import (
	"errors"
	"sync"
	"time"
)

const (
	// DefaultEpoch is an epoch for a default generator.
	DefaultEpoch = 1_500_000_000_000

	// workerIDBits is how many bits are used for worked ID.
	workerIDBits = uint(10)
	// sequenceBits is how many bits are used for sequence number.
	sequenceBits = uint(12)

	// MaxWorkerID is a max possible worked ID.
	MaxWorkerID = (1 << workerIDBits) - 1
	// MaxSequenceID is a max possible sequence number.
	MaxSequenceID = (1 << sequenceBits) - 1

	workerIDShift  = sequenceBits
	timestampShift = sequenceBits + workerIDBits
)

// ID represents an ID value.
type ID uint64

// Parts returns ID's parts: timestamp, worker id, sequence number.
func (id ID) Parts() (int64, int, int) {
	return id.Timestamp(), id.WorkerID(), id.Sequence()
}

// Timestamp of the ID (from generator epoch).
func (id ID) Timestamp() int64 {
	return int64(id) >> timestampShift
}

// WorkerID returns ID's worker id.
func (id ID) WorkerID() int {
	return int(id) >> workerIDShift
}

// Sequence returns ID's sequence number.
func (id ID) Sequence() int {
	return int(id) & MaxSequenceID
}

// Generator represents IDs with a given epoch and workedID.
type Generator struct {
	epoch    uint64
	workerID uint64

	mu            sync.Mutex
	timestamp     uint64
	sequence      uint64
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
	g.mu.Lock()

	g.timestamp = uint64(time.Now().UnixMilli())

	if g.timestamp <= g.lastTimestamp {
		g.sequence = (g.sequence + 1) & MaxSequenceID

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

	g.mu.Unlock()
	return ID(nextID)
}

// LastID returns a last generated ID.
func (g *Generator) LastID() ID {
	g.mu.Lock()
	defer g.mu.Unlock()

	return ID(g.lastID)
}

// LastTimestamp returns a last generated timestamp.
func (g *Generator) LastTimestamp() uint64 {
	g.mu.Lock()
	defer g.mu.Unlock()

	return g.lastTimestamp
}

// LastSequence returns a last generated sequence.
func (g *Generator) LastSequence() int {
	g.mu.Lock()
	defer g.mu.Unlock()

	return int(g.sequence)
}
