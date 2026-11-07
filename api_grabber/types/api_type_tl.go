package types

type ApiTypeTL struct {
	Name        string
	Description []string
	Returns     []string
	Fields      []FieldTL
	Subtypes    []string
	TypeIds     *TypeIdsDescriptor
	IsSend      bool
}

func (x *ApiTypeTL) GetFields() []FieldTL {
	return x.Fields
}

func (x *ApiTypeTL) GetSubTypes() []string {
	return x.Subtypes
}

func (x *ApiTypeTL) GetName() string {
	return x.Name
}

func (x *ApiTypeTL) GetType() string {
	return "types"
}

func (x *ApiTypeTL) GetTypeIds() *TypeIdsDescriptor {
	return x.TypeIds
}

func (x *ApiTypeTL) IsSendMethod() bool {
	return x.IsSend
}

func (x *ApiTypeTL) GetReturns() []string {
	return x.Returns
}

func (x *ApiTypeTL) GetDescription() []string {
	return x.Description
}
