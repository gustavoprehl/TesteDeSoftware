// Suíte gerada pelo ChatGPT na Iteração 2 com o "Prompt Estruturado" (caixa
// branca): o prompt entregou o CFG, as 17 decisões D1–D17 e a meta de 100% de
// Cobertura de Decisão. Ver docs/iteracao2/prompt-estruturado.md (processo e
// resultados), docs/iteracao2/prompt-iteracao2.txt (prompt exato) e
// docs/iteracao2/resposta-ia-iteracao2.md (resposta completa da IA).
//
// Integrada ao repositório sem nenhuma alteração além deste comentário de
// cabeçalho e da formatação aplicada pelo gofmt (só espaços). O prompt já pedia
// package estruturado_test e o import do módulo, então não foi preciso adaptar
// nada, ao contrário das suítes da Iteração 1.
package estruturado_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestUpdateQualityCoberturaEstrutural(t *testing.T) {
	tests := []struct {
		id          string
		name        string
		sellIn      int
		quality     int
		wantSellIn  int
		wantQuality int
		ramos       string
	}{
		{
			id:          "T1",
			name:        "Backstage passes to a TAFKAL80ETC concert",
			sellIn:      0,
			quality:     49,
			wantSellIn:  -1,
			wantQuality: 0,
			ramos:       "D1T, D2F, D5T, D6T, D7T, D8F, D9T, D10F, D11T, D12T, D13T, D14F, D1F",
		},
		{
			id:          "T2",
			name:        "Backstage passes to a TAFKAL80ETC concert",
			sellIn:      5,
			quality:     10,
			wantSellIn:  4,
			wantQuality: 13,
			ramos:       "D1T, D2F, D5T, D6T, D7T, D8T, D9T, D10T, D11T, D12F, D1F",
		},
		{
			id:          "T3",
			name:        "Backstage passes to a TAFKAL80ETC concert",
			sellIn:      11,
			quality:     10,
			wantSellIn:  10,
			wantQuality: 11,
			ramos:       "D1T, D2F, D5T, D6T, D7F, D9F, D11T, D12F, D1F",
		},
		{
			id:          "T4",
			name:        "Aged Brie",
			sellIn:      0,
			quality:     0,
			wantSellIn:  -1,
			wantQuality: 2,
			ramos:       "D1T, D2F, D5T, D6F, D11T, D12T, D13F, D17T, D1F",
		},
		{
			id:          "T5",
			name:        "Aged Brie",
			sellIn:      0,
			quality:     50,
			wantSellIn:  -1,
			wantQuality: 50,
			ramos:       "D1T, D2F, D5F, D11T, D12T, D13F, D17F, D1F",
		},
		{
			id:          "T6",
			name:        "Sulfuras, Hand of Ragnaros",
			sellIn:      -1,
			quality:     10,
			wantSellIn:  -1,
			wantQuality: 10,
			ramos:       "D1T, D2T, D3T, D4F, D11F, D12T, D13T, D14T, D15T, D16F, D1F",
		},
		{
			id:          "T7",
			name:        "Normal Item",
			sellIn:      0,
			quality:     2,
			wantSellIn:  -1,
			wantQuality: 0,
			ramos:       "D1T, D2T, D3T, D4T, D11T, D12T, D13T, D14T, D15T, D16T, D1F",
		},
		{
			id:          "T8",
			name:        "Normal Item",
			sellIn:      0,
			quality:     0,
			wantSellIn:  -1,
			wantQuality: 0,
			ramos:       "D1T, D2T, D3F, D11T, D12T, D13T, D14T, D15F, D1F",
		},
	}

	for _, tc := range tests {
		t.Run(tc.id, func(t *testing.T) {
			item := &gildedrose.Item{
				Name:    tc.name,
				SellIn:  tc.sellIn,
				Quality: tc.quality,
			}

			gildedrose.UpdateQuality([]*gildedrose.Item{item})

			if item.SellIn != tc.wantSellIn {
				t.Errorf(
					"SellIn incorreto: entrada=(Name=%q, SellIn=%d, Quality=%d), obtido=%d, esperado=%d, ramos=%s",
					tc.name,
					tc.sellIn,
					tc.quality,
					item.SellIn,
					tc.wantSellIn,
					tc.ramos,
				)
			}

			if item.Quality != tc.wantQuality {
				t.Errorf(
					"Quality incorreto: entrada=(Name=%q, SellIn=%d, Quality=%d), obtido=%d, esperado=%d, ramos=%s",
					tc.name,
					tc.sellIn,
					tc.quality,
					item.Quality,
					tc.wantQuality,
					tc.ramos,
				)
			}
		})
	}
}
