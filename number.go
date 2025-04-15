// Package tagutil ...
package tagutil

import (
	"strings"

	"github.com/shopspring/decimal"
)

// numberExecutor Numeric Calculator
type numberExecutor struct{}

var (
	numberExecutorImpl = &numberExecutor{}
)

// The tag value of the EQ unit is equal to the tag value configured by the web system. All unitTagValue elements must be satisfied to pass.
func (e *numberExecutor) EQ(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	configValueNumber, err := decimal.NewFromString(configValue)
	if err != nil {
		return false
	}
	for i := range unitTagValue {
		unitTagValueNumber, err := decimal.NewFromString(unitTagValue[i])
		if err != nil {
			return false
		}
		if !unitTagValueNumber.Equal(configValueNumber) { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// The tag value of the LT unit is less than the tag value configured by the web system. All unitTagValue elements must meet the requirement to pass.
func (e *numberExecutor) LT(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	configValueNumber, err := decimal.NewFromString(configValue)
	if err != nil {
		return false
	}
	for i := range unitTagValue {
		unitTagValueNumber, err := decimal.NewFromString(unitTagValue[i])
		if err != nil {
			return false
		}
		if !unitTagValueNumber.LessThan(configValueNumber) { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// The tag value of the LTE unit is less than or equal to the tag value configured by the web system. All unitTagValue elements must meet the requirement to pass.
func (e *numberExecutor) LTE(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	configValueNumber, err := decimal.NewFromString(configValue)
	if err != nil {
		return false
	}
	for i := range unitTagValue {
		unitTagValueNumber, err := decimal.NewFromString(unitTagValue[i])
		if err != nil {
			return false
		}
		if !unitTagValueNumber.LessThanOrEqual(configValueNumber) { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// The tag value of the GT unit is greater than the tag value configured by the web system. All unitTagValue elements must meet this requirement to pass.
func (e *numberExecutor) GT(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	configValueNumber, err := decimal.NewFromString(configValue)
	if err != nil {
		return false
	}
	for i := range unitTagValue {
		unitTagValueNumber, err := decimal.NewFromString(unitTagValue[i])
		if err != nil {
			return false
		}
		if !unitTagValueNumber.GreaterThan(configValueNumber) { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// The tag value of the GTE unit is greater than or equal to the tag value configured by the web system. All unitTagValue elements must meet this requirement to pass.
func (e *numberExecutor) GTE(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	configValueNumber, err := decimal.NewFromString(configValue)
	if err != nil {
		return false
	}
	for i := range unitTagValue {
		unitTagValueNumber, err := decimal.NewFromString(unitTagValue[i])
		if err != nil {
			return false
		}
		if !unitTagValueNumber.GreaterThanOrEqual(configValueNumber) { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// The tag value of the NE unit is not equal to the tag value configured by the web system. All unitTagValue elements must be satisfied to pass.
func (e *numberExecutor) NE(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	configValueNumber, err := decimal.NewFromString(configValue)
	if err != nil {
		return false
	}
	for i := range unitTagValue {
		unitTagValueNumber, err := decimal.NewFromString(unitTagValue[i])
		if err != nil {
			return false
		}
		if unitTagValueNumber.Equal(configValueNumber) { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// IN Determines whether unitTagValue is in configValue. configValue is separated by ;
func (e *numberExecutor) IN(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	configValueList := strings.Split(configValue, splitSeg)
	var configValueNumberList = make([]*decimal.Decimal, len(configValueList))
	for i := range configValueList { // 预处理 configValueList
		itemNumber, err := decimal.NewFromString(configValueList[i])
		if err != nil {
			return false
		}
		configValueNumberList[i] = &itemNumber
	}
	for i := range unitTagValue {
		unitTagValueNumber, err := decimal.NewFromString(unitTagValue[i])
		if err != nil {
			return false
		}
		inFlag := false
		for j := range configValueNumberList {
			if configValueNumberList[j].Equal(unitTagValueNumber) {
				inFlag = true
				break
			}
		}
		if !inFlag {
			return false // 任何一个 unitTagValue 不满足，结果为 false
		}
	}
	return true
}

// NotIN Determine whether unitTagValue is not in configValue. configValue is separated by ;
func (e *numberExecutor) NotIN(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 { // 没有携带用户标签，默认 false
		return false
	}
	configValueList := strings.Split(configValue, splitSeg)
	var configValueNumberList = make([]*decimal.Decimal, len(configValueList))
	for i := range configValueList { // 预处理 configValueList
		itemNumber, err := decimal.NewFromString(configValueList[i])
		if err != nil {
			return false
		}
		configValueNumberList[i] = &itemNumber
	}
	for i := range unitTagValue {
		unitTagValueNumber, err := decimal.NewFromString(unitTagValue[i])
		if err != nil {
			return false
		}
		inFlag := false
		for j := range configValueNumberList {
			if configValueNumberList[j].Equal(unitTagValueNumber) {
				inFlag = true
				break
			}
		}
		if inFlag {
			return false // 任何一个 unitTagValue 不满足，结果为 false
		}
	}
	return true
}

// LORO The left-open and right-open intervals are determined. configValue identifies the interval at this time, separated by : , and there is only one :
// Here we do not determine whether left is smaller than right, which is controlled by the web platform.
func (e *numberExecutor) LORO(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return false
	}
	configValueRange := strings.Split(configValue, rangeSplitSeg)
	if len(configValueRange) != 2 { // 用 : 分割，有且仅有一个 ：，分割left、right
		return false
	}
	// Left
	left, err := decimal.NewFromString(configValueRange[0])
	if err != nil {
		return false
	}
	// Right
	right, err := decimal.NewFromString(configValueRange[1])
	if err != nil {
		return false
	}
	// Traversal
	for i := range unitTagValue {
		item, err := decimal.NewFromString(unitTagValue[i])
		if err != nil {
			return false
		}
		if item.LessThanOrEqual(left) || item.GreaterThanOrEqual(right) { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// LORC The left-open and right-closed interval is determined. configValue identifies the interval at this time, separated by :, and there is only one :
// Here we do not determine whether left is smaller than right, which is controlled by the web platform.
func (e *numberExecutor) LORC(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return false
	}
	configValueRange := strings.Split(configValue, rangeSplitSeg)
	if len(configValueRange) != 2 { // 用 : 分割，有且仅有一个 ：，分割left、right
		return false
	}
	// Left
	left, err := decimal.NewFromString(configValueRange[0])
	if err != nil {
		return false
	}
	// Right
	right, err := decimal.NewFromString(configValueRange[1])
	if err != nil {
		return false
	}
	for i := range unitTagValue {
		item, err := decimal.NewFromString(unitTagValue[i])
		if err != nil {
			return false
		}
		if item.LessThanOrEqual(left) || item.GreaterThan(right) { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// LCRO The left closed and right open interval is determined. configValue identifies the interval at this time, separated by :, and there is only one :
// Here we do not determine whether left is smaller than right, which is controlled by the web platform.
func (e *numberExecutor) LCRO(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return false
	}
	configValueRange := strings.Split(configValue, rangeSplitSeg)
	if len(configValueRange) != 2 { // 用 : 分割，有且仅有一个 ：，分割left、right
		return false
	}
	// Left
	left, err := decimal.NewFromString(configValueRange[0])
	if err != nil {
		return false
	}
	// Right
	right, err := decimal.NewFromString(configValueRange[1])
	if err != nil {
		return false
	}
	for i := range unitTagValue {
		item, err := decimal.NewFromString(unitTagValue[i])
		if err != nil {
			return false
		}
		if item.LessThan(left) || item.GreaterThanOrEqual(right) { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// LCRC Left-closed and right-closed interval determination, configValue identifies the interval at this time, separated by :, there is only one :
// Here we do not determine whether left is smaller than right, which is controlled by the web platform.
func (e *numberExecutor) LCRC(unitTagValue []string, configValue string) bool {
	// If there is no value, it is always false.
	if len(unitTagValue) == 0 {
		return false
	}
	configValueRange := strings.Split(configValue, rangeSplitSeg)
	if len(configValueRange) != 2 { // 用 : 分割，有且仅有一个 ：，分割left、right
		return false
	}
	// Left
	left, err := decimal.NewFromString(configValueRange[0])
	if err != nil {
		return false
	}
	// Right
	right, err := decimal.NewFromString(configValueRange[1])
	if err != nil {
		return false
	}
	// Range
	for i := range unitTagValue {
		item, err := decimal.NewFromString(unitTagValue[i])
		if err != nil {
			return false
		}
		if item.LessThan(left) || item.GreaterThan(right) { // 任何一个 unitTagValue 不满足，结果为 false
			return false
		}
	}
	return true
}

// IsEmpty unitTagValue Determine whether the passed tag key or value is empty or value is not ""
func (e *numberExecutor) IsEmpty(unitTagValue []string, configValue string) bool {
	return IsEmptyTag(unitTagValue, configValue)
}

// IsNotEmpty unitTagValue Determine whether the passed tag key and value is not empty and value is not ""
func (e *numberExecutor) IsNotEmpty(unitTagValue []string, configValue string) bool {
	return !IsEmptyTag(unitTagValue, configValue)
}
