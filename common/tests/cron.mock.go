package tests

import (
	"github.com/stretchr/testify/mock"
	"gopkg.in/robfig/cron.v2"
)

type CronMock struct {
	*mock.Mock
	jobs map[string]cron.Job
}

func (c *CronMock) AddJob(spec string, cmd cron.Job) error {
	args := c.Called(spec, cmd)
	return args.Error(0)
}

func (c *CronMock) Remove(id cron.EntryID) {
	c.Called(id)
}

func (c *CronMock) Schedule(schedule cron.Schedule, cmd cron.Job) {
	c.Called(schedule, cmd)
}

func (c *CronMock) Entries() []*cron.Entry {
	return c.entrySnapshot()
}

func (c *CronMock) Start() {
	c.Called()
}

func (c *CronMock) run() {
	c.Called()
}

func (c *CronMock) Stop() {
	c.Called()
}

func (c *CronMock) entrySnapshot() []*cron.Entry {
	c.Called()
	entries := []*cron.Entry{}
	return entries
}
