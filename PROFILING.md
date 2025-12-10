# Profiling Go Code

## Quick Start

### Profile Benchmarks
```bash
# Generate CPU and memory profiles
go test -bench=BenchmarkPuzzle2 -cpuprofile=cpu.prof -memprofile=mem.prof ./cmd/day09

# View CPU profile (interactive)
go tool pprof cpu.prof

# View memory profile
go tool pprof mem.prof

# Generate web UI (requires graphviz)
go tool pprof -http=:8080 cpu.prof
```

### Profile Main Program
```bash
# Run with CPU profiling
go run ./cmd/day09 -cpuprofile=cpu.prof

# Run with memory profiling  
go run ./cmd/day09 -memprofile=mem.prof
```

## Common pprof Commands (in interactive mode)

- `top` - Show top functions by time/allocations
- `top10` - Show top 10 functions
- `list <function>` - Show source code with annotations
- `web` - Generate SVG graph (requires graphviz)
- `svg` - Generate SVG output
- `png` - Generate PNG output
- `help` - Show all commands
- `exit` or `quit` - Exit pprof

## Useful pprof Flags

- `-http=:8080` - Start web UI on port 8080
- `-top` - Show top functions and exit (non-interactive)
- `-list=<regex>` - List functions matching regex
- `-focus=<regex>` - Focus on functions matching regex
- `-ignore=<regex>` - Ignore functions matching regex

## Example Workflow

1. Generate profile:
   ```bash
   go test -bench=BenchmarkPuzzle2 -cpuprofile=cpu.prof ./cmd/day09
   ```

2. View in web UI:
   ```bash
   go tool pprof -http=:8080 cpu.prof
   ```
   Then open http://localhost:8080 in your browser

3. Or view top functions:
   ```bash
   go tool pprof -top cpu.prof
   ```


