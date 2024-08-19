package utils

import "marmita/types"

// Define enums for initial flow state
const (
	HOME types.FlowState = iota
	REGISTER
	QUERY
	UPDATE
)

// Define action-level enum for REGISTER flow state
const (
	REGISTER_CLIENT types.RegisterState = iota
	REGISTER_MARMITA
	REGISTER_DELIVERY
	REGISTER_ORDER
)
