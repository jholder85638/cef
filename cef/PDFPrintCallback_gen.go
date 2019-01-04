// Code generated - DO NOT EDIT.

package cef

import (
	// #include "PDFPrintCallback_gen.h"
	"C"
	"unsafe"

	"github.com/jholder85638/toolbox/errs"
	"github.com/jholder85638/toolbox/log/jot"
)

// PDFPrintCallbackProxy defines methods required for using PDFPrintCallback.
type PDFPrintCallbackProxy interface {
	OnPdfPrintFinished(self *PDFPrintCallback, path string, ok int32)
}

// PDFPrintCallback (cef_pdf_print_callback_t from .\include/capi/cef_browser_capi.h)
// Callback structure for cef_browser_host_t::PrintToPDF. The functions of this
// structure will be called on the browser process UI thread.
type PDFPrintCallback C.cef_pdf_print_callback_t

// NewPDFPrintCallback creates a new PDFPrintCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewPDFPrintCallback(proxy PDFPrintCallbackProxy) *PDFPrintCallback {
	result := (*PDFPrintCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_pdf_print_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_pdf_print_callback_proxy(result.toNative())
	}
	return result
}

func (d *PDFPrintCallback) toNative() *C.cef_pdf_print_callback_t {
	return (*C.cef_pdf_print_callback_t)(d)
}

func lookupPDFPrintCallbackProxy(obj *BaseRefCounted) PDFPrintCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(PDFPrintCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type PDFPrintCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *PDFPrintCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnPdfPrintFinished (on_pdf_print_finished)
// Method that will be executed when the PDF printing has completed. |path| is
// the output path. |ok| will be true (1) if the printing completed
// successfully or false (0) otherwise.
func (d *PDFPrintCallback) OnPdfPrintFinished(path string, ok int32) {
	lookupPDFPrintCallbackProxy(d.Base()).OnPdfPrintFinished(d, path, ok)
}

//export gocef_pdf_print_callback_on_pdf_print_finished
func gocef_pdf_print_callback_on_pdf_print_finished(self *C.cef_pdf_print_callback_t, path *C.cef_string_t, ok C.int) {
	me__ := (*PDFPrintCallback)(self)
	proxy__ := lookupPDFPrintCallbackProxy(me__.Base())
	path_ := cefstrToString(path)
	proxy__.OnPdfPrintFinished(me__, path_, int32(ok))
}
