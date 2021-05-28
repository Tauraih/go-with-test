package main

import ( 
	"testing"
)

func TestPerimeter(t *testing.T){
	rectangle := Rectangle{12.0, 6.0}
	got := Perimeter(rectangle)
	want := 36.0
	
	if got != want{
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T){
	// checkArea := func(t testing.TB, shape Shape, want float64){
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want{
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T){
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	checkArea(t, rectangle, 72.0)
	// })
	// t.Run("circle", func(t *testing.T){
	// 	circle := Circle{10}
	// 	// got := circle.Area()
	// 	// want := 314.1592653589793

	// 	// if got != want{
	// 	// 	t.Errorf("got %.2f want %.2f", got, want)
	// 	// }
	// 	checkArea(t, circle, 314.1592653589793)
	// })
	areaTests := []struct{
		shape Shape
		want float64
	}{
		{Rectangle{Width: 12, Height: 6}, 72.0},
		{Circle{Radius: 10}, 314.1592653589793},
		{Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
