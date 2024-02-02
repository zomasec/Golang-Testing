package structs_interfaces

import (
	"testify/assert"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{12.5, 5.1}

	got := Perimeter(rectangle)
	want := 127.5

	assert.Equalf(t, want, got, "got %.2f want %.2f", got, want)
}

//func TestArea(t *testing.T) {

//t.Run("rectangle", func(t *testing.T) {
//	rectangle := Rectangle{12.2, 10.6}
//
//	got := Area(rectangle)
//	want := 129.32
//
//	assert.Equalf(t, want, got, "got %.2f want %.2f", got, want)
//})
//
//t.Run("circles", func(t *testing.T) {
//	circle := Circle{10}
//
//	got := Area(circle)
//	want := 129.32
//
//	assert.Equalf(t, want, got, "got %.2f want %.2f", got, want)
//})

// }
func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assert.Equalf(t, tt.want, got, "got %g want %g", got, tt.want)
		//if got != tt.want {
		//	t.Errorf("got %g want %g", got, tt.want)
		//}
	}

}
