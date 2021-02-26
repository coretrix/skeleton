package entity

import "github.com/summer-solutions/orm"

func Init(registry *orm.Registry) {
	registry.RegisterEntity(
		&AdminUserEntity{}, &GoogleOSSBucketCounterEntity{},
	)

	//registry.RegisterEnumStruct("entity.TransactionStatusAll", TransactionStatusAll)
}
