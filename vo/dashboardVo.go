package vo

import (
	"dst-admin-go/constant"
	"dst-admin-go/utils/systemUtils"
)

type DashboardVO struct {
	IsInstallDst bool                  `json:"isInstallDst"`
	MasterStatus bool                  `json:"masterStatus"`
	CavesStatus  bool                  `json:"cavesStatus"`
	HostInfo     *systemUtils.HostInfo `json:"host"`
	CpuInfo      *systemUtils.CpuInfo  `json:"cpu"`
	MemInfo      *systemUtils.MemInfo  `json:"mem"`
	DiskInfo     *systemUtils.DiskInfo `json:"disk"`
	MemStates    uint64                `json:"memStates"`
	MasterLog    string                `json:"masterLog"`
	Version      string                `json:"version"`

	MasterPs *DstPsVo `json:"masterPs"`
	CavesPs  *DstPsVo `json:"cavesPs"`
}

func NewDashboardVO() *DashboardVO {
	return &DashboardVO{
		MasterLog: constant.HOME_PATH + "/" + constant.DST_MASTER_SERVER_LOG_PATH,
	}
}
