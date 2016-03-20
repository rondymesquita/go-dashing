package convergence

import (
	"math/rand"
	"time"

	"github.com/znly/go-dashing/dashingtypes"
)

// Job is the convergence Job structure
type Job struct {
	points []map[string]int
}

// Work implements job interface
func (j *Job) Work(send chan *dashingtypes.Event) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			j.points = j.points[1:]
			j.points = append(j.points, map[string]int{
				"x": j.points[len(j.points)-1]["x"] + 1,
				"y": rand.Intn(50),
			})
			send <- &dashingtypes.Event{"convergence", map[string]interface{}{
				"points": j.points,
			}, ""}
		}
	}
}

// GetJob returns a new job
func GetJob() *Job {
	c := &Job{}
	for i := 0; i < 10; i++ {
		c.points = append(c.points, map[string]int{
			"x": i,
			"y": rand.Intn(50),
		})
	}
	return c
}
