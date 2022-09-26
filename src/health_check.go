package main

import "sync"

type healthCheck struct {
	mu     sync.Mutex
	Status int `json:"status"`
}

func newHealthCheck() *healthCheck {
	hc := new(healthCheck)
	hc.UpdateStatus(200)

	return hc
}

func (c *healthCheck) GetStatus() int {
	return c.Status
}

func (c *healthCheck) UpdateStatus(status int) {
	c.mu.Lock()
	c.Status = status
	c.mu.Unlock()
}
