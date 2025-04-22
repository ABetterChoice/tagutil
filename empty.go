// Package tagutil ...
package tagutil

// emptyExecutor Determine whether the passed tag is an empty executor
type emptyExecutor struct{}

var (
	emptyExecutorImpl = &emptyExecutor{}
)

// IsEmpty unitTagValue Determine whether the passed tag value is empty
func (e *emptyExecutor) IsEmpty(unitTagValue []string, configValue string) bool {
	// 是否有，同时len不为0
	return unitTagValue == nil || len(unitTagValue) == 0
}

// IsNotEmpty unitTagValue Determine whether the passed tag value is not empty
func (e *emptyExecutor) IsNotEmpty(unitTagValue []string, configValue string) bool {
	return len(unitTagValue) > 0
}

// isEmptyTag unitTagValue Determine whether the passed tag key or value is empty or value is not ""
func isEmptyTag(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return true
	}
	for _, val := range unitTagValue {
		if val != "" {
			return false
		}
	}
	return true
}
