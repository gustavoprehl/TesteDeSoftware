# Resposta da IA ao Prompt Estruturado (Iteração 2)

> Registro literal da resposta obtida com o prompt de [prompt-iteracao2.txt](prompt-iteracao2.txt),
> em uma execução única, em contexto novo e isolado (detalhes em
> [prompt-estruturado.md](prompt-estruturado.md), seção 4). Nada foi editado abaixo desta nota.
> O código da Etapa 5 está integrado em
> [prompt_tests/estruturado/estruturado_test.go](../../gilded-rose-go/prompt_tests/estruturado/estruturado_test.go).

---

# Suíte estrutural (caixa branca) para `UpdateQuality`: 100% de Cobertura de Decisão

**Notação:** N = Name, S = SellIn de entrada, Q = Quality de entrada.
Categorias de nome:
- **AB** = "Aged Brie"
- **BP** = "Backstage passes to a TAFKAL80ETC concert"
- **SU** = "Sulfuras, Hand of Ragnaros"
- **C** = qualquer outro nome (item comum)

Todos os casos usam uma lista com exatamente 1 item.

---

## Etapa 1: Condições de alcance (restrições sobre a entrada)

Valores derivados que as guardas realmente leem:
- **Linha 22:** Q₂₂ = Q + 1.
- **Linha 27:** Q₂₇ = Q + 2 se S < 11 e Q ≤ 48. Caso contrário, Q₂₇ = Q + 1.
- **Linha 39:** S₃₉ = S − 1 se N ≠ SU. Se N = SU, S₃₉ = S.
- **Linha 42, item comum:** Q₄₂ = Q − 1 se Q > 0. Caso contrário, Q₄₂ = Q.
- **Linha 42, Sulfuras:** Q₄₂ = Q.
- **Linha 51:** Q₅₁ = Q + 1 se Q < 50. Caso contrário, Q₅₁ = Q.

| Ramo | Restrição de entrada |
|---|---|
| D1T | lista com ≥ 1 item (i = 0 < 1) |
| D1F | sempre ocorre após processar o único item (i = 1 = len) |
| D2T | N ∈ {C, SU} |
| D2F | N ∈ {AB, BP} |
| D3T | N ∈ {C, SU} ∧ Q > 0 |
| D3F | N ∈ {C, SU} ∧ Q ≤ 0 |
| D4T | N = C ∧ Q > 0 |
| D4F | N = SU ∧ Q > 0 |
| D5T | N ∈ {AB, BP} ∧ Q < 50 |
| D5F | N ∈ {AB, BP} ∧ Q ≥ 50 |
| D6T | N = BP ∧ Q < 50 |
| D6F | N = AB ∧ Q < 50 |
| D7T | N = BP ∧ Q < 50 ∧ S < 11 |
| D7F | N = BP ∧ Q < 50 ∧ S ≥ 11 |
| D8T | N = BP ∧ S < 11 ∧ Q+1 < 50, ou seja, Q ≤ 48 |
| D8F | N = BP ∧ S < 11 ∧ Q < 50 ∧ Q+1 ≥ 50, ou seja, Q = 49 |
| D9T | N = BP ∧ Q < 50 ∧ S < 6 |
| D9F | N = BP ∧ Q < 50 ∧ S ≥ 6 |
| D10T | N = BP ∧ S < 6 ∧ Q₂₇ = Q+2 < 50, ou seja, Q ≤ 47 |
| D10F | N = BP ∧ S < 6 ∧ 48 ≤ Q ≤ 49 (Q₂₇ = 50 nos dois casos) |
| D11T | N ≠ SU |
| D11F | N = SU |
| D12T | (N ≠ SU ∧ S − 1 < 0, ou seja, S ≤ 0) ∨ (N = SU ∧ S < 0) |
| D12F | (N ≠ SU ∧ S ≥ 1) ∨ (N = SU ∧ S ≥ 0) |
| D13T | D12T ∧ N ∈ {C, SU, BP} |
| D13F | N = AB ∧ S ≤ 0 |
| D14T | D12T ∧ N ∈ {C, SU} |
| D14F | N = BP ∧ S ≤ 0 |
| D15T | (N = C ∧ S ≤ 0 ∧ Q₄₂ > 0, ou seja, Q ≥ 2) ∨ (N = SU ∧ S < 0 ∧ Q > 0) |
| D15F | (N = C ∧ S ≤ 0 ∧ Q ≤ 1) ∨ (N = SU ∧ S < 0 ∧ Q ≤ 0) |
| D16T | N = C ∧ S ≤ 0 ∧ Q ≥ 2 |
| D16F | N = SU ∧ S < 0 ∧ Q > 0 |
| D17T | N = AB ∧ S ≤ 0 ∧ Q₅₁ < 50, ou seja, Q ≤ 48 |
| D17F | N = AB ∧ S ≤ 0 ∧ Q₅₁ ≥ 50, ou seja, Q ≥ 49 |

Nenhum ramo é inalcançável: todas as 34 restrições são satisfatíveis.

---

## Etapa 2: Seleção e rastreamento dos casos

| Caso | Name | SellIn | Quality |
|---|---|---|---|
| T1 | Sulfuras, Hand of Ragnaros | -1 | 80 |
| T2 | Item comum | 0 | 0 |
| T3 | Item comum | 0 | 10 |
| T4 | Aged Brie | 0 | 50 |
| T5 | Aged Brie | 0 | 10 |
| T6 | Backstage passes to a TAFKAL80ETC concert | 11 | 10 |
| T7 | Backstage passes to a TAFKAL80ETC concert | 5 | 10 |
| T8 | Backstage passes to a TAFKAL80ETC concert | 0 | 49 |

### T1: SU, S = -1, Q = 80
| Linha | Decisão / atribuição | SellIn | Quality |
|---|---|---|---|
| 9 | D1 = T (0 < 1) | -1 | 80 |
| 11 | D2 = T (não é AB nem BP) | -1 | 80 |
| 12 | D3 = T (80 > 0) | -1 | 80 |
| 13 | D4 = F (é SU) | -1 | 80 |
| 35 | D11 = F (é SU, sem decremento) | -1 | 80 |
| 39 | D12 = T (-1 < 0) | -1 | 80 |
| 40 | D13 = T | -1 | 80 |
| 41 | D14 = T | -1 | 80 |
| 42 | D15 = T (80 > 0) | -1 | 80 |
| 43 | D16 = F (é SU) | -1 | 80 |
| 9 | D1 = F (1 < 1) | | |

- Sequência: D1=T, D2=T, D3=T, D4=F, D11=F, D12=T, D13=T, D14=T, D15=T, D16=F, D1=F
- **Saída: SellIn = -1, Quality = 80**

### T2: C, S = 0, Q = 0
| Linha | Decisão / atribuição | SellIn | Quality |
|---|---|---|---|
| 9 | D1 = T | 0 | 0 |
| 11 | D2 = T | 0 | 0 |
| 12 | D3 = F (0 > 0 é falso) | 0 | 0 |
| 35 | D11 = T | | |
| 36 | SellIn = 0 − 1 | -1 | 0 |
| 39 | D12 = T (-1 < 0) | -1 | 0 |
| 40 | D13 = T | | |
| 41 | D14 = T | | |
| 42 | D15 = F (0 > 0 é falso) | -1 | 0 |
| 9 | D1 = F | | |

- Sequência: D1=T, D2=T, D3=F, D11=T, D12=T, D13=T, D14=T, D15=F, D1=F
- **Saída: SellIn = -1, Quality = 0**

### T3: C, S = 0, Q = 10
| Linha | Decisão / atribuição | SellIn | Quality |
|---|---|---|---|
| 9 | D1 = T | 0 | 10 |
| 11 | D2 = T | | |
| 12 | D3 = T (10 > 0) | | |
| 13 | D4 = T (não é SU) | | |
| 14 | Quality = 10 − 1 | 0 | 9 |
| 35 | D11 = T | | |
| 36 | SellIn = 0 − 1 | -1 | 9 |
| 39 | D12 = T | | |
| 40 | D13 = T | | |
| 41 | D14 = T | | |
| 42 | D15 = T (9 > 0) | | |
| 43 | D16 = T | | |
| 44 | Quality = 9 − 1 | -1 | 8 |
| 9 | D1 = F | | |

- Sequência: D1=T, D2=T, D3=T, D4=T, D11=T, D12=T, D13=T, D14=T, D15=T, D16=T, D1=F
- **Saída: SellIn = -1, Quality = 8**

### T4: AB, S = 0, Q = 50
| Linha | Decisão / atribuição | SellIn | Quality |
|---|---|---|---|
| 9 | D1 = T | 0 | 50 |
| 11 | D2 = F (é AB) | | |
| 18 | D5 = F (50 < 50 é falso) | 0 | 50 |
| 35 | D11 = T | | |
| 36 | SellIn = 0 − 1 | -1 | 50 |
| 39 | D12 = T | | |
| 40 | D13 = F (é AB) | | |
| 51 | D17 = F (50 < 50 é falso) | -1 | 50 |
| 9 | D1 = F | | |

- Sequência: D1=T, D2=F, D5=F, D11=T, D12=T, D13=F, D17=F, D1=F
- **Saída: SellIn = -1, Quality = 50**

### T5: AB, S = 0, Q = 10
| Linha | Decisão / atribuição | SellIn | Quality |
|---|---|---|---|
| 9 | D1 = T | 0 | 10 |
| 11 | D2 = F | | |
| 18 | D5 = T (10 < 50) | | |
| 19 | Quality = 10 + 1 | 0 | 11 |
| 20 | D6 = F (não é BP) | | |
| 35 | D11 = T | | |
| 36 | SellIn = 0 − 1 | -1 | 11 |
| 39 | D12 = T | | |
| 40 | D13 = F | | |
| 51 | D17 = T (11 < 50) | | |
| 52 | Quality = 11 + 1 | -1 | 12 |
| 9 | D1 = F | | |

- Sequência: D1=T, D2=F, D5=T, D6=F, D11=T, D12=T, D13=F, D17=T, D1=F
- **Saída: SellIn = -1, Quality = 12**

### T6: BP, S = 11, Q = 10
| Linha | Decisão / atribuição | SellIn | Quality |
|---|---|---|---|
| 9 | D1 = T | 11 | 10 |
| 11 | D2 = F | | |
| 18 | D5 = T | | |
| 19 | Quality = 10 + 1 | 11 | 11 |
| 20 | D6 = T | | |
| 21 | D7 = F (11 < 11 é falso) | | |
| 26 | D9 = F (11 < 6 é falso) | | |
| 35 | D11 = T | | |
| 36 | SellIn = 11 − 1 | 10 | 11 |
| 39 | D12 = F (10 < 0 é falso) | 10 | 11 |
| 9 | D1 = F | | |

- Sequência: D1=T, D2=F, D5=T, D6=T, D7=F, D9=F, D11=T, D12=F, D1=F
- **Saída: SellIn = 10, Quality = 11**

### T7: BP, S = 5, Q = 10
| Linha | Decisão / atribuição | SellIn | Quality |
|---|---|---|---|
| 9 | D1 = T | 5 | 10 |
| 11 | D2 = F | | |
| 18 | D5 = T | | |
| 19 | Quality = 10 + 1 | 5 | 11 |
| 20 | D6 = T | | |
| 21 | D7 = T (5 < 11) | | |
| 22 | D8 = T (11 < 50) | | |
| 23 | Quality = 11 + 1 | 5 | 12 |
| 26 | D9 = T (5 < 6) | | |
| 27 | D10 = T (12 < 50) | | |
| 28 | Quality = 12 + 1 | 5 | 13 |
| 35 | D11 = T | | |
| 36 | SellIn = 5 − 1 | 4 | 13 |
| 39 | D12 = F (4 < 0 é falso) | | |
| 9 | D1 = F | | |

- Sequência: D1=T, D2=F, D5=T, D6=T, D7=T, D8=T, D9=T, D10=T, D11=T, D12=F, D1=F
- **Saída: SellIn = 4, Quality = 13**

### T8: BP, S = 0, Q = 49
| Linha | Decisão / atribuição | SellIn | Quality |
|---|---|---|---|
| 9 | D1 = T | 0 | 49 |
| 11 | D2 = F | | |
| 18 | D5 = T (49 < 50) | | |
| 19 | Quality = 49 + 1 | 0 | 50 |
| 20 | D6 = T | | |
| 21 | D7 = T (0 < 11) | | |
| 22 | D8 = F (50 < 50 é falso) | 0 | 50 |
| 26 | D9 = T (0 < 6) | | |
| 27 | D10 = F (50 < 50 é falso) | 0 | 50 |
| 35 | D11 = T | | |
| 36 | SellIn = 0 − 1 | -1 | 50 |
| 39 | D12 = T | | |
| 40 | D13 = T (não é AB) | | |
| 41 | D14 = F (é BP) | | |
| 48 | Quality = 50 − 50 | -1 | 0 |
| 9 | D1 = F | | |

- Sequência: D1=T, D2=F, D5=T, D6=T, D7=T, D8=F, D9=T, D10=F, D11=T, D12=T, D13=T, D14=F, D1=F
- **Saída: SellIn = -1, Quality = 0**

### Resumo dos ramos cobertos (coluna `ramos` do código)

| Caso | Entrada (N, S, Q) | Saída (S, Q) | Ramos |
|---|---|---|---|
| T1 | SU, -1, 80 | -1, 80 | D1T D1F D2T D3T D4F D11F D12T D13T D14T D15T D16F |
| T2 | C, 0, 0 | -1, 0 | D1T D1F D2T D3F D11T D12T D13T D14T D15F |
| T3 | C, 0, 10 | -1, 8 | D1T D1F D2T D3T D4T D11T D12T D13T D14T D15T D16T |
| T4 | AB, 0, 50 | -1, 50 | D1T D1F D2F D5F D11T D12T D13F D17F |
| T5 | AB, 0, 10 | -1, 12 | D1T D1F D2F D5T D6F D11T D12T D13F D17T |
| T6 | BP, 11, 10 | 10, 11 | D1T D1F D2F D5T D6T D7F D9F D11T D12F |
| T7 | BP, 5, 10 | 4, 13 | D1T D1F D2F D5T D6T D7T D8T D9T D10T D11T D12F |
| T8 | BP, 0, 49 | -1, 0 | D1T D1F D2F D5T D6T D7T D8F D9T D10F D11T D12T D13T D14F |

---

## Etapa 3: Matriz de rastreabilidade

| Ramo | Caso(s) |
|---|---|
| D1T | T1, T2, T3, T4, T5, T6, T7, T8 |
| D1F | T1, T2, T3, T4, T5, T6, T7, T8 |
| D2T | T1, T2, T3 |
| D2F | T4, T5, T6, T7, T8 |
| D3T | T1, T3 |
| D3F | T2 |
| D4T | T3 |
| D4F | T1 |
| D5T | T5, T6, T7, T8 |
| D5F | T4 |
| D6T | T6, T7, T8 |
| D6F | T5 |
| D7T | T7, T8 |
| D7F | T6 |
| D8T | T7 |
| D8F | T8 |
| D9T | T7, T8 |
| D9F | T6 |
| D10T | T7 |
| D10F | T8 |
| D11T | T2, T3, T4, T5, T6, T7, T8 |
| D11F | T1 |
| D12T | T1, T2, T3, T4, T5, T8 |
| D12F | T6, T7 |
| D13T | T1, T2, T3, T8 |
| D13F | T4, T5 |
| D14T | T1, T2, T3 |
| D14F | T8 |
| D15T | T1, T3 |
| D15F | T2 |
| D16T | T3 |
| D16F | T1 |
| D17T | T5 |
| D17F | T4 |

**34 de 34 ramos cobertos (100% de Cobertura de Decisão).**

---

## Etapa 4: Justificativa de minimalidade (limite inferior = 8)

Cada caso tem um item e uma chamada, então executa o corpo do laço uma única vez. Logo, cada caso percorre um único caminho, que toma exatamente um lado de D2.

**1. Casos com D2T precisam ser pelo menos 3.** Depois de D2T, o caminho termina, de forma mutuamente exclusiva, em uma destas três opções:
- (a) D3F
- (b) D3T seguido de D4T
- (c) D3T seguido de D4F

Os ramos D3F, D4T e D4F são todos obrigatórios e cada um só aparece em uma dessas opções. Assim, nenhum caso cobre dois deles.

**2. Casos com D2F precisam ser pelo menos 5.** Depois de D2F, o caminho segue exatamente uma destas cinco opções mutuamente exclusivas:
- (a) D5F
- (b) D5T, D6F
- (c) D5T, D6T, D7F
- (d) D5T, D6T, D7T, D8T
- (e) D5T, D6T, D7T, D8F

Os ramos D5F, D6F, D7F, D8T e D8F são todos obrigatórios. Cada um pertence a uma única opção, e um caminho não pode tomar dois lados da mesma decisão, porque D1 só é True uma vez por caso.

**3. Total.** D2T e D2F não coexistem num mesmo caso, então o mínimo é 3 + 5 = 8. O conjunto T1 a T8 atinge esse limite, logo é mínimo.

Os demais ramos (D9 a D17 e D11F) foram encaixados nesses 8 caminhos sem precisar de caso extra:
- D16F e D11F no T1 (Sulfuras)
- D15F no T2
- D16T no T3
- D17F no T4
- D17T no T5
- D12F e D9F no T6
- D10T no T7
- D14F e D10F no T8

---

## Etapa 5: Código Go

```go
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
			id:     "T1",
			ramos:  "D1T D1F D2T D3T D4F D11F D12T D13T D14T D15T D16F",
			name:   sulfuras, sellIn: -1, quality: 80,
			esperadoSellIn: -1, esperadoQuality: 80,
		},
		{
			id:     "T2",
			ramos:  "D1T D1F D2T D3F D11T D12T D13T D14T D15F",
			name:   comum, sellIn: 0, quality: 0,
			esperadoSellIn: -1, esperadoQuality: 0,
		},
		{
			id:     "T3",
			ramos:  "D1T D1F D2T D3T D4T D11T D12T D13T D14T D15T D16T",
			name:   comum, sellIn: 0, quality: 10,
			esperadoSellIn: -1, esperadoQuality: 8,
		},
		{
			id:     "T4",
			ramos:  "D1T D1F D2F D5F D11T D12T D13F D17F",
			name:   agedBrie, sellIn: 0, quality: 50,
			esperadoSellIn: -1, esperadoQuality: 50,
		},
		{
			id:     "T5",
			ramos:  "D1T D1F D2F D5T D6F D11T D12T D13F D17T",
			name:   agedBrie, sellIn: 0, quality: 10,
			esperadoSellIn: -1, esperadoQuality: 12,
		},
		{
			id:     "T6",
			ramos:  "D1T D1F D2F D5T D6T D7F D9F D11T D12F",
			name:   backstage, sellIn: 11, quality: 10,
			esperadoSellIn: 10, esperadoQuality: 11,
		},
		{
			id:     "T7",
			ramos:  "D1T D1F D2F D5T D6T D7T D8T D9T D10T D11T D12F",
			name:   backstage, sellIn: 5, quality: 10,
			esperadoSellIn: 4, esperadoQuality: 13,
		},
		{
			id:     "T8",
			ramos:  "D1T D1F D2F D5T D6T D7T D8F D9T D10F D11T D12T D13T D14F",
			name:   backstage, sellIn: 0, quality: 49,
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
```

---

## Observações

1. **Condições da D2:** a suíte também exercita as combinações relevantes das duas condições atômicas da linha 11, embora o critério trate D2 como uma decisão única:
   - Item comum (T2, T3): (True, True)
   - Aged Brie (T4, T5): (False, com curto-circuito)
   - Backstage (T6 a T8): (True, False)
2. **Contagem da D1:** em cada caso, D1 é avaliada duas vezes, primeiro True e depois False. Uma lista vazia cobriria só D1F e não é necessária, porque todos os casos já cobrem D1F.
3. **Entradas fora de faixa:** o código não valida entradas. Por exemplo, T1 usa Quality = 80 para Sulfuras e o código apenas preserva esse valor. Valores negativos ou acima de 50 passam pelas guardas conforme as restrições da Etapa 1. Nenhum teste foi gerado para regras de faixa, porque não existe ramo que as imponha.
4. **Limite da cobertura de decisão:** 100% de decisão não garante cobertura de caminhos nem de valores-limite. Exemplos: Backstage com SellIn = 10 ou 6, ou Q = 48 levando a D10F, são limites de interesse que o critério pedido não exige.
