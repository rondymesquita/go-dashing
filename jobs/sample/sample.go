package sample

import (
	"math/rand"
	"time"

	"github.com/znly/go-dashing/dashingtypes"
)

// Job is the sample Job structure
type Job struct{}

// Work implements job interface
func (j *Job) Work(send chan *dashingtypes.Event) {
	ticker := time.NewTicker(1 * time.Second)
	var lastValuation, lastKarma, currentValuation, currentKarma int
	for {
		select {
		case <-ticker.C:
			lastValuation, currentValuation = currentValuation, rand.Intn(100)
			lastKarma, currentKarma = currentKarma, rand.Intn(200000)
			send <- &dashingtypes.Event{"valuation", map[string]interface{}{
				"current": currentValuation,
				"last":    lastValuation,
			}, ""}
			send <- &dashingtypes.Event{"karma", map[string]interface{}{
				"current": currentKarma,
				"last":    lastKarma,
			}, ""}
			send <- &dashingtypes.Event{"synergy", map[string]interface{}{
				"value": rand.Intn(100),
			}, ""}
		}
	}
}

// GetJob returns a new job
func GetJob() *Job {
	return &Job{}
}
