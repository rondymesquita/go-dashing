package dashing

import "github.com/znly/go-dashing/dashingtypes"

// Job does work and sends events to a channel.
type Job interface {
	Work(send chan *dashingtypes.Event)
}

type worker struct {
	broker   *broker
	registry []Job
}

func (w *worker) start() {
	for _, j := range w.registry {
		go j.Work(w.broker.events)
	}
}

func newWorker(b *broker) *worker {
	return &worker{
		broker:   b,
		registry: []Job{},
	}
}

func (w *worker) register(jobs ...Job) {
	var finalsize = 0
	var finalJobs []Job
	for _, job := range jobs {
		if job != nil {
			finalsize++
		}
	}
	if finalsize == len(jobs) {
		finalJobs = jobs
	} else {
		finalJobs = make([]Job, finalsize)
		i := 0
		for _, job := range jobs {
			if job != nil {
				finalJobs[i] = job
				i++
			}
		}
	}
	w.registry = append(w.registry, finalJobs...)
}
