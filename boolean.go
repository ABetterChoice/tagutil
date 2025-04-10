// Package tagutil ...
package tagutil

import "strconv"

// booleanExecutor Boolean Expression Executor
type booleanExecutor struct{}

var (
	booleanExecutorImpl = &booleanExecutor{}
)

// Whether the value of EQ unitTagValue is consistent with configValue, all elements in unitTagValue must be satisfied to be true
func (e *booleanExecutor) EQ(unitTagValue []string, configValue string) bool {
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

// Empty Determine the TagKey for UnitTagValue does not exist
func (e *booleanExecutor) Empty(unitTagValue []string, configValue string) bool {
	return len(unitTagValue) == 0
}

// NotEmpty Determine the TagKey for UnitTagValue exists
func (e *booleanExecutor) NotEmpty(unitTagValue []string, configValue string) bool {
	return len(unitTagValue) != 0
}
