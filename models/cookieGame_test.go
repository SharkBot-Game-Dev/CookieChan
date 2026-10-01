package models

import (
	"gorm.io/gorm/schema"
	"sync"
	"testing"
)

func TestAssociationsReferenceDiscordUserID(t *testing.T) {
	s, err := schema.Parse(&CookieGameUser{}, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"Items", "Achievements"} {
		relation := s.Relationships.Relations[name]
		if relation == nil || len(relation.References) != 1 {
			t.Fatalf("invalid %s relation", name)
		}
		ref := relation.References[0]
		if ref.PrimaryKey.Name != "UserId" || ref.ForeignKey.Name != "UserId" {
			t.Fatalf("%s references %s instead of Discord ID", name, ref.PrimaryKey.Name)
		}
	}
}
