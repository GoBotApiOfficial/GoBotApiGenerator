package types

type ApiTL struct {
	Methods map[string]*ApiMethodTL
	Types   map[string]*ApiTypeTL
	Version string
}
