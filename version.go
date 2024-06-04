// Package tagutil ...
package tagutil

import (
	"regexp"
	"strings"

	"github.com/shopspring/decimal"
)

// versionExecutor 版本执行器，对版本号进行逻辑比较，版本号为
// 分割符为 . 例如 9.9.9.9，同时支持 9.9.9.9-beta
// 版本号用 . 分割
// configValue 后台需要校验是否符合该规则
// 特别的，9.9.9.9-beta > 9.9.9.9，跟 git tag 版本号规则有点区别，需要特别注意
// 版本号不允许比 '0' 还小的 ascii 字符出现，也不支持汉字等非 ascii 字符，如果存在，则结果不可信，
type versionExecutor struct{}

var (
	versionExecutorImpl = &versionExecutor{}
)

// compare 版本俩俩比较
// 如果 unitTagValueVersion >[版本号比较] configVersion 则返回 1
// 如果 unitTagValueVersion =[版本号比较] configVersion 则返回 0
// 如果 unitTagValueVersion <[版本号比较] configVersion 则返回 -1
// 例如：9.9.9 比较 9.8.8，结果返回 1
// 为了支持A-Z a-z，进制改为128进制，所以这个方法只支持长度为9的小版本号 例如 12aabb123.3454aabb12.123456712.123456789
// 超过的限定长度结果不可信，需要用 decimalCompare，为了性能，且大部分不会超过9位，故提供简单的 compare 方法
func (e *versionExecutor) compare(unitTagValueVersion string, configVersion string) int {
	for i, j := 0, 0; i < len(unitTagValueVersion) || j < len(configVersion); {
		var num1 int64 = 0
		for i < len(unitTagValueVersion) && unitTagValueVersion[i] != '.' {
			num1 = num1<<7 + int64(unitTagValueVersion[i]-'0')
			i++
		}
		var num2 int64 = 0
		for j < len(configVersion) && configVersion[j] != '.' {
			num2 = num2<<7 + int64(configVersion[j]-'0')
			j++
		}
		if num1 < num2 {
			return -1
		}
		if num1 > num2 {
			return 1
		}
		i++
		j++
	}
	return 0
}

// compareByDecimal 版本俩俩比较
// 如果 unitTagValueVersion >[版本号比较] configVersion 则返回 1
// 如果 unitTagValueVersion =[版本号比较] configVersion 则返回 0
// 如果 unitTagValueVersion <[版本号比较] configVersion 则返回 -1
// 例如：9.9.9 比较 9.8.8，结果返回 1
// 为了支持A-Z a-z，进制改为128进制，支持超长版本号比较，如 12aabb123424.3454aabb6876786755.1234567gg67867868.1234567gg76768768
func (e *versionExecutor) compareByDecimal(unitTagValueVersion string, configVersion string) int {
	for i, j := 0, 0; i < len(unitTagValueVersion) || j < len(configVersion); {
		var num1 = decimal.NewFromInt(0)
		for i < len(unitTagValueVersion) && unitTagValueVersion[i] != '.' {
			num1 = num1.Mul(decimal.NewFromInt(128)).Add(decimal.NewFromInt(int64(unitTagValueVersion[i] - '0')))
			i++
		}
		var num2 = decimal.NewFromInt(0)
		for j < len(configVersion) && configVersion[j] != '.' {
			num2 = num2.Mul(decimal.NewFromInt(128)).Add(decimal.NewFromInt(int64(configVersion[j] - '0')))
			j++
		}
		if num1.LessThan(num2) {
			return -1
		}
		if num1.GreaterThan(num2) {
			return 1
		}
		i++
		j++
	}
	return 0
}

// EQ unit 的标签值等于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
func (e *versionExecutor) EQ(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	// 遍历
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], configValue) != 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// LT unit 的标签值小于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
func (e *versionExecutor) LT(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	// 遍历
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], configValue) >= 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// LTE unit 的标签值小于等于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
func (e *versionExecutor) LTE(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	// 遍历
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], configValue) > 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// GT unit 的标签值大于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
func (e *versionExecutor) GT(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	// 遍历
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], configValue) <= 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// GTE unit 的标签值大于等于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
func (e *versionExecutor) GTE(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	// 遍历
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], configValue) < 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// NE unit 的标签值不等于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
func (e *versionExecutor) NE(unitTagValue []string, configValue string) bool {
	// 没有携带用户标签，默认 false
	if len(unitTagValue) == 0 {
		return false
	}
	// 匹配
	for i := range unitTagValue {
		// 任何一个 unitTagValue 不满足，结果为 false
		if e.compare(unitTagValue[i], configValue) == 0 {
			return false
		}
	}
	return true
}

// REGEXP unit 的标签值满足 web 系统配置的正则表达式，所有 unitTagValue 元素必须满足才通过
// 例如 unitTagValue[0] = `Hello Regexp` configValue = `^Hello` match == true
func (e *versionExecutor) REGEXP(unitTagValue []string, configValue string) bool {
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
func (e *versionExecutor) IN(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	configValueVersionList := strings.Split(configValue, splitSeg)
	for i := range unitTagValue {
		inFlag := false
		for j := range configValueVersionList {
			if e.compare(unitTagValue[i], configValueVersionList[j]) == 0 { // 任何一个 unitTagValue 不满足，结果为 false
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

// NotIN 判断 unitTagValue 是否都不在 configValue 中，configValue，用 ; 分割
func (e *versionExecutor) NotIN(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	configValueVersionList := strings.Split(configValue, splitSeg)
	for i := range unitTagValue {
		for j := range configValueVersionList {
			if e.compare(unitTagValue[i], configValueVersionList[j]) == 0 { // 任何一个 unitTagValue 不满足，结果为 false
				return false
			}
		}
	}
	return true
}

// LORO 左开右开区间判定，configValue 此时标识区间，用 : 分割，有且仅有一个 ：
// 这里不判定 left 是否小于 right，由 web 平台去控制
func (e *versionExecutor) LORO(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return false
	}
	configValueRange := strings.Split(configValue, rangeSplitSeg)
	if len(configValueRange) != 2 { // 用 : 分割，有且仅有一个 ：，分割left、right
		return false
	}
	minVersion := configValueRange[0]
	maxVersion := configValueRange[1]
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], minVersion) <= 0 || e.compare(unitTagValue[i], maxVersion) >= 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// LORC 左开右闭区间判定，configValue 此时标识区间，用 : 分割，有且仅有一个 ：
// 这里不判定 left 是否小于 right，由 web 平台去控制
func (e *versionExecutor) LORC(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return false
	}
	configValueRange := strings.Split(configValue, rangeSplitSeg)
	if len(configValueRange) != 2 { // 用 : 分割，有且仅有一个 ：，分割left、right
		return false
	}
	minVersion := configValueRange[0]
	maxVersion := configValueRange[1]
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], minVersion) <= 0 || e.compare(unitTagValue[i], maxVersion) > 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// LCRO 左闭右开区间判定，configValue 此时标识区间，用 : 分割，有且仅有一个 ：
// 这里不判定 left 是否小于 right，由 web 平台去控制
func (e *versionExecutor) LCRO(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return false
	}
	configValueRange := strings.Split(configValue, rangeSplitSeg)
	if len(configValueRange) != 2 { // 用 : 分割，有且仅有一个 ：，分割left、right
		return false
	}
	minVersion := configValueRange[0]
	maxVersion := configValueRange[1]
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], minVersion) < 0 || e.compare(unitTagValue[i], maxVersion) >= 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// LCRC 左闭右闭区间判定，configValue 此时标识区间，用 : 分割，有且仅有一个 ：
// 这里不判定 left 是否小于 right，由 web 平台去控制
func (e *versionExecutor) LCRC(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return false
	}
	configValueRange := strings.Split(configValue, rangeSplitSeg)
	if len(configValueRange) != 2 { // 用 : 分割，有且仅有一个 ：，分割left、right
		return false
	}
	minVersion := configValueRange[0]
	maxVersion := configValueRange[1]
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], minVersion) < 0 || e.compare(unitTagValue[i], maxVersion) > 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}
