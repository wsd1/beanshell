package main
import (
    "fmt"
    "log"
    "github.com/kr/beanstalk"
    "time"
    "strings"
    "os/exec"
)

const (
  //uri
  uri =  "127.0.0.1:11300"
  //network
  network =  "tcp"
)
//如果存在错误，则输出
func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
    panic(fmt.Sprintf("%s: %s", msg, err))
  }
}

func main(){

  conn, err := beanstalk.Dial(network, uri)
  failOnError(err, "Failed to connect to beanstalkd")
  defer conn.Close()

  tubeSet := beanstalk.NewTubeSet(conn, "build-req")
  tube := &beanstalk.Tube{conn, "build-ack"}


  for {
    fmt.Println("o")
    //从队列中取出

    id, build_request, terr := tubeSet.Reserve(1 * time.Hour)
//  id, build_request, terr := conn.Reserve(5 * time.Second)
    if terr != nil {
      fmt.Println("[ERR]:", terr.Error())
      if strings.Contains(terr.Error(), "broken pipe"){
        break;
      }

      if strings.Contains(terr.Error(), "timeout"){
        continue;
      }
    }

    //从队列中清掉
    err = conn.Delete(id)
    if err != nil {
      fmt.Println(" [Consumer] Delete err:", err, " id:", id)
    }

    build_result, err := exec.Command(string(build_request)).CombinedOutput()
    if err != nil {
        fmt.Printf("[%d]\n%s", len(build_result), build_result)
        fmt.Println("[ERR]:", err.Error())
        build_result = []byte(err.Error())
        //log.Fatal(err)
    }

    //fmt.Printf("[%d]\n%s", len(build_result), build_result)
    if len(build_result) > 0 {
      id, terr = tube.Put(build_result, 1024, 0, time.Minute)
    }


  }

  fmt.Println("Consumer() end. ")
}

