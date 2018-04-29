package main

import _"fmt"

type marketCurrency struct {
	Success bool        `json:"Success"`
	Message interface{} `json:"Message"`
	Data    []Data      `json:"Data"`
	Error interface{}   `json:"Error"`
}

func (m marketCurrency) getCurrency(Symbol string) Data {
  var currency Data

	for index := range m.Data{
					if m.Data[index].Symbol == Symbol {
						currency = m.Data[index]
					}
	    }
		return currency
}
