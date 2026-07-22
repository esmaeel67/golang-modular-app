package depotpb

import (
	"github.com/esmaeel67/golang-modular-app/internal/registry"
	"github.com/esmaeel67/golang-modular-app/internal/registry/serdes"
)

const (
	CommandChannel = "mallbots.depot.commands"

	CreateShoppingListCommand = "depotapi.CreateShoppingListCommand"
	CancelShoppingListCommand = "depotapi.CancelShoppingListCommand"
	InitiateShoppingCommand   = "depotapi.InitiateShoppingCommand"

	CreatedShoppingListReply = "depotapi.CreatedShoppingListReply"
)

func Registrations(reg registry.Registry) (err error) {
	serde := serdes.NewProtoSerde(reg)

	if err = serde.Register(&CreateShoppingList{}); err != nil {
		return err
	}

}

// commands
func (*CreateShoppingList) Key() string {
	return CreateShoppingListCommand
}
func (*CancelShoppingList) Key() string {
	return CancelShoppingListCommand
}
func (*InitiateShopping) Key() string {
	return InitiateShoppingCommand
}

// Replies
func (*CreatedShoppingList) Key() string {
	return CreatedShoppingListReply
}
