# Evidências visuais — apoio para a Pessoa 4

Esta pasta existe só para gerar e guardar **evidências de execução** (terminal em
vermelho/verde, cobertura) para o vídeo e os slides. Nada aqui entra no build do projeto: o
nome `_evidence` começa com `_`, e o Go ignora por convenção qualquer diretório com esse
prefixo em `go build ./...` / `go test ./...`. Ou seja, pode rodar `go test .` diretamente
dentro das subpastas abaixo sem medo de "contaminar" a suíte principal.

## O que tem aqui

| Arquivo/pasta | Para quê |
|---|---|
| `record-demo.ps1` | Roteiro de gravação passo a passo (PowerShell), com pausas para narrar |
| `few_shot_original/` | Código do Few-shot **exatamente como o ChatGPT gerou** — não compila (RED #1: typo `*testing.t`) |
| `few_shot_typo_fixed_only/` | Mesmo código, só com o typo corrigido — ainda não compila (RED #2: struct inventada) |
| `coverage.out` / `coverage.html` / `coverage_summary.txt` | Cobertura gerada a partir de `prompt_tests/...` |

## Roteiro sugerido para o vídeo (vermelho → verde)

Rodar `./record-demo.ps1` de dentro desta pasta faz exatamente esta sequência, com pausa
entre cada passo:

1. **Mostrar o código legado** (`gildedrose/gildedrose.go`) — a lógica condicional confusa que
   está sendo testada.
2. **RED #1** — entrar em `_evidence/few_shot_original` e rodar `go test .`: reproduz o erro de
   compilação real que o ChatGPT cometeu na técnica Few-shot (`*testing.t` em vez de
   `*testing.T`). Dá pra dar zoom na mensagem de erro do compilador.
3. **RED #2** — entrar em `_evidence/few_shot_typo_fixed_only` e rodar `go test .`: mesmo
   corrigindo só o typo, ainda falha com `undefined: GildedRose` — mostra que a IA usou a
   struct do exemplo do prompt em vez da função real do código fornecido
   (`UpdateQuality(items []*Item)`).
4. **GREEN** — rodar a suíte Few-shot já corrigida e integrada
   (`prompt_tests/few_shot`), mostrando que as asserções da IA estavam corretas o tempo todo;
   o problema era só estrutural.
5. **GREEN (tudo)** — rodar as 5 técnicas juntas (`prompt_tests/...`), 100% passando.
6. **Cobertura** — abrir `coverage.html` no navegador (gerado no passo 5 do script): mostra o
   `gildedrose.go` com as linhas cobertas em verde.

Ver [`../LOG-CORRECOES.md`](../LOG-CORRECOES.md) para o texto de apoio a essa narrativa,
incluindo o achado de que **até o prompt Direto (baseline) chega a 100% de cobertura de
statement**, mesmo sendo a suíte mais superficial — bom gancho para explicar, na cobertura em
verde do passo 6, que "100% coberto" não é o mesmo que "bem testado".

## Comandos manuais (se preferir não usar o script)

```powershell
# RED #1
cd _evidence/few_shot_original
go test .

# RED #2
cd ../few_shot_typo_fixed_only
go test .

# GREEN — saída colorida por teste (requer gotestsum)
go install gotest.tools/gotestsum@latest
cd ../..
gotestsum --format testname -- ./prompt_tests/...

# Cobertura em HTML
go test ./prompt_tests/... -coverpkg=./gildedrose/... -coverprofile=_evidence/coverage.out
go tool cover -html=_evidence/coverage.out -o _evidence/coverage.html
```

## Dicas de gravação

- Aumentar a fonte do terminal antes de gravar (facilita o "zoom" pedido no enunciado).
- Terminais que colorem bem PASS/FAIL: Windows Terminal ou VS Code integrated terminal.
  `gotestsum` detecta automaticamente e usa verde/vermelho quando o terminal suporta.
  Sem `gotestsum`, `go test -v` funciona mas não colore, só escreve `PASS`/`FAIL` em texto.
- Gravador de tela nativo do Windows: `Win + Alt + R` (Xbox Game Bar) grava só a janela
  ativa — útil para não vazar outras informações da tela.
