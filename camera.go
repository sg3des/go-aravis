package aravis

// #cgo pkg-config: aravis-0.10
// #include <arv.h>
// #include <stdlib.h>
// #include <stdio.h>
/*
extern void go_control_lost_handler();

static void control_lost_cb (ArvGvDevice *gv_device)
{
	go_control_lost_handler();
}

static void init_control_lost_cb(ArvCamera *camera)
{
	g_signal_connect(arv_camera_get_device(camera), "control-lost",
		G_CALLBACK (control_lost_cb), NULL);
}

static void stream_cb_rt(void *user_data, ArvStreamCallbackType type, ArvBuffer *buffer)
{
	if (type == ARV_STREAM_CALLBACK_TYPE_INIT) {
		if (!arv_make_thread_realtime (10))
			printf ("Failed to make stream thread realtime\n");
	}
}

static ArvStream* arv_camera_create_rt_stream(ArvCamera *camera, void *user_data, GError **error) {
	return arv_camera_create_stream(camera, stream_cb_rt, user_data, error);
}

static void stream_cb_hp(void *user_data, ArvStreamCallbackType type, ArvBuffer *buffer)
{
	if (type == ARV_STREAM_CALLBACK_TYPE_INIT) {
		if (!arv_make_thread_high_priority (-10))
			printf ("Failed to make stream thread high priority\n");
	}
}

static ArvStream* arv_camera_create_hp_stream(ArvCamera *camera, void *user_data, GError **error) {
	return arv_camera_create_stream(camera, stream_cb_hp, user_data, error);
}
*/
import "C"
import (
	"unsafe"
)

type ThreadPriorityType int

const (
	ThreadPriorityNormal ThreadPriorityType = iota
	ThreadPriorityRealtime
	ThreadPriorityHigh
)

type Camera struct {
	camera         *C.struct__ArvCamera
	ThreadPriority ThreadPriorityType
}

const (
	ACQUISITION_MODE_CONTINUOUS   = C.ARV_ACQUISITION_MODE_CONTINUOUS
	ACQUISITION_MODE_SINGLE_FRAME = C.ARV_ACQUISITION_MODE_SINGLE_FRAME
)

const (
	AUTO_OFF        = C.ARV_AUTO_OFF
	AUTO_ONCE       = C.ARV_AUTO_ONCE
	AUTO_CONTINUOUS = C.ARV_AUTO_CONTINUOUS
)

func NewCamera(name string) (Camera, error) {
	var c Camera
	var gerror *C.GError
	var err error

	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	c.camera = C.arv_camera_new(cs, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return c, err
}

func (c *Camera) CreateStream() (Stream, error) {
	var s Stream
	var gerror *C.GError
	var err error

	switch c.ThreadPriority {
	case ThreadPriorityRealtime:
		s.stream = C.arv_camera_create_rt_stream(
			c.camera,
			nil,
			&gerror,
		)

	case ThreadPriorityHigh:
		s.stream = C.arv_camera_create_hp_stream(
			c.camera,
			nil,
			&gerror,
		)

	default:
		s.stream = C.arv_camera_create_stream(
			c.camera,
			nil,
			nil,
			&gerror,
		)
	}

	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	if s.stream == nil {
		return Stream{}, err
	}

	C.init_control_lost_cb(c.camera)

	return s, err
}

func (c *Camera) GetDevice() (Device, error) {
	var d Device
	var err error

	d.device = C.arv_camera_get_device(c.camera)

	return d, err
}

func (c *Camera) GetVendorName() (string, error) {
	var gerror *C.GError

	name := C.arv_camera_get_vendor_name(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err := errorFromGError(gerror)
		return "", err
	}

	return C.GoString(name), nil
}

func (c *Camera) GetModelName() (string, error) {
	var gerror *C.GError

	name := C.arv_camera_get_model_name(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err := errorFromGError(gerror)
		return "", err
	}

	return C.GoString(name), nil
}

func (c *Camera) GetDeviceId() (string, error) {
	var gerror *C.GError
	var err error

	id := C.arv_camera_get_device_id(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err := errorFromGError(gerror)
		return "", err
	}

	return C.GoString(id), err
}

//
// features
//

func (c *Camera) SetStringFeature(name, value string) error {
	var gerror *C.GError

	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	csval := C.CString(value)
	defer C.free(unsafe.Pointer(cs))

	C.arv_camera_set_string(
		c.camera,
		cs,
		csval,
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		return errorFromGError(gerror)
	}

	return nil
}

func (c *Camera) GetStringFeature(name string) (string, error) {
	var gerror *C.GError

	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	val := C.arv_camera_get_string(
		c.camera,
		cs,
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		return "", errorFromGError(gerror)
	}

	return C.GoString(val), nil
}

func (c *Camera) GetIntegerFeature(name string) (int, error) {
	var gerror *C.GError

	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	val := C.arv_camera_get_integer(
		c.camera,
		cs,
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		return 0, errorFromGError(gerror)
	}

	return int(val), nil
}

func (c *Camera) GetFloatFeature(name string) (float64, error) {
	var gerror *C.GError

	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	val := C.arv_camera_get_float(
		c.camera,
		cs,
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		return 0, errorFromGError(gerror)
	}

	return float64(val), nil
}

func (c *Camera) IsFeatureAvailable(name string) (bool, error) {
	var gerror *C.GError

	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	val := C.arv_camera_is_feature_available(
		c.camera,
		cs,
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		return false, errorFromGError(gerror)
	}

	return toBool(val), nil
}

//
//
//

func (c *Camera) GetSensorSize() (w int, h int, err error) {
	var gerror *C.GError

	C.arv_camera_get_sensor_size(
		c.camera,
		(*C.gint)(unsafe.Pointer(&w)),
		(*C.gint)(unsafe.Pointer(&h)),
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return int(w), int(h), err
}

func (c *Camera) SetRegion(x, y, w, h int) error {
	var gerror *C.GError
	var err error

	C.arv_camera_set_region(c.camera,
		C.gint(x),
		C.gint(y),
		C.gint(w),
		C.gint(h),
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) GetRegion() (x int, y int, w int, h int, err error) {
	var gerror *C.GError

	C.arv_camera_get_region(
		c.camera,
		(*C.gint)(unsafe.Pointer(&x)),
		(*C.gint)(unsafe.Pointer(&y)),
		(*C.gint)(unsafe.Pointer(&w)),
		(*C.gint)(unsafe.Pointer(&h)),
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return int(x), int(y), int(w), int(h), err
}

func (c *Camera) GetHeight() (height int, err error) {
	var gerror *C.GError

	cs := C.CString("Height")
	defer C.free(unsafe.Pointer(cs))

	val := C.arv_camera_get_integer(
		c.camera,
		cs,
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return int(val), err
}

func (c *Camera) GetHeightBounds() (min int, max int, err error) {
	var gerror *C.GError

	C.arv_camera_get_height_bounds(
		c.camera,
		(*C.gint)(unsafe.Pointer(&min)),
		(*C.gint)(unsafe.Pointer(&max)),
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return int(min), int(max), err
}

func (c *Camera) GetWidth() (width int, err error) {
	var gerror *C.GError

	cs := C.CString("Width")
	defer C.free(unsafe.Pointer(cs))

	val := C.arv_camera_get_integer(
		c.camera,
		cs,
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return int(val), err
}

func (c *Camera) GetWidthBounds() (min int, max int, err error) {
	var gerror *C.GError

	C.arv_camera_get_width_bounds(
		c.camera,
		(*C.gint)(unsafe.Pointer(&min)),
		(*C.gint)(unsafe.Pointer(&max)),
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return int(min), int(max), err
}

func (c *Camera) SetBinning() {
	// TODO
}

func (c *Camera) GetBinning() (min int, max int, err error) {
	var gerror *C.GError

	C.arv_camera_get_binning(
		c.camera,
		(*C.gint)(unsafe.Pointer(&min)),
		(*C.gint)(unsafe.Pointer(&max)),
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return int(min), int(max), err
}

func (c *Camera) SetPixelFormat(pixfmt string) error {
	var gerror *C.GError

	cs := C.CString(pixfmt)
	defer C.free(unsafe.Pointer(cs))

	C.arv_camera_set_pixel_format_from_string(
		c.camera,
		cs,
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		return errorFromGError(gerror)
	}

	return nil
}

func (c *Camera) GetPixelFormat() (string, error) {
	var gerror *C.GError

	val := C.arv_camera_get_pixel_format_as_string(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		return "", errorFromGError(gerror)
	}

	return C.GoString(val), nil
}

func (c *Camera) GetAvailablePixelFormats() (list []string, err error) {
	var gerror *C.GError
	var num int

	val := C.arv_camera_dup_available_pixel_formats_as_strings(
		c.camera,
		(*C.guint)(unsafe.Pointer(&num)),
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		return nil, errorFromGError(gerror)
	}

	for _, v := range unsafe.Slice(val, num) {
		list = append(list, C.GoString(v))
	}

	return
}

func (c *Camera) StartAcquisition() error {
	var gerror *C.GError
	var err error

	C.arv_camera_start_acquisition(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) StopAcquisition() error {
	var gerror *C.GError
	var err error

	C.arv_camera_stop_acquisition(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) AbortAcquisition() error {
	var gerror *C.GError
	var err error

	C.arv_camera_abort_acquisition(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) SetAcquisitionMode(mode int) error {
	var gerror *C.GError
	var err error

	C.arv_camera_set_acquisition_mode(c.camera, C.ArvAcquisitionMode(mode), &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) SetFrameRate(frameRate float64) error {
	var gerror *C.GError
	var err error

	C.arv_camera_set_frame_rate(c.camera, C.double(frameRate), &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) GetFrameRate() (float64, error) {
	var gerror *C.GError
	var err error

	fr := C.arv_camera_get_frame_rate(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return float64(fr), err
}

func (c *Camera) GetFrameRateBounds() (float64, float64, error) {
	var gerror *C.GError
	var err error

	var min, max float64
	C.arv_camera_get_frame_rate_bounds(
		c.camera,
		(*C.double)(unsafe.Pointer(&min)),
		(*C.double)(unsafe.Pointer(&max)),
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}
	return float64(min), float64(max), err
}

func (c *Camera) SetLineRate(lineRate float64) {
	c.SetFrameRate(lineRate)
}

func (c *Camera) GetLineRate() (float64, error) {
	return c.GetFrameRate()
}

func (c *Camera) SetTrigger(source string) error {
	var gerror *C.GError
	var err error

	csource := C.CString(source)
	C.arv_camera_set_trigger(c.camera, csource, &gerror)
	C.free(unsafe.Pointer(csource))

	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) SetTriggerSource(source string) error {
	var gerror *C.GError
	var err error

	csource := C.CString(source)
	C.arv_camera_set_trigger_source(c.camera, csource, &gerror)
	C.free(unsafe.Pointer(csource))

	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) GetTriggerSource() (string, error) {
	var gerror *C.GError
	var err error

	csource := C.arv_camera_get_trigger_source(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
		return "", err
	}

	return C.GoString(csource), err
}

func (c *Camera) SoftwareTrigger() error {
	var gerror *C.GError
	var err error

	C.arv_camera_software_trigger(c.camera, &gerror)

	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) ClearTriggers() error {
	var gerror *C.GError
	var err error

	C.arv_camera_clear_triggers(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) IsExposureTimeAvailable() (bool, error) {
	var gerror *C.GError
	var err error

	gboolean := C.arv_camera_is_exposure_time_available(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return toBool(gboolean), err
}

func (c *Camera) IsExposureAutoAvailable() (bool, error) {
	var gerror *C.GError
	var err error

	gboolean := C.arv_camera_is_exposure_auto_available(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}
	return toBool(gboolean), err
}

func (c *Camera) SetExposureTime(time float64) error {
	var gerror *C.GError
	var err error

	C.arv_camera_set_exposure_time(c.camera, C.double(time), &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) GetExposureTime() (float64, error) {
	var gerror *C.GError
	var err error

	cdouble := C.arv_camera_get_exposure_time(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return float64(cdouble), err
}

func (c *Camera) GetExposureTimeBounds() {
	// TODO
}

func (c *Camera) SetExposureTimeAuto(mode int) error {
	var gerror *C.GError
	var err error

	C.arv_camera_set_exposure_time_auto(c.camera, C.ArvAuto(mode), &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) GetExposureTimeAuto() {
	// TODO
}

func (c *Camera) SetGain(gain float64) error {
	var gerror *C.GError
	var err error
	C.arv_camera_set_gain(c.camera, C.double(gain), &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) GetGain() (float64, error) {
	var gerror *C.GError
	var err error

	cgain := C.arv_camera_get_gain(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return float64(cgain), err
}

func (c *Camera) GetGainBounds() (float64, float64, error) {
	var gerror *C.GError
	var err error

	var min, max float64
	C.arv_camera_get_gain_bounds(
		c.camera,
		(*C.double)(unsafe.Pointer(&min)),
		(*C.double)(unsafe.Pointer(&max)),
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return float64(min), float64(max), err
}

func (c *Camera) SetGainAuto() {
	// TODO
}

func (c *Camera) GetPayloadSize() (uint, error) {
	var gerror *C.GError
	var err error

	csize := C.arv_camera_get_payload(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return uint(csize), err
}

func (c *Camera) IsGVDevice() (bool, error) {
	cbool := C.arv_camera_is_gv_device(c.camera)

	return toBool(cbool), nil
}

func (c *Camera) GVGetNumStreamChannels() (int, error) {
	var gerror *C.GError
	var err error

	cint := C.arv_camera_gv_get_n_stream_channels(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return int(cint), err
}

func (c *Camera) GVSelectStreamChannels(id int) error {
	var gerror *C.GError
	var err error

	C.arv_camera_gv_select_stream_channel(c.camera, C.gint(id), &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) GVGetCurrentStreamChannel() (int, error) {
	var gerror *C.GError
	var err error

	cint := C.arv_camera_gv_get_current_stream_channel(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return int(cint), err
}

func (c *Camera) GVGetPacketDelay() (int64, error) {
	var gerror *C.GError
	var err error

	cint64 := C.arv_camera_gv_get_packet_delay(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return int64(cint64), err
}

func (c *Camera) GVSetPacketDelay(delay int64) error {
	var gerror *C.GError
	var err error

	C.arv_camera_gv_set_packet_delay(c.camera, C.gint64(delay), &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}
	return err
}

func (c *Camera) GVGetPacketSize() (int, error) {
	var gerror *C.GError
	var err error

	csize := C.arv_camera_gv_get_packet_size(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return int(csize), err
}

func (c *Camera) GVSetPacketSize(size int) error {
	var gerror *C.GError
	var err error

	C.arv_camera_gv_set_packet_size(c.camera, C.gint(size), &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return err
}

func (c *Camera) GetChunkMode() (bool, error) {
	var gerror *C.GError
	var err error

	mode := C.arv_camera_get_chunk_mode(c.camera, &gerror)
	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return toBool(mode), err
}

func (c *Camera) Close() {
	C.g_object_unref(C.gpointer(c.camera))
}

var controlLostHandler func()

func (c *Camera) SetControlLostHandler(hdl func()) error {
	controlLostHandler = hdl
	return nil
}

func (c *Camera) IsNil() bool {
	return c.camera == nil
}

//export go_control_lost_handler
func go_control_lost_handler() {
	if controlLostHandler != nil {
		controlLostHandler()
	}
}
