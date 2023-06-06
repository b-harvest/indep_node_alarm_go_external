.PHONY: install

install:
	go build -o indep_ext ./cmd/indep_ext.go
	mv indep_ext ${GOPATH}/bin
