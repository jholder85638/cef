// Code generated - DO NOT EDIT.

package cef

import (
	// #include "capi_gen.h"
	"C"
	"unsafe"
)

// AddCrossOriginWhitelistEntry (cef_add_cross_origin_whitelist_entry from include\capi\cef_origin_whitelist_capi.h)
// Add an entry to the cross-origin access whitelist.
//
// The same-origin policy restricts how scripts hosted from different origins
// (scheme + domain + port) can communicate. By default, scripts can only access
// resources with the same origin. Scripts hosted on the HTTP and HTTPS schemes
// (but no other schemes) can use the "Access-Control-Allow-Origin" header to
// allow cross-origin requests. For example, https://source.example.com can make
// XMLHttpRequest requests on http://target.example.com if the
// http://target.example.com request returns an "Access-Control-Allow-Origin:
// https://source.example.com" response header.
//
// Scripts in separate frames or iframes and hosted from the same protocol and
// domain suffix can execute cross-origin JavaScript if both pages set the
// document.domain value to the same domain suffix. For example,
// scheme://foo.example.com and scheme://bar.example.com can communicate using
// JavaScript if both domains set document.domain="example.com".
//
// This function is used to allow access to origins that would otherwise violate
// the same-origin policy. Scripts hosted underneath the fully qualified
// |source_origin| URL (like http://www.example.com) will be allowed access to
// all resources hosted on the specified |target_protocol| and |target_domain|.
// If |target_domain| is non-NULL and |allow_target_subdomains| if false (0)
// only exact domain matches will be allowed. If |target_domain| contains a top-
// level domain component (like "example.com") and |allow_target_subdomains| is
// true (1) sub-domain matches will be allowed. If |target_domain| is NULL and
// |allow_target_subdomains| if true (1) all domains and IP addresses will be
// allowed.
//
// This function cannot be used to bypass the restrictions on local or display
// isolated schemes. See the comments on CefRegisterCustomScheme for more
// information.
//
// This function may be called on any thread. Returns false (0) if
// |source_origin| is invalid or the whitelist cannot be accessed.
func AddCrossOriginWhitelistEntry(source_origin, target_protocol, target_domain string, allow_target_subdomains int32) int32 {
	source_origin_ := C.cef_string_userfree_alloc()
	setCEFStr(source_origin, source_origin_)
	defer func() {
		C.cef_string_userfree_free(source_origin_)
	}()
	target_protocol_ := C.cef_string_userfree_alloc()
	setCEFStr(target_protocol, target_protocol_)
	defer func() {
		C.cef_string_userfree_free(target_protocol_)
	}()
	target_domain_ := C.cef_string_userfree_alloc()
	setCEFStr(target_domain, target_domain_)
	defer func() {
		C.cef_string_userfree_free(target_domain_)
	}()
	return int32(C.cef_add_cross_origin_whitelist_entry((*C.cef_string_t)(source_origin_), (*C.cef_string_t)(target_protocol_), (*C.cef_string_t)(target_domain_), C.int(allow_target_subdomains)))
}

// BinaryValueCreate (cef_binary_value_create from .\include/capi/cef_values_capi.h)
// Creates a new object that is not owned by any other object. The specified
// |data| will be copied.
func BinaryValueCreate(data unsafe.Pointer, data_size uint64) *BinaryValue {
	return (*BinaryValue)(C.cef_binary_value_create(data, C.size_t(data_size)))
}

// BrowserHostCreateBrowser (cef_browser_host_create_browser from .\include/capi/cef_browser_capi.h)
// Create a new browser window using the window parameters specified by
// |windowInfo|. All values will be copied internally and the actual window will
// be created on the UI thread. If |request_context| is NULL the global request
// context will be used. This function can be called on any browser process
// thread and will not block.
func BrowserHostCreateBrowser(windowInfo *WindowInfo, client *Client, url string, settings *BrowserSettings, request_context *RequestContext) int32 {
	url_ := C.cef_string_userfree_alloc()
	setCEFStr(url, url_)
	defer func() {
		C.cef_string_userfree_free(url_)
	}()
	return int32(C.cef_browser_host_create_browser(windowInfo.toNative(&C.cef_window_info_t{}), client.toNative(), (*C.cef_string_t)(url_), settings.toNative(&C.cef_browser_settings_t{}), request_context.toNative()))
}

// BrowserHostCreateBrowserSync (cef_browser_host_create_browser_sync from .\include/capi/cef_browser_capi.h)
// Create a new browser window using the window parameters specified by
// |windowInfo|. If |request_context| is NULL the global request context will be
// used. This function can only be called on the browser process UI thread.
func BrowserHostCreateBrowserSync(windowInfo *WindowInfo, client *Client, url string, settings *BrowserSettings, request_context *RequestContext) *Browser {
	url_ := C.cef_string_userfree_alloc()
	setCEFStr(url, url_)
	defer func() {
		C.cef_string_userfree_free(url_)
	}()
	return (*Browser)(C.cef_browser_host_create_browser_sync(windowInfo.toNative(&C.cef_window_info_t{}), client.toNative(), (*C.cef_string_t)(url_), settings.toNative(&C.cef_browser_settings_t{}), request_context.toNative()))
}

// BrowserViewCreate (cef_browser_view_create from include\capi\views\cef_browser_view_capi.h)
// Create a new BrowserView. The underlying cef_browser_t will not be created
// until this view is added to the views hierarchy.
func BrowserViewCreate(client *Client, url string, settings *BrowserSettings, request_context *RequestContext, delegate *BrowserViewDelegate) *BrowserView {
	url_ := C.cef_string_userfree_alloc()
	setCEFStr(url, url_)
	defer func() {
		C.cef_string_userfree_free(url_)
	}()
	var delegate_ *C.cef_browser_view_delegate_t
	if delegate != nil {
		delegate_ = delegate.toNative()
	}
	return (*BrowserView)(C.cef_browser_view_create(client.toNative(), (*C.cef_string_t)(url_), settings.toNative(&C.cef_browser_settings_t{}), request_context.toNative(), delegate_))
}

// BrowserViewGetForBrowser (cef_browser_view_get_for_browser from include\capi\views\cef_browser_view_capi.h)
// Returns the BrowserView associated with |browser|.
func BrowserViewGetForBrowser(browser *Browser) *BrowserView {
	return (*BrowserView)(C.cef_browser_view_get_for_browser(browser.toNative()))
}

// ClearCrossOriginWhitelist (cef_clear_cross_origin_whitelist from include\capi\cef_origin_whitelist_capi.h)
// Remove all entries from the cross-origin access whitelist. Returns false (0)
// if the whitelist cannot be accessed.
func ClearCrossOriginWhitelist() int32 {
	return int32(C.cef_clear_cross_origin_whitelist())
}

// ClearSchemeHandlerFactories (cef_clear_scheme_handler_factories from .\include/capi/cef_scheme_capi.h)
// Clear all scheme handler factories registered with the global request
// context. Returns false (0) on error. This function may be called on any
// thread in the browser process. Using this function is equivalent to calling c
// ef_request_tContext::cef_request_context_get_global_context()->clear_scheme_h
// andler_factories().
func ClearSchemeHandlerFactories() int32 {
	return int32(C.cef_clear_scheme_handler_factories())
}

// CommandLineCreate (cef_command_line_create from .\include/capi/cef_command_line_capi.h)
// Create a new cef_command_line_t instance.
func CommandLineCreate() *CommandLine {
	return (*CommandLine)(C.cef_command_line_create())
}

// CommandLineGetGlobal (cef_command_line_get_global from .\include/capi/cef_command_line_capi.h)
// Returns the singleton global cef_command_line_t object. The returned object
// will be read-only.
func CommandLineGetGlobal() *CommandLine {
	return (*CommandLine)(C.cef_command_line_get_global())
}

// CookieManagerCreateManager (cef_cookie_manager_create_manager from .\include/capi/cef_cookie_capi.h)
// Creates a new cookie manager. If |path| is NULL data will be stored in memory
// only. Otherwise, data will be stored at the specified |path|. To persist
// session cookies (cookies without an expiry date or validity interval) set
// |persist_session_cookies| to true (1). Session cookies are generally intended
// to be transient and most Web browsers do not persist them. If |callback| is
// non-NULL it will be executed asnychronously on the IO thread after the
// manager's storage has been initialized.
func CookieManagerCreateManager(path string, persist_session_cookies int32, callback *CompletionCallback) *CookieManager {
	path_ := C.cef_string_userfree_alloc()
	setCEFStr(path, path_)
	defer func() {
		C.cef_string_userfree_free(path_)
	}()
	return (*CookieManager)(C.cef_cookie_manager_create_manager((*C.cef_string_t)(path_), C.int(persist_session_cookies), callback.toNative()))
}

// CookieManagerGetBlockingManager (cef_cookie_manager_get_blocking_manager from .\include/capi/cef_cookie_capi.h)
// Returns a cookie manager that neither stores nor retrieves cookies. All usage
// of cookies will be blocked including cookies accessed via the network
// (request/response headers), via JavaScript (document.cookie), and via
// cef_cookie_manager_t functions. No cookies will be displayed in DevTools. If
// you wish to only block cookies sent via the network use the
// cef_request_tHandler CanGetCookies and CanSetCookie functions instead.
func CookieManagerGetBlockingManager() *CookieManager {
	return (*CookieManager)(C.cef_cookie_manager_get_blocking_manager())
}

// CookieManagerGetGlobalManager (cef_cookie_manager_get_global_manager from .\include/capi/cef_cookie_capi.h)
// Returns the global cookie manager. By default data will be stored at
// CefSettings.cache_path if specified or in memory otherwise. If |callback| is
// non-NULL it will be executed asnychronously on the IO thread after the
// manager's storage has been initialized. Using this function is equivalent to
// calling cef_request_tContext::cef_request_context_get_global_context()->get_d
// efault_cookie_manager().
func CookieManagerGetGlobalManager(callback *CompletionCallback) *CookieManager {
	return (*CookieManager)(C.cef_cookie_manager_get_global_manager(callback.toNative()))
}

// CreateContextShared (cef_create_context_shared from .\include/capi/cef_request_context_capi.h)
// Creates a new context object that shares storage with |other| and uses an
// optional |handler|.
func CreateContextShared(other *RequestContext, handler *RequestContextHandler) *RequestContext {
	return (*RequestContext)(C.cef_create_context_shared(other.toNative(), handler.toNative()))
}

// CurrentlyOn (cef_currently_on from .\include/capi/cef_task_capi.h)
// Returns true (1) if called on the specified thread. Equivalent to using
// cef_task_tRunner::GetForThread(threadId)->belongs_to_current_thread().
func CurrentlyOn(threadId ThreadID) int32 {
	return int32(C.cef_currently_on(C.cef_thread_id_t(threadId)))
}

// DictionaryValueCreate (cef_dictionary_value_create from .\include/capi/cef_values_capi.h)
// Creates a new object that is not owned by any other object.
func DictionaryValueCreate() *DictionaryValue {
	return (*DictionaryValue)(C.cef_dictionary_value_create())
}

// DisplayGetAlls (cef_display_get_alls from include\capi\views\cef_display_capi.h)
// Returns all Displays. Mirrored displays are excluded; this function is
// intended to return distinct, usable displays.
func DisplayGetAlls(displaysCount *uint64, displays **Display) {
	displays_ := (*displays).toNative()
	C.cef_display_get_alls((*C.size_t)(displaysCount), &displays_)
}

// DisplayGetCount (cef_display_get_count from include\capi\views\cef_display_capi.h)
// Returns the total number of Displays. Mirrored displays are excluded; this
// function is intended to return the number of distinct, usable displays.
func DisplayGetCount() uint64 {
	return uint64(C.cef_display_get_count())
}

// DisplayGetMatchingBounds (cef_display_get_matching_bounds from include\capi\views\cef_display_capi.h)
// Returns the Display that most closely intersects |bounds|.  Set
// |input_pixel_coords| to true (1) if |bounds| is in pixel coordinates instead
// of density independent pixels (DIP).
func DisplayGetMatchingBounds(bounds *Rect, input_pixel_coords int32) *Display {
	return (*Display)(C.cef_display_get_matching_bounds(bounds.toNative(&C.cef_rect_t{}), C.int(input_pixel_coords)))
}

// DisplayGetNearestPoint (cef_display_get_nearest_point from include\capi\views\cef_display_capi.h)
// Returns the Display nearest |point|. Set |input_pixel_coords| to true (1) if
// |point| is in pixel coordinates instead of density independent pixels (DIP).
func DisplayGetNearestPoint(point *Point, input_pixel_coords int32) *Display {
	return (*Display)(C.cef_display_get_nearest_point(point.toNative(&C.cef_point_t{}), C.int(input_pixel_coords)))
}

// DisplayGetPrimary (cef_display_get_primary from include\capi\views\cef_display_capi.h)
// Returns the primary Display.
func DisplayGetPrimary() *Display {
	return (*Display)(C.cef_display_get_primary())
}

// DoMessageLoopWork (cef_do_message_loop_work from include\capi\cef_app_capi.h)
// Perform a single iteration of CEF message loop processing. This function is
// provided for cases where the CEF message loop must be integrated into an
// existing application message loop. Use of this function is not recommended
// for most users; use either the cef_run_message_loop() function or
// CefSettings.multi_threaded_message_loop if possible. When using this function
// care must be taken to balance performance against excessive CPU usage. It is
// recommended to enable the CefSettings.external_message_pump option when using
// this function so that
// cef_browser_process_handler_t::on_schedule_message_pump_work() callbacks can
// facilitate the scheduling process. This function should only be called on the
// main application thread and only if cef_initialize() is called with a
// CefSettings.multi_threaded_message_loop value of false (0). This function
// will not block.
func DoMessageLoopWork() {
	C.cef_do_message_loop_work()
}

// DragDataCreate (cef_drag_data_create from .\include/capi/cef_drag_data_capi.h)
// Create a new cef_drag_data_t object.
func DragDataCreate() *DragData {
	return (*DragData)(C.cef_drag_data_create())
}

// EnableHighdpiSupport (cef_enable_highdpi_support from include\capi\cef_app_capi.h)
// Call during process startup to enable High-DPI support on Windows 7 or newer.
// Older versions of Windows should be left DPI-unaware because they do not
// support DirectWrite and GDI fonts are kerned very badly.
func EnableHighdpiSupport() {
	C.cef_enable_highdpi_support()
}

// ExecuteProcess (cef_execute_process from include\capi\cef_app_capi.h)
// This function should be called from the application entry point function to
// execute a secondary process. It can be used to run secondary processes from
// the browser client executable (default behavior) or from a separate
// executable specified by the CefSettings.browser_subprocess_path value. If
// called for the browser process (identified by no "type" command-line value)
// it will return immediately with a value of -1. If called for a recognized
// secondary process it will block until the process should exit and then return
// the process exit code. The |application| parameter may be NULL. The
// |windows_sandbox_info| parameter is only used on Windows and may be NULL (see
// cef_sandbox_win.h for details).
func ExecuteProcess(args *MainArgs, application *App, windows_sandbox_info unsafe.Pointer) int32 {
	return int32(C.cef_execute_process(args.toNative(&C.cef_main_args_t{}), application.toNative(), windows_sandbox_info))
}

// ImageCreate (cef_image_create from .\include/capi/cef_image_capi.h)
// Create a new cef_image_t. It will initially be NULL. Use the Add*() functions
// to add representations at different scale factors.
func ImageCreate() *Image {
	return (*Image)(C.cef_image_create())
}

// Initialize (cef_initialize from include\capi\cef_app_capi.h)
// This function should be called on the main application thread to initialize
// the CEF browser process. The |application| parameter may be NULL. A return
// value of true (1) indicates that it succeeded and false (0) indicates that it
// failed. The |windows_sandbox_info| parameter is only used on Windows and may
// be NULL (see cef_sandbox_win.h for details).
func Initialize(args *MainArgs, settings *Settings, application *App, windows_sandbox_info unsafe.Pointer) int32 {
	return int32(C.cef_initialize(args.toNative(&C.cef_main_args_t{}), settings.toNative(&C.cef_settings_t{}), application.toNative(), windows_sandbox_info))
}

// IsCertStatusError (cef_is_cert_status_error from .\include/capi/cef_ssl_info_capi.h)
// Returns true (1) if the certificate status has any error, major or minor.
func IsCertStatusError(status CertStatus) int32 {
	return int32(C.cef_is_cert_status_error(C.cef_cert_status_t(status)))
}

// IsCertStatusMinorError (cef_is_cert_status_minor_error from .\include/capi/cef_ssl_info_capi.h)
// Returns true (1) if the certificate status represents only minor errors (e.g.
// failure to verify certificate revocation).
func IsCertStatusMinorError(status CertStatus) int32 {
	return int32(C.cef_is_cert_status_minor_error(C.cef_cert_status_t(status)))
}

// IsWebPluginUnstable (cef_is_web_plugin_unstable from .\include/capi/cef_web_plugin_capi.h)
// Query if a plugin is unstable. Can be called on any thread in the browser
// process.
func IsWebPluginUnstable(path string, callback *WebPluginUnstableCallback) {
	path_ := C.cef_string_userfree_alloc()
	setCEFStr(path, path_)
	defer func() {
		C.cef_string_userfree_free(path_)
	}()
	C.cef_is_web_plugin_unstable((*C.cef_string_t)(path_), callback.toNative())
}

// LabelButtonCreate (cef_label_button_create from include\capi\views\cef_label_button_capi.h)
// Create a new LabelButton. A |delegate| must be provided to handle the button
// click. |text| will be shown on the LabelButton and used as the default
// accessible name. If |with_frame| is true (1) the button will have a visible
// frame at all times, center alignment, additional padding and a default
// minimum size of 70x33 DIP. If |with_frame| is false (0) the button will only
// have a visible frame on hover/press, left alignment, less padding and no
// default minimum size.
func LabelButtonCreate(delegate *ButtonDelegate, text string, with_frame int32) *LabelButton {
	var delegate_ *C.cef_button_delegate_t
	if delegate != nil {
		delegate_ = delegate.toNative()
	}
	text_ := C.cef_string_userfree_alloc()
	setCEFStr(text, text_)
	defer func() {
		C.cef_string_userfree_free(text_)
	}()
	return (*LabelButton)(C.cef_label_button_create(delegate_, (*C.cef_string_t)(text_), C.int(with_frame)))
}

// ListValueCreate (cef_list_value_create from .\include/capi/cef_values_capi.h)
// Creates a new object that is not owned by any other object.
func ListValueCreate() *ListValue {
	return (*ListValue)(C.cef_list_value_create())
}

// MenuButtonCreate (cef_menu_button_create from include\capi\views\cef_menu_button_capi.h)
// Create a new MenuButton. A |delegate| must be provided to call show_menu()
// when the button is clicked. |text| will be shown on the MenuButton and used
// as the default accessible name. If |with_frame| is true (1) the button will
// have a visible frame at all times, center alignment, additional padding and a
// default minimum size of 70x33 DIP. If |with_frame| is false (0) the button
// will only have a visible frame on hover/press, left alignment, less padding
// and no default minimum size. If |with_menu_marker| is true (1) a menu marker
// will be added to the button.
func MenuButtonCreate(delegate *MenuButtonDelegate, text string, with_frame, with_menu_marker int32) *MenuButton {
	var delegate_ *C.cef_menu_button_delegate_t
	if delegate != nil {
		delegate_ = delegate.toNative()
	}
	text_ := C.cef_string_userfree_alloc()
	setCEFStr(text, text_)
	defer func() {
		C.cef_string_userfree_free(text_)
	}()
	return (*MenuButton)(C.cef_menu_button_create(delegate_, (*C.cef_string_t)(text_), C.int(with_frame), C.int(with_menu_marker)))
}

// MenuModelCreate (cef_menu_model_create from .\include/capi/cef_menu_model_capi.h)
// Create a new MenuModel with the specified |delegate|.
func MenuModelCreate(delegate *MenuModelDelegate) *MenuModel {
	var delegate_ *C.cef_menu_model_delegate_t
	if delegate != nil {
		delegate_ = delegate.toNative()
	}
	return (*MenuModel)(C.cef_menu_model_create(delegate_))
}

// PanelCreate (cef_panel_create from include\capi\views\cef_panel_capi.h)
// Create a new Panel.
func PanelCreate(delegate *PanelDelegate) *Panel {
	var delegate_ *C.cef_panel_delegate_t
	if delegate != nil {
		delegate_ = delegate.toNative(&C.cef_panel_delegate_t{})
	}
	return (*Panel)(C.cef_panel_create(delegate_))
}

// PostDataCreate (cef_post_data_create from .\include/capi/cef_request_capi.h)
// Create a new cef_post_data_t object.
func PostDataCreate() *PostData {
	return (*PostData)(C.cef_post_data_create())
}

// PostDataElementCreate (cef_post_data_element_create from .\include/capi/cef_request_capi.h)
// Create a new cef_post_data_element_t object.
func PostDataElementCreate() *PostDataElement {
	return (*PostDataElement)(C.cef_post_data_element_create())
}

// PostDelayedTask (cef_post_delayed_task from .\include/capi/cef_task_capi.h)
// Post a task for delayed execution on the specified thread. Equivalent to
// using cef_task_tRunner::GetForThread(threadId)->PostDelayedTask(task,
// delay_ms).
func PostDelayedTask(threadId ThreadID, task *Task, delay_ms int64) int32 {
	return int32(C.cef_post_delayed_task(C.cef_thread_id_t(threadId), task.toNative(), C.int64(delay_ms)))
}

// PostTask (cef_post_task from .\include/capi/cef_task_capi.h)
// Post a task for execution on the specified thread. Equivalent to using
// cef_task_tRunner::GetForThread(threadId)->PostTask(task).
func PostTask(threadId ThreadID, task *Task) int32 {
	return int32(C.cef_post_task(C.cef_thread_id_t(threadId), task.toNative()))
}

// PrintSettingsCreate (cef_print_settings_create from .\include/capi/cef_print_settings_capi.h)
// Create a new cef_print_settings_t object.
func PrintSettingsCreate() *PrintSettings {
	return (*PrintSettings)(C.cef_print_settings_create())
}

// ProcessMessageCreate (cef_process_message_create from .\include/capi/cef_process_message_capi.h)
// Create a new cef_process_message_t object with the specified name.
func ProcessMessageCreate(name string) *ProcessMessage {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	return (*ProcessMessage)(C.cef_process_message_create((*C.cef_string_t)(name_)))
}

// QuitMessageLoop (cef_quit_message_loop from include\capi\cef_app_capi.h)
// Quit the CEF message loop that was started by calling cef_run_message_loop().
// This function should only be called on the main application thread and only
// if cef_run_message_loop() was used.
func QuitMessageLoop() {
	C.cef_quit_message_loop()
}

// RefreshWebPlugins (cef_refresh_web_plugins from .\include/capi/cef_web_plugin_capi.h)
// Cause the plugin list to refresh the next time it is accessed regardless of
// whether it has already been loaded. Can be called on any thread in the
// browser process.
func RefreshWebPlugins() {
	C.cef_refresh_web_plugins()
}

// RegisterExtension (cef_register_extension from .\include/capi/cef_v8_capi.h)
// Register a new V8 extension with the specified JavaScript extension code and
// handler. Functions implemented by the handler are prototyped using the
// keyword 'native'. The calling of a native function is restricted to the scope
// in which the prototype of the native function is defined. This function may
// only be called on the render process main thread.
//
// Example JavaScript extension code: <pre>
//   // create the 'example' global object if it doesn't already exist.
//   if (!example)
//     example = {};
//   // create the 'example.test' global object if it doesn't already exist.
//   if (!example.test)
//     example.test = {};
//   (function() {
//     // Define the function 'example.test.myfunction'.
//     example.test.myfunction = function() {
//       // Call CefV8Handler::Execute() with the function name 'MyFunction'
//       // and no arguments.
//       native function MyFunction();
//       return MyFunction();
//     };
//     // Define the getter function for parameter 'example.test.myparam'.
//     example.test.__defineGetter__('myparam', function() {
//       // Call CefV8Handler::Execute() with the function name 'GetMyParam'
//       // and no arguments.
//       native function GetMyParam();
//       return GetMyParam();
//     });
//     // Define the setter function for parameter 'example.test.myparam'.
//     example.test.__defineSetter__('myparam', function(b) {
//       // Call CefV8Handler::Execute() with the function name 'SetMyParam'
//       // and a single argument.
//       native function SetMyParam();
//       if(b) SetMyParam(b);
//     });
//
//     // Extension definitions can also contain normal JavaScript variables
//     // and functions.
//     var myint = 0;
//     example.test.increment = function() {
//       myint += 1;
//       return myint;
//     };
//   })();
// </pre> Example usage in the page: <pre>
//   // Call the function.
//   example.test.myfunction();
//   // Set the parameter.
//   example.test.myparam = value;
//   // Get the parameter.
//   value = example.test.myparam;
//   // Call another function.
//   example.test.increment();
// </pre>
func RegisterExtension(extension_name, javascript_code string, handler *V8handler) int32 {
	extension_name_ := C.cef_string_userfree_alloc()
	setCEFStr(extension_name, extension_name_)
	defer func() {
		C.cef_string_userfree_free(extension_name_)
	}()
	javascript_code_ := C.cef_string_userfree_alloc()
	setCEFStr(javascript_code, javascript_code_)
	defer func() {
		C.cef_string_userfree_free(javascript_code_)
	}()
	return int32(C.cef_register_extension((*C.cef_string_t)(extension_name_), (*C.cef_string_t)(javascript_code_), handler.toNative()))
}

// RegisterSchemeHandlerFactory (cef_register_scheme_handler_factory from .\include/capi/cef_scheme_capi.h)
// Register a scheme handler factory with the global request context. An NULL
// |domain_name| value for a standard scheme will cause the factory to match all
// domain names. The |domain_name| value will be ignored for non-standard
// schemes. If |scheme_name| is a built-in scheme and no handler is returned by
// |factory| then the built-in scheme handler factory will be called. If
// |scheme_name| is a custom scheme then you must also implement the
// cef_app_t::on_register_custom_schemes() function in all processes. This
// function may be called multiple times to change or remove the factory that
// matches the specified |scheme_name| and optional |domain_name|. Returns false
// (0) if an error occurs. This function may be called on any thread in the
// browser process. Using this function is equivalent to calling cef_request_tCo
// ntext::cef_request_context_get_global_context()->register_scheme_handler_fact
// ory().
func RegisterSchemeHandlerFactory(scheme_name, domain_name string, factory *SchemeHandlerFactory) int32 {
	scheme_name_ := C.cef_string_userfree_alloc()
	setCEFStr(scheme_name, scheme_name_)
	defer func() {
		C.cef_string_userfree_free(scheme_name_)
	}()
	domain_name_ := C.cef_string_userfree_alloc()
	setCEFStr(domain_name, domain_name_)
	defer func() {
		C.cef_string_userfree_free(domain_name_)
	}()
	return int32(C.cef_register_scheme_handler_factory((*C.cef_string_t)(scheme_name_), (*C.cef_string_t)(domain_name_), factory.toNative()))
}

// RegisterWebPluginCrash (cef_register_web_plugin_crash from .\include/capi/cef_web_plugin_capi.h)
// Register a plugin crash. Can be called on any thread in the browser process
// but will be executed on the IO thread.
func RegisterWebPluginCrash(path string) {
	path_ := C.cef_string_userfree_alloc()
	setCEFStr(path, path_)
	defer func() {
		C.cef_string_userfree_free(path_)
	}()
	C.cef_register_web_plugin_crash((*C.cef_string_t)(path_))
}

// RegisterWidevineCdm (cef_register_widevine_cdm from .\include/capi/cef_web_plugin_capi.h)
// Register the Widevine CDM plugin.
//
// The client application is responsible for downloading an appropriate
// platform-specific CDM binary distribution from Google, extracting the
// contents, and building the required directory structure on the local machine.
// The cef_browser_host_t::StartDownload function and CefZipArchive structure
// can be used to implement this functionality in CEF. Contact Google via
// https://www.widevine.com/contact.html for details on CDM download.
//
// |path| is a directory that must contain the following files:
//   1. manifest.json file from the CDM binary distribution (see below).
//   2. widevinecdm file from the CDM binary distribution (e.g.
//      widevinecdm.dll on on Windows, libwidevinecdm.dylib on OS X,
//      libwidevinecdm.so on Linux).
//
// If any of these files are missing or if the manifest file has incorrect
// contents the registration will fail and |callback| will receive a |result|
// value of CEF_CDM_REGISTRATION_ERROR_INCORRECT_CONTENTS.
//
// The manifest.json file must contain the following keys:
//   A. "os": Supported OS (e.g. "mac", "win" or "linux").
//   B. "arch": Supported architecture (e.g. "ia32" or "x64").
//   C. "x-cdm-module-versions": Module API version (e.g. "4").
//   D. "x-cdm-interface-versions": Interface API version (e.g. "8").
//   E. "x-cdm-host-versions": Host API version (e.g. "8").
//   F. "version": CDM version (e.g. "1.4.8.903").
//   G. "x-cdm-codecs": List of supported codecs (e.g. "vp8,vp9.0,avc1").
//
// A through E are used to verify compatibility with the current Chromium
// version. If the CDM is not compatible the registration will fail and
// |callback| will receive a |result| value of
// CEF_CDM_REGISTRATION_ERROR_INCOMPATIBLE.
//
// |callback| will be executed asynchronously once registration is complete.
//
// On Linux this function must be called before cef_initialize() and the
// registration cannot be changed during runtime. If registration is not
// supported at the time that cef_register_widevine_cdm() is called then
// |callback| will receive a |result| value of
// CEF_CDM_REGISTRATION_ERROR_NOT_SUPPORTED.
func RegisterWidevineCdm(path string, callback *RegisterCdmCallback) {
	path_ := C.cef_string_userfree_alloc()
	setCEFStr(path, path_)
	defer func() {
		C.cef_string_userfree_free(path_)
	}()
	C.cef_register_widevine_cdm((*C.cef_string_t)(path_), callback.toNative())
}

// RemoveCrossOriginWhitelistEntry (cef_remove_cross_origin_whitelist_entry from include\capi\cef_origin_whitelist_capi.h)
// Remove an entry from the cross-origin access whitelist. Returns false (0) if
// |source_origin| is invalid or the whitelist cannot be accessed.
func RemoveCrossOriginWhitelistEntry(source_origin, target_protocol, target_domain string, allow_target_subdomains int32) int32 {
	source_origin_ := C.cef_string_userfree_alloc()
	setCEFStr(source_origin, source_origin_)
	defer func() {
		C.cef_string_userfree_free(source_origin_)
	}()
	target_protocol_ := C.cef_string_userfree_alloc()
	setCEFStr(target_protocol, target_protocol_)
	defer func() {
		C.cef_string_userfree_free(target_protocol_)
	}()
	target_domain_ := C.cef_string_userfree_alloc()
	setCEFStr(target_domain, target_domain_)
	defer func() {
		C.cef_string_userfree_free(target_domain_)
	}()
	return int32(C.cef_remove_cross_origin_whitelist_entry((*C.cef_string_t)(source_origin_), (*C.cef_string_t)(target_protocol_), (*C.cef_string_t)(target_domain_), C.int(allow_target_subdomains)))
}

// RequestContextCreateContext (cef_request_context_create_context from .\include/capi/cef_request_context_capi.h)
// Creates a new context object with the specified |settings| and optional
// |handler|.
func RequestContextCreateContext(settings *RequestContextSettings, handler *RequestContextHandler) *RequestContext {
	return (*RequestContext)(C.cef_request_context_create_context(settings.toNative(&C.cef_request_context_settings_t{}), handler.toNative()))
}

// RequestContextGetGlobalContext (cef_request_context_get_global_context from .\include/capi/cef_request_context_capi.h)
// Returns the global context object.
func RequestContextGetGlobalContext() *RequestContext {
	return (*RequestContext)(C.cef_request_context_get_global_context())
}

// RequestCreate (cef_request_create from .\include/capi/cef_request_capi.h)
// Create a new cef_request_t object.
func RequestCreate() *Request {
	return (*Request)(C.cef_request_create())
}

// ResourceBundleGetGlobal (cef_resource_bundle_get_global from include\capi\cef_resource_bundle_capi.h)
// Returns the global resource bundle instance.
func ResourceBundleGetGlobal() *ResourceBundle {
	return (*ResourceBundle)(C.cef_resource_bundle_get_global())
}

// ResponseCreate (cef_response_create from .\include/capi/cef_response_capi.h)
// Create a new cef_response_t object.
func ResponseCreate() *Response {
	return (*Response)(C.cef_response_create())
}

// RunMessageLoop (cef_run_message_loop from include\capi\cef_app_capi.h)
// Run the CEF message loop. Use this function instead of an application-
// provided message loop to get the best balance between performance and CPU
// usage. This function should only be called on the main application thread and
// only if cef_initialize() is called with a
// CefSettings.multi_threaded_message_loop value of false (0). This function
// will block until a quit message is received by the system.
func RunMessageLoop() {
	C.cef_run_message_loop()
}

// ScrollViewCreate (cef_scroll_view_create from include\capi\views\cef_scroll_view_capi.h)
// Create a new ScrollView.
func ScrollViewCreate(delegate *ViewDelegate) *ScrollView {
	var delegate_ *C.cef_view_delegate_t
	if delegate != nil {
		delegate_ = delegate.toNative()
	}
	return (*ScrollView)(C.cef_scroll_view_create(delegate_))
}

// ServerCreate (cef_server_create from include\capi\cef_server_capi.h)
// Create a new server that binds to |address| and |port|. |address| must be a
// valid IPv4 or IPv6 address (e.g. 127.0.0.1 or ::1) and |port| must be a port
// number outside of the reserved range (e.g. between 1025 and 65535 on most
// platforms). |backlog| is the maximum number of pending connections. A new
// thread will be created for each CreateServer call (the "dedicated server
// thread"). It is therefore recommended to use a different cef_server_handler_t
// instance for each CreateServer call to avoid thread safety issues in the
// cef_server_handler_t implementation. The
// cef_server_handler_t::OnServerCreated function will be called on the
// dedicated server thread to report success or failure. See
// cef_server_handler_t::OnServerCreated documentation for a description of
// server lifespan.
func ServerCreate(address string, port uint16, backlog int32, handler *ServerHandler) {
	address_ := C.cef_string_userfree_alloc()
	setCEFStr(address, address_)
	defer func() {
		C.cef_string_userfree_free(address_)
	}()
	C.cef_server_create((*C.cef_string_t)(address_), C.uint16(port), C.int(backlog), handler.toNative())
}

// SetOsmodalLoop (cef_set_osmodal_loop from include\capi\cef_app_capi.h)
// Set to true (1) before calling Windows APIs like TrackPopupMenu that enter a
// modal message loop. Set to false (0) after exiting the modal message loop.
func SetOsmodalLoop(osModalLoop int32) {
	C.cef_set_osmodal_loop(C.int(osModalLoop))
}

// Shutdown (cef_shutdown from include\capi\cef_app_capi.h)
// This function should be called on the main application thread to shut down
// the CEF browser process before the application exits.
func Shutdown() {
	C.cef_shutdown()
}

// StreamReaderCreateForData (cef_stream_reader_create_for_data from .\include/capi/cef_stream_capi.h)
// Create a new cef_stream_reader_t object from data.
func StreamReaderCreateForData(data unsafe.Pointer, size uint64) *StreamReader {
	return (*StreamReader)(C.cef_stream_reader_create_for_data(data, C.size_t(size)))
}

// StreamReaderCreateForFile (cef_stream_reader_create_for_file from .\include/capi/cef_stream_capi.h)
// Create a new cef_stream_reader_t object from a file.
func StreamReaderCreateForFile(fileName string) *StreamReader {
	fileName_ := C.cef_string_userfree_alloc()
	setCEFStr(fileName, fileName_)
	defer func() {
		C.cef_string_userfree_free(fileName_)
	}()
	return (*StreamReader)(C.cef_stream_reader_create_for_file((*C.cef_string_t)(fileName_)))
}

// StreamReaderCreateForHandler (cef_stream_reader_create_for_handler from .\include/capi/cef_stream_capi.h)
// Create a new cef_stream_reader_t object from a custom handler.
func StreamReaderCreateForHandler(handler *ReadHandler) *StreamReader {
	return (*StreamReader)(C.cef_stream_reader_create_for_handler(handler.toNative()))
}

// StreamWriterCreateForFile (cef_stream_writer_create_for_file from .\include/capi/cef_stream_capi.h)
// Create a new cef_stream_writer_t object for a file.
func StreamWriterCreateForFile(fileName string) *StreamWriter {
	fileName_ := C.cef_string_userfree_alloc()
	setCEFStr(fileName, fileName_)
	defer func() {
		C.cef_string_userfree_free(fileName_)
	}()
	return (*StreamWriter)(C.cef_stream_writer_create_for_file((*C.cef_string_t)(fileName_)))
}

// StreamWriterCreateForHandler (cef_stream_writer_create_for_handler from .\include/capi/cef_stream_capi.h)
// Create a new cef_stream_writer_t object for a custom handler.
func StreamWriterCreateForHandler(handler *WriteHandler) *StreamWriter {
	return (*StreamWriter)(C.cef_stream_writer_create_for_handler(handler.toNative()))
}

// StringASCIIToUtf16 (cef_string_ascii_to_utf16 from .\include/internal/cef_string_types.h)
func StringASCIIToUtf16(src string, src_len uint64, output *string) int32 {
	return int32(C.cef_string_ascii_to_utf16((*C.char)(src), C.size_t(src_len), (*C.cef_string_utf16_t)(output)))
}

// StringASCIIToWide (cef_string_ascii_to_wide from .\include/internal/cef_string_types.h)
func StringASCIIToWide(src string, src_len uint64, output *string) int32 {
	return int32(C.cef_string_ascii_to_wide((*C.char)(src), C.size_t(src_len), (*C.cef_string_wide_t)(output)))
}

// StringListAlloc (cef_string_list_alloc from .\include/internal/cef_string_list.h)
// Allocate a new string map.
func StringListAlloc() StringList {
	return StringList(C.cef_string_list_alloc())
}

// StringListAppend (cef_string_list_append from .\include/internal/cef_string_list.h)
// Append a new value at the end of the string list.
func StringListAppend(list StringList, value string) {
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(value, value_)
	defer func() {
		C.cef_string_userfree_free(value_)
	}()
	C.cef_string_list_append(C.cef_string_list_t(list), (*C.cef_string_t)(value_))
}

// StringListClear (cef_string_list_clear from .\include/internal/cef_string_list.h)
// Clear the string list.
func StringListClear(list StringList) {
	C.cef_string_list_clear(C.cef_string_list_t(list))
}

// StringListCopy (cef_string_list_copy from .\include/internal/cef_string_list.h)
// Creates a copy of an existing string list.
func StringListCopy(list StringList) StringList {
	return StringList(C.cef_string_list_copy(C.cef_string_list_t(list)))
}

// StringListFree (cef_string_list_free from .\include/internal/cef_string_list.h)
// Free the string list.
func StringListFree(list StringList) {
	C.cef_string_list_free(C.cef_string_list_t(list))
}

// StringListSize (cef_string_list_size from .\include/internal/cef_string_list.h)
// Return the number of elements in the string list.
func StringListSize(list StringList) uint64 {
	return uint64(C.cef_string_list_size(C.cef_string_list_t(list)))
}

// StringListValue (cef_string_list_value from .\include/internal/cef_string_list.h)
// Retrieve the value at the specified zero-based string list index. Returns
// true (1) if the value was successfully retrieved.
func StringListValue(list StringList, index uint64, value *string) int32 {
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(*value, value_)
	defer func() {
		*value = cefstrToString(value_)
		C.cef_string_userfree_free(value_)
	}()
	return int32(C.cef_string_list_value(C.cef_string_list_t(list), C.size_t(index), (*C.cef_string_t)(value_)))
}

// StringMapAlloc (cef_string_map_alloc from .\include/internal/cef_string_map.h)
// Allocate a new string map.
func StringMapAlloc() StringMap {
	return StringMap(C.cef_string_map_alloc())
}

// StringMapAppend (cef_string_map_append from .\include/internal/cef_string_map.h)
// Append a new key/value pair at the end of the string map.
func StringMapAppend(map_r StringMap, key, value string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(value, value_)
	defer func() {
		C.cef_string_userfree_free(value_)
	}()
	return int32(C.cef_string_map_append(C.cef_string_map_t(map_r), (*C.cef_string_t)(key_), (*C.cef_string_t)(value_)))
}

// StringMapClear (cef_string_map_clear from .\include/internal/cef_string_map.h)
// Clear the string map.
func StringMapClear(map_r StringMap) {
	C.cef_string_map_clear(C.cef_string_map_t(map_r))
}

// StringMapFind (cef_string_map_find from .\include/internal/cef_string_map.h)
// Return the value assigned to the specified key.
func StringMapFind(map_r StringMap, key string, value *string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(*value, value_)
	defer func() {
		*value = cefstrToString(value_)
		C.cef_string_userfree_free(value_)
	}()
	return int32(C.cef_string_map_find(C.cef_string_map_t(map_r), (*C.cef_string_t)(key_), (*C.cef_string_t)(value_)))
}

// StringMapFree (cef_string_map_free from .\include/internal/cef_string_map.h)
// Free the string map.
func StringMapFree(map_r StringMap) {
	C.cef_string_map_free(C.cef_string_map_t(map_r))
}

// StringMapKey (cef_string_map_key from .\include/internal/cef_string_map.h)
// Return the key at the specified zero-based string map index.
func StringMapKey(map_r StringMap, index uint64, key *string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(*key, key_)
	defer func() {
		*key = cefstrToString(key_)
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.cef_string_map_key(C.cef_string_map_t(map_r), C.size_t(index), (*C.cef_string_t)(key_)))
}

// StringMapSize (cef_string_map_size from .\include/internal/cef_string_map.h)
// Return the number of elements in the string map.
func StringMapSize(map_r StringMap) uint64 {
	return uint64(C.cef_string_map_size(C.cef_string_map_t(map_r)))
}

// StringMapValue (cef_string_map_value from .\include/internal/cef_string_map.h)
// Return the value at the specified zero-based string map index.
func StringMapValue(map_r StringMap, index uint64, value *string) int32 {
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(*value, value_)
	defer func() {
		*value = cefstrToString(value_)
		C.cef_string_userfree_free(value_)
	}()
	return int32(C.cef_string_map_value(C.cef_string_map_t(map_r), C.size_t(index), (*C.cef_string_t)(value_)))
}

// StringMultimapAlloc (cef_string_multimap_alloc from .\include/internal/cef_string_multimap.h)
// Allocate a new string multimap.
func StringMultimapAlloc() StringMultimap {
	return StringMultimap(C.cef_string_multimap_alloc())
}

// StringMultimapAppend (cef_string_multimap_append from .\include/internal/cef_string_multimap.h)
// Append a new key/value pair at the end of the string multimap.
func StringMultimapAppend(map_r StringMultimap, key, value string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(value, value_)
	defer func() {
		C.cef_string_userfree_free(value_)
	}()
	return int32(C.cef_string_multimap_append(C.cef_string_multimap_t(map_r), (*C.cef_string_t)(key_), (*C.cef_string_t)(value_)))
}

// StringMultimapClear (cef_string_multimap_clear from .\include/internal/cef_string_multimap.h)
// Clear the string multimap.
func StringMultimapClear(map_r StringMultimap) {
	C.cef_string_multimap_clear(C.cef_string_multimap_t(map_r))
}

// StringMultimapEnumerate (cef_string_multimap_enumerate from .\include/internal/cef_string_multimap.h)
// Return the value_index-th value with the specified key.
func StringMultimapEnumerate(map_r StringMultimap, key string, value_index uint64, value *string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(*value, value_)
	defer func() {
		*value = cefstrToString(value_)
		C.cef_string_userfree_free(value_)
	}()
	return int32(C.cef_string_multimap_enumerate(C.cef_string_multimap_t(map_r), (*C.cef_string_t)(key_), C.size_t(value_index), (*C.cef_string_t)(value_)))
}

// StringMultimapFindCount (cef_string_multimap_find_count from .\include/internal/cef_string_multimap.h)
// Return the number of values with the specified key.
func StringMultimapFindCount(map_r StringMultimap, key string) uint64 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return uint64(C.cef_string_multimap_find_count(C.cef_string_multimap_t(map_r), (*C.cef_string_t)(key_)))
}

// StringMultimapFree (cef_string_multimap_free from .\include/internal/cef_string_multimap.h)
// Free the string multimap.
func StringMultimapFree(map_r StringMultimap) {
	C.cef_string_multimap_free(C.cef_string_multimap_t(map_r))
}

// StringMultimapKey (cef_string_multimap_key from .\include/internal/cef_string_multimap.h)
// Return the key at the specified zero-based string multimap index.
func StringMultimapKey(map_r StringMultimap, index uint64, key *string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(*key, key_)
	defer func() {
		*key = cefstrToString(key_)
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.cef_string_multimap_key(C.cef_string_multimap_t(map_r), C.size_t(index), (*C.cef_string_t)(key_)))
}

// StringMultimapSize (cef_string_multimap_size from .\include/internal/cef_string_multimap.h)
// Return the number of elements in the string multimap.
func StringMultimapSize(map_r StringMultimap) uint64 {
	return uint64(C.cef_string_multimap_size(C.cef_string_multimap_t(map_r)))
}

// StringMultimapValue (cef_string_multimap_value from .\include/internal/cef_string_multimap.h)
// Return the value at the specified zero-based string multimap index.
func StringMultimapValue(map_r StringMultimap, index uint64, value *string) int32 {
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(*value, value_)
	defer func() {
		*value = cefstrToString(value_)
		C.cef_string_userfree_free(value_)
	}()
	return int32(C.cef_string_multimap_value(C.cef_string_multimap_t(map_r), C.size_t(index), (*C.cef_string_t)(value_)))
}

// StringUserfreeUtf8Alloc (cef_string_userfree_utf8_alloc from .\include/internal/cef_string_types.h)
func StringUserfreeUtf8Alloc() string {
	return string(C.cef_string_userfree_utf8_alloc())
}

// StringUserfreeUtf8Free (cef_string_userfree_utf8_free from .\include/internal/cef_string_types.h)
func StringUserfreeUtf8Free(str *string) {
	C.cef_string_userfree_utf8_free(C.cef_string_userfree_utf8_t(str))
}

// StringUserfreeUtf16Alloc (cef_string_userfree_utf16_alloc from .\include/internal/cef_string_types.h)
func StringUserfreeUtf16Alloc() string {
	return string(C.cef_string_userfree_utf16_alloc())
}

// StringUserfreeUtf16Free (cef_string_userfree_utf16_free from .\include/internal/cef_string_types.h)
func StringUserfreeUtf16Free(str *string) {
	C.cef_string_userfree_utf16_free(C.cef_string_userfree_utf16_t(str))
}

// StringUserfreeWideAlloc (cef_string_userfree_wide_alloc from .\include/internal/cef_string_types.h)
func StringUserfreeWideAlloc() string {
	return string(C.cef_string_userfree_wide_alloc())
}

// StringUserfreeWideFree (cef_string_userfree_wide_free from .\include/internal/cef_string_types.h)
func StringUserfreeWideFree(str *string) {
	C.cef_string_userfree_wide_free(C.cef_string_userfree_wide_t(str))
}

// StringUtf8Clear (cef_string_utf8_clear from .\include/internal/cef_string_types.h)
func StringUtf8Clear(str *string) {
	C.cef_string_utf8_clear((*C.cef_string_utf8_t)(str))
}

// StringUtf8Cmp (cef_string_utf8_cmp from .\include/internal/cef_string_types.h)
func StringUtf8Cmp(str1, str2 string) int32 {
	return int32(C.cef_string_utf8_cmp((*C.cef_string_utf8_t)(str1), (*C.cef_string_utf8_t)(str2)))
}

// StringUtf8Set (cef_string_utf8_set from .\include/internal/cef_string_types.h)
func StringUtf8Set(src string, src_len uint64, output *string, copy int32) int32 {
	return int32(C.cef_string_utf8_set((*C.char)(src), C.size_t(src_len), (*C.cef_string_utf8_t)(output), C.int(copy)))
}

// StringUtf8ToUtf16 (cef_string_utf8_to_utf16 from .\include/internal/cef_string_types.h)
func StringUtf8ToUtf16(src string, src_len uint64, output *string) int32 {
	return int32(C.cef_string_utf8_to_utf16((*C.char)(src), C.size_t(src_len), (*C.cef_string_utf16_t)(output)))
}

// StringUtf8ToWide (cef_string_utf8_to_wide from .\include/internal/cef_string_types.h)
func StringUtf8ToWide(src string, src_len uint64, output *string) int32 {
	return int32(C.cef_string_utf8_to_wide((*C.char)(src), C.size_t(src_len), (*C.cef_string_wide_t)(output)))
}

// StringUtf16Clear (cef_string_utf16_clear from .\include/internal/cef_string_types.h)
func StringUtf16Clear(str *string) {
	C.cef_string_utf16_clear((*C.cef_string_utf16_t)(str))
}

// StringUtf16Cmp (cef_string_utf16_cmp from .\include/internal/cef_string_types.h)
func StringUtf16Cmp(str1, str2 string) int32 {
	return int32(C.cef_string_utf16_cmp((*C.cef_string_utf16_t)(str1), (*C.cef_string_utf16_t)(str2)))
}

// StringUtf16Set (cef_string_utf16_set from .\include/internal/cef_string_types.h)
func StringUtf16Set(src *int16, src_len uint64, output *string, copy int32) int32 {
	return int32(C.cef_string_utf16_set((*C.char16)(src), C.size_t(src_len), (*C.cef_string_utf16_t)(output), C.int(copy)))
}

// StringUtf16ToLower (cef_string_utf16_to_lower from .\include/internal/cef_string_types.h)
func StringUtf16ToLower(src *int16, src_len uint64, output *string) int32 {
	return int32(C.cef_string_utf16_to_lower((*C.char16)(src), C.size_t(src_len), (*C.cef_string_utf16_t)(output)))
}

// StringUtf16ToUpper (cef_string_utf16_to_upper from .\include/internal/cef_string_types.h)
func StringUtf16ToUpper(src *int16, src_len uint64, output *string) int32 {
	return int32(C.cef_string_utf16_to_upper((*C.char16)(src), C.size_t(src_len), (*C.cef_string_utf16_t)(output)))
}

// StringUtf16ToUtf8 (cef_string_utf16_to_utf8 from .\include/internal/cef_string_types.h)
func StringUtf16ToUtf8(src *int16, src_len uint64, output *string) int32 {
	return int32(C.cef_string_utf16_to_utf8((*C.char16)(src), C.size_t(src_len), (*C.cef_string_utf8_t)(output)))
}

// StringUtf16ToWide (cef_string_utf16_to_wide from .\include/internal/cef_string_types.h)
func StringUtf16ToWide(src *int16, src_len uint64, output *string) int32 {
	return int32(C.cef_string_utf16_to_wide((*C.char16)(src), C.size_t(src_len), (*C.cef_string_wide_t)(output)))
}

// StringWideClear (cef_string_wide_clear from .\include/internal/cef_string_types.h)
func StringWideClear(str *string) {
	C.cef_string_wide_clear((*C.cef_string_wide_t)(str))
}

// StringWideCmp (cef_string_wide_cmp from .\include/internal/cef_string_types.h)
func StringWideCmp(str1, str2 string) int32 {
	return int32(C.cef_string_wide_cmp((*C.cef_string_wide_t)(str1), (*C.cef_string_wide_t)(str2)))
}

// StringWideSet (cef_string_wide_set from .\include/internal/cef_string_types.h)
func StringWideSet(src *wchar_t, src_len uint64, output *string, copy int32) int32 {
	return int32(C.cef_string_wide_set((*C.wchar_t)(src), C.size_t(src_len), (*C.cef_string_wide_t)(output), C.int(copy)))
}

// StringWideToUtf8 (cef_string_wide_to_utf8 from .\include/internal/cef_string_types.h)
func StringWideToUtf8(src *wchar_t, src_len uint64, output *string) int32 {
	return int32(C.cef_string_wide_to_utf8((*C.wchar_t)(src), C.size_t(src_len), (*C.cef_string_utf8_t)(output)))
}

// StringWideToUtf16 (cef_string_wide_to_utf16 from .\include/internal/cef_string_types.h)
func StringWideToUtf16(src *wchar_t, src_len uint64, output *string) int32 {
	return int32(C.cef_string_wide_to_utf16((*C.wchar_t)(src), C.size_t(src_len), (*C.cef_string_utf16_t)(output)))
}

// TaskRunnerGetForCurrentThread (cef_task_runner_get_for_current_thread from .\include/capi/cef_task_capi.h)
// Returns the task runner for the current thread. Only CEF threads will have
// task runners. An NULL reference will be returned if this function is called
// on an invalid thread.
func TaskRunnerGetForCurrentThread() *TaskRunner {
	return (*TaskRunner)(C.cef_task_runner_get_for_current_thread())
}

// TaskRunnerGetForThread (cef_task_runner_get_for_thread from .\include/capi/cef_task_capi.h)
// Returns the task runner for the specified CEF thread.
func TaskRunnerGetForThread(threadId ThreadID) *TaskRunner {
	return (*TaskRunner)(C.cef_task_runner_get_for_thread(C.cef_thread_id_t(threadId)))
}

// TextfieldCreate (cef_textfield_create from include\capi\views\cef_textfield_capi.h)
// Create a new Textfield.
func TextfieldCreate(delegate *TextfieldDelegate) *Textfield {
	var delegate_ *C.cef_textfield_delegate_t
	if delegate != nil {
		delegate_ = delegate.toNative()
	}
	return (*Textfield)(C.cef_textfield_create(delegate_))
}

// TimeDelta (cef_time_delta from .\include/internal/cef_time.h)
// Retrieve the delta in milliseconds between two time values.
//
func TimeDelta(cef_time1, cef_time2 *Time, delta *int64) int32 {
	return int32(C.cef_time_delta(cef_time1.toNative(&C.cef_time_t{}), cef_time2.toNative(&C.cef_time_t{}), (*C.longlong)(delta)))
}

// TimeFromDoublet (cef_time_from_doublet from .\include/internal/cef_time.h)
func TimeFromDoublet(time float64, cef_time *Time) int32 {
	return int32(C.cef_time_from_doublet(C.double(time), cef_time.toNative(&C.cef_time_t{})))
}

// TimeFromTimet (cef_time_from_timet from .\include/internal/cef_time.h)
func TimeFromTimet(time int64, cef_time *Time) int32 {
	return int32(C.cef_time_from_timet(C.time_t(time), cef_time.toNative(&C.cef_time_t{})))
}

// TimeNow (cef_time_now from .\include/internal/cef_time.h)
// Retrieve the current system time.
//
func TimeNow(cef_time *Time) int32 {
	return int32(C.cef_time_now(cef_time.toNative(&C.cef_time_t{})))
}

// TimeToDoublet (cef_time_to_doublet from .\include/internal/cef_time.h)
// Converts cef_time_t to/from a double which is the number of seconds since
// epoch (Jan 1, 1970). Webkit uses this format to represent time. A value of 0
// means "not initialized". Returns true (1) on success and false (0) on
// failure.
func TimeToDoublet(cef_time *Time, time *float64) int32 {
	return int32(C.cef_time_to_doublet(cef_time.toNative(&C.cef_time_t{}), (*C.double)(time)))
}

// TimeToTimet (cef_time_to_timet from .\include/internal/cef_time.h)
// Converts cef_time_t to/from time_t. Returns true (1) on success and false (0)
// on failure.
func TimeToTimet(cef_time *Time, time *int64) int32 {
	return int32(C.cef_time_to_timet(cef_time.toNative(&C.cef_time_t{}), (*C.time_t)(time)))
}

// UnregisterInternalWebPlugin (cef_unregister_internal_web_plugin from .\include/capi/cef_web_plugin_capi.h)
// Unregister an internal plugin. This may be undone the next time
// cef_refresh_web_plugins() is called. Can be called on any thread in the
// browser process.
func UnregisterInternalWebPlugin(path string) {
	path_ := C.cef_string_userfree_alloc()
	setCEFStr(path, path_)
	defer func() {
		C.cef_string_userfree_free(path_)
	}()
	C.cef_unregister_internal_web_plugin((*C.cef_string_t)(path_))
}

// UrlrequestCreate (cef_urlrequest_create from include\capi\cef_urlrequest_capi.h)
// Create a new URL request. Only GET, POST, HEAD, DELETE and PUT request
// functions are supported. Multiple post data elements are not supported and
// elements of type PDE_TYPE_FILE are only supported for requests originating
// from the browser process. Requests originating from the render process will
// receive the same handling as requests originating from Web content -- if the
// response contains Content-Disposition or Mime-Type header values that would
// not normally be rendered then the response may receive special handling
// inside the browser (for example, via the file download code path instead of
// the URL request code path). The |request| object will be marked as read-only
// after calling this function. In the browser process if |request_context| is
// NULL the global request context will be used. In the render process
// |request_context| must be NULL and the context associated with the current
// renderer process' browser will be used.
func UrlrequestCreate(request *Request, client *UrlrequestClient, request_context *RequestContext) *Urlrequest {
	return (*Urlrequest)(C.cef_urlrequest_create(request.toNative(), client.toNative(), request_context.toNative()))
}

// V8contextGetCurrentContext (cef_v8context_get_current_context from .\include/capi/cef_v8_capi.h)
// Returns the current (top) context object in the V8 context stack.
func V8contextGetCurrentContext() *V8context {
	return (*V8context)(C.cef_v8context_get_current_context())
}

// V8contextGetEnteredContext (cef_v8context_get_entered_context from .\include/capi/cef_v8_capi.h)
// Returns the entered (bottom) context object in the V8 context stack.
func V8contextGetEnteredContext() *V8context {
	return (*V8context)(C.cef_v8context_get_entered_context())
}

// V8contextInContext (cef_v8context_in_context from .\include/capi/cef_v8_capi.h)
// Returns true (1) if V8 is currently inside a context.
func V8contextInContext() int32 {
	return int32(C.cef_v8context_in_context())
}

// V8stackTraceGetCurrent (cef_v8stack_trace_get_current from .\include/capi/cef_v8_capi.h)
// Returns the stack trace for the currently active context. |frame_limit| is
// the maximum number of frames that will be captured.
func V8stackTraceGetCurrent(frame_limit int32) *V8stackTrace {
	return (*V8stackTrace)(C.cef_v8stack_trace_get_current(C.int(frame_limit)))
}

// V8valueCreateArray (cef_v8value_create_array from .\include/capi/cef_v8_capi.h)
// Create a new cef_v8value_t object of type array with the specified |length|.
// If |length| is negative the returned array will have length 0. This function
// should only be called from within the scope of a
// cef_render_process_handler_t, cef_v8handler_t or cef_v8accessor_t callback,
// or in combination with calling enter() and exit() on a stored cef_v8context_t
// reference.
func V8valueCreateArray(length int32) *V8value {
	return (*V8value)(C.cef_v8value_create_array(C.int(length)))
}

// V8valueCreateArrayBuffer (cef_v8value_create_array_buffer from .\include/capi/cef_v8_capi.h)
// Create a new cef_v8value_t object of type ArrayBuffer which wraps the
// provided |buffer| of size |length| bytes. The ArrayBuffer is externalized,
// meaning that it does not own |buffer|. The caller is responsible for freeing
// |buffer| when requested via a call to cef_v8array_buffer_release_callback_t::
// ReleaseBuffer. This function should only be called from within the scope of a
// cef_render_process_handler_t, cef_v8handler_t or cef_v8accessor_t callback,
// or in combination with calling enter() and exit() on a stored cef_v8context_t
// reference.
func V8valueCreateArrayBuffer(buffer unsafe.Pointer, length uint64, release_callback *V8arrayBufferReleaseCallback) *V8value {
	return (*V8value)(C.cef_v8value_create_array_buffer(buffer, C.size_t(length), release_callback.toNative()))
}

// V8valueCreateBool (cef_v8value_create_bool from .\include/capi/cef_v8_capi.h)
// Create a new cef_v8value_t object of type bool.
func V8valueCreateBool(value int32) *V8value {
	return (*V8value)(C.cef_v8value_create_bool(C.int(value)))
}

// V8valueCreateDate (cef_v8value_create_date from .\include/capi/cef_v8_capi.h)
// Create a new cef_v8value_t object of type Date. This function should only be
// called from within the scope of a cef_render_process_handler_t,
// cef_v8handler_t or cef_v8accessor_t callback, or in combination with calling
// enter() and exit() on a stored cef_v8context_t reference.
func V8valueCreateDate(date *Time) *V8value {
	return (*V8value)(C.cef_v8value_create_date(date.toNative(&C.cef_time_t{})))
}

// V8valueCreateDouble (cef_v8value_create_double from .\include/capi/cef_v8_capi.h)
// Create a new cef_v8value_t object of type double.
func V8valueCreateDouble(value float64) *V8value {
	return (*V8value)(C.cef_v8value_create_double(C.double(value)))
}

// V8valueCreateFunction (cef_v8value_create_function from .\include/capi/cef_v8_capi.h)
// Create a new cef_v8value_t object of type function. This function should only
// be called from within the scope of a cef_render_process_handler_t,
// cef_v8handler_t or cef_v8accessor_t callback, or in combination with calling
// enter() and exit() on a stored cef_v8context_t reference.
func V8valueCreateFunction(name string, handler *V8handler) *V8value {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	return (*V8value)(C.cef_v8value_create_function((*C.cef_string_t)(name_), handler.toNative()))
}

// V8valueCreateInt (cef_v8value_create_int from .\include/capi/cef_v8_capi.h)
// Create a new cef_v8value_t object of type int.
func V8valueCreateInt(value int32) *V8value {
	return (*V8value)(C.cef_v8value_create_int(C.int32(value)))
}

// V8valueCreateNull (cef_v8value_create_null from .\include/capi/cef_v8_capi.h)
// Create a new cef_v8value_t object of type null.
func V8valueCreateNull() *V8value {
	return (*V8value)(C.cef_v8value_create_null())
}

// V8valueCreateObject (cef_v8value_create_object from .\include/capi/cef_v8_capi.h)
// Create a new cef_v8value_t object of type object with optional accessor
// and/or interceptor. This function should only be called from within the scope
// of a cef_render_process_handler_t, cef_v8handler_t or cef_v8accessor_t
// callback, or in combination with calling enter() and exit() on a stored
// cef_v8context_t reference.
func V8valueCreateObject(accessor *V8accessor, interceptor *V8interceptor) *V8value {
	return (*V8value)(C.cef_v8value_create_object(accessor.toNative(), interceptor.toNative()))
}

// V8valueCreateString (cef_v8value_create_string from .\include/capi/cef_v8_capi.h)
// Create a new cef_v8value_t object of type string.
func V8valueCreateString(value string) *V8value {
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(value, value_)
	defer func() {
		C.cef_string_userfree_free(value_)
	}()
	return (*V8value)(C.cef_v8value_create_string((*C.cef_string_t)(value_)))
}

// V8valueCreateUint (cef_v8value_create_uint from .\include/capi/cef_v8_capi.h)
// Create a new cef_v8value_t object of type unsigned int.
func V8valueCreateUint(value uint32) *V8value {
	return (*V8value)(C.cef_v8value_create_uint(C.uint32(value)))
}

// V8valueCreateUndefined (cef_v8value_create_undefined from .\include/capi/cef_v8_capi.h)
// Create a new cef_v8value_t object of type undefined.
func V8valueCreateUndefined() *V8value {
	return (*V8value)(C.cef_v8value_create_undefined())
}

// ValueCreate (cef_value_create from .\include/capi/cef_values_capi.h)
// Creates a new object.
func ValueCreate() *Value {
	return (*Value)(C.cef_value_create())
}

// VisitWebPluginInfo (cef_visit_web_plugin_info from .\include/capi/cef_web_plugin_capi.h)
// Visit web plugin information. Can be called on any thread in the browser
// process.
func VisitWebPluginInfo(visitor *WebPluginInfoVisitor) {
	C.cef_visit_web_plugin_info(visitor.toNative())
}

// WaitableEventCreate (cef_waitable_event_create from include\capi\cef_waitable_event_capi.h)
// Create a new waitable event. If |automatic_reset| is true (1) then the event
// state is automatically reset to un-signaled after a single waiting thread has
// been released; otherwise, the state remains signaled until reset() is called
// manually. If |initially_signaled| is true (1) then the event will start in
// the signaled state.
func WaitableEventCreate(automatic_reset, initially_signaled int32) *WaitableEvent {
	return (*WaitableEvent)(C.cef_waitable_event_create(C.int(automatic_reset), C.int(initially_signaled)))
}

// WindowCreateTopLevel (cef_window_create_top_level from include\capi\views\cef_window_capi.h)
// Create a new Window.
func WindowCreateTopLevel(delegate *WindowDelegate) *Window {
	var delegate_ *C.cef_window_delegate_t
	if delegate != nil {
		delegate_ = delegate.toNative()
	}
	return (*Window)(C.cef_window_create_top_level(delegate_))
}
