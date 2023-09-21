package ports

import "github.com/KaiBennington2/RabbitGo/sdk/types"

type ISetting interface {
	Load() (*types.Setting, error)
}
