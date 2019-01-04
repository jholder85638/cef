// Code generated - DO NOT EDIT.

package cef

import (
	// #include "BrowserProcessHandler_gen.h"
	"C"
	"unsafe"

	"github.com/jholder85638/toolbox/errs"
	"github.com/jholder85638/toolbox/log/jot"
)

// BrowserProcessHandlerProxy defines methods required for using BrowserProcessHandler.
type BrowserProcessHandlerProxy interface {
	OnContextInitialized(self *BrowserProcessHandler)
	OnBeforeChildProcessLaunch(self *BrowserProcessHandler, command_line *CommandLine)
	OnRenderProcessThreadCreated(self *BrowserProcessHandler, extra_info *ListValue)
	GetPrintHandler(self *BrowserProcessHandler) *PrintHandler
	OnScheduleMessagePumpWork(self *BrowserProcessHandler, delay_ms int64)
}

// BrowserProcessHandler (cef_browser_process_handler_t from .\include/capi/cef_browser_process_handler_capi.h)
// Structure used to implement browser process callbacks. The functions of this
// structure will be called on the browser process main thread unless otherwise
// indicated.
type BrowserProcessHandler C.cef_browser_process_handler_t

// NewBrowserProcessHandler creates a new BrowserProcessHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewBrowserProcessHandler(proxy BrowserProcessHandlerProxy) *BrowserProcessHandler {
	result := (*BrowserProcessHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_browser_process_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_browser_process_handler_proxy(result.toNative())
	}
	return result
}

func (d *BrowserProcessHandler) toNative() *C.cef_browser_process_handler_t {
	return (*C.cef_browser_process_handler_t)(d)
}

func lookupBrowserProcessHandlerProxy(obj *BaseRefCounted) BrowserProcessHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(BrowserProcessHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type BrowserProcessHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *BrowserProcessHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnContextInitialized (on_context_initialized)
// Called on the browser process UI thread immediately after the CEF context
// has been initialized.
func (d *BrowserProcessHandler) OnContextInitialized() {
	lookupBrowserProcessHandlerProxy(d.Base()).OnContextInitialized(d)
}

//export gocef_browser_process_handler_on_context_initialized
func gocef_browser_process_handler_on_context_initialized(self *C.cef_browser_process_handler_t) {
	me__ := (*BrowserProcessHandler)(self)
	proxy__ := lookupBrowserProcessHandlerProxy(me__.Base())
	proxy__.OnContextInitialized(me__)
}

// OnBeforeChildProcessLaunch (on_before_child_process_launch)
// Called before a child process is launched. Will be called on the browser
// process UI thread when launching a render process and on the browser
// process IO thread when launching a GPU or plugin process. Provides an
// opportunity to modify the child process command line. Do not keep a
// reference to |command_line| outside of this function.
func (d *BrowserProcessHandler) OnBeforeChildProcessLaunch(command_line *CommandLine) {
	lookupBrowserProcessHandlerProxy(d.Base()).OnBeforeChildProcessLaunch(d, command_line)
}

//export gocef_browser_process_handler_on_before_child_process_launch
func gocef_browser_process_handler_on_before_child_process_launch(self *C.cef_browser_process_handler_t, command_line *C.cef_command_line_t) {
	me__ := (*BrowserProcessHandler)(self)
	proxy__ := lookupBrowserProcessHandlerProxy(me__.Base())
	proxy__.OnBeforeChildProcessLaunch(me__, (*CommandLine)(command_line))
}

// OnRenderProcessThreadCreated (on_render_process_thread_created)
// Called on the browser process IO thread after the main thread has been
// created for a new render process. Provides an opportunity to specify extra
// information that will be passed to
// cef_render_process_handler_t::on_render_thread_created() in the render
// process. Do not keep a reference to |extra_info| outside of this function.
func (d *BrowserProcessHandler) OnRenderProcessThreadCreated(extra_info *ListValue) {
	lookupBrowserProcessHandlerProxy(d.Base()).OnRenderProcessThreadCreated(d, extra_info)
}

//export gocef_browser_process_handler_on_render_process_thread_created
func gocef_browser_process_handler_on_render_process_thread_created(self *C.cef_browser_process_handler_t, extra_info *C.cef_list_value_t) {
	me__ := (*BrowserProcessHandler)(self)
	proxy__ := lookupBrowserProcessHandlerProxy(me__.Base())
	proxy__.OnRenderProcessThreadCreated(me__, (*ListValue)(extra_info))
}

// GetPrintHandler (get_print_handler)
// Return the handler for printing on Linux. If a print handler is not
// provided then printing will not be supported on the Linux platform.
func (d *BrowserProcessHandler) GetPrintHandler() *PrintHandler {
	return lookupBrowserProcessHandlerProxy(d.Base()).GetPrintHandler(d)
}

//export gocef_browser_process_handler_get_print_handler
func gocef_browser_process_handler_get_print_handler(self *C.cef_browser_process_handler_t) *C.cef_print_handler_t {
	me__ := (*BrowserProcessHandler)(self)
	proxy__ := lookupBrowserProcessHandlerProxy(me__.Base())
	return (proxy__.GetPrintHandler(me__)).toNative()
}

// OnScheduleMessagePumpWork (on_schedule_message_pump_work)
// Called from any thread when work has been scheduled for the browser process
// main (UI) thread. This callback is used in combination with CefSettings.
// external_message_pump and cef_do_message_loop_work() in cases where the CEF
// message loop must be integrated into an existing application message loop
// (see additional comments and warnings on CefDoMessageLoopWork). This
// callback should schedule a cef_do_message_loop_work() call to happen on the
// main (UI) thread. |delay_ms| is the requested delay in milliseconds. If
// |delay_ms| is <= 0 then the call should happen reasonably soon. If
// |delay_ms| is > 0 then the call should be scheduled to happen after the
// specified delay and any currently pending scheduled call should be
// cancelled.
func (d *BrowserProcessHandler) OnScheduleMessagePumpWork(delay_ms int64) {
	lookupBrowserProcessHandlerProxy(d.Base()).OnScheduleMessagePumpWork(d, delay_ms)
}

//export gocef_browser_process_handler_on_schedule_message_pump_work
func gocef_browser_process_handler_on_schedule_message_pump_work(self *C.cef_browser_process_handler_t, delay_ms C.int64) {
	me__ := (*BrowserProcessHandler)(self)
	proxy__ := lookupBrowserProcessHandlerProxy(me__.Base())
	proxy__.OnScheduleMessagePumpWork(me__, int64(delay_ms))
}
