package src

import (
	"fmt"
	"github.com/fwhezfwhez/kara"
	"github.com/gin-gonic/gin"
)

func SingleJob(c *gin.Context) {
	fmt.Println(c.ContentType())
	type Param struct {
		SpotID string `json:"spot_id"`
		Key    string `json:"key" binding:"required"`
	}

	type Result struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
	var param Param
	if e := c.Bind(&param); e != nil {
		c.JSON(400, gin.H{"message": e.Error()})
		return
	}

	if param.SpotID == "" {
		param.SpotID = param.Key
	}
	spt, _ := kara.KaraPool.LoadOrStore(param.SpotID, kara.NewSpot())
	spot :=spt.(*kara.KaraSpot)
	if spot.Type != 1 {
		c.JSON(400,gin.H{"message": fmt.Sprintf("spot_id '%s' is not exist-type, got type %d", param.SpotID, spot.Type)})
		return
	}

	ok, e := spot.SetWhenNotExist(param.Key)
	if e != nil {
		c.JSON(400, Result{ok, e.Error()})
		return
	}
	c.JSON(200, Result{ok, "success"})
}

func MultipleJob(c *gin.Context) {
	type Param struct {
		SpotID string `json:"spot_id"`
		Key    string `json:"key" binding:"required"`
		Limit  int    `json:"limit" binding:"required"`
	}

	type Result struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
	var param Param
	if e := c.Bind(&param); e != nil {
		c.JSON(400, gin.H{"message": e.Error()})
		return
	}

	if param.SpotID == "" {
		param.SpotID = param.Key
	}

	spt, _ := kara.KaraPool.LoadOrStore(param.SpotID, kara.NewTimesSpot(param.Limit))
	spot := spt.(*kara.KaraSpot)
	if spot.Type != 2 {
		c.JSON(400,gin.H{"message": fmt.Sprintf("spot_id '%s' is not times-type, got type %d", param.SpotID, spot.Type)})
		return
	}
	ok, e := spot.AddWhenNotReachedLimit(param.Key)
	if e != nil {
		c.JSON(400, Result{ok, e.Error()})
		return
	}
	c.JSON(200, Result{ok, "success"})
}
