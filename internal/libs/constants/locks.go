package constants

// key used to control access to resources that are accessed concurrently
type Lock string

const (
	InventorySupplyLock Lock = "inventory_supply_lock"
)
