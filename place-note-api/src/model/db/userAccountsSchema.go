package modelDb

import placeNote "placeNote/src/gen/proto"

// UserAccountsEntity user_accountsのEntity
type UserAccountsEntity struct {
	ID            string               `json:"id" bson:"_id"`
	UserSettingId string               `json:"userSettingId" bson:"user_setting_id"`
	Name          string               `json:"name" bson:"name"`
	AuthMethod    placeNote.AuthMethod `json:"authMethod" bson:"auth_method"`
	Email         string               `json:"email" bson:"email"`
	Password      string               `json:"password" bson:"password"`
	Gmail         string               `json:"gmail" bson:"gmail"`
}
