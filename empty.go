// Package tagutil ...
package tagutil

// emptyExecutor 判断传入的标签是否为空执行器
type emptyExecutor struct{}

var (
	emptyExecutorImpl = &emptyExecutor{}
)

// IsEmpty unitTagValue 传入的标签值是否为空
func (e *emptyExecutor) IsEmpty(unitTagValue []string, configValue string) bool {
	return unitTagValue == nil || len(unitTagValue) == 0
}

// IsNotEmpty unitTagValue 传入的标签值是否不为空
func (e *emptyExecutor) IsNotEmpty(unitTagValue []string, configValue string) bool {
	return len(unitTagValue) > 0
}
