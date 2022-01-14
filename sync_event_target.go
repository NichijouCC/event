package events

import (
	"reflect"
	"sync"
	"sync/atomic"
)

type SyncEventTarget struct {
	lisMu         sync.RWMutex
	listeners     []reflect.Value
	onceLen       int32
	onceMu        sync.Mutex
	onceListeners []reflect.Value
}

func NewSyncEventTarget() *SyncEventTarget {
	return &SyncEventTarget{
		listeners:     []reflect.Value{},
		onceListeners: []reflect.Value{},
	}
}

func (s *SyncEventTarget) On(listener interface{}) error {
	if err := isValidListener(listener); err != nil {
		return err
	}
	s.lisMu.Lock()
	defer s.lisMu.Unlock()
	s.listeners = append(s.listeners, reflect.ValueOf(listener))
	return nil
}

func (s *SyncEventTarget) Once(listener interface{}) error {
	if err := isValidListener(listener); err != nil {
		return err
	}
	s.onceMu.Lock()
	defer s.onceMu.Unlock()
	s.onceListeners = append(s.onceListeners, reflect.ValueOf(listener))
	atomic.AddInt32(&s.onceLen, 1)
	return nil
}

func (s *SyncEventTarget) Off(listener interface{}) error {
	if err := isValidListener(listener); err != nil {
		return err
	}
	rv := reflect.ValueOf(listener)
	func() {
		s.lisMu.Lock()
		defer s.lisMu.Unlock()
		for index, el := range s.listeners {
			if el.Pointer() == rv.Pointer() {
				s.listeners = append(s.listeners[:index], s.listeners[index+1:]...)
				break
			}
		}
	}()

	if atomic.LoadInt32(&s.onceLen) > 0 {
		func() {
			s.onceMu.Lock()
			defer s.onceMu.Unlock()
			for index, el := range s.onceListeners {
				if el.Pointer() == rv.Pointer() {
					s.onceListeners = append(s.onceListeners[:index], s.onceListeners[index+1:]...)
					atomic.AddInt32(&s.onceLen, -1)
					break
				}
			}
		}()
	}
	return nil
}

func (s *SyncEventTarget) Emit(args ...interface{}) {
	reflectedArgs := argsToReflectValues(args...)

	func() {
		s.lisMu.RLock()
		defer s.lisMu.RUnlock()
		for _, el := range s.listeners {
			el.Call(reflectedArgs)
		}
	}()

	if atomic.LoadInt32(&s.onceLen) > 0 {
		func(){
			s.onceMu.Lock()
			defer s.onceMu.Unlock()
			for _, el := range s.onceListeners {
				el.Call(reflectedArgs)
			}
			s.onceListeners = nil
			atomic.StoreInt32(&s.onceLen,0)
		}()
	}
}

func (s *SyncEventTarget) Listeners() []interface{} {
	var arr []interface{}
	func(){
		s.lisMu.RLock()
		defer s.lisMu.RUnlock()
		for _, el := range s.listeners {
			arr = append(arr, el.Interface())
		}
	}()

	if atomic.LoadInt32(&s.onceLen) > 0 {
		func(){
			s.onceMu.Lock()
			defer s.onceMu.Unlock()
			for _, el := range s.onceListeners {
				arr = append(arr, el.Interface())
			}
		}()
	}
	return arr
}

func (s *SyncEventTarget) ListenerCount() int {
	s.lisMu.RLock()
	defer s.lisMu.RUnlock()
	return len(s.listeners)+int(atomic.LoadInt32(&s.onceLen))
}