// Defines 3D asset manifest and CDN URL boundaries.
package asset

type Pack struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Version     string `json:"version"`
}

func MatchBiomePack(eraID string, biome string) Pack {
	if eraID == "jurassic" && biome == "humid_conifer_forest" {
		return Pack{
			ID:          "jurassic_humid_conifer_forest_v1",
			DisplayName: "Jurassic Humid Conifer Forest",
			Version:     "v1",
		}
	}

	return Pack{
		ID:          eraID + "_" + biome + "_v1",
		DisplayName: biome,
		Version:     "v1",
	}
}
