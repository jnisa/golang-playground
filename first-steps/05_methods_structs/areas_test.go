
package areas

import "testing"

type polygon_type struct {
	width float64
	length float64
}

type circle_type struct {
	radious float64
}

func Test_get_polygon_area(t *testing.T){
	t.Run("1st Test: low complexity test", func(t *testing.T){		
		var polygon polygon_type
		polygon = polygon_type{10.0, 5.0} 

		result := get_polygon_area(polygon)
		expected := 50.0

		if result != expected {
			t.Errorf("TEST FAILED - Result: %.1f | Expected: %.1f", result, expected)
		}
	})

	t.Run("2nd Test: when zero", func(t *testing.T){
		var polygon polygon_type
		polygon = polygon_type{0.0, 0.0}

		result := get_polygon_area(polygon)
		expected := 0.0

		if result != expected {
			t.Errorf("TEST FAILED - Result: %.1f | Expected: %.1f", result, expected)
		}
	})
}


func Test_get_polygon_perimeter(t *testing.T){
	t.Run("1st Test: low complexity test", func(t *testing.T){
		var polygon polygon_type
		polygon = polygon_type{10.0, 5.0}

		result := get_polygon_perimeter(polygon)
		expected := 30.0

		if result != expected {
			t.Errorf("TEST FAILED - Result: %.1f | Expected: %.1f", result, expected)
		}
	})

	t.Run("2nd Test: when the dimensions are zero", func(t *testing.T){
		var polygon polygon_type
		polygon = polygon_type{0.0, 0.0}

		result := get_polygon_perimeter(polygon)
		expected := 0.0

		if result != expected {
			t.Errorf("TEST FAILED - Result: %.1f | Expected: %.1f", result, expected)
		}
	})
}


func Test_get_circle_area(t *testing.T){
	t.Run("1st Test: low complexity test", func(t *testing.T){
		var circle circle_type
		circle = circle_type{2.0}

		result := roundFloat(get_circle_area(circle), 3)
		expected := 12.566

		if result != expected {
			t.Errorf("TEST FAILED - Result: %.3f | Expected: %.3f", result, expected)
		}
	})

	t.Run("2nd Test: when the dimensions are zero", func(t *testing.T){
		var circle circle_type
		circle = circle_type{0.0}

		result := roundFloat(get_circle_area(circle), 3)
		expected := 0.000

		if result != expected {
			t.Errorf("TEST FAILED - Result: %.3f | Expected: %.3f", result, expected)
		}
	})
}

