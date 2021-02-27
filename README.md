[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/kumadee/blackjack)
[![Coverage Status](https://coveralls.io/repos/github/kumadee/blackjack/badge.svg?branch=main&service=github)](https://coveralls.io/github/kumadee/blackjack?branch=main)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/kumadee/blackjack/Go)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=kumadee_blackjack&metric=alert_status)](https://sonarcloud.io/dashboard?id=kumadee_blackjack)

# blackjack

Console based blackjack game.

To run the CPU and memory profiling, run the tests with the `cpuprofile` and `memprofile` flags.
```bash
# Assuming we are already in the directory with go.mod
go test -benchmem -run=^$ blackjack -bench "^(BenchmarkStartGame)$" -benchmem -cpuprofile cpu.out -memprofile mem.out
```

To view the profile data in browser, run the below command.
```bash
go tool pprof -http=:8080 cpu.out
go tool pprof -http=:8080 mem.out
```

# gitpod known issues
- Run `gp env PIP_USER=false` so that `pre-commit` does not fail with error `Can not perform a '--user' install`.
Related to https://github.com/gitpod-io/gitpod/issues/1997
