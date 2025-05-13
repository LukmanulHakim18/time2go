package constant

import (
	"github.com/LukmanulHakim18/time2go/model"
	"github.com/LukmanulHakim18/time2go/util"
)

const (
	KEY_TYPE_INDEX   util.KeyType = "index"
	KEY_TYPE_TRIGGER util.KeyType = "trigger"
	KEY_TYPE_DATA    util.KeyType = "data"
)

const (
	RETRY_POLICY_TYPE_FIXED       model.RetryPolicyType = "fixed"
	RETRY_POLICY_TYPE_EXPONENTIAL model.RetryPolicyType = "exponential"
)
