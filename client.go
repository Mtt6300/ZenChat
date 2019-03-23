package main
import (
  "net"
  "fmt"
  "bufio"
  "os"
  "github.com/spf13/viper"

)

func main() {
  viper.SetConfigFile("config.json")
  viper.AddConfigPath(".")
  viper.SetConfigName("config")
  viper.ReadInConfig()
  gpServer_name := viper.GetString("gpServer.GroupName")
  gpServer_host := viper.GetString("gpServer.host")
  gpServer_port := viper.GetString("gpServer.port")
  mainServer_port := viper.GetString("mainServer.port")
  mainServer_host := viper.GetString("mainServer.host")

  User_name, _ := os.Hostname()
  fmt.Println("ZenChat Panel")
  fmt.Println("[*] Welcome",User_name,"to",gpServer_name)
  fmt.Println("IP :",gpServer_host+gpServer_port)
  fmt.Println("Type {/server help} and hit enter to see commands")
  fmt.Println("---------------------------------------------------")
  for {
    scanner := bufio.NewScanner(os.Stdin)
    var text_msg string
    fmt.Print(">>")
    scanner.Scan()
    text_msg = scanner.Text()
    if  text_msg == "/server exit"{
      os.Exit(3)
    } else if text_msg == "/server info"{
      fmt.Println("server ip is 46.225.114.57")
      fmt.Println("Name :",gpServer_name)
    } else if text_msg =="/server help"{
      fmt.Println("/server help       show commands")
      fmt.Println("/server info       show server information")
      fmt.Println("/server username       show your name")
      fmt.Println("/server exit       close server connection")

    } else if text_msg =="/server username"{
      fmt.Println("you are ",User_name)
    }
    text_msg1 := " : "+text_msg


    conn, err := net.Dial("tcp", mainServer_host+mainServer_port)
    if err !=nil{
      fmt.Println("[ERROR 400 ] Bad Request The server cannot or will not process the request due to an apparent client error ")
      os.Exit(3)
    }

    User_name1 := "("+User_name+")"
    conn.Write([]byte(User_name1))
    conn.Write([]byte(text_msg1))
    conn.Close()
  }
}
