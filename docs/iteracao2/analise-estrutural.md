# Análise Estrutural - UpdateQuality

## 1. Método analisado

**Arquivo:** [gilded-rose-go/gildedrose/gildedrose.go](../../gilded-rose-go/gildedrose/gildedrose.go), linhas 8–58.

**Nome exato:** `UpdateQuality`, função do pacote `gildedrose` (sem receptor; tecnicamente não é um método Go).

**Assinatura:** `func UpdateQuality(items []*Item)`. A fonte de verdade é exclusivamente essa implementação de produção.

## 2. Visão geral

O projeto utiliza Go, com `go 1.18` declarado em [go.mod](../../gilded-rose-go/go.mod), um pacote de produção e testes com o pacote `testing`. `Item` contém `Name string` e `SellIn, Quality int`. A função percorre o slice por índice e altera os objetos apontados, primeiro ajustando a qualidade, depois o prazo e por fim aplicando os ramos de prazo negativo.

A inspeção da árvore encontrou os seguintes materiais anteriores, identificados pelos READMEs do repositório:

| Local | Conteúdo |
|---|---|
| [README principal](../../README.md) e [README Go](../../gilded-rose-go/README.md) | Organização e execução do trabalho |
| [prompt_tests/](../../gilded-rose-go/prompt_tests/) | Suítes direto, chain_of_thought, persona_pattern, few_shot e spec_driven |
| [manual_tests/manual_audit_test.go](../../gilded-rose-go/manual_tests/manual_audit_test.go) | Testes manuais da auditoria |
| [AUDITORIA-TESTES-IA.md](../../gilded-rose-go/AUDITORIA-TESTES-IA.md) | Auditoria da etapa anterior |
| [docs/](../) | TP-Versao1.pdf, Registro_Prompts.pdf e apresentação AI_Quality_Gilded_Rose em PDF/PPTX |
| [_evidence/](../../gilded-rose-go/_evidence/) | Evidências e cópias históricas do Few-shot |
| [texttest_fixture.go](../../gilded-rose-go/texttest_fixture.go) | Programa demonstrativo que chama a função |

As cópias em `_evidence/` não são a implementação analisada. A documentação existente usa Markdown, títulos, tabelas, links relativos e blocos de código, predominantemente em português. Não havia estrutura da Iteração 2; estes documentos adotam `docs/iteracao2/`. Os materiais anteriores foram localizados para contextualização, não usados para definir os ramos.

## 3. Blocos de fluxo

São **29 blocos lógicos**, incluindo entrada, inicialização, decisões, operações, avanço do loop e saída. A granularidade é por decisão de fonte: a expressão composta de B4 permanece em um nó, com seu curto-circuito detalhado na seção 5. Não se trata de um CFG de instruções de máquina nem de predicados atômicos separados.

Os `else` não avaliam novas condições: são os ramos False de B4 (B8), B21 (B25) e B20 (B26). Junções sem instruções não recebem nós extras. O grafo descreve fluxo normal; exceções implícitas são delimitadas na seção 8.

### B1 - Entrada

Linha 8. Operação/estrutura: `Entrada em UpdateQuality`.

Segue para B2.

### B2 - Inicialização

Linha 9. Operação/estrutura: `i := 0`.

Segue para B3.

### B3 - Condição do loop

Linha 9. Decisão D1; condição original: `i < len(items)`.

True → B4; False → B29.

### B4 - Seleção inicial

Linha 11. Decisão D2; condição original: `items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert"`.

True → B5; False → B8.

### B5 - Qualidade positiva

Linha 12. Decisão D3; condição original: `items[i].Quality > 0`.

True → B6; False → B17.

### B6 - Excluir Sulfuras

Linha 13. Decisão D4; condição original: `items[i].Name != "Sulfuras, Hand of Ragnaros"`.

True → B7; False → B17.

### B7 - Decremento inicial

Linha 14. Operação/estrutura: `items[i].Quality = items[i].Quality - 1`.

Segue para B17.

### B8 - Limite inicial

Linha 18. Decisão D5; condição original: `items[i].Quality < 50`.

True → B9; False → B17.

### B9 - Incremento inicial

Linha 19. Operação/estrutura: `items[i].Quality = items[i].Quality + 1`.

Segue para B10.

### B10 - Identificar Backstage

Linha 20. Decisão D6; condição original: `items[i].Name == "Backstage passes to a TAFKAL80ETC concert"`.

True → B11; False → B17.

### B11 - Primeiro limiar

Linha 21. Decisão D7; condição original: `items[i].SellIn < 11`.

True → B12; False → B14.

### B12 - Limite do primeiro bônus

Linha 22. Decisão D8; condição original: `items[i].Quality < 50`.

True → B13; False → B14.

### B13 - Primeiro bônus

Linha 23. Operação/estrutura: `items[i].Quality = items[i].Quality + 1`.

Segue para B14.

### B14 - Segundo limiar

Linha 26. Decisão D9; condição original: `items[i].SellIn < 6`.

True → B15; False → B17.

### B15 - Limite do segundo bônus

Linha 27. Decisão D10; condição original: `items[i].Quality < 50`.

True → B16; False → B17.

### B16 - Segundo bônus

Linha 28. Operação/estrutura: `items[i].Quality = items[i].Quality + 1`.

Segue para B17.

### B17 - Atualizar prazo?

Linha 35. Decisão D11; condição original: `items[i].Name != "Sulfuras, Hand of Ragnaros"`.

True → B18; False → B19.

### B18 - Decrementar prazo

Linha 36. Operação/estrutura: `items[i].SellIn = items[i].SellIn - 1`.

Segue para B19.

### B19 - Prazo negativo?

Linha 39. Decisão D12; condição original: `items[i].SellIn < 0`.

True → B20; False → B28.

### B20 - Excluir Aged Brie

Linha 40. Decisão D13; condição original: `items[i].Name != "Aged Brie"`.

True → B21; False → B26.

### B21 - Excluir Backstage

Linha 41. Decisão D14; condição original: `items[i].Name != "Backstage passes to a TAFKAL80ETC concert"`.

True → B22; False → B25.

### B22 - Qualidade positiva após prazo

Linha 42. Decisão D15; condição original: `items[i].Quality > 0`.

True → B23; False → B28.

### B23 - Excluir Sulfuras após prazo

Linha 43. Decisão D16; condição original: `items[i].Name != "Sulfuras, Hand of Ragnaros"`.

True → B24; False → B28.

### B24 - Decremento após prazo

Linha 44. Operação/estrutura: `items[i].Quality = items[i].Quality - 1`.

Segue para B28.

### B25 - Zerar Backstage

Linha 48. Operação/estrutura: `items[i].Quality = items[i].Quality - items[i].Quality`.

Segue para B28.

### B26 - Limite de Brie após prazo

Linha 51. Decisão D17; condição original: `items[i].Quality < 50`.

True → B27; False → B28.

### B27 - Incremento de Brie após prazo

Linha 52. Operação/estrutura: `items[i].Quality = items[i].Quality + 1`.

Segue para B28.

### B28 - Avançar e retornar ao loop

Linha 9. Operação/estrutura: `i++`.

Segue para B3.

### B29 - Saída

Linha 58. Operação/estrutura: `Fim da função; retorno implícito, sem valor`.

Sem sucessor no fluxo normal.

## 4. Pontos de decisão

São **17 pontos de decisão**: a condição do `for` e os 16 `if` escritos na função. Predicados repetidos em posições diferentes são decisões diferentes.

Por convenção de modelagem adotada neste trabalho, cada expressão condicional completa do código-fonte representa uma única decisão estrutural no CFG. Assim, a expressão `A && B` de D2 corresponde a um único nó de decisão, B4. Seus predicados atômicos continuam documentados separadamente, junto com o comportamento de curto-circuito, na seção 5; eles não criam nós adicionais no CFG.

Para o cálculo da Complexidade Ciclomática nas etapas seguintes, deve ser mantida exatamente essa granularidade, em consistência com o CFG já produzido: a condição composta com `&&` continuará contando como uma única decisão estrutural, e seus predicados internos não devem ser contados como decisões adicionais. Não se deve usar uma granularidade diferente no cálculo de V(G) sem antes reconstruir formalmente o CFG.

| ID | Bloco | Condição original | Resultado True | Resultado False |
|----|-------|-------------------|----------------|-----------------|
| D1 | B3 | `i < len(items)` | B4 | B29 |
| D2 | B4 | `items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert"` | B5 | B8 |
| D3 | B5 | `items[i].Quality > 0` | B6 | B17 |
| D4 | B6 | `items[i].Name != "Sulfuras, Hand of Ragnaros"` | B7 | B17 |
| D5 | B8 | `items[i].Quality < 50` | B9 | B17 |
| D6 | B10 | `items[i].Name == "Backstage passes to a TAFKAL80ETC concert"` | B11 | B17 |
| D7 | B11 | `items[i].SellIn < 11` | B12 | B14 |
| D8 | B12 | `items[i].Quality < 50` | B13 | B14 |
| D9 | B14 | `items[i].SellIn < 6` | B15 | B17 |
| D10 | B15 | `items[i].Quality < 50` | B16 | B17 |
| D11 | B17 | `items[i].Name != "Sulfuras, Hand of Ragnaros"` | B18 | B19 |
| D12 | B19 | `items[i].SellIn < 0` | B20 | B28 |
| D13 | B20 | `items[i].Name != "Aged Brie"` | B21 | B26 |
| D14 | B21 | `items[i].Name != "Backstage passes to a TAFKAL80ETC concert"` | B22 | B25 |
| D15 | B22 | `items[i].Quality > 0` | B23 | B28 |
| D16 | B23 | `items[i].Name != "Sulfuras, Hand of Ragnaros"` | B24 | B28 |
| D17 | B26 | `items[i].Quality < 50` | B27 | B28 |

## 5. Condições compostas

Há **uma condição composta**, D2/B4 (linha 11):

```go
items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert"
```

- Predicado A: `items[i].Name != "Aged Brie"`.
- Predicado B: `items[i].Name != "Backstage passes to a TAFKAL80ETC concert"`.
- A=False: Go não avalia B; o resultado completo é False e segue para B8.
- A=True: Go avalia B; B=True leva a B5 e B=False leva a B8.

A avaliação ocorre da esquerda para a direita, com curto-circuito de `&&`. O nó B4 abstrai essa avaliação interna, preservando os destinos da expressão completa. Não há `||`, negação lógica unária `!` nem comparações encadeadas. `!=` é comparação de desigualdade, não uma negação unária separada. Os `if` aninhados da segunda metade continuam sendo decisões distintas, não condições compostas inventadas.

## 6. Alterações de estado

Todas as linhas abaixo pressupõem B3=True na iteração atual. As condições são avaliadas no instante em que cada decisão é alcançada: `Quality` pode ter sido alterada anteriormente e D12 usa `SellIn` após B18 quando esse bloco executa. As condições listadas identificam os requisitos de acesso a cada escrita; não enumeram caminhos completos.

| Bloco | Campo alterado | Operação original | Condição necessária |
|-------|----------------|-------------------|--------------------|
| B7 | Quality | `items[i].Quality = items[i].Quality - 1` | D2=True, D3=True e D4=True. |
| B9 | Quality | `items[i].Quality = items[i].Quality + 1` | D2=False e D5=True. |
| B13 | Quality | `items[i].Quality = items[i].Quality + 1` | D2=False, D5=True, D6=True, D7=True e D8=True. |
| B16 | Quality | `items[i].Quality = items[i].Quality + 1` | D2=False, D5=True, D6=True, D9=True e D10=True; D7/D8 podem ter produzido um incremento anterior. |
| B18 | SellIn | `items[i].SellIn = items[i].SellIn - 1` | D11=True, após a junção de todos os ramos iniciais. |
| B24 | Quality | `items[i].Quality = items[i].Quality - 1` | D12=True, D13=True, D14=True, D15=True e D16=True. |
| B25 | Quality | `items[i].Quality = items[i].Quality - items[i].Quality` | D12=True, D13=True e D14=False. |
| B27 | Quality | `items[i].Quality = items[i].Quality + 1` | D12=True, D13=False e D17=True. |

Há sete atribuições a `Quality` e uma a `SellIn`. B2 e B28 alteram apenas a variável local `i`. `Name` e o slice não recebem atribuições. B25 preserva a expressão de subtração do próprio campo, cujo resultado é zero.

## 7. Fluxo textual

Visão resumida; as transições completas constam em [cfg.md](cfg.md).

```text
B1 Entrada → B2 i := 0 → B3 Loop
  B3 False → B29 Saída
  B3 True → B4 Seleção inicial
    B4 True → B5/B6 → B7 quando ambas True → B17
      Qualquer False em B5/B6 → B17
    B4 False → B8 Limite inicial
      B8 False → B17
      B8 True → B9 Incremento → B10 Backstage?
        B10 False → B17
        B10 True → B11/B12 → B13 quando ambas True
          Junção em B14 (inclusive se B11 ou B12 False)
          B14/B15 → B16 quando ambas True → B17
          Qualquer False em B14/B15 → B17
  B17 True → B18 SellIn - 1 → B19
  B17 False → B19
  B19 False → B28
  B19 True → B20 Não é Brie?
    B20 False → B26 → B27 se True → B28
      B26 False → B28
    B20 True → B21 Não é Backstage?
      B21 False → B25 Zerar → B28
      B21 True → B22/B23 → B24 quando ambas True → B28
        Qualquer False em B22/B23 → B28
  B28 i++ → B3
```

## 8. Observações estruturais

- B14 é um segundo `if`, não um `else if` de B11. Após B13, B14 ainda é avaliado. As verificações de limite em B8, B12 e B15 leem o estado atualizado e foram preservadas.
- Os limiares B11/B14 usam o prazo anterior ao decremento. B19 usa o prazo posterior, exceto para Sulfuras, que não passa por B18.
- Sulfuras pode alcançar as decisões de prazo negativo, mas suas escritas são impedidas pelas verificações de nome. O código não normaliza sua qualidade para um valor específico.
- Backstage pode receber incrementos antes de B25 zerar sua qualidade. A escrita final não elimina do grafo as operações anteriores.
- O código não contém uma validação geral dos limites de qualidade de entrada; as guardas locais não equivalem a normalização de todos os valores.
- A auditoria anterior menciona uma regra esperada específica para `Conjured`. **O código faz:** esse nome segue o ramo genérico; não existe decisão específica para ele. **Expectativa registrada na documentação anterior:** tratamento diferenciado. Essa expectativa não acrescenta nós ao CFG.
- Slice vazio ou `nil` encerra em B3=False. Um elemento `nil` em slice não vazio causa `panic` ao acessar `Name` em B4. O CFG modela o fluxo normal com ponteiros válidos e não adiciona arestas de falhas implícitas do runtime. Essa delimitação deve ser revista por uma pessoa caso se deseje um CFG excepcional.
- Objetos são alterados por ponteiro. Se índices diferentes apontarem para o mesmo objeto, ele será atualizado novamente ao ser alcançado; não há cópia nem deduplicação.
- Não há `break`, `continue`, `return` explícito ou chamada auxiliar no corpo. O retorno ao loop ocorre por B28; B29 representa o encerramento implícito.
