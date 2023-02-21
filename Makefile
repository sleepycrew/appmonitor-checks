SOURCES := $(wildcard ./*.go)

.Phony: clean

build/systemd.so: $(SOURCES)
	go build -buildmode=plugin -o build/systemd.so ./pkg/systemd/plugin

clean:
	rm -rf build
