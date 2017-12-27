// Code generated by go-bindata.
// sources:
// compiled/Authority.abi
// compiled/Authority.bin
// compiled/Energy.abi
// compiled/Energy.bin
// DO NOT EDIT!

package gen

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _compiledAuthorityAbi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x94\x41\x6a\xf3\x40\x0c\x85\xef\xa2\xf5\xac\x7e\xf8\xbb\xf0\xae\x07\x28\x85\x6e\x43\x28\xe3\x8c\x92\x0a\x1c\xc9\x8c\x34\x49\x4d\xc8\xdd\x4b\x8a\xc7\xe3\x34\x6e\x5a\x0a\x4e\xe9\xd6\x7e\xd2\xfb\x2c\x3d\x6b\x71\x80\x95\xb0\x9a\x67\x83\x6a\xed\x1b\x45\x07\xc4\x6d\x32\x85\x6a\x71\x00\xf6\x5b\x84\x0a\x9e\x19\xf7\x8f\x7b\xc6\x08\x0e\xac\x6b\x4f\x8f\x7c\x08\x11\x55\xe1\xb8\x74\x59\xa5\x68\x59\x24\xc9\xfa\x1e\x4b\x07\xad\xef\x7c\xdd\xe0\xd0\x5f\xcd\x1b\x3e\x24\xf3\x35\x35\x64\x1d\x54\xc0\xc2\x59\x34\x18\xac\x13\xaf\x8c\x84\xe1\xe8\xc6\x8c\x16\xd3\x24\x62\x29\x4c\xc4\xf6\xef\xff\xdd\x98\xcc\xd7\x8a\x6c\x88\x67\x64\x13\xa5\xe3\x8f\xfa\x9a\x7a\x47\xb8\x9f\x87\xb7\x8d\xd2\x8a\x62\xd4\x5f\x00\xfe\x3c\x04\x99\x6a\xc2\xde\x15\x55\x5d\x5e\xd7\x22\xcd\xd9\x1a\x92\xbd\x48\x24\xc5\x5b\x26\xa4\xd8\xcb\x45\x38\xff\xfe\x44\xdf\x83\x3d\xeb\x38\x3f\xd2\x17\x73\x62\xba\xad\x75\x19\x49\x40\x5d\x95\x72\xb5\x48\xbc\x19\xcf\x25\xe2\x86\xd4\x66\x39\x45\x9e\x85\xbb\xad\x24\x9d\xe2\x23\x0e\xf8\x8a\x21\x87\x30\xe3\x9e\xb6\x7a\xf5\x74\x3e\xf5\xbc\x18\x8a\x0c\x77\xa7\xd5\xfe\xd4\xf2\x7a\xb4\x86\xa2\xbe\xdf\x37\xa2\x76\x9f\x7f\xde\x4b\xc4\xe5\x5b\x00\x00\x00\xff\xff\x53\x5c\xb3\x30\x46\x06\x00\x00")

func compiledAuthorityAbiBytes() ([]byte, error) {
	return bindataRead(
		_compiledAuthorityAbi,
		"compiled/Authority.abi",
	)
}

func compiledAuthorityAbi() (*asset, error) {
	bytes, err := compiledAuthorityAbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Authority.abi", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledAuthorityBin = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x0b\x72\xdc\x30\x08\xbd\x12\x1f\x21\xc1\x71\xf4\xbd\xff\x11\x3a\x48\xf2\x7a\xe3\x24\x6d\xb6\x4d\xb3\x9f\x99\x8c\xc3\xca\x06\x3d\xe0\xf1\xac\x08\xfe\x0d\x20\xc4\x1c\x01\x40\x21\x22\x20\x00\x64\x45\x09\x8a\x89\xc7\x17\x3f\x40\x68\x18\x0d\x94\xbf\x7e\x0f\x46\x20\x4c\x06\x22\xe2\x7e\xab\x2e\xff\x20\x33\x16\x36\xff\x3b\x18\xe0\x8c\x32\x02\x04\x8e\x38\x57\x69\x97\x34\xd7\x49\xaa\x33\xe6\xbf\xfa\x18\x40\x88\x7c\x46\xa4\x10\x19\x39\x8f\x00\x2c\x18\xdc\x8f\xb1\x24\xb7\x52\xd1\xc4\xa2\x61\x59\x6b\x5d\xd6\x48\xdc\xb9\x61\x9e\x56\xa4\xb1\xad\xa3\x65\x30\xe3\x65\x35\x5a\x56\x6d\x59\x6a\x91\xb2\xac\x2d\x2e\x6b\x23\x0e\x21\x86\x31\xad\x44\x65\x59\x3b\xd6\xc4\x46\xeb\xb9\x14\xf7\x73\x07\x55\x32\x2d\x7d\x59\x35\x48\x92\xb2\xf2\x36\x9a\x14\x0e\x28\x33\xe2\x8d\xcc\xb2\xce\x68\xb3\x23\xe7\xf8\xb2\x3c\x92\x1f\x83\x08\x04\x80\x06\x86\x06\x02\x33\x4f\xd4\x51\xa2\x14\x80\xd3\x63\x4b\x57\x8f\xbd\x9d\x1e\x3f\x7c\x0a\xa7\xe1\x4f\x99\x79\x45\x05\xa5\x47\xe2\x7a\x64\xad\xa2\xd7\xcd\xf4\x8f\xd3\xf7\xf6\x68\x08\x6c\x30\xf8\xb6\x0b\xe4\x7c\xd9\x05\xfa\xea\xdf\xef\xa2\xf4\x57\xdb\x85\xb5\xeb\x2e\x5a\xf8\x8e\xec\xfb\xdd\x28\xf8\x09\x0e\xa3\x5d\x6a\x02\x3b\x5e\xe3\xe8\xe6\xad\xdb\xf0\xc5\x10\x23\x8e\x97\x48\x29\xb6\xff\x8e\x58\x1c\xf1\x82\x18\xa5\x6b\xfd\x91\x52\x44\x48\x02\xd7\x95\x3a\xae\x2b\xdb\xf8\xa8\x52\x95\x00\x2f\x16\xe7\x58\xc0\x01\x4e\xbe\x30\xd1\x08\xb0\x71\x5a\x08\x19\x28\xee\x6b\x52\x30\x36\x9a\x7b\xc1\x13\x4f\x65\x65\xcf\x9c\x06\x4e\xee\x61\xe1\x7b\x7c\xef\xf6\x98\x8c\x8f\xc8\x57\xb4\x06\x12\xec\x36\x63\x9c\x7b\xff\x57\xf6\xf9\xa1\x39\x84\xc1\xf3\xe4\x75\xcc\xf5\x1e\x59\xc7\xea\x99\x53\x71\x65\x7e\x66\x16\x97\x6f\x27\xa5\x15\xa9\xcf\xbf\xd1\xa5\xd8\x8c\x71\xe5\xc6\xaf\xc8\x67\x1a\xfa\xd5\xcc\xcb\x3f\xe0\xad\xbb\x53\x81\xde\x7b\xaf\xed\xc7\xbc\xc3\xca\xc2\xcb\xd7\x4e\x90\xb7\xfd\x0b\xb0\x34\xd5\x7f\x8b\xf3\x8e\xe1\xe0\xbc\x5e\x59\x70\xac\x04\x14\x67\x64\x12\x69\x47\x86\x67\xb4\x05\x2f\xd1\x92\x82\x84\x89\x35\x02\x2a\x29\xfa\xaa\x2a\xde\xcf\x11\xa1\x48\xf2\x7c\x18\x7e\x96\x6f\x95\xc5\x6f\x86\xcf\xeb\x96\x63\x17\xb6\xae\x17\xfe\x0f\x31\xf8\x77\xe1\xaf\x78\xd3\xb5\x31\xe5\x5b\x25\x23\xe2\xc6\x5f\xd2\x3b\xfc\xe3\xba\xfb\xb6\x03\xe0\x4b\xd7\x89\xea\x1f\xbb\xee\x9f\xba\xe4\x88\x42\xf9\xbd\xef\x1a\xff\xe4\xfb\x69\xef\x0e\x07\x6e\xa8\x78\x62\xe7\xfc\x33\x6b\x38\x52\x3a\x6a\x58\xe7\x3c\x9a\x77\xbc\x52\x7d\xa8\xf2\xc6\x4f\x6d\xac\x08\xa5\x3c\xa8\x8c\x46\x02\xc9\x06\x59\x62\xab\x2c\xb1\x27\x19\x0d\xb0\xbb\xe4\x31\x23\x1d\x31\x85\xde\xdb\x60\x8d\x8d\x47\x1f\xa1\x49\x43\x8d\xb1\xab\xd5\x3c\x12\x25\xf4\x18\x0e\x4d\x86\xb2\xbe\xbf\xd7\x50\x99\xf6\xcc\x8f\xdf\x33\xdf\x37\xe3\xbb\x36\xea\x9e\x2d\x30\xff\xff\x07\x98\x3b\xf1\xf5\xfd\x25\x05\x67\x40\x3a\x33\x72\xac\x0c\x6f\xf4\xc1\x6d\xf7\xae\xd2\xf4\x67\x63\xb6\x37\xfc\x71\x32\x8c\x9c\x1c\x93\x32\x5f\x38\xe6\xe8\x15\x78\xcc\xeb\xb7\x55\x3d\x48\x98\xfc\xe2\xff\x23\xc6\xf5\xa6\x30\x7f\xe2\xc5\x3e\x5e\x33\xde\xbf\xb7\x7d\x2a\xc8\x3b\x45\xf6\xc4\x1d\xb8\x82\x76\x8d\x7c\xfc\xee\x11\x8a\x2e\x8d\xec\xf9\xcf\xbb\x7f\x1f\x8c\x6e\x50\xe3\xc4\x21\x6b\x0f\x29\x57\x65\x8c\x5d\x00\x72\x25\xc6\x6a\x90\x63\x47\x0d\x5a\x33\x91\xca\x08\xd0\x20\xf5\x9c\x85\x00\xa4\x87\xc2\x19\xba\xad\xce\xfc\xa0\x3f\xdf\x68\xa8\x1b\xe7\x3d\x84\xc7\xf7\x69\x12\x5c\xef\x2f\x3a\x76\x5d\x3a\xab\xcc\xae\xa1\x1d\xa7\x3c\x8f\x91\xb7\x62\xa2\x19\xa1\x8d\x78\xaf\x98\xa6\x65\xf5\x5b\xf8\x40\x25\x99\xe8\x97\x55\x52\x7c\x09\x95\x14\xee\x54\x92\x3c\xaf\x22\x8e\x29\x38\x19\x34\xaf\x33\x84\x53\x27\x4d\xdb\xe2\xdb\xb0\xbb\x3e\x7c\xaa\x8d\x32\xfe\x51\x9f\xfc\xb3\x36\x0a\x9f\x68\xa3\x2c\xe1\xc5\xb5\x51\xf8\x44\x1b\xe5\x22\x9f\x68\xa3\x97\xa8\x8a\x59\x0d\x4b\x7b\x14\x01\x67\xd9\x7b\xbd\x91\x46\x4c\x71\xfc\xed\xc9\xaf\x7f\x52\xff\xfb\x93\xe3\xdb\xe9\x31\x8e\x7d\xf6\x23\x5e\x61\x51\xca\xca\xb1\xb2\x88\xa3\xee\x03\xd9\xf1\x4d\x5d\x92\xe2\x52\x7a\x6f\x2a\x04\xd7\x19\x0a\xac\x55\x6d\xe7\xa3\xd2\x9a\x24\xe5\xd4\x58\x1f\x3f\x37\xe7\xaf\x3c\x37\xdb\x6f\x9e\x4b\x5f\x9d\xcb\xef\xab\xfb\x3c\x4b\xb2\x59\x65\x80\xe6\xba\x09\x70\xcc\x53\xfa\x32\x40\x92\xce\xc9\x34\x86\x57\xb2\xb2\xba\x87\xa4\x22\xee\xa3\x62\x3f\x22\x98\xd9\x77\x4e\x15\x91\xcd\xc2\xfe\x6b\x5a\xfb\x98\x6b\x0e\x56\xa8\xd8\x24\x29\x09\x2a\x89\x18\x1e\x0a\xd5\x36\x2b\xaf\x7d\xae\x99\x32\xeb\x66\x9d\x46\x55\x2a\x17\x0c\x76\xbe\xdc\x22\xf3\x7e\xd7\x17\x4a\x87\x97\x70\x9c\xa4\x1e\xec\xb1\xfb\x69\x22\x5a\x59\xee\x9f\x61\xb3\x61\x9c\x84\x22\xa5\x9c\x72\x22\x06\x51\x82\x22\xa3\x0b\x12\xc9\xa8\xa5\x8e\x6e\xaa\x25\x41\x80\xc1\x6c\xd6\x5b\x43\x49\xd6\xb5\x66\xc8\xa1\x3b\x64\x6c\x01\x5a\x27\xcb\x15\x43\xe9\x02\xd9\xa1\xff\x15\x00\x00\xff\xff\x0d\xad\xec\xb4\xa6\x19\x00\x00")

func compiledAuthorityBinBytes() ([]byte, error) {
	return bindataRead(
		_compiledAuthorityBin,
		"compiled/Authority.bin",
	)
}

func compiledAuthorityBin() (*asset, error) {
	bytes, err := compiledAuthorityBinBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Authority.bin", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledEnergyAbi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\xd1\x6a\xdb\x30\x14\xfd\x17\x3d\xfb\xa9\x63\x63\xe4\xad\x0f\xe9\xd8\xc3\xba\x41\x52\x18\x94\x50\xae\xed\x9b\x4c\x4c\xbe\x12\xd2\x55\x52\x53\xfa\xef\xc3\xa9\x13\x3b\x43\xb2\x3d\x2f\x89\xfb\x94\x80\xae\xe4\x73\xae\xce\x91\x8e\x1e\x5f\x44\xa6\xc9\x31\x10\x8b\x19\x5b\x8f\x89\x90\x64\x3c\x3b\x31\x7b\x5c\x25\x82\xa0\x40\x31\x7b\xfb\x49\x84\xf6\x5c\x0f\xbd\x1c\x47\x44\x22\xb8\x34\xd5\x5f\xc7\x56\xd2\x46\xbc\xae\x12\x61\xa0\x84\x54\xa1\x98\xad\x41\x39\x4c\x84\x63\x60\xfc\xe6\x19\x52\xa9\x24\x97\x62\x26\x8c\xb7\xd8\x4c\x5d\x7b\xca\x58\x6a\x12\xaf\x49\x1b\x4f\x3d\xfb\x08\xe8\xf8\xd5\x27\x67\x90\x72\xb4\xcd\x0a\x90\xe7\x16\x9d\xdb\x2f\x70\x28\x82\x42\x7b\xe2\xa6\xc6\x4b\xe2\x9b\x8f\x9f\xf6\x08\xeb\x1a\x30\xc6\xea\x6d\x84\x9b\xf3\x59\x56\xad\x79\x5c\x20\xd5\x5a\x0d\xe4\x47\x9a\x0e\x45\x7d\x2c\xa3\x5d\x7f\xb8\xff\xba\x5c\xfe\xbc\x9b\xcf\xc3\xf0\xc2\xc4\xfa\xa1\x6d\x25\xee\xc6\x83\x62\xcd\xa0\x16\xde\x18\x55\x86\x61\xed\x0b\xe6\x84\x76\x53\x5e\x14\x61\x5c\x1c\x6b\xab\x8b\x6e\x65\xb0\xee\x1c\x67\xb0\x1b\xe4\x91\xe2\x6a\xd5\xe0\xb3\x91\x6d\x99\x07\x04\xe8\x7e\x81\xc5\xbb\x37\xc0\x67\xdc\xe3\x7f\x90\xdf\x05\xfb\x38\xc4\x81\x6c\x81\xdc\x1a\x6d\xbc\x07\xd3\xd8\x30\xb0\x05\x69\xc9\xe8\x3e\xdc\xb4\x19\xf6\xef\xad\x0b\x93\xea\xed\xed\x00\x7d\x75\xc9\xeb\xd2\xc7\x40\x8e\x99\x2c\x40\x45\xd8\xe5\xa7\xa0\x3e\x5f\xe5\x52\xf8\x7f\x53\x0f\x90\x6b\x85\xc3\xff\x75\x19\x5e\x4b\x89\x4f\x7a\x47\xc1\x3b\xaf\x81\x97\x82\x02\xca\xf0\xfb\x3a\xbc\x2f\xf5\xf0\x34\x92\x71\x65\x91\x6a\x15\xf1\xf8\x34\x31\xe2\x8c\xe7\xd7\xbb\x3c\xbb\x7a\x09\x46\x4d\xd1\x10\xdc\x20\x2f\xaa\x83\xec\xf6\xd0\x8b\x89\xa3\x48\xbf\x1d\x92\x21\x39\xb1\x95\x01\x95\xd2\xbb\xda\x16\x01\x6e\x16\x0b\x90\x54\x49\x72\x12\xd7\x54\x21\x70\x7e\xff\xe5\xe1\xc7\xd5\x3a\x3f\x52\x99\xa0\x54\x0a\xd9\xef\xfd\x12\x40\x9a\xca\x42\x7b\x17\x32\x9f\xa4\x1c\x9f\x31\x3f\x70\xee\x4f\x1b\x91\x09\x31\x6d\x1f\xcb\xeb\x4f\x1f\xea\xb7\xa0\x7c\x77\x18\x5b\x36\x5e\xae\x8b\x70\x8b\xc4\xa3\x29\x75\x08\x34\x32\xa3\xf3\x5d\x33\x9e\xd8\xed\xfe\x99\x03\xea\x5c\xc4\xce\xbf\x57\x27\xe5\x5d\x57\x75\xa4\x09\x5d\x71\x29\x32\x65\x40\x3a\x5f\xb4\xd2\xf9\x69\xe3\x56\x7f\x02\x00\x00\xff\xff\x2d\x66\x89\x32\x39\x0f\x00\x00")

func compiledEnergyAbiBytes() ([]byte, error) {
	return bindataRead(
		_compiledEnergyAbi,
		"compiled/Energy.abi",
	)
}

func compiledEnergyAbi() (*asset, error) {
	bytes, err := compiledEnergyAbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Energy.abi", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledEnergyBin = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\x09\x76\xf3\x2c\x8c\x57\xd2\x0a\xe2\x38\xac\xf7\x3f\xc2\x3c\xc0\x4e\x1c\x37\x4d\xe3\x7e\xe9\xef\xcc\x4c\x92\x57\xd7\x95\x41\x08\x21\x09\x49\xc8\x75\xd0\xbf\x02\x4a\x2c\xa8\x0e\x01\xa0\xa9\x77\x00\x60\xd0\x8a\x26\x87\x24\x0d\x0c\xfa\x13\xac\x1d\xce\xa1\x5f\x1b\x03\x5c\xfb\x3a\x00\x61\x87\xa3\x55\x49\xb3\x3f\xab\xcf\xd0\xf1\xfd\xea\x13\x00\xc4\x71\x5b\x3e\xe8\x0c\x1c\x83\x6b\xa5\x54\x60\x94\x3e\x4e\x4d\xea\x07\x34\x68\x8d\x3e\x4d\x28\xfa\xb0\x40\x5b\x85\xa6\xac\x13\x5a\x78\x42\xd1\xd0\x41\x29\x65\x42\x5b\x5e\xa0\xa9\x11\xfb\xec\x07\x94\x48\x27\x94\x38\x99\xa7\xa5\x2d\x65\x5a\xa0\x8e\x8b\x36\xc4\x01\x65\x5e\x68\x60\xe4\x5c\xd5\x4d\x0c\x9c\x16\x1a\x24\xc7\xa0\x21\xc2\x84\x56\x9b\x50\x0f\x11\x8c\x78\x62\x10\x59\xda\x06\x2d\x16\x92\x2c\xd0\xe0\x26\x34\x06\xd0\x90\x53\x1a\x50\x25\x99\xd0\xc4\xa5\x5a\x5d\x66\xac\xbe\x4e\x68\x29\x8e\x6a\xe1\x3a\xa1\x35\x4e\x68\x73\xa1\x56\x95\x49\x99\x53\xa7\x5e\xd3\xba\xd2\xd5\xdd\xac\xf4\xe5\x6e\x7d\xde\x6e\x9f\x77\x48\xed\x58\x7c\x53\xd7\xdb\x0b\x28\x5a\x97\x0d\x20\x00\x34\x32\x04\x36\x52\x32\x36\x34\xd4\xfe\x43\xf3\x59\x40\x05\x03\xc5\xb0\xb4\x0d\x60\x60\x6c\xdc\xb1\x6b\xea\xed\x11\xc6\x98\xc8\x63\x36\x46\x30\xfa\x9b\xc0\x82\xc3\xb0\xf7\xd2\x2e\x61\x48\xdc\xc7\x57\x98\xdf\x30\x7e\xe6\x73\x07\x38\x64\x65\xe2\x72\x69\xc1\xc5\x7d\x6c\xd7\x69\xe4\x31\x7e\x17\xd5\x2e\x99\x11\x18\x03\xba\x5b\x3a\x3b\xe6\x40\x13\xf7\x3a\xc7\x80\xc0\x01\x1a\x5f\x78\x83\x26\x3b\xde\x60\xea\x7a\x21\x9d\x1f\xac\xfe\x22\xb9\x3f\x7d\xd0\x5d\x79\x12\x06\x5f\x58\x6f\x21\x3a\x67\xed\x32\xdd\x70\x9d\x50\xe7\x77\x47\xfd\x03\x9a\x4b\xdd\xd3\x5c\x9d\x43\x88\x7c\xbb\x9e\xf4\x34\x46\x02\xbf\xc3\x48\xd0\x3a\x46\x91\xdf\x62\x64\xd8\x63\x8c\xf9\x55\x7c\x3d\xb3\xf7\x93\xab\x1c\xa5\xfe\x96\x73\xb9\xec\x38\xc7\x84\xef\xc1\xb9\xbb\x33\x2d\xee\x5f\xe5\x99\x65\x6f\x9f\xd8\x6d\x74\xb0\x3f\xe9\xda\xfd\x1c\x3d\x28\x31\xdd\xd0\x23\x47\xe6\x7c\xa4\xed\x75\x7e\xdd\xee\x5d\xee\xb7\x2b\xcd\xab\x6d\xfb\x7e\xee\x79\x6f\x7f\x38\x67\x87\xa8\x88\xb7\x5c\x75\xd0\xc7\x9c\xd7\xe7\x39\xdb\x78\x87\x5d\xc4\xbf\xaf\x2c\xa1\x62\xec\xb3\xee\x7b\xc9\x32\x03\xd1\x3d\x7f\xa4\xef\x53\xff\x3e\x83\x65\x44\x2b\xe9\x97\x7a\x2a\x11\xf7\x94\xc5\xe0\x10\x23\xd3\x9f\xed\xaa\x52\xc3\x8f\xbb\xaa\xe4\xfa\xd4\xae\xaa\xe8\xfe\x6c\x57\x55\xda\xf9\x9e\xa0\x4e\xfe\x72\x57\xc5\xe8\xf5\x1f\xad\x90\x5a\xd8\xd3\x5c\x5e\x46\xf3\x3f\xcb\x69\x2b\xf9\x97\x72\xaa\x4d\x77\xf3\x72\xf2\x12\x0d\x7a\xc5\xbc\x08\xf1\xb7\xf3\x72\x6e\xaf\x7f\xce\x85\x8e\x31\xf2\x51\x8c\xbd\xaf\xf9\x1e\x21\x91\xad\x7d\xa7\x3e\x76\xbd\x99\x91\xd1\xb0\x3a\xf1\x8a\xab\xfb\xce\xa2\x2a\x04\xa2\xae\x3b\xe0\xe4\xbc\x0f\x47\x22\xa3\x8e\x6b\xd5\xcf\x39\x2a\x0c\x6b\x01\xc0\xfc\xf7\x7b\x56\x80\xeb\x7d\x1f\x93\xfa\xc8\x3c\xef\x03\xa8\x04\x58\xad\xc0\x88\xe0\xc6\xbe\xd3\xf5\xc9\x21\x98\x4e\xff\x13\x27\xc5\x67\x53\xbb\xd2\x69\xa8\xd2\xc3\xc2\xd6\x80\xa6\xbf\x60\xdc\x29\x06\x42\x1f\x40\x55\xe1\x0d\x38\xdc\xad\x2c\x4e\x6a\x8e\xd1\xc0\x67\x73\x9b\xae\x94\x0b\x9d\x4c\x0b\x5e\x69\xd1\x64\xb8\xca\xe2\x79\xab\x6a\x87\x2c\xe1\xab\x46\x5e\x79\x60\x74\x48\x92\x8e\x71\xa9\x59\xd6\x54\xb1\x6a\x4d\x35\xfb\xa2\xa9\xa0\x34\x8f\x42\xbe\x60\x35\x69\x5c\x0a\x30\x4a\x86\xe6\x13\x51\xc0\xaa\x89\x00\x62\xb6\xec\x33\xa7\x40\x6a\xf2\xac\x2d\x8e\xdd\xe3\x59\xfd\x98\xa0\xc8\xbc\xf5\xd0\x4e\x95\xb8\x8f\xde\xbe\x82\x96\x8d\xde\x06\x9a\x12\x30\xd7\xd6\x89\xe1\xba\xca\x0e\xc9\xe3\xdd\x7d\x71\xb9\xda\x21\x8a\xe2\x91\xd6\x28\xd3\x6b\xec\x11\x7c\xd8\xfb\x17\x31\x44\x9b\x1e\xc6\xf0\x12\xae\x2d\xe3\x3e\x7e\x8d\xb1\x99\xdd\xb6\xec\xd0\x14\xb7\xed\xcc\x1f\x9a\x47\x38\xd4\x3a\x1d\xd3\x71\x57\x72\xe3\xd6\x9d\xcc\xd4\xc8\x47\x74\xad\x02\x5a\x08\x4e\xad\xc6\x4c\x2e\x81\x39\xad\x56\x7a\xbb\x58\x5d\x42\x97\xd4\x59\xc6\xc0\x22\xbe\xa5\x62\xc1\xae\x1e\xd7\x77\xd1\xf0\xdd\xa8\x21\x8a\x99\xd9\xb5\xef\x5f\x49\xa6\x7b\x94\xc3\x5e\xa8\x44\xe9\x14\x9f\x4d\xc1\x7d\x3e\x11\x04\x1e\x1e\xcc\xd4\xe1\x4b\x56\x64\xfd\xfd\x48\xf3\x3a\xc6\xe1\xd7\x8d\xf8\xcd\xc8\x86\x7f\xd7\xf7\x8d\x25\xae\xcc\x21\xab\x37\x1a\xab\x25\x33\xda\x9c\x70\x04\xf5\xad\x6a\x0a\xa3\xff\xc4\xba\xe0\x9c\x1e\x18\x4d\x7b\x6d\x8b\x8d\xbc\xef\x3d\xfe\x99\xb5\x39\xac\xdb\x7d\x4e\xd6\xd4\x9b\x9b\x5e\xc3\xd5\x1e\x99\x0e\xc8\xc6\x5a\x3a\x84\xa2\x23\xaa\x34\xb4\xc5\x7a\xcd\xb5\x71\x08\xa9\xb9\xf1\x64\xf0\x72\x7d\xda\xf9\xd7\x47\x88\xad\x47\x38\x3d\x9e\xe0\x11\xdd\x04\xfc\x9e\x7b\xc3\xe2\xad\xe7\x2f\xbb\x88\xc3\xea\x2f\xf3\x51\x69\x73\x1f\xd7\x38\x23\x8c\xc8\x6b\xf1\x82\x97\x0c\xc2\x75\xd5\x6e\xbd\xe8\xe7\xc7\xbd\xfa\xda\x47\x68\xdd\xfa\xe4\xb4\xa5\xa6\x8b\x9e\x8e\xf8\xeb\x0a\xeb\x77\xba\xe4\x33\xae\x79\x0d\x0d\xb2\xde\xeb\x97\x7d\xe2\xec\x98\xe4\x87\x08\xaa\xb6\xf6\x89\xa0\x7e\x49\xf3\xc7\x13\x7b\x05\x2d\x37\x11\x94\x43\x68\x60\x26\x77\x7c\x9b\x86\xfb\xfc\x6b\xc3\x62\x7a\xf5\xcb\xaf\xd6\xdf\xe0\x98\x47\xf3\x37\x32\x71\x2c\xdf\xbf\xca\xc4\x69\x94\x6f\xf7\x9b\x4e\x09\x23\x2e\xf9\xd9\x56\x48\xbd\x8e\x98\x69\x18\xc5\x34\xe1\x08\x1a\x27\xfc\x62\x3f\x4e\xe3\xfa\x39\xd2\xab\x72\xcb\x8f\x16\x3a\x3f\x96\x75\x3c\x25\x0e\x9f\x12\xa8\x32\xa8\xf1\x3d\xde\x1d\xb4\x1c\x8a\x53\x5e\x4f\x0b\x61\x95\xb9\xeb\x6f\xeb\x2f\x86\x76\x63\xd7\x70\x87\xc8\xb1\xa8\x77\x88\xa8\x34\x4f\x1b\xde\x80\x6a\x02\xba\x4f\xf5\xfb\xd8\x18\x87\x48\xe4\x06\xc7\xf0\x78\x44\xfa\x7f\x41\x03\x7f\x58\xa7\xff\x97\x76\x69\x2b\x1d\x2d\x5e\xf4\xe9\x44\x79\x7d\x6c\x05\x16\xfa\x0e\x8d\xf9\x47\x7b\xf6\xa1\xf5\x3a\x68\xe3\x5b\x29\x8d\x94\x62\xc1\x54\x29\x5b\x48\x2e\x64\x4a\xe0\xac\x65\xf6\x56\x62\x0c\x4a\x29\xfa\xee\x0c\x64\x89\x88\x8e\xac\xa9\x46\x29\x4d\x89\x13\xd7\x66\x7a\x30\x7f\x89\x43\x06\x38\xd1\xca\xe5\x19\x77\x0f\x58\xfa\xe4\x34\x5f\x40\xf9\xbb\x7a\xd2\x10\x78\xcd\xde\x2c\xeb\x3b\x5a\x2b\xd9\x25\x03\x20\xb0\x5a\x8b\xdb\x5c\x0b\x4a\x76\x3f\xe7\x5a\x56\x69\x9a\xd2\xf7\xba\xac\x4b\x80\x35\x87\xd1\xb1\xd8\xc2\xe3\x25\x4b\xc4\x97\xcc\x2c\x20\x3d\xc8\xcb\x9e\x1e\xc7\x3e\x8e\xba\xd1\x2d\xf9\xd7\xb3\xa5\xf7\x13\x75\xff\x97\x94\xbf\xab\xad\xe8\xfb\x81\xcb\x6c\x6e\x7f\x4e\x80\x2e\xef\xea\x3b\xd1\x15\x3f\xea\x13\x76\xd1\x39\xba\x7a\x53\xdf\x65\xce\xdc\x27\x93\xfe\x64\x26\x1d\xd1\x1b\xdb\xe0\xeb\x5a\x5d\xd3\xdb\x9b\x18\x5f\x6c\xb2\x0f\x75\xc4\x47\x3e\x98\xc9\x58\x29\xc2\x34\xab\xab\x3a\xb4\xce\x33\x2b\x44\x1f\x27\x9e\x35\x27\x82\x1d\x4f\x87\xa7\x62\x6c\xfa\x30\x16\xbb\x8c\x55\x61\x8e\x95\xeb\xed\x58\xf3\x3c\x14\x7d\x89\xe6\x8d\x3a\xbc\xf9\x3d\x0d\x93\x8a\x1a\x3f\xde\xc5\x0b\x28\x7f\x57\x8b\x71\xfb\xbd\xbb\x03\x9f\xe9\xcd\xab\xac\xe7\x4c\x27\xd2\x31\x3d\x18\x3c\x9d\x8e\x8b\xf7\xe4\x10\x43\x8d\x0e\x5c\xb7\x07\xa1\x64\xe3\x63\xd9\x02\xc6\x87\xb1\x36\x12\x37\xfd\xee\x19\x46\xa2\x7e\xc5\x71\x05\x36\x15\x7a\x84\x6d\x64\x80\xbf\xc7\xd7\xe7\xf7\xe0\x39\xa0\xf1\xb4\xb9\x06\xd7\xca\xe7\xa5\xd6\x6f\xa1\x87\xe3\x13\x55\x76\xbc\xad\xb2\x13\x95\x2a\xfe\xfb\x2a\xba\x9f\x3e\xdf\x57\xd9\x9d\xef\x01\xfe\xe0\xaf\x66\xa4\x8f\xbf\xfa\x4b\x9a\x3f\xbb\xcf\x2b\x68\xd9\xfb\xab\x19\xd3\xbd\x53\x22\xcc\xb4\x7b\x7f\x08\x33\xc3\xd6\x17\xd9\x9e\x12\x7d\x64\xe2\x30\xe5\xdf\x9e\x12\x61\xae\x7a\xf7\x94\xa8\x98\xbc\xd1\xa9\x48\x31\xfa\x5f\x74\x2a\x52\x4b\x53\x6f\x67\xeb\xde\x6a\xf7\xc9\x48\x05\x96\xbc\xd2\x26\xc7\x5c\x29\x7f\x72\xcc\x7f\x95\x63\x3e\x5a\x23\x7b\x52\x8e\xb9\x56\xf9\x92\x63\xae\xb5\x7c\xa2\xc0\x17\x50\xfe\xae\xfb\xf0\xae\x72\xf6\x4b\x04\xe8\xcd\x7f\x72\x40\x4f\xe4\x80\x64\xa9\x89\xdc\x56\x55\x5e\xaa\x29\xdd\x0b\xaa\x2a\x09\xa1\xde\xab\xaa\x24\x48\xe9\xb9\xaa\x4a\xa1\x4b\xfd\x9b\xac\x3b\x3e\x41\xa5\xcb\xce\xbe\x56\xcc\xc9\xba\x77\x8d\xa1\x7d\x6f\xd5\xea\xf5\xe9\xe3\x0c\xd4\x88\xd5\xee\xd4\x38\x12\xc4\x91\x5b\x32\x0e\x37\xd9\x87\xbd\xfc\x5d\x22\x94\xff\xe0\x4d\xdb\x6f\xde\xbe\x38\xe9\x14\x3c\x7c\xe5\x06\x5e\xeb\xd8\xed\x17\xb4\xcd\x1a\x55\x1a\xff\xe7\x62\x95\xad\xf1\x77\x9d\xd5\xa6\x3c\xb3\x81\x63\xed\xf1\xfa\x3e\xe6\x35\xc6\x9e\x36\x80\x4c\x96\xba\x5c\x5e\xe5\x6e\xe2\x6d\x36\x25\xcf\x96\x13\xa4\x2f\xb6\x84\x8c\x10\x2f\xed\x69\xad\xff\xb5\x1e\xdb\xf3\x9d\xf9\x6e\xad\xcf\xd4\x21\xff\x52\x1d\xa2\xf1\xdf\x41\xbe\xea\x10\x69\x79\x4e\x87\xf6\x7a\xb2\x68\x11\x4d\xff\xf8\x56\xc3\x16\x1d\xa2\xd2\xf5\x56\x36\xbd\x2e\xa3\x8e\x13\x24\x24\x8a\xc5\x74\x56\x60\xfd\x54\xb5\xb2\xaf\x33\x5e\xe6\x34\x73\x37\x54\x70\xa3\xa5\xee\x11\x9e\xd5\x5e\xed\xf1\xdd\xd5\x5c\x12\x9e\xf9\x9c\x4d\x1e\xe6\xeb\xba\x21\x31\xf0\xc8\x53\xdf\xab\xe5\x7b\x07\x9f\x8d\x36\x2b\xce\x75\xe4\xe2\x89\xa3\xbd\x8d\xdf\xfb\x53\xad\xd2\x1b\xf0\x70\xbe\xaf\xbd\x5f\xfb\x61\x21\xe8\xa2\x1f\x82\x32\xe4\x41\xd0\x2d\xd1\xa4\x0c\xde\x4f\xce\x0b\xf2\xd4\x35\x10\x94\xf5\x94\x8a\x04\xeb\x23\x5b\x42\x77\x2b\xda\xdd\x36\x23\x37\xb4\xa3\x6b\xb5\xb1\xaa\xe1\x18\x0f\x3b\x66\x37\xdf\xc4\x66\x98\xda\xd9\x7f\xf3\x17\x2d\x47\xeb\x1a\x8e\x30\x7b\xd8\x5a\x71\xef\xe6\xdb\xef\x69\x93\x25\x47\x92\x8c\xc3\x46\xa6\xf1\x66\x38\x4e\xfa\xd3\xe5\xff\x6c\x18\x7d\xcd\x70\x05\x78\x7e\x3d\xd6\x3c\x98\xea\xd8\x01\xe8\x92\xdd\xeb\x7f\xd3\xf6\x6f\x5d\xf3\x69\x48\xe2\x65\x9e\xe2\x4c\x4e\xf4\x2b\x40\xec\xe2\x43\x3e\xfa\xe8\x89\x41\x8d\xc0\x23\x38\x0a\x32\x4a\xf6\x63\x8b\x4d\x12\x17\x64\xe5\xe0\x82\x17\x20\x4b\x10\x55\xa1\xd5\x60\x58\x20\xa1\x27\xc9\xde\xbb\x6a\x08\xc9\x0a\x97\x06\x40\xe1\x7f\x02\x00\x00\xff\xff\xda\xfb\xc1\xb4\x1c\x4a\x00\x00")

func compiledEnergyBinBytes() ([]byte, error) {
	return bindataRead(
		_compiledEnergyBin,
		"compiled/Energy.bin",
	)
}

func compiledEnergyBin() (*asset, error) {
	bytes, err := compiledEnergyBinBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Energy.bin", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"compiled/Authority.abi": compiledAuthorityAbi,
	"compiled/Authority.bin": compiledAuthorityBin,
	"compiled/Energy.abi": compiledEnergyAbi,
	"compiled/Energy.bin": compiledEnergyBin,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"compiled": &bintree{nil, map[string]*bintree{
		"Authority.abi": &bintree{compiledAuthorityAbi, map[string]*bintree{}},
		"Authority.bin": &bintree{compiledAuthorityBin, map[string]*bintree{}},
		"Energy.abi": &bintree{compiledEnergyAbi, map[string]*bintree{}},
		"Energy.bin": &bintree{compiledEnergyBin, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
