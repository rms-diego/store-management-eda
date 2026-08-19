.PHONY: run-watch-order run-watch-inventory run-watch-billing

clean-tmp:
	rm -rf ./tmp/order \
	rm -rf ./tmp/inventory \
	rm -rf ./tmp/billing 


run-watch-billing:
	PORT=8082 air -c .air.billing.toml

run-watch-inventory:
	PORT=8081 air -c .air.inventory.toml

run-watch-order:
	PORT=8080 air -c .air.order.toml
