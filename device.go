package aravis

// #cgo pkg-config: aravis-0.8
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
import "unsafe"

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
	var err error

	cbool := C.arv_device_take_control(d.device, &gerror)

	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return toBool(cbool), err
}

func (d *Device) LeaveControl() (bool, error) {
	var gerror *C.GError
	var err error

	cbool := C.arv_device_leave_control(d.device, &gerror)

	if unsafe.Pointer(gerror) != nil {
		err = errorFromGError(gerror)
	}

	return toBool(cbool), err
}

func (d *Device) SetStringFeatureValue(feature, value string) {
	cfeature := C.CString(feature)
	cvalue := C.CString(value)
	C.arv_device_set_string_feature_value(d.device, cfeature, cvalue, nil)
	C.free(unsafe.Pointer(cfeature))
	C.free(unsafe.Pointer(cvalue))
}

func (d *Device) GetStringFeatureValue(feature string) (string, error) {
	cfeature := C.CString(feature)
	cvalue, err := C.arv_device_get_string_feature_value(d.device, cfeature, nil)
	C.free(unsafe.Pointer(cfeature))
	return C.GoString(cvalue), err
}

func (d *Device) SetIntegerFeatureValue(feature string, value int64) {
	cfeature := C.CString(feature)
	cvalue := C.long(value)
	C.arv_device_set_integer_feature_value(d.device, cfeature, cvalue, nil)
	C.free(unsafe.Pointer(cfeature))
}

func (d *Device) GetIntegerFeatureValue(feature string) (int64, error) {
	cfeature := C.CString(feature)
	cvalue, err := C.arv_device_get_integer_feature_value(d.device, cfeature, nil)
	C.free(unsafe.Pointer(cfeature))
	return int64(cvalue), err
}

func (d *Device) SetFloatFeatureValue(feature string, value float64) {
	cfeature := C.CString(feature)
	cvalue := C.double(value)
	C.arv_device_set_float_feature_value(d.device, cfeature, cvalue, nil)
	C.free(unsafe.Pointer(cfeature))
}

func (d *Device) GetFloatFeatureValue(feature string) (float64, error) {
	cfeature := C.CString(feature)
	cvalue, err := C.arv_device_get_float_feature_value(d.device, cfeature, nil)
	C.free(unsafe.Pointer(cfeature))
	return float64(cvalue), err
}

func (d *Device) SetNodeFeatureValue(feature, value string) {
	cfeature := C.CString(feature)
	cvalue := C.CString(value)
	C.arv_set_node_feature_value(d.device, cfeature, cvalue)
	C.free(unsafe.Pointer(cfeature))
	C.free(unsafe.Pointer(cvalue))
}
