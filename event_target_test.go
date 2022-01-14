package events

import (
	"reflect"
	"testing"
	"time"
)

func TestNewEventTarget(t *testing.T) {
	target:=NewEventTarget()
	if target==nil {
		t.Fail()
	}
}

func TestEvenTarget_On(t *testing.T) {
	target:=NewEventTarget()
	err:=target.On(func() {})
	if err!=nil {
		t.Fatal(err)
	}
	if target.ListenerCount()!=1 {
		t.Fatal()
	}
}

func TestEvenTarget_Off(t *testing.T) {
	target:=NewEventTarget()
	listener:= func() {}
	err:=target.On(listener)
	if err!=nil {
		t.Fatal(err)
	}
	err =target.Off(listener)
	if err!=nil {
		t.Fatal(err)
	}
	if target.ListenerCount()!=0 {
		t.Fatal()
	}
}

func TestEvenTarget_Once(t *testing.T) {
	target:=NewEventTarget()
	listener:= func() {}
	err:=target.Once(listener)
	if err!=nil {
		t.Fatal(err)
	}
	target.Emit()
	if target.ListenerCount()!=0 {
		t.Fatal()
	}
}

func TestEventTarget_Emit(t *testing.T) {
	target:=NewEventTarget()
	end:=make(chan struct{})
	err:=target.On(func() {
		go func() {
			end<- struct{}{}
		}()
	})
	if err!=nil {
		t.Fatal(err)
	}
	target.Emit()
	select {
	case <-end:
	case <-time.After(time.Second*1):
		t.Fatal()
	}
}

func TestEvenTarget_Listeners(t *testing.T) {
	target:=NewEventTarget()
	listener:= func() {}
	err:=target.On(listener)
	if err!=nil {
		t.Fatal(err)
	}
	listeners:=target.Listeners()
	if reflect.ValueOf(listeners[0]).Pointer()!=reflect.ValueOf(listener).Pointer() {
		t.Fatal()
	}
}



