package events

import (
	"testing"
	"time"
)

func TestNewSyncEventTarget(t *testing.T) {
	target:=NewSyncEventTarget()
	if target==nil {
		t.Failed()
	}
}

func TestSyncEventTarget_On(t *testing.T) {
	target:=NewSyncEventTarget()
	target.On(func() {})
	if target.ListenerCount()!=1 {
		t.Failed()
	}
}

func TestSyncEventTarget_Once(t *testing.T) {
	target:=NewSyncEventTarget()
	target.Once(func() {})
	if target.ListenerCount()!=1 {
		t.Failed()
	}
}

func TestSyncEventTarget_ListenerCount(t *testing.T) {
	target:=NewSyncEventTarget()
	if target.ListenerCount()!=0 {
		t.Failed()
	}
	target.On(func() {})
	if target.ListenerCount()!=1 {
		t.Failed()
	}
}

func TestSyncEventTarget_Off(t *testing.T) {
	target:=NewSyncEventTarget()
	listener:=func() {}
	target.On(listener)
	target.Off(listener)
	if target.ListenerCount()!=0 {
		t.Failed()
	}
}

func TestSyncEventTarget_Emit(t *testing.T) {
	target:=NewSyncEventTarget()
	finishCh:=make(chan struct{},1)
	target.On(func() {
		finishCh<- struct{}{}
	})
	target.Emit()
	select {
	case <-finishCh:
	case <-time.After(time.Second):
		t.Failed()
	}
}