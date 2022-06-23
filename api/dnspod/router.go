package dnspod

import (
	"tdp-cloud/core/midware"

	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/dnspod")

	rg.Use(midware.Auth())
	rg.Use(midware.Secret())

	{
		rg.POST("/describeDomainList", describeDomainList)
		rg.POST("/describeRecordList", describeRecordList)

		rg.POST("/describeRecordLineList", describeRecordLineList)
		rg.POST("/describeRecordType", describeRecordType)

		rg.POST("/modifyRecord", modifyRecord)
	}

}
