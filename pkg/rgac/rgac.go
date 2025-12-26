package rgac
import("crypto/sha256";"encoding/hex";"sync")
type Chain struct{mu sync.Mutex;Head string;Log []string}
func (c *Chain)Append(h string)string{c.mu.Lock();defer c.mu.Unlock();if c.Head==""{c.Head=h}else{sum:=sha256.Sum256([]byte(c.Head+h));c.Head=hex.EncodeToString(sum[:])};c.Log=append(c.Log,c.Head);return c.Head}
