# Log de Correções — Integração dos testes gerados pela IA

Responsabilidade da Dupla 2: "Corrigir erros de compilação/sintaxe que a IA cometer para os
testes rodarem." Este documento registra exatamente o que precisou de intervenção manual, para
alimentar a análise crítica da Dupla 3 e o "Veredito" do vídeo.

## Como os testes foram integrados

As 5 suítes descritas no `Registro_Prompts.pdf` foram colocadas em pacotes Go separados sob
`gilded-rose-go/prompt_tests/<tecnica>/`, cada um testando o pacote real `gildedrose` de fora
(black-box), no estilo do placeholder original do kata (`package gildedrose_test`).

**Por que pacotes separados, e não um arquivo único:** todas as 5 conversas foram escritas
como `package gildedrose` (white-box) de forma independente umas das outras. Colocá-las juntas
no mesmo diretório causaria colisão de nomes — em particular, **Few-shot** e **Spec-driven**
geraram funções com nomes idênticos (`TestQualityDecreasesByOneForNormalItemBeforeSellDate` e
`TestQualityDecreasesByTwoForNormalItemAfterSellDate`), o que não compila em Go se estiverem no
mesmo pacote.

**Regra seguida em todas as correções:** só o necessário para compilar/rodar foi alterado.
Nenhuma asserção, nenhum valor esperado (`want`/expected) foi tocado — se um teste falhasse por
expectativa errada da IA, isso ficaria evidente no `go test` e seria material de auditoria, não
algo para a Dupla 2 "consertar" silenciosamente. Na prática, depois das correções abaixo, os 5
pacotes passam 100%.

## Correções aplicadas

### Todas as 5 suítes
Adaptação mecânica de white-box para black-box: qualificação de `Item` → `gildedrose.Item` e de
`UpdateQuality(...)` → `gildedrose.UpdateQuality(...)`, com o import
`github.com/emilybache/gildedrose-refactoring-kata/gildedrose`. Não é uma correção de erro da
IA — é a adaptação necessária para reunir as 5 suítes num único repositório sem colisão de
pacote/nome.

### Few-shot (`prompt_tests/few_shot/few_shot_test.go`) — única que não compilava
A própria Dupla 1 já havia identificado os dois problemas no `Registro_Prompts.pdf`; a correção
efetiva foi feita aqui:

1. **Typo de sintaxe**: `func TestQualityIncreasesByTwoForAgedBrieAtQualityFortyNineAfterSellDate(t *testing.t)`
   — `*testing.t` (minúsculo) não existe em Go; deveria ser `*testing.T`. Esse erro sozinho
   impedia a compilação de todo o pacote.
2. **Erro estrutural**: a IA seguiu o exemplo do prompt (que usava `gr := GildedRose{Items: items}; gr.UpdateQuality()`)
   em vez de usar a assinatura real do código fornecido (`func UpdateQuality(items []*Item)`).
   O tipo `GildedRose` não existe neste pacote — só `Item` e a função `UpdateQuality`. Todas as
   ~20 chamadas foram trocadas para `gildedrose.UpdateQuality(items)`.

Depois dessas duas correções, os 20 testes do Few-shot passam sem nenhuma outra alteração — ou
seja, a lógica dos testes em si (asserções e valores esperados) estava correta; o problema era
puramente estrutural/sintático, confirmando a observação da Dupla 1.

## Resultado após integração

```
go test ./prompt_tests/... -v
ok  .../prompt_tests/chain_of_thought
ok  .../prompt_tests/direto
ok  .../prompt_tests/few_shot
ok  .../prompt_tests/persona_pattern
ok  .../prompt_tests/spec_driven
```

Cobertura (statement) do pacote `gildedrose`, por técnica:

| Técnica | Compila? | Passa 100%? | Cobertura de statement |
|---|---|---|---|
| Direto | Sim | Sim | 100% |
| Chain-of-Thought | Sim | Sim | 100% |
| Persona Pattern | Sim | Sim | 100% |
| Few-shot | **Não → corrigido** | Sim | 100% |
| Spec-driven | Sim | Sim | 100% |

## Achado para o "Veredito" (Fase 2 / Dupla 3)

Um ponto que vale destacar no vídeo: **mesmo o prompt Direto (baseline)**, que a própria Dupla 1
classificou como "superficial" (testa só valores já parados nos limites, não a transição até
eles), **atinge 100% de cobertura de statement**. Isso é evidência concreta de que cobertura de
linha/statement não mede qualidade de teste — o `UpdateQuality` tem só um caminho de código por
tipo de item, então basta *executar* cada ramo uma vez para chegar a 100%, independentemente de
testar os valores de fronteira corretos (`Quality` chegando exatamente a 0/50/50→51 etc.). As
suítes mais profundas (Persona Pattern, Chain-of-Thought) não aumentam a cobertura de statement
(já estava em 100%), mas aumentam a cobertura de *casos de borda* — uma métrica que `go tool
cover` não captura. Vale citar isso explicitamente na análise crítica: "100% de cobertura não
significa suíte adequada".
