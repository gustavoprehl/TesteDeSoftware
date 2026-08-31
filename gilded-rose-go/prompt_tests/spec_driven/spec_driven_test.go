// Suíte gerada pelo ChatGPT com a técnica "Spec-driven" (bônus) — em vez de
// pedir para a IA inferir as regras de negócio lendo o código, as regras
// foram dadas prontas no prompt. Ver Registro_Prompts.pdf, seção 5.
//
// Única alteração feita para integrar ao repositório: adaptado de white-box
// (package gildedrose) para black-box (package spec_driven_test),
// qualificando Item e UpdateQuality com o prefixo do pacote. Nenhuma
// asserção foi alterada.
package spec_driven_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestQualityDecreasesByOneForNormalItemBeforeSellDate(t *testing.T) {
	items := []*gildedrose.Item{{Name: "foo", SellIn: 10, Quality: 20}}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 19 {
		t.Errorf("expected quality 19, got %d", items[0].Quality)
	}
}

func TestQualityDecreasesByTwoForNormalItemAfterSellDate(t *testing.T) {
	items := []*gildedrose.Item{{Name: "foo", SellIn: -1, Quality: 20}}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 18 {
		t.Errorf("expected quality 18, got %d", items[0].Quality)
	}
}

func TestNormalItemQualityNeverBecomesNegative(t *testing.T) {
	items := []*gildedrose.Item{{Name: "foo", SellIn: 10, Quality: 0}}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 0 {
		t.Errorf("expected quality 0, got %d", items[0].Quality)
	}
}

func TestNormalItemQualityDoesNotBecomeNegativeAfterSellDate(t *testing.T) {
	items := []*gildedrose.Item{{Name: "foo", SellIn: -1, Quality: 1}}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 0 {
		t.Errorf("expected quality 0, got %d", items[0].Quality)
	}
}

func TestNormalItemSellInDecreasesByOne(t *testing.T) {
	items := []*gildedrose.Item{{Name: "foo", SellIn: 10, Quality: 20}}
	gildedrose.UpdateQuality(items)

	if items[0].SellIn != 9 {
		t.Errorf("expected SellIn 9, got %d", items[0].SellIn)
	}
}

func TestAgedBrieQualityIncreasesByOneBeforeSellDate(t *testing.T) {
	items := []*gildedrose.Item{{Name: "Aged Brie", SellIn: 10, Quality: 20}}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 21 {
		t.Errorf("expected quality 21, got %d", items[0].Quality)
	}
}

func TestAgedBrieQualityIncreasesByTwoAfterSellDate(t *testing.T) {
	items := []*gildedrose.Item{{Name: "Aged Brie", SellIn: -1, Quality: 20}}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 22 {
		t.Errorf("expected quality 22, got %d", items[0].Quality)
	}
}

func TestAgedBrieQualityNeverExceedsFifty(t *testing.T) {
	items := []*gildedrose.Item{{Name: "Aged Brie", SellIn: 10, Quality: 50}}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 50 {
		t.Errorf("expected quality 50, got %d", items[0].Quality)
	}
}

func TestAgedBrieQualityDoesNotExceedFiftyAfterSellDate(t *testing.T) {
	items := []*gildedrose.Item{{Name: "Aged Brie", SellIn: -1, Quality: 49}}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 50 {
		t.Errorf("expected quality 50, got %d", items[0].Quality)
	}
}

func TestSulfurasQualityAlwaysRemainsEighty(t *testing.T) {
	items := []*gildedrose.Item{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 10, Quality: 80},
	}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 80 {
		t.Errorf("expected quality 80, got %d", items[0].Quality)
	}
}

func TestSulfurasSellInNeverChanges(t *testing.T) {
	items := []*gildedrose.Item{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 10, Quality: 80},
	}
	gildedrose.UpdateQuality(items)

	if items[0].SellIn != 10 {
		t.Errorf("expected SellIn 10, got %d", items[0].SellIn)
	}
}

func TestSulfurasRemainsUnchangedAfterSellDate(t *testing.T) {
	items := []*gildedrose.Item{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80},
	}
	gildedrose.UpdateQuality(items)

	if items[0].SellIn != -1 {
		t.Errorf("expected SellIn -1, got %d", items[0].SellIn)
	}

	if items[0].Quality != 80 {
		t.Errorf("expected quality 80, got %d", items[0].Quality)
	}
}

func TestBackstagePassQualityIncreasesByOneWhenSellInIsGreaterThanTen(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  15,
			Quality: 20,
		},
	}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 21 {
		t.Errorf("expected quality 21, got %d", items[0].Quality)
	}
}

func TestBackstagePassQualityIncreasesByTwoWhenSellInIsTen(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  10,
			Quality: 20,
		},
	}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 22 {
		t.Errorf("expected quality 22, got %d", items[0].Quality)
	}
}

func TestBackstagePassQualityIncreasesByThreeWhenSellInIsFive(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  5,
			Quality: 20,
		},
	}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 23 {
		t.Errorf("expected quality 23, got %d", items[0].Quality)
	}
}

func TestBackstagePassQualityBecomesZeroAfterConcert(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  -1,
			Quality: 20,
		},
	}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 0 {
		t.Errorf("expected quality 0, got %d", items[0].Quality)
	}
}

func TestBackstagePassQualityNeverExceedsFifty(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  10,
			Quality: 49,
		},
	}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 50 {
		t.Errorf("expected quality 50, got %d", items[0].Quality)
	}
}

func TestBackstagePassQualityDoesNotExceedFiftyWhenThreeIncrementsApply(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  5,
			Quality: 48,
		},
	}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 50 {
		t.Errorf("expected quality 50, got %d", items[0].Quality)
	}
}

func TestBackstagePassSellInDecreasesByOne(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  10,
			Quality: 20,
		},
	}
	gildedrose.UpdateQuality(items)

	if items[0].SellIn != 9 {
		t.Errorf("expected SellIn 9, got %d", items[0].SellIn)
	}
}

func TestQualityAtSellInZeroUsesNormalRulesBeforeExpiration(t *testing.T) {
	items := []*gildedrose.Item{{Name: "foo", SellIn: 0, Quality: 20}}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 18 {
		t.Errorf("expected quality 18, got %d", items[0].Quality)
	}
}

func TestAgedBrieAtSellInZeroIncreasesByTwo(t *testing.T) {
	items := []*gildedrose.Item{{Name: "Aged Brie", SellIn: 0, Quality: 20}}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 22 {
		t.Errorf("expected quality 22, got %d", items[0].Quality)
	}
}

func TestBackstagePassAtSellInZeroBecomesZero(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  0,
			Quality: 20,
		},
	}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 0 {
		t.Errorf("expected quality 0, got %d", items[0].Quality)
	}
}

func TestQualityNeverBecomesNegativeForAnySupportedItem(t *testing.T) {
	items := []*gildedrose.Item{
		{Name: "foo", SellIn: -1, Quality: 0},
		{Name: "Aged Brie", SellIn: 10, Quality: 0},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 10, Quality: 80},
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  -1,
			Quality: 0,
		},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality < 0 {
		t.Errorf("normal item quality became negative: %d", items[0].Quality)
	}

	if items[1].Quality < 0 {
		t.Errorf("Aged Brie quality became negative: %d", items[1].Quality)
	}

	if items[2].Quality != 80 {
		t.Errorf("Sulfuras quality changed: %d", items[2].Quality)
	}

	if items[3].Quality < 0 {
		t.Errorf("Backstage pass quality became negative: %d", items[3].Quality)
	}
}
