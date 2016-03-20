# Go-Sprockets <img src="logo.png" alt="logo" width="80px">
Sprockets' asset pipeline for Golang

## About Sprockets
Sprockets is a Ruby library for compiling and serving web assets.
It features declarative dependency management for JavaScript and CSS
assets, as well as a powerful preprocessor pipeline that allows you to
write assets in languages like CoffeeScript, Sass and SCSS.
And now it's in Golang!

## Installation

Get the package:

```bash
$ go get github.com/znly/go-sprockets
```

## Warning.
it is using [go-libsass](http://github.com/wellington/go-libsass). it compiles C/C++ libsass on every build. Except if you install the packages.

it is using [duktape](https://github.com/olebedev/go-duktape). it compiles C duktape on every build. Except if you install the packages.

## Compilation Pipeline
* Get the asset extension info
    * use the last \\..* as the extension
    * if no extension info is found, use the default one

* Search for the asset using the extenstion info
    * search for the asset in each path defined in the extension info
        * for each path try current extension then alternate extension
        * Stop and return the first result, changing the extension info to the new extension of the asset if needed
* Read the asset and all it s requirement

    * Read the raw asset
        * read the whole asset file content
        * apply all the content contentTreatment from the extension info

    * Find requirements
        * use the require pattern Head to retrieve the HEADER
        * apply all the header ContentTreatment from the extension info if HEADER is found
        * use the require pattern Requires to retrieve each directive lines in the HEADER
        * apply the filecompiler of the extension info
        * search for the requirements, read them and find their own requirements.

* Bundle and compile
    * resolve the dependency graph and build the full content of the asset
    * apply the bundlecompiler of the extension info
    * apply all the post compile ContentTreatment

* Return the result

## Examples:
If you want a simple example of extensions look at the [NewWithDefault Function](default.go)

You can also look at this [project](https://github.com/znly/go-dashing)

## Doc:
You can find the doc on godoc.org  [![GoDoc](https://godoc.org/github.com/znly/go-sprockets?status.png)](https://godoc.org/github.com/znly/go-sprockets)

## Todo:

* find extension info by iterating on each extension info name and check if the asset is ending by it (sort by size .min.js > .js)
* implement [Index files are proxies for folders](https://github.com/rails/sprockets#index-files-are-proxies-for-folders)
* implement [require self](https://github.com/rails/sprockets#the-require_self-directive)
* implement [depend_on](https://github.com/rails/sprockets#the-depend_on-directive)
* implement [depend_on_asset](https://github.com/rails/sprockets#the-depend_on_asset-directive)
* implement [stub](https://github.com/rails/sprockets#the-stub-directive)
* Make the production mode build assets in the public path
* Write Tests
* Make the cache faster!
