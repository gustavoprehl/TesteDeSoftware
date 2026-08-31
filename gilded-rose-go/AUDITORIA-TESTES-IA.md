# Auditoria crítica dos testes gerados por IA

## Veredito

Ainda não é possível considerar a suíte gerada pela IA completa. As cinco
suítes passam e atingem 100% de cobertura de statements porque reproduzem os
quatro comportamentos presentes em `gildedrose.go`. Nenhuma delas, porém,
especifica o comportamento de itens `Conjured`, que aparece no fixture do kata
como `Conjured Mana Cake`. A implementação trata esse item como um item normal.

O resultado é um falso positivo importante: 100% de cobertura de statements
sem cobertura de 100% das regras de negócio.

## Lacunas encontradas

### 1. `Conjured` não foi testado por nenhuma técnica (crítica)

Busca nas cinco suítes de `prompt_tests/`: não há nenhuma ocorrência de
`Conjured`. A omissão também existe na implementação: não há um ramo para esse
tipo em `gildedrose/gildedrose.go`.

Regra esperada:

A [especificação canônica do kata em português](https://github.com/emilybache/GildedRose-Refactoring-Kata/blob/main/GildedRoseRequirements_pt-BR.md)
declara que itens expirados degradam duas vezes mais rápido e que itens
`Conjured` degradam duas vezes mais rápido que os demais. Combinando as duas
regras, a taxa de um `Conjured` expirado é de 4 pontos por dia.

- antes da data de venda, a qualidade cai 2 pontos por dia;
- a partir da expiração, cai 4 pontos por dia (duas vezes a taxa de um item
  normal expirado);
- a qualidade nunca fica abaixo de 0;
- `SellIn` continua caindo 1 por dia.

Impacto observado no código atual:

| Cenário | Esperado | Atual |
|---|---:|---:|
| `SellIn=3`, `Quality=6` | `SellIn=2`, `Quality=4` | `SellIn=2`, `Quality=5` |
| `SellIn=0`, `Quality=6` | `SellIn=-1`, `Quality=2` | `SellIn=-1`, `Quality=4` |
| `SellIn=0`, `Quality=3` | `SellIn=-1`, `Quality=0` | `SellIn=-1`, `Quality=1` |

### 2. Os testes derivaram o oráculo do código, não da especificação (crítica)

Os prompts 1 a 4 pediram que a IA deduzisse as regras do próprio código. Como
o código já omitia `Conjured`, os testes copiaram a mesma omissão. Até o prompt
Spec-driven enumerou somente item normal, Aged Brie, Sulfuras e Backstage Pass.
Assim, todas as suítes podem passar quando código e testes compartilham o mesmo
erro de requisito.

O fixture de demonstração já continha `Conjured Mana Cake`, mas ele não foi
usado como fonte de casos de teste.

### 3. A afirmação de "suíte definitiva" da Persona é forte demais (média)

A suíte Persona é a mais completa para os quatro tipos implementados: cobre os
dois lados dos limiares 11/10 e 6/5 de Backstage, `SellIn` 0 e negativo, e as
transições para qualidade 0 e 50. Mesmo assim, omite completamente `Conjured`.
Logo, o ranking relativo das técnicas pode ser mantido, mas a conclusão de que
ela não tem ponto fraco deve ser corrigida.

### 4. Um teste de Sulfuras legitima um estado inválido (média)

A suíte Chain-of-Thought cria Sulfuras com `Quality=100` e exige que 100 seja
preservado. A especificação diz que Sulfuras tem qualidade imutável de 80. O
kata normalmente pressupõe que o item já entra no estado válido, portanto não
é obrigatório normalizar 100 para 80; ainda assim, chamar esse comportamento de
regra esperada amplia o contrato sem base na especificação e pode mascarar dados
inválidos. O teste deveria usar 80 ou ser marcado explicitamente como teste de
caracterização fora do domínio.

### 5. Predominam testes de uma única atualização (baixa)

As suítes verificam quase sempre um único chamado de `UpdateQuality`. Os casos
unitários cobrem bem as transições locais dos quatro tipos conhecidos, mas não
há um teste de trajetória que atravesse vários limiares ao longo de dias. Esse
tipo de teste não substitui os unitários; ele protege a integração temporal e é
útil como teste de caracterização.

### 6. Robustez de ponteiros não está coberta (fora da regra de negócio)

A Persona testa slice `nil` e slice vazio, que são seguros. Nenhuma suíte testa
um slice não vazio contendo um item `nil`; o código atual causa panic ao
desreferenciá-lo. Como o kata não define esse input como válido, isso deve ser
registrado como decisão de contrato, não automaticamente como defeito.

## O que as suítes cobrem bem

Consideradas em conjunto, as suítes da IA cobrem adequadamente:

- itens normais antes/depois da expiração e piso 0;
- Aged Brie antes/depois da expiração e teto 50;
- Sulfuras sem alteração de `SellIn` ou `Quality` nos estados válidos;
- Backstage nos limiares 11/10 e 6/5, no dia do show e após o show;
- transições de qualidade para 0 e 50;
- slice vazio, e slice `nil` na suíte Persona.

Isso torna `Conjured` a lacuna funcional mais clara e demonstrável da cobertura
conjunta, não apenas de uma técnica isolada.

## Testes manuais adicionados

Os testes estão em `manual_tests/manual_audit_test.go`. Eles exercitam
`Conjured` antes da expiração, na passagem para expirado e no piso zero.

Executar somente a evidência da auditoria:

```shell
go test ./manual_tests -v
```

No estado atual, os três subtestes devem falhar. Esse vermelho é intencional:
ele demonstra que a suíte da IA estava verde apesar de uma regra não
implementada. Para comparar:

```shell
go test ./prompt_tests/... -v
```

As cinco suítes geradas pela IA devem continuar verdes.

## Recomendação

Manter os testes manuais vermelhos como evidência da auditoria até que a equipe
decida implementar a regra `Conjured`. Depois da correção do código de produção,
esses mesmos testes devem permanecer como regressão e passar sem alteração.
