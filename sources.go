package main

import (
	"fmt"
	"strings"
)

func getSources() map[string]Source {
	sources := map[string]Source{
		// ABasin [requires HTML parsing]: https://www.arapahoebasin.com/snow-conditions/terrain/
		// Aspen [requires two queries]:
		//   https://www.aspensnowmass.com/aspensnowmass/liftstatus/feed?mountain=AspenMountain
		//   https://www.aspensnowmass.com/aspensnowmass/groomingreport/feed?mountain=AspenMountain
		// Aspen Highlands [requires two queries]
		//   https://www.aspensnowmass.com/aspensnowmass/liftstatus/feed?mountain=AspenHighlands
		//   https://www.aspensnowmass.com/aspensnowmass/groomingreport/feed?mountain=AspenHighlands
		// Buttermilk [requires two queries]
		//   https://www.aspensnowmass.com/aspensnowmass/liftstatus/feed?mountain=Buttermilk
		//   https://www.aspensnowmass.com/aspensnowmass/groomingreport/feed?mountain=Buttermilk
		// Echo Mountain [requires HTML parsing]: https://echomntn.com/mountain-conditions
		// Loveland [requires HTML parsing]: https://skiloveland.com/trail-lift-report/
		// Monarch [requires HTML parsing]: https://www.skimonarch.com/conditions/
		// Powderhorn [requires two queries and HTML parsing]:
		//   https://www.powderhorn.com/explore/conditions/lift-status.html
		//   https://powderhorn.com/explore/conditions/grooming-report.html
		// Purgatory [requires HTML parsing - WP]: https://www.purgatoryresort.com/snow-report/
		// Snowmass [requires two queries]
		//   https://www.aspensnowmass.com/aspensnowmass/liftstatus/feed?mountain=Snowmass
		//   https://www.aspensnowmass.com/aspensnowmass/groomingreport/feed?mountain=Snowmass
		// Sunlight [requires HTML parsing - Drupal]: https://sunlightmtn.com/mountain/winter-fun/snow-report
		// Telluride [requires HTML parsing]: https://www.tellurideskiresort.com/the-mountain/snow-report/
		//
		// Howelsen Hill [not available online]
		// Silverton [not available online]
		// Wolf Creek [not available online]
		// Cooper [requires JPG parsing]: https://www.skicooper.com/snow-trail-report/
		"Eldora":       Source{Name: "Eldora", Provider: "powder", Url: "https://www.eldora.com/api/v1/dor/status"},
		"Copper":       Source{Name: "Copper", Provider: "powder", Url: "https://www.coppercolorado.com/api/v1/dor/status"},
		"Stratton":     Source{Name: "Stratton", Provider: "alterra", Url: "https://mtnpowder.com/feed?resortId=1"},
		"Snowshoe":     Source{Name: "Snowshoe", Provider: "alterra", Url: "https://mtnpowder.com/feed?resortId=2"},
		"Blue":         Source{Name: "Blue", Provider: "alterra", Url: "https://mtnpowder.com/feed?resortId=3"},
		"Tremblant":    Source{Name: "Tremblant", Provider: "alterra", Url: "https://mtnpowder.com/feed?resortId=4"},
		"WinterPark":   Source{Name: "Winter Park", Provider: "alterra", Url: "https://mtnpowder.com/feed?resortId=5"},
		"Steamboat":    Source{Name: "Steamboat", Provider: "alterra", Url: "https://mtnpowder.com/feed?resortId=6"},
		"Vail":         Source{Name: "Vail", Provider: "snowcom", Url: "https://cache.snow.com/api/TerrainApi/GetTerrainStatus", Method: "POST", Args: "SiteId=1"},
		"BeaverCreek":  Source{Name: "BeaverCreek", Provider: "snowcom", Url: "https://cache.snow.com/api/TerrainApi/GetTerrainStatus", Method: "POST", Args: "SiteId=2"},
		"Keystone":     Source{Name: "Keystone", Provider: "snowcom", Url: "https://cache.snow.com/api/TerrainApi/GetTerrainStatus", Method: "POST", Args: "SiteId=3"},
		"Breckenridge": Source{Name: "Breckenridge", Provider: "snowcom", Url: "https://cache.snow.com/api/TerrainApi/GetTerrainStatus", Method: "POST", Args: "SiteId=4"},
		"CrestedButte": Source{Name: "Crested Butte", Provider: "snowcom", Url: "https://cache.snow.com/api/TerrainApi/GetTerrainStatus", Method: "POST", Args: "SiteId=15"},
		"Dev":          Source{Name: "Dev", Provider: "snowcom", Url: "http://127.0.0.1:24358/dev.json"},
	}
	return sources
}

func getSource(name string) (Source, error) {
	sources := getSources()
	if source, ok := sources[name]; ok {
		return source, nil
	}
	names := []string{}
	for name, _ := range sources {
		names = append(names, name)
	}
	return Source{}, fmt.Errorf("Unsupported resort, try one of: %s", strings.Join(names, ", "))
}
