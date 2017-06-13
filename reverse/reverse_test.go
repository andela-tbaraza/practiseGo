package reverse

import "testing"

func TestReverse(t *testing.T) {
	actualResult := Reverse("golang")
	var expectedResult = "gnalog"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}

}
