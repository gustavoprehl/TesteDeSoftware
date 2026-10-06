# AI Quality: O Desafio Gilded Rose - Go + ChatGPT

Trabalho da disciplina de Teste de Software (PUC) - auditoria de uma suíte de testes gerada por
IA sobre o kata [Gilded Rose](https://github.com/emilybache/GildedRose-Refactoring-Kata).

**Configuração do grupo**

| Item | Valor |
|---|---|
| Linguagem | Go |
| LLM | ChatGPT (a mesma nas Iterações 1 e 2) |
| Kata base | [Gilded Rose Refactoring Kata (Emily Bache)](https://github.com/emilybache/GildedRose-Refactoring-Kata) |

## Estrutura do repositório

```
TP1/
├── docs/
│   ├── TP-Versao1.pdf                    # enunciado da Iteração 1
│   ├── TP-Versao2 (1).pdf                # enunciado da Iteração 2
│   ├── Registro_Prompts.pdf              # prompts da Iteração 1 no ChatGPT + links das conversas
│   ├── AI_Quality_Gilded_Rose.{pdf,pptx} # slides da Iteração 1
│   └── iteracao2/                        # CFG, critério, prompt estruturado, comparação e slides finais
└── gilded-rose-go/
    ├── gildedrose/                       # código do kata (produção)
    ├── prompt_tests/                     # suítes geradas pelo ChatGPT: 5 técnicas (Iteração 1) + estruturado (Iteração 2)
    ├── manual_tests/                     # testes escritos manualmente após a auditoria
    ├── _decisioncov/                     # medidor de Cobertura de Decisão (Iteração 2)
    ├── AUDITORIA-TESTES-IA.md            # veredito da auditoria: o que a IA acertou e errou
    └── _evidence/                        # prints e relatório de cobertura para vídeo/slides
```

## Rodando o projeto

```bash
cd gilded-rose-go
go build ./...

# as suítes geradas pela IA (Iterações 1 e 2)
go test ./prompt_tests/... -v -cover

# Cobertura de Decisão de cada suíte (Iteração 2)
go run ./_decisioncov ./prompt_tests/...

# os testes manuais da auditoria (falham de propósito - ver "Veredito" abaixo)
go test ./manual_tests -v
```

Detalhes de como rodar cada técnica isoladamente e gerar o relatório de cobertura em HTML estão
em [`gilded-rose-go/README.md`](gilded-rose-go/README.md).

## Metodologia: 5 técnicas de prompt

O código de `gildedrose.go` foi colado em 5 conversas independentes do ChatGPT (sem reaproveitar
contexto entre elas), cada uma pedindo a suíte de testes com uma técnica diferente. Prompts
completos, links das conversas e observações estão em
[`docs/Registro_Prompts.pdf`](docs/Registro_Prompts.pdf).

| Técnica | Compilou? | Cobertura de casos de borda | Ponto forte | Ponto fraco |
|---|:---:|---|---|---|
| 1. Direto (baseline) | Sim | Só valores já nos limites, não a transição até eles | Simplicidade | Superficial |
| 2. Chain-of-Thought | Sim | Testou transições (ex.: 49→50) | Raciocínio explícito antes do código | Não pegou `nil` × lista vazia |
| 3. Persona Pattern | Sim | Testou os dois lados de cada limiar | Único a identificar `nil` × lista vazia | - |
| 4. Few-shot | **Não** (corrigido, ver Veredito) | Boa cobertura, na teoria | Nomenclatura e padrão AAA consistentes | Typo de sintaxe + struct do exemplo copiada por engano |
| 5. Spec-driven (bônus) | Sim | Boa cobertura de fronteiras | Regras de negócio prontas, sem inferir do código | Texto disse que 3 testes "deveriam falhar" quando na verdade passam |

**Ranking de qualidade** (Dupla 1, antes da auditoria): Persona Pattern > Chain-of-Thought >
Spec-driven > Direto > Few-shot.

**Ranking revisto na Iteração 2**, pela Cobertura de Decisão medida (desempate: compilar sem
correção e não alucinar): Persona Pattern (100%) > Spec-driven (100%) > Few-shot (100%) >
Chain-of-Thought (94,1%) > Direto (91,2%). Detalhes em
[`docs/iteracao2/comparacao-final.md`](docs/iteracao2/comparacao-final.md).

## Veredito da auditoria

Análise completa em [`gilded-rose-go/AUDITORIA-TESTES-IA.md`](gilded-rose-go/AUDITORIA-TESTES-IA.md).
Resumo:

- **As 5 suítes passam 100% e atingem 100% de cobertura de statement, mas isso é um falso
  positivo.** Nenhuma das 5 técnicas testa itens `Conjured` (presentes no fixture do kata), e a
  implementação também não trata esse tipo, código e testes compartilham a mesma omissão de
  requisito. Mesmo o prompt Direto (o mais superficial, por avaliação da própria Dupla 1) chega a
  100% de cobertura de statement, evidenciando que essa métrica não mede qualidade de teste.
- **O Few-shot foi o único que não compilou.** Dois erros reais da IA: um typo de sintaxe
  (`*testing.t` em vez de `*testing.T`) e o uso da struct de exemplo do prompt
  (`GildedRose{Items: items}`) em vez da função real fornecida (`UpdateQuality(items []*Item)`).
  Corrigidos os dois problemas, sem alterar nenhuma asserção, os 24 testes passaram, confirmando
  que a lógica da IA estava correta; o problema era puramente estrutural.
- Testes manuais cobrindo `Conjured` foram adicionados em `manual_tests/` e falham
  intencionalmente no estado atual do código, evidência visual da lacuna encontrada.

## Iteração 2: análise estrutural × prompt estruturado

Enunciado em [`docs/TP-Versao2 (1).pdf`](<docs/TP-Versao2 (1).pdf>). Toda a documentação está em
[`docs/iteracao2/`](docs/iteracao2/):

| Etapa | Documento | Resultado |
|---|---|---|
| Análise estrutural | [`analise-estrutural.md`](docs/iteracao2/analise-estrutural.md) | 17 decisões (D1–D17) em `UpdateQuality` |
| CFG e complexidade | [`cfg.md`](docs/iteracao2/cfg.md) | 29 nós, 45 arestas, V(G) = 18, caminhos C1–C18 |
| Critério e casos manuais | [`criterio-cobertura.md`](docs/iteracao2/criterio-cobertura.md) | Cobertura de Decisão: 34 ramos, mínimo de 8 casos (M1–M8) |
| Prompt estruturado | [`prompt-estruturado.md`](docs/iteracao2/prompt-estruturado.md) | ChatGPT gerou 8 testes, 34/34 ramos, sem alucinação |
| Comparação final | [`comparacao-final.md`](docs/iteracao2/comparacao-final.md) | Quadro Iteração 1 × manual × Iteração 2 |

A Iteração 2 usou a mesma LLM da Iteração 1 (ChatGPT); só o prompt mudou. A cobertura de decisão
é medida com [`gilded-rose-go/_decisioncov/`](gilded-rose-go/_decisioncov/).

| Critério | Iteração 1 (Prompt Ingênuo) | Abordagem Manual | Iteração 2 (Prompt Estruturado) |
|---|---|---|---|
| Cobertura de Decisão | 91,2% detectada (31/34) | 100% mapeada (34/34) | 100% detectada (34/34) |
| Testes | 13 gerados | 8 necessários | 8 gerados |
| Alucinações | Não | N/A | Não |

"Prompt Ingênuo" é o prompt Direto da Iteração 1. Entre as outras técnicas da Iteração 1, Few-shot
(API inventada) e Spec-driven (afirmou falhas que não ocorrem) alucinaram.

Mesmo com 100% de decisão, `Conjured` continua sem teste na suíte estruturada, porque o código não
tem ramo para esse item: cobertura estrutural não revela requisito ausente.

## Evidências visuais

- Prints da execução (vermelho/verde) e cobertura:
  [`gilded-rose-go/_evidence/assets/Imagens/`](gilded-rose-go/_evidence/assets/Imagens/)
- Relatório de cobertura em HTML:
  [`gilded-rose-go/_evidence/coverage.html`](gilded-rose-go/_evidence/coverage.html)
- Reprodução ao vivo do erro de compilação original do Few-shot (antes da correção):
  [`gilded-rose-go/_evidence/few_shot_original/`](gilded-rose-go/_evidence/few_shot_original/) e
  [`gilded-rose-go/_evidence/few_shot_typo_fixed_only/`](gilded-rose-go/_evidence/few_shot_typo_fixed_only/)
  (ambos fora do build normal, só compilam se você entrar na pasta de propósito)

## Entregáveis

- Slides da Iteração 2 (versão final): [`docs/iteracao2/AI_Quality_Gilded_Rose_IT2.pptx`](docs/iteracao2/AI_Quality_Gilded_Rose_IT2.pptx)
- Slides + relatório consolidado da Iteração 1: [`docs/AI_Quality_Gilded_Rose.pdf`](docs/AI_Quality_Gilded_Rose.pdf) / [`.pptx`](docs/AI_Quality_Gilded_Rose.pptx)
- [Vídeo-demonstração](https://drive.google.com/file/d/1MBs7JPvyEo5RFLzwKYxWZbS09PH1rDXj/view?usp=sharing)
- Log de Crítica individual: entregue em sala, fora deste repositório
