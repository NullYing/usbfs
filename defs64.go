// +build linux,amd64 linux,arm64,darwin

package usbfs

const (
	usbdevfsControl = 0xc0185500
	usbdevfsBulk    = 0xc0185502
)
