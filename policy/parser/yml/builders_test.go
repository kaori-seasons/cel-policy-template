package yml

import (
	"reflect"
	"testing"

	"github.com/google/cel-policy-templates-go/policy/model"
)

func TestBuilders_ModelMapValue(t *testing.T) {
	sv := model.NewMapValue()
	sb := &mapBuilder{
		baseBuilder: &baseBuilder{declType: model.MapType},
		mv:          sv,
	}

	// Simulate setting a role binding on an IAM grant policy
	sb.id(1)
	r, _ := sb.field(2, "role")
	r.assign("role/storage.bucket.admin")
	r.id(3)
	m, _ := sb.field(4, "members")
	m.id(5)
	m.initList()
	m0, _ := m.entry(0)
	m0.id(6)
	m0.assign("user:wiley@acme.co")

	role := model.NewField(2, "role")
	role.Ref, _ = model.NewDynValue(3, "role/storage.bucket.admin")

	members := model.NewField(4, "members")
	memberList := model.NewListValue()
	elem, _ := model.NewDynValue(6, "user:wiley@acme.co")
	memberList.Append(elem)
	members.Ref, _ = model.NewDynValue(5, memberList)

	want := model.NewMapValue()
	want.AddField(role)
	want.AddField(members)
	if !reflect.DeepEqual(sv.Fields, want.Fields) {
		t.Errorf("got %v, wanted %v", sv.Fields, want.Fields)
	}
}
