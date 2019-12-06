package aravis

// #cgo pkg-config: aravis-0.6
// #include <arv.h>
// #include <stdlib.h>
import "C"
import "unsafe"

type Device struct {
	device *C.struct__ArvDevice
}

func (d *Device) SetStringFeatureValue(feature, value string) {
	cfeature := C.CString(feature)
	cvalue := C.CString(value)
	C.arv_device_set_string_feature_value(d.device, cfeature, cvalue)
	C.free(unsafe.Pointer(cfeature))
	C.free(unsafe.Pointer(cvalue))
}

func (d *Device) GetStringFeatureValue(feature string) (string, error) {
	cfeature := C.CString(feature)
	cvalue, err := C.arv_device_get_string_feature_value(d.device, cfeature)
	C.free(unsafe.Pointer(cfeature))
	return C.GoString(cvalue), err
}

func (d *Device) SetIntegerFeatureValue(feature string, value int64) {
	cfeature := C.CString(feature)
	cvalue := C.long(value)
	C.arv_device_set_integer_feature_value(d.device, cfeature, cvalue)
	C.free(unsafe.Pointer(cfeature))
}

func (d *Device) GetIntegerFeatureValue(feature string) (int64, error) {
	cfeature := C.CString(feature)
	cvalue, err := C.arv_device_get_integer_feature_value(d.device, cfeature)
	C.free(unsafe.Pointer(cfeature))
	return int64(cvalue), err
}
