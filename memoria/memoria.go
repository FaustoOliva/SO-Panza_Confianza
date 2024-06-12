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

	http.HandleFunc("GET /getInstructionFromPidASD", utils.GetInstruction)
	http.HandleFunc("POST /createProcess", utils.CreateProcessHandler)
	http.HandleFunc("POST /terminateProcess", utils.TerminateProcessHandler)
	http.HandleFunc("POST /resizeProcess", utils.ResizeProcessHandler)
	http.HandleFunc("POST /readMemory", utils.ReadMemoryHandler)
	http.HandleFunc("POST /writeMemory", utils.WriteMemoryHandler)
	http.HandleFunc("POST /getFramefromCPU", utils.GetPageFromCPU)                 //Recive la pagina desde "MMU" para devolver el frame
	http.HandleFunc("POST /SendInputSTDINToMemory", utils.RecieveInputSTDINFromIO) // Escribir en memoria el input de un proceso
	http.HandleFunc("POST /SendAdressSTDOUTToMemory", utils.RecieveAdressSTDOUTFromIO)
	http.HandleFunc("POST /SendPortOfInterfaceToMemory", utils.RecievePortOfInterfaceFromKernel) // Recive el puerto de la interfaz para despues saber a que interfaz mandar

	http.ListenAndServe(":"+strconv.Itoa(puerto), nil)
}
