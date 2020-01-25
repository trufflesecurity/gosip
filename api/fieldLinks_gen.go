// Package api :: This is auto generated file, do not edit manually
package api

// Conf receives custom request config definition, e.g. custom headers, custom OData mod
func (fieldLinks *FieldLinks) Conf(config *RequestConfig) *FieldLinks {
	fieldLinks.config = config
	return fieldLinks
}

// Select adds $select OData modifier
func (fieldLinks *FieldLinks) Select(oDataSelect string) *FieldLinks {
	fieldLinks.modifiers.AddSelect(oDataSelect)
	return fieldLinks
}

// Filter adds $filter OData modifier
func (fieldLinks *FieldLinks) Filter(oDataFilter string) *FieldLinks {
	fieldLinks.modifiers.AddFilter(oDataFilter)
	return fieldLinks
}

// Top adds $top OData modifier
func (fieldLinks *FieldLinks) Top(oDataTop int) *FieldLinks {
	fieldLinks.modifiers.AddTop(oDataTop)
	return fieldLinks
}