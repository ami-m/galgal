package models

type Address struct {
	Street   string `json:"street"`
	Line1    string `json:"line1"`
	Line2    string `json:"line2"`
	Country  string `json:"country"`
	PostCode string `json:postcode"`
}

// Although address is not possessing any DB table, we deliberately
// implementing the GetTableName function in order to bind Address to Model interface
func (a Address) GetTableName() string {
	return ""
}
