# Alien-Invasion

A simulation of mad aliens invading our earth.

## Build, Test and Run

```bash
go build
go test -v ./...
./alien-invasion -f <map_file | optional> -n <number_of_aliens | optional> -max <max _allowed_steps | optional>
```

## Assumptions

- In each step, all the aliens are moving to their closest neighboring city. So, a city might have more than 2 aliens at any given step.
- Aliens don't move to their neighboring city randomly, instead, they prefer the closest one (first available in the direction map).
- The final result will print all the non-deleted cities with at least one direction map, just like the input file.
