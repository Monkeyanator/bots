// Code generated by go-bindata.
// sources:
// base.html
// error.html
// notfound.html
// primary.html
// DO NOT EDIT!

package layout

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
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

var _baseHtml = []byte(`<!DOCTYPE html>
<html lang="en" itemscope itemtype="https://schema.org/WebPage">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <meta name="theme-color" content="#466BB0"/>

        <meta name="description" content="{{ .Description }}">

		{{ if .Title }}
	        <meta name="title" content="{{ .Title }}">
			<title>Eng / {{ .Title }}</title>
		{{ else }}
	        <meta name="title" content="Eng Dashboard">
			<title>Eng Dashboard</title>
		{{ end }}

        <!-- Favicons: generated from img/istio-whitelogo-bluebackground-framed.svg by http://cthedot.de/icongen -->
        <link rel="shortcut icon" href="/favicons/favicon.ico" >
        <link rel="apple-touch-icon" href="/favicons/apple-touch-icon-180x180.png" sizes="180x180">
        <link rel="icon" type="image/png" href="/favicons/favicon-16x16.png" sizes="16x16">
        <link rel="icon" type="image/png" href="/favicons/favicon-32x32.png" sizes="32x32">
        <link rel="icon" type="image/png" href="/favicons/android-36x36.png" sizes="36x36">
        <link rel="icon" type="image/png" href="/favicons/android-48x48.png" sizes="48x48">
        <link rel="icon" type="image/png" href="/favicons/android-72x72.png" sizes="72x72">
        <link rel="icon" type="image/png" href="/favicons/android-96x96.png" sizes="96xW96">
        <link rel="icon" type="image/png" href="/favicons/android-144x144.png" sizes="144x144">
        <link rel="icon" type="image/png" href="/favicons/android-192x192.png" sizes="192x192">

        <!-- app manifests -->
        <link rel="manifest" href="/manifest.json">
        <meta name="apple-mobile-web-app-title" content="Istio Eng Dashboard">
        <meta name="application-name" content="Istio Eng Dashboard">

        <!-- style sheets -->
        <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Work+Sans:400|Chivo:400|Work+Sans:500,300,600,300italic,400italic`+"`"+` +
	`+"`"+`,500italic,600italic|Chivo:500,300,600,300italic,400italic,500italic,600italic">
        <link rel="stylesheet" href="/css/all.css">
    </head>

    <body>
        <!-- set the color theme as soon as possible -->
        <script src="/js/themes_init.min.js"></script>

		<!-- libraries we use -->
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" defer></script>

        <!-- our own stuff -->
        <script src="/js/all.min.js" defer></script>
        <script src="/js/fitty.min.js" defer></script>

        {{ template "main" . }}
    </body>
</html>
`)

func baseHtmlBytes() ([]byte, error) {
	return _baseHtml, nil
}

func baseHtml() (*asset, error) {
	bytes, err := baseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "base.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _errorHtml = []byte(`{{ define "main" }}
{{ template "header" . }}

<main class="notfound" role="main">
    <svg class="icon">
        <use xlink:href="/icons/icons.svg#exclamation-mark"/>
    </svg>

    <div class="error">
        We're sorry, the server has experienced an error.
    </div>

    <div class="explanation">
        {{ .Content }}
    </div>
</main>
{{ end }}
`)

func errorHtmlBytes() ([]byte, error) {
	return _errorHtml, nil
}

func errorHtml() (*asset, error) {
	bytes, err := errorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "error.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _notfoundHtml = []byte(`{{ define "main" }}
{{ template "header" . }}

<main class="notfound" role="main">
    <svg class="icon">
        <use xlink:href="/icons/icons.svg#exclamation-mark"/>
    </svg>

    <div class="error">
        We're sorry, the page you requested cannot be found
    </div>

    <div class="explanation">
        The URL may be misspelled or the page you're looking for is no longer available
    </div>
</main>
{{ end }}
`)

func notfoundHtmlBytes() ([]byte, error) {
	return _notfoundHtml, nil
}

func notfoundHtml() (*asset, error) {
	bytes, err := notfoundHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "notfound.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _primaryHtml = []byte(`{{ define "main" }}
{{ template "header" . }}

<main class="primary notoc">
	<div id="sidebar-container" class="sidebar-container sidebar-offcanvas">
        {{ template "sidebar" . }}
    </div>

	<div class="article-container">
        {{ if .Control }}
            <nav aria-label="Breadcrumb">
                {{ .Control }}
            </nav>
        {{ end }}

        <article aria-labelledby="title">
            <div class="title-area">
                <div>
                    <h1 id="title">
                        {{- .Title -}}
                    </h1>

					{{ .Content }}
				</div>
			</div>
		</article>
    </div>
</main>
{{ end }}`)

func primaryHtmlBytes() ([]byte, error) {
	return _primaryHtml, nil
}

func primaryHtml() (*asset, error) {
	bytes, err := primaryHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "primary.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"base.html": baseHtml,
	"error.html": errorHtml,
	"notfound.html": notfoundHtml,
	"primary.html": primaryHtml,
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
	"base.html": &bintree{baseHtml, map[string]*bintree{}},
	"error.html": &bintree{errorHtml, map[string]*bintree{}},
	"notfound.html": &bintree{notfoundHtml, map[string]*bintree{}},
	"primary.html": &bintree{primaryHtml, map[string]*bintree{}},
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

