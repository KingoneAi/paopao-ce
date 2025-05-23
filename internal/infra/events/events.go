// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package events

import (
	"sync"

	"github.com/alimy/tryst/cfg"
	"github.com/alimy/tryst/pool"
	"github.com/robfig/cron/v3"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/infra/metrics/statistics"
	"github.com/sirupsen/logrus"
)

var (
	_defaultEventManager EventManager
	_defaultJobManager   JobManager = emptyJobManager{}
	_onceInitial         sync.Once
)

func StartEventManager() {
	_defaultEventManager.Start()
}

func StopEventManager() {
	_defaultEventManager.Stop()
}

// OnEvent push event to gorotine pool then handled automatic.
func OnEvent(event Event) {
	_defaultEventManager.OnEvent(event)
}

func StartJobManager() {
	_defaultJobManager.Start()
}

func StopJobManager() {
	_defaultJobManager.Stop()
}

// NewJob create new Job instance
func NewJob(s cron.Schedule, fn JobFn) Job {
	return &simpleJob{
		Schedule: s,
		Job:      fn,
	}
}

// RemoveJob an entry from being run in the future.
func RemoveJob(id EntryID) {
	_defaultJobManager.Remove(id)
}

// Schedule adds a Job to the Cron to be run on the given schedule.
// The job is wrapped with the configured Chain.
func Schedule(job Job) EntryID {
	return _defaultJobManager.Schedule(job)
}

// OnTask adds a Job to the Cron to be run on the given schedule.
// The job is wrapped with the configured Chain.
func OnTask(s cron.Schedule, fn JobFn) EntryID {
	job := &simpleJob{
		Schedule: s,
		Job:      fn,
	}
	return _defaultJobManager.Schedule(job)
}

func Initial() {
	_onceInitial.Do(func() {
		initEventManager()
		cfg.Not("DisableJobManager", func() {
			initJobManager()
			logrus.Debugln("initial JobManager")
		})
	})
}

func initJobManager() {
	_defaultJobManager = NewJobManager()
	StartJobManager()
}

func initEventManager() {
	var opts []pool.Option
	s := conf.EventManagerSetting
	if s.MinWorker > 5 {
		opts = append(opts, pool.WithMinWorker(s.MinWorker))
	} else {
		opts = append(opts, pool.WithMinWorker(5))
	}
	if s.MaxEventBuf > 10 {
		opts = append(opts, pool.WithMaxRequestBuf(s.MaxEventBuf))
	} else {
		opts = append(opts, pool.WithMaxRequestBuf(10))
	}
	if s.MaxTempEventBuf > 10 {
		opts = append(opts, pool.WithMaxRequestTempBuf(s.MaxTempEventBuf))
	} else {
		opts = append(opts, pool.WithMaxRequestTempBuf(10))
	}
	if cfg.If("Metrics") {
		opts = append(opts, pool.WithWorkerHook(NewEventWorkerHook("default", statistics.NewMetricCache())))
	}
	opts = append(opts, pool.WithMaxIdelTime(s.MaxIdleTime), pool.WithMaxTempWorker(s.MaxTempWorker))
	_defaultEventManager = NewEventManager(func(req Event, err error) {
		if err != nil {
			logrus.Errorf("handle event[%s] occurs error: %s", req.Name(), err)
		}
	}, opts...)
}
