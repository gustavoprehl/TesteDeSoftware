// Suíte gerada pela IA na Iteração 2 com o "Prompt Estruturado" (caixa branca):
// o prompt entregou o CFG, as 17 decisões D1–D17 e a meta de 100% de
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

const (
	agedBrie  = "Aged Brie"
	backstage = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras  = "Sulfuras, Hand of Ragnaros"
	comum     = "Item comum"
)

// TestUpdateQuality_CoberturaDeDecisao percorre os 34 ramos (D1T..D17F) do
// grafo de fluxo de controle de UpdateQuality. Cada linha da tabela é UM item
// em UMA chamada. Os valores esperados caracterizam o comportamento atual do
// código legado (não corrigem nem pressupõem regras de negócio).
func TestUpdateQuality_CoberturaDeDecisao(t *testing.T) {
	casos := []struct {
		id              string
		ramos           string
		name            string
		sellIn          int
		quality         int
		esperadoSellIn  int
		esperadoQuality int
	}{
		{
			id:    "T1",
			ramos: "D1T D1F D2T D3T D4F D11F D12T D13T D14T D15T D16F",
			name:  sulfuras, sellIn: -1, quality: 80,
			esperadoSellIn: -1, esperadoQuality: 80,
		},
		{
			id:    "T2",
			ramos: "D1T D1F D2T D3F D11T D12T D13T D14T D15F",
			name:  comum, sellIn: 0, quality: 0,
			esperadoSellIn: -1, esperadoQuality: 0,
		},
		{
			id:    "T3",
			ramos: "D1T D1F D2T D3T D4T D11T D12T D13T D14T D15T D16T",
			name:  comum, sellIn: 0, quality: 10,
			esperadoSellIn: -1, esperadoQuality: 8,
		},
		{
			id:    "T4",
			ramos: "D1T D1F D2F D5F D11T D12T D13F D17F",
			name:  agedBrie, sellIn: 0, quality: 50,
			esperadoSellIn: -1, esperadoQuality: 50,
		},
		{
			id:    "T5",
			ramos: "D1T D1F D2F D5T D6F D11T D12T D13F D17T",
			name:  agedBrie, sellIn: 0, quality: 10,
			esperadoSellIn: -1, esperadoQuality: 12,
		},
		{
			id:    "T6",
			ramos: "D1T D1F D2F D5T D6T D7F D9F D11T D12F",
			name:  backstage, sellIn: 11, quality: 10,
			esperadoSellIn: 10, esperadoQuality: 11,
		},
		{
			id:    "T7",
			ramos: "D1T D1F D2F D5T D6T D7T D8T D9T D10T D11T D12F",
			name:  backstage, sellIn: 5, quality: 10,
			esperadoSellIn: 4, esperadoQuality: 13,
		},
		{
			id:    "T8",
			ramos: "D1T D1F D2F D5T D6T D7T D8F D9T D10F D11T D12T D13T D14F",
			name:  backstage, sellIn: 0, quality: 49,
			esperadoSellIn: -1, esperadoQuality: 0,
		},
	}

	for _, c := range casos {
		c := c
		t.Run(c.id, func(t *testing.T) {
			item := &gildedrose.Item{Name: c.name, SellIn: c.sellIn, Quality: c.quality}

			gildedrose.UpdateQuality([]*gildedrose.Item{item})

			if item.SellIn != c.esperadoSellIn {
				t.Errorf("%s [%s] entrada (%q, SellIn=%d, Quality=%d): SellIn = %d, esperado %d",
					c.id, c.ramos, c.name, c.sellIn, c.quality, item.SellIn, c.esperadoSellIn)
			}
			if item.Quality != c.esperadoQuality {
				t.Errorf("%s [%s] entrada (%q, SellIn=%d, Quality=%d): Quality = %d, esperado %d",
					c.id, c.ramos, c.name, c.sellIn, c.quality, item.Quality, c.esperadoQuality)
			}
		})
	}
}
