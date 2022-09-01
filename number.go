// Package tagutil ...
package tagutil

import (
	"github.com/shopspring/decimal"
	"strings"
)

// numberExecutor 数值类型计算器
type numberExecutor struct{}

var (
	numberExecutorImpl = &numberExecutor{}
)

// EQ unit 的标签值等于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
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

// LT unit 的标签值小于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
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

// LTE unit 的标签值小于等于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
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

// GT unit 的标签值大于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
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

// GTE unit 的标签值大于等于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
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

// NE unit 的标签值不等于 web 系统配置的标签值，所有 unitTagValue 元素必须满足才通过
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

// IN 判断 unitTagValue 是否都在 configValue 中，configValue，用 ; 分割
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

// NotIN 判断 unitTagValue 是否都不在 configValue 中，configValue，用 ; 分割
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

// LORO 左开右开区间判定，configValue 此时标识区间，用 : 分割，有且仅有一个 ：
// 这里不判定 left 是否小于 right，由 web 平台去控制
func (e *numberExecutor) LORO(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return false
	}
	configValueRange := strings.Split(configValue, rangeSplitSeg)
	if len(configValueRange) != 2 { // 用 : 分割，有且仅有一个 ：，分割left、right
		return false
	}
	left, err := decimal.NewFromString(configValueRange[0])
	if err != nil {
		return false
	}
	right, err := decimal.NewFromString(configValueRange[1])
	if err != nil {
		return false
	}
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

// LORC 左开右闭区间判定，configValue 此时标识区间，用 : 分割，有且仅有一个 ：
// 这里不判定 left 是否小于 right，由 web 平台去控制
func (e *numberExecutor) LORC(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return false
	}
	configValueRange := strings.Split(configValue, rangeSplitSeg)
	if len(configValueRange) != 2 { // 用 : 分割，有且仅有一个 ：，分割left、right
		return false
	}
	left, err := decimal.NewFromString(configValueRange[0])
	if err != nil {
		return false
	}
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

// LCRO 左闭右开区间判定，configValue 此时标识区间，用 : 分割，有且仅有一个 ：
// 这里不判定 left 是否小于 right，由 web 平台去控制
func (e *numberExecutor) LCRO(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return false
	}
	configValueRange := strings.Split(configValue, rangeSplitSeg)
	if len(configValueRange) != 2 { // 用 : 分割，有且仅有一个 ：，分割left、right
		return false
	}
	left, err := decimal.NewFromString(configValueRange[0])
	if err != nil {
		return false
	}
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

// LCRC 左闭右闭区间判定，configValue 此时标识区间，用 : 分割，有且仅有一个 ：
// 这里不判定 left 是否小于 right，由 web 平台去控制
func (e *numberExecutor) LCRC(unitTagValue []string, configValue string) bool {
	if len(unitTagValue) == 0 {
		return false
	}
	configValueRange := strings.Split(configValue, rangeSplitSeg)
	if len(configValueRange) != 2 { // 用 : 分割，有且仅有一个 ：，分割left、right
		return false
	}
	left, err := decimal.NewFromString(configValueRange[0])
	if err != nil {
		return false
	}
	right, err := decimal.NewFromString(configValueRange[1])
	if err != nil {
		return false
	}
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
