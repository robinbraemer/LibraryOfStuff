package main

import "fmt"
import "strconv"

type Vertex struct {
	Lat, Long float64
}

type Location struct {
	Name string
	Loc  Vertex
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(toString(m))

	var locs []Location
	for k, v := range m {
		locs = append(locs, Location{k, v})
	}

	for _, loc := range locs {
		fmt.Println(loc.toString())
	}
}

func (l Location) toString() string {
	var s string
	s += "Name: " + l.Name + "\n"
	s += "Location:\n"
	s += "  Lat: " + FloatToString(l.Loc.Lat) + "\n"
	s += "  Long: " + FloatToString(l.Loc.Long) + "\n"
	return s
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func toString(m map[string]Vertex) string {
	var s string
	for k, v := range m {
		s += "---------\n"
		s += "Name: " + k + "\n"
		s += "Location:\n"
		s += "  Lat: " + FloatToString(v.Lat) + "\n"
		s += "  Long: " + FloatToString(v.Long) + "\n"
	}
	return s
}
