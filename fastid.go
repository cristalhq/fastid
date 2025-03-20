package fastid

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

const (
	// DefaultEpoch is an epoch for a default generator.
	DefaultEpoch = 1_700_000_000_000

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

// Parse given string as [ID].
func Parse(s string) (ID, error) {
	n, err := strconv.ParseUint(s, 10, 64)
	return ID(n), err
}

func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// Timestamp of the ID (from generator epoch).
func (id ID) Timestamp() int64 {
	return int64(id) >> timestampShift
}

// WorkerID returns ID's worker id.
func (id ID) WorkerID() int {
	return (int(id) >> workerIDShift) & MaxWorkerID
}

// Sequence returns ID's sequence number.
func (id ID) Sequence() int {
	return int(id) & MaxSequenceID
}

// Generator represents IDs with a given epoch and workedID.
type Generator struct {
	epoch    uint64
	workerID uint64

	mu       sync.Mutex
	sequence uint64
	lastTS   uint64
	lastID   uint64
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

	now := uint64(time.Now().UnixMilli())

	switch {
	case now > g.lastTS:
		g.sequence = 0
	case now == g.lastTS:
		if (g.sequence + 1) <= MaxSequenceID {
			g.sequence++
		} else {
			g.sequence = 0
			now++
		}
	}
	g.lastTS = now

	ts := (now - g.epoch) << timestampShift
	id := g.workerID << workerIDShift
	seq := g.sequence

	nextID := ts | id | seq

	g.lastID = nextID

	g.mu.Unlock()
	return ID(nextID)
}

// LastID returns a last generated ID.
func (g *Generator) LastID() ID {
	g.mu.Lock()
	defer g.mu.Unlock()

	return ID(g.lastID)
}

// LastTS returns a last generated timestamp.
func (g *Generator) LastTS() uint64 {
	g.mu.Lock()
	defer g.mu.Unlock()

	return g.lastTS
}

// LastSequence returns a last generated sequence.
func (g *Generator) LastSequence() int {
	g.mu.Lock()
	defer g.mu.Unlock()

	return int(g.sequence)
}
