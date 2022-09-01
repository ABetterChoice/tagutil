// Package tagutil ...
package tagutil

import (
	"regexp"
	"strings"
)

// stringExecutor 字符串表达式执行器，正常来说，字符串类型，用户标签值应该是一个字符串
// 由于用户标签是字符串数组，数组任何一个不满足表达式，都判定为 false
// 例如，用户标签值 [ab,aa] 要求大于 aa 的才能满足，虽然 ab 满足 ab > aa, 但 aa 不满足 aa > aa，故结果为 false
// 只有全部满足才为 true
type stringExecutor struct{}

var (
	stringExecutorImpl = stringExecutor{}
)

// EG unit 的标签值等于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
func (e *stringExecutor) EG(unitTagValue []string, configValue string) bool {
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

// LT unit 的标签值小于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
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

// LTE unit 的标签值小于等于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
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

// GT unit 的标签值大于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
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

// GTE unit 的标签值大于等于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
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

// NE unit 的标签值不等于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
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

// REGEXP unit 的标签值满足 web 系统配置的正则表达式，所有 unitTagValue 元素必须满足才通过
// 例如 unitTagValue[0] = `Hello Regexp` configValue = `^Hello` match == true
// 特别的，当 configValue 为空时 ""，结果为 true；当 configValue 为*时 "*"，结果为 false
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

// IN 判断 unitTagValue 是否都在 configValue 中，configValue，用 ; 分割
func (e *stringExecutor) IN(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
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

// NotIN 判断 unitTagValue 是否都不在 configValue 中，configValue，用 ; 分割
func (e *stringExecutor) NotIN(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
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
