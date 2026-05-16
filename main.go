package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleConnection(conn net.Conn){
	defer conn.Close();
	for{
		conn.SetReadDeadline(time.Now().Add(time.Second))
		buff:= make([]byte ,1024);
		n ,err:= conn.Read(buff);

		if err!=nil{
			if netErr, ok := err.(net.Error); ok && netErr.Timeout(){
				log.Println("Time out error" ,err);
				continue;
			}else{
				fmt.Println("Error in connection: " ,err);
				break;
			}
		}

		reqPath := string(buff[:n]);

		if reqPath == "/api"{
			
		}

		fmt.Println("variable val: ", reqPath);

		fmt.Println("Recieved data: ", string(buff[:n]));
	}
}

func main(){
	ln ,err:= net.Listen("tcp", ":8080");

	if err!=nil{
		fmt.Println("error in starting server:", err);
		return
	}

	fmt.Println("Tcp server started at post 8080");

	for{
		conn , err:= ln.Accept();

		if err!= nil{
			fmt.Println("error in accept connection:", err);
			continue;
		}

		go handleConnection(conn)

	}
}