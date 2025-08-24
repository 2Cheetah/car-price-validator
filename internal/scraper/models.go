package scraper

type Root struct {
	Data []Ad `json:"data"`
	Meta Meta `json:"meta"`
}

type Ad struct {
	Attributes Attributes `json:"attributes"`
}

type Meta struct {
	TotalResults int `json:"total-results"`
	TotalShowing int `json:"total-showing"`
}

type Attributes struct {
	MakeName         string  `json:"make_name"`
	ModelName        string  `json:"model_name"`
	Price            int     `json:"price"`
	ManufacturedYear string  `json:"manufactured_year"`
	Mileage          Mileage `json:"mileage"`
}

type Mileage struct {
	GTE string `json:"gte"`
	LTE string `json:"lte"`
}

type Car struct {
	Brand     string
	Model     string
	MilageGT  string
	MilageLT  string
	ModelYear string
	Price     int
}

type BarData struct {
	Name       string
	LowerLimit int
	UpperLimit int
	Items      int
}

type Filters struct {
	From        string
	Limit       string
	MakeID      string
	MfgYearFrom string
	MfgYearTo   string
	MilageFrom  string
	MilageTo    string
	ModelID     string
	PriceFrom   string
	PriceTo     string
	Type        string
}
