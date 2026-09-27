// Validação do medidor de Cobertura de Decisão: os 8 casos manuais M1–M8 de
// docs/iteracao2/criterio-cobertura.md (seção 5) cobrem os 34 ramos, então
//
//	go run ./_decisioncov ./_decisioncov/validacao
//
// deve informar 34/34. Se o medidor informar outro número, ele está errado.
// Fica fora de ./... (pasta com "_"), não entra na execução normal dos testes.
package validacao_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestCasosManuaisM1aM8(t *testing.T) {
	casos := []struct {
		id, nome                string
		sellIn, quality         int
		wantSellIn, wantQuality int
	}{
		{"M1", "Elixir of the Mongoose", 0, 10, -1, 8},
		{"M2", "Elixir of the Mongoose", 0, 0, -1, 0},
		{"M3", "Sulfuras, Hand of Ragnaros", -1, 80, -1, 80},
		{"M4", "Aged Brie", 0, 10, -1, 12},
		{"M5", "Aged Brie", 0, 50, -1, 50},
		{"M6", "Backstage passes to a TAFKAL80ETC concert", 15, 10, 14, 11},
		{"M7", "Backstage passes to a TAFKAL80ETC concert", 3, 49, 2, 50},
		{"M8", "Backstage passes to a TAFKAL80ETC concert", 0, 10, -1, 0},
	}
	for _, c := range casos {
		t.Run(c.id, func(t *testing.T) {
			item := &gildedrose.Item{Name: c.nome, SellIn: c.sellIn, Quality: c.quality}
			gildedrose.UpdateQuality([]*gildedrose.Item{item})
			if item.SellIn != c.wantSellIn || item.Quality != c.wantQuality {
				t.Errorf("(%d, %d), esperado (%d, %d)", item.SellIn, item.Quality, c.wantSellIn, c.wantQuality)
			}
		})
	}
}
