// Code generated - DO NOT EDIT.

#include "RenderHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_render_handler_proxy(cef_render_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->get_accessibility_handler = (void *)&gocef_render_handler_get_accessibility_handler;
	self->get_root_screen_rect = (void *)&gocef_render_handler_get_root_screen_rect;
	self->get_view_rect = (void *)&gocef_render_handler_get_view_rect;
	self->get_screen_point = (void *)&gocef_render_handler_get_screen_point;
	self->get_screen_info = (void *)&gocef_render_handler_get_screen_info;
	self->on_popup_show = (void *)&gocef_render_handler_on_popup_show;
	self->on_popup_size = (void *)&gocef_render_handler_on_popup_size;
	self->on_paint = (void *)&gocef_render_handler_on_paint;
	self->on_cursor_change = (void *)&gocef_render_handler_on_cursor_change;
	self->start_dragging = (void *)&gocef_render_handler_start_dragging;
	self->update_drag_cursor = (void *)&gocef_render_handler_update_drag_cursor;
	self->on_scroll_offset_changed = (void *)&gocef_render_handler_on_scroll_offset_changed;
	self->on_ime_composition_range_changed = (void *)&gocef_render_handler_on_ime_composition_range_changed;
	self->on_text_selection_changed = (void *)&gocef_render_handler_on_text_selection_changed;
}
