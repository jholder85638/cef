// Code generated - DO NOT EDIT.

package cef

import (
	// #include "capi_gen.h"
	// cef_resource_handler_t * gocef_scheme_handler_factory_create(cef_scheme_handler_factory_t * self, cef_browser_t * browser, cef_frame_t * frame, cef_string_t * scheme_name, cef_request_t * request, cef_resource_handler_t * (CEF_CALLBACK *callback__)(cef_scheme_handler_factory_t *, cef_browser_t *, cef_frame_t *, cef_string_t *, cef_request_t *)) { return callback__(self, browser, frame, scheme_name, request); }
	"C"
)

// SchemeHandlerFactory (cef_scheme_handler_factory_t from .\include/capi/cef_scheme_capi.h)
// Structure that creates cef_resource_handler_t instances for handling scheme
// requests. The functions of this structure will always be called on the IO
// thread.
type SchemeHandlerFactory C.cef_scheme_handler_factory_t

func (d *SchemeHandlerFactory) toNative() *C.cef_scheme_handler_factory_t {
	return (*C.cef_scheme_handler_factory_t)(d)
}

// Base (base)
// Base structure.
func (d *SchemeHandlerFactory) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Create (create)
// Return a new resource handler instance to handle the request or an NULL
// reference to allow default handling of the request. |browser| and |frame|
// will be the browser window and frame respectively that originated the
// request or NULL if the request did not originate from a browser window (for
// example, if the request came from cef_urlrequest_t). The |request| object
// passed to this function will not contain cookie data.
func (d *SchemeHandlerFactory) Create(browser *Browser, frame *Frame, scheme_name string, request *Request) *ResourceHandler {
	scheme_name_ := C.cef_string_userfree_alloc()
	setCEFStr(scheme_name, scheme_name_)
	defer func() {
		C.cef_string_userfree_free(scheme_name_)
	}()
	return (*ResourceHandler)(C.gocef_scheme_handler_factory_create(d.toNative(), browser.toNative(), frame.toNative(), (*C.cef_string_t)(scheme_name_), request.toNative(), d.create))
}
