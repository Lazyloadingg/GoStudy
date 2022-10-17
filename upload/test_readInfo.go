package main

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/mitchellh/mapstructure"
)

type IPAInfo struct {
	Version      string `json:"version"`     //版本号
	Build        string `json:"build"`       //build号
	DisplayName  string `json:"displayName"` //名称
	BundleID     string `json:"bundleID"`
	MinOSVersion string `json:"minOSVersion"`
}

func ReadInfoPList(path string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		fmt.Printf("\"etree读取失败\": %v\n", "etree读取失败")
		return
	}
	dict := doc.SelectElement("plist")
	fmt.Printf("dict: %v\n", dict.Tag)
	elements := dict.ChildElements()
	fmt.Printf("elements: %v\n", elements)
	element := elements[0]
	fmt.Printf("element: %v\n", element.Tag)

	dictchild := element.ChildElements()
	infoMap := make(map[string]interface{})
	for index, v := range dictchild {

		if v.Text() == "CFBundleDisplayName" {
			infoMap["displayName"] = dictchild[index+1].Text()
		} else if v.Text() == "CFBundleShortVersionString" && v.Tag == "key"{
			infoMap["version"] = dictchild[index+1].Text()
		} else if v.Text() == "CFBundleVersion" && v.Tag == "key"{
			infoMap["build"] = dictchild[index+1].Text()
		} else if v.Text() == "CFBundleIdentifier" && v.Tag == "key"{
			infoMap["bundleID"] = dictchild[index+1].Text()
		} else if v.Text() == "MinimumOSVersion" && v.Tag == "key"{
			infoMap["minOSVersion"] = dictchild[index+1].Text()
		}

		fmt.Printf("%v--child: %v--%v\n", index, v.Text(), v.Tag)
	}

	var ipainfo IPAInfo
	err := mapstructure.Decode(infoMap, &ipainfo)
	if err != nil {
		fmt.Printf("\"ipa信息系列化失败\": %v\n", "ipa信息系列化失败")
	}
	fmt.Printf("infoMap: %v\n", infoMap)
	fmt.Printf("ipainfo成功: %v\n", ipainfo)
}
