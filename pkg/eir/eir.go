package eir
import ("encoding/json";"time")
type Record struct{R,D,O,T string}
func New(r,d,o string)*Record{return &Record{R:r,D:d,O:o,T:time.Now().UTC().Format(time.RFC3339Nano)}}
func (e *Record)JSON()[]byte{b,_:=json.Marshal(e);return b}
