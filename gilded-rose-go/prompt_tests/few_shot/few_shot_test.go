// Suíte gerada pelo ChatGPT com a técnica "Few-shot" — o modelo recebeu um
// exemplo de teste já formatado e foi instruído a seguir o mesmo padrão.
// Ver Registro_Prompts.pdf, seção 4.
//
// Esta foi a única das cinco conversas que NÃO compilou. Duas correções
// manuais foram necessárias para integrar ao repositório (documentadas
// também no PLANO-DUPLA2.md, responsabilidade da Dupla 2):
//
//  1. Typo de sintaxe: `*testing.t` -> `*testing.T` em
//     TestQualityIncreasesByTwoForAgedBrieAtQualityFortyNineAfterSellDate.
//  2. Erro estrutural mais sério: a IA copiou a struct do exemplo do prompt
//     (`gr := GildedRose{Items: items}; gr.UpdateQuality()`) em vez de usar a
//     função que existe de fato no código real fornecido
//     (`UpdateQuality(items []*Item)`). O tipo GildedRose nem existe neste
//     pacote. Todas as chamadas foram trocadas para
//     `gildedrose.UpdateQuality(items)`.
//
// Nenhuma asserção/valor esperado foi alterado — só o necessário para
// compilar e rodar.
package few_shot_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestQualityDecreasesByOneForNormalItemBeforeSellDate(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{{Name: "foo", SellIn: 10, Quality: 20}}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 19 {
		t.Errorf("esperado 19, obtido %d", items[0].Quality)
	}
}

func TestQualityDecreasesByTwoForNormalItemAfterSellDate(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{{Name: "foo", SellIn: 0, Quality: 20}}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 18 {
		t.Errorf("esperado 18, obtido %d", items[0].Quality)
	}
}

func TestQualityDoesNotDecreaseBelowZeroForNormalItem(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{{Name: "foo", SellIn: 10, Quality: 0}}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 0 {
		t.Errorf("esperado 0, obtido %d", items[0].Quality)
	}
}

func TestQualityDecreasesByTwoForNormalItemAtZeroSellIn(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{{Name: "foo", SellIn: 0, Quality: 1}}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 0 {
		t.Errorf("esperado 0, obtido %d", items[0].Quality)
	}
}

func TestSellInDecreasesByOneForNormalItem(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{{Name: "foo", SellIn: 10, Quality: 20}}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].SellIn != 9 {
		t.Errorf("esperado 9, obtido %d", items[0].SellIn)
	}
}

func TestQualityIncreasesByOneForAgedBrieBeforeSellDate(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{{Name: "Aged Brie", SellIn: 10, Quality: 20}}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 21 {
		t.Errorf("esperado 21, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByTwoForAgedBrieAfterSellDate(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{{Name: "Aged Brie", SellIn: 0, Quality: 20}}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 22 {
		t.Errorf("esperado 22, obtido %d", items[0].Quality)
	}
}

func TestQualityDoesNotExceedFiftyForAgedBrie(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{{Name: "Aged Brie", SellIn: 10, Quality: 50}}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 50 {
		t.Errorf("esperado 50, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByOneForAgedBrieAtQualityFortyNineBeforeSellDate(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{{Name: "Aged Brie", SellIn: 10, Quality: 49}}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 50 {
		t.Errorf("esperado 50, obtido %d", items[0].Quality)
	}
}

// Correção manual: a IA escreveu "*testing.t" (minúsculo) aqui, o que não
// compila em Go (o tipo correto é *testing.T). Erro de digitação -- a IA
// travou a compilação de todo o pacote por causa dele.
func TestQualityIncreasesByTwoForAgedBrieAtQualityFortyNineAfterSellDate(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{{Name: "Aged Brie", SellIn: 0, Quality: 49}}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 50 {
		t.Errorf("esperado 50, obtido %d", items[0].Quality)
	}
}

func TestSellInDecreasesByOneForAgedBrie(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{{Name: "Aged Brie", SellIn: 10, Quality: 20}}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].SellIn != 9 {
		t.Errorf("esperado 9, obtido %d", items[0].SellIn)
	}
}

func TestQualityDoesNotChangeForSulfurasBeforeSellDate(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 10, Quality: 80},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 80 {
		t.Errorf("esperado 80, obtido %d", items[0].Quality)
	}
}

func TestQualityDoesNotChangeForSulfurasAfterSellDate(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 80 {
		t.Errorf("esperado 80, obtido %d", items[0].Quality)
	}
}

func TestSellInDoesNotChangeForSulfuras(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 10, Quality: 80},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].SellIn != 10 {
		t.Errorf("esperado 10, obtido %d", items[0].SellIn)
	}
}

func TestQualityIncreasesByOneForBackstagePassMoreThanTenDaysBeforeConcert(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  15,
			Quality: 20,
		},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 21 {
		t.Errorf("esperado 21, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByTwoForBackstagePassTenDaysBeforeConcert(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  10,
			Quality: 20,
		},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 22 {
		t.Errorf("esperado 22, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByTwoForBackstagePassSixDaysBeforeConcert(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  6,
			Quality: 20,
		},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 22 {
		t.Errorf("esperado 22, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByThreeForBackstagePassFiveDaysBeforeConcert(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  5,
			Quality: 20,
		},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 23 {
		t.Errorf("esperado 23, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesByThreeForBackstagePassOneDayBeforeConcert(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  1,
			Quality: 20,
		},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 23 {
		t.Errorf("esperado 23, obtido %d", items[0].Quality)
	}
}

func TestQualityDropsToZeroForBackstagePassOnConcertDate(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  0,
			Quality: 20,
		},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 0 {
		t.Errorf("esperado 0, obtido %d", items[0].Quality)
	}
}

func TestQualityRemainsZeroForBackstagePassAfterConcertDate(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  -1,
			Quality: 20,
		},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 0 {
		t.Errorf("esperado 0, obtido %d", items[0].Quality)
	}
}

func TestQualityDoesNotExceedFiftyForBackstagePass(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  10,
			Quality: 50,
		},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 50 {
		t.Errorf("esperado 50, obtido %d", items[0].Quality)
	}
}

func TestQualityIncreasesOnlyUpToFiftyForBackstagePass(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  5,
			Quality: 49,
		},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].Quality != 50 {
		t.Errorf("esperado 50, obtido %d", items[0].Quality)
	}
}

func TestSellInDecreasesByOneForBackstagePass(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  10,
			Quality: 20,
		},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	if items[0].SellIn != 9 {
		t.Errorf("esperado 9, obtido %d", items[0].SellIn)
	}
}
