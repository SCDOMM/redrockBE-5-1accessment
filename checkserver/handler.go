package checkserver

import (
	service "checkserver/kitex_gen/checkserver/service"
	"context"
)

// CheckServiceImpl implements the last service interface defined in the IDL.
type CheckServiceImpl struct{}

// CheckOrder implements the CheckServiceImpl interface.
func (s *CheckServiceImpl) CheckOrder(ctx context.Context, orderData *service.OrderData) (err error) {
	// TODO: Your code here...
	return
}
