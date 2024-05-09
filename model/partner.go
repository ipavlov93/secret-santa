package model

import (
	"fmt"
	"math/rand"
	"time"
)

// Partner or Participant
type Partner struct {
	Id             string // hash or id
	FullName       string
	NickName       string // required
	Description    string // short intro
	PresentAdvices string // some advices about present
}

func NewPartner(nickName, fullName string) *Partner {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Partner{
		Id: fmt.Sprint(r.Int()),
		NickName: nickName,
		FullName: fullName,
	}
}

func (p Partner) String() string {
	return fmt.Sprintf("id:%s nickname:%s", p.Id, p.NickName)
}

func (p *Partner) Equals(partner *Partner) bool {
	if p.Id != partner.Id {
		return false
	}
	return p.NickName == partner.NickName
}
