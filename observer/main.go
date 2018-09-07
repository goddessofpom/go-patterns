package main

import (
	"fmt"
	"time"
)

type (
	// 事件为一个即时的协程
	Event struct {
		Data int64
	}

	//观察者方法
	Observer interface {
		//收到通知
		OnNotify(Event)
	}

	//被观察的对象方法
	Notifier interface {
		Register(Observer)

		Unregister(Observer)

		Notify(Event)
	}
)

type (
	//观察者结构体
	eventObsever struct {
		id int
	}

	//被观察者结构体
	eventNotifier struct {
		observers map[Observer]struct{}
	}
)

func (o *eventObsever) OnNotify(e Event) {
	fmt.Printf("*** Observer %d receive: %d\n", o.id, e.Data)
}

func (n *eventNotifier) Register(o Observer) {
	n.observers[o] = struct{}{}
}

func (n *eventNotifier) Unregister(o Observer) {
	delete(n.observers, o)
}

func (n *eventNotifier) Notify(e Event) {
	for o := range n.observers {
		o.OnNotify(e)
	}
}

func main() {
	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	n.Register(&eventObsever{id: 1})
	n.Register(&eventObsever{id: 2})

	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C

	for {
		select {
		case <-stop:
			return

		case t := <-tick:
			n.Notify(Event{Data: t.UnixNano()})
		}
	}
}
