package camera

import (
	"log"
)

func ExampleNewCamera() {
	cam := NewCamera()
	res, err := cam.Action(ActTakePicture)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
