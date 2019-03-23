package main
import (
  "net"
  "fmt"
  "bufio"
  "log"
  "github.com/spf13/viper"
)

func accept_conn (){
  fmt.Println("Powered By ZenChat")
  log.Println("[*]locate Config.json File")
  viper.SetConfigFile("config.json")
  viper.AddConfigPath(".")
  viper.SetConfigName("config")
  viper.ReadInConfig()
  log.Printf("[*]Using config: %s\n", viper.ConfigFileUsed())
  mainServer_host := viper.GetString("mainServer.host")
  mainServer_port := viper.GetString("mainServer.port")
  gpServer_host := viper.GetString("gpServer.host")
  gpServer_port := viper.GetString("gpServer.port")

  log.Println("[*]Lunching Main Server",mainServer_host+mainServer_port)
  ln, err := net.Listen("tcp", mainServer_port)
  if err != nil{
    log.Println("[ERROR 502] Bad Gateway : Cant Create MainServer")
  }
  log.Println("[*] Lunching Group Server",gpServer_host+gpServer_port)
  gp_conn, err := net.Listen("tcp", gpServer_port)
  if err != nil{
    fmt.Println("[ERROR 502] Bad Gateway : Cant Create GroupServer")
  }

  for {
    conn, err := ln.Accept()
    if err != nil{
      fmt.Println("[ERROR 400 ] Bad Request : Cant Accept Group Requests")
    }
    fmt.Println(conn,"Send request to Server")

    var (
      buf = make([]byte, 1024)
      r   = bufio.NewReader(conn)
    )

    send_conn_accept, err := gp_conn.Accept()
    if err !=nil{
      fmt.Println("[ERROR 400 ] Bad Request : Cant Accept Clients Requests")
    }
    n_User_name, err := r.Read(buf)
    if err != nil{
      fmt.Println("[ERROR] Cant Recive Data From Clients")
    }
    data_User_name := string(buf[:n_User_name])

    n_msg, err := r.Read(buf)
    if err != nil{
      fmt.Println("[ERROR] Cant Recive Data From Clients")
    }
    data_msg := string(buf[:n_msg])
    log.Println("Receive From [",data_User_name,"]:", data_msg)

    User_name := data_User_name
    text := data_msg
    send_conn_accept.Write([]byte(User_name))

    send_conn_accept.Write([]byte(text))
    //log.Printf("Send: %s", text)
    send_conn_accept.Close()

    }
  }

func main(){
  accept_conn()


}
