package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	// Map com o nome do evento como string e
	// um array de EventHandlerInterface
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (e *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	// Verifica se o evento já foi registrado
	if _, ok := e.handlers[eventName]; ok {
		for _, h := range e.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	// Adiciona o handler ao evento
	e.handlers[eventName] = append(e.handlers[eventName], handler)
	return nil
}

func (e *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := e.handlers[eventName]; ok {
		for _, h := range e.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (e *EventDispatcher) Dispatch(event EventInterface) {
	if handlers, ok := e.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
}

func (e *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := e.handlers[eventName]; ok {
		for i, h := range e.handlers[eventName] {
			if h == handler {
				e.handlers[eventName] = append(e.handlers[eventName][:i], e.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func (e *EventDispatcher) Clear() {
	e.handlers = make(map[string][]EventHandlerInterface)
}
