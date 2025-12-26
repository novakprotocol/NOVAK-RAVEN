package hvet
import("crypto/sha256";"crypto/sha512";"encoding/hex";"lukechampine.com/blake3")
func Compute(data []byte)map[string]string{out:=make(map[string]string);s256:=sha256.Sum256(data);s384:=sha512.Sum384(data);s512:=sha512.Sum512(data);b3:=blake3.Sum256(data);out["sha256"]=hex.EncodeToString(s256[:]);out["sha384"]=hex.EncodeToString(s384[:]);out["sha512"]=hex.EncodeToString(s512[:]);out["blake3"]=hex.EncodeToString(b3[:]);return out}
