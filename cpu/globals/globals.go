package globals

type Config struct {
	Puerto           string `json:"puerto"`
	IPMemory         string `json:"ip_memory"`
	PortMemory       string `json:"port_memory"`
	NumberFellingTLB int    `json:"number_felling_tlb"`
	AlgorithmTLB     string `json:"algorithm_tlb"`
}

var ClientConfig *Config
