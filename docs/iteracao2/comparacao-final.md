# Comparação final — Iteração 1 × Análise manual × Iteração 2

## 1. Objetivo da consolidação

Este documento fecha o projeto **AI Quality — O Desafio Gilded Rose** comparando três momentos:

1. **Iteração 1:** cinco suítes geradas por IA a partir de técnicas de prompt engineering.
2. **Análise manual:** modelagem estrutural da função `UpdateQuality`, escolha do critério de cobertura e definição do menor conjunto de casos.
3. **Iteração 2:** nova execução com prompt estrutural, baseado no grafo de fluxo de controle e em rastreabilidade obrigatória.

A comparação usa como eixo principal a **Cobertura de Decisão**, porque ela permite medir os dois lados de cada decisão do código, e não apenas se uma linha foi executada.

---

## 2. Base técnica usada na comparação

A análise estrutural identificou a função `UpdateQuality(items []*Item)` no pacote `gildedrose`, com 17 decisões relevantes, D1 a D17. A condição composta da linha 11 foi mantida como uma única decisão, D2.

O grafo de fluxo de controle foi modelado com:

| Métrica | Valor |
|---|---:|
| Nós | 29 |
| Arestas | 45 |
| Decisões | 17 |
| Ramos exigidos por Cobertura de Decisão | 34 |
| Complexidade Ciclomática V(G) | 18 |

A análise manual concluiu que todos os 34 ramos são alcançáveis e que o menor conjunto necessário para atingir 100% de Cobertura de Decisão tem **8 casos**.

---

## 3. Iteração 1 — Resultado das suítes geradas por IA

Na Iteração 1, foram avaliadas cinco técnicas de prompt:

| Técnica | Nº de testes | Resultado original |
|---|---:|---|
| Direto | 13 testes | Compilou e passou |
| Chain-of-Thought | 27 testes | Compilou e passou |
| Persona Pattern | 36 testes | Compilou e passou |
| Spec-driven | 23 testes | Compilou e passou, mas com afirmação incorreta de falha esperada |
| Few-shot | 24 testes | Não compilou inicialmente; passou após correção estrutural |

Cada teste é um cenário executável. Chain-of-Thought e Persona organizaram seus testes dentro de funções agrupadoras (`t.Run`); contamos os testes, não as funções.

O resultado agregado era forte sob a métrica tradicional: **123 testes, 0 falhas e 100% de cobertura de statements**. O runner reporta 131 porque também conta as 8 funções que apenas agrupam testes (4 do Chain-of-Thought e 4 da Persona). Porém, essa métrica não mostra se cada decisão foi exercitada nos dois sentidos.

### Medição posterior por Cobertura de Decisão

| Suíte | Testes | Cobertura de Decisão | Ramos não cobertos |
|---|---:|---:|---|
| Direto | 13 | 31/34 = **91,2%** | D15F, D16F, D17F |
| Chain-of-Thought | 27 | 32/34 = **94,1%** | D15F, D17F |
| Persona Pattern | 36 | 34/34 = **100%** | — |
| Spec-driven | 23 | 34/34 = **100%** | — |
| Few-shot | 24 | 34/34 = **100%** | — |

### Leitura da Iteração 1

O principal achado é que o ranking qualitativo inicial (Persona > Chain-of-Thought > Spec-driven > Direto > Few-shot) não coincide com a cobertura estrutural real. Persona, Spec-driven e Few-shot atingiram 100% de decisão, enquanto Direto e Chain-of-Thought deixaram ramos sem cobertura.

Refazendo o ranking com a Cobertura de Decisão como primeiro critério e, no empate, compilar sem correção e não alucinar:

| Posição | Técnica | Decisão | Desempate |
|---|---|---:|---|
| 1º | Persona Pattern | 100% | Compilou e não alucinou; único com slice nil e vazio |
| 2º | Spec-driven | 100% | Compilou, mas afirmou que 3 testes falhariam (eles passam) |
| 3º | Few-shot | 100% | Não compilou: struct inventada e `*testing.t` |
| 4º | Chain-of-Thought | 94,1% | Faltam D15F e D17F |
| 5º | Direto | 91,2% | Faltam D15F, D16F e D17F |

Mesmo assim, todas as suítes compartilham a mesma limitação: nenhuma testa `Conjured`. Esse ponto não aparece na cobertura estrutural porque o código não tem um ramo específico para esse item.

---

## 4. Análise manual — O papel da auditoria humana

A análise manual não começou escrevendo testes. Primeiro ela modelou o código:

1. Identificou as decisões D1–D17.
2. Construiu o CFG.
3. Calculou a complexidade ciclomática.
4. Escolheu Cobertura de Decisão como critério intermediário: mais forte que statement coverage e mais barato que cobertura de caminhos.
5. Mapeou as condições de alcance dos 34 ramos.
6. Selecionou o menor conjunto de casos.

### Casos manuais mínimos

| Caso | Tipo | Entrada `(SellIn, Quality)` | Saída esperada | Papel estrutural |
|---|---|---:|---:|---|
| M1 | Comum | `(0, 10)` | `(-1, 8)` | Item comum vencido com dupla degradação |
| M2 | Comum | `(0, 0)` | `(-1, 0)` | D3F e D15F sem qualidade negativa |
| M3 | Sulfuras | `(-1, 80)` | `(-1, 80)` | D11F e D16F |
| M4 | Aged Brie | `(0, 10)` | `(-1, 12)` | D17T |
| M5 | Aged Brie | `(0, 50)` | `(-1, 50)` | D5F e D17F |
| M6 | Backstage | `(15, 10)` | `(14, 11)` | D7F e D9F |
| M7 | Backstage | `(3, 49)` | `(2, 50)` | D8F e D10F |
| M8 | Backstage | `(0, 10)` | `(-1, 0)` | D14F e zeragem após show |

A análise manual também deixou claro o limite da métrica: **100% de Cobertura de Decisão não garante que uma regra ausente do código será descoberta**. O caso `Conjured` depende de comparação com a especificação do kata, não apenas com a estrutura implementada.

---

## 5. Iteração 2 — Prompt estrutural

A Iteração 2 usou a mesma LLM da Iteração 1 (ChatGPT) e um prompt novo, com papel de Engenheiro de QA Sênior e foco explícito em teste estrutural. O prompt forneceu:

- o código sob teste;
- o modelo estrutural D1–D17;
- a meta de 100% de Cobertura de Decisão;
- a definição de caso como um item por chamada;
- o processo obrigatório em cinco etapas;
- as regras exatas da API Go;
- restrições contra alucinação.

A IA não recebeu os casos manuais M1–M8 nem os valores de fronteira prontos. Ela recebeu a estrutura do problema e precisou escolher os casos.

### Resultado da Iteração 2

| Critério | Resultado |
|---|---:|
| Compilou de primeira | Sim |
| Testes passaram | 8/8 |
| Cobertura de statements | 100% |
| Cobertura de Decisão | 34/34 = **100%** |
| Testes gerados | 8 |
| Alucinações encontradas | Não |

A suíte estruturada chegou ao mesmo número mínimo da análise manual: **8 testes**. A coluna `ramos` de cada teste foi conferida isoladamente com o medidor (8/8 idênticas) e as 34 condições de alcance escritas pela IA foram comparadas com o código em 7.644 entradas, sem divergência.

Como a LLM foi a mesma nas duas iterações, a diferença observada vem do prompt.

### Onde a IA ainda falhou

- Percorreu 5 dos 18 caminhos independentes: 100% de decisão não é 100% de caminhos.
- Não testou as fronteiras SellIn 10 e 6 de Backstage, porque o critério não exigia.
- Usou Sulfuras com Quality 10, estado que a especificação não admite (fixa em 80): o mesmo padrão do Chain-of-Thought na Iteração 1.
- Não fala de `Conjured`, que não tem ramo no código.

---

## 6. Quadro comparativo final

| Critério | Iteração 1 — IA | Análise manual | Iteração 2 — IA estruturada |
|---|---|---|---|
| Base de entrada | Prompts por técnica, sem CFG formal | Código + CFG + critério escolhido | Prompt com CFG, 34 ramos e processo obrigatório |
| Objetivo explícito | Gerar testes unitários | Mapear e cobrir ramos | Gerar suíte mínima com 100% de decisão |
| Cobertura de statements | 100% em todas as suítes | Considerada insuficiente | 100% |
| Cobertura de Decisão | 91,2% a 100% | 100% mapeada | 100% medida |
| Nº de testes | 13, 27, 36, 23, 24 | 8 necessários | 8 gerados |
| Rastreabilidade | Parcial ou ausente | Completa, ramo a ramo | Completa e verificável |
| Alucinação | Encontrada em Few-shot e Spec-driven | Não aplicável | Não encontrada |
| Capacidade de achar `Conjured` | Não achou | Achou por auditoria de requisito | Não achou, pois não há ramo no código |
| Principal força | Rapidez e variedade | Rigor e critério | Rigor com automação |
| Principal limite | Pode parecer completo sem ser | Mais trabalhoso | Depende do modelo estrutural dado |

---

## 7. Conclusões finais

### 7.1 O que melhorou da Iteração 1 para a Iteração 2

A Iteração 2 melhorou principalmente em três pontos:

1. **Precisão estrutural:** a suíte atingiu 34/34 ramos com apenas 8 casos.
2. **Rastreabilidade:** cada caso declarou os ramos cobertos, e isso pôde ser conferido.
3. **Redução de alucinação:** a API real foi usada corretamente, sem inventar `GildedRose`, construtor ou bibliotecas externas.

### 7.2 O que a análise manual continua fazendo melhor

A análise manual continua indispensável porque define o critério, valida o CFG e compara o código contra a especificação do problema. Foi essa comparação que revelou a lacuna de `Conjured`.

### 7.3 Veredito consolidado

A IA é útil para acelerar a escrita de testes, especialmente quando recebe um prompt estrutural com metas verificáveis. Porém, **cobertura não substitui auditoria**. A suíte da Iteração 2 é excelente para caracterizar o código legado como ele está, mas não prova que o código implementa todos os requisitos esperados do Gilded Rose.

O melhor processo observado é híbrido:

1. o humano define o critério e audita a especificação;
2. a IA gera uma suíte rastreável;
3. ferramentas medem a cobertura;
4. a equipe mantém testes de requisito, como `Conjured`, para evidenciar lacunas reais do produto.

---

## 8. Frase final para apresentação

> A Iteração 2 mostra que a IA melhora quando recebe engenharia de teste, não apenas instruções genéricas. Mas o caso `Conjured` mostra o limite: 100% de cobertura do código não significa 100% de cobertura da especificação.
