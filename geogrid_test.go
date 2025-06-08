package geogrid

import "testing"

func TestEncode(t *testing.T) {
	tests := []struct {
		name      string
		lat       float64
		lon       float64
		precision int
		want      string
	}{
		{
			name:      "Beijing coordinates with precision 6",
			lat:       39.9042,
			lon:       116.4074,
			precision: 6,
			want:      "wx4g0b",
		},
		{
			name:      "New York coordinates with precision 6",
			lat:       40.7128,
			lon:       -74.0060,
			precision: 6,
			want:      "dr5reg",
		},
		{
			name:      "Tokyo coordinates with precision 8",
			lat:       35.6762,
			lon:       139.6503,
			precision: 8,
			want:      "xn76cydh",
		},
		{
			name:      "London coordinates with precision 8",
			lat:       51.5074,
			lon:       -0.1278,
			precision: 8,
			want:      "gcpvj0du",
		},
		{
			name:      "Sydney coordinates with precision 12",
			lat:       -33.8688,
			lon:       151.2093,
			precision: 12,
			want:      "r3gx2f77bn44",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(tt.lat, tt.lon, tt.precision)
			if got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeBin(t *testing.T) {
	tests := []struct {
		name     string
		lat      float64
		latRange [3]float64
		want     [32]int
	}{
		{
			name:     "Latitude at equator",
			lat:      0.0,
			latRange: [3]float64{-90.0, 0.0, 90.0},
			want:     [32]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:     "Latitude at 45 degrees",
			lat:      45.0,
			latRange: [3]float64{-90.0, 0.0, 90.0},
			want:     [32]int{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var latBin [32]int
			got := encodeBin(tt.lat, latBin, tt.latRange)
			if got != tt.want {
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterleave(t *testing.T) {
	tests := []struct {
		name   string
		latBin [32]int
		lonBin [32]int
		want   []int
	}{
		{
			name:   "Simple interleaving",
			latBin: [32]int{1, 0, 1, 0},
			lonBin: [32]int{0, 1, 0, 1},
			want:   make([]int, 64), // Initialize with zeros
		},
	}

	// Set up the expected interleaved pattern
	tests[0].want[0] = 0 // lon
	tests[0].want[1] = 1 // lat
	tests[0].want[2] = 1 // lon
	tests[0].want[3] = 0 // lat
	tests[0].want[4] = 0 // lon
	tests[0].want[5] = 1 // lat
	tests[0].want[6] = 1 // lon
	tests[0].want[7] = 0 // lat

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := interleave(tt.latBin, tt.lonBin)
			if len(got) != len(tt.want) {
				t.Errorf("interleave() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("interleave()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestGenerateGeohash(t *testing.T) {
	tests := []struct {
		name        string
		interleaved []int
		precision   int
		want        string
	}{
		{
			name:        "Simple geohash generation",
			interleaved: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			precision:   4,
			want:        "0000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateGeohash(tt.interleaved, tt.precision)
			if got != tt.want {
				t.Errorf("generateGeohash() = %v, want %v", got, tt.want)
			}
		})
	}
}
