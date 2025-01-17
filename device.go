package aravis

// #cgo pkg-config: aravis-0.10
// #include <arv.h>
// #include <stdlib.h>
/*
void arv_set_node_feature_value(ArvDevice *device, char *name, char *value) {
	ArvGcNode *feature;
	feature = arv_device_get_feature (device, name);
	arv_gc_feature_node_set_value_from_string (ARV_GC_FEATURE_NODE (feature), value, NULL);
}

gboolean arv_device_take_control(ArvDevice *device, GError **error) {
	return arv_gv_device_take_control(ARV_GV_DEVICE(device), error);
}

gboolean arv_device_leave_control(ArvDevice *device, GError **error) {
	return arv_gv_device_leave_control(ARV_GV_DEVICE(device), error);
}

*/
import "C"
import (
	"unsafe"
)

const (
	DEVICE_ERROR_WRONG_FEATURE     = C.ARV_DEVICE_ERROR_WRONG_FEATURE
	DEVICE_ERROR_FEATURE_NOT_FOUND = C.ARV_DEVICE_ERROR_FEATURE_NOT_FOUND
	DEVICE_ERROR_NOT_CONNECTED     = C.ARV_DEVICE_ERROR_NOT_CONNECTED
	DEVICE_ERROR_PROTOCOL_ERROR    = C.ARV_DEVICE_ERROR_PROTOCOL_ERROR
	DEVICE_ERROR_TRANSFER_ERROR    = C.ARV_DEVICE_ERROR_TRANSFER_ERROR
	DEVICE_ERROR_TIMEOUT           = C.ARV_DEVICE_ERROR_TIMEOUT
	DEVICE_ERROR_NOT_FOUND         = C.ARV_DEVICE_ERROR_NOT_FOUND
	DEVICE_ERROR_INVALID_PARAMETER = C.ARV_DEVICE_ERROR_INVALID_PARAMETER
	DEVICE_ERROR_GENICAM_NOT_FOUND = C.ARV_DEVICE_ERROR_GENICAM_NOT_FOUND
	DEVICE_ERROR_NO_STREAM_CHANNEL = C.ARV_DEVICE_ERROR_NO_STREAM_CHANNEL
	DEVICE_ERROR_NOT_CONTROLLER    = C.ARV_DEVICE_ERROR_NOT_CONTROLLER
	DEVICE_ERROR_UNKNOWN           = C.ARV_DEVICE_ERROR_UNKNOWN
)

type Device struct {
	device *C.struct__ArvDevice
}

func (d *Device) TakeControl() (bool, error) {
	var gerror *C.GError

	cbool := C.arv_device_take_control(d.device, &gerror)
	if unsafe.Pointer(gerror) != nil {
		return false, errorFromGError(gerror)
	}

	return toBool(cbool), nil
}

func (d *Device) LeaveControl() (bool, error) {
	var gerror *C.GError

	cbool := C.arv_device_leave_control(d.device, &gerror)
	if unsafe.Pointer(gerror) != nil {
		return false, errorFromGError(gerror)
	}

	return toBool(cbool), nil
}

func (d *Device) SetFeatures(value string) error {
	var gerror *C.GError

	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))

	C.arv_device_set_features_from_string(d.device, cvalue, &gerror)
	if unsafe.Pointer(gerror) != nil {
		return errorFromGError(gerror)
	}

	return nil
}

func (d *Device) IsFeatureAvailable(name string) (bool, error) {
	var gerror *C.GError

	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	val := C.arv_device_is_feature_available(
		d.device,
		cs,
		&gerror,
	)
	if unsafe.Pointer(gerror) != nil {
		return false, errorFromGError(gerror)
	}

	return toBool(val), nil
}

func (d *Device) SetStringFeatureValue(feature, value string) error {
	var gerror *C.GError

	cfeature := C.CString(feature)
	defer C.free(unsafe.Pointer(cfeature))
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))

	C.arv_device_set_string_feature_value(d.device, cfeature, cvalue, &gerror)
	if unsafe.Pointer(gerror) != nil {
		return errorFromGError(gerror)
	}

	return nil
}

func (d *Device) GetStringFeatureValue(feature string) (string, error) {
	var gerror *C.GError

	cfeature := C.CString(feature)
	defer C.free(unsafe.Pointer(cfeature))

	cvalue := C.arv_device_get_string_feature_value(d.device, cfeature, &gerror)
	if unsafe.Pointer(gerror) != nil {
		return "", errorFromGError(gerror)
	}

	return C.GoString(cvalue), nil
}

func (d *Device) SetIntegerFeatureValue(feature string, value int64) error {
	var gerror *C.GError

	cfeature := C.CString(feature)
	defer C.free(unsafe.Pointer(cfeature))
	cvalue := C.long(value)

	C.arv_device_set_integer_feature_value(d.device, cfeature, cvalue, &gerror)
	if unsafe.Pointer(gerror) != nil {
		return errorFromGError(gerror)
	}

	return nil
}

func (d *Device) GetIntegerFeatureValue(feature string) (int64, error) {
	var gerror *C.GError

	cfeature := C.CString(feature)
	defer C.free(unsafe.Pointer(cfeature))

	cvalue := C.arv_device_get_integer_feature_value(d.device, cfeature, &gerror)
	if unsafe.Pointer(gerror) != nil {
		return 0, errorFromGError(gerror)
	}

	return int64(cvalue), nil
}

func (d *Device) SetFloatFeatureValue(feature string, value float64) error {
	var gerror *C.GError

	cfeature := C.CString(feature)
	defer C.free(unsafe.Pointer(cfeature))
	cvalue := C.double(value)

	C.arv_device_set_float_feature_value(d.device, cfeature, cvalue, &gerror)
	if unsafe.Pointer(gerror) != nil {
		return errorFromGError(gerror)
	}

	return nil
}

func (d *Device) GetFloatFeatureValue(feature string) (float64, error) {
	var gerror *C.GError

	cfeature := C.CString(feature)
	defer C.free(unsafe.Pointer(cfeature))

	cvalue := C.arv_device_get_float_feature_value(d.device, cfeature, &gerror)
	if unsafe.Pointer(gerror) != nil {
		return 0, errorFromGError(gerror)
	}

	return float64(cvalue), nil
}

func (d *Device) SetNodeFeatureValue(feature, value string) error {
	cfeature := C.CString(feature)
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cfeature))
	defer C.free(unsafe.Pointer(cvalue))

	C.arv_set_node_feature_value(d.device, cfeature, cvalue)

	return nil
}

func (d *Device) ExecuteCommand(feature string) error {
	var gerror *C.GError

	cfeature := C.CString(feature)
	defer C.free(unsafe.Pointer(cfeature))

	C.arv_device_execute_command(d.device, cfeature, &gerror)
	if unsafe.Pointer(gerror) != nil {
		return errorFromGError(gerror)
	}

	return nil
}

func (d *Device) GetGenicamXML() string {
	var size int

	cvalue := C.arv_device_get_genicam_xml(
		d.device,
		(*C.size_t)(unsafe.Pointer(&size)),
	)

	return C.GoString(cvalue)
}

func (d *Device) IsNil() bool {
	return d.device == nil
}
