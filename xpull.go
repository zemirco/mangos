// Copyright 2014 Garrett D'Amore
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sp

type xpull struct {
	sock ProtocolSocket
}

func (x *xpull) Init(sock ProtocolSocket) {
	x.sock = sock
}

func (x *xpull) receiver(ep Endpoint) {
	for {

		msg := ep.RecvMsg()
		if msg == nil {
			return
		}

		select {
		case x.sock.RecvChannel() <- msg:
		case <-x.sock.CloseChannel():
			return
		}
	}
}

func (*xpull) Name() string {
	return XPullName
}

func (*xpull) Number() uint16 {
	return ProtoPull
}

func (*xpull) IsRaw() bool {
	return true
}

func (*xpull) ValidPeer(peer uint16) bool {
	if peer == ProtoPush {
		return true
	}
	return false
}

func (x *xpull) AddEndpoint(ep Endpoint) {
	go x.receiver(ep)
}

func (x *xpull) RemEndpoint(ep Endpoint) {}

func (*xpull) SendHook(msg *Message) bool {
	return false
}

type xpullFactory int

func (xpullFactory) NewProtocol() Protocol {
	return &xpull{}
}

// XPullFactory implements the Protocol Factory for the XPULL protocol.
// The XPULL Protocol is the raw form of the PULL (Pull) protocol.
var XPullFactory xpullFactory