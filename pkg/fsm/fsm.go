package fsm

import (
	"errors"
	"fmt"
	"github.com/solost23/protopb/gen/go/protos/order_machine"
	"sync"
)

type Handler func(opt *Opt) (order_machine.OrderStatus, error)

// FSM 有限状态机
type FSM struct {
	mu       sync.Mutex                                                         // 排他锁
	status   order_machine.OrderStatus                                          // 当前状态
	handlers map[order_machine.OrderStatus]map[order_machine.OrderEvent]Handler // 当前状态可触发的有限个事件
}

func NewFSM(initStatus order_machine.OrderStatus) (fsm *FSM, err error) {
	fsm = new(FSM)
	fsm.status = initStatus
	fsm.handlers = make(map[order_machine.OrderStatus]map[order_machine.OrderEvent]Handler)
	fsm, err = fsm.addHandlers()
	if err != nil {
		return nil, err
	}
	return fsm, nil
}

// 获取当前状态
func (f *FSM) getState() order_machine.OrderStatus {
	return f.status
}

// 设置当前状态
func (f *FSM) setState(newState order_machine.OrderStatus) {
	f.status = newState
}

// addHandlers 添加事件和处理方法
func (f *FSM) addHandlers() (*FSM, error) {
	if StatusEvent == nil || len(StatusEvent) <= 0 {
		return nil, errors.New("[警告] 未定义 statusEvent")
	}

	for state, events := range StatusEvent {
		if len(events) <= 0 {
			return nil, errors.New(fmt.Sprintf("[警告] 状态(%s)未定义事件", StatusText(state)))
		}

		for _, event := range events {
			handler := eventHandler[event]
			if handler == nil {
				return nil, errors.New(fmt.Sprintf("[警告] 事件(%s)未定义处理方法", event))
			}

			if _, ok := f.handlers[state]; !ok {
				f.handlers[state] = make(map[order_machine.OrderEvent]Handler)
			}

			if _, ok := f.handlers[state][event]; ok {
				return nil, errors.New(fmt.Sprintf("[警告] 状态(%s)事件(%s)已定义过", StatusText(state), event))
			}

			f.handlers[state][event] = handler
		}
	}

	return f, nil
}

// Call 事件处理
func (f *FSM) Call(event order_machine.OrderEvent, opts ...Option) (order_machine.OrderStatus, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	events := f.handlers[f.getState()]
	if events == nil {
		return 0, errors.New(fmt.Sprintf("[警告] 状态(%s)未定义任何事件", StatusText(f.getState())))
	}

	// 将所有传入的参数生成了一个opt对象
	opt := new(Opt)
	for _, f := range opts {
		f(opt)
	}

	fn, ok := events[event]
	if !ok {
		return 0, errors.New(fmt.Sprintf("[警告] 状态(%s)不允许操作(%s)", StatusText(f.getState()), event))
	}

	status, err := fn(opt)
	if err != nil {
		return 0, err
	}

	oldState := f.getState()
	f.setState(status)
	newState := f.getState()

	fmt.Println(fmt.Sprintf("操作[%s]，状态从 [%s] 变成 [%s]", event, StatusText(oldState), StatusText(newState)))

	return f.getState(), nil
}
