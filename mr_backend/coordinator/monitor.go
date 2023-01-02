package coordinator

import "time"

func (c *Coordinator) monitor(quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			c.mu.Lock()
			for _, mTask := range c.mTasks {
				if mTask.state == stateInProgress && time.Now().Sub(mTask.started) > c.taskTimeout {
					mTask.state = stateIdle
					mTask.started = time.Time{}
				}
			}
			for _, rTask := range c.rTasks {
				if rTask.state == stateInProgress && time.Now().Sub(rTask.started) > c.taskTimeout {
					rTask.state = stateIdle
					rTask.started = time.Time{}
				}
			}
			c.mu.Unlock()
		}
	}
}
