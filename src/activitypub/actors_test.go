package activitypub

import (
	"testing"
)

func TestActorNew(t *testing.T) {
	var testValue = ObjectId("test")
	var testType = ApplicationType

	o := ActorNew(testValue, testType)

	if o.Id != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != testType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, testType)
	}

	n := ActorNew(testValue, "")
	if n.Id != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", n.Id, testValue)
	}
	if n.Type != ActorType {
		t.Errorf("APObject Type '%v' different than expected '%v'", n.Type, ActorType)
	}
}

func TestPersonNew(t *testing.T) {
	var testValue = ObjectId("test")

	o := PersonNew(testValue)
	if o.Id != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != PersonType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, PersonType)
	}
}

func TestApplicationNew(t *testing.T) {
	var testValue = ObjectId("test")

	o := ApplicationNew(testValue)
	if o.Id != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != ApplicationType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, ApplicationType)
	}
}

func TestGroupNew(t *testing.T) {
	var testValue = ObjectId("test")

	o := GroupNew(testValue)
	if o.Id != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != GroupType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, GroupType)
	}
}

func TestOrganizationNew(t *testing.T) {
	var testValue = ObjectId("test")

	o := OrganizationNew(testValue)
	if o.Id != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != OrganizationType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, OrganizationType)
	}
}

func TestServiceNew(t *testing.T) {
	var testValue = ObjectId("test")

	o := ServiceNew(testValue)
	if o.Id != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != ServiceType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, ServiceType)
	}
}

func TestValidActorType(t *testing.T) {
	var invalidType ActivityVocabularyType = "RandomType"

	if ValidActorType(invalidType) {
		t.Errorf("APObject Type '%v' should not be valid", invalidType)
	}
	for _, validType := range validActorTypes {
		if !ValidActorType(validType) {
			t.Errorf("APObject Type '%v' should be valid", validType)
		}
	}
}
