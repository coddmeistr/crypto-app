package ws

import (
	"fmt"
	"log"
	"time"
)

type ISubscribtion interface {
	Start()
	End()
	GetMessage() ([]byte, bool)
}

type Subscribtion struct {
	DataPump            chan []byte
	closePump           chan struct{}
	subscribeCallback   func()
	actionCallback      func() ([]byte, error)
	unsubscribeCallback func()
}

func NewSubscription(subscribeCallback func(), actionCallback func() ([]byte, error), unsubscribeCallback func()) ISubscribtion {
	return &Subscribtion{
		DataPump:            make(chan []byte, 256),
		closePump:           make(chan struct{}),
		subscribeCallback:   subscribeCallback,
		actionCallback:      actionCallback,
		unsubscribeCallback: unsubscribeCallback,
	}
}

func (s *Subscribtion) subscribe() {
	s.subscribeCallback()
}

func (s *Subscribtion) action() {
	defer func() {
		close(s.DataPump)
		fmt.Println("Exiting subscripting action main function")
	}()

	bytesChan := make(chan []byte, 10)
	doneMsgChan := make(chan struct{})
	go func() {
		defer func() {
			close(bytesChan)
			fmt.Println("Exiting subscripting action helping gouruting")
		}()
		for {
			data, err := s.actionCallback()
			if err != nil {
				log.Printf("error: %v", err)
				break
			}

			select {
			case <-doneMsgChan:
				return
			case bytesChan <- data:
			}
		}
	}()

	for {
		select {
		case <-s.closePump:
			go func() {
				select {
				case doneMsgChan <- struct{}{}:
					return
				case <-time.After(1 * time.Second):
					fmt.Println("Cannot push to doneMsgChan. Timeted out")
					return
				}
			}()
			return
		case v, ok := <-bytesChan:
			if !ok {
				return
			}
			s.DataPump <- v

		default:
		}
	}
}

func (s *Subscribtion) GetMessage() ([]byte, bool) {
	data, ok := <-s.DataPump
	return data, ok
}

func (s *Subscribtion) End() {
	go s.unsubscribeCallback() // Creating gourutine because it may contain async logic
	s.closePump <- struct{}{}
}

func (s *Subscribtion) Start() {
	go func() {
		log.Println("Starting subscription loop...")
		s.subscribe()
		s.action()
		log.Println("Ending subscription loop...")
	}()
}
