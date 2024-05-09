package model

import (
	"testing"
)

func TestRoom(t *testing.T) {
	partner1 := &Partner{
		Id:          "1",
		FullName:    "1",
		NickName:    "1",
		Description: "1",
	}

	r, err := NewRoom(partner1)
	if err != nil {
		t.Error(err)
	}
	if len(r.Participants) != 1 {
		t.Error("room participant creator was expected")
	}

	room := NewRoomNoParticipants(partner1.Id)
	if len(room.Participants) > 0 {
		t.Error("no room participants were expected")
	}
	if room.createdBy != partner1.Id {
		t.Errorf("expected creator id: %s", partner1.Id)
	}

	// failed flow
	// roll not enough participants, 0
	_, err = room.Roll()
	if err == nil {
		t.Errorf("expected: not enough participants")
	}

	// happy flow
	// nil map
	err = room.DeletePartner(partner1.Id)
	if err != nil {
		t.Error(err)
	}

	// happy flow
	err = room.AddPartner(partner1)
	if err != nil {
		t.Error(err)
	}

	// failed flow
	err = room.AddPartner(partner1)
	if err == nil {
		t.Errorf("expected: duplicate error")
	}

	ps := room.Participants

	p, ok := ps[partner1.Id]
	if !ok {
		t.Errorf("not found participant id: %s", partner1.Id)
	}
	if !p.Equals(partner1) {
		t.Errorf("got %s but want %s", p, partner1)
	}

	// failed update flow
	//oldNickname := p.NickName

	//newNickname := "2"
	//p.NickName = newNickname
	//t.Log(p.NickName)
	//t.Log(partner1.NickName)

	p, ok = ps[partner1.Id]
	if !ok {
		t.Errorf("not found participant id: %s", partner1.Id)
	}
	if !p.Equals(partner1) {
		t.Errorf("got nickname %s but want %s", p.NickName, partner1.NickName)
	}

	// happy flow
	err = room.DeletePartner(partner1.Id)
	if err != nil {
		t.Error(err)
	}
	p, ok = ps[partner1.Id]
	if ok {
		t.Errorf("found participant id: %s", partner1.Id)
	}

	// failed flow
	// roll not enough participants, 1
	rr, err := room.Roll()
	if err == nil {
		t.Errorf("not found participant id: %s", partner1.Id)
	}
	if rr != nil {
		t.Errorf("error expected")
	}

	partner2 := &Partner{
		Id:          "2",
		FullName:    "2",
		NickName:    "2",
		Description: "2",
	}
	// happy flow
	err = room.AddPartners(partner1, partner2)
	if err != nil {
		t.Error(err)
	}
	if len(room.Participants) != 2 {
		t.Errorf("expected 2 participants")
	}

	// happy flow
	rr, err = room.Roll()
	if err != nil {
		t.Error(err)
	}

	resultMapId := "1"
	pair, ok := rr.ResultMap[resultMapId]
	if !ok {
		t.Errorf("result map id: %s not found", resultMapId)
	}

	if rr == nil || len(rr.ResultMap) != 1 || len(pair) != 2 {
		t.Errorf("array with 2 items expected")
	}
	for i, participant := range pair {
		if participant == "" {
			t.Errorf("got empty participant in pair:%d", i)
		}
	}

	// happy flow
	room.ClearRoom()
	if room.createdBy != partner1.Id {
		t.Errorf("expected creator id: %s", partner1.Id)
	}
	if len(room.Participants) != 0 {
		t.Error("expected empty participant list")
	}
	if room.lastResult != nil {
		t.Error("expected empty last result")
	}
	if len(room.RollResult) != 0 {
		t.Error("expected empty roll result list")
	}
}
