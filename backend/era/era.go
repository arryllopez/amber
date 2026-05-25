// Defines geological era metadata and time ranges.
package era

type Era struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	StartMa     float64 `json:"startMa"`
	EndMa       float64 `json:"endMa"`
	Description string  `json:"description"`
}

func List() []Era {
	return []Era{
		{
			ID:          "pleistocene",
			Name:        "Pleistocene",
			StartMa:     2.58,
			EndMa:       0.0117,
			Description: "Ice ages, megafauna, and early human environments.",
		},
		{
			ID:          "cretaceous",
			Name:        "Cretaceous",
			StartMa:     145,
			EndMa:       66,
			Description: "High sea levels, warm climate, dinosaurs, and flowering plants.",
		},
		{
			ID:          "jurassic",
			Name:        "Jurassic",
			StartMa:     201,
			EndMa:       145,
			Description: "Warm humid climates, conifer forests, cycads, and major dinosaur radiation.",
		},
	}
}

func Find(id string) (Era, bool) {
	for _, current := range List() {
		if current.ID == id {
			return current, true
		}
	}

	return Era{}, false
}
