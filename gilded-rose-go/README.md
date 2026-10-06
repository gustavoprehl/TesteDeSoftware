# GO Starter

- Run :

```shell
go run texttest_fixture.go [<number-of-days>; default: 2]
```

- Run tests :

```shell
go test ./...
```

- Run tests and coverage :

```shell
go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```

## Testes gerados por IA (prompt_tests/)

Além do kata original em [gildedrose/](gildedrose/), este diretório contém as suítes de testes
geradas pelo ChatGPT, a mesma LLM nas duas iterações: uma por técnica de prompt na Iteração 1
(documentadas em [`../docs/Registro_Prompts.pdf`](../docs/Registro_Prompts.pdf)) e a do prompt
estruturado na Iteração 2:

| Pacote | Técnica |
|---|---|
| [`prompt_tests/direto`](prompt_tests/direto) | Direto (baseline) |
| [`prompt_tests/chain_of_thought`](prompt_tests/chain_of_thought) | Chain-of-Thought |
| [`prompt_tests/persona_pattern`](prompt_tests/persona_pattern) | Persona Pattern |
| [`prompt_tests/few_shot`](prompt_tests/few_shot) | Few-shot |
| [`prompt_tests/spec_driven`](prompt_tests/spec_driven) | Spec-driven (bônus) |
| [`prompt_tests/estruturado`](prompt_tests/estruturado) | **Iteração 2:** Prompt Estruturado (caixa branca, baseado no CFG). Ver [`../docs/iteracao2/prompt-estruturado.md`](../docs/iteracao2/prompt-estruturado.md) |

Cada suíte é um pacote Go independente (testes black-box, importando
`github.com/emilybache/gildedrose-refactoring-kata/gildedrose`), para que nenhuma técnica
interfira nas outras — importante porque duas delas (Few-shot e Spec-driven) geraram funções
de teste com nomes idênticos.

Rodar só uma técnica:
```shell
go test ./prompt_tests/few_shot/... -v
```

Rodar tudo com cobertura atribuída ao pacote testado (os testes são black-box, então é
necessário `-coverpkg`):
```shell
go test ./prompt_tests/... -coverpkg=./gildedrose/... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

Resultado atual: as 6 suítes compilam e passam 100%, com 100% de cobertura de statement do
pacote `gildedrose`. Ver as seções "Intervenção manual na integração" e "Cobertura de statement
como falso positivo" em [`AUDITORIA-TESTES-IA.md`](AUDITORIA-TESTES-IA.md) para as correções
manuais que foram necessárias (responsabilidade da Dupla 2) e para uma observação sobre o que
essa cobertura de 100% não garante.

## Cobertura de Decisão (Iteração 2)

O `go test -cover` só mede statements. O medidor em [`_decisioncov/`](_decisioncov/) informa
quais dos 34 ramos (D1–D17, True/False) de `UpdateQuality` cada suíte percorre, usando uma
cópia instrumentada da função em uma pasta temporária (nenhum arquivo do repositório é
alterado):

```shell
go run ./_decisioncov ./_decisioncov/validacao   # validação: os 8 casos manuais dão 34/34
go run ./_decisioncov ./prompt_tests/...         # uma linha por suíte
go run ./_decisioncov -matriz ./prompt_tests/... # com a matriz ramo × suíte
```

## Auditoria manual (Dupla 3)

A auditoria crítica e a lista de lacunas estão em
[`AUDITORIA-TESTES-IA.md`](AUDITORIA-TESTES-IA.md). Os testes manuais de `Conjured`,
que expõem uma regra ausente na implementação e nas cinco suítes da IA, podem ser executados com:

```shell
go test ./manual_tests -v
```

Esses testes falham intencionalmente no estado atual do código, servindo como evidência visual da
lacuna encontrada.
