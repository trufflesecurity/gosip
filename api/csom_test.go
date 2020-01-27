package api

import (
	"testing"

	"github.com/koltyakov/gosip/csom"
)

func TestCsomRequest(t *testing.T) {
	t.Parallel()
	checkClient(t)

	sp := NewHTTPClient(spClient)

	b := csom.NewBuilder()

	b.AddObject(csom.NewObject(`<Property Id="{{.ID}}" ParentId="{{.ParentID}}" Name="Web" />`), nil)
	b.AddAction(csom.NewAction(`
		<Query Id="{{.ID}}" ObjectPathId="{{.ObjectID}}">
			<Query SelectAllProperties="true">
				<Properties />
			</Query>
		</Query>
	`), nil)

	csomXML, err := b.Compile()
	if err != nil {
		t.Error(err)
	}

	if _, err := sp.ProcessQuery(spClient.AuthCnfg.GetSiteURL(), []byte(csomXML)); err != nil {
		t.Error(err)
	}
}