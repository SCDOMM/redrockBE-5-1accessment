package main

import (
	"Storage/dao"
	service "Storage/kitex_gen/storage/service"
	"context"
	"log"
)

// StorageServiceImpl implements the last service interface defined in the IDL.
type StorageServiceImpl struct{}

// CheckOrder implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) CheckOrder(ctx context.Context, orderData *service.OrderData) (err error) {
	err = dao.CheckHandler(orderData)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
