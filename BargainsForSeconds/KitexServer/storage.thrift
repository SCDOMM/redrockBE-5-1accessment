namespace go storage.service

struct OrderData {
	1: required i32 UserId,
	2: required i32 ProductId 
}

service StorageService {
	void CheckOrder(1: required OrderData orderData)
}