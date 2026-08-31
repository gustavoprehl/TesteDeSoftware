// Suíte gerada pelo ChatGPT com o prompt "Chain-of-Thought (CoT)" — o modelo
// foi instruído a listar tipos de item, regras de negócio e limites antes de
// escrever o código de teste. Ver Registro_Prompts.pdf, seção 2.
//
// Única alteração feita para integrar ao repositório: adaptado de white-box
// (package gildedrose) para black-box (package chain_of_thought_test),
// qualificando Item e UpdateQuality com o prefixo do pacote. Nenhuma
// asserção foi alterada.
package chain_of_thought_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestNormalItem(t *testing.T) {
	tests := []struct {
		name         string
		item         gildedrose.Item
		expectedSell int
		expectedQual int
	}{
		{
			name:         "normal item before expiration",
			item:         gildedrose.Item{Name: "Elixir of the Mongoose", SellIn: 10, Quality: 20},
			expectedSell: 9,
			expectedQual: 19,
		},
		{
			name:         "normal item with SellIn zero",
			item:         gildedrose.Item{Name: "Elixir of the Mongoose", SellIn: 0, Quality: 20},
			expectedSell: -1,
			expectedQual: 18,
		},
		{
			name:         "normal item after expiration",
			item:         gildedrose.Item{Name: "Elixir of the Mongoose", SellIn: -1, Quality: 20},
			expectedSell: -2,
			expectedQual: 18,
		},
		{
			name:         "normal item with zero quality",
			item:         gildedrose.Item{Name: "Elixir of the Mongoose", SellIn: 10, Quality: 0},
			expectedSell: 9,
			expectedQual: 0,
		},
		{
			name:         "normal item with quality 50",
			item:         gildedrose.Item{Name: "Elixir of the Mongoose", SellIn: 10, Quality: 50},
			expectedSell: 9,
			expectedQual: 49,
		},
		{
			name:         "unknown item is treated as normal",
			item:         gildedrose.Item{Name: "Unknown Item", SellIn: 10, Quality: 20},
			expectedSell: 9,
			expectedQual: 19,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := tt.item

			gildedrose.UpdateQuality([]*gildedrose.Item{&item})

			if item.SellIn != tt.expectedSell {
				t.Errorf(
					"SellIn = %d, want %d",
					item.SellIn,
					tt.expectedSell,
				)
			}

			if item.Quality != tt.expectedQual {
				t.Errorf(
					"Quality = %d, want %d",
					item.Quality,
					tt.expectedQual,
				)
			}
		})
	}
}

func TestAgedBrie(t *testing.T) {
	tests := []struct {
		name         string
		item         gildedrose.Item
		expectedSell int
		expectedQual int
	}{
		{
			name:         "increases quality before expiration",
			item:         gildedrose.Item{Name: "Aged Brie", SellIn: 10, Quality: 20},
			expectedSell: 9,
			expectedQual: 21,
		},
		{
			name:         "SellIn zero gets normal increase and expiration increase",
			item:         gildedrose.Item{Name: "Aged Brie", SellIn: 0, Quality: 20},
			expectedSell: -1,
			expectedQual: 22,
		},
		{
			name:         "increases quality after expiration",
			item:         gildedrose.Item{Name: "Aged Brie", SellIn: -1, Quality: 20},
			expectedSell: -2,
			expectedQual: 22,
		},
		{
			name:         "quality zero",
			item:         gildedrose.Item{Name: "Aged Brie", SellIn: 10, Quality: 0},
			expectedSell: 9,
			expectedQual: 1,
		},
		{
			name:         "quality 49 reaches maximum",
			item:         gildedrose.Item{Name: "Aged Brie", SellIn: 10, Quality: 49},
			expectedSell: 9,
			expectedQual: 50,
		},
		{
			name:         "quality 50 stays at maximum",
			item:         gildedrose.Item{Name: "Aged Brie", SellIn: 10, Quality: 50},
			expectedSell: 9,
			expectedQual: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := tt.item

			gildedrose.UpdateQuality([]*gildedrose.Item{&item})

			if item.SellIn != tt.expectedSell {
				t.Errorf(
					"SellIn = %d, want %d",
					item.SellIn,
					tt.expectedSell,
				)
			}

			if item.Quality != tt.expectedQual {
				t.Errorf(
					"Quality = %d, want %d",
					item.Quality,
					tt.expectedQual,
				)
			}
		})
	}
}

func TestSulfuras(t *testing.T) {
	tests := []struct {
		name         string
		item         gildedrose.Item
		expectedSell int
		expectedQual int
	}{
		{
			name: "does not change",
			item: gildedrose.Item{
				Name:    "Sulfuras, Hand of Ragnaros",
				SellIn:  10,
				Quality: 80,
			},
			expectedSell: 10,
			expectedQual: 80,
		},
		{
			name: "does not change with SellIn zero",
			item: gildedrose.Item{
				Name:    "Sulfuras, Hand of Ragnaros",
				SellIn:  0,
				Quality: 80,
			},
			expectedSell: 0,
			expectedQual: 80,
		},
		{
			name: "does not change after expiration",
			item: gildedrose.Item{
				Name:    "Sulfuras, Hand of Ragnaros",
				SellIn:  -10,
				Quality: 80,
			},
			expectedSell: -10,
			expectedQual: 80,
		},
		{
			name: "quality above 50 is preserved",
			item: gildedrose.Item{
				Name:    "Sulfuras, Hand of Ragnaros",
				SellIn:  10,
				Quality: 100,
			},
			expectedSell: 10,
			expectedQual: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := tt.item

			gildedrose.UpdateQuality([]*gildedrose.Item{&item})

			if item.SellIn != tt.expectedSell {
				t.Errorf(
					"SellIn = %d, want %d",
					item.SellIn,
					tt.expectedSell,
				)
			}

			if item.Quality != tt.expectedQual {
				t.Errorf(
					"Quality = %d, want %d",
					item.Quality,
					tt.expectedQual,
				)
			}
		})
	}
}

func TestBackstagePasses(t *testing.T) {
	tests := []struct {
		name         string
		item         gildedrose.Item
		expectedSell int
		expectedQual int
	}{
		{
			name: "more than 10 days remaining increases quality by 1",
			item: gildedrose.Item{
				Name:    "Backstage passes to a TAFKAL80ETC concert",
				SellIn:  11,
				Quality: 20,
			},
			expectedSell: 10,
			expectedQual: 21,
		},
		{
			name: "10 days remaining increases quality by 2",
			item: gildedrose.Item{
				Name:    "Backstage passes to a TAFKAL80ETC concert",
				SellIn:  10,
				Quality: 20,
			},
			expectedSell: 9,
			expectedQual: 22,
		},
		{
			name: "6 days remaining increases quality by 2",
			item: gildedrose.Item{
				Name:    "Backstage passes to a TAFKAL80ETC concert",
				SellIn:  6,
				Quality: 20,
			},
			expectedSell: 5,
			expectedQual: 22,
		},
		{
			name: "5 days remaining increases quality by 3",
			item: gildedrose.Item{
				Name:    "Backstage passes to a TAFKAL80ETC concert",
				SellIn:  5,
				Quality: 20,
			},
			expectedSell: 4,
			expectedQual: 23,
		},
		{
			name: "quality 49 is capped at 50",
			item: gildedrose.Item{
				Name:    "Backstage passes to a TAFKAL80ETC concert",
				SellIn:  5,
				Quality: 49,
			},
			expectedSell: 4,
			expectedQual: 50,
		},
		{
			name: "quality 50 stays at 50",
			item: gildedrose.Item{
				Name:    "Backstage passes to a TAFKAL80ETC concert",
				SellIn:  5,
				Quality: 50,
			},
			expectedSell: 4,
			expectedQual: 50,
		},
		{
			name: "quality zero increases normally",
			item: gildedrose.Item{
				Name:    "Backstage passes to a TAFKAL80ETC concert",
				SellIn:  11,
				Quality: 0,
			},
			expectedSell: 10,
			expectedQual: 1,
		},
		{
			name: "SellIn zero increases then becomes zero after expiration",
			item: gildedrose.Item{
				Name:    "Backstage passes to a TAFKAL80ETC concert",
				SellIn:  0,
				Quality: 20,
			},
			expectedSell: -1,
			expectedQual: 0,
		},
		{
			name: "expired backstage pass has zero quality",
			item: gildedrose.Item{
				Name:    "Backstage passes to a TAFKAL80ETC concert",
				SellIn:  -1,
				Quality: 20,
			},
			expectedSell: -2,
			expectedQual: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := tt.item

			gildedrose.UpdateQuality([]*gildedrose.Item{&item})

			if item.SellIn != tt.expectedSell {
				t.Errorf(
					"SellIn = %d, want %d",
					item.SellIn,
					tt.expectedSell,
				)
			}

			if item.Quality != tt.expectedQual {
				t.Errorf(
					"Quality = %d, want %d",
					item.Quality,
					tt.expectedQual,
				)
			}
		})
	}
}

func TestUpdateQualityMultipleItems(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    "Elixir of the Mongoose",
			SellIn:  10,
			Quality: 20,
		},
		{
			Name:    "Aged Brie",
			SellIn:  10,
			Quality: 20,
		},
		{
			Name:    "Sulfuras, Hand of Ragnaros",
			SellIn:  10,
			Quality: 80,
		},
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  5,
			Quality: 20,
		},
	}

	gildedrose.UpdateQuality(items)

	expected := []gildedrose.Item{
		{
			Name:    "Elixir of the Mongoose",
			SellIn:  9,
			Quality: 19,
		},
		{
			Name:    "Aged Brie",
			SellIn:  9,
			Quality: 21,
		},
		{
			Name:    "Sulfuras, Hand of Ragnaros",
			SellIn:  10,
			Quality: 80,
		},
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  4,
			Quality: 23,
		},
	}

	for i, item := range items {
		if item.SellIn != expected[i].SellIn {
			t.Errorf(
				"item %d SellIn = %d, want %d",
				i,
				item.SellIn,
				expected[i].SellIn,
			)
		}

		if item.Quality != expected[i].Quality {
			t.Errorf(
				"item %d Quality = %d, want %d",
				i,
				item.Quality,
				expected[i].Quality,
			)
		}
	}
}

func TestUpdateQualityEmptyItems(t *testing.T) {
	var items []*gildedrose.Item

	gildedrose.UpdateQuality(items)
}
