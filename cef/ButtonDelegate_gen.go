// Code generated - DO NOT EDIT.

package cef

import (
	// #include "ButtonDelegate_gen.h"
	"C"
	"unsafe"

	"github.com/jholder85638/toolbox/errs"
	"github.com/jholder85638/toolbox/log/jot"
)

// ButtonDelegateProxy defines methods required for using ButtonDelegate.
type ButtonDelegateProxy interface {
	OnButtonPressed(self *ButtonDelegate, button *Button)
	OnButtonStateChanged(self *ButtonDelegate, button *Button)
}

// ButtonDelegate (cef_button_delegate_t from include\capi\views\cef_button_delegate_capi.h)
// Implement this structure to handle Button events. The functions of this
// structure will be called on the browser process UI thread unless otherwise
// indicated.
type ButtonDelegate C.cef_button_delegate_t

// NewButtonDelegate creates a new ButtonDelegate with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewButtonDelegate(proxy ButtonDelegateProxy) *ButtonDelegate {
	result := (*ButtonDelegate)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_button_delegate_t, proxy)))
	if proxy != nil {
		C.gocef_set_button_delegate_proxy(result.toNative())
	}
	return result
}

func (d *ButtonDelegate) toNative() *C.cef_button_delegate_t {
	return (*C.cef_button_delegate_t)(d)
}

func lookupButtonDelegateProxy(obj *BaseRefCounted) ButtonDelegateProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ButtonDelegateProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ButtonDelegateProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ButtonDelegate) Base() *ViewDelegate {
	return (*ViewDelegate)(&d.base)
}

// OnButtonPressed (on_button_pressed)
// Called when |button| is pressed.
func (d *ButtonDelegate) OnButtonPressed(button *Button) {
	lookupButtonDelegateProxy(d.Base().Base()).OnButtonPressed(d, button)
}

//export gocef_button_delegate_on_button_pressed
func gocef_button_delegate_on_button_pressed(self *C.cef_button_delegate_t, button *C.cef_button_t) {
	me__ := (*ButtonDelegate)(self)
	proxy__ := lookupButtonDelegateProxy(me__.Base().Base())
	proxy__.OnButtonPressed(me__, (*Button)(button))
}

// OnButtonStateChanged (on_button_state_changed)
// Called when the state of |button| changes.
func (d *ButtonDelegate) OnButtonStateChanged(button *Button) {
	lookupButtonDelegateProxy(d.Base().Base()).OnButtonStateChanged(d, button)
}

//export gocef_button_delegate_on_button_state_changed
func gocef_button_delegate_on_button_state_changed(self *C.cef_button_delegate_t, button *C.cef_button_t) {
	me__ := (*ButtonDelegate)(self)
	proxy__ := lookupButtonDelegateProxy(me__.Base().Base())
	proxy__.OnButtonStateChanged(me__, (*Button)(button))
}
