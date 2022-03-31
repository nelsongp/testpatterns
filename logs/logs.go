package logs

import (
	"github.com/nelsongp/testpatterns/credentialValidator"
	"github.com/nelsongp/testpatterns/message"
	"github.com/nelsongp/testpatterns/notifier"
)

type PrintLog struct {
	Val credentialValidator.CredentialValidator
}

func (p *PrintLog) Validate(ID string, password string) error {
	err := p.Val.Validate(ID, password)
	if err != nil {
		m := message.Message{}
		m.AddObserver("texto", &notifier.Text{})
		m.AddObserver("bd", &notifier.DB{})
		m.ID = ID
		m.Password = password
		m.NotifyObservers()
	}
	return err
}
