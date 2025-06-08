# GeoGrid

A Go library for encoding geographic coordinates into geohashes. Geohashing is a method of encoding latitude and longitude coordinates into a short string of letters and digits.

## Installation

```bash
go get github.com/aanurraj/geogrid
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/aanurraj/geogrid"
)

func main() {
    // Encode coordinates with different precision levels
    // Beijing coordinates
    beijingHash := geogrid.Encode(39.9042, 116.4074, 6)
    fmt.Printf("Beijing geohash (precision 6): %s\n", beijingHash)
    // Output: Beijing geohash (precision 6): wx4g0b

    // New York coordinates
    nyHash := geogrid.Encode(40.7128, -74.0060, 6)
    fmt.Printf("New York geohash (precision 6): %s\n", nyHash)
    // Output: New York geohash (precision 6): dr5reg

    // Tokyo coordinates with higher precision
    tokyoHash := geogrid.Encode(35.6762, 139.6503, 8)
    fmt.Printf("Tokyo geohash (precision 8): %s\n", tokyoHash)
    // Output: Tokyo geohash (precision 8): xn76cydh
}
```

### Understanding Precision

The precision parameter determines the length of the geohash string and the size of the geographic area it represents:

- Precision 1: ±2500km
- Precision 2: ±630km
- Precision 3: ±78km
- Precision 4: ±20km
- Precision 5: ±2.4km
- Precision 6: ±610m
- Precision 7: ±76m
- Precision 8: ±19m
- Precision 9: ±2.4m
- Precision 10: ±60cm
- Precision 11: ±7.5cm
- Precision 12: ±1.9cm

### Example with Different Precisions

```go
package main

import (
    "fmt"
    "github.com/aanurraj/geogrid"
)

func main() {
    // London coordinates
    lat, lon := 51.5074, -0.1278

    // Generate geohashes with different precisions
    for precision := 1; precision <= 6; precision++ {
        hash := geogrid.Encode(lat, lon, precision)
        fmt.Printf("London geohash (precision %d): %s\n", precision, hash)
    }
}
```

Output:
```
London geohash (precision 1): g
London geohash (precision 2): gc
London geohash (precision 3): gcp
London geohash (precision 4): gcpv
London geohash (precision 5): gcpvj
London geohash (precision 6): gcpvj0
```

### Use Cases

1. **Location-based Services**
   - Store and query locations efficiently
   - Find nearby points of interest
   - Implement location-based search

2. **Spatial Indexing**
   - Create spatial indexes for databases
   - Optimize geographic queries
   - Implement geofencing

3. **Data Visualization**
   - Aggregate geographic data
   - Create heat maps
   - Implement clustering

### Best Practices

1. **Choose Appropriate Precision**
   - Use lower precision (1-6) for general area representation
   - Use higher precision (7-12) for precise location encoding
   - Consider your use case when selecting precision

2. **Error Handling**
   - Always validate input coordinates
   - Handle edge cases (coordinates outside valid ranges)
   - Consider precision limits

3. **Performance Considerations**
   - Higher precision means longer strings and more processing
   - Balance precision with your application's needs
   - Consider caching frequently used geohashes

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
