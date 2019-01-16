package rpi

import (
	common "github.com/gbbirkisson/rpi/pkg/common"
)

var Revision string = "development"
var Version string = "development"

func GetCommonServer() *common.CommonServer {
	return &common.CommonServer{Version: Version, Revision: Revision}
}
