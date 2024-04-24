package main

import "fmt"

func greedyAlgo(needToCover map[string]bool, stations map[string]map[string]bool) map[string]bool {
	// результат:
	result := make(map[string]bool)

	// пока есть непокрытые города
	for len(needToCover) > 0 {
		// лучшаяя станция - станция покрывающая самое большое колво городов которые еще не покрыты
		bestStation := ""

		// какие покрыты этой станцией
		citiesCovered := map[string]bool{}

		// итерируемся по станциям
		for station, stationCoverage := range stations {
			// какие станции пересекаются с теме что остались
			coveredByStation := intersectCoverage(needToCover, stationCoverage)

			//если длина покрытой станцией - больше чем предыдущая
			if len(coveredByStation) > len(citiesCovered) {
				// это лучшая станция
				bestStation = station
				// ее пересечение с нужными - лучшее
				citiesCovered = coveredByStation
			}
		}

		// удаляем из списка на покрытие найденое лучшее покрытие станцией
		for city, _ := range citiesCovered {
			delete(needToCover, city)
		}

		// вписываем станцию в результат
		result[bestStation] = true
	}

	return result
}

func intersectCoverage(neededCities, stationsCities map[string]bool) map[string]bool {
	res := make(map[string]bool)

	for station, _ := range neededCities {
		if _, ok := stationsCities[station]; ok {
			res[station] = true
		}
	}

	return res
}

func main() {
	needToCover := map[string]bool{
		"nsk":   true,
		"msk":   true,
		"piter": true,
		"barn":  true,
		"ord":   true,
		"tms":   true,
		"krs":   true,
		"vlg":   true,
		"sochi": true,
		"samar": true,
	}

	stations := map[string]map[string]bool{}

	stations["fm1"] = map[string]bool{"nsk": true, "piter": true}
	stations["fm2"] = map[string]bool{"vlg": true, "sochi": true, "barn": true}
	stations["fm3"] = map[string]bool{"tms": true, "nsk": true}
	stations["fm4"] = map[string]bool{"ord": true, "krs": true, "msk": true, "samar": true}
	stations["fm5"] = map[string]bool{"tms": true, "vlg": true}
	stations["fm7"] = map[string]bool{"ord": true, "tms": true}
	stations["fm8"] = map[string]bool{"msk": true}
	stations["fm9"] = map[string]bool{"barn": true, "samar": true}

	res := greedyAlgo(needToCover, stations)
	for i, val := range res {
		if val {
			fmt.Print(i)
			fmt.Print(" ")
		}
	}
}
