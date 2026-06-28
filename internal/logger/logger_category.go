package logger

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	Io              Category = "Io"
	Internal        Category = "Internal"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
	Prometheus      Category = "Prometheus"
	Baskets         Category = "Baskets"
	Stores          Category = "Stores"
	Customers       Category = "Customers"
	Notifications   Category = "Notifications"
	Payments        Category = "Payments"
)

const (
	// General
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"

	// Postgres
	Migration SubCategory = "Migration"
	Select    SubCategory = "Select"
	Rollback  SubCategory = "Rollback"
	Update    SubCategory = "Update"
	Delete    SubCategory = "Delete"
	Insert    SubCategory = "Insert"

	// Internal
	Api                 SubCategory = "Api"
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound             = "DefaultRoleNotFound"

	// Validation
	MobileValidation   SubCategory = "MobileValidation"
	PasswordValidation SubCategory = "PasswordValidation"

	// IO
	RemoveFile SubCategory = "RemoveFile"

	// basket sub category
	StartBasket      SubCategory = "StartBasket"
	CancelBasket     SubCategory = "CancelBasket"
	CheckoutBasket   SubCategory = "CheckoutBasket"
	BasketAddItem    SubCategory = "BasketAddItem"
	BasketRemoveItem SubCategory = "BasketRemoveItem"
	GetBasket        SubCategory = "GetBasket"

	// Stores
	CreateStore            SubCategory = "CreateStore"
	GetStore               SubCategory = "GetStore"
	GetStores              SubCategory = "GetStores"
	GetParticipatingStores SubCategory = "GetParticipatingStores"
	EnableParticipation    SubCategory = "EnableParticipation"
	DisableParticipation   SubCategory = "DisableParticipation"
	AddProduct             SubCategory = "AddProduct"
	RemoveProduct          SubCategory = "RemoveProduct"
	GetCatalog             SubCategory = "GetCatalog"
	GetProduct             SubCategory = "GetProduct"

	// Customer
	RegisterCustomer  SubCategory = "RegisterCustomer"
	AuthorizeCustomer SubCategory = "AuthorizeCustomer"
	GetCustomer       SubCategory = "GetCustomer"
	EnableCustomer    SubCategory = "EnableCustomer"
	DisableCustomer   SubCategory = "DisableCustomer"

	// Notifications
	NotifyOrderCreated  SubCategory = "NotifyOrderCreated"
	NotifyOrderCanceled SubCategory = "NotifyOrderCanceled"
	NotifyOrderReady    SubCategory = "NotifyOrderReady"

	// Payments
	AuthorizePayment SubCategory = "AuthorizePayment"
	ConfirmPayment   SubCategory = "ConfirmPayment"
	CreateInvoice    SubCategory = "CreateInvoice"
	AdjustInvoice    SubCategory = "AdjustInvoice"
	PayInvoice       SubCategory = "PayInvoice"
	CancelInvoice    SubCategory = "CancelInvoice"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "Logger"
	ClientIp     ExtraKey = "ClientIp"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	RequestBody  ExtraKey = "RequestBody"
	ResponseBody ExtraKey = "ResponseBody"
	ErrorMessage ExtraKey = "ErrorMessage"
)
