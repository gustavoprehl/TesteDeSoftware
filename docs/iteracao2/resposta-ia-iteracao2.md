# Resposta da IA ao Prompt Estruturado (Iteração 2)

> Registro literal da resposta do **ChatGPT** ao prompt de [prompt-iteracao2.txt](prompt-iteracao2.txt),
> em uma execução única, em conversa nova (detalhes em [prompt-estruturado.md](prompt-estruturado.md),
> seção 4). Conversa original: <https://chatgpt.com/share/6ac4e403-4234-83e9-87e6-89506b49550b>.
>
> O texto foi copiado da conversa. Na cópia, as tabelas chegaram separadas por tabulação e os
> blocos de código sem delimitadores; abaixo, só essa marcação Markdown foi reconstituída
> (tabelas, títulos e blocos de código). Nenhuma palavra, valor ou ramo foi alterado.
> O código da Etapa 5 está integrado em
> [prompt_tests/estruturado/estruturado_test.go](../../gilded-rose-go/prompt_tests/estruturado/estruturado_test.go).

---

A análise abaixo trata exclusivamente da estrutura de controle apresentada. Não uso nenhuma regra de negócio externa ao código.

Para simplificar as restrições, vou usar:

- A = "Aged Brie"
- B = "Backstage passes to a TAFKAL80ETC concert"
- S = "Sulfuras, Hand of Ragnaros"
- O = qualquer outro nome
- s = SellIn de entrada
- q = Quality de entrada

Como SellIn e Quality são int, expressões como s < 1 equivalem a s <= 0.

## Etapa 1 — Condições de alcance dos 34 ramos

As condições abaixo são condições sobre os valores de entrada, mesmo quando a decisão lê um valor já alterado.

| Ramo | Restrição de entrada que permite tomar o ramo |
|---|---|
| D1T | Sempre, pois cada teste contém exatamente 1 item e inicialmente 0 < len(items) |
| D1F | Sempre após processar o único item, pois i passa a 1 e 1 < 1 é falso |
| D2T | Name = O ou Name = S |
| D2F | Name = A ou Name = B |
| D3T | (Name = O ou S) ∧ q > 0 |
| D3F | (Name = O ou S) ∧ q <= 0 |
| D4T | Name = O ∧ q > 0 |
| D4F | Name = S ∧ q > 0 |
| D5T | (Name = A ou B) ∧ q < 50 |
| D5F | (Name = A ou B) ∧ q >= 50 |
| D6T | Name = B ∧ q < 50 |
| D6F | Name = A ∧ q < 50 |
| D7T | Name = B ∧ q < 50 ∧ s < 11 |
| D7F | Name = B ∧ q < 50 ∧ s >= 11 |
| D8T | Name = B ∧ s < 11 ∧ q <= 48 |
| D8F | Name = B ∧ s < 11 ∧ q = 49 |
| D9T | Name = B ∧ q < 50 ∧ s < 6 |
| D9F | Name = B ∧ q < 50 ∧ s >= 6 |
| D10T | Name = B ∧ s < 6 ∧ q <= 47 |
| D10F | Name = B ∧ s < 6 ∧ (q = 48 ou q = 49) |
| D11T | Name != S |
| D11F | Name = S |
| D12T | (Name != S ∧ s <= 0) ou (Name = S ∧ s < 0) |
| D12F | (Name != S ∧ s >= 1) ou (Name = S ∧ s >= 0) |
| D13T | (Name = B ∧ s <= 0) ou (Name = O ∧ s <= 0) ou (Name = S ∧ s < 0) |
| D13F | Name = A ∧ s <= 0 |
| D14T | (Name = O ∧ s <= 0) ou (Name = S ∧ s < 0) |
| D14F | Name = B ∧ s <= 0 |
| D15T | (Name = O ∧ s <= 0 ∧ q > 1) ou (Name = S ∧ s < 0 ∧ q > 0) |
| D15F | (Name = O ∧ s <= 0 ∧ q <= 1) ou (Name = S ∧ s < 0 ∧ q <= 0) |
| D16T | Name = O ∧ s <= 0 ∧ q > 1 |
| D16F | Name = S ∧ s < 0 ∧ q > 0 |
| D17T | Name = A ∧ s <= 0 ∧ q <= 48 |
| D17F | Name = A ∧ s <= 0 ∧ q >= 49 |

### Pontos em que a mutação anterior altera a restrição

Há quatro casos particularmente importantes.

Em D8, Quality já passou pela linha 19. Logo:

```text
Quality_D8 = q + 1
```

Portanto:

```text
D8T: q + 1 < 50  => q <= 48
D8F: q + 1 >= 50 => como D5T exige q < 50, q = 49
```

Em D10, para s < 6, D7 necessariamente é verdadeiro. Assim:

```text
q <= 48:
linha 19 => q + 1
D8T      => linha 23 => q + 2
```

Consequentemente:

```text
D10T: q + 2 < 50 => q <= 47
D10F: q = 48 ou q = 49
```

Em D12, SellIn já foi decrementado para qualquer item exceto Sulfuras:

```text
item não-Sulfuras: SellIn_D12 = s - 1
Sulfuras:          SellIn_D12 = s
```

Finalmente, para um item comum expirado, D15 lê Quality depois da linha 14:

```text
q > 0  => Quality_D15 = q - 1
q <= 0 => Quality_D15 = q
```

Por isso um item comum somente consegue D15T se q > 1.

## Etapa 2 — Seleção e rastreamento dos casos

O conjunto mínimo encontrado contém 8 casos.

| Caso | Name | SellIn | Quality | Saída SellIn | Saída Quality |
|---|---|---|---|---|---|
| T1 | Backstage | 0 | 49 | -1 | 0 |
| T2 | Backstage | 5 | 10 | 4 | 13 |
| T3 | Backstage | 11 | 10 | 10 | 11 |
| T4 | Aged Brie | 0 | 0 | -1 | 2 |
| T5 | Aged Brie | 0 | 50 | -1 | 50 |
| T6 | Sulfuras | -1 | 10 | -1 | 10 |
| T7 | item comum | 0 | 2 | -1 | 0 |
| T8 | item comum | 0 | 0 | -1 | 0 |

### T1 — Backstage, SellIn=0, Quality=49

Entrada:

```text
Name    = Backstage passes to a TAFKAL80ETC concert
SellIn  = 0
Quality = 49
```

| Linha/decisão | Resultado / estado |
|---|---|
| D1 | T |
| D2 | F |
| D5 | 49 < 50 → T |
| linha 19 | Quality: 49 -> 50 |
| D6 | T |
| D7 | 0 < 11 → T |
| D8 | 50 < 50 → F |
| D9 | 0 < 6 → T |
| D10 | 50 < 50 → F |
| D11 | T |
| linha 36 | SellIn: 0 -> -1 |
| D12 | -1 < 0 → T |
| D13 | T |
| D14 | F |
| linha 48 | Quality: 50 -> 0 |
| D1 | F |

Sequência:

```text
D1T, D2F, D5T, D6T, D7T, D8F, D9T, D10F,
D11T, D12T, D13T, D14F, D1F
```

Saída:

```text
SellIn  = -1
Quality = 0
```

### T2 — Backstage, SellIn=5, Quality=10

| Linha/decisão | Resultado / estado |
|---|---|
| D1 | T |
| D2 | F |
| D5 | T |
| linha 19 | Quality: 10 -> 11 |
| D6 | T |
| D7 | 5 < 11 → T |
| D8 | 11 < 50 → T |
| linha 23 | Quality: 11 -> 12 |
| D9 | 5 < 6 → T |
| D10 | 12 < 50 → T |
| linha 28 | Quality: 12 -> 13 |
| D11 | T |
| linha 36 | SellIn: 5 -> 4 |
| D12 | 4 < 0 → F |
| D1 | F |

Sequência:

```text
D1T, D2F, D5T, D6T, D7T, D8T, D9T, D10T,
D11T, D12F, D1F
```

Saída:

```text
SellIn  = 4
Quality = 13
```

### T3 — Backstage, SellIn=11, Quality=10

| Linha/decisão | Resultado / estado |
|---|---|
| D1 | T |
| D2 | F |
| D5 | T |
| linha 19 | Quality: 10 -> 11 |
| D6 | T |
| D7 | 11 < 11 → F |
| D9 | 11 < 6 → F |
| D11 | T |
| linha 36 | SellIn: 11 -> 10 |
| D12 | 10 < 0 → F |
| D1 | F |

Sequência:

```text
D1T, D2F, D5T, D6T, D7F, D9F, D11T, D12F, D1F
```

Saída:

```text
SellIn  = 10
Quality = 11
```

### T4 — Aged Brie, SellIn=0, Quality=0

| Linha/decisão | Resultado / estado |
|---|---|
| D1 | T |
| D2 | F |
| D5 | 0 < 50 → T |
| linha 19 | Quality: 0 -> 1 |
| D6 | F |
| D11 | T |
| linha 36 | SellIn: 0 -> -1 |
| D12 | T |
| D13 | F |
| D17 | 1 < 50 → T |
| linha 52 | Quality: 1 -> 2 |
| D1 | F |

Sequência:

```text
D1T, D2F, D5T, D6F, D11T, D12T, D13F, D17T, D1F
```

Saída:

```text
SellIn  = -1
Quality = 2
```

### T5 — Aged Brie, SellIn=0, Quality=50

| Linha/decisão | Resultado / estado |
|---|---|
| D1 | T |
| D2 | F |
| D5 | 50 < 50 → F |
| D11 | T |
| linha 36 | SellIn: 0 -> -1 |
| D12 | T |
| D13 | F |
| D17 | 50 < 50 → F |
| D1 | F |

Sequência:

```text
D1T, D2F, D5F, D11T, D12T, D13F, D17F, D1F
```

Saída:

```text
SellIn  = -1
Quality = 50
```

### T6 — Sulfuras, SellIn=-1, Quality=10

| Linha/decisão | Resultado / estado |
|---|---|
| D1 | T |
| D2 | T |
| D3 | 10 > 0 → T |
| D4 | Name != Sulfuras → F |
| D11 | Name != Sulfuras → F |
| SellIn | não é alterado; continua -1 |
| D12 | -1 < 0 → T |
| D13 | T |
| D14 | T |
| D15 | 10 > 0 → T |
| D16 | Name != Sulfuras → F |
| D1 | F |

Sequência:

```text
D1T, D2T, D3T, D4F, D11F, D12T,
D13T, D14T, D15T, D16F, D1F
```

Saída:

```text
SellIn  = -1
Quality = 10
```

Esse caso demonstra explicitamente a observação estrutural do enunciado: Sulfuras pode chegar a D12T quando o SellIn já entra negativo, mesmo não sofrendo o decremento da linha 36.

### T7 — Item comum, SellIn=0, Quality=2

Usaremos "Normal Item" como nome.

| Linha/decisão | Resultado / estado |
|---|---|
| D1 | T |
| D2 | T |
| D3 | 2 > 0 → T |
| D4 | T |
| linha 14 | Quality: 2 -> 1 |
| D11 | T |
| linha 36 | SellIn: 0 -> -1 |
| D12 | T |
| D13 | T |
| D14 | T |
| D15 | 1 > 0 → T |
| D16 | T |
| linha 44 | Quality: 1 -> 0 |
| D1 | F |

Sequência:

```text
D1T, D2T, D3T, D4T, D11T, D12T,
D13T, D14T, D15T, D16T, D1F
```

Saída:

```text
SellIn  = -1
Quality = 0
```

### T8 — Item comum, SellIn=0, Quality=0

| Linha/decisão | Resultado / estado |
|---|---|
| D1 | T |
| D2 | T |
| D3 | 0 > 0 → F |
| D11 | T |
| linha 36 | SellIn: 0 -> -1 |
| D12 | T |
| D13 | T |
| D14 | T |
| D15 | 0 > 0 → F |
| D1 | F |

Sequência:

```text
D1T, D2T, D3F, D11T, D12T,
D13T, D14T, D15F, D1F
```

Saída:

```text
SellIn  = -1
Quality = 0
```

## Etapa 3 — Matriz de rastreabilidade

Nenhum dos 34 ramos fica sem cobertura.

| Ramo | Caso(s) que cobrem |
|---|---|
| D1T | T1, T2, T3, T4, T5, T6, T7, T8 |
| D1F | T1, T2, T3, T4, T5, T6, T7, T8 |
| D2T | T6, T7, T8 |
| D2F | T1, T2, T3, T4, T5 |
| D3T | T6, T7 |
| D3F | T8 |
| D4T | T7 |
| D4F | T6 |
| D5T | T1, T2, T3, T4 |
| D5F | T5 |
| D6T | T1, T2, T3 |
| D6F | T4 |
| D7T | T1, T2 |
| D7F | T3 |
| D8T | T2 |
| D8F | T1 |
| D9T | T1, T2 |
| D9F | T3 |
| D10T | T2 |
| D10F | T1 |
| D11T | T1, T2, T3, T4, T5, T7, T8 |
| D11F | T6 |
| D12T | T1, T4, T5, T6, T7, T8 |
| D12F | T2, T3 |
| D13T | T1, T6, T7, T8 |
| D13F | T4, T5 |
| D14T | T6, T7, T8 |
| D14F | T1 |
| D15T | T6, T7 |
| D15F | T8 |
| D16T | T7 |
| D16F | T6 |
| D17T | T4 |
| D17F | T5 |

Portanto:

```text
34 / 34 ramos cobertos
Cobertura de decisão = 100%
```

## Etapa 4 — Justificativa de minimalidade

O conjunto de 8 casos não é apenas suficiente; com as restrições fornecidas, ele é mínimo.

### 1. Backstage exige pelo menos 3 casos

Para alcançar D7, D8, D9 e D10, obrigatoriamente precisamos de Backstage.

O ramo:

```text
D7F
```

exige:

```text
SellIn >= 11
```

Já D10 só é alcançado quando:

```text
D9T => SellIn < 6
```

Logo, o caso usado para D7F não pode cobrir D10.

Além disso, precisamos de ambos:

```text
D10T
D10F
```

Uma única execução que chega a D10 só pode tomar um desses dois ramos.

Portanto:

```text
1 caso para D7F
+
1 caso para D10T
+
1 caso para D10F
=
mínimo de 3 Backstage
```

São T1, T2 e T3.

### 2. Aged Brie exige pelo menos 2 casos

D17 somente é alcançada por "Aged Brie".

Precisamos de:

```text
D17T
D17F
```

Como uma execução só pode escolher um dos lados da decisão:

```text
mínimo de 2 Aged Brie
```

São T4 e T5.

Até aqui:

```text
3 + 2 = 5 casos
```

### 3. D16F exige um Sulfuras específico

Para obter:

```text
D16F
```

é necessário que D16 seja avaliada e:

```text
Name != "Sulfuras, Hand of Ragnaros"
```

seja falso.

Logo obrigatoriamente:

```text
Name = Sulfuras
```

Além disso, para chegar a D16:

```text
D12T
D13T
D14T
D15T
```

Portanto precisamos de Sulfuras com:

```text
SellIn < 0
Quality > 0
```

Esse caso não pode ser nenhum dos cinco anteriores.

T6 é obrigatório estruturalmente.

Total mínimo:

```text
6 casos
```

### 4. D16T exige um item comum

Para D16T, o nome deve satisfazer:

```text
Name != Sulfuras
```

Mas para chegar a D16 também já precisamos de:

```text
D13T => não Aged Brie
D14T => não Backstage
```

Logo D16T exige necessariamente um item que não seja:

```text
Aged Brie
Backstage
Sulfuras
```

ou seja, um item comum.

Para ele alcançar D15/D16 depois da primeira redução:

```text
SellIn <= 0
Quality > 1
```

Assim precisamos de pelo menos um caso adicional, T7.

Total mínimo:

```text
7 casos
```

### 5. D3F não pode ser coberto por T7 nem pelo Sulfuras de D16F

Ainda falta:

```text
D3F
```

D3F exige:

```text
Quality <= 0
```

O caso comum usado para D16T, porém, exige:

```text
Quality > 1
```

Logo o mesmo item não pode satisfazer os dois.

Também não pode ser o Sulfuras usado para D16F, pois esse obrigatoriamente exige:

```text
Quality > 0
```

E não pode ser nenhum Aged Brie ou Backstage, porque ambos seguem por D2F, não chegando a D3.

Consequentemente existe a necessidade de um oitavo caso, T8.

Temos então o limite inferior:

```text
3 Backstage
+ 2 Aged Brie
+ 1 Sulfuras
+ 1 item comum para D16T
+ 1 item comum/Sulfuras com Q<=0 para D3F
= 8 casos
```

Como a suíte apresentada usa exatamente 8 casos e cobre todos os 34 ramos:

$$
\boxed{\text{mínimo = 8 casos de teste}}
$$

## Etapa 5 — Código Go

Arquivo completo estruturado_test.go:

```go
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
```

## Observações

Todos os 17 pontos de decisão possuem os dois resultados exercitados, totalizando os 34 ramos do modelo fornecido. Não há ramo inalcançável nesse CFG.

A suíte é de caracterização: os valores esperados representam exatamente as mutações produzidas pelo código legado apresentado, inclusive comportamentos peculiares como Backstage com Quality=49 chegar a 50 antes de D8/D10 e depois ser zerado quando expirado.
