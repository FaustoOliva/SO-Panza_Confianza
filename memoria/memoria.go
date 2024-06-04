package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/sisoputnfrba/tp-golang/memoria/globals"
	"github.com/sisoputnfrba/tp-golang/memoria/utils"
)

func main() {
	utils.ConfigurarLogger()
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}

	puerto := globals.ClientConfig.Puerto

	http.HandleFunc("POST /setInstructionFromFileToMap", utils.SetInstructionsFromFileToMap)
	http.HandleFunc("GET  /getInstructionFromPid", utils.GetInstruction)
	http.HandleFunc("POST /createProcess", utils.CreateProcessHandler)
	http.HandleFunc("POST /terminateProcess", utils.TerminateProcessHandler)
	http.HandleFunc("POST /resizeProcess", utils.ResizeProcessHandler)
	http.HandleFunc("POST /readMemory", utils.ReadMemoryHandler)
	http.HandleFunc("POST /writeMemory", utils.WriteMemoryHandler)

	http.ListenAndServe(":"+strconv.Itoa(puerto), nil)
}
