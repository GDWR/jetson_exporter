.PHONY: clean
dist/jetson-exporter_0.0.1_arm64.deb:
	mkdir -p dist/jetson-exporter/DEBIAN
	mkdir -p dist/jetson-exporter/usr/bin
	mkdir -p dist/jetson-exporter/lib/systemd/system

	cp deploy/debian/* dist/jetson-exporter/DEBIAN/
	cp deploy/jetson-exporter.service dist/jetson-exporter/lib/systemd/system
	GOARCH=arm64 go build -o dist/jetson-exporter/usr/bin/jetson-exporter ./cmd/jetson_exporter.go

	dpkg-deb --build dist/jetson-exporter dist/jetson-exporter_0.0.1_arm64.deb

clean:
	rm -rf dist
