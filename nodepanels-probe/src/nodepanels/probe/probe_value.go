package probe

var memTotal uint

var diskInitReadBytes map[string]uint64  //记录初始磁盘读速率
var diskInitWriteBytes map[string]uint64 //记录初始磁盘写速率

var netInitRxBytes uint64 //下行流量初始值
var netInitTxBytes uint64 //上行流量初始值

var initProcessUsageMap map[string]ProcessUsage
