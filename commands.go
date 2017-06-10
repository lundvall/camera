package camera

const (
	id      = 1
	version = "1.0"
)

// Commands without parameters
var (
	GetShootMode                      = NewCommand("getShootMode")
	GetSupportedShootMode             = NewCommand("getSupportedShootMode")
	GetAvailableShootMode             = NewCommand("getAvailableShootMode")
	ActTakePicture                    = NewCommand("actTakePicture")
	AwaitTakePicture                  = NewCommand("awaitTakePicture")
	StartContShooting                 = NewCommand("startContShooting")
	StopContShooting                  = NewCommand("stopContShooting")
	StartMovieRec                     = NewCommand("startMovieRec")
	StopMovieRec                      = NewCommand("stopMovieRec")
	StartAudioRec                     = NewCommand("startAudioRec")
	StopAudioRec                      = NewCommand("stopAudioRec")
	StartIntervalStillRec             = NewCommand("startIntervalStillRec")
	StopIntervalStillRec              = NewCommand("stopIntervalStillRec")
	StartLoopRec                      = NewCommand("startLoopRec")
	StopLoopRec                       = NewCommand("stopLoopRec")
	StartLiveview                     = NewCommand("startLiveview")
	StopLiveview                      = NewCommand("stopLiveview")
	GetLiveviewSize                   = NewCommand("getLiveviewSize")
	GetSupportedLiveviewSize          = NewCommand("getSupportedLiveviewSize")
	GetAvailableLiveviewSize          = NewCommand("getAvailableLiveviewSize")
	GetLiveviewFrameInfo              = NewCommand("getLiveviewFrameInfo")
	GetZoomSetting                    = NewCommand("getZoomSetting")
	GetSupportedZoomSetting           = NewCommand("getSupportedZoomSetting")
	GetAvailableZoomSetting           = NewCommand("getAvailableZoomSetting")
	ActHalfPressShutter               = NewCommand("actHalfPressShutter")
	CancelHalfPressShutter            = NewCommand("cancelHalfPressShutter")
	GetTouchAFPosition                = NewCommand("getTouchAFPosition")
	CancelTouchAFPosition             = NewCommand("cancelTouchAFPosition")
	CancelTrackingFocus               = NewCommand("cancelTrackingFocus")
	GetTrackingFocus                  = NewCommand("getTrackingFocus")
	GetSupportedTrackingFocus         = NewCommand("getSupportedTrackingFocus")
	GetAvailableTrackingFocus         = NewCommand("getAvailableTrackingFocus")
	GetContShootingMode               = NewCommand("getContShootingMode")
	GetSupportedContShootingMode      = NewCommand("getSupportedContShootingMode")
	GetAvailableContShootingMode      = NewCommand("getAvailableContShootingMode")
	GetContShootingSpeed              = NewCommand("getContShootingSpeed")
	GetSupportedContShootingSpeed     = NewCommand("getSupportedContShootingSpeed")
	GetAvailableContShootingSpeed     = NewCommand("getAvailableContShootingSpeed")
	GetSelfTimer                      = NewCommand("getSelfTimer")
	GetSupportedSelfTimer             = NewCommand("getSupportedSelfTimer")
	GetAvailableSelfTimer             = NewCommand("getAvailableSelfTimer")
	GetExposureMode                   = NewCommand("getExposureMode")
	GetSupportedExposureMode          = NewCommand("getSupportedExposureMode")
	GetAvailableExposureMode          = NewCommand("getAvailableExposureMode")
	GetFocusMode                      = NewCommand("getFocusMode")
	GetSupportedFocusMode             = NewCommand("getSupportedFocusMode")
	GetAvailableFocusMode             = NewCommand("getAvailableFocusMode")
	GetExposureCompensation           = NewCommand("getExposureCompensation")
	GetSupportedExposureCompensation  = NewCommand("getSupportedExposureCompensation")
	GetAvailableExposureCompensation  = NewCommand("getAvailableExposureCompensation")
	GetFNumber                        = NewCommand("getFNumber")
	GetSupportedFNumber               = NewCommand("getSupportedFNumber")
	GetAvailableFNumber               = NewCommand("getAvailableFNumber")
	GetShutterSpeed                   = NewCommand("getShutterSpeed")
	GetSupportedShutterSpeed          = NewCommand("getSupportedShutterSpeed")
	GetAvailableShutterSpeed          = NewCommand("getAvailableShutterSpeed")
	GetIsoSpeedRate                   = NewCommand("getIsoSpeedRate")
	GetSupportedIsoSpeedRate          = NewCommand("getSupportedIsoSpeedRate")
	GetAvailableIsoSpeedRate          = NewCommand("getAvailableIsoSpeedRate")
	GetWhiteBalance                   = NewCommand("getWhiteBalance")
	GetSupportedWhiteBalance          = NewCommand("getSupportedWhiteBalance")
	GetAvailableWhiteBalance          = NewCommand("getAvailableWhiteBalance")
	ActWhiteBalanceOnePushCustom      = NewCommand("actWhiteBalanceOnePushCustom")
	GetSupportedProgramShift          = NewCommand("getSupportedProgramShift")
	GetFlashMode                      = NewCommand("getFlashMode")
	GetSupportedFlashMode             = NewCommand("getSupportedFlashMode")
	GetAvailableFlashMode             = NewCommand("getAvailableFlashMode")
	GetStillSize                      = NewCommand("getStillSize")
	GetSupportedStillSize             = NewCommand("getSupportedStillSize")
	GetAvailableStillSize             = NewCommand("getAvailableStillSize")
	GetStillQuality                   = NewCommand("getStillQuality")
	GetSupportedStillQuality          = NewCommand("getSupportedStillQuality")
	GetAvailableStillQuality          = NewCommand("getAvailableStillQuality")
	GetPostviewImageSize              = NewCommand("getPostviewImageSize")
	GetSupportedPostviewImageSize     = NewCommand("getSupportedPostviewImageSize")
	GetAvailablePostviewImageSize     = NewCommand("getAvailablePostviewImageSize")
	GetMovieFileFormat                = NewCommand("getMovieFileFormat")
	GetSupportedMovieFileFormat       = NewCommand("getSupportedMovieFileFormat")
	GetAvailableMovieFileFormat       = NewCommand("getAvailableMovieFileFormat")
	GetMovieQuality                   = NewCommand("getMovieQuality")
	GetSupportedMovieQuality          = NewCommand("getSupportedMovieQuality")
	GetAvailableMovieQuality          = NewCommand("getAvailableMovieQuality")
	GetSteadyMode                     = NewCommand("getSteadyMode")
	GetSupportedSteadyMode            = NewCommand("getSupportedSteadyMode")
	GetAvailableSteadyMode            = NewCommand("getAvailableSteadyMode")
	GetViewAngle                      = NewCommand("getViewAngle")
	GetSupportedViewAngle             = NewCommand("getSupportedViewAngle")
	GetAvailableViewAngle             = NewCommand("getAvailableViewAngle")
	GetSceneSelection                 = NewCommand("getSceneSelection")
	GetSupportedSceneSelection        = NewCommand("getSupportedSceneSelection")
	GetAvailableSceneSelection        = NewCommand("getAvailableSceneSelection")
	GetColorSetting                   = NewCommand("getColorSetting")
	GetSupportedColorSetting          = NewCommand("getSupportedColorSetting")
	GetAvailableColorSetting          = NewCommand("getAvailableColorSetting")
	GetIntervalTime                   = NewCommand("getIntervalTime")
	GetSupportedIntervalTime          = NewCommand("getSupportedIntervalTime")
	GetAvailableIntervalTime          = NewCommand("getAvailableIntervalTime")
	GetLoopRecTime                    = NewCommand("getLoopRecTime")
	GetSupportedLoopRecTime           = NewCommand("getSupportedLoopRecTime")
	GetAvailableLoopRecTime           = NewCommand("getAvailableLoopRecTime")
	GetWindNoiseReduction             = NewCommand("getWindNoiseReduction")
	GetSupportedWindNoiseReduction    = NewCommand("getSupportedWindNoiseReduction")
	GetAvailableWindNoiseReduction    = NewCommand("getAvailableWindNoiseReduction")
	GetAudioRecording                 = NewCommand("getAudioRecording")
	GetSupportedAudioRecording        = NewCommand("getSupportedAudioRecording")
	GetAvailableAudioRecording        = NewCommand("getAvailableAudioRecording")
	GetFlipSetting                    = NewCommand("getFlipSetting")
	GetSupportedFlipSetting           = NewCommand("getSupportedFlipSetting")
	GetAvailableFlipSetting           = NewCommand("getAvailableFlipSetting")
	GetTvColorSystem                  = NewCommand("getTvColorSystem")
	GetSupportedTvColorSystem         = NewCommand("getSupportedTvColorSystem")
	GetAvailableTvColorSystem         = NewCommand("getAvailableTvColorSystem")
	StartRecMode                      = NewCommand("startRecMode")
	StopRecMode                       = NewCommand("stopRecMode")
	GetCameraFunction                 = NewCommand("getCameraFunction")
	GetSupportedCameraFunction        = NewCommand("getSupportedCameraFunction")
	GetAvailableCameraFunction        = NewCommand("getAvailableCameraFunction")
	GetInfraredRemoteControl          = NewCommand("getInfraredRemoteControl")
	GetSupportedInfraredRemoteControl = NewCommand("getSupportedInfraredRemoteControl")
	GetAvailableInfraredRemoteControl = NewCommand("getAvailableInfraredRemoteControl")
	GetAutoPowerOff                   = NewCommand("getAutoPowerOff")
	GetSupportedAutoPowerOff          = NewCommand("getSupportedAutoPowerOff")
	GetAvailableAutoPowerOff          = NewCommand("getAvailableAutoPowerOff")
	GetBeepMode                       = NewCommand("getBeepMode")
	GetSupportedBeepMode              = NewCommand("getSupportedBeepMode")
	GetAvailableBeepMode              = NewCommand("getAvailableBeepMode")
	GetStorageInformation             = NewCommand("getStorageInformation")
	GetAvailableAPIList               = NewCommand("getAvailableApiList")
	GetApplicationInfo                = NewCommand("getApplicationInfo")
	GetVersions                       = NewCommand("getVersions")
	GetMethodTypes                    = NewCommand("getMethodTypes")
)

// Commands with parameters
var (
	GetEvent                 = NewParamsCommand("getEvent", Params{true})
	SetShootMode             = NewParamsCommand("setShootMode", Params{"still"})
	StartLiveviewWithSize    = NewParamsCommand("startLiveviewWithSize", Params{"M"})
	SetLiveviewFrameInfo     = NewParamsCommand("setLiveviewFrameInfo", LiveviewFrameInfo{true})
	ActZoom                  = NewParamsCommand("actZoom", Params{"in", "start"})
	SetZoomSetting           = NewParamsCommand("setZoomSetting", ZoomSetting{"Optical Zoom Only"})
	SetTouchAFPosition       = NewParamsCommand("setTouchAFPosition", Params{23.4, 45.6})
	ActTrackingFocus         = NewParamsCommand("actTrackingFocus", TrackingFocusPosition{23.4, 45.6})
	SetTrackingFocus         = NewParamsCommand("setTrackingFocus", TrackingFocusState{"On"})
	SetContShootingMode      = NewParamsCommand("setContShootingMode", ContShootingMode{"Spd Priority Cont."})
	SetContShootingSpeed     = NewParamsCommand("setContShootingSpeed", ContShootingSpeed{"Hi"})
	SetSelfTimer             = NewParamsCommand("setSelfTimer", Params{2})
	SetExposureMode          = NewParamsCommand("setExposureMode", Params{"Intelligent Auto"})
	SetFocusMode             = NewParamsCommand("setFocusMode", Params{"MF"})
	SetExposureCompensation  = NewParamsCommand("setExposureCompensation", Params{2})
	SetFNumber               = NewParamsCommand("setFNumber", Params{"5.4"})
	SetShutterSpeed          = NewParamsCommand("setShutterSpeed", Params{"1/2"})
	SetIsoSpeedRate          = NewParamsCommand("setIsoSpeedRate", Params{"400"})
	SetWhiteBalance          = NewParamsCommand("setWhiteBalance", Params{"Color Temperature", true, 2500})
	SetProgramShift          = NewParamsCommand("setProgramShift", Params{1})
	SetFlashMode             = NewParamsCommand("setFlashMode", Params{"off"})
	SetStillSize             = NewParamsCommand("setStillSize", Params{"4:3", "5M"})
	SetStillQuality          = NewParamsCommand("setStillQuality", StillQuality{"RAW+JPEG"})
	SetPostviewImageSize     = NewParamsCommand("setPostviewImageSize", Params{"Original"})
	SetMovieFileFormat       = NewParamsCommand("SetMovieFileFormat", MovieFileFormat{"XAVC S"})
	SetMovieQuality          = NewParamsCommand("setMovieQuality", Params{"HQ"})
	SetSteadyMode            = NewParamsCommand("setSteadyMode", Params{"off"})
	SetViewAngle             = NewParamsCommand("setViewAngle", Params{120})
	SetSceneSelection        = NewParamsCommand("setSceneSelection", SceneSelection{"Normal"})
	SetColorSetting          = NewParamsCommand("setColorSetting", ColorSetting{"Vivid"})
	SetIntervalTime          = NewParamsCommand("setIntervalTime", IntervalTimeSec{"10"})
	SetLoopRecTime           = NewParamsCommand("setLoopRecTime", LoopRecTimeMin{"5"})
	SetWindNoiseReduction    = NewParamsCommand("setWindNoiseReduction", WindNoiseReduction{"On"})
	SetAudioRecording        = NewParamsCommand("setAudioRecording", AudioRecording{"On"})
	SetFlipSetting           = NewParamsCommand("setFlipSetting", Flip{"On"})
	SetTvColorSystem         = NewParamsCommand("setTvColorSystem", TvColorSystem{"NTSC"})
	SetCameraFunction        = NewParamsCommand("setCameraFunction", Params{"Remote Shooting"})
	SetInfraredRemoteControl = NewParamsCommand("setInfraredRemoteControl", InfraredRemoteControl{"Off"})
	SetAutoPowerOff          = NewParamsCommand("setAutoPowerOff", AutoPowerOff{60})
	SetBeepMode              = NewParamsCommand("setBeepMode", Params{"Off"})
	SetCurrentTime           = NewParamsCommand("setCurrentTime", CurrentTime{"2015-21-10T07:28:00Z", 540, 0})
)

// Params ...
type Params []interface{}

// Command to execute on camera
type Command struct {
	ID      int         `json:"id"`
	Version string      `json:"version"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

// NewCommand ...
func NewCommand(method string) Command {
	params := make(Params, 0)
	return NewParamsCommand(method, params)
}

// NewParamsCommand ...
func NewParamsCommand(method string, params interface{}) Command {
	return Command{id, version, method, params}
}

// LiveviewFrameInfo for setLiveviewFrameInfo
type LiveviewFrameInfo struct {
	FrameInfo bool `json:"frameInfo"`
}

// ZoomSetting for setZoomSetting
type ZoomSetting struct {
	Zoom string `json:"zoom"`
}

// TrackingFocusPosition for actTrackingFocus
type TrackingFocusPosition struct {
	XPosition float64 `json:"xPosition"`
	YPosition float64 `json:"yPosition"`
}

// TrackingFocusState  for setTrackingFocus
type TrackingFocusState struct {
	TrackingFocus string `json:"trackingFocus"`
}

// ContShootingMode for setContShootingMode
type ContShootingMode struct {
	ContShootingMode string `json:"contShootingMode"`
}

// ContShootingSpeed for contShootingSpeed
type ContShootingSpeed struct {
	ContShootingSpeed string `json:"contShootingSpeed"`
}

// StillQuality for setStillQuality
type StillQuality struct {
	StillQuality string `json:"stillQuality"`
}

// MovieFileFormat for setMovieFileFormat
type MovieFileFormat struct {
	MovieFileFormat string `json:"movieFileFormat"`
}

// SceneSelection for setSceneSelection
type SceneSelection struct {
	Scene string `json:"scene"`
}

// ColorSetting for setColorSetting
type ColorSetting struct {
	ColorSetting string `json:"colorSetting"`
}

// IntervalTimeSec for setIntervalTime
type IntervalTimeSec struct {
	IntervalTimeSec string `json:"intervalTimeSec"`
}

// LoopRecTimeMin for setLoopRecTime
type LoopRecTimeMin struct {
	LoopRecTimeMin string `json:"loopRecTimeMin"`
}

// WindNoiseReduction for setWindNoiseReduction
type WindNoiseReduction struct {
	WindNoiseReduction string `json:"windNoiseReduction"`
}

// AudioRecording for setAudioRecording
type AudioRecording struct {
	AudioRecording string `json:"audioRecording"`
}

// Flip for setFlipSetting
type Flip struct {
	Flip string `json:"flip"`
}

// TvColorSystem for setTvColorSystem
type TvColorSystem struct {
	TvColorSystem string `json:"tvColorSystem"`
}

// InfraredRemoteControl for setInfraredRemoteControl
type InfraredRemoteControl struct {
	InfraredRemoteControl string `json:"infraredRemoteControl"`
}

// AutoPowerOff for setAutoPowerOff
type AutoPowerOff struct {
	AutoPowerOff int `json:"autoPowerOff"`
}

// CurrentTime for setCurrentTime
type CurrentTime struct {
	DateTime             string `json:"dateTime"`
	TimeZoneOffsetMinute int    `json:"timeZoneOffsetMinute"`
	DSTOffsetMinute      int    `json:"dstOffsetMinute"`
}
