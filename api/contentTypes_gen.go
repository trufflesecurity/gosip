// Code generated by `ggen -ent ContentTypes -item ContentType -conf -coll -mods Select,Expand,Filter,Top,Skip,OrderBy -helpers Data,Normalized,Pagination`; DO NOT EDIT.

package api

import (
	"fmt"
)

// Conf receives custom request config definition, e.g. custom headers, custom OData mod
func (contentTypes *ContentTypes) Conf(config *RequestConfig) *ContentTypes {
	contentTypes.config = config
	return contentTypes
}

// Select adds $select OData modifier
func (contentTypes *ContentTypes) Select(oDataSelect string) *ContentTypes {
	contentTypes.modifiers.AddSelect(oDataSelect)
	return contentTypes
}

// Expand adds $expand OData modifier
func (contentTypes *ContentTypes) Expand(oDataExpand string) *ContentTypes {
	contentTypes.modifiers.AddExpand(oDataExpand)
	return contentTypes
}

// Filter adds $filter OData modifier
func (contentTypes *ContentTypes) Filter(oDataFilter string) *ContentTypes {
	contentTypes.modifiers.AddFilter(oDataFilter)
	return contentTypes
}

// Top adds $top OData modifier
func (contentTypes *ContentTypes) Top(oDataTop int) *ContentTypes {
	contentTypes.modifiers.AddTop(oDataTop)
	return contentTypes
}

// Skip adds $skiptoken OData modifier
func (contentTypes *ContentTypes) Skip(skipToken string) *ContentTypes {
	contentTypes.modifiers.AddSkip(skipToken)
	return contentTypes
}

// OrderBy adds $orderby OData modifier
func (contentTypes *ContentTypes) OrderBy(oDataOrderBy string, ascending bool) *ContentTypes {
	contentTypes.modifiers.AddOrderBy(oDataOrderBy, ascending)
	return contentTypes
}

/* Response helpers */

// Data response helper
func (contentTypesResp *ContentTypesResp) Data() []ContentTypeResp {
	collection, _ := normalizeODataCollection(*contentTypesResp)
	contentTypes := []ContentTypeResp{}
	for _, item := range collection {
		contentTypes = append(contentTypes, ContentTypeResp(item))
	}
	return contentTypes
}

// Normalized returns normalized body
func (contentTypesResp *ContentTypesResp) Normalized() []byte {
	normalized, _ := NormalizeODataCollection(*contentTypesResp)
	return normalized
}

/* Pagination helpers */

// ContentTypesPage - paged items
type ContentTypesPage struct {
	Items       ContentTypesResp
	HasNextPage func() bool
	GetNextPage func() (*ContentTypesPage, error)
}

// GetPaged gets Paged Items collection
func (contentTypes *ContentTypes) GetPaged() (*ContentTypesPage, error) {
	data, err := contentTypes.Get()
	if err != nil {
		return nil, err
	}
	res := &ContentTypesPage{
		Items: data,
		HasNextPage: func() bool {
			return data.HasNextPage()
		},
		GetNextPage: func() (*ContentTypesPage, error) {
			nextURL := data.NextPageURL()
			if nextURL == "" {
				return nil, fmt.Errorf("unable to get next page")
			}
			return NewContentTypes(contentTypes.client, nextURL, contentTypes.config).GetPaged()
		},
	}
	return res, nil
}

// NextPageURL gets next page OData collection
func (contentTypesResp *ContentTypesResp) NextPageURL() string {
	return getODataCollectionNextPageURL(*contentTypesResp)
}

// HasNextPage returns is true if next page exists
func (contentTypesResp *ContentTypesResp) HasNextPage() bool {
	return contentTypesResp.NextPageURL() != ""
}
