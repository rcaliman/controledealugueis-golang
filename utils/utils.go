package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func TimeParse(s string) time.Time {
	data := strings.Split(s, "-")
	dia, _ := strconv.Atoi(data[2])
	mes := func() time.Month {
		m, _ := strconv.Atoi(data[1])
		return time.Month(m)
	}()
	ano, _ := strconv.Atoi(data[0])
	novaData := time.Date(ano, mes, dia, 0, 0, 0, 0, time.Local)
	log.Println("data", data)
	return novaData
}

func FormataDataUS(s time.Time) string {
	timeStamp := s.String()
	data := strings.Split(timeStamp, " ")[0]
	return data
}

func FormataDataBR(s time.Time) string {
	timeStamp := s.String()
	data := strings.Split(timeStamp, " ")[0]
	dataFatiada := strings.Split(data, "-")
	novaData := fmt.Sprintf("%s/%s/%s", dataFatiada[2], dataFatiada[1], dataFatiada[0])
	return novaData
}

func FormataFloat(v float64) string {
	s := fmt.Sprintf("%.2f", v)
	stringFormatada := strings.Replace(s, ".", ",", 1)
	return stringFormatada
}
