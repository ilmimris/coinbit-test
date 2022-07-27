.PHONY: init-topic

init-topic:
	docker exec -it broker kafka-topics --bootstrap-server localhost:9092 --create --topic deposits --partitions 1 --replication-factor 1 || true
	docker exec -it broker kafka-topics --bootstrap-server localhost:9092 --create --topic aboveThreshold-table --partitions 1 --replication-factor 1 || true
	