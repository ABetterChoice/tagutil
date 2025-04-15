// Package tagutil ...
package tagutil

import (
	"regexp"
	"strings"
)

// stringExecutor String expression executor. Normally, the string type and the user tag value should be a string.
// Since the user tag is a string array, any string in the array that does not satisfy the expression will be judged as false.
// For example, the user tag value [ab,aa] must be greater than aa to satisfy. Although ab satisfies ab > aa, aa does not satisfy aa > aa, so the result is false.
// True only if all are met
type stringExecutor struct{}

var (
	stringExecutorImpl = stringExecutor{}
)

// EQ The tag value of the unit is equal to the tag value configured by the web system. All unitTagValue elements must meet this requirement to pass.
func (e *stringExecutor) EQ(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	for i := range unitTagValue {
		if unitTagValue[i] != configValue { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// LT The tag value of the unit is less than the tag value configured by the web system. All unitTagValue elements must meet this requirement to pass.
func (e *stringExecutor) LT(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	for i := range unitTagValue {
		if strings.Compare(unitTagValue[i], configValue) >= 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// LTE The tag value of the unit is less than or equal to the tag value configured by the web system. All unitTagValue elements must meet this requirement to pass.
func (e *stringExecutor) LTE(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	for i := range unitTagValue {
		if strings.Compare(unitTagValue[i], configValue) > 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// GT The tag value of the unit is greater than the tag value configured by the web system. All unitTagValue elements must satisfy this condition to pass.
func (e *stringExecutor) GT(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	for i := range unitTagValue {
		if strings.Compare(unitTagValue[i], configValue) <= 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// GTE The tag value of the unit is greater than or equal to the tag value configured by the web system. All unitTagValue elements must meet this requirement to pass.
func (e *stringExecutor) GTE(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	for i := range unitTagValue {
		if strings.Compare(unitTagValue[i], configValue) < 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// NE The tag value of the unit is not equal to the tag value configured by the web system. All unitTagValue elements must be satisfied for the request to pass.
func (e *stringExecutor) NE(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	for i := range unitTagValue {
		if strings.Compare(unitTagValue[i], configValue) == 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// REGEXP The tag value of the unit satisfies the regular expression configured by the web system. All unitTagValue elements must satisfy this to pass.
// such as unitTagValue[0] = `Hello Regexp` configValue = `^Hello` match == true
// In particular, when configValue is empty "", the result is true; when configValue is * "*", the result is false
func (e *stringExecutor) REGEXP(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	for i := range unitTagValue {
		if match, _ := regexp.MatchString(configValue, unitTagValue[i]); !match { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// IN Determine whether unitTagValue is in configValue. configValue is separated by ;
func (e *stringExecutor) IN(unitTagValue []string, configValue string) bool {
	// No user tag is carried, default is false
	if len(unitTagValue) == 0 {
		return false
	}
	configValueList := strings.Split(configValue, splitSeg)
	for i := range unitTagValue {
		inFlag := false
		for j := range configValueList {
			if unitTagValue[i] == configValueList[j] {
				inFlag = true
				break
			}
		}
		if !inFlag { // unitTagValue 有任何一个不存在于 configValueList，都为 false
			return false
		}
	}
	return true
}

// NotIN Determine whether unitTagValue is not in configValue. configValue is separated by ;
func (e *stringExecutor) NotIN(unitTagValue []string, configValue string) bool {
	// No user tag is carried, default is false
	if len(unitTagValue) == 0 {
		return false
	}
	configValueList := strings.Split(configValue, splitSeg)
	for i := range unitTagValue {
		inFlag := false
		for j := range configValueList {
			if unitTagValue[i] == configValueList[j] {
				inFlag = true
				break
			}
		}
		if inFlag { // unitTagValue 有任何一个存在于 configValueList，都为 false
			return false
		}
	}
	return true
}

// IsEmpty unitTagValue Determine whether the passed tag key or value is empty or value is not ""
func (e *stringExecutor) IsEmpty(unitTagValue []string, configValue string) bool {
	return IsEmptyTag(unitTagValue, configValue)
}

// IsNotEmpty unitTagValue Determine whether the passed tag key and value is not empty and value is not ""
func (e *stringExecutor) IsNotEmpty(unitTagValue []string, configValue string) bool {
	return !IsEmptyTag(unitTagValue, configValue)
}
