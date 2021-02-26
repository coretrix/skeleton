package entity

import (
	"github.com/summer-solutions/orm"
)

type GoogleOSSBucketCounterEntity struct {
	orm.ORM `orm:"table=google_oss_buckets_counters"`
	ID      uint64
	Counter uint64 `orm:"required"`
}
