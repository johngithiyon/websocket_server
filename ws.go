package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"

)

var upgrade = websocket.Upgrader {
	 CheckOrigin: func(r *http.Request) bool {
		return true 
	 },
}

var clients = make(map[*websocket.Conn]bool)


func handlews(w http.ResponseWriter,r *http.Request) {

      conn,err := upgrade.Upgrade(w,r,nil)

	  clients[conn] = true 

	  defer func() {
		conn.Close()
		delete(clients,conn)
	  }()

	  if err != nil {
		fmt.Println("Error",err)
		return 
	  }

	  for {
	     

		   _,message,err := conn.ReadMessage()

		   if err != nil {
			fmt.Println("Error",err)
			return
		   }


          for client := range clients {
			 err := client.WriteMessage(websocket.TextMessage,message)

			 if err != nil {
				fmt.Println("Write Error",err)
			 }
		  }


	  }
}

func main() {

	http.HandleFunc("/ws",handlews)
	fmt.Println("Server is Listening")
    http.ListenAndServe(":8090",nil)
}



