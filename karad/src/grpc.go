package src

import (
	"context"
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"kara"
)

type JobService struct{}

func (js *JobService) SingleTimesJob(ctx context.Context, req *SingleTimesJobRequest) (*SingleTimesJobResponse, error) {
	if req.Key == "" {
		return nil, fmt.Errorf("SingleTimesJobReqeust's 'key' field is required")
	}

	if req.SpotId == "" {
		req.SpotId = req.Key
	}

	spot, _ := kara.KaraPool.LoadOrStore(req.SpotId, kara.NewSpot())
	ok, e := spot.(*kara.KaraSpot).SetWhenNotExist(req.Key)
	if e != nil {
		return nil, e
	}
	return &SingleTimesJobResponse{
		Status:  ok,
		Message: "success",
	}, nil
}

func (js *JobService) MultipleTimesJob(ctx context.Context, req *MultipleTimesJobRequest) (*MultipleTimesJobResponse, error) {
	if req.Limit == 0 {
		return nil, fmt.Errorf("MultipleTimesJobReqeust's 'limit' field is required")
	}

	if req.Key == "" {
		return nil, fmt.Errorf("MultipleTimesJobReqeust's 'key' field is required")
	}

	if req.SpotId == "" {
		req.SpotId = req.Key
	}

	spot, _ := kara.KaraPool.LoadOrStore(req.SpotId, kara.NewTimesSpot(int(req.Limit)))
	ok, e := spot.(*kara.KaraSpot).AddWhenNotReachedLimit(req.Key)
	if e != nil {
		return nil, errorx.Wrap(e)
	}
	return &MultipleTimesJobResponse{
		Status:  ok,
		Message: "success",
	}, nil
}
