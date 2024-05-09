package service

import (
	"fmt"
	"net/http"
	"secret-santa/model"
	"secret-santa/utils"
)

type IRoomService interface {
	CreateRoom(ownerKey string) (roomKey string, err error)
	JoinRoom(roomKey string, partners ...*model.Partner) error
	JoinSecretRoom(roomKey string, partners ...*model.Partner) error
	LeftRoom(roomKey, partnerKey string) error
	Roll(roomKey string) (*model.RollResult, error)

	////KickPartner owner method
	//KickPartner(roomKey string, partnerKeys ...string) error
	//// Cancel party, owner method
	//Cancel(roomKey string) error
}

type RoomServiceImp struct {
	rooms map[string]model.Room
}

func NewRoomServiceImp() *RoomServiceImp {
	return &RoomServiceImp{
		rooms: make(map[string]model.Room),
	}
}

func (r *RoomServiceImp) CreateRoom(createdBy string) (roomKey string, err error) {
	if r == nil {
		return "", nil
	}
	room := model.NewRoomNoParticipants(createdBy)
	r.rooms[room.Id] = *room
	return room.Id, nil
}

func (r *RoomServiceImp) JoinRoom(roomKey string, partners ...*model.Partner) error {
	if r == nil {
		return nil
	}
	room, ok := r.rooms[roomKey]
	if !ok {
		return fmt.Errorf("room id:%s not found", roomKey)
	}
	err := room.AddPartners(partners...)
	if err != nil {
		return err
	}

	return nil
}

func (r *RoomServiceImp) JoinSecretRoom(roomKey string, password string, partners ...*model.Partner) error {
	if r == nil {
		return nil
	}
	room, ok := r.rooms[roomKey]
	if !ok {
		return fmt.Errorf("room id:%s not found", roomKey)
	}
	err := room.AddPartners(partners...)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword(password, []byte(request.Password)); err != nil {
		return tokens, utils.AppError{
			StatusCode: http.StatusUnauthorized,
		}
	}

	return nil
}

func (r *RoomServiceImp) LeftRoom(roomKey, partnerKey string) error {
	if r == nil {
		return nil
	}
	room, ok := r.rooms[roomKey]
	if !ok {
		return nil
	}
	err := room.DeletePartner(partnerKey)
	if err != nil {
		return err
	}

	return nil
}

func (r *RoomServiceImp) Roll(roomKey string) (*model.RollResult, error) {
	if r == nil {
		return nil, nil
	}
	room, ok := r.rooms[roomKey]
	if !ok {
		return nil, fmt.Errorf("room id:%s not found", roomKey)
	}

	rollResult, err := room.Roll()
	if err != nil {
		return nil, err
	}

	return rollResult, nil
}

func (r *RoomServiceImp) AddRooms(rooms ...*model.Room) error {
	if r == nil {
		return nil
	}
	for _, room := range rooms {
		err := r.AddRoom(room)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RoomServiceImp) AddRoom(room *model.Room) error {
	if r == nil {
		return nil
	}
	if room == nil {
		return fmt.Errorf("validation error: empty request")
	}
	_, found := r.rooms[room.Id]
	if found {
		return fmt.Errorf("duplicate error: room id:%s already exists", room.Id)
	}
	r.rooms[room.Id] = *room
	return nil
}

func (r *RoomServiceImp) DeleteRoom(roomId string) error {
	if roomId == "" {
		return fmt.Errorf("validation error: zero room id")
	}
	delete(r.rooms, roomId)

	return nil
}
