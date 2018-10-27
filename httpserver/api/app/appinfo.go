package app

import (
	"github.com/gin-gonic/gin"
	"dep/commd"
	"dep/model"
	"net/http"
	"dep/module"
	"github.com/gorilla/websocket"
	"html/template"
	"time"
)

func Info(c *gin.Context) {
	db := commd.GetDefaultDB()
	db.LogMode(true)
	var apps []model.Apps
	db.Find(&apps)

	c.JSON(http.StatusOK, module.ApiResp{
		ErrNo:  commd.SuccesCode,
		ErrMsg: "",
		Data:   apps,
	})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Echo(c *gin.Context) {
	// 升级 get 请求为 webSocket 协议
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		commd.Logger.Error("upgrader:", err)
		return
	}
	defer ws.Close()
	for {
		yh := model.User{Name: "you hao", Age: 30, Ip: ws.RemoteAddr().String()}
		err = ws.WriteJSON(yh)
		if err != nil {
			commd.Logger.Error("write:", err)
			return
		}
		time.Sleep(1 * time.Second)
		// 读取 ws 中的数据
		//mt, message, err := ws.ReadMessage()
		//if err != nil {
		//	break
		//}
		//if string(message) == "ping" {
		//	message = []byte("pong")
		//}
		// 写入 ws 数据
		//err = ws.WriteMessage(mt, message)
		//if err != nil {
		//	commd.Logger.Error("write:", err)
		//	break
		//}
	}
}


var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  

//var ws = new WebSocket("ws://localhost:10000/echo");  
//// 连接打开时触发 
//ws.onopen = function(evt) {  
//    console.log("Connection open ...");  
//    ws.send("Hello WebSockets!");};  
//// 接收到消息时触发  
//ws.onmessage = function(evt) {  
//    console.log("Received Message:" + evt.data);};  
//// 连接关闭时触发  
//ws.onclose = function(evt) {  
//    console.log("Connection closed.");}; 

window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))

func Home(c *gin.Context){
	homeTemplate.Execute(c.Writer, "ws://"+c.Request.Host+"/echo")
}

