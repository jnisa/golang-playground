
package areas

import "math"

const base = 10

func get_polygon_area(polygon polygon_type) float64 {
	return polygon.width * polygon.length
}

func get_polygon_perimeter(polygon polygon_type) float64 {
	return 2 * (polygon.width + polygon.length)
}

func get_circle_area(circle circle_type) float64 {
	return math.Pi * math.Pow(circle.radious, 2) 
}

func get_circle_perimeter(circle circle_type) float64 { 
	return 2 * math.Pi * circle.radious
}

func roundFloat(val float64, precision float64) float64 {
	int_part, dec_part := math.Modf(val)
	dec_part, residual := math.Modf(dec_part * math.Pow(base, precision))
	return int_part + (dec_part / math.Pow(base, precision)) + (residual * 0) 
}
