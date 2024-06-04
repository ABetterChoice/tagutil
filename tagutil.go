// Package tagutil ...
package tagutil

import (
	protoctabcacheserver "git.tencent.com/abetterchoice/protocol/protoc_cache_server"
)

// IsHit 标签表达式是否为 true，如果是不支持的标签表达式，结果为 false
func IsHit(tagType protoctabcacheserver.TagType, operator protoctabcacheserver.Operator, unitTagValue []string, configValue string) bool {
	e, ok := executorIndex[tagType][operator]
	if !ok {
		return false
	}
	return e(unitTagValue, configValue)
}

type executor func(unitTagValue []string, configValue string) bool

var executorIndex = map[protoctabcacheserver.TagType]map[protoctabcacheserver.Operator]executor{
	protoctabcacheserver.TagType_TAG_TYPE_STRING: {
		protoctabcacheserver.Operator_OPERATOR_EQ:      stringExecutorImpl.EQ,
		protoctabcacheserver.Operator_OPERATOR_LT:      stringExecutorImpl.LT,
		protoctabcacheserver.Operator_OPERATOR_LTE:     stringExecutorImpl.LTE,
		protoctabcacheserver.Operator_OPERATOR_GT:      stringExecutorImpl.GT,
		protoctabcacheserver.Operator_OPERATOR_GTE:     stringExecutorImpl.GTE,
		protoctabcacheserver.Operator_OPERATOR_NE:      stringExecutorImpl.NE,
		protoctabcacheserver.Operator_OPERATOR_REGULAR: stringExecutorImpl.REGEXP,
		protoctabcacheserver.Operator_OPERATOR_IN:      stringExecutorImpl.IN,
		protoctabcacheserver.Operator_OPERATOR_NOT_IN:  stringExecutorImpl.NotIN,
	},
	protoctabcacheserver.TagType_TAG_TYPE_NUMBER: {
		protoctabcacheserver.Operator_OPERATOR_EQ:     numberExecutorImpl.EQ,
		protoctabcacheserver.Operator_OPERATOR_LT:     numberExecutorImpl.LT,
		protoctabcacheserver.Operator_OPERATOR_LTE:    numberExecutorImpl.LTE,
		protoctabcacheserver.Operator_OPERATOR_GT:     numberExecutorImpl.GT,
		protoctabcacheserver.Operator_OPERATOR_GTE:    numberExecutorImpl.GTE,
		protoctabcacheserver.Operator_OPERATOR_NE:     numberExecutorImpl.NE,
		protoctabcacheserver.Operator_OPERATOR_IN:     numberExecutorImpl.IN,
		protoctabcacheserver.Operator_OPERATOR_NOT_IN: numberExecutorImpl.NotIN,
		protoctabcacheserver.Operator_OPERATOR_LORO:   numberExecutorImpl.LORO,
		protoctabcacheserver.Operator_OPERATOR_LORC:   numberExecutorImpl.LORC,
		protoctabcacheserver.Operator_OPERATOR_LCRO:   numberExecutorImpl.LCRO,
		protoctabcacheserver.Operator_OPERATOR_LCRC:   numberExecutorImpl.LCRC,
	},
	protoctabcacheserver.TagType_TAG_TYPE_BOOLEAN: {
		protoctabcacheserver.Operator_OPERATOR_EQ: booleanExecutorImpl.EQ,
	},
	protoctabcacheserver.TagType_TAG_TYPE_VERSION: {
		protoctabcacheserver.Operator_OPERATOR_EQ:      versionExecutorImpl.EQ,
		protoctabcacheserver.Operator_OPERATOR_LT:      versionExecutorImpl.LT,
		protoctabcacheserver.Operator_OPERATOR_LTE:     versionExecutorImpl.LTE,
		protoctabcacheserver.Operator_OPERATOR_GT:      versionExecutorImpl.GT,
		protoctabcacheserver.Operator_OPERATOR_GTE:     versionExecutorImpl.GTE,
		protoctabcacheserver.Operator_OPERATOR_NE:      versionExecutorImpl.NE,
		protoctabcacheserver.Operator_OPERATOR_REGULAR: versionExecutorImpl.REGEXP,
		protoctabcacheserver.Operator_OPERATOR_IN:      versionExecutorImpl.IN,
		protoctabcacheserver.Operator_OPERATOR_NOT_IN:  versionExecutorImpl.NotIN,
		protoctabcacheserver.Operator_OPERATOR_LORO:    versionExecutorImpl.LORO,
		protoctabcacheserver.Operator_OPERATOR_LORC:    versionExecutorImpl.LORC,
		protoctabcacheserver.Operator_OPERATOR_LCRO:    versionExecutorImpl.LCRO,
		protoctabcacheserver.Operator_OPERATOR_LCRC:    versionExecutorImpl.LCRC,
	},
	protoctabcacheserver.TagType_TAG_TYPE_EMPTY: {
		protoctabcacheserver.Operator_OPERATOR_EMPTY:     emptyExecutorImpl.IsEmpty,
		protoctabcacheserver.Operator_OPERATOR_NOT_EMPTY: emptyExecutorImpl.IsNotEmpty,
	},
	protoctabcacheserver.TagType_TAG_TYPE_SET: {
		protoctabcacheserver.Operator_OPERATOR_SUB_SET:   setExecutorImpl.IsSubset,
		protoctabcacheserver.Operator_OPERATOR_SUPER_SET: setExecutorImpl.IsSuperset,
		protoctabcacheserver.Operator_OPERATOR_EQ:        setExecutorImpl.EQ,
	},
}

const (
	splitSeg      = ";" // 分割符，多个值分割
	rangeSplitSeg = ":" // 区间分割符，有且仅有一个:，左右分别是 number 类型，代表 left、right
)
