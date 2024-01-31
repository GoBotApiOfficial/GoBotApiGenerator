package types

type ApiMethodTL struct {
	Name        string
	Description []string
	Params      []FieldTL
	Returns     []string
	Subtypes    []string
}

func (x *ApiMethodTL) GetFields() []FieldTL {
	return x.Params
}

func (x *ApiMethodTL) GetSubTypes() []string {
	return x.Subtypes
}

func (x *ApiMethodTL) GetName() string {
	return x.Name
}

func (x *ApiMethodTL) GetType() string {
	return "methods"
}

func (x *ApiMethodTL) GetTypeIds() *TypeIdsDescriptor {
	return &TypeIdsDescriptor{}
}

func (x *ApiMethodTL) IsSendMethod() bool {
	return true
}

func (x *ApiMethodTL) GetReturns() []string {
	return x.Returns
}

func (x *ApiMethodTL) GetDescription() []string {
	return x.Description
}
