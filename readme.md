# TTO Resize

## Requirement

- Golang 1.11.2
- libvips 8.6.5

## Usage options

- Scale with profile: `/<project>/i/<profile>/<year>/<month>/<day>/<image>`
- Origin: `/<project>/r/<year>/<month>/<day>/<image>`

## Install

- Install libvips before build:

```bash
    ./preinstall.sh
```

Or with Centos 7 `yum install libvips-devel-8.6.5`

- Install dependencies Go

```bash
    dep ensure -update
```

## Build

```bash
    make build
```

## Testing

```bash
    make test
```

## Deploy

```bash
    make build
    make deploy
```

## Benchmark

```bash
    ./benchmark.sh
```

Or

```bash
    make benchmark
```

## Docker Usage

- Build

```bash
    docker build -t youname/resize .
    docker-compose -f docker-compose-build.yml up
```

- Run

```bash
    docker-compose up -d
```
