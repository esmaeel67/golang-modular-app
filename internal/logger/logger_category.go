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
	Orders          Category = "Orders"
	Depot           Category = "Depot"
	Search          Category = "Search"
	Stream          Category = "Stream"
	COSEC           Category = "COSEC"
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
	// baskets events
	OnBasketStarted     SubCategory = "OnBasketStarted"
	OnBasketItemAdded   SubCategory = "OnBasketItemAdded"
	OnBasketItemRemoved SubCategory = "OnBasketItemRemoved"
	OnBasketCanceled    SubCategory = "OnBasketCanceled"
	OnBasketCheckedOut  SubCategory = "OnBasketCheckedOut"

	// Stores
	CreateStore            SubCategory = "CreateStore"
	GetStore               SubCategory = "GetStore"
	GetStores              SubCategory = "GetStores"
	GetParticipatingStores SubCategory = "GetParticipatingStores"
	EnableParticipation    SubCategory = "EnableParticipation"
	DisableParticipation   SubCategory = "DisableParticipation"
	GetCatalog             SubCategory = "GetCatalog"
	GetProduct             SubCategory = "GetProduct"
	ProductRebranded       SubCategory = "ProductRebranded"
	ProductPriceChanged    SubCategory = "ProductPriceChanged"
	ProductAdded           SubCategory = "ProductAdded"
	ProductRemoved         SubCategory = "ProductRemoved"
	// stores events
	StoreHandleEvent           SubCategory = "StoreHandleEvent"
	StoreCreated               SubCategory = "StoreCreated"
	StoreParticipationEnabled  SubCategory = "StoreParticipationEnabled"
	StoreParticipationDisabled SubCategory = "StoreParticipationDisabled"
	StoreParticipationToggled  SubCategory = "StoreParticipationToggled"
	StoreRebranded             SubCategory = "StoreRebranded"
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

	// Orders
	CreateOrder   SubCategory = "CreateOrder"
	CancelOrder   SubCategory = "CancelOrder"
	ReadyOrder    SubCategory = "ReadyOrder"
	CompleteOrder SubCategory = "CompleteOrder"
	GetOrder      SubCategory = "GetOrder"
	// orders events
	OrderHandleEvent SubCategory = "OrderHandleEvent"
	OnOrderCreated   SubCategory = "OnOrderCreated"
	OnOrderReadied   SubCategory = "OnOrderReadied"
	OnOrderCanceled  SubCategory = "OnOrderCanceled"

	// Depot
	CreateShoppingList   SubCategory = "CreateShoppingList"
	CancelShoppingList   SubCategory = "CancelShoppingList"
	AssignShoppingList   SubCategory = "AssignShoppingList"
	CompleteShoppingList SubCategory = "CompleteShoppingList"
	GetShoppingList      SubCategory = "GetShoppingList"
	// depot events
	DepotHandleEvent        SubCategory = "DepotHandleEvent"
	OnShoppingListCreated   SubCategory = "OnShoppingListCreated"
	OnShoppingListCanceled  SubCategory = "OnShoppingListCanceled"
	OnShoppingListAssigned  SubCategory = "OnShoppingListAssigned"
	OnShoppingListCompleted SubCategory = "OnShoppingListCompleted"

	// Search
	SearchOrders SubCategory = "SearchOrders"
	HandleEvent  SubCategory = "HandleEvent"

	// Stream
	StreamError SubCategory = "StreamError"

	HandleReply SubCategory = "HandleReply"
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
