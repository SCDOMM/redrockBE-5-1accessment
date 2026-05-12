namespace go checkserver.service

struct OrderData {
	1: required i32 UserId,
	2: required i32 ProductId 
}

service CheckService {
	void CheckOrder(1: required OrderData orderData)
}