# GoRTSPBrute - Bruteforce de rutas RTSP

RTSPBrute es una herramienta escrita en Go para realizar ataques de fuerza bruta sobre rutas RTSP en cámaras IP y servidores de streaming. Utiliza el método **DESCRIBE** para verificar si realmente existe un stream en la ruta probada, evitando falsos positivos.

## Características
✅ Escaneo rápido y multihilo.  
✅ Utiliza el método **DESCRIBE** en lugar de **OPTIONS**, mejorando la precisión.  
✅ Filtra rutas falsas verificando que la respuesta contenga información de stream.  
✅ Compatible con servidores RTSP estándar.

## Requisitos
- Go 1.18 o superior

## Instalación
Clona el repositorio y compila el binario:
```bash
git clone https://github.com/vpanal/rtspbrute.git
cd rtspbrute
go build main.go -o gortspbrute
```

## Uso
Ejecuta el binario indicando la IP:PUERTO del servidor RTSP y la wordlist de rutas:
```bash
./gortspbrute 192.168.1.54:8554 wordlist.txt
```

Ejemplo de wordlist:
```
stream
live.sdp
video
cam1
axis-media/media.amp
```

## Ejemplo de salida
```bash
[+] Stream válido encontrado: rtsp://192.168.1.54:8554/live.sdp
[-] Ruta inválida o sin stream: rtsp://192.168.1.54:8554/test
[+] Stream válido encontrado: rtsp://192.168.1.54:8554/axis-media/media.amp
Bruteforce finalizado.
```

## Notas
- Si una ruta devuelve **401 Unauthorized**, significa que es válida pero requiere autenticación.
- Si VLC no puede abrir la URL, es probable que el stream no sea accesible.

## Autor
Creado por vpanal para fines educativos y pruebas de seguridad en entornos controlados. No me hago responsable del uso indebido.

