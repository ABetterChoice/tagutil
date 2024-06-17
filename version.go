// Package tagutil ...
package tagutil

import (
	"regexp"
	"strings"

	"github.com/shopspring/decimal"
)

// versionExecutor The version executor performs a logical comparison on the version number. The version number is
// The separator is . For example, 9.9.9.9, and 9.9.9.9-beta is also supported.
// The version number is separated by .
// configValue The background needs to verify whether it complies with the rule
// In particular, 9.9.9.9-beta > 9.9.9.9 is slightly different from the git tag version number rule, so you need to pay special attention to it.
// The version number does not allow ASCII characters smaller than '0' to appear, and does not support non-ASCII characters such as Chinese characters. If they exist, the result is unreliable.
type versionExecutor struct{}

var (
	versionExecutorImpl = &versionExecutor{}
)

// compare Compare versions
// If unitTagValueVersion > [version number comparison] configVersion, return 1
// If unitTagValueVersion = [version number comparison] configVersion, return 0
// If unitTagValueVersion < [version number comparison] configVersion, return -1
// For example: 9.9.9 compares 9.8.8, and the result returns 1
// In order to support A-Z a-z, the base is changed to 128, so this method only supports small version numbers with a length of 9, such as 12aabb123.3454aabb12.123456712.123456789
// The result exceeding the limited length is unreliable, and decimalCompare needs to be used. For performance reasons, most of them will not exceed 9 digits, so a simple compare method is provided
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

// compareByDecimal Compare versions
// If unitTagValueVersion > [version number comparison] configVersion, return 1
// If unitTagValueVersion = [version number comparison] configVersion, return 0
// If unitTagValueVersion < [version number comparison] configVersion, return -1
// For example: 9.9.9 compares 9.8.8, the result returns 1
// In order to support A-Z a-z, the base is changed to base 128, and ultra-long version number comparison is supported, such as 12aabb123424.3454aabb6876786755.1234567gg67867868.1234567gg76768768
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

// EQ The tag value of the unit is equal to the tag value configured by the web system. All unitTagValue elements must meet this requirement to pass.
func (e *versionExecutor) EQ(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], configValue) != 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// LT The tag value of the unit is less than the tag value configured by the web system. All unitTagValue elements must meet this requirement to pass.
func (e *versionExecutor) LT(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], configValue) >= 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// LTE The tag value of the unit is less than or equal to the tag value configured by the web system. All unitTagValue elements must meet this requirement to pass.
func (e *versionExecutor) LTE(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], configValue) > 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// GT The tag value of the unit is greater than the tag value configured by the web system. All unitTagValue elements must satisfy this condition to pass.
func (e *versionExecutor) GT(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], configValue) <= 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// GTE The tag value of the unit is greater than or equal to the tag value configured by the web system. All unitTagValue elements must meet this requirement to pass.
func (e *versionExecutor) GTE(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	for i := range unitTagValue {
		if e.compare(unitTagValue[i], configValue) < 0 { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// NE The tag value of the unit is not equal to the tag value configured by the web system. All unitTagValue elements must be satisfied for the request to pass.
func (e *versionExecutor) NE(unitTagValue []string, configValue string) bool {
	// No user tag is carried, default is false
	if len(unitTagValue) == 0 {
		return false
	}

	for i := range unitTagValue {
		// If any unitTagValue is not satisfied, the result is false
		if e.compare(unitTagValue[i], configValue) == 0 {
			return false
		}
	}
	return true
}

// REGEXP The tag value of unit satisfies the regular expression configured by the web system. All unitTagValue elements must satisfy the regular expression to pass.
// For example, unitTagValue[0] = `Hello Regexp` configValue = `^Hello` match == true
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

// IN Determine whether unitTagValue is in configValue. configValue is separated by ;
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

// NotIN Determine whether unitTagValue is not in configValue. configValue is separated by ;
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

// LORO The left-open and right-open intervals are determined. configValue identifies the interval at this time, separated by : , and there is only one :
// Here we do not determine whether left is smaller than right, which is controlled by the web platform.
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

// LORC The left-open and right-closed interval is determined. configValue identifies the interval at this time, separated by :, and there is only one :
// Here we do not determine whether left is smaller than right, which is controlled by the web platform.
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

// LCRO The left closed and right open interval is determined. configValue identifies the interval at this time, separated by :, and there is only one :
// Here we do not determine whether left is smaller than right, which is controlled by the web platform.
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

// LCRC Left-closed and right-closed interval determination, configValue identifies the interval at this time, separated by :, there is only one :
// Here we do not determine whether left is smaller than right, which is controlled by the web platform.
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
