package main

type mapping struct {
	SourceRangeStart      int64
	SourceRangeEnd        int64
	DestinationRangeStart int64
	RangeLength           int64
}

func (m *mapping) lookup(n int64) (int64, bool) {
	if n >= m.SourceRangeStart && n < m.SourceRangeEnd {
		distance := n - m.SourceRangeStart
		return m.DestinationRangeStart + distance, true
	}
	return 0, false
}

type almanac struct {
	seeds []int64
	maps  map[string][]mapping
}

type seedInfo struct {
	Seed        int64
	Soil        int64
	Fertilizer  int64
	Water       int64
	Light       int64
	Temperature int64
	Humidity    int64
	Location    int64
}

type seedIter struct {
	current int64
	from    int64
	to      int64
}

func createSeedIter(from int64, to int64) *seedIter {
	return &seedIter{
		from:    from,
		to:      to,
		current: -1,
	}
}

func (i *seedIter) Current() int64 {
	return i.current
}

func (i *seedIter) Next() bool {
	if i.current == -1 {
		i.current = i.from
		return true
	}
	if i.current >= i.from && i.current < i.to {
		i.current++
		return true
	}
	return false
}

func (a *almanac) getSeedRanges() [][]int64 {
	ranges := make([][]int64, 0)
	for i := 0; i < len(a.seeds); i += 2 {
		from := a.seeds[i]
		to := from + a.seeds[i+1]
		ranges = append(ranges, []int64{from, to})
	}
	return ranges
}

func (a *almanac) resolveSeedInfo(seed int64) seedInfo {
	resolve := func(mapName string, defaultValue int64) int64 {
		mappings := a.maps[mapName]
		for _, m := range mappings {
			n, ok := m.lookup(defaultValue)
			if ok {
				return n
			}
		}
		return defaultValue
	}
	soil := resolve("seed-to-soil map", seed)
	fertilizer := resolve("soil-to-fertilizer map", soil)
	water := resolve("fertilizer-to-water map", fertilizer)
	light := resolve("water-to-light map", water)
	temperature := resolve("light-to-temperature map", light)
	humidity := resolve("temperature-to-humidity map", temperature)
	location := resolve("humidity-to-location map", humidity)

	return seedInfo{
		Seed:        seed,
		Soil:        soil,
		Fertilizer:  fertilizer,
		Water:       water,
		Light:       light,
		Temperature: temperature,
		Humidity:    humidity,
		Location:    location,
	}
}
