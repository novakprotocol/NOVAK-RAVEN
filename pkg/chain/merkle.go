package chain
import("crypto/sha256";"encoding/hex";"encoding/json";"os";"sync")
type Chain struct{mu sync.Mutex;Path string;Nodes []string;Root string}
func New(path string)(*Chain,error){c:=&Chain{Path:path};_ = c.load();return c,nil}
func (c *Chain)load()error{d,err:=os.ReadFile(c.Path);if err!=nil{return nil};return json.Unmarshal(d,c)}
func (c *Chain)persist(){b,_:=json.MarshalIndent(c,"","  ");_ = os.WriteFile(c.Path,b,0600)}
func (c *Chain)Append(d []byte)string{c.mu.Lock();defer c.mu.Unlock();h:=sha256.Sum256(d);c.Nodes=append(c.Nodes,hex.EncodeToString(h[:]));layer:=c.Nodes;for len(layer)>1{var n []string;for i:=0;i<len(layer);i+=2{if i+1<len(layer){s:=sha256.Sum256([]byte(layer[i]+layer[i+1]));n=append(n,hex.EncodeToString(s[:]));}else{n=append(n,layer[i])}};layer=n};c.Root=layer[0];c.persist();return c.Root}
func (c *Chain)Depth()int{return len(c.Nodes)};func (c *Chain)RootHash()string{c.mu.Lock();defer c.mu.Unlock();return c.Root}
