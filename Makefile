.PHONY: clean
dist/jetsonexporter_0.0.1_arm64.deb:
	mkdir -p dist/jetsonexporter/DEBIAN
	mkdir -p dist/jetsonexporter/usr/bin
	mkdir -p dist/jetsonexporter/lib/systemd/system

	cp deploy/debian/* dist/jetsonexporter/DEBIAN/
	cp deploy/jetsonexporter.service dist/jetsonexporter/lib/systemd/system
	GOARCH=arm64 go build -o dist/jetsonexporter/usr/bin/jetsonexporter cmd/jetsonexporter.go

	dpkg-deb --build dist/jetsonexporter dist/jetsonexporter_0.0.1_arm64.deb

clean:
	rm -rf dist
