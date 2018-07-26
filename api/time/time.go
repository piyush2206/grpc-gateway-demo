package time

import (
	"time"

	"golang.org/x/net/context"
)

// EntityTime is used to implement time.Time
type EntityTime struct{}

// Timezone implements time.Time
func (*EntityTime) Timezone(ctx context.Context, req *ReqTimezone) (*ResTimezone, error) {

	layout := "2006-01-02T15:04:05.000Z"

	t, err := time.Parse(layout, req.Time)
	if err != nil {
		return nil, err
	}

	//init the loc
	loc, _ := time.LoadLocation(req.Zone)

	//set timezone
	now := t.In(loc)

	return &ResTimezone{Time: now.String()}, nil
}
