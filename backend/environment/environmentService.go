// Assembles tectonic, climate, and fauna data into environment profiles.
package environment

import (
	"errors"
	"fmt"

	"amber/backend/asset"
	"amber/backend/era"
)

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PreviewRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	EraID     string  `json:"eraId"`
}

type Classification struct {
	Biome              string  `json:"biome"`
	TemperatureCelsius float64 `json:"temperatureCelsius"`
	Precipitation      string  `json:"precipitation"`
	WaterPresence      string  `json:"waterPresence"`
	TerrainProfile     string  `json:"terrainProfile"`
	Confidence         string  `json:"confidence"`
}

type SceneAssembly struct {
	AssetPack    asset.Pack         `json:"assetPack"`
	TerrainMode  string             `json:"terrainMode"`
	ScatterRules map[string]float64 `json:"scatterRules"`
	Atmosphere   map[string]string  `json:"atmosphere"`
	Soundscape   string             `json:"soundscape"`
}

type CreatureLayer struct {
	Presence        string   `json:"presence"`
	Confidence      string   `json:"confidence"`
	Species         []string `json:"species"`
	Interaction     string   `json:"interaction"`
	MaxVisibleCount int      `json:"maxVisibleCount"`
}

type PreviewResponse struct {
	ModernCoordinate Coordinate     `json:"modernCoordinate"`
	PaleoCoordinate  Coordinate     `json:"paleoCoordinate"`
	Era              era.Era        `json:"era"`
	Environment      Classification `json:"environment"`
	SceneAssembly    SceneAssembly  `json:"sceneAssembly"`
	CreatureLayer    CreatureLayer  `json:"creatureLayer"`
	Notes            []string       `json:"notes"`
}

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) Preview(request PreviewRequest) (PreviewResponse, error) {
	selectedEra, ok := era.Find(request.EraID)
	if !ok {
		return PreviewResponse{}, fmt.Errorf("unknown era: %s", request.EraID)
	}

	if request.Latitude < -90 || request.Latitude > 90 {
		return PreviewResponse{}, errors.New("latitude must be between -90 and 90")
	}

	if request.Longitude < -180 || request.Longitude > 180 {
		return PreviewResponse{}, errors.New("longitude must be between -180 and 180")
	}

	classification := Classification{
		Biome:              "humid_conifer_forest",
		TemperatureCelsius: 32,
		Precipitation:      "high",
		WaterPresence:      "none",
		TerrainProfile:     "lowland",
		Confidence:         "mock",
	}

	return PreviewResponse{
		ModernCoordinate: Coordinate{
			Latitude:  request.Latitude,
			Longitude: request.Longitude,
		},
		PaleoCoordinate: Coordinate{
			Latitude:  request.Latitude - 8.4,
			Longitude: request.Longitude + 24.2,
		},
		Era:         selectedEra,
		Environment: classification,
		SceneAssembly: SceneAssembly{
			AssetPack:   asset.MatchBiomePack(selectedEra.ID, classification.Biome),
			TerrainMode: "flattened_lowland",
			ScatterRules: map[string]float64{
				"coniferDensity": 0.70,
				"fernUnderstory": 0.40,
				"cycadClusters":  0.15,
			},
			Atmosphere: map[string]string{
				"haze":             "medium",
				"colorTemperature": "warm",
				"sunAngleSource":   "paleolatitude",
			},
			Soundscape: "humid_forest_day_v1",
		},
		CreatureLayer: CreatureLayer{
			Presence:        "ambient_distant",
			Confidence:      "low",
			Species:         []string{"ornithopod_group", "pterosaur_group"},
			Interaction:     "none",
			MaxVisibleCount: 3,
		},
		Notes: []string{
			"Mocked environment profile for the first vertical slice.",
			"Replace tectonic, climate, and fauna values with real data integrations in later phases.",
		},
	}, nil
}
