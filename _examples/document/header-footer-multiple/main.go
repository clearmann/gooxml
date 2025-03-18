// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"fmt"

	"github.com/clearmann/gooxml/color"
	"github.com/clearmann/gooxml/common"
	"github.com/clearmann/gooxml/document"
	"github.com/clearmann/gooxml/measurement"
	"github.com/clearmann/gooxml/schema/soo/wml"
)

var lorem = `Lorem ipsum dolor sit amet, Vestibulum t反反复复反反复复凤飞飞发发发344444444444444444444444444444444444442GVvvvvvvvvvvvvvvvvvvv了解覅额黑白 问佛嗯empus sagittis elementum`

func main() {

	doc := document.New()

	var footer document.Footer
	//页脚,可内部设置字体大小样式,注意分节情况下进行继承调用
	setFooter := func() {
		footer = doc.AddFooter()
		para := footer.AddParagraph()
		para.Properties().SetAlignment(wml.ST_JcCenter)

		// 添加页码字段
		run := para.AddRun()
		run.Properties().SetFontFamily("Times New Roman")
		run.Properties().SetSize(10.5 * measurement.Point)
		run.AddField(document.FieldCurrentPage)
	}
	var hdr document.Header
	//页眉,可内部设置字体大小,字体,并且自带分割线与正文分割等
	setHeader := func(text string) {
		hdr = doc.AddHeader()
		para := hdr.AddParagraph()
		para.Properties().SetAlignment(wml.ST_JcCenter)
		run := para.AddRun()
		run.Properties().SetFontFamily("SimHei")
		run.Properties().SetSize(14 * measurement.Point)
		run.AddTab()
		run.AddText(text)
		//增加页眉下方的黑线
		borderPara := hdr.AddParagraph()
		borderPara.Properties().SetStyle("Normal")
		borderPara.Properties().SetBottomBorder(wml.ST_BorderThick, color.Black, 10*measurement.Point)

		// 添加一个空行
		emptyPara := hdr.AddParagraph()
		emptyPara.Properties().SetStyle("Normal")
	}

	//分节,并且继承了不同页眉和相同页脚
	setSection := func() {
		para := doc.AddParagraph()
		section := para.Properties().AddSection(wml.ST_SectionMarkNextPage)
		section.SetHeader(hdr, wml.ST_HdrFtrDefault)
		section.SetFooter(footer, wml.ST_HdrFtrDefault)
		section.SetLandscapeA4PageSize()
	}

	//插入图片文本(宽为6英寸并自适应缩放), 并添加编号,
	imageImage := func(imgPath string, imgIndex int) {
		para := doc.AddParagraph()
		para.Properties().SetAlignment(wml.ST_JcCenter)
		run := para.AddRun()
		// 加载本地图片
		img, err := common.ImageFromFile(imgPath)
		if err != nil {
			fmt.Printf("Error loading image: %v\n", err)
			return
		}
		// 插入图片到段落
		imgRef, err := doc.AddImage(img)
		if err != nil {
			fmt.Printf("Error adding image to document: %v\n", err)
			return
		}
		inl, err := run.AddDrawingInline(imgRef)
		if err != nil {
			fmt.Printf("Error adding image: %v\n", err)
			return
		}
		// 设置图片大小，长为6英寸，宽自适应
		width := measurement.Distance(6 * measurement.Inch)
		height := width * measurement.Distance(img.Size.Y) / measurement.Distance(img.Size.X)
		inl.SetSize(width, height)
		// 添加图片编号
		if imgIndex != 0 {
			para = doc.AddParagraph()
			para.Properties().SetAlignment(wml.ST_JcCenter)
			run = para.AddRun()
			run.Properties().SetSize(14.0)
			run.AddText(fmt.Sprintf("图 %d", imgIndex))
		}
	}
	setFooter()
	document.AddIndentedMultilineText(doc, "宋体", lorem, 14)
	setHeader("说  明  书  摘  要")
	setSection()
	document.AddIndentedMultilineText(doc, "宋体", lorem, 14)
	setHeader("摘  要  附  图")
	setSection()
	len := 3
	if len > 0 {
		for i := 0; i < len; i++ {
			imageImage(fmt.Sprintf("%d号.png", i+1), i+1)
			setHeader("说  明  书  附  图")
			if i != len-1 {
				setSection()
			}
		}
	} else {
		setHeader("说  明  书  附  图")
	}

	//这三个方法做为页眉页脚以及A4大小的收尾方法,很重要!
	doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)
	doc.BodySection().SetLandscapeA4PageSize()
	doc.BodySection().SetFooter(footer, wml.ST_HdrFtrDefault)

	doc.SaveToFile("simple2.docx")

}
