package report

import (
	"fmt"
	"time"
)

func Generate(startTime, endTime time.Time, totalRequests int, results map[int]int) {
	fmt.Println("Relatório de Teste de Carga:")
	fmt.Println("------------------------------")
	fmt.Printf("Tempo total gasto na execução: %v\n", endTime.Sub(startTime))
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalRequests)
	var status200, otherStatus int
	for statusCode, count := range results {
		if statusCode == 200 {
			status200 += count
		} else {
			otherStatus += count
		}
	}
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", status200)
	fmt.Printf("Distribuição de outros códigos de status HTTP: %d\n", otherStatus)
}
