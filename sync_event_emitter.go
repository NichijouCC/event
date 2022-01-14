package events

import (
	"sync"
)

type SyncEventEmitter struct {
	listeners   sync.Map
}

func NewSyncEventEmitter() *SyncEventEmitter {
	return &SyncEventEmitter{}
}

type wrapSyncEventTarget struct {
	*SyncEventTarget
}

func (t *wrapSyncEventTarget) Init()  {
	t.SyncEventTarget=NewSyncEventTarget()
}


func (s *SyncEventEmitter) findOrCreateTarget(event string) *wrapSyncEventTarget {
	target,loaded:=s.listeners.LoadOrStore(event,&wrapSyncEventTarget{})
	if !loaded {
		target.(*wrapSyncEventTarget).Init()
	}
	return target.(*wrapSyncEventTarget)
}

func (s *SyncEventEmitter) AddListener(event string, listener interface{}) error {
	err := isValidListener(listener)
	if err != nil {
		return err
	}
	s.findOrCreateTarget(event).On(listener)
	return nil
}

func (s *SyncEventEmitter) On(event string, listener interface{}) error {
	return s.AddListener(event, listener)
}

func (s *SyncEventEmitter) Once(event string, listener interface{}) error {
	err := isValidListener(listener)
	if err != nil {
		return err
	}
	s.findOrCreateTarget(event).Once(listener)
	return nil
}

func (s *SyncEventEmitter) RemoveListener(event string, listener interface{}) error {
	err := isValidListener(listener)
	if err != nil {
		return err
	}
	if target,ok:=s.listeners.Load(event);ok {
		target.(*wrapSyncEventTarget).Off(listener)
	}
	return nil
}

func (s *SyncEventEmitter) Off(event string, listener interface{}) error {
	return s.RemoveListener(event, listener)
}

func (s *SyncEventEmitter) EventNames() []string {
	var keys []string
	s.listeners.Range(func(key, value interface{}) bool {
		keys=append(keys,key.(string))
		return true
	})
	return keys
}

func (s *SyncEventEmitter) RemoveAllListeners(events ...string) {
	if len(events)>0 {
		for _,el:=range events{
			s.listeners.Delete(el)
		}
	}else {
		s.listeners.Range(func(key interface{}, value interface{}) bool {
			s.listeners.Delete(key)
			return true
		})
	}
}

func (s *SyncEventEmitter) Emit(event string, args ...interface{}) {
	if target,ok:=s.listeners.Load(event);ok {
		target.(*wrapSyncEventTarget).Emit(args...)
	}
}

func (s *SyncEventEmitter) Listeners(event string) []interface{} {
	if target,ok:=s.listeners.Load(event);ok {
		return target.(*wrapSyncEventTarget).Listeners()
	}
	return []interface{}{}
}

func (s *SyncEventEmitter) ListenerCount(event string) int {
	if target,ok:=s.listeners.Load(event);ok {
		return target.(*wrapSyncEventTarget).ListenerCount()
	}
	return 0
}
