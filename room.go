package main

import (
  "log"
  "net"
  "fmt"
  "bufio"
  "github.com/spf13/viper"

)

func main(){

  viper.SetConfigFile("config.json")
  viper.AddConfigPath(".")
  viper.SetConfigName("config")
  viper.ReadInConfig()
  gpServer_port := viper.GetString("gpServer.port")
  gpServer_host := viper.GetString("gpServer.host")
  gpServer_name := viper.GetString("gpServer.GroupName")
  gpServer_bio := viper.GetString("gpServer.Bio")

  fmt.Println("Welcome to ",gpServer_name)
  fmt.Println("This Room Powered By ZenChat")
  fmt.Println("Ip :",gpServer_host+gpServer_port)
  fmt.Println(gpServer_bio)
  fmt.Println("----------------------------------------------")
  for {
    conn, err := net.Dial("tcp", gpServer_host+gpServer_port)
    if err != nil {
      fmt.Println("[ERROR 400 ] Bad Request The server cannot or will not process the request due to an apparent client error ")
    }
    var (
      buf = make([]byte, 1024)
      r   = bufio.NewReader(conn)
    )

    n_text, _ := r.Read(buf)
    data_text := string(buf[:n_text])

    n_User_name, _ := r.Read(buf)
    data_User_name := string(buf[:n_User_name])


    log.Println(data_User_name, data_text)

    }

}
