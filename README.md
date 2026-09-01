# AI Quality: O Desafio Gilded Rose - Go + ChatGPT

Trabalho da disciplina de Teste de Software (PUC) - auditoria de uma suíte de testes gerada por
IA sobre o kata [Gilded Rose](https://github.com/emilybache/GildedRose-Refactoring-Kata).

**Configuração do grupo**

|---|---|
| Linguagem | Go |
| LLM | ChatGPT |
| Kata base | [Gilded Rose Refactoring Kata (Emily Bache)](https://github.com/emilybache/GildedRose-Refactoring-Kata) |

## Estrutura do repositório

```
TP1/
├── docs/
│   ├── TP-Versao1.pdf                    # enunciado do trabalho
│   ├── Registro_Prompts.pdf              # prompts usados no ChatGPT + links das conversas
│   └── AI_Quality_Gilded_Rose.{pdf,pptx} # slides finais da apresentação
└── gilded-rose-go/
    ├── gildedrose/                       # código do kata (produção)
    ├── prompt_tests/                     # as 5 suítes geradas pelo ChatGPT, uma por técnica
    ├── manual_tests/                     # testes escritos manualmente após a auditoria
    ├── AUDITORIA-TESTES-IA.md            # veredito da auditoria: o que a IA acertou e errou
    └── _evidence/                        # prints e relatório de cobertura para vídeo/slides
```

## Rodando o projeto

```bash
cd gilded-rose-go
go build ./...

# as 5 suítes geradas pela IA
go test ./prompt_tests/... -v -cover

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
  Corrigidos os dois problemas, sem alterar nenhuma asserção, os 20 testes passaram, confirmando
  que a lógica da IA estava correta; o problema era puramente estrutural.
- Testes manuais cobrindo `Conjured` foram adicionados em `manual_tests/` e falham
  intencionalmente no estado atual do código, evidência visual da lacuna encontrada.

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

- Slides + relatório consolidado: [`docs/AI_Quality_Gilded_Rose.pdf`](docs/AI_Quality_Gilded_Rose.pdf) / [`.pptx`](docs/AI_Quality_Gilded_Rose.pptx)
- [`Vídeo-demonstração`]([docs/AI_Quality_Gilded_Rose.pdf](https://drive.google.com/file/d/1MBs7JPvyEo5RFLzwKYxWZbS09PH1rDXj/view?usp=sharing))
- Log de Crítica individual: entregue em sala, fora deste repositório
