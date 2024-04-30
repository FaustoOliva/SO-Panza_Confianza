package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/sisoputnfrba/tp-golang/cpu/globals"
	"github.com/sisoputnfrba/tp-golang/cpu/utils"
)

func main() {
	utils.ConfigurarLogger()
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}
	puerto := globals.ClientConfig.Puerto

	http.HandleFunc("GET /input", utils.Prueba)
	http.HandleFunc("/savePCB", utils.ProcessSavedPCBFromKernel)
	http.ListenAndServe(":"+strconv.Itoa(puerto), nil)
}
