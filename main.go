package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

// Función para probar una ruta RTSP
func checkRTSP(target, route string, wg *sync.WaitGroup) {
	defer wg.Done()

	rtspURL := fmt.Sprintf("rtsp://%s/%s", target, route)
	conn, err := net.DialTimeout("tcp", target, 5*time.Second)
	if err != nil {
		fmt.Printf("[-] No conexión con %s\n", rtspURL)
		return
	}
	defer conn.Close()

	// Enviar una petición DESCRIBE para verificar si el stream es válido
	request := fmt.Sprintf("DESCRIBE rtsp://%s/%s RTSP/1.0\r\nCSeq: 2\r\nUser-Agent: VLC/3.0.11\r\nAccept: application/sdp\r\n\r\n", target, route)
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Printf("[-] Error al enviar petición a %s\n", rtspURL)
		return
	}

	response := make([]byte, 4096)
	n, err := conn.Read(response)
	if err != nil || n == 0 {
		fmt.Printf("[-] Sin respuesta válida de %s\n", rtspURL)
		return
	}

	responseStr := string(response[:n])

	// Comprobamos si la respuesta tiene "RTSP/1.0 200 OK" y si devuelve datos SDP
	if strings.Contains(responseStr, "RTSP/1.0 200 OK") && strings.Contains(responseStr, "m=video") {
		fmt.Printf("[+] Stream válido encontrado: %s\n", rtspURL)
	} else {
		fmt.Printf("[-] Ruta inválida o sin stream: %s\n", rtspURL)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: go run main.go <IP:PUERTO> <WORDLIST>")
		os.Exit(1)
	}

	target := os.Args[1]
	wordlistPath := os.Args[2]
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("[-] Error al abrir la wordlist:", err)
		os.Exit(1)
	}
	defer file.Close()

	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		route := scanner.Text()
		wg.Add(1)
		go checkRTSP(target, route, &wg)
	}

	wg.Wait()
	fmt.Println("Bruteforce finalizado.")
}

