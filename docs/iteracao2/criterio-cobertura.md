# Critério de Cobertura e Casos de Teste Manuais da UpdateQuality

## 1. Ponto de partida

Este documento dá continuidade aos três anteriores e não refaz nenhum deles. A [análise estrutural](analise-estrutural.md) identificou as 17 decisões da função `UpdateQuality` (D1 a D17). O [CFG](cfg.md) transformou essas decisões em um grafo com 29 nós e 45 arestas, e a seção 6 do mesmo arquivo calculou a Complexidade Ciclomática, V(G) = 18, com os 18 caminhos independentes C1 a C18.

A pergunta que respondemos aqui é outra: qual critério de cobertura vamos usar como referência para julgar os testes, e qual é o menor conjunto de casos que atinge 100% desse critério. Mantivemos a mesma granularidade dos documentos anteriores. A condição composta da linha 11 continua sendo uma única decisão (D2), exatamente como foi modelada no CFG e contada no cálculo de V(G).

## 2. Critério escolhido: Cobertura de Decisão

Escolhemos a **Cobertura de Decisão**, também chamada de cobertura de arestas ou de ramos. A regra é simples: toda decisão do código precisa ser avaliada como verdadeira em pelo menos um teste e como falsa em pelo menos um teste. No [CFG](cfg.md), isso equivale a percorrer todas as arestas que saem dos 17 nós de decisão.

Consideramos também as outras duas opções sugeridas no enunciado e tivemos três motivos para ficar com esta.

**Motivo 1: é o critério que o quadro comparativo mede.** Na seção 3 do enunciado da Iteração 2 ("O Entregável Final"), a primeira linha do quadro comparativo é "Cobertura de Decisão", com as colunas "% detectada" para a IA e "% mapeada" para a abordagem manual. Se escolhêssemos outro critério, a coluna manual não seria comparável com as colunas da IA.

**Motivo 2: a Cobertura de Comandos já se mostrou insuficiente.** Na Iteração 1, todas as suítes geradas pela IA atingiram 100% de cobertura de statements, como registrado em [coverage_summary.txt](../../gilded-rose-go/_evidence/coverage_summary.txt). Esse número parecia ótimo, mas a ferramenta `go test -cover` só verifica se cada linha foi executada, e não se cada condição foi testada nos dois sentidos. Um exemplo do próprio código mostra a diferença:

```go
if items[i].Name != "Sulfuras, Hand of Ragnaros" {   // linha 35, decisão D11
    items[i].SellIn = items[i].SellIn - 1
}
```

Qualquer teste com um item comum executa a linha 36, e a cobertura de comandos já considera esse trecho completo. Mas o lado falso da decisão, que é justamente a regra de que o Sulfuras nunca perde prazo, pode ficar sem nenhum teste. A Cobertura de Decisão obriga a existência desse teste. Por isso ela é o critério que permite descobrir o que o "100%" da Iteração 1 escondia.

**Motivo 3: a Cobertura de Caminhos já está representada.** Os 18 caminhos independentes calculados a partir de V(G) formam uma base de caminhos, que é um critério mais forte. Repetir esse trabalho não acrescentaria nada. Com a Cobertura de Decisão conseguimos mostrar um nível intermediário: mais exigente que a cobertura de comandos da Iteração 1 e mais barato que a base de caminhos. A seção 7 compara os três.

## 3. O que precisa ser coberto

Cada uma das 17 decisões tem dois resultados possíveis, então o critério exige **34 ramos** (17 × 2). A tabela abaixo lista, para cada decisão, o que precisa acontecer na entrada para que ela dê verdadeiro e para que dê falso. As condições foram copiadas da análise estrutural, e o bloco correspondente no CFG aparece entre parênteses.

| Decisão | Condição no código | Para dar verdadeiro | Para dar falso |
|---|---|---|---|
| D1 (B3) | `i < len(items)` | lista com pelo menos um item | fim da lista (acontece em toda execução) |
| D2 (B4) | não é Brie **e** não é Backstage | item comum ou Sulfuras | Aged Brie ou Backstage |
| D3 (B5) | `Quality > 0` | item do grupo de D2 verdadeiro com Quality > 0 | item do grupo de D2 verdadeiro com Quality = 0 |
| D4 (B6) | não é Sulfuras | item comum com Quality > 0 | Sulfuras com Quality > 0 |
| D5 (B8) | `Quality < 50` | Brie ou Backstage com Quality < 50 | Brie ou Backstage com Quality ≥ 50 |
| D6 (B10) | é Backstage | Backstage com Quality < 50 | Aged Brie com Quality < 50 |
| D7 (B11) | `SellIn < 11` | Backstage com SellIn ≤ 10 | Backstage com SellIn ≥ 11 |
| D8 (B12) | `Quality < 50` | Backstage que ainda está abaixo de 50 após o primeiro incremento | Backstage que chega a 50 no primeiro incremento (Quality inicial 49) |
| D9 (B14) | `SellIn < 6` | Backstage com SellIn ≤ 5 | Backstage com SellIn ≥ 6 |
| D10 (B15) | `Quality < 50` | Backstage ainda abaixo de 50 antes do segundo bônus | Backstage já em 50 antes do segundo bônus |
| D11 (B17) | não é Sulfuras | qualquer item que não seja Sulfuras | Sulfuras |
| D12 (B19) | `SellIn < 0` | prazo negativo após o decremento (SellIn inicial ≤ 0) | prazo ainda não negativo (SellIn inicial ≥ 1) |
| D13 (B20) | não é Brie | item vencido que não é Aged Brie | Aged Brie vencido |
| D14 (B21) | não é Backstage | item vencido que não é Brie nem Backstage | Backstage vencido |
| D15 (B22) | `Quality > 0` | item comum ou Sulfuras vencido com Quality > 0 neste ponto | item comum ou Sulfuras vencido com Quality = 0 neste ponto |
| D16 (B23) | não é Sulfuras | item comum vencido com Quality > 0 | Sulfuras vencido com Quality > 0 |
| D17 (B26) | `Quality < 50` | Aged Brie vencido abaixo de 50 neste ponto | Aged Brie vencido já em 50 neste ponto |

Antes de montar os casos, verificamos se algum desses 34 ramos seria impossível de alcançar, porque um ramo inalcançável tornaria 100% inatingível e teria de ser justificado à parte. Não encontramos nenhum. Os mais delicados são D8 falso e D10 falso. Como a decisão D5 só deixa passar itens com Quality abaixo de 50, e o incremento da linha 19 soma 1, a única forma de D8 dar falso é o Backstage começar com Quality exatamente 49. Já D10 falso aparece com Quality inicial 49, ou com 48 quando o primeiro bônus leva o ingresso a 50. É um valor de fronteira que um testador só encontra lendo o código, e por isso ele é um bom indicador para avaliar os testes da IA.

## 4. Como montamos os casos

Três decisões guiaram a construção dos casos.

**O que conta como um caso de teste.** Como `UpdateQuality` recebe uma lista, seria possível colocar oito itens em uma única chamada e cobrir tudo com "um teste". Não fizemos isso. Contamos cada caso como **um item em uma chamada de `UpdateQuality`**, com uma entrada e uma saída esperada. Assim, quando um caso falha, sabemos exatamente qual regra quebrou. Além disso, é assim que as suítes da IA na Iteração 1 organizaram seus testes, então a contagem fica comparável no quadro final.

**Começamos pelos caminhos independentes.** Em vez de inventar entradas do zero, partimos dos 18 caminhos de [cfg.md](cfg.md) e perguntamos quais deles eram necessários para cobrir os 34 ramos. Alguns ramos só aparecem em um único caminho da lista (D15 falso só em C14, D17 verdadeiro só em C17, D17 falso só em C18, D5 falso só em C7, D14 falso só em C16), o que obriga a incluir esses caminhos. Usando apenas caminhos daquela lista, o menor conjunto que encontramos tem 10: C3, C7, C8, C12, C13, C14, C15, C16, C17 e C18. Isso acontece porque cada caminho da base difere dos outros em poucas decisões, e por isso ramos que poderiam ser cobertos juntos acabam em caminhos separados.

**Juntamos o que podia ser juntado.** Em dois pontos, trocamos dois caminhos da base por um único caso, escolhendo valores de entrada que forçam os dois ramos na mesma execução:

| Caso novo | Substitui | Por que funciona |
|---|---|---|
| Comum, SellIn 0, Quality 0 | C3 (Comum, 5, 0) e C14 (Comum, 0, 1) | Com Quality 0, D3 dá falso logo no início. Com SellIn 0, o item vence e chega em D15, que também dá falso porque a Quality continua 0. |
| Aged Brie, SellIn 0, Quality 50 | C7 (Aged Brie, 5, 50) e C18 (Aged Brie, 0, 49) | Com Quality 50, D5 dá falso. Como o item vence, chega em D17, que também dá falso. |

Com essas duas fusões, os 10 caminhos viram 8 casos. Os outros seis casos são exatamente caminhos da lista de caminhos independentes, com os mesmos valores de entrada. Mantivemos esses valores de propósito, para que qualquer pessoa consiga conferir um caso daqui diretamente contra o caminho correspondente no CFG.

## 5. Casos de teste manuais

Nos casos com item comum usamos o nome `Elixir of the Mongoose`, que é um dos itens do próprio fixture do kata e não recebe nenhum tratamento especial no código. Nas tabelas ele aparece como "Comum", seguindo a convenção do [cfg.md](cfg.md).

| Caso | Nome | SellIn | Quality | SellIn esperado | Quality esperada | Origem | O que o caso verifica |
|---|---|---|---|---|---|---|---|
| M1 | Comum | 0 | 10 | -1 | 8 | C13 | Item comum no dia do vencimento perde qualidade em dobro |
| M2 | Comum | 0 | 0 | -1 | 0 | fusão de C3 e C14 | Qualidade nunca fica negativa, nem antes nem depois do prazo |
| M3 | Sulfuras | -1 | 80 | -1 | 80 | C15 | Sulfuras não perde prazo nem qualidade, mesmo vencido |
| M4 | Aged Brie | 0 | 10 | -1 | 12 | C17 | Aged Brie vencido ganha qualidade em dobro |
| M5 | Aged Brie | 0 | 50 | -1 | 50 | fusão de C7 e C18 | Aged Brie não passa de 50, nem antes nem depois do prazo |
| M6 | Backstage | 15 | 10 | 14 | 11 | C8 | Ingresso com mais de 10 dias ganha só 1 ponto |
| M7 | Backstage | 3 | 49 | 2 | 50 | C12 | Os bônus do ingresso param em 50 |
| M8 | Backstage | 0 | 10 | -1 | 0 | C16 | Ingresso é zerado depois do show, mesmo tendo recebido os bônus no mesmo dia |

## 6. Rastreabilidade: 34 de 34 ramos

A tabela abaixo mostra, para cada decisão, qual caso cobre o lado verdadeiro e qual cobre o lado falso. Quando mais de um caso cobre o mesmo ramo, listamos todos.

| Decisão | Verdadeiro coberto por | Falso coberto por |
|---|---|---|
| D1 | todos | todos |
| D2 | M1, M2, M3 | M4, M5, M6, M7, M8 |
| D3 | M1, M3 | M2 |
| D4 | M1 | M3 |
| D5 | M4, M6, M7, M8 | M5 |
| D6 | M6, M7, M8 | M4 |
| D7 | M7, M8 | M6 |
| D8 | M8 | M7 |
| D9 | M7, M8 | M6 |
| D10 | M8 | M7 |
| D11 | M1, M2, M4, M5, M6, M7, M8 | M3 |
| D12 | M1, M2, M3, M4, M5, M8 | M6, M7 |
| D13 | M1, M2, M3, M8 | M4, M5 |
| D14 | M1, M2, M3 | M8 |
| D15 | M1, M3 | M2 |
| D16 | M1 | M3 |
| D17 | M4 | M5 |

Nenhuma célula ficou vazia, portanto os 8 casos cobrem os **34 ramos, ou seja, 100% de Cobertura de Decisão**.

Para não depender só da leitura do código, executamos os 8 casos em uma cópia de `UpdateQuality` com um contador em cada decisão. As saídas obtidas foram exatamente as da seção 5, e o contador registrou os 34 ramos. Os ramos marcados na tabela acima são os que essa execução registrou.

## 7. Por que 8 é o mínimo

Depois de reduzir de 10 para 8, precisávamos saber se era possível reduzir mais. Não é, e o motivo está na forma como o código separa os itens por nome.

As decisões D3 e D4 só são alcançadas por itens que não são Aged Brie nem Backstage. As decisões D6 a D10 só são alcançadas por Aged Brie e Backstage, e D7 a D10 só por Backstage. Então os casos se dividem em três grupos que não se misturam, e cada grupo tem um mínimo próprio.

| Grupo | Ramos que forçam casos separados | Mínimo |
|---|---|---|
| Comum e Sulfuras | D3 falso exige Quality 0, e nesse caso D4 nem é avaliado. D4 verdadeiro exige um item comum com Quality > 0. D4 falso exige um Sulfuras com Quality > 0. São três itens diferentes. | 3 |
| Aged Brie | D6 falso só acontece quando D5 dá verdadeiro (Quality < 50). D5 falso exige Quality ≥ 50. Um mesmo Brie não pode estar nas duas situações. | 2 |
| Backstage | D7 falso exige SellIn ≥ 11, e aí D8 nem é avaliado. D8 falso exige que o ingresso chegue a 50, e daí em diante D10 também dá falso. D10 verdadeiro exige um ingresso que ainda esteja abaixo de 50 depois de D8. São três ingressos diferentes. | 3 |

Somando os grupos, 3 + 2 + 3 = **8 casos no mínimo**, que é exatamente o tamanho do conjunto da seção 5.

Uma observação: não incluímos o caso da lista vazia (C1). O lado falso de D1 é atingido no final de toda execução, quando o laço termina, então a lista vazia não acrescenta nenhum ramo novo para este critério. Ela continua sendo um teste útil de robustez, mas não é necessária para os 100%.

## 8. Comparação entre os critérios

| Critério | O que exige | Casos necessários | Situação da IA na Iteração 1 |
|---|---|---|---|
| Comandos | Toda linha executada pelo menos uma vez | poucos | 100%, segundo o `go test -cover` |
| **Decisão (este documento)** | Toda decisão avaliada como verdadeira e como falsa | **8** | ainda não medido; será feito com a tabela da seção 6 |
| Caminhos básicos (V(G)) | Todos os 18 caminhos independentes percorridos | 18 | ainda não medido |

A tabela mostra que cada critério pede mais testes que o anterior, porque é mais exigente. Os 8 casos garantem que nenhuma condição ficou testada só pela metade, mas não percorrem todas as combinações de caminhos. Para isso seriam necessários os 18 casos da base.

## 9. Uso deste documento nas próximas etapas

**Quadro comparativo.** Na coluna "Abordagem Manual (Teórica)", a linha de Cobertura de Decisão fica com **100% mapeada** e a linha de Casos de Teste fica com **8 casos necessários**.

**Medição da "% detectada" nas suítes da IA.** Como o Go não mede cobertura de decisão, a medição é feita com a tabela da seção 3. Para cada teste de uma suíte, identificamos a entrada usada e marcamos quais dos 34 ramos ela percorre. No final, a porcentagem da suíte é o número de ramos marcados dividido por 34. O mesmo procedimento vale para as suítes da Iteração 1 e para a suíte gerada pelo prompt estruturado, e assim as duas colunas da IA ficam comparáveis com a coluna manual.

**Prompt estruturado.** Os pontos que mais provavelmente vão diferenciar uma suíte superficial de uma suíte completa são os ramos que exigem valores de fronteira ou combinações pouco óbvias. São eles: Backstage com Quality 49 (D8 e D10 falsos), Backstage com SellIn 11, 6 e 0 (limites de D7, D9 e D12), Sulfuras já vencido e com Quality positiva (D16 falso), item comum vencido com Quality 0 (D15 falso) e Aged Brie vencido com Quality 50 (D17 falso). Vale pedir explicitamente esses casos no prompt, ou pedir que a IA justifique quais decisões cada teste cobre.
