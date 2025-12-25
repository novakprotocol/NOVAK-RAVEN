package main
import (
    "fmt"
    "log"
    "net/http"
    "os"
    "sync/atomic"
    "time"
)
var (
    heartbeat uint64
    version = "dev"
)
func metrics(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","text/plain; version=0.0.4")
    fmt.Fprintf(w,"novak_heartbeat_total %d\n",atomic.LoadUint64(&heartbeat))
    fmt.Fprintf(w,"novak_build_info{version=%q} 1\n",version)
}
func health(w http.ResponseWriter, r *http.Request){w.WriteHeader(http.StatusOK);w.Write([]byte("ok"))}
func main(){
    mux:=http.NewServeMux()
    mux.HandleFunc("/healthz",health)
    mux.HandleFunc("/metrics",metrics)
    go func(){
        log.Println("HTTP :9105 up")
        if err:=http.ListenAndServe(":9105",mux);err!=nil{log.Fatal(err)}
    }()
    logger:=log.New(os.Stdout,"",0)
    for{
        atomic.AddUint64(&heartbeat,1)
        logger.Println(time.Now().UTC().Format(time.RFC3339),"hello Novak v8 dev
        time.Sleep(5*time.Second)
    }
}
