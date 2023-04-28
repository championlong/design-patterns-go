package service

import (
	"github.com/championlong/design-patterns-go/counter/dao"
)

// ScheduledReporter 作为父类解决代码重复问题
type ScheduledReporter struct {
	metricsStorage dao.MetricsStorage
	aggregator     Aggregator
	viewer         StatViewer
}

func NewScheduledReporter(metricsStorage dao.MetricsStorage, aggregator Aggregator, viewer StatViewer) *ScheduledReporter {
	return &ScheduledReporter{
		metricsStorage: metricsStorage,
		aggregator:     aggregator,
		viewer:         viewer,
	}
}

func (self *ScheduledReporter) doStatAndReport(startTimeInMillis, endTimeInMillis float64) {
	durationInMillis := endTimeInMillis - startTimeInMillis
	requestInfos := self.metricsStorage.GetAllRequestInfosByDuration(startTimeInMillis, endTimeInMillis)
	requestStats := self.aggregator.Aggregate(requestInfos, durationInMillis)
	self.viewer.output(requestStats, startTimeInMillis, endTimeInMillis)
}
