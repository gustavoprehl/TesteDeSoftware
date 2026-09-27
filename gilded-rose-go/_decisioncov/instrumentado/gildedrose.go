// Cópia instrumentada de gildedrose/gildedrose.go usada SOMENTE pelo medidor
// de Cobertura de Decisão (_decisioncov/main.go), via `go test -overlay`.
//
// A lógica é idêntica à de produção. A única diferença é que cada lado (True e
// False) das 17 decisões D1–D17 de docs/iteracao2/analise-estrutural.md
// ganhou uma chamada marcador(...) em um bloco próprio. Com isso, a cobertura
// de statements desta cópia equivale à Cobertura de Decisão do original: um
// ramo foi percorrido se, e somente se, o bloco do seu marcador foi executado.
//
// Para manter a granularidade do CFG (docs/iteracao2/cfg.md), a condição
// composta D2 (`&&`) continua sendo uma única decisão.
//
// O laço `for` foi reescrito com um `if` explícito para que D1 também tenha um
// bloco para cada resultado; o comportamento é o mesmo.
package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

// marcador não faz nada; só existe para que cada ramo tenha um statement próprio.
func marcador(string) {}

func UpdateQuality(items []*Item) {
	for i := 0; ; i++ {
		if i < len(items) {
			marcador("D1T")
		} else {
			marcador("D1F")
			break
		}

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			marcador("D2T")
			if items[i].Quality > 0 {
				marcador("D3T")
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					marcador("D4T")
					items[i].Quality = items[i].Quality - 1
				} else {
					marcador("D4F")
				}
			} else {
				marcador("D3F")
			}
		} else {
			marcador("D2F")
			if items[i].Quality < 50 {
				marcador("D5T")
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					marcador("D6T")
					if items[i].SellIn < 11 {
						marcador("D7T")
						if items[i].Quality < 50 {
							marcador("D8T")
							items[i].Quality = items[i].Quality + 1
						} else {
							marcador("D8F")
						}
					} else {
						marcador("D7F")
					}
					if items[i].SellIn < 6 {
						marcador("D9T")
						if items[i].Quality < 50 {
							marcador("D10T")
							items[i].Quality = items[i].Quality + 1
						} else {
							marcador("D10F")
						}
					} else {
						marcador("D9F")
					}
				} else {
					marcador("D6F")
				}
			} else {
				marcador("D5F")
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			marcador("D11T")
			items[i].SellIn = items[i].SellIn - 1
		} else {
			marcador("D11F")
		}

		if items[i].SellIn < 0 {
			marcador("D12T")
			if items[i].Name != "Aged Brie" {
				marcador("D13T")
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					marcador("D14T")
					if items[i].Quality > 0 {
						marcador("D15T")
						if items[i].Name != "Sulfuras, Hand of Ragnaros" {
							marcador("D16T")
							items[i].Quality = items[i].Quality - 1
						} else {
							marcador("D16F")
						}
					} else {
						marcador("D15F")
					}
				} else {
					marcador("D14F")
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				marcador("D13F")
				if items[i].Quality < 50 {
					marcador("D17T")
					items[i].Quality = items[i].Quality + 1
				} else {
					marcador("D17F")
				}
			}
		} else {
			marcador("D12F")
		}
	}
}
