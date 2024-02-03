Jetson Exporter
===============

<div align="flex">
    <img align="right" src="./assets/elroy_jetson.webp" width="40%" alt="Elroy Jetson">
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/GDWR/jetson-exporter)](https://goreportcard.com/report/github.com/GDWR/jetson-exporter)
[![Go Reference](https://pkg.go.dev/badge/github.com/GDWR/jetson-exporter.svg)](https://pkg.go.dev/github.com/GDWR/jetson-exporter)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/GDWR/jetson-exporter)](https://github.com/GDWR/jetson_exporter/releases)
[![GitHub](https://img.shields.io/github/license/GDWR/jetson-exporter)](https://github.com/GDWR/jetson_exporter/blob/main/LICENSE)

> [!CAUTION]
> Project is currently under creation, not ready for use.

Prometheus exporter for the Nvidia Jetson family.

Supported devices:
- Jetson Orin

## Install

```shell
# Download the latest release
wget https://github.com/GDWR/jetson-exporter/releases/download/v0.0.1/jetson-exporter_0.0.1_arm64.deb
# Install the package
sudo dpkg -i jetson-exporter_0.0.1_arm64.deb
# Clean up
rm jetson-exporter_0.0.1_arm64.deb
```


Resources
---------
* [Repository Structure](https://github.com/golang-standards/project-layout)
* [`tegrastats` Utility](https://docs.nvidia.com/drive/drive-os-5.2.0.0L/drive-os/index.html#page/DRIVE_OS_Linux_SDK_Development_Guide/Utilities/util_tegrastats.html)
* [ridgerun's Evaluating Performance docs using `tegrastats`](https://developer.ridgerun.com/wiki/index.php/Xavier/JetPack_5.0.2/Performance_Tuning/Evaluating_Performance)