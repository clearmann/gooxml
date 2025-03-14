// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"fmt"
	"github.com/clearmann/gooxml/color"
	"github.com/clearmann/gooxml/document"
	"github.com/clearmann/gooxml/schema/soo/wml"
)

func main() {
	doc := document.New()
	para := doc.AddParagraph()
	run := para.AddRun()
	para.SetStyle("Title")
	run.AddText("Simple Document\n Formatting")
	run.AddBreak()
	run.AddMultiLineText("Simple Document\n Formatting")
	run.AddBreak()
	run.AddText("Simple Document\n Formatting")

	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	run = para.AddRun()
	run.AddText("一级标题")

	para = doc.AddParagraph()
	para.SetStyle("Heading2")
	run = para.AddRun()
	run.AddText("二级标题")

	para = doc.AddParagraph()
	para.SetStyle("Heading3")
	run = para.AddRun()
	run.AddText("三级标题")

	para = doc.AddParagraph()

	text := "1、一种数据处理方法，其特征在于，包括步骤一获取输入数据，步骤二对所述输入数据进行预处理，步骤三将预处理后的数据输入至训练好的9流模型中进行分析，步骤四输出分析结果。\n\n2、根据权利要求1所述的数据处理方法，其特征在于，步骤一中获取输入数据的方式包括从数据库中提取数据或通过传感器实时采集数据。\n\n3、根据权利要求1所述的数据处理方法，其特征在于，步骤二中对所述输入数据进行预处理的方式包括对数据进行归一化处理，归一化处理的计算公式为$X_{normalized} = \\\\frac{X - X_{min}}{X_{max} - X_{min}}$，其中$X$为原始数据，$X_{min}$为数据最小值，$X_{max}$为数据最大值。\n\n4、根据权利要求1所述的数据处理方法，其特征在于，步骤三中所述9流模型的训练数据量在一个范围内，训练数据的来源包括历史数据和实时数据。\n\n5、根据权利要求1所述的数据处理方法，其特征在于，步骤四中输出分析结果的方式包括将结果存储在数据库中或通过可视化界面展示。"

	document.AddIndentedMultilineText(doc, "宋体", text, 14)

	document.AddIndentedMultilineText(doc, "宋体", text, 12)
	document.AddIndentedMultilineText(doc, "宋体", text, 8)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("Courier")
	run.Properties().SetSize(15)
	run.Properties().SetColor(color.Red)
	run.AddText("Multiple runs with different formatting can exist in the same paragraph. ")

	run = para.AddRun()
	run.AddText("Adding breaks \nto a run will insert line breaks after the run. ")
	run.AddBreak()
	run.AddBreak()

	run = createParaRun(doc, "Runs support styling options:")

	run = createParaRun(doc, "small caps")
	run.Properties().SetSmallCaps(true)

	run = createParaRun(doc, "strike")
	run.Properties().SetStrikeThrough(true)

	run = createParaRun(doc, "double strike")
	run.Properties().SetDoubleStrikeThrough(true)

	run = createParaRun(doc, "outline")
	run.Properties().SetOutline(true)

	run = createParaRun(doc, "emboss")
	run.Properties().SetEmboss(true)

	run = createParaRun(doc, "shadow")
	run.Properties().SetShadow(true)

	run = createParaRun(doc, "imprint")
	run.Properties().SetImprint(true)

	run = createParaRun(doc, "highlighting")
	run.Properties().SetHighlight(wml.ST_HighlightColorYellow)

	run = createParaRun(doc, "underline")
	run.Properties().SetUnderline(wml.ST_UnderlineWavyDouble, color.Red)

	run = createParaRun(doc, "text effects")
	run.Properties().SetEffect(wml.ST_TextEffectAntsRed)

	nd := doc.Numbering.Definitions()[0]

	for i := 1; i < 5; i++ {
		p := doc.AddParagraph()
		p.SetNumberingLevel(i - 1)
		p.SetNumberingDefinition(nd)
		run := p.AddRun()
		run.AddText(fmt.Sprintf("Level %d", i))
	}
	doc.SaveToFile("simple2.docx")
}

func createParaRun(doc *document.Document, s string) document.Run {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(s)
	return run
}
