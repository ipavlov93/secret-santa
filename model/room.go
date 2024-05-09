package model

import (
	"fmt"
	"math/rand"
	"time"
)

type Room struct {
	Id           string
	createdBy    string
	Participants map[string]Partner
	lastResult   *RollResult
	// history
	RollResult map[string]RollResult
}

func NewRoom(partner *Partner) (*Room, error) {
	room := NewRoomNoParticipants(partner.Id)
	return room, room.AddPartner(partner)
}

func NewRoomNoParticipants(createdBy string) *Room {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	room := &Room{
		Id: fmt.Sprint(r.Int()),
		createdBy:    createdBy,
		RollResult:   make(map[string]RollResult),
		Participants: make(map[string]Partner),
	}
	return room
}

func (r *Room) AddPartners(partners ...*Partner) error {
	if r == nil {
		return nil
	}
	for _, participant := range partners {
		err := r.AddPartner(participant)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Room) AddPartner(partner *Partner) error {
	if r == nil {
		return nil
	}
	if partner == nil {
		return fmt.Errorf("validation error: empty request")
	}
	_, found := r.Participants[partner.Id]
	if found {
		return fmt.Errorf("duplicate error: particapant id:%s already exists", partner.Id)
	}
	r.Participants[partner.Id] = *partner
	return nil
}

func (r *Room) DeletePartner(participantId string) error {
	if participantId == "" {
		return fmt.Errorf("validation error: zero participant id")
	}
	delete(r.Participants, participantId)

	return nil
}

func (r *Room) ClearRoom() {
	*r = *NewRoomNoParticipants(r.createdBy)
}

// Roll returns random participant/partner couples
func (r *Room) Roll() (rollResult *RollResult, err error) {
	if len(r.Participants)%2 != 0 {
		return nil, fmt.Errorf("not enough participants to start, even number of participants required")
	}
	return r.roll("all participants have got a partner")
}

// IgnoreOddParticipantsNumber same as Roll() for odd number of participants
func (r *Room) IgnoreOddParticipantsNumber() (rollResult *RollResult, err error) {
	return r.roll("odd participants number")
}

// roll returns random participant/partner couples
// description is optional
func (r *Room) roll(description string) (rollResult *RollResult, err error) {
	if len(r.Participants) < 2 {
		return nil, fmt.Errorf("not enough participants to start")
	}

	rollResultId := fmt.Sprint(len(r.RollResult) + 1)
	rollResult = &RollResult{
		Id:          rollResultId,
		ResultMap:   make(map[string][2]string),
		Description: description,
		CreatedAt:   time.Now(),
	}

	i := 0
	var previousPartnerId string

	// todo: edit fori loop for better readability
	for _, partner := range r.Participants {
		i++
		if i%2 == 0 {
			rollResult.ResultMap[fmt.Sprint(i/2)] = [2]string{partner.Id, previousPartnerId}
			continue
		}
		previousPartnerId = partner.Id
	}
	rollResult.Description = fmt.Sprintf("%s: participant id: %s has no partner", rollResult.Description, previousPartnerId)

	r.RollResult[rollResultId] = *rollResult
	r.lastResult = rollResult

	return rollResult, nil
}

//func (r *Room) GetParticipants() map[string]Partner {
//	if r == nil {
//		return nil
//	}
//
//	return r.Participants
//}
