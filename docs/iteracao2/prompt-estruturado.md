# Prompt Estruturado e Nova Execução da IA - Iteração 2

## 1. Ponto de partida

Este documento usa o que os anteriores produziram e não refaz nenhum deles:

- A [análise estrutural](analise-estrutural.md) identificou as 17 decisões D1–D17 de `UpdateQuality`.
- O [CFG](cfg.md) montou o grafo (29 nós, 45 arestas) e calculou V(G) = 18, com os caminhos C1–C18.
- O [critério de cobertura](criterio-cobertura.md) escolheu a **Cobertura de Decisão** (34 ramos) e mostrou que o mínimo são **8 casos** (M1–M8).

A pergunta desta parte é a da seção 2.2 do enunciado: **a IA melhora quando recebe orientações estruturais?** Para responder, fizemos três coisas:

1. Escrevemos um novo prompt, técnico e estrutural, baseado no CFG.
2. Rodamos o prompt uma vez e integramos a suíte gerada, sem mexer nela.
3. Medimos a suíte e conferimos cada afirmação da IA contra o código real.

| Artefato | Local |
|---|---|
| Prompt exato, pronto para colar | [prompt-iteracao2.txt](prompt-iteracao2.txt) |
| Resposta completa da IA, sem edição | [resposta-ia-iteracao2.md](resposta-ia-iteracao2.md) |
| Suíte gerada | [prompt_tests/estruturado/estruturado_test.go](../../gilded-rose-go/prompt_tests/estruturado/estruturado_test.go) |
| Medidor de Cobertura de Decisão | [_decisioncov/](../../gilded-rose-go/_decisioncov/) |

## 2. Por que um prompt novo, e o que mudou

Os prompts da Iteração 1 ([Registro_Prompts.pdf](../Registro_Prompts.pdf)) pediam testes a partir do código ou da especificação, sem nenhuma informação sobre a estrutura. O prompt da Iteração 2 foi montado para responder, item por item, ao que deu errado ou funcionou naquela rodada:

| Observado na Iteração 1 | Elemento do novo prompt |
|---|---|
| O Direto testou só os casos óbvios. A IA não tinha como saber quais caminhos existiam. | Seção 2: tabela das 17 decisões, com linha, condição e destino True/False (o CFG em forma de tabela), mais as métricas do grafo. |
| Nenhuma técnica tinha uma meta mensurável. "Suíte completa" foi julgada no olho. | Seção 3: meta explícita de 100% de Cobertura de Decisão (34 ramos, com nomes D1T…D17F) e pedido do menor conjunto. |
| As suítes colocavam vários itens na mesma chamada, o que deixa a contagem de casos incomparável. | Definição de caso igual à usada no critério de cobertura: um item em uma chamada. |
| O CoT e a Persona melhoraram porque obrigaram a IA a raciocinar antes do código. | Seção 4: processo obrigatório em 5 etapas (condições de alcance, rastreamento linha a linha, matriz de rastreabilidade, minimalidade e só então o código). A persona foi mantida, agora como especialista em caixa branca. |
| O Few-shot usou `GildedRose{Items: items}` (struct inexistente) e `*testing.t`. | Seção 5: API real descrita explicitamente, com a proibição de inventar outra; `*testing.T` e pacote/import definidos. |
| As suítes da Iteração 1 precisaram ser adaptadas de `package gildedrose` para pacotes externos. | Seção 5 já pede `package estruturado_test` com o import do módulo. |
| O Spec-driven afirmou que 3 testes "deveriam falhar", mas eles passam. | Seção 6: todos os testes devem passar; a coluna `ramos` de cada caso tem que bater com o rastreamento. Essa coluna vira uma afirmação verificável (seção 5.2). |
| Nas guardas de `Quality`, o valor já foi alterado antes da comparação. É onde uma leitura superficial erra. | Seção 2, observações estruturais tiradas da seção 8 da análise estrutural. |

**O que o prompt deliberadamente não entrega.** O prompt tem a estrutura do código (decisões, destinos, V(G)), mas **não tem nenhum valor de entrada**, nem os casos M1–M8, nem as fronteiras que o critério de cobertura apontou como difíceis (Backstage com Quality 49, Aged Brie vencido com 50 etc.). Se esses valores estivessem no prompt, estaríamos medindo se a IA sabe copiar a resposta, e não se ela melhora com orientação estrutural. Pelo mesmo motivo, o prompt não menciona `Conjured`: ele pede só a estrutura do código, e o item `Conjured` não tem nenhum ramo no código.

## 3. Estrutura do prompt

O texto completo está em [prompt-iteracao2.txt](prompt-iteracao2.txt). Resumo das seções:

1. **Persona e tarefa.** Engenheiro de QA sênior de teste estrutural. A tarefa não é descobrir regras de negócio, e sim percorrer a estrutura com rastreabilidade.
2. **Código sob teste**, com números de linha (a tabela de decisões refere-se a eles).
3. **Modelo estrutural.** Granularidade (D2 composta = uma decisão), as 17 decisões com destinos True/False, as métricas do CFG e as observações estruturais.
4. **Critério.** 100% de Cobertura de Decisão, definição de caso, menor conjunto e prova de ramo inalcançável, se houver.
5. **Processo obrigatório.** Condições de alcance, rastreamento, matriz, minimalidade e código.
6. **Regras do código.** Pacote, imports, API real, table-driven com campos `id` e `ramos`, valores de caracterização.
7. **Restrições contra alucinação.**
8. **Formato da resposta.**

## 4. Protocolo de execução

| Item | Como foi feito |
|---|---|
| LLM | Claude (Anthropic), pelo Claude Code |
| Contexto | Conversa nova e isolada. A IA recebeu **somente o texto do prompt**, sem acesso ao repositório, aos documentos do grupo ou a ferramentas (não leu arquivos nem executou código). |
| Tentativas | Uma única execução, sem nova tentativa e sem pergunta de acompanhamento. O resultado é o primeiro que saiu. |
| Data | 27/09/2026 |
| Integração | Código da Etapa 5 copiado para `prompt_tests/estruturado/`. Acrescentamos só um comentário de cabeçalho e rodamos o `gofmt`, que mudou apenas espaços (conferido com `diff -w`). Nenhuma asserção foi alterada. |

**Limitação: a LLM mudou.** A Iteração 1 foi feita no ChatGPT, e esta execução não teve acesso a ele. Por isso a comparação Iteração 1 × Iteração 2 muda duas variáveis ao mesmo tempo, o prompt e o modelo, e isso precisa aparecer nos slides. Para isolar o efeito do prompt, basta repetir a execução no ChatGPT:

1. Abra uma conversa nova no ChatGPT, cole o conteúdo de [prompt-iteracao2.txt](prompt-iteracao2.txt) e guarde o link de compartilhamento.
2. Substitua o código de `prompt_tests/estruturado/estruturado_test.go` pelo bloco `go` da resposta.
3. Rode os comandos da seção 7. A medição é a mesma, e a tabela da seção 6 pode ser preenchida de novo.

## 5. Resultados

| Pergunta | Resultado |
|---|---|
| Compilou? | **Sim, de primeira**, sem nenhuma adaptação. Na Iteração 1, todas as suítes precisaram de adaptação de pacote, e o Few-shot não compilava. |
| Os testes passam? | **8 de 8** passam contra o código atual. |
| Cobertura de statements (`go test -cover`) | 100% |
| **Cobertura de Decisão (medida)** | **34/34 ramos = 100%** |
| Casos de teste gerados | **8**, igual ao mínimo mapeado na abordagem manual |
| Alucinações | **Não** foram encontradas (verificação na seção 5.2) |

### 5.1 Como a Cobertura de Decisão foi medida

O Go só mede statements. Para medir decisão, criamos o medidor `_decisioncov`:

- [`_decisioncov/instrumentado/gildedrose.go`](../../gilded-rose-go/_decisioncov/instrumentado/gildedrose.go) é uma cópia de `UpdateQuality` com a mesma lógica. A diferença é que cada lado (True e False) de D1–D17 ganhou um bloco próprio com um marcador. Nessa cópia, cobertura de statement dos marcadores = cobertura de decisão do original. A granularidade é a mesma do CFG (D2 continua sendo uma decisão).
- [`_decisioncov/main.go`](../../gilded-rose-go/_decisioncov/main.go) copia o módulo para uma pasta temporária e troca lá o `gildedrose.go` pela cópia instrumentada. Depois roda a suíte com `-coverprofile` e traduz os blocos executados para D1T…D17F. **Nenhum arquivo do repositório é modificado.**

Validação do medidor, antes de usá-lo na suíte da IA:

| Verificação | Esperado | Obtido |
|---|---|---|
| Os 8 casos manuais M1–M8 ([`_decisioncov/validacao`](../../gilded-rose-go/_decisioncov/validacao/validacao_test.go)) | 34/34 (seção 6 do critério de cobertura) | 34/34 |
| M1–M7, sem o M8 | faltar D8T, D10T e D14F (os ramos que só o M8 cobre) | faltaram D8T, D10T e D14F |
| Suíte da IA sem o T8 | faltar D8F, D10F e D14F (os ramos que só o T8 cobre) | faltaram D8F, D10F e D14F |

### 5.2 Conferência das afirmações da IA (alucinações)

Não nos limitamos a rodar os testes: conferimos tudo o que a IA **afirmou** na resposta.

| Afirmação da IA | Como conferimos | Resultado |
|---|---|---|
| API usada (`gildedrose.Item`, `gildedrose.UpdateQuality`) | Compilação | Correta, sem struct ou função inventada |
| Saídas esperadas dos 8 casos (Etapa 2) | Execução dos testes | 8/8 corretas |
| Coluna `ramos` de cada caso (o que cada teste cobre) | Medidor rodando **cada caso isoladamente** (`-run .../T1$` etc.) e comparando o conjunto medido com o afirmado | 8/8 idênticos. Nenhum ramo afirmado e não coberto, nenhum coberto e não declarado. |
| Matriz de rastreabilidade: 34/34 (Etapa 3) | Medidor na suíte inteira | 34/34 |
| As 34 condições de alcance (Etapa 1), incluindo as sutis: D8F só com Q = 49, D10F com 48 ≤ Q ≤ 49, D15T de item comum só com Q ≥ 2, D17F de Brie com Q ≥ 49 | Enumeração exaustiva: 4 tipos de item × SellIn de −5 a 15 × Quality de −5 a 85 = 7.644 entradas, comparando o ramo real (cópia instrumentada) com a condição escrita pela IA para cada um dos 34 ramos | **0 divergências** |
| Mínimo de 8 casos (Etapa 4) | Comparação com a prova do [critério de cobertura](criterio-cobertura.md#7-por-que-8-é-o-mínimo) | Correto. A IA usou outro argumento (3 caminhos após D2T + 5 após D2F), mas chegou ao mesmo limite do grupo (3 + 2 + 3). |
| "Nenhum ramo é inalcançável" | Suíte com 34/34 | Correto |
| Observação 4: "Q = 48 leva a D10F" e SellIn 10/6 não são exigidos pelo critério | Rastreamento manual e enumeração acima | Correto |

**Veredito: sem alucinação.** Na Iteração 1, houve alucinação de API (Few-shot) e de comportamento (Spec-driven).

### 5.3 Comparação com os casos manuais M1–M8

| Caso IA | Entrada (Nome, SellIn, Quality) | Equivalente manual | Comentário |
|---|---|---|---|
| T1 | Sulfuras, −1, 80 | M3 (C15) | Mesma entrada |
| T2 | Comum, 0, 0 | M2 (fusão C3 + C14) | Mesma entrada. A IA chegou sozinha à mesma fusão que o grupo. |
| T3 | Comum, 0, 10 | M1 (C13) | Mesma entrada |
| T4 | Aged Brie, 0, 50 | M5 (fusão C7 + C18) | Mesma entrada. Também chegou à mesma fusão. |
| T5 | Aged Brie, 0, 10 | M4 (C17) | Mesma entrada |
| T6 | Backstage, **11**, 10 | M6 (Backstage, 15, 10) | Mesmo caminho (C8), mas a IA escolheu o **valor exato da fronteira** de D7 (`SellIn < 11`) |
| T7 | Backstage, 5, 10 | M8 cobre D8T/D10T | A IA cobre os bônus sem vencer o ingresso (caminho C11) |
| T8 | Backstage, 0, **49** | M7 (D8F/D10F) + M8 (D14F) | Junta a saturação em 50 e o zeramento pós-show no mesmo caso |

A distribuição dos ramos entre os casos de Backstage é diferente da manual, mas o total é o mesmo: 8 casos. As duas fronteiras que o critério de cobertura apontou como indicadores de um teste que só se acha lendo o código (Backstage com Quality 49 e Aged Brie vencido com 50) **apareceram sem que o prompt as mencionasse**.

### 5.4 O que a suíte não cobre

Isto não é falha da IA: é o limite do critério pedido. Ainda assim, vale para o quadro comparativo e para a apresentação.

- **Caminhos básicos:** dos 18 caminhos de [cfg.md](cfg.md#7-caminhos-independentes), a suíte percorre exatamente 5 (C8, C11, C13, C15 e C17). Os outros 3 casos fazem fusões que geram caminhos fora dessa base. A base de V(G) não é única, então esse número serve só como referência de que 100% de decisão ≠ 100% de caminhos.
- **Fronteiras não exigidas pelo critério:** Backstage com SellIn 10 e 6 (o lado True mais próximo de D7 e D9) não foi testado. A própria IA registrou isso na Observação 4.
- **Requisitos ausentes do código:** a suíte não fala de `Conjured`. É o esperado, porque o prompt pediu cobertura da estrutura, e `Conjured` não tem ramo nenhum. Isso reforça a conclusão da auditoria da Iteração 1: **cobertura estrutural, mesmo 100% de decisão, não revela requisito que o código esqueceu.** Os testes de `manual_tests/` continuam sendo a única evidência dessa lacuna.

## 6. Dados para o quadro comparativo (coluna "Iteração 2 - Prompt Estruturado")

| Critério | Iteração 2 (Prompt Estruturado) |
|---|---|
| Cobertura de Decisão | **100% detectada** (34/34 ramos, medidos com `_decisioncov`) |
| Casos de Teste | **8 gerados** (mínimo necessário: 8) |
| Alucinações | **Não** |

Para a coluna da Iteração 1, o mesmo medidor pode ser usado nas 5 suítes antigas (comando abaixo), e assim os três números saem do mesmo instrumento.

## 7. Como reproduzir

A partir de `gilded-rose-go/`:

```bash
# roda a suíte da Iteração 2
go test ./prompt_tests/estruturado -v

# valida o medidor (deve dar 34/34)
go run ./_decisioncov ./_decisioncov/validacao

# Cobertura de Decisão da suíte da Iteração 2, com a matriz ramo × suíte
go run ./_decisioncov -matriz ./prompt_tests/estruturado

# um caso isolado (usado para conferir a coluna `ramos` de cada teste)
go run ./_decisioncov -run 'TestUpdateQuality_CoberturaDeDecisao/T8$' ./prompt_tests/estruturado

# todas as suítes de IA, das duas iterações (uma linha por suíte)
go run ./_decisioncov ./prompt_tests/...
```

A pasta `_decisioncov` começa com `_`, então o Go a ignora em `./...`. `go build ./...` e `go test ./...` continuam exatamente como antes.
