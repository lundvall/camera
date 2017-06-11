package camera

const (
	id      = 1
	version = "1.0"
)

// Commands without parameters
var (
	GetShootMode                      = NewRequest("getShootMode")
	GetSupportedShootMode             = NewRequest("getSupportedShootMode")
	GetAvailableShootMode             = NewRequest("getAvailableShootMode")
	ActTakePicture                    = NewRequest("actTakePicture")
	AwaitTakePicture                  = NewRequest("awaitTakePicture")
	StartContShooting                 = NewRequest("startContShooting")
	StopContShooting                  = NewRequest("stopContShooting")
	StartMovieRec                     = NewRequest("startMovieRec")
	StopMovieRec                      = NewRequest("stopMovieRec")
	StartAudioRec                     = NewRequest("startAudioRec")
	StopAudioRec                      = NewRequest("stopAudioRec")
	StartIntervalStillRec             = NewRequest("startIntervalStillRec")
	StopIntervalStillRec              = NewRequest("stopIntervalStillRec")
	StartLoopRec                      = NewRequest("startLoopRec")
	StopLoopRec                       = NewRequest("stopLoopRec")
	StartLiveview                     = NewRequest("startLiveview")
	StopLiveview                      = NewRequest("stopLiveview")
	GetLiveviewSize                   = NewRequest("getLiveviewSize")
	GetSupportedLiveviewSize          = NewRequest("getSupportedLiveviewSize")
	GetAvailableLiveviewSize          = NewRequest("getAvailableLiveviewSize")
	GetLiveviewFrameInfo              = NewRequest("getLiveviewFrameInfo")
	GetZoomSetting                    = NewRequest("getZoomSetting")
	GetSupportedZoomSetting           = NewRequest("getSupportedZoomSetting")
	GetAvailableZoomSetting           = NewRequest("getAvailableZoomSetting")
	ActHalfPressShutter               = NewRequest("actHalfPressShutter")
	CancelHalfPressShutter            = NewRequest("cancelHalfPressShutter")
	GetTouchAFPosition                = NewRequest("getTouchAFPosition")
	CancelTouchAFPosition             = NewRequest("cancelTouchAFPosition")
	CancelTrackingFocus               = NewRequest("cancelTrackingFocus")
	GetTrackingFocus                  = NewRequest("getTrackingFocus")
	GetSupportedTrackingFocus         = NewRequest("getSupportedTrackingFocus")
	GetAvailableTrackingFocus         = NewRequest("getAvailableTrackingFocus")
	GetContShootingMode               = NewRequest("getContShootingMode")
	GetSupportedContShootingMode      = NewRequest("getSupportedContShootingMode")
	GetAvailableContShootingMode      = NewRequest("getAvailableContShootingMode")
	GetContShootingSpeed              = NewRequest("getContShootingSpeed")
	GetSupportedContShootingSpeed     = NewRequest("getSupportedContShootingSpeed")
	GetAvailableContShootingSpeed     = NewRequest("getAvailableContShootingSpeed")
	GetSelfTimer                      = NewRequest("getSelfTimer")
	GetSupportedSelfTimer             = NewRequest("getSupportedSelfTimer")
	GetAvailableSelfTimer             = NewRequest("getAvailableSelfTimer")
	GetExposureMode                   = NewRequest("getExposureMode")
	GetSupportedExposureMode          = NewRequest("getSupportedExposureMode")
	GetAvailableExposureMode          = NewRequest("getAvailableExposureMode")
	GetFocusMode                      = NewRequest("getFocusMode")
	GetSupportedFocusMode             = NewRequest("getSupportedFocusMode")
	GetAvailableFocusMode             = NewRequest("getAvailableFocusMode")
	GetExposureCompensation           = NewRequest("getExposureCompensation")
	GetSupportedExposureCompensation  = NewRequest("getSupportedExposureCompensation")
	GetAvailableExposureCompensation  = NewRequest("getAvailableExposureCompensation")
	GetFNumber                        = NewRequest("getFNumber")
	GetSupportedFNumber               = NewRequest("getSupportedFNumber")
	GetAvailableFNumber               = NewRequest("getAvailableFNumber")
	GetShutterSpeed                   = NewRequest("getShutterSpeed")
	GetSupportedShutterSpeed          = NewRequest("getSupportedShutterSpeed")
	GetAvailableShutterSpeed          = NewRequest("getAvailableShutterSpeed")
	GetIsoSpeedRate                   = NewRequest("getIsoSpeedRate")
	GetSupportedIsoSpeedRate          = NewRequest("getSupportedIsoSpeedRate")
	GetAvailableIsoSpeedRate          = NewRequest("getAvailableIsoSpeedRate")
	GetWhiteBalance                   = NewRequest("getWhiteBalance")
	GetSupportedWhiteBalance          = NewRequest("getSupportedWhiteBalance")
	GetAvailableWhiteBalance          = NewRequest("getAvailableWhiteBalance")
	ActWhiteBalanceOnePushCustom      = NewRequest("actWhiteBalanceOnePushCustom")
	GetSupportedProgramShift          = NewRequest("getSupportedProgramShift")
	GetFlashMode                      = NewRequest("getFlashMode")
	GetSupportedFlashMode             = NewRequest("getSupportedFlashMode")
	GetAvailableFlashMode             = NewRequest("getAvailableFlashMode")
	GetStillSize                      = NewRequest("getStillSize")
	GetSupportedStillSize             = NewRequest("getSupportedStillSize")
	GetAvailableStillSize             = NewRequest("getAvailableStillSize")
	GetStillQuality                   = NewRequest("getStillQuality")
	GetSupportedStillQuality          = NewRequest("getSupportedStillQuality")
	GetAvailableStillQuality          = NewRequest("getAvailableStillQuality")
	GetPostviewImageSize              = NewRequest("getPostviewImageSize")
	GetSupportedPostviewImageSize     = NewRequest("getSupportedPostviewImageSize")
	GetAvailablePostviewImageSize     = NewRequest("getAvailablePostviewImageSize")
	GetMovieFileFormat                = NewRequest("getMovieFileFormat")
	GetSupportedMovieFileFormat       = NewRequest("getSupportedMovieFileFormat")
	GetAvailableMovieFileFormat       = NewRequest("getAvailableMovieFileFormat")
	GetMovieQuality                   = NewRequest("getMovieQuality")
	GetSupportedMovieQuality          = NewRequest("getSupportedMovieQuality")
	GetAvailableMovieQuality          = NewRequest("getAvailableMovieQuality")
	GetSteadyMode                     = NewRequest("getSteadyMode")
	GetSupportedSteadyMode            = NewRequest("getSupportedSteadyMode")
	GetAvailableSteadyMode            = NewRequest("getAvailableSteadyMode")
	GetViewAngle                      = NewRequest("getViewAngle")
	GetSupportedViewAngle             = NewRequest("getSupportedViewAngle")
	GetAvailableViewAngle             = NewRequest("getAvailableViewAngle")
	GetSceneSelection                 = NewRequest("getSceneSelection")
	GetSupportedSceneSelection        = NewRequest("getSupportedSceneSelection")
	GetAvailableSceneSelection        = NewRequest("getAvailableSceneSelection")
	GetColorSetting                   = NewRequest("getColorSetting")
	GetSupportedColorSetting          = NewRequest("getSupportedColorSetting")
	GetAvailableColorSetting          = NewRequest("getAvailableColorSetting")
	GetIntervalTime                   = NewRequest("getIntervalTime")
	GetSupportedIntervalTime          = NewRequest("getSupportedIntervalTime")
	GetAvailableIntervalTime          = NewRequest("getAvailableIntervalTime")
	GetLoopRecTime                    = NewRequest("getLoopRecTime")
	GetSupportedLoopRecTime           = NewRequest("getSupportedLoopRecTime")
	GetAvailableLoopRecTime           = NewRequest("getAvailableLoopRecTime")
	GetWindNoiseReduction             = NewRequest("getWindNoiseReduction")
	GetSupportedWindNoiseReduction    = NewRequest("getSupportedWindNoiseReduction")
	GetAvailableWindNoiseReduction    = NewRequest("getAvailableWindNoiseReduction")
	GetAudioRecording                 = NewRequest("getAudioRecording")
	GetSupportedAudioRecording        = NewRequest("getSupportedAudioRecording")
	GetAvailableAudioRecording        = NewRequest("getAvailableAudioRecording")
	GetFlipSetting                    = NewRequest("getFlipSetting")
	GetSupportedFlipSetting           = NewRequest("getSupportedFlipSetting")
	GetAvailableFlipSetting           = NewRequest("getAvailableFlipSetting")
	GetTvColorSystem                  = NewRequest("getTvColorSystem")
	GetSupportedTvColorSystem         = NewRequest("getSupportedTvColorSystem")
	GetAvailableTvColorSystem         = NewRequest("getAvailableTvColorSystem")
	StartRecMode                      = NewRequest("startRecMode")
	StopRecMode                       = NewRequest("stopRecMode")
	GetCameraFunction                 = NewRequest("getCameraFunction")
	GetSupportedCameraFunction        = NewRequest("getSupportedCameraFunction")
	GetAvailableCameraFunction        = NewRequest("getAvailableCameraFunction")
	GetInfraredRemoteControl          = NewRequest("getInfraredRemoteControl")
	GetSupportedInfraredRemoteControl = NewRequest("getSupportedInfraredRemoteControl")
	GetAvailableInfraredRemoteControl = NewRequest("getAvailableInfraredRemoteControl")
	GetAutoPowerOff                   = NewRequest("getAutoPowerOff")
	GetSupportedAutoPowerOff          = NewRequest("getSupportedAutoPowerOff")
	GetAvailableAutoPowerOff          = NewRequest("getAvailableAutoPowerOff")
	GetBeepMode                       = NewRequest("getBeepMode")
	GetSupportedBeepMode              = NewRequest("getSupportedBeepMode")
	GetAvailableBeepMode              = NewRequest("getAvailableBeepMode")
	GetStorageInformation             = NewRequest("getStorageInformation")
	GetAvailableAPIList               = NewRequest("getAvailableApiList")
	GetApplicationInfo                = NewRequest("getApplicationInfo")
	GetVersions                       = NewRequest("getVersions")
	GetMethodTypes                    = NewRequest("getMethodTypes")
)

// Commands with parameters
var (
	GetEvent                 = NewParamsRequest("getEvent", Params{true})
	SetShootMode             = NewParamsRequest("setShootMode", Params{"still"})
	StartLiveviewWithSize    = NewParamsRequest("startLiveviewWithSize", Params{"M"})
	SetLiveviewFrameInfo     = NewParamsRequest("setLiveviewFrameInfo", LiveviewFrameInfo{true})
	ActZoom                  = NewParamsRequest("actZoom", Params{"in", "start"})
	SetZoomSetting           = NewParamsRequest("setZoomSetting", ZoomSetting{"Optical Zoom Only"})
	SetTouchAFPosition       = NewParamsRequest("setTouchAFPosition", Params{23.4, 45.6})
	ActTrackingFocus         = NewParamsRequest("actTrackingFocus", TrackingFocusPosition{23.4, 45.6})
	SetTrackingFocus         = NewParamsRequest("setTrackingFocus", TrackingFocusState{"On"})
	SetContShootingMode      = NewParamsRequest("setContShootingMode", ContShootingMode{"Spd Priority Cont."})
	SetContShootingSpeed     = NewParamsRequest("setContShootingSpeed", ContShootingSpeed{"Hi"})
	SetSelfTimer             = NewParamsRequest("setSelfTimer", Params{2})
	SetExposureMode          = NewParamsRequest("setExposureMode", Params{"Intelligent Auto"})
	SetFocusMode             = NewParamsRequest("setFocusMode", Params{"MF"})
	SetExposureCompensation  = NewParamsRequest("setExposureCompensation", Params{2})
	SetFNumber               = NewParamsRequest("setFNumber", Params{"5.4"})
	SetShutterSpeed          = NewParamsRequest("setShutterSpeed", Params{"1/2"})
	SetIsoSpeedRate          = NewParamsRequest("setIsoSpeedRate", Params{"400"})
	SetWhiteBalance          = NewParamsRequest("setWhiteBalance", Params{"Color Temperature", true, 2500})
	SetProgramShift          = NewParamsRequest("setProgramShift", Params{1})
	SetFlashMode             = NewParamsRequest("setFlashMode", Params{"off"})
	SetStillSize             = NewParamsRequest("setStillSize", Params{"4:3", "5M"})
	SetStillQuality          = NewParamsRequest("setStillQuality", StillQuality{"RAW+JPEG"})
	SetPostviewImageSize     = NewParamsRequest("setPostviewImageSize", Params{"Original"})
	SetMovieFileFormat       = NewParamsRequest("SetMovieFileFormat", MovieFileFormat{"XAVC S"})
	SetMovieQuality          = NewParamsRequest("setMovieQuality", Params{"HQ"})
	SetSteadyMode            = NewParamsRequest("setSteadyMode", Params{"off"})
	SetViewAngle             = NewParamsRequest("setViewAngle", Params{120})
	SetSceneSelection        = NewParamsRequest("setSceneSelection", SceneSelection{"Normal"})
	SetColorSetting          = NewParamsRequest("setColorSetting", ColorSetting{"Vivid"})
	SetIntervalTime          = NewParamsRequest("setIntervalTime", IntervalTimeSec{"10"})
	SetLoopRecTime           = NewParamsRequest("setLoopRecTime", LoopRecTimeMin{"5"})
	SetWindNoiseReduction    = NewParamsRequest("setWindNoiseReduction", WindNoiseReduction{"On"})
	SetAudioRecording        = NewParamsRequest("setAudioRecording", AudioRecording{"On"})
	SetFlipSetting           = NewParamsRequest("setFlipSetting", Flip{"On"})
	SetTvColorSystem         = NewParamsRequest("setTvColorSystem", TvColorSystem{"NTSC"})
	SetCameraFunction        = NewParamsRequest("setCameraFunction", Params{"Remote Shooting"})
	SetInfraredRemoteControl = NewParamsRequest("setInfraredRemoteControl", InfraredRemoteControl{"Off"})
	SetAutoPowerOff          = NewParamsRequest("setAutoPowerOff", AutoPowerOff{60})
	SetBeepMode              = NewParamsRequest("setBeepMode", Params{"Off"})
	SetCurrentTime           = NewParamsRequest("setCurrentTime", CurrentTime{"2015-21-10T07:28:00Z", 540, 0})
)

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
