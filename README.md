# TesteDeSoftware — TP1: AI Quality — O Desafio Gilded Rose

Trabalho da disciplina de Teste de Software (PUC).

**Configuração do grupo**
- Linguagem: Go
- LLM: ChatGPT
- Kata base: [Gilded Rose Refactoring Kata (Emily Bache)](https://github.com/emilybache/GildedRose-Refactoring-Kata)

## Estrutura

- [`gilded-rose-go/`](gilded-rose-go/) — código do kata em Go (obtido via sparse-checkout do
  repositório oficial, pasta `go/`), com os testes gerados pela IA integrados.
- [`Registro_Prompts.pdf`](Registro_Prompts.pdf) — prompts utilizados com o ChatGPT (Direto,
  Chain-of-Thought, Persona Pattern, Few-shot, Spec-driven), links das conversas e análise
  comparativa das técnicas.
- [`TP-Versao1.pdf`](TP-Versao1.pdf) — enunciado do trabalho.

## Rodando os testes

```bash
cd gilded-rose-go
go build ./...
go test ./... -v -cover
```
