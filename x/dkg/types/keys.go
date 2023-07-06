package types

const (
	// ModuleName defines the module name
	ModuleName = "dkg"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dkg"
)
const (
	EventTypeKeygen = "keygen"
)
const (
	AttributeValueSend     = "send"
	AttributeValueStart    = "start"
	AttributeValueMsg      = "message"
	AttributeValueDecided  = "decided"
	AttributeValueReject   = "reject"
	AttributeValueAssigned = "assigned"
	AttributeValueDispute  = "dispute"
)
const (
	DisputeKey      = "Dispute/value/"
	DisputeCountKey = "Commit/count/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
