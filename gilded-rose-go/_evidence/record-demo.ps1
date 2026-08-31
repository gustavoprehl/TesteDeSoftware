# Roteiro de gravação para o vídeo — Pessoa 4 (evidências visuais).
# Rode cada bloco COM PAUSA entre eles (Read-Host abaixo espera Enter),
# para dar tempo de narrar e dar zoom no terminal antes de seguir.
#
# Pré-requisito: gotestsum instalado (dá saída colorida por teste).
#   go install gotest.tools/gotestsum@latest
# Se não tiver, troque "gotestsum --format testname --" por "go test" nos
# blocos abaixo (funciona igual, só sem as cores).

$ErrorActionPreference = "Continue"
Set-Location (Split-Path $PSScriptRoot -Parent)   # .../gilded-rose-go

function Pause($msg) {
    Write-Host ""
    Write-Host ">>> $msg" -ForegroundColor Cyan
    Read-Host "Pressione Enter para continuar"
}

Pause "1/6 - Código legado original (gildedrose/gildedrose.go) - mostrar a lógica condicional confusa"
Get-Content gildedrose\gildedrose.go

Pause "2/6 - RED #1: rodando a suite Few-shot EXATAMENTE como o ChatGPT gerou (erro de sintaxe)"
Push-Location _evidence\few_shot_original
go test .
Pop-Location

Pause "3/6 - RED #2: corrigindo so o typo (*testing.t -> *testing.T) ainda falha - struct inventada pela IA"
Push-Location _evidence\few_shot_typo_fixed_only
go test .
Pop-Location

Pause "4/6 - GREEN: suite Few-shot ja corrigida e integrada ao projeto"
& "$(go env GOPATH)\bin\gotestsum.exe" --format testname -- ./prompt_tests/few_shot/...

Pause "5/6 - GREEN: as 5 tecnicas de prompt rodando juntas, 100% passando"
& "$(go env GOPATH)\bin\gotestsum.exe" --format testname -- ./prompt_tests/...

Pause "6/6 - Cobertura: abrindo o relatorio HTML (verde = coberto)"
go test ./prompt_tests/... -coverpkg=./gildedrose/... -coverprofile=_evidence\coverage.out | Out-Null
go tool cover -html=_evidence\coverage.out -o _evidence\coverage.html
Start-Process _evidence\coverage.html

Write-Host ""
Write-Host "Fim do roteiro." -ForegroundColor Green
