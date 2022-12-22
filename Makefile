challenges := $(wildcard */day*/main.go)
benchmarkFiles := $(patsubst %/main.go,benchmarks/results/%-ns,$(challenges)) $(patsubst %/main.go,benchmarks/results/%-mem-b,$(challenges))

benchmarks/README.md: benchmarks/README.template.md benchmarks/benchmarks.md
	cat $^ > $@

benchmarks/benchmarks.md: $(benchmarkFiles) benchmarks/main.go
	cd benchmarks && go run .

benchmarks/results/%-ns benchmarks/results/%-mem-b: %/main.go %/main_test.go
	@mkdir -p $(@D)
	$(eval resultPrefix := ../../benchmarks/results/$*)
	cd $* && go test -bench=BenchmarkMain -benchmem . | grep "BenchmarkMain-" | awk '{print $$3>"$(resultPrefix)-ns"}{print $$5>"$(resultPrefix)-mem-b"}' || rm $(resultPrefix)-ns $(resultPrefix)-mem-b


benchmarks/benchmarks.snippet.md: $(benchmarkFiles) | benchmarks
	jq --raw-input -r '(input_filename|split("/")|first|sub("day";""))+" | "+.' $^ > $@
