// Package tagutil ...
package tagutil

import "strconv"

// booleanExecutor 布尔表达式执行器
type booleanExecutor struct{}

var (
	booleanExecutorImpl = &booleanExecutor{}
)

// EG unitTagValue 的值是否跟 configValue 一致，需要 unitTagValue 中所有元素都满足才为 true
func (e *booleanExecutor) EG(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return false
	}
	configValueBoolean, err := strconv.ParseBool(configValue)
	if err != nil {
		return false
	}
	for i := range unitTagValue {
		itemBoolean, err := strconv.ParseBool(unitTagValue[i])
		if err != nil {
			return false
		}
		if itemBoolean != configValueBoolean { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}
