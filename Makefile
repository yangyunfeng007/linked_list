test:
	go test -run TestIntSet

test_race:
	go test -run TestIntSet -race

bench:
	go test -run=NOTEST -bench=. -benchtime=100000x -benchmem -count=10 -timeout=60m  > x.txt && benchstat x.txt