
# Camera

[![GoDoc](https://godoc.org/github.com/lundvall/camera?status.svg)](https://godoc.org/github.com/lundvall/camera)

Camera Remote API for Go.


## About

Automate your photography with Go and create things like sophisticated time-lapses or user interfaces.

## Usage

To use this library with your camera, you must enable the remote control feature. Please see [developer.sony.com](https://developer.sony.com/devices/cameras/) for a full list of supported devices.

```
cam := camera.NewCamera()
cam.Execute(camera.ActTakePicture)
```

# License

MIT