package camera

import (
	"log"
)

func ExampleNewCamera() {
	cam := NewCamera()
	res, err := cam.Execute(ActTakePicture)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
