// EVIDÊNCIA (NÃO integrado ao build principal): variação de
// _evidence/few_shot_original com APENAS o typo `*testing.t` corrigido para
// `*testing.T`, mantendo o segundo erro da IA (a struct `GildedRose{Items:
// items}` que não existe neste pacote — a função real é
// `UpdateQuality(items []*Item)`). Isola o segundo problema para o vídeo:
// mesmo corrigindo o erro de sintaxe, o código ainda não compila.
//
// Este diretório começa com "_" de propósito — ver comentário equivalente em
// _evidence/few_shot_original/few_shot_test.go.
package gildedrose

import "testing"

func TestQualityDecreasesByOneForNormalItemBeforeSellDate(t *testing.T) {
	// Arrange
	items := []*Item{{Name: "foo", SellIn: 10, Quality: 20}}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 19 {
		t.Errorf("esperado 19, obtido %d", items[0].Quality)
	}
}

func TestQualityDecreasesByTwoForNormalItemAfterSellDate(t *testing.T) {
	// Arrange
	items := []*Item{{Name: "foo", SellIn: 0, Quality: 20}}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 18 {
		t.Errorf("esperado 18, obtido %d", items[0].Quality)
	}
}

func TestQualityDoesNotDecreaseBelowZeroForNormalItem(t *testing.T) {
	// Arrange
	items := []*Item{{Name: "foo", SellIn: 10, Quality: 0}}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 0 {
		t.Errorf("esperado 0, obtido %d", items[0].Quality)
	}
}

func TestQualityDecreasesByTwoForNormalItemAtZeroSellIn(t *testing.T) {
	// Arrange
	items := []*Item{{Name: "foo", SellIn: 0, Quality: 1}}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 0 {
		t.Errorf("esperado 0, obtido %d", items[0].Quality)
	}
}

func TestSellInDecreasesByOneForNormalItem(t *testing.T) {
	// Arrange
	items := []*Item{{Name: "foo", SellIn: 10, Quality: 20}}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].SellIn != 9 {
		t.Errorf("esperado 9, obtido %d", items[0].SellIn)
	}
}

func TestQualityIncreasesByOneForAgedBrieBeforeSellDate(t *testing.T) {
	// Arrange
	items := []*Item{{Name: "Aged Brie", SellIn: 10, Quality: 20}}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 21 {
		t.Errorf("esperado 21, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByTwoForAgedBrieAfterSellDate(t *testing.T) {
	// Arrange
	items := []*Item{{Name: "Aged Brie", SellIn: 0, Quality: 20}}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 22 {
		t.Errorf("esperado 22, obtido %d", items[0].Quality)
	}
}

func TestQualityDoesNotExceedFiftyForAgedBrie(t *testing.T) {
	// Arrange
	items := []*Item{{Name: "Aged Brie", SellIn: 10, Quality: 50}}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 50 {
		t.Errorf("esperado 50, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByOneForAgedBrieAtQualityFortyNineBeforeSellDate(t *testing.T) {
	// Arrange
	items := []*Item{{Name: "Aged Brie", SellIn: 10, Quality: 49}}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 50 {
		t.Errorf("esperado 50, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByTwoForAgedBrieAtQualityFortyNineAfterSellDate(t *testing.T) {
	// Arrange
	items := []*Item{{Name: "Aged Brie", SellIn: 0, Quality: 49}}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 50 {
		t.Errorf("esperado 50, obtido %d", items[0].Quality)
	}
}

func TestSellInDecreasesByOneForAgedBrie(t *testing.T) {
	// Arrange
	items := []*Item{{Name: "Aged Brie", SellIn: 10, Quality: 20}}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].SellIn != 9 {
		t.Errorf("esperado 9, obtido %d", items[0].SellIn)
	}
}

func TestQualityDoesNotChangeForSulfurasBeforeSellDate(t *testing.T) {
	// Arrange
	items := []*Item{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 10, Quality: 80},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 80 {
		t.Errorf("esperado 80, obtido %d", items[0].Quality)
	}
}

func TestQualityDoesNotChangeForSulfurasAfterSellDate(t *testing.T) {
	// Arrange
	items := []*Item{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 80 {
		t.Errorf("esperado 80, obtido %d", items[0].Quality)
	}
}

func TestSellInDoesNotChangeForSulfuras(t *testing.T) {
	// Arrange
	items := []*Item{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 10, Quality: 80},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].SellIn != 10 {
		t.Errorf("esperado 10, obtido %d", items[0].SellIn)
	}
}

func TestQualityIncreasesByOneForBackstagePassMoreThanTenDaysBeforeConcert(t *testing.T) {
	// Arrange
	items := []*Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  15,
			Quality: 20,
		},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 21 {
		t.Errorf("esperado 21, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByTwoForBackstagePassTenDaysBeforeConcert(t *testing.T) {
	// Arrange
	items := []*Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  10,
			Quality: 20,
		},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 22 {
		t.Errorf("esperado 22, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByTwoForBackstagePassSixDaysBeforeConcert(t *testing.T) {
	// Arrange
	items := []*Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  6,
			Quality: 20,
		},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 22 {
		t.Errorf("esperado 22, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByThreeForBackstagePassFiveDaysBeforeConcert(t *testing.T) {
	// Arrange
	items := []*Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  5,
			Quality: 20,
		},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 23 {
		t.Errorf("esperado 23, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByThreeForBackstagePassOneDayBeforeConcert(t *testing.T) {
	// Arrange
	items := []*Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  1,
			Quality: 20,
		},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 23 {
		t.Errorf("esperado 23, obtido %d", items[0].Quality)
	}
}

func TestQualityDropsToZeroForBackstagePassOnConcertDate(t *testing.T) {
	// Arrange
	items := []*Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  0,
			Quality: 20,
		},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 0 {
		t.Errorf("esperado 0, obtido %d", items[0].Quality)
	}
}

func TestQualityRemainsZeroForBackstagePassAfterConcertDate(t *testing.T) {
	// Arrange
	items := []*Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  -1,
			Quality: 20,
		},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 0 {
		t.Errorf("esperado 0, obtido %d", items[0].Quality)
	}
}

func TestQualityDoesNotExceedFiftyForBackstagePass(t *testing.T) {
	// Arrange
	items := []*Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  10,
			Quality: 50,
		},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 50 {
		t.Errorf("esperado 50, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesOnlyUpToFiftyForBackstagePass(t *testing.T) {
	// Arrange
	items := []*Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  5,
			Quality: 49,
		},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].Quality != 50 {
		t.Errorf("esperado 50, obtido %d", items[0].Quality)
	}
}

func TestSellInDecreasesByOneForBackstagePass(t *testing.T) {
	// Arrange
	items := []*Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  10,
			Quality: 20,
		},
	}
	gr := GildedRose{Items: items}

	// Act
	gr.UpdateQuality()

	// Assert
	if items[0].SellIn != 9 {
		t.Errorf("esperado 9, obtido %d", items[0].SellIn)
	}
}
