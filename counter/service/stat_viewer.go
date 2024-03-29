package service

import (
	"fmt"
	"github.com/championlong/design-patterns-go/counter/vo"
)

// StatViewer 负责显示发送
type StatViewer interface {
	output(requestStats map[string]vo.RequestStat, startTimeInMillis float64, endTimeInMills float64)
}

type ConsoleViewer struct {
}

func (self *ConsoleViewer) output(requestStats map[string]vo.RequestStat, startTimeInMillis float64, endTimeInMills float64) {
	fmt.Println("ConsoleViewer", requestStats, startTimeInMillis, endTimeInMills)
}

type EmailViewer struct {
}

func (self *EmailViewer) output(requestStats map[string]vo.RequestStat, startTimeInMillis float64, endTimeInMills float64) {
	fmt.Println("EmailViewer", requestStats, startTimeInMillis, endTimeInMills)
}
