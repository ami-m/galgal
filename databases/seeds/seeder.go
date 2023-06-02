package seeds

var seedsArray []Seed = []Seed{
	TimeSlot{Common{Name: "TimeSlot"}},
}

// Iterate over the exists seeds and execute they Run fucntions
func Run() {
	for _, seed := range seedsArray {
		seed.Run()
	}
}
