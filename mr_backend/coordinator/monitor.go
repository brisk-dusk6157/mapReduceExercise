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
				if mTask.state == STATE_IN_PROGRESS && time.Now().Sub(mTask.started) > c.taskTimeout {
					mTask.state = STATE_IDLE
					mTask.started = time.Time{}
				}
			}
			for _, rTask := range c.rTasks {
				if rTask.state == STATE_IN_PROGRESS && time.Now().Sub(rTask.started) > c.taskTimeout {
					rTask.state = STATE_IDLE
					rTask.started = time.Time{}
				}
			}
			c.mu.Unlock()
		}
	}
}
