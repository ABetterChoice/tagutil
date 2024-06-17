// Package tagutil ...
package tagutil

import "strings"

// setExecutor Collection Executor
type setExecutor struct{}

var (
	setExecutorImpl = &setExecutor{}
)

// IsSubset Whether unitTagValue is a subset of configValue. In particular, an empty unitTagValue always returns true.
func (e *setExecutor) IsSubset(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return true // 特殊的，不同于其他执行器，空也算子集
	}
	configValueItemList := strings.Split(configValue, splitSeg)
	for i := range unitTagValue {
		inFlag := false
		for j := range configValueItemList {
			if unitTagValue[i] == configValueItemList[j] {
				inFlag = true
				break
			}
		}
		if !inFlag {
			return false
		}
	}
	return true
}

// IsSuperset Whether unitTagValue is a superset of configValue. In particular, an empty configValue always returns true.
func (e *setExecutor) IsSuperset(unitTagValue []string, configValue string) bool {
	if len(configValue) == 0 {
		return true // 空的 configValue 结果恒为 true
	}
	configValueItemList := strings.Split(configValue, splitSeg)
	for i := range configValueItemList {
		inFlag := false
		for j := range unitTagValue {
			if configValueItemList[i] == unitTagValue[j] {
				inFlag = true
				break
			}
		}
		if !inFlag {
			return false
		}
	}
	return true
}

// EQ equal
func (e *setExecutor) EQ(unitTagValue []string, configValue string) bool {
	return e.IsSubset(unitTagValue, configValue) && e.IsSuperset(unitTagValue, configValue)
}
