package main

// Center Detail
type Center struct {
	CenterId     int           `json:"center_id"`
	Name         string        `json:"name"`
	Address      string        `json:"address"`
	StateName    string        `json:"state_name"`
	DistrictName string        `json:"district_name"`
	BlockName    string        `json:"block_name"`
	PinCode      int           `json:"pincode"`
	Latitude     int           `json:"lat"`
	Longitude    int           `json:"long"`
	From         string        `json:"from"`
	To           string        `json:"to"`
	FeeType      string        `json:"fee_type"`
	Sessions     []Sessionlist `json:"sessions"`
}

//Sessionlist
type Sessionlist struct {
	Sessionid         string   `json:"session_id"`
	Date              string   `json:"date"`
	AvailableCapacity int      `json:"available_capacity"`
	MinAgeLimit       int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}

//Centerlist
type Centerlist struct {
	CentersArray []Center `json:"centers"`
}
