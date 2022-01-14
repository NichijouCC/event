package events

import "testing"

func TestNewSyncEventEmitter(t *testing.T) {
	emitter:=NewSyncEventEmitter()
	if emitter==nil {
		t.Failed()
	}
}

func TestSyncEventEmitter_ListenerCount(t *testing.T) {
	emitter:=NewSyncEventEmitter()
	emitter.On("ev", func() {})
	if emitter.ListenerCount("ev")!=1 {
		t.Failed()
	}
	emitter.On("ev", func() {})
	if emitter.ListenerCount("ev")!=2 {
		t.Failed()
	}
}

func TestSyncEventEmitter_On(t *testing.T) {
	emitter:=NewSyncEventEmitter()
	emitter.On("ev", func() {})
	if emitter.ListenerCount("ev")!=1 {
		t.Failed()
	}
}

func TestSyncEventEmitter_RemoveAllListeners(t *testing.T) {
	emitter:=NewSyncEventEmitter()
	emitter.On("ev", func() {})
	emitter.RemoveAllListeners("ev")
	if emitter.ListenerCount("ev")!=0 {
		t.Failed()
	}
	emitter.On("ev", func() {})
	emitter.RemoveAllListeners()
	if emitter.ListenerCount("ev")!=0 {
		t.Failed()
	}
}
func TestSyncEventEmitter_AddListener(t *testing.T) {
	emitter:=NewSyncEventEmitter()
	emitter.AddListener("ev", func() {})
	if emitter.ListenerCount("ev")!=1 {
		t.Failed()
	}
}

func TestSyncEventEmitter_EventNames(t *testing.T) {
	emitter:=NewSyncEventEmitter()
	emitter.AddListener("ev", func() {})
	emitter.AddListener("ev1", func() {})
	nameArr:=emitter.EventNames()
	if len(nameArr)!=2 {
		t.Failed()
	}
	if nameArr[0]!="ev"||nameArr[1]!="ev1" {
		t.Failed()
	}
}

func TestSyncEventEmitter_Once(t *testing.T) {
	emitter:=NewSyncEventEmitter()
	count:=0
	emitter.Once("ev", func() {
		count++
	})
	if emitter.ListenerCount("ev")!=1 {
		t.Failed()
	}
	emitter.Emit("ev")
	emitter.Emit("ev")
	if count!=1 {
		t.Failed()
	}
}

func TestSyncEventEmitter_Off(t *testing.T) {
	emitter:=NewSyncEventEmitter()
	listener:= func() {}
	emitter.On("ev",listener)
	emitter.Off("ev",listener)
	if emitter.ListenerCount("ev")!=0 {
		t.Failed()
	}
}

func TestSyncEventEmitter_Emit(t *testing.T) {
	emitter:=NewSyncEventEmitter()
	count:=0
	emitter.Once("ev", func() {
		count++
	})
	emitter.On("ev", func() {
		count++
	})
	emitter.Emit("ev")
	if count!=2 {
		t.Failed()
	}
	emitter.Emit("ev")
	if count!=3 {
		t.Failed()
	}
}