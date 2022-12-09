#!/bin/bash -e

go test -bench=BenchmarkMain -cpuprofile=cpu.prof .
go tool pprof -http localhost:8081 ./day*.test cpu.prof
