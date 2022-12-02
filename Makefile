challenges := $(wildcard */day*/main.go)
benchmarkFiles := $(patsubst %/main.go,benchmarks/results/%-ns,$(challenges)) $(patsubst %/main.go,benchmarks/results/%-mem-b,$(challenges))

benchmarks/README.md: benchmarks/README.template.md benchmarks/benchmarks.md
	cat $^ > $@

benchmarks/benchmarks.md: $(benchmarkFiles) benchmarks/main.go
	cd benchmarks && go run .

benchmarks/results/%-ns: %/main.go %/main_test.go
	@mkdir -p $(@D)
	cd $* && go test -bench=BenchmarkMain . | grep "BenchmarkMain-" | awk '{print $$3}' > ../../$@ || rm ../../$@

benchmarks/results/%-mem-b: %/main.go %/main_test.go
	@mkdir -p $(@D)
	cd $* && go test -bench=BenchmarkMain -benchmem . | grep "BenchmarkMain-" | awk '{print $$5>"../../$@"}' > ../../$@ || rm ../../$@

benchmarks/benchmarks.snippet.md: $(benchmarkFiles) | benchmarks
	jq --raw-input -r '(input_filename|split("/")|first|sub("day";""))+" | "+.' $^ > $@
