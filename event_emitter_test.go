package events

import (
	"sync"
	"testing"
	"time"
)

func TestNewEventEmitter(t *testing.T) {
	emitter:=NewEventEmitter()
	if emitter==nil {
		t.Fail()
	}
}

func TestEventEmitter_On(t *testing.T) {
	emitter:=NewEventEmitter()
	emitter.On("ev", func() {})
	if len(emitter.Listeners("ev"))!=1 {
		t.Failed()
	}
	emitter.On("ev", func() {})
	if len(emitter.Listeners("ev"))!=2 {
		t.Failed()
	}
}

func TestEventEmitter_ListenerCount(t *testing.T) {
	emitter:=NewEventEmitter()
	emitter.On("ev", func() {})
	emitter.On("ev", func() {})
	if emitter.ListenerCount("ev")!=2 {
		t.Failed()
	}
	emitter.On("ev", func() {})
	if emitter.ListenerCount("ev")!=3 {
		t.Failed()
	}
}

func TestEventEmitter_EventNames(t *testing.T) {
	emitter:=NewEventEmitter()
	emitter.On("ev", func() {})
	evs:= emitter.EventNames()
	if len(evs)!=1 {
		t.Failed()
	}
	if evs[0]!="ev" {
		t.Failed()
	}
	emitter.On("ev1", func() {})
	evs= emitter.EventNames()
	if len(evs)!=2 {
		t.Failed()
	}
	if evs[0]!="ev"||evs[1]!="ev1" {
		t.Failed()
	}
}


func TestEventEmitter_Listeners(t *testing.T) {
	emitter:=NewEventEmitter()
	emitter.On("ev", func() {})
	emitter.On("ev", func() {})
	if len(emitter.Listeners("ev"))!=2 {
		t.Failed()
	}
	emitter.On("ev", func() {})
	if len(emitter.Listeners("ev"))!=3 {
		t.Failed()
	}
}


func TestEventEmitter_AddListener(t *testing.T) {
	emitter:=NewEventEmitter()
	emitter.On("ev", func() {})
	if len(emitter.Listeners("ev"))!=1 {
		t.Failed()
	}
	emitter.On("ev", func() {})
	if len(emitter.Listeners("ev"))!=2 {
		t.Failed()
	}
}

func TestEventEmitter_Once(t *testing.T) {
	emitter:=NewEventEmitter()
	emitter.Once("ev", func() {})
	if len(emitter.Listeners("ev"))!=1 {
		t.Failed()
	}
	emitter.Once("ev", func() {})
	if len(emitter.Listeners("ev"))!=2 {
		t.Failed()
	}
}

func TestEventEmitter_Off(t *testing.T) {
	emitter:=NewEventEmitter()
	listener:= func() {}
	emitter.On("ev", listener)
	if len(emitter.Listeners("ev"))!=1 {
		t.Failed()
	}
	emitter.Off("ev",listener)
	if len(emitter.Listeners("ev"))!=0 {
		t.Failed()
	}

	listener2:= func() {}
	emitter.Once("ev2", listener2)
	if len(emitter.Listeners("ev"))!=1 {
		t.Failed()
	}
	emitter.Off("ev2",listener2)
	if len(emitter.Listeners("ev2"))!=0 {
		t.Failed()
	}
}

func TestEventEmitter_RemoveListener(t *testing.T) {
	emitter:=NewEventEmitter()
	listener:= func() {}
	emitter.On("ev", listener)
	if len(emitter.Listeners("ev"))!=1 {
		t.Failed()
	}
	emitter.Off("ev",listener)
	if len(emitter.Listeners("ev"))!=0 {
		t.Failed()
	}

	listener2:= func() {}
	emitter.Once("ev2", listener2)
	if len(emitter.Listeners("ev"))!=1 {
		t.Failed()
	}
	emitter.Off("ev2",listener2)
	if len(emitter.Listeners("ev2"))!=0 {
		t.Failed()
	}
}

func TestEventEmitter_RemoveAllListeners(t *testing.T) {
	emitter:=NewEventEmitter()
	emitter.On("ev", func() {})
	emitter.On("ev1", func() {})
	emitter.RemoveAllListeners("ev")
	if len(emitter.Listeners("ev"))!=0 {
		t.Failed()
	}
	emitter.Once("ev1", func() {})
	emitter.RemoveAllListeners("ev1")
	if len(emitter.Listeners("ev1"))!=0 {
		t.Failed()
	}

	emitter.On("ev", func() {})
	emitter.On("ev1", func() {})

	emitter.RemoveAllListeners()
	if len(emitter.Listeners("ev"))!=0 {
		t.Failed()
	}
	if len(emitter.Listeners("ev1"))!=0 {
		t.Failed()
	}
}

func TestEventEmitter_Emit(t *testing.T) {
	emitter:=NewEventEmitter()
	wg:=sync.WaitGroup{}
	wg.Add(1)
	emitter.On("ev", func() {
		wg.Done()
	})
	wg.Add(1)
	emitter.Once("ev", func() {
		wg.Done()
	})
	emitter.Emit("ev")

	finish:=make(chan struct{},1)
	go func() {
		wg.Wait()
		finish<- struct{}{}
	}()

	select {
	case <-time.After(time.Second):
		t.Failed()
	case <-finish:
	}
}




