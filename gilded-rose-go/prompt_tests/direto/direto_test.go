// Suíte gerada pelo ChatGPT com o prompt "Direto (baseline)" — sem nenhuma
// técnica de prompt engineering, apenas o código colado e o pedido de testes.
// Ver Registro_Prompts.pdf, seção 1.
//
// Única alteração feita para integrar ao repositório: adaptado de white-box
// (package gildedrose) para black-box (package direto_test), qualificando
// Item e UpdateQuality com o prefixo do pacote. Nenhuma asserção foi alterada.
package direto_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestCommonItemQualityDecreases(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Elixir of the Mongoose",
		SellIn:  10,
		Quality: 20,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.SellIn != 9 {
		t.Errorf("expected SellIn 9, got %d", item.SellIn)
	}

	if item.Quality != 19 {
		t.Errorf("expected Quality 19, got %d", item.Quality)
	}
}

func TestCommonItemQualityDecreasesTwiceAfterExpiration(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Elixir of the Mongoose",
		SellIn:  0,
		Quality: 20,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.SellIn != -1 {
		t.Errorf("expected SellIn -1, got %d", item.SellIn)
	}

	if item.Quality != 18 {
		t.Errorf("expected Quality 18, got %d", item.Quality)
	}
}

func TestCommonItemQualityNeverNegative(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Elixir of the Mongoose",
		SellIn:  10,
		Quality: 0,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.Quality != 0 {
		t.Errorf("expected Quality 0, got %d", item.Quality)
	}
}

func TestAgedBrieQualityIncreases(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Aged Brie",
		SellIn:  10,
		Quality: 20,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.SellIn != 9 {
		t.Errorf("expected SellIn 9, got %d", item.SellIn)
	}

	if item.Quality != 21 {
		t.Errorf("expected Quality 21, got %d", item.Quality)
	}
}

func TestAgedBrieQualityIncreasesTwiceAfterExpiration(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Aged Brie",
		SellIn:  0,
		Quality: 20,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.SellIn != -1 {
		t.Errorf("expected SellIn -1, got %d", item.SellIn)
	}

	if item.Quality != 22 {
		t.Errorf("expected Quality 22, got %d", item.Quality)
	}
}

func TestAgedBrieQualityNeverExceeds50(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Aged Brie",
		SellIn:  10,
		Quality: 50,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.Quality != 50 {
		t.Errorf("expected Quality 50, got %d", item.Quality)
	}
}

func TestBackstagePassQualityIncreasesByOneWhenMoreThanTenDays(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		SellIn:  15,
		Quality: 20,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.SellIn != 14 {
		t.Errorf("expected SellIn 14, got %d", item.SellIn)
	}

	if item.Quality != 21 {
		t.Errorf("expected Quality 21, got %d", item.Quality)
	}
}

func TestBackstagePassQualityIncreasesByTwoWhenTenDaysOrLess(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		SellIn:  10,
		Quality: 20,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.Quality != 22 {
		t.Errorf("expected Quality 22, got %d", item.Quality)
	}
}

func TestBackstagePassQualityIncreasesByThreeWhenFiveDaysOrLess(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		SellIn:  5,
		Quality: 20,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.Quality != 23 {
		t.Errorf("expected Quality 23, got %d", item.Quality)
	}
}

func TestBackstagePassesBecomeZeroAfterConcert(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		SellIn:  0,
		Quality: 20,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.Quality != 0 {
		t.Errorf("expected Quality 0, got %d", item.Quality)
	}
}

func TestBackstagePassQualityNeverExceeds50(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		SellIn:  5,
		Quality: 49,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.Quality != 50 {
		t.Errorf("expected Quality 50, got %d", item.Quality)
	}
}

func TestSulfurasDoesNotChange(t *testing.T) {
	item := &gildedrose.Item{
		Name:    "Sulfuras, Hand of Ragnaros",
		SellIn:  10,
		Quality: 80,
	}

	gildedrose.UpdateQuality([]*gildedrose.Item{item})

	if item.SellIn != 10 {
		t.Errorf("expected SellIn 10, got %d", item.SellIn)
	}

	if item.Quality != 80 {
		t.Errorf("expected Quality 80, got %d", item.Quality)
	}
}

func TestMultipleItems(t *testing.T) {
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

	if items[0].Quality != 19 {
		t.Errorf("common item: expected Quality 19, got %d", items[0].Quality)
	}

	if items[1].Quality != 21 {
		t.Errorf("Aged Brie: expected Quality 21, got %d", items[1].Quality)
	}

	if items[2].Quality != 22 {
		t.Errorf("Backstage Pass: expected Quality 22, got %d", items[2].Quality)
	}

	if items[3].Quality != 80 {
		t.Errorf("Sulfuras: expected Quality 80, got %d", items[3].Quality)
	}
}
