package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _photo_jpg = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x90\x79\x38\xd4\xdf\xff\xfe\x5f\x63\x16\x4c\x61\x66\xcc\x08\x65\x1f\x64\x49\x85\xec\xc4\x08\x63\x10\xc6\x96\xad\x45\x51\x23\x7b\x46\xd9\xc7\x52\xb2\x67\xac\xc9\x12\xcd\x64\x86\x44\x22\x4b\x89\x10\x51\x52\x96\xb1\x97\x5d\x48\xc5\x34\xca\x96\x9f\xf7\xef\xfb\xc7\xe7\x3e\xcf\x3f\xee\xeb\x3c\xef\xf3\xb8\xee\xeb\xec\x8d\xec\x4d\x01\x08\x82\xa9\xb9\x29\x00\x02\x01\x00\x68\xff\x00\x7b\xd3\x80\xa2\xc9\x2d\x92\x17\x00\x98\x9b\x2b\x02\x3c\xc0\x7f\xd7\x27\x41\x5c\x00\xf7\xbe\x13\x05\xfe\x4f\x66\xfe\xfe\x57\x7d\x3c\xf7\xb3\xe3\x40\x3c\x00\xe6\xe2\xfa\x6f\xf6\x05\xd9\x1f\x28\x0f\x14\x0a\x81\x40\xe1\xdc\x08\x18\xcf\x41\x04\xe2\xe0\x01\xf8\x01\x14\x1f\x3f\x0a\x81\xe4\x47\x0a\xf0\x1f\x40\x08\x21\x05\xd0\x48\x61\x0c\xe6\xa0\xc0\x21\x14\x06\x8d\x40\x21\x51\x48\x10\x18\x02\x81\x41\x61\x70\x18\x0c\x8e\xe0\x3f\xc0\x87\x12\x44\xa1\x50\x18\x84\x00\x02\x21\x88\x41\xf0\x21\x30\x68\x24\x02\x85\x40\xa0\x50\x82\x02\xfc\x08\xc4\x7f\x56\x00\x29\x80\x40\xa2\xf6\x0d\x82\x1f\x21\x88\x42\xf3\xa3\xf6\x5a\x01\x24\x0f\x60\x08\x18\x82\x41\x48\x80\x6b\x9f\x87\x04\xed\x75\xfc\xd7\x15\x0c\x80\x41\xc0\xff\x04\x85\x71\x83\x21\x3c\xbc\xfb\x3b\x5d\x04\x00\xda\xef\xcb\x05\xf9\x4f\xdc\xff\x65\x40\x5c\xfb\xd5\x61\x48\x6e\x94\x14\x70\x52\x10\x6d\x64\x2b\xad\x7a\x29\x86\x47\x06\x17\x88\x35\x8e\xcd\x10\x22\x96\x94\x1e\x92\x57\x0f\x7a\xf6\xe6\x33\x66\xff\xf1\x11\x10\xc0\x05\x06\x81\x40\xff\x43\x43\xf7\xb9\x30\x2e\x10\x0f\xb0\xbf\xd5\x43\xee\xa3\xf7\xc1\x60\x30\x1c\x0a\xfd\xff\x91\x7d\x34\x12\x22\x05\x40\x51\x27\x61\xd2\x46\xb6\x82\x97\x02\x55\x63\x32\xb8\x31\x32\x38\x62\xc9\x3e\x70\x1a\x2d\xe4\x11\xfb\xeb\x4c\x10\xb5\x6d\x0d\xbb\x37\x06\x1c\xdc\x2f\xcc\x85\x04\x23\x81\xd3\x00\x7b\xdb\xe8\xf8\xd1\xc1\x9e\xc1\x07\x5e\xd5\xc6\x59\xf0\xf5\x88\xe7\x8c\x0b\x57\x74\x6e\x11\x5e\xdf\xdf\xb4\xfe\xbd\x4d\xbe\x44\x17\xbb\xb3\x7b\xfa\xe4\x6e\xdb\xd7\x83\x1a\x51\xd6\xec\xc7\x4b\xcb\x4e\xd9\x43\x5f\x8a\x68\x24\xc9\xef\xe8\x33\x02\x56\x97\xb3\x36\x97\xd4\x7c\x80\x4a\xe2\xb2\x8e\x1a\x93\x70\x3d\x5b\x1b\xf4\x13\x49\x17\xd7\x6b\xe7\x5e\xdd\xa2\xe2\xb5\xbe\xdf\x48\x9b\x0c\x28\xa1\xc5\x16\xa6\x76\x6d\x4e\x90\xd9\x9d\xd6\x5a\xef\xc5\xa6\x77\x7a\x83\x57\x50\x28\xdf\xdd\x6b\x4e\x2f\x57\xe4\x8c\x52\xb3\x4f\x06\xc6\x36\x75\x9b\x78\xd8\xfd\xfa\xd1\x13\x00\x8f\x73\x04\x77\x12\x4d\x4f\xfc\x53\x2b\x7e\xfc\x7b\xe7\x4c\xd8\x51\x2d\x7d\x1d\x4e\x3b\xe3\xfa\xca\xf0\xf0\x9f\xaf\x5c\x6c\x21\xc9\x04\x68\xe7\xc0\x97\x33\xf4\x0b\x13\x85\x5d\x38\x1e\xcd\xb7\x3d\xbd\x81\x85\xa4\xcf\x09\x8c\x9a\x02\x78\xce\x4e\x10\x62\x66\x0f\xc8\x55\x78\xda\xf5\x33\xbe\x02\xfc\xa9\x38\xea\xae\x57\x34\x3e\xa7\xa2\xcb\xb6\x24\x3a\x7e\x3b\x45\xe2\x56\xd7\xe0\x68\xa0\x19\x26\x60\xc4\x77\xe0\xa3\x5d\x89\x59\xcd\xe7\x9a\x94\x50\x0a\x63\xf9\x20\xb9\xc3\xda\xa4\x76\xce\xf1\xc7\xe6\x0d\xa5\xca\xbb\xb5\xae\xe1\xc1\x49\x89\xe1\x14\x74\xca\xf4\xf3\x3f\x90\xe2\x52\xeb\x16\xe3\x7e\x1e\xa7\x82\x56\xce\x2e\x6c\xf1\x4f\x9b\x45\xec\xdf\x1b\x7f\x0b\xde\xf8\x3b\xf1\x29\xf7\x77\x07\x85\x64\xf5\x29\xac\x43\x30\x4a\x81\xbc\x92\x3f\x4e\xc8\x71\x1a\x16\x57\xc4\x4b\x1b\xd8\xda\x16\xfe\x8d\x53\x1d\xbe\x91\x79\x9e\xda\xda\x06\xad\x00\xbd\xe2\xe7\xea\x70\xda\x51\x41\xf0\xa5\xfa\x95\x65\xab\x7f\xdf\xc1\x33\xca\xd3\x1f\xa5\x3e\x0c\xfe\x6e\xd3\x38\xf3\x3d\x4b\xce\xb2\x38\xb8\xd2\xf5\x8b\xe1\x52\x5c\x7e\xdd\x05\x8f\x18\xd3\xb2\x91\xac\x82\x3a\x51\x84\x45\xe6\x50\x91\xdb\xcd\x06\xdc\xb3\x65\xb5\x46\xe3\x8c\xc6\x73\xbb\x60\xe6\x11\x1d\x22\x72\xb8\x6a\xe9\xc7\x4e\x00\xaa\x79\x8d\x97\xc4\x05\x70\x4d\x1a\x4e\x80\x17\xae\x30\xdd\xb1\xe3\x30\xf5\x32\x5c\x89\xe3\x31\xd2\x1f\x93\x11\xf6\x3d\xa6\x72\xf5\x79\x5f\x5e\xc7\x0f\x65\x22\x07\x4e\xb7\x04\x9c\xf6\x0f\x1c\x4c\x36\x0e\x9e\xae\x5e\x65\x2d\x74\x7f\x2f\xaa\x70\x9a\xd3\x3f\x2b\x5a\xa0\xf0\x06\xdd\x67\x5b\xa1\xf6\x91\xd7\xc2\x04\x4b\x03\x97\x65\xf1\x2a\xd9\xa6\x96\x2e\x89\xdd\x6c\x38\xe1\xe1\xfc\xe6\xdd\x14\xc7\x9a\xf7\xa0\xbb\x8a\xc4\xe6\xd0\xa3\xef\x4e\x87\xc5\x8f\x0d\xf8\xd0\xd4\xb4\x65\xa8\xdb\x12\x8b\xff\xda\x6f\xfe\x10\xe2\x8e\x2a\x93\x77\x78\x39\x11\xc6\x91\x98\x6c\x60\xdf\x3b\xf1\xf7\x83\x35\xfb\xbc\x61\xfe\xe0\x97\xa7\xd4\xf7\x3f\x68\x6f\xfb\x5e\xf1\x5d\xfe\xa0\x05\xbc\x16\xf6\x69\x50\x4f\x1e\xc7\x5f\x6a\x77\xb5\x9d\x55\x9d\x3d\xfd\x2f\x22\xb3\x16\x68\xff\xa6\xed\xcb\x4e\x0b\xc2\xd4\x11\xe6\x3c\x55\x1a\x8b\xc2\x9e\xdb\x8c\x38\xa8\x2a\x8a\x1e\x32\xc9\xe1\xb6\x9d\xe8\x9c\xbb\xd0\x37\xda\xf1\xc1\xe6\xf3\x82\x88\x98\xbe\x87\x4d\x16\x4d\xbe\x35\x0c\x98\x97\x32\x36\x7e\x0c\xa8\x66\x67\x29\xbc\x4f\x34\x4e\x56\x94\x3b\xd1\x57\x26\x34\x97\xa4\x3f\x8c\xc9\x87\x40\x6b\x7a\x49\x1d\xf4\x26\x6f\x1e\x03\xc7\x30\xfe\xb7\x75\xbb\xe7\x88\xd6\x5a\xd0\xe7\x1f\x79\x56\x08\xb5\x03\xba\x96\x33\x03\x5e\x8e\xd8\xf4\x1a\xc3\x57\x61\xc9\xe9\x27\x64\x02\xef\x68\x9c\x69\x17\xb9\x5c\x1e\x1a\x53\x7a\xff\x17\x21\x53\x55\x94\x5c\x8d\x1b\x20\xd5\x50\x6c\x74\x29\x36\x86\x9a\x60\xbe\x96\x94\x4a\xed\x3d\x00\xec\x70\xcd\x11\x9a\x65\xa7\x62\xd2\x33\xe6\x8d\xd9\x7c\x12\x0b\x0e\x0a\x38\x70\xd7\x56\x8c\x2a\x58\x67\xa6\x37\x15\x96\xc9\x82\xc7\x33\xc7\x97\x33\xaf\xa4\x2b\xec\xf4\x26\x14\x98\x9e\x74\xae\x3b\x7b\x29\x7f\x20\x49\x51\xd1\xff\xa0\xc3\xab\xc6\xf3\xad\x74\x51\xed\xac\x68\x86\xbb\x69\x3f\x92\x2e\x45\x34\x45\x3a\x97\x76\x3d\x5b\xb5\xef\x5f\xfd\x56\x7a\xa5\x33\x8b\xd8\xaa\x6b\x94\xad\xe9\x77\x53\xb9\x88\xcb\xd5\xba\xf3\xc9\xac\x44\x16\x25\x4a\x8f\x4c\x0f\xa5\x1f\xb2\x35\xb2\xbc\x6f\xb3\x60\xcb\x4e\xe6\xad\xde\xfe\x19\xd5\xea\xe1\xf3\x5e\xd9\xb7\x29\x3a\x6f\xbd\x03\x03\x6b\x66\xf4\x1f\x48\x47\xe3\xe1\x72\x6f\x8e\xc9\x17\x76\x4b\xcb\x3e\xee\x67\x9e\x4c\x34\x4d\x5e\x0a\xae\xc4\x35\x58\x39\x32\x6e\x9f\x86\x1e\x47\x45\xd8\x15\xab\x6f\x04\xdd\xf8\x91\x31\x69\x36\x5c\xff\x5a\x9a\xee\xc9\x57\xe2\x40\x36\x68\x0b\x8b\x7c\x34\x3e\xd0\x12\xf9\xb2\x53\x26\xe6\x96\x95\xf3\x5b\xa3\xca\x27\x12\x3b\xdf\x4e\xb3\x9d\xfc\xd5\xd3\x2b\xbf\x8e\xfe\xe4\x93\x48\x4c\xb5\x24\x9d\x17\x4b\x36\x0d\x16\x3e\x91\x85\xb1\x21\x4b\x7c\xbc\x15\x84\xdc\x96\x6f\xdc\x9a\x8a\x2a\x8a\xc5\x34\xb0\x93\x98\xe1\xe3\x17\x26\x8a\x28\x51\xd1\xdb\x71\x2d\x0a\x7a\x04\xa1\xaa\x7c\x85\x74\xc5\xcb\x1d\x9c\x43\x3c\xff\x26\x87\x76\x6d\x94\x89\xa6\x87\x3f\x97\x86\xd9\xb9\x98\xfe\x55\x2a\x8a\x47\xde\x63\xf4\xb7\xb7\x6c\x3b\x99\xc0\xec\xf2\x4c\xa8\x63\xcd\xe8\x47\xa9\x6d\x5a\x07\x12\xbc\x3b\x06\x26\x6e\x5a\xd0\xfe\xc9\x9e\x23\xf5\x5b\x49\xd6\x3c\xd1\x9d\xc1\xd6\x4e\x38\x78\x91\x88\xe9\xd9\x45\xe3\xab\xb2\x6f\xa6\x2a\x62\x9f\x40\xd2\x73\x17\x76\x63\xc4\x9b\xed\xa2\xee\xd4\x62\xcb\xda\xe2\x2e\x4b\x79\x74\x88\xb5\x7e\x09\x30\x53\x9e\xea\x5c\x97\xd5\x75\xc8\xa9\xf4\x76\x36\x17\x45\xd3\xb2\x94\x18\x30\xc4\x71\x5d\xf1\xcc\x70\xc1\x24\x3e\xf9\x5f\x91\xf8\x6a\x8e\x5f\xcb\xc5\x52\x4e\x4a\x39\xa6\xed\x6b\xba\x79\x90\x1b\xee\x9e\xe3\x43\x02\xb2\x70\xe4\xf0\x6b\x39\x76\x9f\x82\x59\x8b\x5b\xff\x70\x03\x88\x5c\x8b\xa2\xe0\xcf\x8f\x5d\x4c\xde\x4c\x8b\x0f\xf9\x79\xa8\xd3\x5f\x86\x66\xc2\xef\xe4\x60\x9f\xd3\x33\x72\xab\x18\x55\x04\xa9\x0f\xc2\xd0\x89\x64\xe0\xf6\x01\x12\x8c\xbb\xee\x7d\x04\x6b\x62\xcb\x8b\x72\x20\xc9\xd7\xd6\xcf\x6d\x18\x4b\xe6\xb1\x99\x37\xf1\xf1\xf1\xc3\xf3\x98\xa5\xaf\x31\x8e\x30\xeb\x64\x42\xcf\xe2\x0e\x89\xf5\x8e\xa8\x1c\x4d\x3a\x88\x5a\x07\x2b\xda\x5a\x10\xec\x99\x92\x3d\xb5\xb1\x1e\x26\x0f\x25\xca\x99\xe9\x5e\x5d\xf8\xb6\x99\xb8\x34\x6a\x7e\x7d\xcd\x55\xd7\x50\x23\x0a\x7a\x3d\xe1\xce\x43\x1a\xb8\x44\xa9\xe9\xcd\xa0\x59\x7c\x2e\xac\x83\x3a\x45\xc1\xc8\x4a\x0f\x92\x4e\xb9\x63\x2b\x5b\x43\x87\x9a\x19\x7f\xc0\xa5\x57\x9c\x6f\x88\xe8\x5e\x2a\xc3\x4e\xbf\x4d\xa1\xf3\x5e\xe8\xa8\x63\x14\x40\xe2\x17\xd3\xc2\x9a\xb6\x4b\x4e\x91\x42\xb2\xab\x89\x85\x8d\x99\x35\x20\xd1\xda\xde\x1e\xac\xe9\x1f\x31\xb9\xc0\x70\x9d\x4d\xf7\x87\x65\x81\x8a\x8f\x9c\x71\xc8\xdf\xe6\x84\x78\xae\x3a\x48\xaa\x29\xaf\xa2\xb1\xaa\xf2\x7d\xec\x07\xb7\x4e\xa4\xec\xbf\x23\x55\x1b\x24\xf3\x54\xd1\xf1\xa2\x81\x9a\xfb\xca\x66\xb7\x7c\x88\x79\x49\x55\xa9\x5d\x67\xd1\x01\x2b\x8d\x2a\xe2\x1d\x0a\xbe\x74\xe3\xf5\x2d\x99\x6c\x27\x62\x4d\x2e\x9a\xdd\xaa\x55\xf9\x27\x4f\x37\xc3\x3f\x60\x15\x28\xfc\x28\x3c\x68\x35\x70\x92\x30\x8b\x67\x80\xde\xa2\xbe\x34\x36\x11\x1d\x23\xcc\xe4\x3d\x66\x7b\xe0\xf6\x8f\xfb\x5d\x0e\x11\x4e\x5d\x0c\x29\x19\xab\x3f\x1b\x37\xa2\xdc\x16\x77\xcd\xe1\x91\xe3\x42\xcd\xb9\xb2\x6a\x29\xe5\xcd\x90\x66\x77\xbd\xe8\x45\x6d\xec\xd5\x61\x86\x92\xf7\x67\xc7\x6f\x24\xe4\xb4\x26\x42\x23\x73\x20\xb2\x5f\x97\xa3\xf6\x15\x56\xc1\x36\xfd\x16\x54\xcc\x2d\x69\x96\x42\x5b\xc7\x6a\x9c\x8a\x17\xde\xc8\x15\x7c\xf1\x87\x4a\x3f\xf3\x2e\x33\xcd\xf0\xa6\xd5\x71\xcd\x42\xd7\x34\x57\x87\xec\x86\x1a\x7b\xfa\xd3\xc3\xf1\x8c\xa2\xb9\xb5\x41\x51\xb4\xff\x66\x4a\x44\xe7\x6b\x91\x8f\x05\xb6\x2e\x4c\x71\xec\x04\x75\xac\x1b\xcb\xd1\x98\x98\x4e\xf8\x07\x71\x1a\x0d\x94\x3d\x45\x90\xb6\x17\x80\x9f\x02\x52\x22\x07\xb9\xcf\x2b\x18\x08\x0f\x49\x4f\xa6\x14\x37\xfd\xa5\xb1\x9c\x95\x00\x7b\xf5\xcc\x09\xa6\xb9\x04\xea\x1e\x37\x53\xab\x62\xfc\xc5\x2c\x1a\x6f\x4b\xdd\x7e\x4e\x33\x88\xc7\x6a\x5d\x9e\x58\x63\x65\x1f\x3f\x32\x25\x72\x54\xb8\x6f\xe0\xd1\xea\x62\x70\x2a\x0d\x34\xf3\xe1\x3c\x61\x26\xda\x3a\x44\x7c\xd4\xc1\xcb\xd9\xb4\x1f\x14\x7a\xfd\x1c\x57\xf0\xb0\x26\x2d\xef\x95\x80\x08\x97\x3c\x09\x7c\xa8\x67\x77\xd9\x5f\x4c\xa8\xa4\x3b\x79\xf1\xdc\x4b\xdf\x4d\xc5\xa4\x81\xf5\xab\xa6\x17\x7e\x75\x83\xc2\xd7\x07\x13\x9c\xb2\x52\x29\xe8\x12\xb3\x10\xb5\x18\x2a\x30\x30\x48\xd9\x5e\xf6\x44\x8d\x3f\xc4\xa7\x5f\x61\xb4\x26\x82\xc7\xb0\x23\x37\x57\x10\xc6\x8b\x7a\xd7\xad\x73\x74\x8e\x32\x59\xe3\xd7\x2b\xe7\x0c\x89\x6d\xbf\xf3\x62\xeb\x84\x8c\x99\xb7\x05\xe9\x3a\xb9\x8c\x27\xdd\x96\x5c\x01\x1a\xe1\x1b\x93\xef\xf1\x0f\xdd\x2c\x9b\x95\x87\x10\x8b\x79\xf9\xe3\x6b\xf7\x3a\xaf\xd0\x3c\xe7\xb9\x08\x69\x9f\x22\x54\x35\xeb\x6f\xff\xbc\xa0\x97\x56\xe2\xda\x43\xf7\xdd\x72\x42\xe0\x06\x01\x18\x06\xc9\x88\xa7\x64\x43\xe2\x78\x59\xb7\xd1\x0c\xd3\x39\x3e\x39\x3b\x0c\x66\x75\x4e\xbd\xe7\xfd\x37\x8a\x15\xe1\x22\x73\xe8\x18\xce\xc5\xf1\xe9\x45\x38\xee\xf5\xdd\x5c\x11\xf3\x0c\x32\x54\x0e\x9f\x82\xac\x1c\x38\xcf\x72\x69\xc6\x55\xaa\xce\x6a\xec\xdc\x10\x9b\x5c\xf2\xd2\xbe\xa5\xc5\x4e\xe4\x31\x34\x0d\xa1\x66\x2b\x45\xb7\xa1\x22\xd4\xe9\x92\x05\x90\x1f\xf3\xf5\xb1\xf9\xbf\x44\x2e\x61\x63\x3e\x94\xd2\x68\xeb\xdd\x71\xd7\xf0\xbc\x6a\x15\xb6\x2c\xba\x37\x17\xac\x40\xb4\x47\xcd\x89\x5d\x26\x68\x39\xa9\xab\xf8\x69\xd4\xc2\x4f\x61\xa1\x0e\x42\xaf\x33\xee\xe1\x97\x19\x1f\x92\xf2\xd4\x54\x29\xf7\xd4\x91\xac\x5f\xea\x3d\xb6\xb1\xea\x3c\x8c\x6e\x76\x2d\xb6\x34\x91\x32\x39\xfe\x5c\xea\xfe\x75\x95\xe3\xda\xbe\x40\xb6\x58\xc7\x5c\x50\x78\xc1\xa2\x5e\x51\x8b\xae\xe3\x24\xd1\xce\x72\x80\x87\x55\xd2\x63\x7a\x5f\x7f\xcb\x19\x27\x93\x77\x37\xb1\xfd\x72\x5c\xc3\xfd\xa6\x2d\x39\x86\x7b\x09\xa1\xd0\xec\x73\x46\x25\x13\xa4\x01\x97\x9e\x0b\x1a\xd9\x25\x71\x0f\x1c\x99\x65\x3a\xa4\x7a\x36\x3f\x4d\x24\x67\x32\x52\xad\x55\xa6\x8f\xbd\x58\x1e\x2c\xe6\x7d\xb6\x5e\xa2\xfb\x40\x4f\xe0\x85\xb9\xfc\xf0\x23\xab\xb9\x2a\xf4\xed\x5b\x86\x12\x27\xb2\x25\xee\x0e\xa5\x62\xcc\xf8\xd7\xf8\x74\x58\x41\xe5\x8f\x5d\xaf\x51\x71\x5c\xc1\x31\x6b\xbe\xc4\x8b\xbc\xd8\xf4\x90\x91\xb4\xc9\xb0\xdc\xf2\x10\xfb\xb9\x33\xe5\x6e\xfc\x52\x5d\x82\xcd\x70\x9f\xea\x29\x62\xf8\x55\x4a\xf4\xb0\x98\x7d\x64\x82\xc2\x10\xbd\x03\x88\x1e\x2b\x1e\x38\xe0\xb3\x78\xb7\xa1\x42\x2e\x58\xcd\xc9\x06\x24\x7d\x06\x66\xcf\x93\xf7\x00\xcb\xf9\x76\x71\xf3\x32\xb4\xec\xda\xb2\xe7\x03\xa2\xfd\x3b\xc7\x89\x32\x16\x4d\x68\x70\xe2\x31\xf8\x61\x9a\x40\x0b\x77\xcc\xbb\x2e\xe2\x89\xd9\xda\x65\x56\xec\x7d\x8e\xf5\xca\x5d\xd5\x94\x47\x89\x35\x65\x3b\x21\x4c\x78\xca\x99\x63\x0a\x2e\x17\x33\x3e\x3d\x38\xec\xc7\x63\x70\xab\x72\x64\xe0\x94\xc6\x4a\x55\x2b\xad\xca\x42\x8c\x49\xb8\xd7\xdc\xd2\x56\xbb\xcc\x65\xb7\x95\x75\xcb\xb9\xf5\x48\x41\x38\x99\xfc\x61\x72\x12\x8b\x91\x14\xdf\x49\xf8\xb1\x72\xc7\x5d\x13\x79\x95\x47\xc1\xf4\x52\x8e\x46\x3c\x15\xe4\x6c\x20\x2a\x28\x4f\xf6\x66\x29\xeb\xcd\x7d\x87\xe4\x75\x99\x7d\xc5\xb8\x78\xaa\xf4\x24\xf8\x1a\xf5\x37\xa5\xca\x8e\x79\x27\x17\xdd\x78\x7b\xce\xee\x63\x91\x18\x31\x26\x8d\x97\xbb\xa0\x4c\x4b\x0b\xff\x9d\xdb\x48\xa4\x14\xd3\x54\x5e\x40\x16\xbd\xdb\x50\xdb\x3a\xee\xa1\xce\x40\x4a\x43\x62\xe5\x8c\x5a\x55\xd4\xa6\x3a\xd9\x3a\x66\x15\xce\x0c\xd2\xd9\x52\xa1\xbb\x9e\x30\x2e\x7f\x1a\xae\xf7\x77\xb7\xf9\x57\x1b\xee\xd7\x42\x24\xbb\x19\xef\x55\x22\x54\x28\xd3\x4f\xdf\x94\x70\xff\x5a\xc3\x6c\xaa\x7a\x15\x72\x31\x61\xe1\xd6\x03\x9b\xec\x8b\xf7\x7a\xec\x5e\x8a\x8b\x9b\xcf\xce\xc8\x74\xb9\x84\x6e\xe8\x5d\x2d\x82\x60\xc5\xa4\xd9\xdd\x1b\x64\x14\x00\x50\x1f\x5e\xb5\xff\xf6\xe0\x8d\x9d\xd7\xa8\xac\x57\x70\xf5\x5f\xf5\x8d\x6b\x18\xca\x28\x21\xee\x52\x35\x5d\x26\x23\xc8\xd1\x52\x0b\x14\xa7\xe8\x0b\x58\x7e\x4c\x51\x34\x79\xe6\x50\xe6\xf6\x62\xa2\x6c\xe1\x6d\x85\xa9\x53\xae\xe7\x89\xf7\xb6\x7b\x40\x12\xfe\xbd\x55\x61\xb7\x1d\x05\x8a\x91\xf2\x26\xfc\xa2\xfd\xbd\x05\xb9\x52\xf1\xf8\x1e\x47\xc9\x67\x48\xb5\x06\xdd\x08\xd5\xdb\xc6\x17\x9a\xe8\xd8\x8a\x72\xb9\x11\x3c\x86\x81\x58\xf5\x29\x31\xe8\x81\x50\x13\xc9\x5e\xef\x77\xb8\xf1\x0d\x1b\x70\x09\xc1\xf3\xa2\x18\xf6\x36\x8d\xee\x64\x82\x4b\x1c\x12\x4d\x89\x35\x2d\xf0\x1f\xd7\xd0\x0b\xef\x33\x53\x91\xf3\x3e\x31\xf0\x4f\xa9\xd4\x60\xe2\xf6\x09\xd2\x93\x8a\x60\xaf\x3f\x58\x59\xf3\xc0\xec\xc6\x46\xdd\x42\x2d\xe2\x47\x01\x00\xeb\xa9\x4f\xb6\xb3\x34\xfe\xc0\xed\xa8\xae\x63\x8f\x4b\xc0\xd5\x99\x17\x2a\x42\x96\x33\x82\xc2\xbe\xe3\x5f\xc6\x49\x4f\x85\x11\x08\xc5\x85\x8a\x23\x3e\xa2\x5f\x3c\xe0\x16\x31\x90\xc4\x79\x8b\x8a\x73\xf0\xe0\xb6\x6c\xba\xce\x53\xa4\x46\x87\xea\x4b\xb2\xad\x2f\x83\x2a\x94\x53\x6a\x46\xe6\xc1\xca\xae\x78\x1a\x1e\x4d\x1b\x55\xc0\x71\x6f\xfd\x0e\xfc\x40\xbd\x4f\xa3\xae\x3c\xb5\x8c\xcc\x3d\xad\x1a\x8d\x4c\x15\xf7\xb4\xca\x98\x5f\x3d\x76\x5e\x94\x8b\x72\x52\xdd\x02\xc8\xfb\x49\x07\x7f\x95\x41\xc5\x55\x71\xc8\x47\x33\x8d\x4b\x28\xd1\x04\xd1\x48\xf0\xee\xd3\x34\x8a\x06\x0e\x02\x08\x54\xfe\x4a\x0e\xd7\x56\x39\xca\x50\x1d\xbb\x5a\x4a\xa6\x85\x53\xce\x83\xf9\xbc\x5f\x27\x02\x6e\x03\xec\xdf\xe7\xf6\x80\x45\x77\x8c\xb5\xd1\x1e\x80\xfe\x97\x77\x7c\x67\xb5\xa8\x98\x63\xf1\x57\xfe\x77\x97\xdb\x12\x74\x2d\xe2\x5b\x6f\xce\xbc\x77\x09\xb9\x37\x21\xc0\x5e\x64\x4d\xdc\x35\xd9\x7c\x3d\x23\x41\xf7\xac\x19\xe4\xcf\xe9\x0b\xdb\x61\xf9\x69\x89\x18\xff\x6e\x40\xaa\x71\xf0\x0c\xaf\x3e\xcd\x24\x73\x78\x0f\x38\xf9\x3b\xc4\x39\x5b\xc1\xf9\x42\x72\x87\xf7\x1a\x5a\x7b\xa8\xfc\x0d\x42\xf9\xd8\x21\xf0\xd8\x97\x05\x15\x48\x26\xfa\xeb\x5b\x83\x12\x78\x98\xa2\xff\x84\xb5\xda\x82\xeb\x71\xc9\x97\x22\x74\xe1\x4f\x03\x94\x23\x9b\x46\x6b\x0e\x8d\x52\xc0\x8d\x19\x9d\xf4\x3c\x29\xb3\x0a\x91\xc7\xee\x67\xda\xf1\x76\xbe\x37\x18\xeb\x93\x76\x62\x72\xdb\x15\xd1\x4b\xa2\xa5\x33\x01\xf0\xaa\xbf\x42\xdf\xec\x88\x2f\xb8\x8c\x13\xce\x59\x75\xf3\xdd\x96\xc3\xe2\xb3\x9b\x7b\xb2\x1e\xf6\xa9\xc5\xe8\xd8\x67\x31\xa6\x9f\xd8\x29\xa8\x06\x87\xdc\xbf\xa1\x2e\x9e\x18\x3e\x74\xc1\x41\x3f\x49\xef\x10\x6c\x2c\xc8\x00\x3d\x5f\x1b\xbf\xa3\xc0\xf3\x6b\xdc\x50\x58\x9b\xd8\x28\x19\xac\x90\x6f\xe3\xbd\x03\xaf\x93\x8c\x8f\xd1\x63\xe6\xe1\xbc\x78\x65\x02\x55\xbd\x4c\xaa\x9f\x75\x47\xe1\x99\xb7\xed\x62\xcb\x3a\xae\x29\x2d\x4d\x28\xd8\x24\xda\xda\x5a\x10\x69\x85\x46\x35\x0e\x8c\xd5\x4b\x2f\xa0\x4e\x95\x7f\x45\x79\x61\x13\xd7\x6a\xe6\x7b\x83\xf4\x5a\x4e\xab\xe4\xcc\xf4\x8f\xb9\x90\x0f\xb3\xee\xce\x9b\x9b\x3b\xda\x53\x9b\x1f\xee\x54\xa4\xa4\x20\x8e\x10\xc2\xaa\x6c\xaa\x44\x94\x02\x2b\x34\x2b\x08\x8e\x67\xf2\x5e\x94\x3e\xb8\x0b\xdf\x6a\xeb\xa7\xde\x8c\x3d\x29\x7e\xfd\x0d\x3c\x84\x59\xf0\xb9\x7e\xd2\x85\x54\x27\x9b\xc5\x54\xb0\x09\x10\xac\x24\x65\x1c\xe9\x01\x85\xc1\x63\x08\x34\x75\x0f\x98\x86\xe2\x1e\x90\xfc\xd2\xa3\x6b\xc8\xe2\x86\x4d\x36\x49\x67\x52\xe8\xf5\x48\xba\x93\xdf\x91\x57\xb1\x75\xc0\x91\xf3\x27\x0a\xa0\x7c\x4f\x2f\x17\x11\x5b\xa2\x9e\x0e\x7e\xd6\x43\x89\x04\x06\x8d\xc4\x3e\xe0\xc3\x22\x5f\xfd\xd0\x89\x3d\x2e\x9e\x71\xc7\x2b\xd0\x31\x49\xb9\xa8\xa8\xb0\x0b\x3e\x46\x36\x6e\x0b\x83\x48\xa2\xe0\xa7\x12\xb1\x14\xaa\xdf\x71\x30\xe9\x50\xb1\x37\xab\xfc\x9c\x90\x3e\x5b\x55\xf2\x31\xab\xa4\x5f\x56\x42\xd3\xfa\x26\x4f\x80\x90\xd3\xbb\xb3\xb8\xf4\x9e\xa0\xe2\x11\xae\x70\x8e\x5f\xb1\x5a\x15\x01\xc4\x0f\x6a\x88\xcd\x94\xe6\xef\x48\x6b\x18\x64\xad\xc5\x69\x1f\xbf\x6d\x95\xbf\xbc\x06\x07\xb1\xb2\x96\x7d\xdd\x79\x61\x5e\xbe\x58\x12\x52\xe3\x19\x37\x6f\xcd\x29\x4d\x33\x15\x00\xb2\xa9\xa4\xad\xff\xd8\x2e\xc2\x32\xba\x9a\x73\xc9\xb9\xd9\xfb\x49\x6d\x6d\x8b\x4d\xb4\x1e\x77\x74\xf2\x5f\xc7\x5a\x13\xeb\x3b\x55\x7f\x54\x9f\xef\x01\xc6\x89\x7b\xc0\x74\x95\xc5\x83\x4a\xeb\xa0\x9f\x4f\xea\xf2\x4f\x9d\x8c\x6e\x7f\xf1\x6c\x0f\xb0\x95\x0c\xe9\x6a\xb0\x96\x88\x08\x17\x8f\xc7\xb1\x6a\x6b\x7a\x27\xdc\x13\x72\xc9\xe1\x4d\x51\x23\x6e\x11\x05\xbf\xba\xd6\xfc\x23\x16\x6d\x0c\x52\x9f\xbf\xe3\xd9\xc8\x32\xa8\x2f\xd5\xb9\x69\x7f\x21\x2b\x85\x74\x9d\xf5\xa4\xf9\x96\x43\x51\xa9\x5d\x8e\x13\x1f\x3f\x65\x6b\x5c\x7b\x2d\x47\x57\xca\x60\xc9\xfc\xdf\x62\xb1\xa1\x79\x53\x76\x69\x61\x89\xb1\xf5\xf7\xf4\x6b\xbb\x8f\xb6\xc0\x6b\x06\x86\x97\xfb\xb2\xf8\x55\x36\xdf\xf4\x3d\x38\x9a\xfe\xe9\x21\x1b\xc9\x19\xe4\x6f\x6f\x22\x46\x29\xb6\x18\x5a\x9d\xca\x8f\x0b\x69\x2e\x6f\x6a\xfe\xb8\xec\x20\x4b\x7b\xf1\x7d\x32\x91\xf7\xab\x94\xe4\x9d\xd9\x17\xba\x97\xaa\xcf\xc2\xf0\x6b\x47\x82\x0b\x7c\xc5\x2c\xc7\x93\x43\x39\x4e\x0e\x48\xbe\x75\x5a\x70\xde\x1e\x50\x35\x8b\xf0\x0b\xac\x78\xaf\xf4\x69\xf6\xb4\x9e\x97\x77\xce\x48\xac\x0f\xac\xde\x50\x8a\xef\xa1\x0f\xb1\xc7\xbc\x5f\xee\x40\xb2\xcb\xc7\x42\xa1\xd0\x1d\x51\x6f\x0b\xfb\x07\x4d\x8d\x8c\xdc\x7e\xf0\x60\x79\xb2\x69\xa1\xf1\xe3\x4a\x37\x5e\xd2\xbb\xab\x7e\x91\x30\xfb\x82\xd9\x46\xc5\x2e\x31\xaf\x8a\xf9\xc9\x4c\x6c\x4f\x33\xe4\x70\xf1\x29\x89\xd0\xd4\x17\x1d\x22\x1e\x6c\xcb\xf5\xc2\xe5\x3f\xe2\x7c\xd2\x74\xd6\xa7\xac\xde\xda\x12\x02\x44\xfe\x43\x26\x42\xf2\xc9\xed\xb3\x42\x1c\xdc\xf7\x9c\xde\x45\x15\x87\xb7\x5a\x35\x49\xb3\xe5\x2a\xf6\x18\x8a\xf4\x59\xa3\x2f\xe3\x71\xfa\x31\x6d\x84\x48\x10\xc9\x20\x9e\xd3\x78\xd1\x4e\x1b\x3f\x6e\xda\x7b\xd4\x50\x63\xfe\xa4\x47\xa4\x9f\x40\x9e\x5a\x63\xe7\xb8\xb6\xab\x88\xc7\xf8\xe3\x50\x17\x85\x22\x90\x82\x7b\x49\x99\x6b\xec\xed\x0b\xab\x00\xdf\x1c\x54\x52\xb0\x51\xcb\xe1\xa5\x3b\xfd\xcb\x96\xb5\x6d\x28\xbd\x8b\x94\x2b\xec\x60\x94\x9a\x24\xec\xb9\xd1\xb7\x50\x5b\xb0\x05\xf5\xe3\x9d\x8a\x10\xd0\x4c\xbb\xf1\x09\xee\xb9\x5c\x16\x94\xdd\x1e\xf2\x23\x6e\x22\x84\xfb\x89\x6b\xdf\x98\x56\x6b\xa3\xad\x45\xde\x95\x3a\x7b\x45\x22\x2e\x38\x77\xb4\x99\x6a\xf3\x68\x46\x43\x6e\xa7\x62\x10\x29\xa5\x3a\x9e\xc8\x2b\xa4\xe5\x54\xcf\xc3\xaa\x94\x88\x8a\x5d\x5c\x98\x14\xea\xb7\xf8\x84\x0e\x7f\xd5\xb5\x79\xba\xbe\x7d\xc3\x48\xdf\x0e\xec\xb0\xdd\x6a\xd0\x2c\xd0\x72\x79\x11\x32\x15\xe2\x72\x3a\xbf\x06\x6a\x53\xce\x09\xae\xf7\x6f\xca\xfc\x14\xde\x25\xc5\xe6\xa3\x06\x1e\xe8\xe4\xe7\x79\xb3\xb1\x7c\x50\x1a\x9f\xe2\x5e\x77\xd8\x8d\xe9\x73\xee\xf7\xcb\xb3\xef\xd6\x84\xe3\xdf\x83\xea\xaf\xa0\xd1\xe8\x94\x20\x58\xe8\x97\xd0\x9e\xd7\x65\x22\xf2\xd2\x6c\x6f\xbf\xbc\x86\x15\x53\x37\xf7\xb2\xef\x65\xd5\x6c\x7d\xeb\xb3\xb9\xaa\xe2\x13\x26\x9d\xb5\x3e\x9f\x64\x94\xdf\x25\x7e\xa6\x61\xcb\xe1\x78\xcb\x70\xe9\x8b\xcb\x0e\x1c\x25\x16\xaf\xbc\x8d\xd3\xe7\x6d\x6f\x3d\xaf\xab\xe7\xb3\x3d\x72\x48\xb5\xf7\x22\x1a\x03\xc0\xaf\xcd\x4a\x80\x3f\x47\xdf\x49\x86\x55\xe2\xc5\x77\x0d\x47\x28\x4e\xbf\x38\xbb\xd5\x25\xef\x47\x60\xee\xce\x14\x9d\x8d\x1f\x55\xb6\x35\x66\x0e\x08\x50\x01\x6f\x69\x7b\xc5\x75\x95\x87\x6a\x50\xa5\x40\xbe\xf7\xe1\x1d\x2e\x9a\xd9\x3c\x1f\xca\x62\xa3\x03\xee\xc1\xac\x6f\xa4\xac\x87\x4d\x73\xdd\xd4\x2b\x5d\xdc\xf9\x11\x2d\x18\x27\xa2\x0f\xde\xdc\x0d\x0f\xb6\xa2\x1e\x96\xc3\x53\x71\xa8\x8c\x1b\x63\x1e\xfd\x91\x74\xa4\xa6\xd6\xa5\xe1\xc3\x9c\x32\x3f\x51\x72\x8b\xbf\xf6\xfc\xb7\x17\x2c\xd6\x23\xa1\xf4\x0f\x5b\xe1\x8d\x9c\x37\xc2\xdc\xdf\x1e\xad\x0c\xcf\xa3\x0e\x72\x26\xfe\x19\x7f\x96\xfe\xf2\x3c\xf5\xc2\x91\xab\xb3\x7f\x3b\x19\xaa\x6c\x7a\x3e\x9e\x37\xfd\x45\xd1\xeb\x30\x19\xbc\x3e\x47\x44\x2d\x59\x8c\x35\xae\x33\xfa\x48\x99\x25\xfd\x50\x84\x2a\x43\x99\x11\xdc\x69\xb4\xda\x8a\x7c\xa7\xd0\x57\x93\xd0\x72\x46\xb5\xfd\x89\xae\x61\x54\xf5\x60\x7a\xa7\x7e\xa7\xd9\xab\xe9\xf7\x2e\xc8\xb7\x11\x5a\x8e\x23\x8b\x44\x32\x68\xf4\xda\xa4\x74\x34\x03\xdd\x67\x7f\x7f\xe5\xae\x63\x1a\xeb\x5c\xa6\x9c\x27\xd3\xb6\x6d\xca\x21\x86\x4f\xd6\x4c\x38\x5a\x18\xbc\xd4\xac\x8d\x37\x48\x1a\x89\xd4\x29\xfd\x7a\xe2\xf3\xfe\x07\x6c\x74\xf5\x6c\x86\xff\xb2\xf0\x9a\xfa\xa2\xf1\x00\x97\xb9\xdb\x9d\xa7\x7f\x10\x19\x71\xf0\x91\x7d\xb9\x8b\x1a\xd3\xfc\x66\xb6\x36\x68\x37\x0d\x1b\xb8\x20\x7c\x3a\x7a\x90\xca\xd3\x77\x39\x65\x5d\x67\x5e\x1e\x57\xe2\x7a\x58\x82\xdd\xb0\xb4\x18\x7b\xdf\xe0\xca\x4f\x81\x85\x6b\x4c\xa3\x7b\x99\x51\xf9\x9f\xaa\xa2\x59\x56\x1d\xdb\x8a\x89\x0d\xed\xb0\xe6\xd1\x67\xcc\xba\xbc\x7a\xfd\xc3\xfc\x36\x4a\xef\xd8\xc7\x8e\x89\xdd\x5b\x9a\x37\xdf\xba\xba\x20\x31\xb4\x3a\xec\x47\x4b\x2e\x8e\x9e\x8b\x44\x1f\x5f\xeb\x7c\x3e\x0f\x96\xf5\x4d\x8a\xf7\xca\xca\x2d\x63\x8a\xc6\x1c\x86\x22\xe5\x07\xce\x2d\x88\x87\xc5\xbb\x4a\xec\x8d\xfe\xbf\x00\x00\x00\xff\xff\x67\xf5\xf9\xa8\xdd\x13\x00\x00")

func photo_jpg_bytes() ([]byte, error) {
	return bindata_read(
		_photo_jpg,
		"photo.jpg",
	)
}

func photo_jpg() (*asset, error) {
	bytes, err := photo_jpg_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "photo.jpg", size: 5085, mode: os.FileMode(436), modTime: time.Unix(1476052406, 0)}
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
	"photo.jpg": photo_jpg,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"photo.jpg": &_bintree_t{photo_jpg, map[string]*_bintree_t{}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}