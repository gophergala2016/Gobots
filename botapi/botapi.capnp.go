package botapi

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	server "zombiezen.com/go/capnproto2/server"
)

type AiConnector struct{ Client capnp.Client }

func (c AiConnector) Connect(ctx context.Context, params func(ConnectRequest) error, opts ...capnp.CallOption) AiConnector_connect_Results_Promise {
	if c.Client == nil {
		return AiConnector_connect_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0x9804b41cc3cba212,
			MethodID:      0,
			InterfaceName: "botapi.capnp:AiConnector",
			MethodName:    "connect",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(ConnectRequest{Struct: s}) }
	}
	return AiConnector_connect_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type AiConnector_Server interface {
	Connect(AiConnector_connect) error
}

func AiConnector_ServerToClient(s AiConnector_Server) AiConnector {
	c, _ := s.(server.Closer)
	return AiConnector{Client: server.New(AiConnector_Methods(nil, s), c)}
}

func AiConnector_Methods(methods []server.Method, s AiConnector_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0x9804b41cc3cba212,
			MethodID:      0,
			InterfaceName: "botapi.capnp:AiConnector",
			MethodName:    "connect",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := AiConnector_connect{c, opts, ConnectRequest{Struct: p}, AiConnector_connect_Results{Struct: r}}
			return s.Connect(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// AiConnector_connect holds the arguments for a server call to AiConnector.connect.
type AiConnector_connect struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  ConnectRequest
	Results AiConnector_connect_Results
}

type AiConnector_connect_Results struct{ capnp.Struct }

func NewAiConnector_connect_Results(s *capnp.Segment) (AiConnector_connect_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return AiConnector_connect_Results{}, err
	}
	return AiConnector_connect_Results{st}, nil
}

func NewRootAiConnector_connect_Results(s *capnp.Segment) (AiConnector_connect_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return AiConnector_connect_Results{}, err
	}
	return AiConnector_connect_Results{st}, nil
}

func ReadRootAiConnector_connect_Results(msg *capnp.Message) (AiConnector_connect_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return AiConnector_connect_Results{}, err
	}
	st := capnp.ToStruct(root)
	return AiConnector_connect_Results{st}, nil
}

// AiConnector_connect_Results_List is a list of AiConnector_connect_Results.
type AiConnector_connect_Results_List struct{ capnp.List }

// NewAiConnector_connect_Results creates a new list of AiConnector_connect_Results.
func NewAiConnector_connect_Results_List(s *capnp.Segment, sz int32) (AiConnector_connect_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return AiConnector_connect_Results_List{}, err
	}
	return AiConnector_connect_Results_List{l}, nil
}

func (s AiConnector_connect_Results_List) At(i int) AiConnector_connect_Results {
	return AiConnector_connect_Results{s.List.Struct(i)}
}
func (s AiConnector_connect_Results_List) Set(i int, v AiConnector_connect_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// AiConnector_connect_Results_Promise is a wrapper for a AiConnector_connect_Results promised by a client call.
type AiConnector_connect_Results_Promise struct{ *capnp.Pipeline }

func (p AiConnector_connect_Results_Promise) Struct() (AiConnector_connect_Results, error) {
	s, err := p.Pipeline.Struct()
	return AiConnector_connect_Results{s}, err
}

type ConnectRequest struct{ capnp.Struct }

func NewConnectRequest(s *capnp.Segment) (ConnectRequest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return ConnectRequest{}, err
	}
	return ConnectRequest{st}, nil
}

func NewRootConnectRequest(s *capnp.Segment) (ConnectRequest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return ConnectRequest{}, err
	}
	return ConnectRequest{st}, nil
}

func ReadRootConnectRequest(msg *capnp.Message) (ConnectRequest, error) {
	root, err := msg.Root()
	if err != nil {
		return ConnectRequest{}, err
	}
	st := capnp.ToStruct(root)
	return ConnectRequest{st}, nil
}

func (s ConnectRequest) Credentials() (Credentials, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Credentials{}, err
	}

	ss := capnp.ToStruct(p)

	return Credentials{Struct: ss}, nil
}

func (s ConnectRequest) SetCredentials(v Credentials) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewCredentials sets the credentials field to a newly
// allocated Credentials struct, preferring placement in s's segment.
func (s ConnectRequest) NewCredentials() (Credentials, error) {

	ss, err := NewCredentials(s.Struct.Segment())
	if err != nil {
		return Credentials{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s ConnectRequest) Ai() Ai {
	p, err := s.Struct.Pointer(1)
	if err != nil {

		return Ai{}
	}
	c := capnp.ToInterface(p).Client()
	return Ai{Client: c}
}

func (s ConnectRequest) SetAi(v Ai) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPointer(1, in)
}

// ConnectRequest_List is a list of ConnectRequest.
type ConnectRequest_List struct{ capnp.List }

// NewConnectRequest creates a new list of ConnectRequest.
func NewConnectRequest_List(s *capnp.Segment, sz int32) (ConnectRequest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return ConnectRequest_List{}, err
	}
	return ConnectRequest_List{l}, nil
}

func (s ConnectRequest_List) At(i int) ConnectRequest           { return ConnectRequest{s.List.Struct(i)} }
func (s ConnectRequest_List) Set(i int, v ConnectRequest) error { return s.List.SetStruct(i, v.Struct) }

// ConnectRequest_Promise is a wrapper for a ConnectRequest promised by a client call.
type ConnectRequest_Promise struct{ *capnp.Pipeline }

func (p ConnectRequest_Promise) Struct() (ConnectRequest, error) {
	s, err := p.Pipeline.Struct()
	return ConnectRequest{s}, err
}

func (p ConnectRequest_Promise) Credentials() Credentials_Promise {
	return Credentials_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p ConnectRequest_Promise) Ai() Ai {
	return Ai{Client: p.Pipeline.GetPipeline(1).Client()}
}

type Credentials struct{ capnp.Struct }

func NewCredentials(s *capnp.Segment) (Credentials, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Credentials{}, err
	}
	return Credentials{st}, nil
}

func NewRootCredentials(s *capnp.Segment) (Credentials, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Credentials{}, err
	}
	return Credentials{st}, nil
}

func ReadRootCredentials(msg *capnp.Message) (Credentials, error) {
	root, err := msg.Root()
	if err != nil {
		return Credentials{}, err
	}
	st := capnp.ToStruct(root)
	return Credentials{st}, nil
}

func (s Credentials) SecretToken() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Credentials) SetSecretToken(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// Credentials_List is a list of Credentials.
type Credentials_List struct{ capnp.List }

// NewCredentials creates a new list of Credentials.
func NewCredentials_List(s *capnp.Segment, sz int32) (Credentials_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Credentials_List{}, err
	}
	return Credentials_List{l}, nil
}

func (s Credentials_List) At(i int) Credentials           { return Credentials{s.List.Struct(i)} }
func (s Credentials_List) Set(i int, v Credentials) error { return s.List.SetStruct(i, v.Struct) }

// Credentials_Promise is a wrapper for a Credentials promised by a client call.
type Credentials_Promise struct{ *capnp.Pipeline }

func (p Credentials_Promise) Struct() (Credentials, error) {
	s, err := p.Pipeline.Struct()
	return Credentials{s}, err
}

type Ai struct{ Client capnp.Client }

func (c Ai) TakeTurn(ctx context.Context, params func(Ai_takeTurn_Params) error, opts ...capnp.CallOption) Ai_takeTurn_Results_Promise {
	if c.Client == nil {
		return Ai_takeTurn_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xd403ce7bb5b69f1f,
			MethodID:      0,
			InterfaceName: "botapi.capnp:Ai",
			MethodName:    "takeTurn",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Ai_takeTurn_Params{Struct: s}) }
	}
	return Ai_takeTurn_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Ai_Server interface {
	TakeTurn(Ai_takeTurn) error
}

func Ai_ServerToClient(s Ai_Server) Ai {
	c, _ := s.(server.Closer)
	return Ai{Client: server.New(Ai_Methods(nil, s), c)}
}

func Ai_Methods(methods []server.Method, s Ai_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xd403ce7bb5b69f1f,
			MethodID:      0,
			InterfaceName: "botapi.capnp:Ai",
			MethodName:    "takeTurn",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Ai_takeTurn{c, opts, Ai_takeTurn_Params{Struct: p}, Ai_takeTurn_Results{Struct: r}}
			return s.TakeTurn(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Ai_takeTurn holds the arguments for a server call to Ai.takeTurn.
type Ai_takeTurn struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Ai_takeTurn_Params
	Results Ai_takeTurn_Results
}

type Ai_takeTurn_Params struct{ capnp.Struct }

func NewAi_takeTurn_Params(s *capnp.Segment) (Ai_takeTurn_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Ai_takeTurn_Params{}, err
	}
	return Ai_takeTurn_Params{st}, nil
}

func NewRootAi_takeTurn_Params(s *capnp.Segment) (Ai_takeTurn_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Ai_takeTurn_Params{}, err
	}
	return Ai_takeTurn_Params{st}, nil
}

func ReadRootAi_takeTurn_Params(msg *capnp.Message) (Ai_takeTurn_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return Ai_takeTurn_Params{}, err
	}
	st := capnp.ToStruct(root)
	return Ai_takeTurn_Params{st}, nil
}

func (s Ai_takeTurn_Params) Board() (Board, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Board{}, err
	}

	ss := capnp.ToStruct(p)

	return Board{Struct: ss}, nil
}

func (s Ai_takeTurn_Params) SetBoard(v Board) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewBoard sets the board field to a newly
// allocated Board struct, preferring placement in s's segment.
func (s Ai_takeTurn_Params) NewBoard() (Board, error) {

	ss, err := NewBoard(s.Struct.Segment())
	if err != nil {
		return Board{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// Ai_takeTurn_Params_List is a list of Ai_takeTurn_Params.
type Ai_takeTurn_Params_List struct{ capnp.List }

// NewAi_takeTurn_Params creates a new list of Ai_takeTurn_Params.
func NewAi_takeTurn_Params_List(s *capnp.Segment, sz int32) (Ai_takeTurn_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Ai_takeTurn_Params_List{}, err
	}
	return Ai_takeTurn_Params_List{l}, nil
}

func (s Ai_takeTurn_Params_List) At(i int) Ai_takeTurn_Params {
	return Ai_takeTurn_Params{s.List.Struct(i)}
}
func (s Ai_takeTurn_Params_List) Set(i int, v Ai_takeTurn_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Ai_takeTurn_Params_Promise is a wrapper for a Ai_takeTurn_Params promised by a client call.
type Ai_takeTurn_Params_Promise struct{ *capnp.Pipeline }

func (p Ai_takeTurn_Params_Promise) Struct() (Ai_takeTurn_Params, error) {
	s, err := p.Pipeline.Struct()
	return Ai_takeTurn_Params{s}, err
}

func (p Ai_takeTurn_Params_Promise) Board() Board_Promise {
	return Board_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Ai_takeTurn_Results struct{ capnp.Struct }

func NewAi_takeTurn_Results(s *capnp.Segment) (Ai_takeTurn_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Ai_takeTurn_Results{}, err
	}
	return Ai_takeTurn_Results{st}, nil
}

func NewRootAi_takeTurn_Results(s *capnp.Segment) (Ai_takeTurn_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Ai_takeTurn_Results{}, err
	}
	return Ai_takeTurn_Results{st}, nil
}

func ReadRootAi_takeTurn_Results(msg *capnp.Message) (Ai_takeTurn_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return Ai_takeTurn_Results{}, err
	}
	st := capnp.ToStruct(root)
	return Ai_takeTurn_Results{st}, nil
}

func (s Ai_takeTurn_Results) Turn() (Turn_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Turn_List{}, err
	}

	l := capnp.ToList(p)

	return Turn_List{List: l}, nil
}

func (s Ai_takeTurn_Results) SetTurn(v Turn_List) error {

	return s.Struct.SetPointer(0, v.List)
}

// Ai_takeTurn_Results_List is a list of Ai_takeTurn_Results.
type Ai_takeTurn_Results_List struct{ capnp.List }

// NewAi_takeTurn_Results creates a new list of Ai_takeTurn_Results.
func NewAi_takeTurn_Results_List(s *capnp.Segment, sz int32) (Ai_takeTurn_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Ai_takeTurn_Results_List{}, err
	}
	return Ai_takeTurn_Results_List{l}, nil
}

func (s Ai_takeTurn_Results_List) At(i int) Ai_takeTurn_Results {
	return Ai_takeTurn_Results{s.List.Struct(i)}
}
func (s Ai_takeTurn_Results_List) Set(i int, v Ai_takeTurn_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Ai_takeTurn_Results_Promise is a wrapper for a Ai_takeTurn_Results promised by a client call.
type Ai_takeTurn_Results_Promise struct{ *capnp.Pipeline }

func (p Ai_takeTurn_Results_Promise) Struct() (Ai_takeTurn_Results, error) {
	s, err := p.Pipeline.Struct()
	return Ai_takeTurn_Results{s}, err
}

type Board struct{ capnp.Struct }

func NewBoard(s *capnp.Segment) (Board, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Board{}, err
	}
	return Board{st}, nil
}

func NewRootBoard(s *capnp.Segment) (Board, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Board{}, err
	}
	return Board{st}, nil
}

func ReadRootBoard(msg *capnp.Message) (Board, error) {
	root, err := msg.Root()
	if err != nil {
		return Board{}, err
	}
	st := capnp.ToStruct(root)
	return Board{st}, nil
}

func (s Board) GameId() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Board) SetGameId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s Board) Width() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s Board) SetWidth(v int16) {

	s.Struct.SetUint16(0, uint16(v))
}

func (s Board) Height() int16 {
	return int16(s.Struct.Uint16(2))
}

func (s Board) SetHeight(v int16) {

	s.Struct.SetUint16(2, uint16(v))
}

func (s Board) Robots() (Robot_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Robot_List{}, err
	}

	l := capnp.ToList(p)

	return Robot_List{List: l}, nil
}

func (s Board) SetRobots(v Robot_List) error {

	return s.Struct.SetPointer(0, v.List)
}

func (s Board) Round() int32 {
	return int32(s.Struct.Uint32(4))
}

func (s Board) SetRound(v int32) {

	s.Struct.SetUint32(4, uint32(v))
}

// Board_List is a list of Board.
type Board_List struct{ capnp.List }

// NewBoard creates a new list of Board.
func NewBoard_List(s *capnp.Segment, sz int32) (Board_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	if err != nil {
		return Board_List{}, err
	}
	return Board_List{l}, nil
}

func (s Board_List) At(i int) Board           { return Board{s.List.Struct(i)} }
func (s Board_List) Set(i int, v Board) error { return s.List.SetStruct(i, v.Struct) }

// Board_Promise is a wrapper for a Board promised by a client call.
type Board_Promise struct{ *capnp.Pipeline }

func (p Board_Promise) Struct() (Board, error) {
	s, err := p.Pipeline.Struct()
	return Board{s}, err
}

type Robot struct{ capnp.Struct }

func NewRobot(s *capnp.Segment) (Robot, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	if err != nil {
		return Robot{}, err
	}
	return Robot{st}, nil
}

func NewRootRobot(s *capnp.Segment) (Robot, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	if err != nil {
		return Robot{}, err
	}
	return Robot{st}, nil
}

func ReadRootRobot(msg *capnp.Message) (Robot, error) {
	root, err := msg.Root()
	if err != nil {
		return Robot{}, err
	}
	st := capnp.ToStruct(root)
	return Robot{st}, nil
}

func (s Robot) Id() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s Robot) SetId(v int32) {

	s.Struct.SetUint32(0, uint32(v))
}

func (s Robot) X() int16 {
	return int16(s.Struct.Uint16(4))
}

func (s Robot) SetX(v int16) {

	s.Struct.SetUint16(4, uint16(v))
}

func (s Robot) Y() int16 {
	return int16(s.Struct.Uint16(6))
}

func (s Robot) SetY(v int16) {

	s.Struct.SetUint16(6, uint16(v))
}

func (s Robot) Health() int16 {
	return int16(s.Struct.Uint16(8))
}

func (s Robot) SetHealth(v int16) {

	s.Struct.SetUint16(8, uint16(v))
}

func (s Robot) Faction() Faction {
	return Faction(s.Struct.Uint16(10))
}

func (s Robot) SetFaction(v Faction) {

	s.Struct.SetUint16(10, uint16(v))
}

// Robot_List is a list of Robot.
type Robot_List struct{ capnp.List }

// NewRobot creates a new list of Robot.
func NewRobot_List(s *capnp.Segment, sz int32) (Robot_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	if err != nil {
		return Robot_List{}, err
	}
	return Robot_List{l}, nil
}

func (s Robot_List) At(i int) Robot           { return Robot{s.List.Struct(i)} }
func (s Robot_List) Set(i int, v Robot) error { return s.List.SetStruct(i, v.Struct) }

// Robot_Promise is a wrapper for a Robot promised by a client call.
type Robot_Promise struct{ *capnp.Pipeline }

func (p Robot_Promise) Struct() (Robot, error) {
	s, err := p.Pipeline.Struct()
	return Robot{s}, err
}

type Faction uint16

// Values of Faction.
const (
	Faction_mine     Faction = 0
	Faction_opponent Faction = 1
)

// String returns the enum's constant name.
func (c Faction) String() string {
	switch c {
	case Faction_mine:
		return "mine"
	case Faction_opponent:
		return "opponent"

	default:
		return ""
	}
}

// FactionFromString returns the enum value with a name,
// or the zero value if there's no such value.
func FactionFromString(c string) Faction {
	switch c {
	case "mine":
		return Faction_mine
	case "opponent":
		return Faction_opponent

	default:
		return 0
	}
}

type Faction_List struct{ capnp.List }

func NewFaction_List(s *capnp.Segment, sz int32) (Faction_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	if err != nil {
		return Faction_List{}, err
	}
	return Faction_List{l.List}, nil
}

func (l Faction_List) At(i int) Faction {
	ul := capnp.UInt16List{List: l.List}
	return Faction(ul.At(i))
}

func (l Faction_List) Set(i int, v Faction) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type Turn struct{ capnp.Struct }
type Turn_Which uint16

const (
	Turn_Which_wait         Turn_Which = 0
	Turn_Which_move         Turn_Which = 1
	Turn_Which_attack       Turn_Which = 2
	Turn_Which_selfDestruct Turn_Which = 3
)

func (w Turn_Which) String() string {
	const s = "waitmoveattackselfDestruct"
	switch w {
	case Turn_Which_wait:
		return s[0:4]
	case Turn_Which_move:
		return s[4:8]
	case Turn_Which_attack:
		return s[8:14]
	case Turn_Which_selfDestruct:
		return s[14:26]

	}
	return "Turn_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewTurn(s *capnp.Segment) (Turn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Turn{}, err
	}
	return Turn{st}, nil
}

func NewRootTurn(s *capnp.Segment) (Turn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Turn{}, err
	}
	return Turn{st}, nil
}

func ReadRootTurn(msg *capnp.Message) (Turn, error) {
	root, err := msg.Root()
	if err != nil {
		return Turn{}, err
	}
	st := capnp.ToStruct(root)
	return Turn{st}, nil
}

func (s Turn) Which() Turn_Which {
	return Turn_Which(s.Struct.Uint16(0))
}

func (s Turn) Id() int32 {
	return int32(s.Struct.Uint32(4))
}

func (s Turn) SetId(v int32) {

	s.Struct.SetUint32(4, uint32(v))
}

func (s Turn) SetWait() {
	s.Struct.SetUint16(0, 0)
}

func (s Turn) Move() Direction {
	return Direction(s.Struct.Uint16(2))
}

func (s Turn) SetMove(v Direction) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetUint16(2, uint16(v))
}

func (s Turn) Attack() Direction {
	return Direction(s.Struct.Uint16(2))
}

func (s Turn) SetAttack(v Direction) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint16(2, uint16(v))
}

func (s Turn) SetSelfDestruct() {
	s.Struct.SetUint16(0, 3)
}

// Turn_List is a list of Turn.
type Turn_List struct{ capnp.List }

// NewTurn creates a new list of Turn.
func NewTurn_List(s *capnp.Segment, sz int32) (Turn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return Turn_List{}, err
	}
	return Turn_List{l}, nil
}

func (s Turn_List) At(i int) Turn           { return Turn{s.List.Struct(i)} }
func (s Turn_List) Set(i int, v Turn) error { return s.List.SetStruct(i, v.Struct) }

// Turn_Promise is a wrapper for a Turn promised by a client call.
type Turn_Promise struct{ *capnp.Pipeline }

func (p Turn_Promise) Struct() (Turn, error) {
	s, err := p.Pipeline.Struct()
	return Turn{s}, err
}

type Direction uint16

// Values of Direction.
const (
	Direction_north Direction = 0
	Direction_south Direction = 1
	Direction_east  Direction = 2
	Direction_west  Direction = 3
)

// String returns the enum's constant name.
func (c Direction) String() string {
	switch c {
	case Direction_north:
		return "north"
	case Direction_south:
		return "south"
	case Direction_east:
		return "east"
	case Direction_west:
		return "west"

	default:
		return ""
	}
}

// DirectionFromString returns the enum value with a name,
// or the zero value if there's no such value.
func DirectionFromString(c string) Direction {
	switch c {
	case "north":
		return Direction_north
	case "south":
		return Direction_south
	case "east":
		return Direction_east
	case "west":
		return Direction_west

	default:
		return 0
	}
}

type Direction_List struct{ capnp.List }

func NewDirection_List(s *capnp.Segment, sz int32) (Direction_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	if err != nil {
		return Direction_List{}, err
	}
	return Direction_List{l.List}, nil
}

func (l Direction_List) At(i int) Direction {
	ul := capnp.UInt16List{List: l.List}
	return Direction(ul.At(i))
}

func (l Direction_List) Set(i int, v Direction) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}
