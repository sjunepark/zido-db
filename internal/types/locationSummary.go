package types

type LocationSummary struct {
	ID            int     `db:"id"`
	Address       string  `db:"address"`
	AddressGroup  string  `db:"sdSggEm"`
	RoadNameGroup string  `db:"addrDetail"`
	Lat           float64 `db:"lat"`
	Long          float64 `db:"long"`
	X             float64 `db:"x"`
	Y             float64 `db:"y"`
}

func (ls LocationSummary) TableName() string {
	return "locations_summary"
}
