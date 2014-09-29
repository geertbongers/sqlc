package sqlc

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
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

func sqlc_tmpl_fields_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x94, 0x52,
		0xcb, 0x6e, 0xdb, 0x30, 0x10, 0x3c, 0x9b, 0x5f, 0x31, 0x87, 0xa0, 0xb0,
		0x0a, 0xd5, 0xb9, 0x0b, 0xe8, 0xc1, 0x69, 0x98, 0x46, 0x05, 0x61, 0x17,
		0x36, 0x5b, 0xa3, 0xa7, 0x82, 0xa5, 0x68, 0x97, 0x88, 0x4a, 0xd9, 0x14,
		0xdd, 0x07, 0x84, 0xfc, 0x7b, 0x97, 0x74, 0x6a, 0x2b, 0x6e, 0x20, 0x24,
		0x17, 0x89, 0xbb, 0x9c, 0x9d, 0x9d, 0xe1, 0xee, 0xe5, 0x25, 0xe4, 0x6d,
		0xb9, 0xc4, 0x4d, 0x29, 0x38, 0x56, 0xd3, 0x25, 0xa6, 0x9f, 0xe4, 0xfc,
		0x3d, 0x9f, 0xf1, 0xc5, 0x54, 0xf2, 0x6b, 0xbc, 0xc1, 0x74, 0xf6, 0x05,
		0xfc, 0xba, 0x94, 0x4b, 0xc8, 0xf9, 0x01, 0xba, 0x2a, 0x85, 0xc0, 0x15,
		0x87, 0x98, 0x2f, 0x25, 0x56, 0xb7, 0x7c, 0x86, 0x52, 0x82, 0xf2, 0x0b,
		0x7e, 0xac, 0x63, 0x6c, 0xab, 0xf4, 0x9d, 0xda, 0x18, 0xb4, 0xbb, 0x5a,
		0x33, 0xd6, 0x75, 0xf0, 0xca, 0x51, 0x78, 0xf1, 0x35, 0xc7, 0x45, 0x40,
		0xf1, 0x16, 0x93, 0xf0, 0x67, 0x6b, 0x5a, 0xdc, 0xdf, 0x33, 0x16, 0x4f,
		0x20, 0x48, 0x68, 0x44, 0xf3, 0xcb, 0x78, 0x02, 0x4c, 0x3e, 0x7a, 0xb3,
		0xb6, 0xbf, 0xe9, 0xf6, 0xc6, 0x9a, 0xba, 0x42, 0x1b, 0xfc, 0x5e, 0x07,
		0x74, 0x6c, 0xe4, 0xd4, 0x0f, 0x13, 0x43, 0xeb, 0x36, 0x6c, 0x14, 0xd4,
		0xb7, 0xda, 0x40, 0xc6, 0xaf, 0xb0, 0x77, 0x86, 0xf5, 0xb8, 0xfe, 0xe7,
		0xb0, 0x2e, 0x18, 0xbf, 0x56, 0xda, 0x44, 0x9a, 0x54, 0x93, 0xf2, 0x6c,
		0xc4, 0x77, 0xe3, 0x9f, 0xaa, 0xde, 0xff, 0x2b, 0x13, 0x96, 0x70, 0xaa,
		0xa6, 0xba, 0x0c, 0xef, 0x1a, 0x57, 0xd9, 0x60, 0x1b, 0xc7, 0x46, 0x65,
		0x7b, 0x86, 0x7b, 0x4c, 0x9f, 0xe1, 0x43, 0x63, 0xdd, 0x09, 0x4f, 0x52,
		0xd6, 0x7b, 0xa7, 0x31, 0xd6, 0x78, 0x3d, 0x64, 0x2d, 0xc3, 0x8c, 0x1c,
		0x8d, 0xb3, 0x07, 0x4f, 0x51, 0x9b, 0x37, 0x61, 0xef, 0x1d, 0xf4, 0x24,
		0x7a, 0x7d, 0x01, 0x51, 0xf2, 0xf4, 0x34, 0x53, 0x7a, 0xa8, 0xc9, 0xa1,
		0xd3, 0x0b, 0x18, 0xc9, 0xf1, 0xd6, 0x9b, 0x6a, 0xe8, 0x61, 0x7a, 0x6d,
		0x8e, 0xb9, 0xee, 0xca, 0xd2, 0xc9, 0x6d, 0x0a, 0x24, 0x9e, 0x87, 0xa8,
		0xfb, 0x1c, 0x1f, 0xaf, 0x40, 0x64, 0xcc, 0x0f, 0x37, 0x05, 0x34, 0xcd,
		0xff, 0xf9, 0x7a, 0xd2, 0x0c, 0x7a, 0x8a, 0x06, 0x47, 0xd0, 0x53, 0xf6,
		0x28, 0xdf, 0x89, 0xef, 0x2d, 0x35, 0xce, 0xb1, 0x88, 0xff, 0x83, 0x1a,
		0x22, 0xaa, 0xac, 0x56, 0x81, 0xe4, 0xf1, 0xdd, 0x31, 0x38, 0x49, 0x3b,
		0x6b, 0x37, 0x3e, 0x5b, 0xbc, 0x1c, 0xbd, 0xb5, 0xcc, 0x9e, 0x5e, 0xbf,
		0x93, 0x9a, 0x57, 0x43, 0x26, 0xbb, 0xc8, 0x54, 0x24, 0xbe, 0x1c, 0xa9,
		0x4d, 0x91, 0xbe, 0x49, 0x0b, 0x15, 0x1a, 0x57, 0x11, 0xf4, 0x6f, 0x00,
		0x00, 0x00, 0xff, 0xff, 0xa5, 0xea, 0xd0, 0x69, 0xb6, 0x03, 0x00, 0x00,
	},
		"sqlc/tmpl/fields.tmpl",
	)
}

func sqlc_tmpl_schema_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x9c, 0x52,
		0x61, 0x6b, 0xea, 0x30, 0x14, 0xfd, 0xdc, 0xfc, 0x8a, 0x4b, 0x91, 0x47,
		0xfb, 0xf0, 0xb5, 0xdf, 0x05, 0x3f, 0xf4, 0x61, 0x7c, 0x16, 0x4a, 0x7d,
		0xd8, 0x38, 0x19, 0x63, 0x48, 0xac, 0x51, 0xcb, 0x6a, 0xdb, 0xa5, 0xe9,
		0xc6, 0x28, 0xfe, 0xf7, 0x25, 0xa9, 0x68, 0x1d, 0x75, 0xca, 0xa0, 0x94,
		0xdc, 0x9b, 0x73, 0xce, 0x3d, 0x39, 0x89, 0xeb, 0x02, 0x99, 0xf8, 0x11,
		0x8c, 0xfd, 0x00, 0xc3, 0xc2, 0x8b, 0xc0, 0x9b, 0x93, 0xe9, 0x3f, 0x1c,
		0xe2, 0x99, 0x47, 0xf0, 0x08, 0xfe, 0x80, 0x17, 0x3e, 0x02, 0x1e, 0xf9,
		0x24, 0x02, 0x32, 0x6d, 0xa0, 0x0b, 0x3f, 0x08, 0xe0, 0x2f, 0x86, 0x60,
		0x1a, 0x11, 0x58, 0x4c, 0x70, 0x08, 0x3e, 0x01, 0xd9, 0x9f, 0xe1, 0x13,
		0x0f, 0x15, 0x34, 0x7e, 0xa1, 0x5b, 0x06, 0x75, 0x0d, 0xce, 0xff, 0xe3,
		0xfa, 0x70, 0x40, 0x28, 0xd9, 0x17, 0x39, 0x17, 0x60, 0x21, 0xc3, 0xdc,
		0x26, 0x62, 0x57, 0xad, 0x9c, 0x38, 0xdf, 0xbb, 0x9c, 0xa5, 0x79, 0x51,
		0xba, 0xe5, 0x6b, 0x1a, 0xeb, 0x9f, 0x89, 0x6c, 0x84, 0x24, 0x95, 0xd3,
		0x4c, 0xf2, 0x7a, 0xcb, 0x3e, 0xf4, 0x04, 0x0c, 0x86, 0xe0, 0x10, 0xba,
		0x4a, 0x59, 0xa9, 0x95, 0xc4, 0x47, 0xa1, 0xe5, 0x45, 0x1e, 0xe4, 0xef,
		0x8c, 0x4b, 0x84, 0x13, 0xd2, 0xbd, 0x9a, 0x02, 0xa5, 0xe0, 0x55, 0x2c,
		0xa0, 0x46, 0xc6, 0xa5, 0xc6, 0x46, 0x69, 0x48, 0xdc, 0x38, 0x61, 0xe9,
		0x5a, 0xab, 0x18, 0x5a, 0x60, 0x5e, 0x14, 0x4a, 0x60, 0x73, 0x16, 0x90,
		0x1e, 0x9c, 0x07, 0xca, 0xe3, 0x1d, 0xe5, 0x1a, 0xac, 0x81, 0x2c, 0x5b,
		0x2b, 0x8e, 0x9c, 0xbd, 0xa9, 0xb2, 0x18, 0x2c, 0x01, 0xbf, 0x3b, 0xe7,
		0xdb, 0xe0, 0x97, 0x11, 0x4b, 0x59, 0x2c, 0x94, 0x5b, 0xcb, 0x86, 0xfa,
		0x0e, 0x8a, 0x5a, 0x48, 0xa8, 0xb4, 0x9e, 0x64, 0x5b, 0x65, 0x9d, 0x33,
		0x51, 0xf1, 0x0c, 0xcc, 0x4e, 0xbc, 0x79, 0x8f, 0x8b, 0xe6, 0x98, 0x52,
		0xf4, 0xe9, 0x59, 0x1f, 0x48, 0xd7, 0x2d, 0xe9, 0x76, 0x5b, 0x76, 0x6f,
		0x87, 0x65, 0xb4, 0x73, 0xb1, 0xda, 0xd1, 0x9d, 0xa6, 0xf6, 0x8f, 0x86,
		0xbf, 0x46, 0x6a, 0xda, 0xfd, 0x66, 0xc2, 0x31, 0x44, 0x43, 0xe7, 0x78,
		0xae, 0x6f, 0xdc, 0xf6, 0x1b, 0xe5, 0xb0, 0x5c, 0x76, 0xdf, 0xf6, 0x10,
		0x7e, 0x75, 0x6e, 0xd4, 0x0d, 0xad, 0xd3, 0xe6, 0x55, 0x92, 0x8a, 0x07,
		0xfd, 0xec, 0xd5, 0x0c, 0x2e, 0x9e, 0x8d, 0x75, 0xc5, 0xee, 0x77, 0x01,
		0xb5, 0xe2, 0x90, 0xdf, 0xb9, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x2f,
		0x57, 0xa6, 0x1b, 0xa4, 0x03, 0x00, 0x00,
	},
		"sqlc/tmpl/schema.tmpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"sqlc/tmpl/fields.tmpl": sqlc_tmpl_fields_tmpl,
	"sqlc/tmpl/schema.tmpl": sqlc_tmpl_schema_tmpl,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"sqlc": &_bintree_t{nil, map[string]*_bintree_t{
		"tmpl": &_bintree_t{nil, map[string]*_bintree_t{
			"fields.tmpl": &_bintree_t{sqlc_tmpl_fields_tmpl, map[string]*_bintree_t{
			}},
			"schema.tmpl": &_bintree_t{sqlc_tmpl_schema_tmpl, map[string]*_bintree_t{
			}},
		}},
	}},
}}