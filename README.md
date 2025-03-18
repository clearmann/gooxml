**Notice** Special thanks to the original author for the open source library:
https://github.com/carmel/gooxml. As the original repository is not maintained , 
can not repair the original repository by submitting the code in the way of problems , 
so I used the original repository of the code , for secondary development 、 repairfeat some of the possible bugs 
、add some needed features.

特别感谢原作者提供的开源库：https://github.com/carmel/gooxml 由于原始仓库不在维护,无法通过提交代码的
方式修复原始仓库中存在的问题，因此我使用了原始仓库的代码，进行二次开发、修复一些可能存在的bug、增加一些需要的功能。



**gooxml** is a library for creation of Office Open XML documents (.docx, .xlsx
and .pptx).  It's goal is to be the most compatible and highest performance Go
library for creation and editing of docx/xlsx/pptx files.

Requires **go1.8+**, builds are tested with 1.8, 1.9 and tip.

[![Build Status](https://travis-ci.org/baliance/gooxml.svg?branch=master)](https://travis-ci.org/baliance/gooxml)
[![GitHub (pre-)release](https://img.shields.io/github/release/baliance/gooxml/all.svg)](https://github.com/clearmann/gooxml/releases)
[![License: AGPL v3](https://img.shields.io/badge/License-Dual%20AGPL%20v3/Commercial-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
[![GoDoc](https://godoc.org/baliance.com/gooxml?status.svg)](https://godoc.org/baliance.com/gooxml)
[![go 1.8+](https://img.shields.io/badge/go-1.8%2B-blue.svg)](http://golang.org)

![https://baliance.com/gooxml/](./_examples/preview.png "gooxml")

## Status ##

- Documents (docx) [Word]
	- Read/Write/Edit
	- Formatting
	- Images
	- Tables
- Spreadsheets (xlsx) [Excel]
 	- Read/Write/Edit
 	- Cell formatting including conditional formatting
	- Cell validation (drop down combobox, rules, etc.)
    - Retrieve cell values as formatted by Excel (e.g. retrieve a date or number as displayed in Excel)
 	- Formula Evaluation (100+ functions supported currently, more will be added as required)
 	- Embedded Images
 	- All chart types
- PowerPoint (pptx) [PowerPoint]
	- Creation from templates
	- Textboxes/shapes


## Performance ##

There has been a great deal of interest in performance numbers for spreadsheet
creation/reading lately, so here are gooxml numbers for this
[benchmark](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/lots-of-rows)
which creates a sheet with 30k rows, each with 100 columns.

    creating 30000 rows * 100 cells took 3.92506863s
    saving took 89ns
    reading took 9.522383048s

Creation is fairly fast, saving is very quick due to no reflection usage, and
reading is a bit slower. The downside is that the binary is large (33MB) as it
contains generated structs, serialization and deserialization code for all of
DOCX/XLSX/PPTX.

## Installation ##

    go get github.com/clearmann/gooxml

## Docx 文档示例 ##

- [简单文本格式](https://github.com/clearmann/gooxml/tree/master/_examples/document/simple) 文本字体颜色、大小、高亮显示等
- [自动生成目录](https://github.com/clearmann/gooxml/tree/master/_examples/document/toc) 创建文档标题，并基于标题自动生成目录
- [图像处理](https://github.com/clearmann/gooxml/tree/master/_examples/document/image) 在页面上绝对定位图像，支持不同文字环绕方式
- [页眉页脚](https://github.com/clearmann/gooxml/tree/master/_examples/document/header-footer) 创建包含页码的页眉和页脚
- [多重页眉+页脚+分割线+页面设置为A4+图片等比例处理](https://github.com/clearmann/gooxml/tree/master/_examples/document/header-footer-multiple) 根据文档章节使用不同的页眉页脚和黑色分割线以及页面大小设置
- [内联表格](https://github.com/clearmann/gooxml/tree/master/_examples/document/tables) 添加带边框和不带边框的表格
- [使用现有Word文档作为模板](https://github.com/clearmann/gooxml/tree/master/_examples/document/use-template) 打开文档作为模板以复用其中的样式
- [填写表单字段](https://github.com/clearmann/gooxml/tree/master/_examples/document/fill-out-form) 打开包含嵌入式表单字段的文档，填写字段并将结果另存为已填写的新表单
- [编辑现有文档](https://github.com/clearmann/gooxml/tree/master/_examples/document/edit-document) 打开现有文档并替换/删除文本，同时保持原有格式不变

## Excel 文档示例 ##
- [简单excel表格](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/simple) 包含少量单元格的简单工作表
- [命名单元格](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/named-cells) 行列和单元格的不同引用方式
- [数字/日期/时间格式](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/number-date-time-formats) 设置多种数字/日期/时间格式的单元格
- [折线图](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/line-chart)/[3D折线图](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/line-chart-3d) 折线图绘制示例
- [柱状图](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/bar-chart) 柱状图绘制示例
- [多图表布局](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/multiple-charts) 单工作表中创建多个图表
- [命名单元格区域](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/named-ranges) 定义单元格区域名称
- [合并单元格](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/merged) 合并与取消合并单元格
- [条件格式](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/conditional-formatting) 条件格式设置（样式/渐变色/图标集/数据条）
- [综合应用](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/complex) 多图表、自动筛选与条件格式的组合应用
- [单元格边框](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/borders) 单个单元格边框和单元格区域矩形边框设置
- [数据验证](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/validation) 包含组合框下拉列表的数据验证
- [冻结行列](https://github.com/clearmann/gooxml/tree/master/_examples/spreadsheet/freeze-rows-cols) 冻结表头行和列的视图设置

## Presentation Examples ##

- [Simple Text Boxes](https://github.com/clearmann/gooxml/tree/master/_examples/presentation/simple) Simple text boxes and shapes
- [Images](https://github.com/clearmann/gooxml/tree/master/_examples/presentation/image) Simple image insertion
- [Template](https://github.com/clearmann/gooxml/tree/master/_examples/presentation/use-template/simple) Creating a presentation from a template

## Raw Types ##

The OOXML specification is large and creating a friendly API to cover the entire
specification is a very time consuming endeavor.  This library attempts to
provide an easy to use API for common use cases in creating OOXML documents
while allowing users to fall back to raw document manipulation should the
library's API not cover a specific use case.

The raw XML based types reside in the ```schema/``` directory. These types are
accessible from the wrapper types via a ```X()``` method that returns the raw
type. 

For example, the library currently doesn't have an API for setting a document
background color. However it's easy to do manually via editing the
```CT_Background``` element of the document.

    dox := document.New()
    doc.X().Background = wordprocessingml.NewCT_Background()
	doc.X().Background.ColorAttr = &wordprocessingml.ST_HexColor{}
	doc.X().Background.ColorAttr.ST_HexColorRGB = color.RGB(50, 50, 50).AsRGBString()

### Contribution guidelines ###

[![CLA assistant](https://cla-assistant.io/readme/badge/baliance/gooxml)](https://cla-assistant.io/baliance/gooxml)

All contributors are must sign a contributor license agreement before their code
will be reviewed and merged.


### Licensing ###

This library is offered under a dual license. It is freely available for use
under the terms of AGPLv3. If you would like to use this library for a closed
source project, please contact sales@unidoc.io.

There are no differences in functionality between the open source and commercial 
versions. You are encouraged to use the open source version to evaluate the library
before purchasing a commercial license.

