package main

import (
	"fmt"
	"io"
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

		fmt.Println("variable val: ", reqPath);

		fmt.Println("Recieved data: ", string(buff[:n]));

		targetConn, err := net.Dial("tcp", reqPath);

		if err!= nil{
			log.Println("Error during request client api: " ,err);
		}

		defer targetConn.Close()

		errChan := make(chan error ,2);

		go func (){
			_ ,err:= io.Copy(targetConn , conn);
			// if err != nil{
			// 	fmt.Println("error occured during copying from tatget to conn");
			// }
			errChan <-err
		}()

		go func ()  {
			_ ,err := io.Copy(conn , targetConn);
			
			errChan <-err
		}()

		<-errChan
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