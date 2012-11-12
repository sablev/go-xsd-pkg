//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/Math/XMLSchema/mathml2/presentation/error.xsd
package gopkg_WwwW3OrgMathXMLSchemaMathml2Mathml2Xsd

//	This is the XML Schema module for the MathML "merror" element.
//	Author: Stéphane Dalmas, INRIA.
import (
)

type XsdGoPkgHasAtts_MerrorAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

}

type TmerrorType struct {
	XsdGoPkgHasGroup_MerrorContent

	XsdGoPkgHasAtts_MerrorAttlist

}

type XsdGoPkgHasElem_Merror struct {
	Merror *TmerrorType `xml:"http://www.w3.org/1998/Math/MathML merror"`

}

type XsdGoPkgHasElems_Merror struct {
	Merrors []*TmerrorType `xml:"http://www.w3.org/1998/Math/MathML merror"`

}

type XsdGoPkgHasGroup_MerrorContent struct {
	XsdGoPkgHasGroup_PresentationExprClass

}
