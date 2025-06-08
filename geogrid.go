package geogrid

import (
	"strings"
)

const (
	base32       = "0123456789bcdefghjkmnpqrstuvwxyz"
	latitudeMin  = -90.0
	latitudeMax  = 90.0
	longitudeMin = -180.0
	longitudeMax = 180.0
)

func encodeBin(lat float64, latBin [32]int, latRange [3]float64) [32]int {
	for i := range 32 {
		if lat >= latRange[1] {
			latBin[i] = 1
			latRange[0] = latRange[1]
		} else {
			latBin[i] = 0
			latRange[2] = latRange[1]
		}
		latRange[1] = (latRange[0] + latRange[2]) / 2
	}
	return latBin
}

func interleave(latBin, lonBin [32]int) []int {
	interleaved := make([]int, 64)
	for i := range 32 {
		interleaved[i*2] = lonBin[i]
		interleaved[i*2+1] = latBin[i]
	}
	return interleaved
}

func generateGeohash(interleaved []int, precision int) string {
	geohash := strings.Builder{}
	for i := range precision {
		value := 0
		for j := range 5 {
			bit := interleaved[i*5+j]
			value = (value << 1) | bit
		}
		geohash.WriteByte(base32[value])
	}
	return geohash.String()
}

func Encode(lat, lon float64, precision int) string {

	latRange := [3]float64{latitudeMin, 0, latitudeMax}
	lonRange := [3]float64{longitudeMin, 0, longitudeMax}
	geohash := ""
	latBin := [32]int{}
	lonBin := [32]int{}

	latBin = encodeBin(lat, latBin, latRange)
	lonBin = encodeBin(lon, lonBin, lonRange)


	// Interleave latitude and longitude bits
	interleaved := interleave(latBin, lonBin)
	geohash = generateGeohash(interleaved, precision)
	return geohash
}