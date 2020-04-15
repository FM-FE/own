package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "regexp"
    "strconv"
    "strings"
    "time"

    "github.com/gorilla/websocket"
    L "hualu.com/logger"
    //"hualu.com/network/netInfo"
    "hualu.com/salt-go"
)

type CommonRequest struct {
    Ip string `json:"ip,omitempty"`
}

type CommonResponse struct {
    Result string `json:"result"`
    Errors string `json:"errors,omitempty"`
}

//interfaceflow Info
type InterfaceFlowInfo struct {
    Label string `json:"label,omitempty"`
    Rx    string `json:"rx,omitempty"`
    Tx    string `json:"tx,omitempty"`
    Avirx string `json:"avirx,omitempty"`
    Avitx string `json:"avitx,omitempty"`
}
type InterfaceFlowListRsp struct {
    CommonResponse
    InterfaceFlowList []InterfaceFlowInfo `json:"interfaceflowlist,omitempty"`
}

//memory info
type MemoryInfo struct {
    Mem          string `json:"memory,omitempty"`
    Total        string `json:"total,omitempty"`
    Used         string `json:"used,omitempty"`
    Free         string `json:"free,omitempty"`
    BuffAndCache string `json:"buffandcache,omitempty"`
}
type MemoryInfoListRsp struct {
    CommonResponse
    MemoryInfoList []MemoryInfo `json:"memoryinfolist,omitempty"`
}

//cpu info
type CpuInfo struct {
    Cpu string `json:"cpu,omitempty"`
    Us  string `json:"us,omitempty"`
    Sy  string `json:"sy,omitempty"`
    Id  string `json:"id,omitempty"`
    Ni  string `json:"ni,omitempty"`
    Wa  string `json:"wa,omitempty"`
    Hi  string `json:"hi,omitempty"`
    Si  string `json:"si,omitempty"`
    St  string `json:"st,omitempty"`
}
type CpuInfoListRsp struct {
    CommonResponse
    CpuInfoList []CpuInfo `json:"cpuinfolist,omitempty"`
}

//sysruntime info
type RuntimeInfo struct {
    Runtime string `json:"runtime,omitempty"`
}
type RuntimeInfoListRsp struct {
    CommonResponse
    RuntimeInfoList []RuntimeInfo `json:"runtimeinfolist,omitempty"`
}

var addr = flag.String("addr", "192.168.146.128:7070", "http service address")
var mess map[string]interface{}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    }} // resolve request origin not allowed

func handleInterfaceString(m interface{}) string {
    var v string
    switch value := m.(type) {
    case string:
        v = value
        //fmt.Println(v)
        //fmt.Println(reflect.TypeOf(v))
    }
    return v
}

func DeleteExtraSpace(s string) string {
    //删除字符串中的多余空格，有多个空格时，仅保留一个空格
    s1 := strings.Replace(s, "  ", " ", -1)      //替换tab为空格
    regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
    reg, _ := regexp.Compile(regstr)             //编译正则表达式
    s2 := make([]byte, len(s1))                  //定义字符数组切片
    copy(s2, s1)                                 //将字符串复制到切片
    spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
    for len(spc_index) > 0 {                     //找到适配项
        s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
        spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
    }
    return string(s2)
}

func TimeProcess(s int64) string {
    day := s / 86400
    hour := (s % 86400) / 3600
    minute := (s % 3600) / 60
    seconds := s % 60
    str := strconv.FormatInt(day, 10) + "天" + strconv.FormatInt(hour, 10) + "时" + strconv.FormatInt(minute, 10) + "分" + strconv.FormatInt(seconds, 10) + "秒"
    return str
}

func NewProcessMemoryInfoList(w http.ResponseWriter, r *http.Request) {
    var rsp MemoryInfoListRsp
    var memoryinfolist []MemoryInfo
    var upgrader = websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true
        }} //resolve request origin not allowed
    args := []string{}
    var mess map[string]interface{}

    // interrupt := make(chan os.Signal, 1)
    // signal.Notify(interrupt, os.Interrupt)

    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer c.Close()

    done := make(chan struct{})
    // go func() {
    //defer close(done)
    // }()
    // for {
    _, message, err := c.ReadMessage()
    if err != nil {
        log.Println("read:", err)
        // break
    }
    log.Printf("recv from client: %s", message)
    fmt.Println(string(message))
    if e := json.Unmarshal(message, &mess); e != nil {
        rsp.Result = "ERROR"
        rsp.Errors = e.Error()
        return
    }
    fmt.Println(mess["ip"])
    fmt.Println(mess["timeinterval"])
    //     break
    // }

    go func() {
        defer close(done)
        for {
            time.Sleep(time.Millisecond)
            _, message, err := c.ReadMessage()
            if err != nil {
                log.Println("read:", err)
                return
            }
            log.Printf("recv: %s", message)
        }
    }()

    seconds, _ := strconv.Atoi(handleInterfaceString(mess["timeinterval"]))
    ticker := time.NewTicker(time.Duration(seconds) * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-done:
            return
        case t := <-ticker.C:
            log.Println(t)
            //mss := "response cpuinfo"
            //err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
            // cmd
            saltClient, _ := saltgo.NewClient(handleInterfaceString(mess["ip"]))
            args = append(args, " free -h | grep Mem ")
            ret, _ := saltClient.Reader("cmd.run", args)
            str := handleInterfaceString(ret)
            str = DeleteExtraSpace(str)
            lines := strings.Split(str, " ")
            // response
            memoryinfolist = append(memoryinfolist, MemoryInfo{Mem: "Mem", Total: lines[1], Used: lines[2], Free: lines[3], BuffAndCache: lines[5]})

            rsp.Result = "OK"
            rsp.MemoryInfoList = memoryinfolist
            buf, _ := json.Marshal(&rsp)

            err := c.WriteMessage(websocket.TextMessage, buf)
            if err != nil {
                log.Println("write:", err)
                return
            }
            memoryinfolist, args = nil, nil
            // case <-interrupt:
            //     log.Println("interrupt")
            //     // Cleanly close the connection by sending a close message and then
            //     // waiting (with timeout) for the server to close the connection.
            //     err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
            //     if err != nil {
            //         log.Println("write close:", err)
            //         return
            //     }
            //     select {
            //     case <-done:
            //     case <-time.After(time.Duration(seconds) * time.Second):
            //     }
            //     return
        }
    }
}
func NewProcessInterfaceFlowList(w http.ResponseWriter, r *http.Request) {
    var rsp InterfaceFlowListRsp
    var interfaceflowlist []InterfaceFlowInfo
    RX_preargs := []string{}
    TX_preargs := []string{}
    RX_nextargs := []string{}
    TX_nextargs := []string{}
    var upgrader = websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true
        }} //resolve request origin not allowed
    var mess map[string]interface{}

    // interrupt := make(chan os.Signal, 1)
    // signal.Notify(interrupt, os.Interrupt)

    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer c.Close()

    done := make(chan struct{})
    // go func() {
    //defer close(done)
    // }()
    // for {
    _, message, err := c.ReadMessage()
    if err != nil {
        log.Println("read:", err)
        // break
    }
    log.Printf("recv from client: %s", message)
    fmt.Println(string(message))
    if e := json.Unmarshal(message, &mess); e != nil {
        rsp.Result = "ERROR"
        rsp.Errors = e.Error()
        return
    }
    fmt.Println(mess["ip"])
    fmt.Println(mess["timeinterval"])
    //     break
    // }

    go func() {
        defer close(done)
        for {
            time.Sleep(time.Millisecond)
            _, message, err := c.ReadMessage()
            if err != nil {
                log.Println("read:", err)
                return
            }
            log.Printf("recv: %s", message)
        }
    }()

    seconds, _ := strconv.Atoi(handleInterfaceString(mess["timeinterval"]))
    ticker := time.NewTicker(time.Duration(seconds) * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-done:
            return
        case t := <-ticker.C:
            log.Println(t)
            //mss := "response cpuinfo"
            //err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
            // cmd
            saltClient, _ := saltgo.NewClient(handleInterfaceString(mess["ip"]))
            RX_preargs = append(RX_preargs, " cat /proc/net/dev | grep "+handleInterfaceString(mess["ethn"])+" | sed 's/:/ /g' | awk '{print $2}' ")
            RX_pre, _ := saltClient.Reader("cmd.run", RX_preargs)
            TX_preargs = append(TX_preargs, " cat /proc/net/dev | grep "+handleInterfaceString(mess["ethn"])+" | sed 's/:/ /g' | awk '{print $10}' ")
            TX_pre, _ := saltClient.Reader("cmd.run", TX_preargs)

            time.Sleep(time.Duration(seconds) * time.Second)

            RX_nextargs = append(RX_nextargs, " cat /proc/net/dev | grep "+handleInterfaceString(mess["ethn"])+" | sed 's/:/ /g' | awk '{print $2}' ")
            RX_next, _ := saltClient.Reader("cmd.run", RX_nextargs)
            TX_nextargs = append(TX_nextargs, " cat /proc/net/dev | grep "+handleInterfaceString(mess["ethn"])+" | sed 's/:/ /g' | awk '{print $10}' ")
            TX_next, _ := saltClient.Reader("cmd.run", TX_nextargs)

            RX_nextmp := handleInterfaceString(RX_next)
            RX_nextint, _ := strconv.Atoi(RX_nextmp)
            RX_pretmp := handleInterfaceString(RX_pre)
            RX_preint, _ := strconv.Atoi(RX_pretmp)
            RX := RX_nextint - RX_preint
            AVIRX := RX / seconds
            fmt.Println("RX: ", RX)
            fmt.Println("AVIRX: ", AVIRX)

            TX_nextmp := handleInterfaceString(TX_next)
            TX_nextint, _ := strconv.Atoi(TX_nextmp)
            TX_pretmp := handleInterfaceString(TX_pre)
            TX_preint, _ := strconv.Atoi(TX_pretmp)
            TX := TX_nextint - TX_preint
            AVITX := TX / seconds
            fmt.Println("TX: ", TX)
            fmt.Println("AVITX: ", AVITX)

            // response
            interfaceflowlist = append(interfaceflowlist, InterfaceFlowInfo{Label: handleInterfaceString(mess["ethn"]), Rx: strconv.Itoa(RX) + "B", Tx: strconv.Itoa(TX) + "B", Avirx: strconv.Itoa(AVIRX) + "B/s", Avitx: strconv.Itoa(AVITX) + "B/s"})

            rsp.Result = "OK"
            rsp.InterfaceFlowList = interfaceflowlist
            buf, _ := json.Marshal(&rsp)

            err := c.WriteMessage(websocket.TextMessage, buf)
            if err != nil {
                log.Println("write:", err)
                return
            }
            interfaceflowlist, RX_preargs, TX_preargs, RX_nextargs, TX_nextargs = nil, nil, nil, nil, nil
            // case <-interrupt:
            //     log.Println("interrupt")
            //     // Cleanly close the connection by sending a close message and then
            //     // waiting (with timeout) for the server to close the connection.
            //     err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
            //     if err != nil {
            //         log.Println("write close:", err)
            //         return
            //     }
            //     select {
            //     case <-done:
            //     case <-time.After(time.Duration(seconds) * time.Second):
            //     }
            //     return
        }
    }
}
func NewProcesCpuInfoList(w http.ResponseWriter, r *http.Request) {
    var rsp CpuInfoListRsp
    var cpuinfolist []CpuInfo
    var upgrader = websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true
        }} //resolve request origin not allowed
    args := []string{}
    var mess map[string]interface{}

    // interrupt := make(chan os.Signal, 1)
    // signal.Notify(interrupt, os.Interrupt)

    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer c.Close()

    done := make(chan struct{})
    // go func() {
    //defer close(done)
    // }()
    // for {
    _, message, err := c.ReadMessage()
    if err != nil {
        log.Println("read:", err)
        // break
    }
    log.Printf("recv from client: %s", message)
    fmt.Println(string(message))
    if e := json.Unmarshal(message, &mess); e != nil {
        rsp.Result = "ERROR"
        rsp.Errors = e.Error()
        return
    }
    fmt.Println(mess["ip"])
    fmt.Println(mess["timeinterval"])
    // break
    // }

    go func() {
        defer close(done)
        for {
            time.Sleep(time.Millisecond * 1)
            _, message, err := c.ReadMessage()
            if err != nil {
                log.Println("read:", err)
                return
            }
            log.Printf("recv: %s", message)
        }
    }()

    seconds, _ := strconv.Atoi(handleInterfaceString(mess["timeinterval"]))
    ticker := time.NewTicker(time.Duration(seconds) * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-done:
            return
        case t := <-ticker.C:
            log.Println(t)
            // cmd
            saltClient, _ := saltgo.NewClient(handleInterfaceString(mess["ip"]))
            args = append(args, " top -bn 1 | head -n 4 | grep Cpu ")
            ret, _ := saltClient.Reader("cmd.run", args)
            str := handleInterfaceString(ret)
            str = DeleteExtraSpace(strings.Replace(str, ",", " ", -1))
            lines := strings.Split(str, " ")
            // response
            cpuinfolist = append(cpuinfolist, CpuInfo{Cpu: "%Cpu(s)", Us: lines[1], Sy: lines[3], Id: lines[7], Ni: lines[5], Wa: lines[9], Hi: lines[11], Si: lines[13], St: lines[15]})

            rsp.Result = "OK"
            rsp.CpuInfoList = cpuinfolist
            buf, _ := json.Marshal(&rsp)

            err := c.WriteMessage(websocket.TextMessage, buf)
            if err != nil {
                log.Println("write:", err)
                return
            }
            cpuinfolist, args = nil, nil
            // case <-interrupt:
            //     log.Println("interrupt")
            //     // Cleanly close the connection by sending a close message and then
            //     // waiting (with timeout) for the server to close the connection.
            //     err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
            //     if err != nil {
            //         log.Println("write close:", err)
            //         return
            //     }
            //     select {
            //     case <-done:
            //     case <-time.After(time.Duration(seconds) * time.Second):
            //     }
            //     return
        }
    }
}

//test function
func echo(w http.ResponseWriter, r *http.Request) {
    var rsp CpuInfoListRsp
    var cpuinfolist []CpuInfo
    args := []string{}
    var mess map[string]interface{}
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt)

    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer c.Close()

    done := make(chan struct{})

    // go func() {
    //defer close(done)
    // }()

    for {
        _, message, err := c.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }
        log.Printf("recv from client: %s", message)
        fmt.Println(string(message))
        if e := json.Unmarshal(message, &mess); e != nil {
            rsp.Result = "ERROR"
            rsp.Errors = e.Error()
            return
        }
        fmt.Println(mess["ip"])
        fmt.Println(mess["timeinterval"])
        break
    }

    seconds, _ := strconv.Atoi(handleInterfaceString(mess["timeinterval"]))
    ticker := time.NewTicker(time.Duration(seconds) * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-done:
            return
        case t := <-ticker.C:
            log.Println(t)
            //mss := "response cpuinfo"
            //err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
            // cmd
            saltClient, _ := saltgo.NewClient(handleInterfaceString(mess["ip"]))
            args = append(args, " top -bn 1 | head -n 4 | grep Cpu ")
            ret, _ := saltClient.Reader("cmd.run", args)
            str := handleInterfaceString(ret)
            str = DeleteExtraSpace(str)
            lines := strings.Split(str, " ")
            // response
            cpuinfolist = append(cpuinfolist, CpuInfo{Cpu: "%Cpu(s)", Us: lines[1], Sy: lines[3], Id: lines[7], Ni: lines[5], Wa: lines[9], Hi: lines[11], Si: lines[13], St: lines[15]})

            rsp.Result = "OK"
            rsp.CpuInfoList = cpuinfolist
            buf, _ := json.Marshal(&rsp)

            err := c.WriteMessage(websocket.TextMessage, buf)
            if err != nil {
                log.Println("write:", err)
                return
            }
            cpuinfolist, args = nil, nil
        case <-interrupt:
            log.Println("interrupt")

            // Cleanly close the connection by sending a close message and then
            // waiting (with timeout) for the server to close the connection.
            err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
            if err != nil {
                log.Println("write close:", err)
                return
            }
            select {
            case <-done:
            case <-time.After(time.Duration(seconds) * time.Second):
            }
            return
        }
    }
}
func ProcessSysRunTimeInfoList(w http.ResponseWriter, r *http.Request) {
    var rsp RuntimeInfoListRsp
    var runtimeinfolist []RuntimeInfo

    var upgrader = websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true
        }} //resolve request origin not allowed
    args := []string{}
    var mess map[string]interface{}

    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        L.System.Fatal("upgrade:", err)
        return
    }
    defer c.Close()

    _, message, err := c.ReadMessage()
    if err != nil {
        L.System.Warning("read:", err)
    }
    L.System.Debug("recv from client: %s", message)
    fmt.Println(string(message))
    if e := json.Unmarshal(message, &mess); e != nil {
        rsp.Result = "ERROR"
        rsp.Errors = e.Error()
        return
    }
    seconds, _ := strconv.Atoi(handleInterfaceString(mess["timeinterval"]))

    go func() {
        //for {
        // time.Sleep(time.Millisecond * 1)
        _, message, err := c.ReadMessage()
        if err != nil {
            L.System.Debug("read:", err)
            return
        }
        L.System.Debug("recv from client: %s", message)
        //}
    }()

    for {
        // cmd
        //saltClient, _ := saltgo.NewClient(handleInterfaceString(mess["ip"]))
        saltClient, _ := saltgo.NewClient("192.168.146.128")
        args = append(args, " cat /proc/uptime|awk -F. '{print $1}' ")
        ret, _ := saltClient.Reader("cmd.run", args)
        str := handleInterfaceString(ret)
        secds, _ := strconv.ParseInt(str, 10, 64)

        // response
        runtimeinfolist = append(runtimeinfolist, RuntimeInfo{Runtime: TimeProcess(secds)})

        rsp.Result = "OK"
        rsp.RuntimeInfoList = runtimeinfolist
        buf, _ := json.Marshal(&rsp)

        err := c.WriteMessage(websocket.TextMessage, buf)
        if err != nil {
            L.System.Warning("write:", err)
            return
        }
        runtimeinfolist = make([]RuntimeInfo, 0)
        args = make([]string, 0)
        ret = nil
        time.Sleep(time.Duration(seconds) * time.Second)
    }
}

func main() {
    flag.Parse()
    log.SetFlags(0)

    http.HandleFunc("/echo", echo)
    http.HandleFunc("/cpuinfolist", NewProcesCpuInfoList)
    http.HandleFunc("/memoryinfolist", NewProcessMemoryInfoList)
    http.HandleFunc("/interfaceflowlist", NewProcessInterfaceFlowList)
    http.HandleFunc("/sysruntimeinfolist", ProcessSysRunTimeInfoList)
    //log.Fatal(http.ListenAndServe(*addr, nil))
}
