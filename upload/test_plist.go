package main

import (
	"os"

	"github.com/beevik/etree"
)

// 创建plist
func CreatePlist() {
	newdoc := etree.NewDocument()
	newdoc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	newdoc.CreateDirective(`DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd"`)
	plist := newdoc.CreateElement("plist")
	plist.CreateAttr("version", "1.0")
	newdoc.Indent(2)

	dict := plist.CreateElement("dict")
	key := dict.CreateElement("key")
	_ = key.CreateText("items")
	array := dict.CreateElement("array")
	item := array.CreateElement("dict")

	//assets
	assetsK := item.CreateElement("key")
	assetsK.CreateText("assets")
	assetArr := item.CreateElement("array")

	//pkg
	pkgDict := assetArr.CreateElement("dict")
	kindK := pkgDict.CreateElement("key")
	kindK.CreateText("kind")
	softPkg := pkgDict.CreateElement("string")
	softPkg.CreateText("software-package")

	urlK := pkgDict.CreateElement("key")
	urlK.CreateText("url")
	urlV := pkgDict.CreateElement("string")
	urlV.CreateText("https://sit-f01.hongxinshop.com/image/")

	//full-image
	fullDict := assetArr.CreateElement("dict")
	fullKindK := fullDict.CreateElement("key")
	fullKindK.CreateText("kind")
	fullKindV := fullDict.CreateElement("string")
	fullKindV.CreateText("full-size-image")

	fullShineK := fullDict.CreateElement("key")
	fullShineK.CreateText("needs-shine")
	fullDict.CreateElement("true")

	fullImgK := fullDict.CreateElement("key")
	fullImgK.CreateText("url")
	fullImgV := fullDict.CreateElement("string")
	fullImgV.CreateText("你的大图url")

	//display-image
	dispalayDict := assetArr.CreateElement("dict")
	displayKindK := dispalayDict.CreateElement("key")
	displayKindK.CreateText("kind")
	displayKindV := dispalayDict.CreateElement("string")
	displayKindV.CreateText("display-image")

	displayShineK := dispalayDict.CreateElement("key")
	displayShineK.CreateText("needs-shine")
	dispalayDict.CreateElement("true")

	displayImgK := dispalayDict.CreateElement("key")
	displayImgK.CreateText("url")
	displayImgV := dispalayDict.CreateElement("string")
	displayImgV.CreateText("你的小图url")

	//metadata
	metadataK := item.CreateElement("key")
	metadataK.CreateText("metadata")
	metadataDict := item.CreateElement("dict")
	createMetadata("baidu.com", "1.0.0", "四六级", metadataDict)
	newdoc.WriteTo(os.Stdout)
	newdoc.WriteToFile("./test.plist")
}

// 创建metadata
func createMetadata(bundleid, version, title string, metadata *etree.Element) {

	bundleidK := metadata.CreateElement("key")
	bundleidK.CreateText("bundle-identifier")
	bundleidV := metadata.CreateElement("string")
	bundleidV.CreateText(bundleid)

	versionK := metadata.CreateElement("key")
	versionK.CreateText("bundle-version")
	versionV := metadata.CreateElement("string")
	versionV.CreateText(version)

	kindK := metadata.CreateElement("key")
	kindK.CreateText("kind")
	kindV := metadata.CreateElement("string")
	kindV.CreateText("software")

	titleK := metadata.CreateElement("key")
	titleK.CreateText("title")
	titleV := metadata.CreateElement("string")
	titleV.CreateText(title)

}
