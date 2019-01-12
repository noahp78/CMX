// generated by go-extpoints -- DO NOT EDIT
package extpoints

import (
	"reflect"
	"runtime"
	"strings"
	"sync"
)

var extRegistry = &registryType{m: make(map[string]*extensionPoint)}

type registryType struct {
	sync.Mutex
	m map[string]*extensionPoint
}

// Top level registration

func extensionTypes(extension interface{}) []string {
	var ifaces []string
	typ := reflect.TypeOf(extension)
	for name, ep := range extRegistry.m {
		if ep.iface.Kind() == reflect.Func && typ.AssignableTo(ep.iface) {
			ifaces = append(ifaces, name)
		}
		if ep.iface.Kind() != reflect.Func && typ.Implements(ep.iface) {
			ifaces = append(ifaces, name)
		}
	}
	return ifaces
}

func RegisterExtension(extension interface{}, name string) []string {
	extRegistry.Lock()
	defer extRegistry.Unlock()
	var ifaces []string
	for _, iface := range extensionTypes(extension) {
		if extRegistry.m[iface].register(extension, name) {
			ifaces = append(ifaces, iface)
		}
	}
	return ifaces
}

func UnregisterExtension(name string) []string {
	extRegistry.Lock()
	defer extRegistry.Unlock()
	var ifaces []string
	for iface, extpoint := range extRegistry.m {
		if extpoint.unregister(name) {
			ifaces = append(ifaces, iface)
		}
	}
	return ifaces
}


// Base extension point

type extensionPoint struct {
	sync.Mutex
	iface      reflect.Type
	extensions map[string]interface{}
}

func newExtensionPoint(iface interface{}) *extensionPoint {
	ep := &extensionPoint{
		iface:      reflect.TypeOf(iface).Elem(),
		extensions: make(map[string]interface{}),
	}
	extRegistry.Lock()
	extRegistry.m[ep.iface.Name()] = ep
	extRegistry.Unlock()
	return ep
}

func (ep *extensionPoint) lookup(name string) interface{} {
	ep.Lock()
	defer ep.Unlock()
	ext, ok := ep.extensions[name]
	if !ok {
		return nil
	}
	return ext
}

func (ep *extensionPoint) all() map[string]interface{} {
	ep.Lock()
	defer ep.Unlock()
	all := make(map[string]interface{})
	for k, v := range ep.extensions {
		all[k] = v
	}
	return all
}

func (ep *extensionPoint) register(extension interface{}, name string) bool {
	ep.Lock()
	defer ep.Unlock()
	if name == "" {
		typ := reflect.TypeOf(extension)
		if typ.Kind() == reflect.Func {
			nameParts := strings.Split(runtime.FuncForPC(
				reflect.ValueOf(extension).Pointer()).Name(), ".")
			name = nameParts[len(nameParts)-1]
		} else {
			name = typ.Elem().Name()
		}
	}
	_, exists := ep.extensions[name]
	if exists {
		return false
	}
	ep.extensions[name] = extension
	return true
}

func (ep *extensionPoint) unregister(name string) bool {
	ep.Lock()
	defer ep.Unlock()
	_, exists := ep.extensions[name]
	if !exists {
		return false
	}
	delete(ep.extensions, name)
	return true
}

// Addon

var Addons = &addonExt{
	newExtensionPoint(new(Addon)),
}

type addonExt struct {
	*extensionPoint
}

func (ep *addonExt) Unregister(name string) bool {
	return ep.unregister(name)
}

func (ep *addonExt) Register(extension Addon, name string) bool {
	return ep.register(extension, name)
}

func (ep *addonExt) Lookup(name string) Addon {
	ext := ep.lookup(name)
	if ext == nil {
		return nil
	}
	return ext.(Addon)
}

func (ep *addonExt) Select(names []string) []Addon {
	var selected []Addon
	for _, name := range names {
		selected = append(selected, ep.Lookup(name))
	}
	return selected
}

func (ep *addonExt) All() map[string]Addon {
	all := make(map[string]Addon)
	for k, v := range ep.all() {
		all[k] = v.(Addon)
	}
	return all
}

func (ep *addonExt) Names() []string {
	var names []string
	for k := range ep.all() {
		names = append(names, k)
	}
	return names
}


// EventData

var EventDatas = &eventDataExt{
	newExtensionPoint(new(EventData)),
}

type eventDataExt struct {
	*extensionPoint
}

func (ep *eventDataExt) Unregister(name string) bool {
	return ep.unregister(name)
}

func (ep *eventDataExt) Register(extension EventData, name string) bool {
	return ep.register(extension, name)
}

func (ep *eventDataExt) Lookup(name string) EventData {
	ext := ep.lookup(name)
	if ext == nil {
		return nil
	}
	return ext.(EventData)
}

func (ep *eventDataExt) Select(names []string) []EventData {
	var selected []EventData
	for _, name := range names {
		selected = append(selected, ep.Lookup(name))
	}
	return selected
}

func (ep *eventDataExt) All() map[string]EventData {
	all := make(map[string]EventData)
	for k, v := range ep.all() {
		all[k] = v.(EventData)
	}
	return all
}

func (ep *eventDataExt) Names() []string {
	var names []string
	for k := range ep.all() {
		names = append(names, k)
	}
	return names
}


// EventListener

var EventListeners = &eventListenerExt{
	newExtensionPoint(new(EventListener)),
}

type eventListenerExt struct {
	*extensionPoint
}

func (ep *eventListenerExt) Unregister(name string) bool {
	return ep.unregister(name)
}

func (ep *eventListenerExt) Register(extension EventListener, name string) bool {
	return ep.register(extension, name)
}

func (ep *eventListenerExt) Lookup(name string) EventListener {
	ext := ep.lookup(name)
	if ext == nil {
		return nil
	}
	return ext.(EventListener)
}

func (ep *eventListenerExt) Select(names []string) []EventListener {
	var selected []EventListener
	for _, name := range names {
		selected = append(selected, ep.Lookup(name))
	}
	return selected
}

func (ep *eventListenerExt) All() map[string]EventListener {
	all := make(map[string]EventListener)
	for k, v := range ep.all() {
		all[k] = v.(EventListener)
	}
	return all
}

func (ep *eventListenerExt) Names() []string {
	var names []string
	for k := range ep.all() {
		names = append(names, k)
	}
	return names
}


// Route

var Routes = &routeExt{
	newExtensionPoint(new(Route)),
}

type routeExt struct {
	*extensionPoint
}

func (ep *routeExt) Unregister(name string) bool {
	return ep.unregister(name)
}

func (ep *routeExt) Register(extension Route, name string) bool {
	return ep.register(extension, name)
}

func (ep *routeExt) Lookup(name string) Route {
	ext := ep.lookup(name)
	if ext == nil {
		return nil
	}
	return ext.(Route)
}

func (ep *routeExt) Select(names []string) []Route {
	var selected []Route
	for _, name := range names {
		selected = append(selected, ep.Lookup(name))
	}
	return selected
}

func (ep *routeExt) All() map[string]Route {
	all := make(map[string]Route)
	for k, v := range ep.all() {
		all[k] = v.(Route)
	}
	return all
}

func (ep *routeExt) Names() []string {
	var names []string
	for k := range ep.all() {
		names = append(names, k)
	}
	return names
}


