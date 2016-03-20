package buzzwords

import (
	"math/rand"
	"time"

	"github.com/znly/go-dashing/dashingtypes"
)

// Job is the buzzword Job structure
type Job struct {
	words []map[string]interface{}
}

// Work implements job interface
func (j *Job) Work(send chan *dashingtypes.Event) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:

			for i := 0; i < len(j.words); i++ {
				if 1 < rand.Intn(3) {
					value := j.words[i]["value"].(int)
					j.words[i]["value"] = (value + 1) % 30
				}
			}
			send <- &dashingtypes.Event{"buzzwords", map[string]interface{}{
				"items": j.words,
			}, ""}
		}
	}
}

// GetJob returns a new job
func GetJob() *Job {
	return &Job{[]map[string]interface{}{
		{"label": "Paradigm shift", "value": 0},
		{"label": "Leverage", "value": 0},
		{"label": "Pivoting", "value": 0},
		{"label": "Turn-key", "value": 0},
		{"label": "Streamlininess", "value": 0},
		{"label": "Exit strategy", "value": 0},
		{"label": "Synergy", "value": 0},
		{"label": "Enterprise", "value": 0},
		{"label": "Web 2.0", "value": 0},
	}}
}
