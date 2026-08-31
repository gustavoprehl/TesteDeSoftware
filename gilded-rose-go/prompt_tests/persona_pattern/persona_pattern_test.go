// Suíte gerada pelo ChatGPT com a técnica "Persona Pattern" — o modelo foi
// instruído a assumir o papel de um Engenheiro de QA Sênior cético.
// Ver Registro_Prompts.pdf, seção 3.
//
// Única alteração feita para integrar ao repositório: adaptado de white-box
// (package gildedrose) para black-box (package persona_pattern_test),
// qualificando Item e UpdateQuality com o prefixo do pacote. Nenhuma
// asserção foi alterada.
package persona_pattern_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestUpdateQuality_NormalItem(t *testing.T) {
	tests := []struct {
		name           string
		initialSellIn  int
		initialQuality int
		wantSellIn     int
		wantQuality    int
	}{
		{
			name:           "normal item before expiration",
			initialSellIn:  10,
			initialQuality: 20,
			wantSellIn:     9,
			wantQuality:    19,
		},
		{
			name:           "normal item reaches expiration",
			initialSellIn:  1,
			initialQuality: 20,
			wantSellIn:     0,
			wantQuality:    19,
		},
		{
			name:           "normal item expires",
			initialSellIn:  0,
			initialQuality: 20,
			wantSellIn:     -1,
			wantQuality:    18,
		},
		{
			name:           "normal item already expired",
			initialSellIn:  -1,
			initialQuality: 20,
			wantSellIn:     -2,
			wantQuality:    18,
		},
		{
			name:           "quality cannot become negative",
			initialSellIn:  10,
			initialQuality: 0,
			wantSellIn:     9,
			wantQuality:    0,
		},
		{
			name:           "expired item with zero quality",
			initialSellIn:  -1,
			initialQuality: 0,
			wantSellIn:     -2,
			wantQuality:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &gildedrose.Item{
				Name:    "Normal Item",
				SellIn:  tt.initialSellIn,
				Quality: tt.initialQuality,
			}

			gildedrose.UpdateQuality([]*gildedrose.Item{item})

			if item.SellIn != tt.wantSellIn {
				t.Errorf("SellIn = %d, want %d", item.SellIn, tt.wantSellIn)
			}

			if item.Quality != tt.wantQuality {
				t.Errorf("Quality = %d, want %d", item.Quality, tt.wantQuality)
			}
		})
	}
}

func TestUpdateQuality_AgedBrie(t *testing.T) {
	tests := []struct {
		name           string
		initialSellIn  int
		initialQuality int
		wantSellIn     int
		wantQuality    int
	}{
		{
			name:           "increases quality before expiration",
			initialSellIn:  10,
			initialQuality: 20,
			wantSellIn:     9,
			wantQuality:    21,
		},
		{
			name:           "increases twice when it expires",
			initialSellIn:  0,
			initialQuality: 20,
			wantSellIn:     -1,
			wantQuality:    22,
		},
		{
			name:           "increases twice when already expired",
			initialSellIn:  -1,
			initialQuality: 20,
			wantSellIn:     -2,
			wantQuality:    22,
		},
		{
			name:           "quality 49 reaches exactly 50",
			initialSellIn:  10,
			initialQuality: 49,
			wantSellIn:     9,
			wantQuality:    50,
		},
		{
			name:           "quality 50 does not increase",
			initialSellIn:  10,
			initialQuality: 50,
			wantSellIn:     9,
			wantQuality:    50,
		},
		{
			name:           "expired quality 49 reaches exactly 50",
			initialSellIn:  0,
			initialQuality: 49,
			wantSellIn:     -1,
			wantQuality:    50,
		},
		{
			name:           "expired quality 50 remains 50",
			initialSellIn:  0,
			initialQuality: 50,
			wantSellIn:     -1,
			wantQuality:    50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &gildedrose.Item{
				Name:    "Aged Brie",
				SellIn:  tt.initialSellIn,
				Quality: tt.initialQuality,
			}

			gildedrose.UpdateQuality([]*gildedrose.Item{item})

			if item.SellIn != tt.wantSellIn {
				t.Errorf("SellIn = %d, want %d", item.SellIn, tt.wantSellIn)
			}

			if item.Quality != tt.wantQuality {
				t.Errorf("Quality = %d, want %d", item.Quality, tt.wantQuality)
			}
		})
	}
}

func TestUpdateQuality_BackstagePasses(t *testing.T) {
	tests := []struct {
		name           string
		initialSellIn  int
		initialQuality int
		wantSellIn     int
		wantQuality    int
	}{
		{
			name:           "more than 10 days remaining",
			initialSellIn:  15,
			initialQuality: 20,
			wantSellIn:     14,
			wantQuality:    21,
		},
		{
			name:           "exactly 11 days remaining",
			initialSellIn:  11,
			initialQuality: 20,
			wantSellIn:     10,
			wantQuality:    21,
		},
		{
			name:           "exactly 10 days remaining",
			initialSellIn:  10,
			initialQuality: 20,
			wantSellIn:     9,
			wantQuality:    22,
		},
		{
			name:           "exactly 6 days remaining",
			initialSellIn:  6,
			initialQuality: 20,
			wantSellIn:     5,
			wantQuality:    22,
		},
		{
			name:           "exactly 5 days remaining",
			initialSellIn:  5,
			initialQuality: 20,
			wantSellIn:     4,
			wantQuality:    23,
		},
		{
			name:           "exactly 1 day remaining",
			initialSellIn:  1,
			initialQuality: 20,
			wantSellIn:     0,
			wantQuality:    23,
		},
		{
			name:           "expiration causes quality to drop to zero",
			initialSellIn:  0,
			initialQuality: 20,
			wantSellIn:     -1,
			wantQuality:    0,
		},
		{
			name:           "already expired causes quality to drop to zero",
			initialSellIn:  -1,
			initialQuality: 20,
			wantSellIn:     -2,
			wantQuality:    0,
		},
		{
			name:           "quality 49 with more than 10 days",
			initialSellIn:  15,
			initialQuality: 49,
			wantSellIn:     14,
			wantQuality:    50,
		},
		{
			name:           "quality 49 in double increase range",
			initialSellIn:  10,
			initialQuality: 49,
			wantSellIn:     9,
			wantQuality:    50,
		},
		{
			name:           "quality 49 in triple increase range",
			initialSellIn:  5,
			initialQuality: 49,
			wantSellIn:     4,
			wantQuality:    50,
		},
		{
			name:           "quality 50 never exceeds maximum",
			initialSellIn:  5,
			initialQuality: 50,
			wantSellIn:     4,
			wantQuality:    50,
		},
		{
			name:           "expired pass with zero quality stays zero",
			initialSellIn:  0,
			initialQuality: 0,
			wantSellIn:     -1,
			wantQuality:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &gildedrose.Item{
				Name:    "Backstage passes to a TAFKAL80ETC concert",
				SellIn:  tt.initialSellIn,
				Quality: tt.initialQuality,
			}

			gildedrose.UpdateQuality([]*gildedrose.Item{item})

			if item.SellIn != tt.wantSellIn {
				t.Errorf("SellIn = %d, want %d", item.SellIn, tt.wantSellIn)
			}

			if item.Quality != tt.wantQuality {
				t.Errorf("Quality = %d, want %d", item.Quality, tt.wantQuality)
			}
		})
	}
}

func TestUpdateQuality_Sulfuras(t *testing.T) {
	tests := []struct {
		name    string
		sellIn  int
		quality int
	}{
		{
			name:    "normal values",
			sellIn:  10,
			quality: 80,
		},
		{
			name:    "zero sell in",
			sellIn:  0,
			quality: 80,
		},
		{
			name:    "expired",
			sellIn:  -10,
			quality: 80,
		},
		{
			name:    "quality below 50",
			sellIn:  10,
			quality: 40,
		},
		{
			name:    "quality exactly 50",
			sellIn:  10,
			quality: 50,
		},
		{
			name:    "quality above 50",
			sellIn:  10,
			quality: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &gildedrose.Item{
				Name:    "Sulfuras, Hand of Ragnaros",
				SellIn:  tt.sellIn,
				Quality: tt.quality,
			}

			gildedrose.UpdateQuality([]*gildedrose.Item{item})

			if item.SellIn != tt.sellIn {
				t.Errorf("SellIn = %d, want %d", item.SellIn, tt.sellIn)
			}

			if item.Quality != tt.quality {
				t.Errorf("Quality = %d, want %d", item.Quality, tt.quality)
			}
		})
	}
}

func TestUpdateQuality_UnknownItem(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Completely New Item",
		SellIn:  10,
		Quality: 20,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.SellIn != 9 {
		t.Errorf("SellIn = %d, want 9", item.SellIn)
	}

	if item.Quality != 19 {
		t.Errorf("Quality = %d, want 19", item.Quality)
	}
}

func TestUpdateQuality_MultipleItems(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    "Normal Item",
			SellIn:  10,
			Quality: 20,
		},
		{
			Name:    "Aged Brie",
			SellIn:  10,
			Quality: 20,
		},
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  10,
			Quality: 20,
		},
		{
			Name:    "Sulfuras, Hand of Ragnaros",
			SellIn:  10,
			Quality: 80,
		},
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn != 9 || items[0].Quality != 19 {
		t.Errorf(
			"normal item: got SellIn=%d Quality=%d, want SellIn=9 Quality=19",
			items[0].SellIn,
			items[0].Quality,
		)
	}

	if items[1].SellIn != 9 || items[1].Quality != 21 {
		t.Errorf(
			"Aged Brie: got SellIn=%d Quality=%d, want SellIn=9 Quality=21",
			items[1].SellIn,
			items[1].Quality,
		)
	}

	if items[2].SellIn != 9 || items[2].Quality != 22 {
		t.Errorf(
			"Backstage Pass: got SellIn=%d Quality=%d, want SellIn=9 Quality=22",
			items[2].SellIn,
			items[2].Quality,
		)
	}

	if items[3].SellIn != 10 || items[3].Quality != 80 {
		t.Errorf(
			"Sulfuras: got SellIn=%d Quality=%d, want SellIn=10 Quality=80",
			items[3].SellIn,
			items[3].Quality,
		)
	}
}

func TestUpdateQuality_EmptyItems(t *testing.T) {
	// Garante que a função não falhe ao receber uma lista vazia.
	gildedrose.UpdateQuality([]*gildedrose.Item{})
}

func TestUpdateQuality_NilItems(t *testing.T) {
	// Slice nil é um input válido em Go e deve ser tratado
	// como uma lista vazia pelo loop.
	gildedrose.UpdateQuality(nil)
}
