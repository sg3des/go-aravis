package aravis

// #cgo pkg-config: aravis-0.8
// #include <arv.h>
import "C"
import "errors"

func toBool(x C.gboolean) bool {
	if int(x) != 0 {
		return true
	} else {
		return false
	}
}

func errorFromGError(gerr *C.GError) error {
	defer C.g_error_free(gerr)
	return errors.New(goString(gerr.message))
}

func goString(cstr *C.gchar) string {
	return C.GoString((*C.char)(cstr))
}
