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

Além do kata original em [gildedrose/](gildedrose/), este diretório contém as 5 suítes de
testes geradas pelo ChatGPT (uma por técnica de prompt), documentadas em
[`../Registro_Prompts.pdf`](../Registro_Prompts.pdf):

| Pacote | Técnica |
|---|---|
| [`prompt_tests/direto`](prompt_tests/direto) | Direto (baseline) |
| [`prompt_tests/chain_of_thought`](prompt_tests/chain_of_thought) | Chain-of-Thought |
| [`prompt_tests/persona_pattern`](prompt_tests/persona_pattern) | Persona Pattern |
| [`prompt_tests/few_shot`](prompt_tests/few_shot) | Few-shot |
| [`prompt_tests/spec_driven`](prompt_tests/spec_driven) | Spec-driven (bônus) |

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

Resultado atual: as 5 suítes compilam e passam 100%, com 100% de cobertura de statement do
pacote `gildedrose`. Ver [`../LOG-CORRECOES.md`](../LOG-CORRECOES.md) para as correções manuais
que foram necessárias (responsabilidade da Dupla 2) e para uma observação sobre o que essa
cobertura de 100% não garante — material de apoio para a auditoria da Dupla 3.

## Auditoria manual (Dupla 3)

A auditoria crítica e a lista de lacunas estão em
[`AUDITORIA-TESTES-IA.md`](AUDITORIA-TESTES-IA.md). Os testes manuais de `Conjured`,
que expõem uma regra ausente na implementação e nas cinco suítes da IA, podem ser executados com:

```shell
go test ./manual_tests -v
```

Esses testes falham intencionalmente no estado atual do código, servindo como evidência visual da
lacuna encontrada.
