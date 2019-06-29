package fastid

const (
	// DefaultEpoch is an epoch for a default generator.
	DefaultEpoch = 1500000000000

	// WorkerIDBits is how many bits are used for worked ID.
	WorkerIDBits = uint(10)
	// SequenceBits is how many bits are used for sequence number.
	SequenceBits = uint(12)

	workerIDShift  = SequenceBits
	timestampShift = SequenceBits + WorkerIDBits

	// MaxWorkerID is a max possible worked ID.
	MaxWorkerID = (1 << WorkerIDBits) - 1
	// MaxSequenceID is a max possible sequence number.
	MaxSequenceID = (1 << SequenceBits) - 1
	sequenceMask  = (1 << SequenceBits) - 1
)

// ID represents an ID value.
type ID uint64
