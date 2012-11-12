//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/Math/XMLSchema/mathml2/content/relations.xsd
package gopkg_WwwW3OrgMathXMLSchemaMathml2Mathml2Xsd

//	This is an XML Schema module for the relational operators of content MathML.
//	Author: Stéphane Dalmas, INRIA.
import (
)

type TRelationsType struct {
	XsdGoPkgHasAtts_CommonAttrib

	XsdGoPkgHasAtts_DefinitionAttrib

}

type XsdGoPkgHasElem_Eq struct {
	Eq *TRelationsType `xml:"http://www.w3.org/1998/Math/MathML eq"`

}

type XsdGoPkgHasElems_Eq struct {
	Eqs []*TRelationsType `xml:"http://www.w3.org/1998/Math/MathML eq"`

}

type XsdGoPkgHasElems_Neq struct {
	Neqs []*TRelationsType `xml:"http://www.w3.org/1998/Math/MathML neq"`

}

type XsdGoPkgHasElem_Neq struct {
	Neq *TRelationsType `xml:"http://www.w3.org/1998/Math/MathML neq"`

}

type XsdGoPkgHasElem_Leq struct {
	Leq *TRelationsType `xml:"http://www.w3.org/1998/Math/MathML leq"`

}

type XsdGoPkgHasElems_Leq struct {
	Leqs []*TRelationsType `xml:"http://www.w3.org/1998/Math/MathML leq"`

}

type XsdGoPkgHasElem_Lt struct {
	Lt *TRelationsType `xml:"http://www.w3.org/1998/Math/MathML lt"`

}

type XsdGoPkgHasElems_Lt struct {
	Lts []*TRelationsType `xml:"http://www.w3.org/1998/Math/MathML lt"`

}

type XsdGoPkgHasElem_Geq struct {
	Geq *TRelationsType `xml:"http://www.w3.org/1998/Math/MathML geq"`

}

type XsdGoPkgHasElems_Geq struct {
	Geqs []*TRelationsType `xml:"http://www.w3.org/1998/Math/MathML geq"`

}

type XsdGoPkgHasElems_Gt struct {
	Gts []*TRelationsType `xml:"http://www.w3.org/1998/Math/MathML gt"`

}

type XsdGoPkgHasElem_Gt struct {
	Gt *TRelationsType `xml:"http://www.w3.org/1998/Math/MathML gt"`

}

type XsdGoPkgHasElem_Equivalent struct {
	Equivalent *TRelationsType `xml:"http://www.w3.org/1998/Math/MathML equivalent"`

}

type XsdGoPkgHasElems_Equivalent struct {
	Equivalents []*TRelationsType `xml:"http://www.w3.org/1998/Math/MathML equivalent"`

}

type XsdGoPkgHasElems_Approx struct {
	Approxs []*TRelationsType `xml:"http://www.w3.org/1998/Math/MathML approx"`

}

type XsdGoPkgHasElem_Approx struct {
	Approx *TRelationsType `xml:"http://www.w3.org/1998/Math/MathML approx"`

}

type XsdGoPkgHasElems_Factorof struct {
	Factorofs []*TRelationsType `xml:"http://www.w3.org/1998/Math/MathML factorof"`

}

type XsdGoPkgHasElem_Factorof struct {
	Factorof *TRelationsType `xml:"http://www.w3.org/1998/Math/MathML factorof"`

}

type XsdGoPkgHasGroup_ContentRelationsClass struct {
	XsdGoPkgHasElem_Eq

	XsdGoPkgHasElem_Neq

	XsdGoPkgHasElem_Leq

	XsdGoPkgHasElem_Lt

	XsdGoPkgHasElem_Geq

	XsdGoPkgHasElem_Gt

	XsdGoPkgHasElem_Equivalent

	XsdGoPkgHasElem_Approx

	XsdGoPkgHasElem_Factorof

}
