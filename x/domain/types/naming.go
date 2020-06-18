package types

// Module names
const (
	// ModuleName is the name of the module
	ModuleName = "domain"
	// DomainStore key defines the store key used to store domains information
	DomainStoreKey = "domain"
	// AccountStoreKey defines the store key used to store account information
	AccountStoreKey = "account"
	// IndexStoreKey defines the store key used to store indexing information
	IndexStoreKey = ModuleName + "index"
	// RouterKey defines the path used to interact with the domain module
	RouterKey    = ModuleName
	QuerierAlias = "starname"
	// QuerierRoute defines the query path used to interact with the domain module
	QuerierRoute = ModuleName
	// DefaultParamSpace defines the key for the default param space
	DefaultParamSpace = ModuleName
)
