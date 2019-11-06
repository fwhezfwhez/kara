package src

import (
	"context"
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"github.com/fwhezfwhez/kara"
)

type JobService struct{}

func (js *JobService) SingleTimesJob(ctx context.Context, req *SingleTimesJobRequest) (*SingleTimesJobResponse, error) {
	if req.Key == "" {
		return nil, fmt.Errorf("SingleTimesJobReqeust's 'key' field is required")
	}

	if req.SpotId == "" {
		req.SpotId = req.Key
	}

	spt, _ := kara.KaraPool.LoadOrStore(req.SpotId, kara.NewSpot())
	spot := spt.(*kara.KaraSpot)
	if spot.Type != 1 {
		return &SingleTimesJobResponse{
			Status:  false,
			Message: fmt.Sprintf("spot_id '%s' is not times-type, got type %d", req.SpotId, spot.Type),
		}, nil
	}
	ok, e := spot.SetWhenNotExist(req.Key)
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

	spt, _ := kara.KaraPool.LoadOrStore(req.SpotId, kara.NewTimesSpot(int(req.Limit)))
	spot := spt.(*kara.KaraSpot)
	if spot.Type != 2 {
		return &MultipleTimesJobResponse{
			Status: false,
			Message: fmt.Sprintf("spot_id '%s' is not times-type, got type %d", req.SpotId, spot.Type),
		}, nil
	}
	ok, e := spot.AddWhenNotReachedLimit(req.Key)
	if e != nil {
		return nil, errorx.Wrap(e)
	}
	return &MultipleTimesJobResponse{
		Status:  ok,
		Message: "success",
	}, nil
}
