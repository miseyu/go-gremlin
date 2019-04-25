package gremlin

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// Request ...
type Request struct {
	RequestID string       `json:"requestId"`
	Op        string       `json:"op"`
	Processor string       `json:"processor"`
	Args      *RequestArgs `json:"args"`
}

// RequestArgs ...
type RequestArgs struct {
	Gremlin           string            `json:"gremlin,omitempty"`
	Session           string            `json:"session,omitempty"`
	Bindings          Bind              `json:"bindings,omitempty"`
	Language          string            `json:"language,omitempty"`
	Rebindings        Bind              `json:"rebindings,omitempty"`
	Sasl              string            `json:"sasl,omitempty"`
	BatchSize         int               `json:"batchSize,omitempty"`
	ManageTransaction bool              `json:"manageTransaction,omitempty"`
	Aliases           map[string]string `json:"aliases,omitempty"`
}

// FormattedReq the requests in the appropriate way
type FormattedReq struct {
	Op        string       `json:"op"`
	RequestID interface{}  `json:"requestId"`
	Args      *RequestArgs `json:"args"`
	Processor string       `json:"processor"`
}

// GraphSONSerializer ...
func GraphSONSerializer(req *Request) ([]byte, error) {
	form := NewFormattedReq(req)
	msg, err := json.Marshal(form)
	if err != nil {
		return nil, err
	}
	mimeType := []byte("application/vnd.gremlin-v2.0+json")
	var mimeLen = []byte{0x21}
	res := append(mimeLen, mimeType...)
	res = append(res, msg...)
	return res, nil

}

// NewFormattedReq ...
func NewFormattedReq(req *Request) FormattedReq {
	rID := map[string]string{"@type": "g:UUID", "@value": req.RequestID}
	sr := FormattedReq{RequestID: rID, Processor: req.Processor, Op: req.Op, Args: req.Args}

	return sr
}

// Bind ...
type Bind map[string]interface{}

// BuildQuery ...
func BuildQuery(query string) (*Request, error) {
	args := &RequestArgs{
		Gremlin:  query,
		Language: "gremlin-groovy",
	}
	u, err := uuid.NewV4()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create uuid version4")
	}
	uuidString := u.String()
	req := &Request{
		RequestID: uuidString,
		Op:        "eval",
		Processor: "",
		Args:      args,
	}
	return req, nil
}

// Bindings is args setter
func (req *Request) Bindings(bindings Bind) *Request {
	req.Args.Bindings = bindings
	return req
}

// ManageTransaction is args setter
func (req *Request) ManageTransaction(flag bool) *Request {
	req.Args.ManageTransaction = flag
	return req
}

// Aliases is args setter
func (req *Request) Aliases(aliases map[string]string) *Request {
	req.Args.Aliases = aliases
	return req
}

// Session is args setter
func (req *Request) Session(session string) *Request {
	req.Args.Session = session
	return req
}

// SetProcessor is args setter
func (req *Request) SetProcessor(processor string) *Request {
	req.Processor = processor
	return req
}
