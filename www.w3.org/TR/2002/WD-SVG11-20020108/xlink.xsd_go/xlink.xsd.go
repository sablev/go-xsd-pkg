//	Auto-generated by the "go-xsd" package located at:
//		github.com/sablev/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/TR/2002/WD-SVG11-20020108/xlink.xsd
package go_Xlink

import (
	xsdt "github.com/sablev/go-xsd/types"
)

type TxsdType xsdt.String

//	Returns true if the value of this enumerated TxsdType is "locator".
func (me TxsdType) IsLocator() bool { return me == "locator" }

//	Since TxsdType is just a simple String type, this merely returns the current string value.
func (me TxsdType) String() string { return xsdt.String(me).String() }

//	Returns true if the value of this enumerated TxsdType is "arc".
func (me TxsdType) IsArc() bool { return me == "arc" }

//	Since TxsdType is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdType) Set(s string) { (*xsdt.String)(me).Set(s) }

//	Returns true if the value of this enumerated TxsdType is "extended".
func (me TxsdType) IsExtended() bool { return me == "extended" }

//	Returns true if the value of this enumerated TxsdType is "simple".
func (me TxsdType) IsSimple() bool { return me == "simple" }

//	This convenience method just performs a simple type conversion to TxsdType's alias type xsdt.String.
func (me TxsdType) ToXsdtString() xsdt.String { return xsdt.String(me) }

type XsdGoPkgHasAttr_Type struct {
	Type TxsdType `xml:"http://www.w3.org/1999/xlink type,attr"`
}

//	Returns the default value for Type -- "simple"
func (me XsdGoPkgHasAttr_Type) TypeDefault() TxsdType { return TxsdType("simple") }

type XsdGoPkgHasAttr_Href struct {
	Href xsdt.AnyURI `xml:"http://www.w3.org/1999/xlink href,attr"`
}

type XsdGoPkgHasAttr_Role struct {
	Role xsdt.String `xml:"http://www.w3.org/1999/xlink role,attr"`
}

type XsdGoPkgHasAttr_Arcrole struct {
	Arcrole xsdt.String `xml:"http://www.w3.org/1999/xlink arcrole,attr"`
}

type XsdGoPkgHasAttr_Title struct {
	Title xsdt.String `xml:"http://www.w3.org/1999/xlink title,attr"`
}

type TxsdShow xsdt.String

//	Returns true if the value of this enumerated TxsdShow is "other".
func (me TxsdShow) IsOther() bool { return me == "other" }

//	Since TxsdShow is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdShow) Set(s string) { (*xsdt.String)(me).Set(s) }

//	Returns true if the value of this enumerated TxsdShow is "replace".
func (me TxsdShow) IsReplace() bool { return me == "replace" }

//	Since TxsdShow is just a simple String type, this merely returns the current string value.
func (me TxsdShow) String() string { return xsdt.String(me).String() }

//	Returns true if the value of this enumerated TxsdShow is "none".
func (me TxsdShow) IsNone() bool { return me == "none" }

//	Returns true if the value of this enumerated TxsdShow is "new".
func (me TxsdShow) IsNew() bool { return me == "new" }

//	Returns true if the value of this enumerated TxsdShow is "embed".
func (me TxsdShow) IsEmbed() bool { return me == "embed" }

//	This convenience method just performs a simple type conversion to TxsdShow's alias type xsdt.String.
func (me TxsdShow) ToXsdtString() xsdt.String { return xsdt.String(me) }

type XsdGoPkgHasAttr_Show struct {
	Show TxsdShow `xml:"http://www.w3.org/1999/xlink show,attr"`
}

//	Returns the default value for Show -- "embed"
func (me XsdGoPkgHasAttr_Show) ShowDefault() TxsdShow { return TxsdShow("embed") }

type TxsdActuate xsdt.String

//	This convenience method just performs a simple type conversion to TxsdActuate's alias type xsdt.String.
func (me TxsdActuate) ToXsdtString() xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TxsdActuate is "other".
func (me TxsdActuate) IsOther() bool { return me == "other" }

//	Since TxsdActuate is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdActuate) Set(s string) { (*xsdt.String)(me).Set(s) }

//	Returns true if the value of this enumerated TxsdActuate is "onLoad".
func (me TxsdActuate) IsOnLoad() bool { return me == "onLoad" }

//	Returns true if the value of this enumerated TxsdActuate is "onRequest".
func (me TxsdActuate) IsOnRequest() bool { return me == "onRequest" }

//	Returns true if the value of this enumerated TxsdActuate is "none".
func (me TxsdActuate) IsNone() bool { return me == "none" }

//	Since TxsdActuate is just a simple String type, this merely returns the current string value.
func (me TxsdActuate) String() string { return xsdt.String(me).String() }

type XsdGoPkgHasAttr_Actuate struct {
	Actuate TxsdActuate `xml:"http://www.w3.org/1999/xlink actuate,attr"`
}

//	Returns the default value for Actuate -- "onLoad"
func (me XsdGoPkgHasAttr_Actuate) ActuateDefault() TxsdActuate { return TxsdActuate("onLoad") }

type XsdGoPkgHasAttr_From struct {
	From xsdt.String `xml:"http://www.w3.org/1999/xlink from,attr"`
}

type XsdGoPkgHasAttr_To struct {
	To xsdt.String `xml:"http://www.w3.org/1999/xlink to,attr"`
}
