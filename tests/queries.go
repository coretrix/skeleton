package tests

//func CreateUser(_ *testing.T, row map[string]interface{}) *entity.UserEntity {
//	ormService := ioc.GetOrmEngineGlobalService()
//
//	insertedEntity := &entity.UserEntity{
//		Email:     "hymn@abv.bg",
//		CreatedAt: time.Now(),
//	}
//
//	if len(row) != 0 {
//		for field, value := range row {
//			switch field {
//			case "Email":
//				insertedEntity.Email = value.(string)
//			case "Name":
//				insertedEntity.Name = value.(string)
//			}
//		}
//	}
//
//	ormService.Flush(insertedEntity)
//
//	return insertedEntity
//}
