// Package tagutil ...
package tagutil

import "strings"

// setExecutor 集合执行器
type setExecutor struct{}

var (
	setExecutorImpl = &setExecutor{}
)

// IsSubSet unitTagValue 是否为 configValue 的子集，特别的，空的 unitTagValue 结果恒为 true
func (e *setExecutor) IsSubSet(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return true // 特殊的，不同于其他执行器，空也算子集
	}
	configValueItemList := strings.Split(configValue, splitSeg) // 数组统一用 ; 分割
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

// IsSuperSet unitTagValue 是否为 configValue 的超集，特别的，空的 configValue 结果恒为 true
func (e *setExecutor) IsSuperSet(unitTagValue []string, configValue string) bool {
	if len(configValue) == 0 {
		return true // 空的 configValue 结果恒为 true
	}
	configValueItemList := strings.Split(configValue, splitSeg) // 数组统一用 ; 分割
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

func (e *setExecutor) EG(unitTagValue []string, configValue string) bool {
	return e.IsSubSet(unitTagValue, configValue) && e.IsSuperSet(unitTagValue, configValue)
	//if len(unitTagValue) == 0 && len(configValue) == 0 {
	//	return true
	//}
	//if len(unitTagValue) == 0 {
	//	return false
	//}
	//if len(configValue) == 0 {
	//	return false
	//}
	//configValueItemList := strings.Split(configValue, splitSeg)
	//var countMap = make(map[string]uint8, len(configValueItemList))
	//for i := range unitTagValue {
	//	if countMap[unitTagValue[i]] == 0 { // 去重
	//		countMap[unitTagValue[i]] = 1
	//	}
	//}
	//for i := range configValueItemList {
	//	if countMap[configValueItemList[i]] == 0 { // 配置中有的值，用户标签值中没有传入
	//		return false
	//	} else { // == 1
	//		countMap[configValueItemList[i]] = 2 // 同个key，在 unitTagValue\configValue 都存在
	//	}
	//}
	//for _, v := range countMap {
	//	if v != 2 { // 2 代表key，在两个set中都出现，如果不等于2，则key在 unitTagValue 中出现，但配置中没出现
	//		return false
	//	}
	//}
	//return true
}
