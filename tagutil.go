// Package tagutil ...
package tagutil

import (
	protoctabcacheserver "github.com/abetterchoice/protoc_cache_server"
)

// IsHit Whether the label expression is true. If it is an unsupported label expression, the result is false.
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
		protoctabcacheserver.Operator_OPERATOR_EQ:        stringExecutorImpl.EQ,
		protoctabcacheserver.Operator_OPERATOR_LT:        stringExecutorImpl.LT,
		protoctabcacheserver.Operator_OPERATOR_LTE:       stringExecutorImpl.LTE,
		protoctabcacheserver.Operator_OPERATOR_GT:        stringExecutorImpl.GT,
		protoctabcacheserver.Operator_OPERATOR_GTE:       stringExecutorImpl.GTE,
		protoctabcacheserver.Operator_OPERATOR_NE:        stringExecutorImpl.NE,
		protoctabcacheserver.Operator_OPERATOR_REGULAR:   stringExecutorImpl.REGEXP,
		protoctabcacheserver.Operator_OPERATOR_IN:        stringExecutorImpl.IN,
		protoctabcacheserver.Operator_OPERATOR_NOT_IN:    stringExecutorImpl.NotIN,
		protoctabcacheserver.Operator_OPERATOR_EMPTY:     stringExecutorImpl.IsEmpty,
		protoctabcacheserver.Operator_OPERATOR_NOT_EMPTY: stringExecutorImpl.IsNotEmpty,
		protoctabcacheserver.Operator_OPERATOR_LIKE:      stringExecutorImpl.Like,
		protoctabcacheserver.Operator_OPERATOR_NOT_LIKE:  stringExecutorImpl.NotLike,
	},
	protoctabcacheserver.TagType_TAG_TYPE_NUMBER: {
		protoctabcacheserver.Operator_OPERATOR_EQ:        numberExecutorImpl.EQ,
		protoctabcacheserver.Operator_OPERATOR_LT:        numberExecutorImpl.LT,
		protoctabcacheserver.Operator_OPERATOR_LTE:       numberExecutorImpl.LTE,
		protoctabcacheserver.Operator_OPERATOR_GT:        numberExecutorImpl.GT,
		protoctabcacheserver.Operator_OPERATOR_GTE:       numberExecutorImpl.GTE,
		protoctabcacheserver.Operator_OPERATOR_NE:        numberExecutorImpl.NE,
		protoctabcacheserver.Operator_OPERATOR_IN:        numberExecutorImpl.IN,
		protoctabcacheserver.Operator_OPERATOR_NOT_IN:    numberExecutorImpl.NotIN,
		protoctabcacheserver.Operator_OPERATOR_LORO:      numberExecutorImpl.LORO,
		protoctabcacheserver.Operator_OPERATOR_LORC:      numberExecutorImpl.LORC,
		protoctabcacheserver.Operator_OPERATOR_LCRO:      numberExecutorImpl.LCRO,
		protoctabcacheserver.Operator_OPERATOR_LCRC:      numberExecutorImpl.LCRC,
		protoctabcacheserver.Operator_OPERATOR_EMPTY:     numberExecutorImpl.IsEmpty,
		protoctabcacheserver.Operator_OPERATOR_NOT_EMPTY: numberExecutorImpl.IsNotEmpty,
	},
	protoctabcacheserver.TagType_TAG_TYPE_BOOLEAN: {
		protoctabcacheserver.Operator_OPERATOR_EQ:        booleanExecutorImpl.EQ,
		protoctabcacheserver.Operator_OPERATOR_EMPTY:     booleanExecutorImpl.IsEmpty,
		protoctabcacheserver.Operator_OPERATOR_NOT_EMPTY: booleanExecutorImpl.IsNotEmpty,
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
	splitSeg      = ";" // Separator, multiple value separation
	rangeSplitSeg = ":" // There is only one interval separator:, and the left and right are number types, representing left and right
)
