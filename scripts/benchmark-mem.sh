#!/bin/bash -e

go test -bench=BenchmarkMain -memprofile=mem.prof .
go tool pprof -http localhost:8081 ./day*.test mem.prof
