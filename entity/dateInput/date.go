package dateInput

import "time"

type Date struct {
	start time.Time `json:"start"`
	end   time.Time `json:"end"`
}
