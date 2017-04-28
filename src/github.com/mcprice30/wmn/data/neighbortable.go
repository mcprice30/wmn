package data

const LinkTypeUnidirectional = 0
const LinkTypeBidirectional = 1
const LinkTypeMPR = 2

type NeighborTableRow struct {
	LinkType        int
	TwoHopNeigbhors []ManetAddr
}

type NeighborTable struct {
	Table     map[ManetAddr]*NeighborTableRow
	Selectors map[ManetAddr]bool
}

func CreateNeighborTable() *NeighborTable {
	return &NeighborTable{
		Table:     map[ManetAddr]*NeighborTableRow{},
		Selectors: map[ManetAddr]bool{},
	}
}
