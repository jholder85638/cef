// Code generated - DO NOT EDIT.

package cef

import (
	// #include "TextfieldDelegate_gen.h"
	"C"
	"unsafe"

	"github.com/jholder85638/toolbox/errs"
	"github.com/jholder85638/toolbox/log/jot"
)

// TextfieldDelegateProxy defines methods required for using TextfieldDelegate.
type TextfieldDelegateProxy interface {
	OnKeyEvent(self *TextfieldDelegate, textfield *Textfield, event *KeyEvent) int32
	OnAfterUserAction(self *TextfieldDelegate, textfield *Textfield)
}

// TextfieldDelegate (cef_textfield_delegate_t from .\include/capi/views/cef_textfield_delegate_capi.h)
// Implement this structure to handle Textfield events. The functions of this
// structure will be called on the browser process UI thread unless otherwise
// indicated.
type TextfieldDelegate C.cef_textfield_delegate_t

// NewTextfieldDelegate creates a new TextfieldDelegate with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewTextfieldDelegate(proxy TextfieldDelegateProxy) *TextfieldDelegate {
	result := (*TextfieldDelegate)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_textfield_delegate_t, proxy)))
	if proxy != nil {
		C.gocef_set_textfield_delegate_proxy(result.toNative())
	}
	return result
}

func (d *TextfieldDelegate) toNative() *C.cef_textfield_delegate_t {
	return (*C.cef_textfield_delegate_t)(d)
}

func lookupTextfieldDelegateProxy(obj *BaseRefCounted) TextfieldDelegateProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(TextfieldDelegateProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type TextfieldDelegateProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *TextfieldDelegate) Base() *ViewDelegate {
	return (*ViewDelegate)(&d.base)
}

// OnKeyEvent (on_key_event)
// Called when |textfield| recieves a keyboard event. |event| contains
// information about the keyboard event. Return true (1) if the keyboard event
// was handled or false (0) otherwise for default handling.
func (d *TextfieldDelegate) OnKeyEvent(textfield *Textfield, event *KeyEvent) int32 {
	return lookupTextfieldDelegateProxy(d.Base().Base()).OnKeyEvent(d, textfield, event)
}

//export gocef_textfield_delegate_on_key_event
func gocef_textfield_delegate_on_key_event(self *C.cef_textfield_delegate_t, textfield *C.cef_textfield_t, event *C.cef_key_event_t) C.int {
	me__ := (*TextfieldDelegate)(self)
	proxy__ := lookupTextfieldDelegateProxy(me__.Base().Base())
	event_ := event.toGo()
	return C.int(proxy__.OnKeyEvent(me__, (*Textfield)(textfield), event_))
}

// OnAfterUserAction (on_after_user_action)
// Called after performing a user action that may change |textfield|.
func (d *TextfieldDelegate) OnAfterUserAction(textfield *Textfield) {
	lookupTextfieldDelegateProxy(d.Base().Base()).OnAfterUserAction(d, textfield)
}

//export gocef_textfield_delegate_on_after_user_action
func gocef_textfield_delegate_on_after_user_action(self *C.cef_textfield_delegate_t, textfield *C.cef_textfield_t) {
	me__ := (*TextfieldDelegate)(self)
	proxy__ := lookupTextfieldDelegateProxy(me__.Base().Base())
	proxy__.OnAfterUserAction(me__, (*Textfield)(textfield))
}
