all: bench

bench:
	go test ./... -bench=. -test.benchmem | tee bench.txt
