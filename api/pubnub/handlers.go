package pubnub

import (
	"fmt"
	"sync"
	"time"
)

// handleResult handle result from publish action
func handleResult(successChan, errorChan chan []byte, timeoutVal uint16, action, channel string, wg *sync.WaitGroup) {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(timeoutVal) * time.Second)
		timeout <- true
	}()
	for {
		select {
		case success, ok := <-successChan:
			if !ok {
				break
			}
			if string(success) != "[]" {
				// fmt.Println(fmt.Sprintf("%s Response: %s ", action, success))
				close(successChan)
				close(errorChan)
				wg.Done()
				return
			}
		case failure, ok := <-errorChan:
			if !ok {
				break
			}
			if string(failure) != "[]" {
				fmt.Println(fmt.Sprintf("%s Error Callback: %s chan: %v", action, failure, channel))
				wg.Done()
				return
			}
		case <-timeout:
			fmt.Println(fmt.Sprintf("%s Handler timeout after %d secs chan: %v", action, timeoutVal, channel))
			wg.Done()
			return
		}
	}
}