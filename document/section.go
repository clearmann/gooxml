// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"log"

	"github.com/clearmann/gooxml"
	"github.com/clearmann/gooxml/measurement"
	"github.com/clearmann/gooxml/schema/soo/ofc/sharedTypes"
	"github.com/clearmann/gooxml/schema/soo/wml"
)

// Section is the beginning of a new section.
type Section struct {
	d *Document
	x *wml.CT_SectPr
}

// 定义 PageNumbering 结构体，用于存储页码的起始值
type PageNumbering struct {
	Start int
}

// MakeSection constructs a new section.
func (s Section) SetPageNumbering(pn PageNumbering) {
	pageNumber := wml.NewCT_PageNumber()
	pageNumber.StartAttr = gooxml.Int64(int64(pn.Start))
	s.x.PgNumType = pageNumber
}

// X returns the internally wrapped *wml.CT_SectPr.
func (s Section) X() *wml.CT_SectPr {
	return s.x
}

// SetHeader sets a section header.
func (s Section) SetHeader(h Header, t wml.ST_HdrFtr) {
	hdrRef := wml.NewEG_HdrFtrReferences()
	s.x.EG_HdrFtrReferences = append(s.x.EG_HdrFtrReferences, hdrRef)
	hdrRef.HeaderReference = wml.NewCT_HdrFtrRef()
	hdrRef.HeaderReference.TypeAttr = t
	hdrID := s.d.docRels.FindRIDForN(h.Index(), gooxml.HeaderType)
	if hdrID == "" {
		log.Print("unable to determine header ID")
	}
	hdrRef.HeaderReference.IdAttr = hdrID
}

// SetFooter sets a section footer.
func (s Section) SetFooter(f Footer, t wml.ST_HdrFtr) {
	ftrRef := wml.NewEG_HdrFtrReferences()
	s.x.EG_HdrFtrReferences = append(s.x.EG_HdrFtrReferences, ftrRef)
	ftrRef.FooterReference = wml.NewCT_HdrFtrRef()
	ftrRef.FooterReference.TypeAttr = t
	hdrID := s.d.docRels.FindRIDForN(f.Index(), gooxml.FooterType)
	if hdrID == "" {
		log.Print("unable to determine footer ID")
	}
	ftrRef.FooterReference.IdAttr = hdrID
}

// SetPageMargins sets the page margins for a section
func (s Section) SetPageMargins(top, right, bottom, left, header, footer, gutter measurement.Distance) {

	margins := wml.NewCT_PageMar()
	margins.TopAttr.Int64 = gooxml.Int64(int64(top / measurement.Twips))
	margins.BottomAttr.Int64 = gooxml.Int64(int64(bottom / measurement.Twips))
	margins.RightAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(right / measurement.Twips))
	margins.LeftAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(left / measurement.Twips))
	margins.HeaderAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(header / measurement.Twips))
	margins.FooterAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(footer / measurement.Twips))
	margins.GutterAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(gutter / measurement.Twips))

	s.x.PgMar = margins
}

// 设置页面为A4纸张大小,如果出现分节问题,需要在每个分节中都调用
func (section Section) SetLandscapeA4PageSize() {
	sectPr := section.X()
	sectPr.PgSz = wml.NewCT_PageSz()
	sectPr.PgSz.OrientAttr = wml.ST_PageOrientationLandscape
	// A4：9 A5：11
	sectPr.PgSz.CodeAttr = gooxml.Int64(9)
	// 纸张宽高，A4横向长宽为21x29.7cm
	// 计算为厘米转磅数再乘以20
	// 长：21 * 28.35 * 20 = 11907
	// 宽：29.7 * 28.35 * 20 = 16839.9 ≈ 16840
	sectPr.PgSz.WAttr = &sharedTypes.ST_TwipsMeasure{ST_UnsignedDecimalNumber: gooxml.Uint64(11907)}
	sectPr.PgSz.HAttr = &sharedTypes.ST_TwipsMeasure{ST_UnsignedDecimalNumber: gooxml.Uint64(16839)}
	// 页边距只需要厘米转为磅数，即厘米x28.35
	section.SetPageMargins(72, 54, 72, 54, 42.55, 49.6, 0)
}
