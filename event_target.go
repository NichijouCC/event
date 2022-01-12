package events

import "reflect"

type IEventTarget interface {
	On(listener interface{}) error
	Once(listener interface{}) error
	Off(listener interface{}) error
	Emit(args ...interface{})
	Listeners() []interface{}
	ListenerCount() int
}

type EvenTarget struct {
	listeners []*eventListener
}


func NewEventTarget() *EvenTarget {
	return &EvenTarget{
		listeners: []*eventListener{},
	}
}

func (e *EvenTarget) On(listener interface{}) error {
	if err := isValidListener(listener); err != nil {
		return err
	}
	e.listeners = append(e.listeners, &eventListener{callback: reflect.ValueOf(listener)})
	return nil
}

func (e *EvenTarget) Once(listener interface{}) error {
	if err := isValidListener(listener); err != nil {
		return err
	}
	e.listeners = append(e.listeners, &eventListener{callback: reflect.ValueOf(listener)})
	return nil
}

func (e *EvenTarget) Off(listener interface{}) error {
	if err := isValidListener(listener); err != nil {
		return err
	}
	rv := reflect.ValueOf(listener)
	for index, el := range e.listeners {
		if el.callback == rv {
			e.listeners = append(e.listeners[:index], e.listeners[index+1:]...)
			return nil
		}
	}
	return nil
}

func (e *EvenTarget) Emit(args ...interface{}) {
	reflectedArgs := argsToReflectValues(args)
	for i := 0; i < len(e.listeners); i++ {
		e.listeners[i].callback.Call(reflectedArgs)
		if e.listeners[i].once {
			e.listeners = append(e.listeners[:i], e.listeners[i+1:]...)
			i--
		}
	}
}

func (e *EvenTarget) Listeners() []interface{} {
	var arr []interface{}
	for _, el := range e.listeners {
		arr = append(arr, el.callback.Interface())
	}
	return arr
}

func (e *EvenTarget) ListenerCount() int {
	return len(e.listeners)
}
