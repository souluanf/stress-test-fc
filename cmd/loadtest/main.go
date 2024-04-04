package main

import (
	"flag"
	"github.com/souluanf/stress-test-fc/internal/httpclient"
	"github.com/souluanf/stress-test-fc/internal/report"
	"log"
	"os"
	"time"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 0, "Número total de requests")
	concurrency := flag.Int("concurrency", 0, "Número de chamadas simultâneas")
	flag.Parse()
	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	client := httpclient.NewClient(*concurrency)
	startTime := time.Now()
	results, err := client.LoadTest(*url, *requests)
	if err != nil {
		log.Fatalf("Erro ao executar testes de carga: %v", err)
	}
	report.Generate(startTime, time.Now(), *requests, results)
}
