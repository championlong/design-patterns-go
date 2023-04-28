package main

import (
	collector "github.com/championlong/design-patterns-go/counter/collector"
	"github.com/championlong/design-patterns-go/counter/dao"
	"github.com/championlong/design-patterns-go/counter/service"
	"github.com/championlong/design-patterns-go/counter/vo"
	"time"
)

// MVC三层开发: 分层能起到代码复用的作用、分层能起到隔离变化的作用、分层能起到隔离关注点的作用、分层能提高代码的可测试性、分层能应对系统的复杂性
// VO、BO、Entity的设计思路并不违反DRY原则，为了分层清晰、减少耦合，多维护几个类的成本也并不是不能接受的。
// 功能性需求: 合理地将功能划分到不同模块、设计模块与模块之间的交互关系、设计模块的接口、数据库、业务模型
// 非功能性的需求: 考虑易用性、性能、扩展性、容错性、通用性
// 面向对象设计和实现要做的事情，就是把合适的代码放到合适的类中。
/*
能够获取接口调用的各种统计信息，比如，响应时间的最大值（max）、最小值（min）、平均值（avg）、百分位值（percentile）、接口调用次数（count）、频率（tps） 等，
并且支持将统计结果以各种显示格式（比如：JSON格式、网页格式、自定义显示格式等）输出到各种终端（Console命令行、HTTP网页、Email、日志文件、自定义输出终端等），以方便查看。
*/
func main() {
	dao.RequestInfos = make(map[string][]vo.RequestInfo)

	consoleReporter := service.NewConsoleRedisReporter()
	// 每10秒计算60s的请求
	go consoleReporter.StartRepeatedReport(10, 60)

	metrics := collector.NewMetricsRedisCollector()
	metrics.RecordRequest(vo.NewRequestInfo("register", 123, 10234))
	metrics.RecordRequest(vo.NewRequestInfo("register", 223, 11234))
	metrics.RecordRequest(vo.NewRequestInfo("login", 23, 12434))

	time.Sleep(10 * time.Minute)
}
