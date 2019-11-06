package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kara"
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
	spot, _ := kara.KaraPool.LoadOrStore(param.SpotID, kara.NewSpot())
	ok, e := spot.(*kara.KaraSpot).SetWhenNotExist(param.Key)
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

	spot, _ := kara.KaraPool.LoadOrStore(param.SpotID, kara.NewTimesSpot(param.Limit))
	ok, e := spot.(*kara.KaraSpot).AddWhenNotReachedLimit(param.Key)
	if e != nil {
		c.JSON(400, Result{ok, e.Error()})
		return
	}
	c.JSON(200, Result{ok, "success"})
}
