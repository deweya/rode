package eventmanager

import (
	grafeas "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
)

type EventManager interface {
	Initialize(attesterName string) error
	PublishAttestation(attesterName string, occurrence *grafeas.Occurrence) error
	PublishPublicKey(attesterName string, publicKey []byte) error
	Subscribe(attesterName string) error
	Unsubscribe(attesterName string) error
}
