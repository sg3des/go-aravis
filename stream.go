package aravis

// #cgo pkg-config: aravis-0.6
// #include <arv.h>
// #include <stdlib.h>
/*
void arv_set_stream_property_long(ArvStream *stream, char *property, long value) {
	g_object_set (stream, property, value, NULL);
}

void arv_set_stream_property_double(ArvStream *stream, char *property, double value) {
	g_object_set (stream, property, value, NULL);
}
*/
import "C"
import (
	"errors"
	"time"
	"unsafe"
)

type Stream struct {
	stream *C.struct__ArvStream
}

func (s *Stream) PushBuffer(b Buffer) {
	C.arv_stream_push_buffer(s.stream, b.buffer)
}

func (s *Stream) PopBuffer() (Buffer, error) {
	var b Buffer
	var err error

	b.buffer, err = C.arv_stream_pop_buffer(s.stream)

	return b, err
}

func (s *Stream) TryPopBuffer() (Buffer, error) {
	var b Buffer
	var err error

	b.buffer, err = C.arv_stream_try_pop_buffer(s.stream)

	return b, err
}

func (s *Stream) TimeoutPopBuffer(t time.Duration) (Buffer, error) {
	var b Buffer
	var err error

	b.buffer, err = C.arv_stream_timeout_pop_buffer(s.stream, C.guint64(t/1000))

	if b.buffer == nil {
		return Buffer{}, errors.New("Aravis returned null pointer")
	}

	return b, err
}

func (s *Stream) Close() {
	C.g_object_unref(C.gpointer(s.stream))
}

func (s *Stream) SetPropertyLong(property string, value int64) {
	cprop := C.CString(property)
	cvalue := C.long(value)
	C.arv_set_stream_property_long(s.stream, cprop, cvalue)
	C.free(unsafe.Pointer(cprop))
}

func (s *Stream) SetPropertyDouble(property string, value float32) {
	cprop := C.CString(property)
	cvalue := C.double(value)
	C.arv_set_stream_property_double(s.stream, cprop, cvalue)
	C.free(unsafe.Pointer(cprop))
}
