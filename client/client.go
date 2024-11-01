package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// peer ile bağlantı kurmaya çalışıyoruz
	/*
		net.Dial(), belirtilen adrese (localhost:8080) TCP bağlantısı kurar.
		localhost: Aynı makineyi belirtir.
		8080: Bağlanılacak port numarası.
	*/
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Bağlantı Hatası", err)
		os.Exit(1)
	}

	defer conn.Close()

	// Mesaj Okuma
	/*
		bufio.NewReader: Bağlantıdan veri okumak için bufio.Reader nesnesi oluşturur.
		ReadString('\n'): conn nesnesinden \n karakterine kadar olan veriyi okur. Sunucunun gönderdiği mesajı (Merhaba Peer!\n) alırız.
	*/
	message, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Alınan Mesaj:", message)
}
