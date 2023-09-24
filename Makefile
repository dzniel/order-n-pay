build-order:
	go build -a -o order ./services/order/main.go

build-payment:
	go build -a -o payment ./services/payment/main.go

order:
	./order

payment:
	./payment

.PHONY: order payment
