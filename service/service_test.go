package service

import (
	"secret-santa/model"
	"testing"
)

func TestRoomService(t *testing.T) {
	partner1 := &model.Partner{
		Id:          "1",
		FullName:    "1",
		NickName:    "1",
		Description: "1",
	}

	room1, err := model.NewRoom(partner1)
	if err != nil {
		t.Error(err)
	}
	if len(room1.Participants) != 1 {
		t.Error("room participant creator was expected")
	}
	roomService := NewRoomServiceImp()
	if err != nil {
		t.Error(err)
	}
	//if len(roomService.R) != 1 {
	//	t.Error("room participant creator was expected")
	//}

	// happy flow
	err = roomService.DeleteRoom(room1.Id)
	if err != nil {
		t.Error(err)
	}

	// happy flow
	err = roomService.AddRoom(room1)
	if err != nil {
		t.Error(err)
	}

	// failed flow
	err = roomService.AddRoom(room1)
	if err == nil {
		t.Errorf("expected: duplicate error")
	}

	rooms := roomService.rooms

	_, ok := rooms[room1.Id]
	if !ok {
		t.Errorf("not found room id: %s", room1.Id)
	}

	// failed update flow

	// happy flow
	err = roomService.DeleteRoom(room1.Id)
	if err != nil {
		t.Error(err)
	}
	_, ok = rooms[room1.Id]
	if ok {
		t.Errorf("found room id: %s", room1.Id)
	}

	room2, err := model.NewRoom(partner1)
	if err != nil {
		t.Error(err)
	}
	// happy flow
	err = roomService.AddRooms(room1, room2)
	if err != nil {
		t.Error(err)
	}
	if len(roomService.rooms) != 2 {
		t.Errorf("expected 2 rooms")
	}

	rooms = roomService.rooms

	partner2 := &model.Partner{
		Id:          "2",
		FullName:    "2",
		NickName:    "2",
		Description: "2",
	}

	// happy flow
	err = roomService.JoinRoom(room1.Id, partner2)
	if err != nil {
		t.Error(err)
	}

	room, ok := rooms[room1.Id]
	if !ok {
		t.Errorf("room id: %s not found", room1.Id)
	}

	if len(room.Participants) != 2 {
		t.Errorf("expected added participant id:%s to the room id:%s", partner2.Id, room1.Id)
	}

	partner, ok := room.Participants[partner2.Id]
	if !ok {
		t.Errorf("partner id: %s not found", partner.Id)
	}


	// happy flow
	rr, err := roomService.Roll(room1.Id)
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
	err = roomService.LeftRoom(room1.Id, partner2.Id)
	if err != nil {
		t.Error(err)
	}
	if len(room.Participants) != 1 {
		t.Error("expected one participant")
	}
}
