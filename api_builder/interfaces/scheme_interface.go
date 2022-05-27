package interfaces

import "BotApiCompiler/api_grabber/types"

type SchemeInterface interface {
	*types.ApiTypeTL | *types.ApiMethodTL
	GetFields() []types.FieldTL
	GetSubTypes() []string
	GetName() string
	GetType() string
	GetTypeIds() *types.TypeIdsDescriptor
	GetReturns() []string
	IsSendMethod() bool
}
