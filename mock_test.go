package main 

import "testing"


func TestAdds(t *testing.T) {
	sum := Add(10, 10)
	exp := 20

	if sum!= exp {
		t.Errorf("expected '%d' but got '%d'", exp, sum)
	}
}
