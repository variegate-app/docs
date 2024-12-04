# Snippet

## [<<< ---](gobyexample.md)

```go
func(interval int) {
		ticker := time.NewTicker(time.Second * time.Duration(interval))
		for range ticker.C {
				log.WithFields(log.Fields{
					"level": "debug",
					"interval": interval,
				}).Error("Some usefull info")
		}
}(10)
```

ticker

```go
%t: для вывода значений типа boolean (true или false)

%b: для вывода целых чисел в двоичной системе

%c: для вывода символов, представленных числовым кодом

%d: для вывода целых чисел в десятичной системе

%o: для вывода целых чисел в восьмеричной системе

%q: для вывода символов в одинарных кавычках

%x: для вывода целых чисел в шестнадцатиричной системе, буквенные символы числа имеют нижний регистр a-f

%X: для вывода целых чисел в шестнадцатиричной системе, буквенные символы числа имеют верхний регистр A-F

%U: для вывода символов в формате кодов Unicode, например, U+1234

%e: для вывода чисел с плавающей точкой в экспоненциальном представлении, например, -1.234456e+78

%E: для вывода чисел с плавающей точкой в экспоненциальном представлении, например, -1.234456E+78

%f: для вывода чисел с плавающей точкой, например, 123.456

%F: то же самое, что и %f

%g: для длинных чисел с плаващей точкой используется %e, для других - %f

%G: для длинных чисел с плаващей точкой используется %E, для других - %F

%s: для вывода строки

%p: для вывода значения указателя - адреса в шестнадцатеричном представлении

Для чисел с плавающей точкой можно указать точность или количество символов в дробной части. Для этого количество символов указывается после точки: %.2f - две цифры в дробной части после точки. Например, варианты форматирования чисел с плавающей точкой:

%f: точность и ширина значения по умолчанию

%9f: ширина - 9 символов и точность по умолчанию

%.2f: ширина по умолчанию и точность - 2 символа

%9.2f: ширина - 9 и точность - 2

%9.f: ширина - 9 и точность - 0

fmt.Printf("%06d", 12) // Prints to stdout '000012'
fmt.Printf("%d", 12)   // Uses default width, prints '12'
fmt.Printf("%6d", 12)  // prints '    12'
fmt.Printf("%-6d", 12)   // Padding right-justified, prints '12    '

myWidth := 6
fmt.Printf("%0*d", myWidth, 12) // Prints '000012' as before
```


# Сборка

Go-сервисы собираются в RPM-пакеты на серверах CI и публикуются в [http://repo.dmr/](http://repo.dmr/). Для сборки на CI в корне репозитория должен лежать файл .gitlab-ci.yml ([пример](https://gitlab.dev.dmr/dmrdev/sbp/-/blob/master/.gitlab-ci.yml)) как минимум с двумя этапами build и public, соотвественно сборка RPM и загрузка его в repo.dmr.

Сервисы работают в проде под systemd, для этого в пакете должен быть service файл ([пример](https://gitlab.dev.dmr/dmrdev/sbp/-/blob/master/configs/sbp.service))

```yaml
include:
  - project: ci-cd/build
    file: pipeline/common.yaml

image: registry.dev.dmr/images/buildgolang:el7-go1.21-latest

stages:
  - lint
  - test
  - build
  - publish

lint:
  stage: lint
  image: golangci/golangci-lint:v1.56
  allow_failure: false
  except:
    - tags
  script:
    - make lint

test:
  stage: test
  when: on_success
  allow_failure: false
  except:
    - tags
  script:
    - useradd -m golang
    - su golang -c 'make test'

build:
  stage: build
  when: on_success
  allow_failure: false
  before_script:
    - yum install -y rpm-build yum-utils git
  script:
    - make build
    - VERSION=$(make version)
    - yum-builddep -y build/smev3-api.spec
    - tar --exclude='build/*.tar.gz' --transform "s/^./smev3-api-$VERSION/" -czf build/smev3-api-$VERSION.tar.gz .
    - rpmbuild -bb --define "_sourcedir $PWD/build" --define "__version $VERSION" --define "__release $CI_PIPELINE_IID" build/smev3-api.spec
    - "mv $HOME/rpmbuild/RPMS/x86_64/*rpm ./"
  artifacts:
    expire_in: 1 day
    paths:
      - "*.rpm"

publish:
  stage: publish
  when: manual
  script:
    - ls -1 *.rpm
    - rpm_file=$(ls -1 *.rpm)
    - for file in $rpm_file; do /usr/bin/curl -X POST -F "myfile=@${file}" -F "file_root=/7/dmr-dev/x86_64/" -i "${RPM_REPO_URL}/upload"; done
    - /usr/bin/curl -i ${RPM_REPO_URL}/update

```

```makefile
PROJECT_NAME=api
VERSION=1.8
CI_PIPELINE_IID?=1
LDFLAGS="-X 'main.version=$(VERSION)-$(CI_PIPELINE_IID)-$(CI_COMMIT_REF_NAME)-$(shell git rev-parse --short HEAD)-$(shell date +"%Y-%m-%d")-$(shell go version | cut -d ' ' -f3)'"

.PHONY: all build test vendor version

version:
	@echo $(VERSION)

all: build test

up:
	@docker-compose up api

up-consumer:
	@docker-compose up consumer

migrate-up:
	@docker-compose up migrate_up

migrate-down:
	@docker-compose up migrate_down

build-docs:
	docker run --name $(PROJECT_NAME)-docs --rm -i -v "$(PWD):/smev3-api" -w /sbp onrik/gaws:1.6.0 sh -c "gaws -t 'SMEV3 api' -v '$(VERSION)' -s http://gotest.dev.dmr/smev3 -path=/smev3-api/pkg/api/handlers > /smev3-api/docs/openapi.yml"

build:
	go build -mod=vendor -ldflags $(LDFLAGS) ./cmd/smev3-api
	go build -mod=vendor -ldflags $(LDFLAGS) ./cmd/smev3-consumer

lint:
	golangci-lint run ./pkg/... -E gofmt -E bodyclose -E gosec -E goconst -E unparam -E unconvert -E asciicheck -E exportloopref -E nilerr --timeout=10m

docker-lint:
	docker run --name $(PROJECT_NAME)-lint --rm -i -v "$(PWD):/smev3-api" -w /smev3-api golangci/golangci-lint:v1.56 make lint

test:
	go test ./pkg/... -coverprofile=coverage.out && go tool cover -func=coverage.out

docker-test:
	docker run --name $(PROJECT_NAME)-test --rm --platform=linux/amd64 -i -v "$(PWD):/smev3-api" -w /smev3-api golang:1.21 bash -c "useradd -m golang && su golang -c 'make test'"

vendor:
	export GOPRIVATE=gitlab.dev.dmr GOINSECURE=gitlab.dev.dmr && go mod tidy && go mod vendor

```

```yaml
version: "3"

services:
  api:
    image: images/buildgolang:el7-go1.21-latest
    entrypoint: bash -c "go install -buildvcs=false ./cmd/api && /go/bin/api -c /api/configs/docker/config.yml"
    working_dir: /api

  consumer:
    image: images/buildgolang:el7-go1.21-latest
    entrypoint: bash -c "go install -buildvcs=false ./cmd/consumer && /go/bin/consumer -c /api/configs/docker/config.yml"
    working_dir: /api

  migrate_up:
    image: onrik/goose:3.6.1
    entrypoint: goose -dir=/migrations postgres "postgres://postgres:password@postgres/api?sslmode=disable" up
    volumes: 
      - ./migrations:/migrations

  migrate_down:
    image: onrik/goose:3.6.1
    entrypoint: goose -dir=/migrations postgres "postgres://postgres:password@postgres/api?sslmode=disable" down
    volumes: 
      - ./migrations:/migrations

networks:
  default:
    name: dev_network
    external: true

```