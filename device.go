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
*/
import "C"
import "unsafe"

const (
	// DEVICE_STATUS_UNKNOWN        = C.ARV_DEVICE_STATUS_UNKNOWN
	// DEVICE_STATUS_SUCCESS        = C.ARV_DEVICE_STATUS_SUCCESS
	// DEVICE_STATUS_TIMEOUT        = C.ARV_DEVICE_STATUS_TIMEOUT
	// DEVICE_STATUS_WRITE_ERROR    = C.ARV_DEVICE_STATUS_WRITE_ERROR
	// DEVICE_STATUS_TRANSFER_ERROR = C.ARV_DEVICE_STATUS_TRANSFER_ERROR
	// DEVICE_STATUS_NOT_CONNECTED  = C.ARV_DEVICE_STATUS_NOT_CONNECTED
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

// func (d *Device) GetStatus() (int, error) {
// 	var gerror *C.GError
// 	var err error
// 	var cvalue C.int
// 	cvalue = C.arv_device_get_status(d.device, &gerror)
// 	if unsafe.Pointer(gerror) != nil {
// 		err = errorFromGError(unsafe.Pointer(gerror))
// 	}

// 	return int(cvalue), err
// }
