package main

import (
	"fmt"
	"net"
)

func main() {
	// sunucu rolünde bir bağlantı dinleyici başlatıyoruz

	// net.Listen fonksiyonu, belirtilen protokol ve adreste (tcp, :8080) bir ağ dinleyicisi oluşturur.
	// :8080: Dinlenecek port numarası. Bir bağlantı isteği geldiğinde bu port üzerinden kabul edilecek.
	listiner, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("dinleme başarısız", err)
		return
	}
	//defer anahtar kelimesi, listener.Close() fonksiyonunu program sonlandığında otomatik olarak çalıştırır.
	// Böylece program sona erdiğinde dinleyici kapatılır ve port boşaltılır.
	defer listiner.Close()
	fmt.Println("Sunucu başlatıldı, bağlantılar bekleniyor...")

	for {
		/*
		   listener.Accept(), dinleyiciyi kullanarak gelen bağlantıları kabul eder. Yeni bir bağlantı geldiğinde net.Conn türünde bir conn nesnesi döner.
		   net.Conn: İki nokta arasındaki bağlantıyı temsil eder; veri gönderme ve alma işlemlerini bu nesne üzerinden yapabiliriz.
		   Döngü içinde her yeni bağlantı geldiğinde, yeni bir conn nesnesi oluşturulur.
		*/
		conn, err := listiner.Accept()

		if err != nil {
			fmt.Println("Bağlantı kabul hatası", err)
			continue
		}
		/*
		   go anahtar kelimesiyle bir goroutine başlatılır. Bu, handleConnection fonksiyonunun paralel olarak çalışmasını sağlar.
		   Her bağlantı kendi goroutine’inde işlenir, böylece aynı anda birden fazla bağlantıyı kabul edebiliriz.
		*/
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// RemoteAddr: Bağlanan istemcinin IP adresini ve portunu döner.
	fmt.Println("yeni bir bağlantı kabul edildi", conn.RemoteAddr())
	conn.Write([]byte("Merhaba peer\n"))

}
