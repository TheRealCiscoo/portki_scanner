<h1 align="center">Portki Scanner</h1>

<p align="center">Un escaner de puertos desarrollado en <b>Golang</b> que identifica los puertos abiertos en un servidor a través de un escaneo <b>TCP</b> simple.</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Networking-1BA0D7?style=for-the-badge&logo=openwrt&logoColor=white" alt="Networking" />
  <img src="https://img.shields.io/badge/Pentesting-000000?style=for-the-badge&logo=kalilinux&logoColor=white" alt="Pentesting" />
</p>

## Compilación
```bash
$ git clone https://github.com/TheRealCiscoo/portki_scanner

$ go build -ldflags="-s -w" -o .\builds\portki_scanner.exe .\main.go
```

## Uso
```bash
# .\portki_scanner.exe -t <TARGET_HOST> -r <PORT_RANGE>

$ .\portki_scanner.exe -t scanme.nmap.org -r 21-80
```

## Contáctame

- LinkedIn - [Francisco De Jesús](https://www.linkedin.com/in/franciscopauldejesus/)
- Github - [TheRealCiscoo](https://github.com/TheRealCiscoo)


## License

La licencia del proyecto es [MIT licensed](https://github.com/TheRealCiscoo/portki_scanner/blob/main/LICENSE).
