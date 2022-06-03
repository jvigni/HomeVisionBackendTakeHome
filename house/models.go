package house

type Houses []House

type House struct {
	Id        int
	Address   string
	Homeowner string
	Price     int
	PhotoURL  string
}

type HousesResponse struct {
	Houses []House
	Ok     bool
}
