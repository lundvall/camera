package camera

import (
	"testing"
	"time"
)

func TestErrorResponse(t *testing.T) {
	cam := NewCamera()

	if available("actTakePicture", &cam) {
		cam.Action(StopRecMode)
		time.Sleep(3 * time.Second)
	}

	res, err := cam.Action(ActTakePicture)
	if len(res.Result) > 0 {
		t.Fatalf("unexpected result length: got %d want %d", len(res.Result), 0)
	}

	expected := "response: 1 Not Available Now"
	if err.Error() != expected {
		t.Errorf("expected error message: got %s want %s", err.Error(), expected)
	}
}

func TestAvailableAPIs(t *testing.T) {
	cam := NewCamera()

	name := "getAvailableApiList"
	if !available(name, &cam) {
		t.Errorf("expected %s to be avaialbe", name)
	}
}

func available(name string, cam *Camera) bool {
	res, _ := cam.Action(GetAvailableAPIList)
	list := res.Result[0].([]interface{})
	for _, item := range list {
		if item == name {
			return true
		}
	}
	return false
}
