// Medidor de Cobertura de Decisão para UpdateQuality.
//
// O Go só mede cobertura de statements. Este programa copia o módulo para uma
// pasta temporária, troca lá o arquivo gildedrose/gildedrose.go pela cópia
// instrumentada em _decisioncov/instrumentado/gildedrose.go e roda as suítes
// nessa cópia. Nenhum arquivo do repositório é alterado. Em seguida lê o
// coverprofile e informa quais dos 34 ramos (D1–D17, lados True e False) cada
// suíte percorreu.
//
// (go test -overlay seria mais simples, mas a instrumentação de cobertura do
// Go ignora o overlay e mede o arquivo original.)
//
// Uso, a partir de gilded-rose-go/:
//
//	go run ./_decisioncov ./prompt_tests/estruturado
//	go run ./_decisioncov ./prompt_tests/...          (uma linha por suíte)
//	go run ./_decisioncov -matriz ./prompt_tests/...  (inclui a matriz ramo × suíte)
//	go run ./_decisioncov -run 'TestX/T3' ./pacote     (mede um único caso)
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const arquivoProducao = "gildedrose/gildedrose.go"
const arquivoInstrumentado = "_decisioncov/instrumentado/gildedrose.go"

type posicao struct{ linha, coluna int }

type bloco struct {
	inicio, fim posicao
	contagem    int
}

type resultado struct {
	pacote   string
	passou   bool
	cobertos map[string]bool
}

var reMarcador = regexp.MustCompile(`marcador\("(D\d+[TF])"\)`)
var reBloco = regexp.MustCompile(`^(.+):(\d+)\.(\d+),(\d+)\.(\d+) (\d+) (\d+)$`)

func main() {
	matriz := flag.Bool("matriz", false, "imprime a matriz ramo × suíte")
	filtro := flag.String("run", "", "repassado ao go test -run (ex.: TestX/T3 mede um caso isolado)")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "uso: go run ./_decisioncov [-matriz] <pacotes de teste...>")
		os.Exit(2)
	}

	raiz, err := raizDoModulo()
	checar(err)
	marcadores, err := lerMarcadores(filepath.Join(raiz, arquivoInstrumentado))
	checar(err)
	ramos := ordenarRamos(marcadores)

	tmp, err := os.MkdirTemp("", "decisioncov")
	checar(err)
	defer os.RemoveAll(tmp)

	pacotes, err := expandirPacotes(raiz, flag.Args())
	checar(err)

	copia := filepath.Join(tmp, "modulo")
	checar(copiarModulo(raiz, copia))
	instrumentado, err := os.ReadFile(filepath.Join(raiz, arquivoInstrumentado))
	checar(err)
	checar(os.WriteFile(filepath.Join(copia, arquivoProducao), instrumentado, 0o644))

	var resultados []resultado
	for i, pacote := range pacotes {
		perfil := filepath.Join(tmp, fmt.Sprintf("cover%d.out", i))
		args := []string{"test", "-count=1", "-coverpkg=./gildedrose", "-coverprofile", perfil}
		if *filtro != "" {
			args = append(args, "-run", *filtro)
		}
		cmd := exec.Command("go", append(args, pacote)...)
		cmd.Dir = copia
		saida, errTeste := cmd.CombinedOutput()
		blocos, errPerfil := lerPerfil(perfil)
		if errPerfil != nil {
			fmt.Fprintf(os.Stderr, "%s não gerou coverprofile (erro de compilação?):\n%s\n", pacote, saida)
			continue
		}
		r := resultado{pacote: pacote, passou: errTeste == nil, cobertos: map[string]bool{}}
		for ramo, pos := range marcadores {
			for _, b := range blocos {
				if b.contagem > 0 && !antes(pos, b.inicio) && !antes(b.fim, pos) {
					r.cobertos[ramo] = true
				}
			}
		}
		resultados = append(resultados, r)
	}

	fmt.Printf("%-32s %-8s %-15s %s\n", "SUÍTE", "TESTES", "DECISÃO", "RAMOS NÃO COBERTOS")
	for _, r := range resultados {
		status := "passam"
		if !r.passou {
			status = "FALHAM"
		}
		var faltando []string
		for _, ramo := range ramos {
			if !r.cobertos[ramo] {
				faltando = append(faltando, ramo)
			}
		}
		n := len(ramos) - len(faltando)
		fmt.Printf("%-32s %-8s %2d/%d (%5.1f%%)  %s\n", nomeCurto(r.pacote), status, n, len(ramos),
			100*float64(n)/float64(len(ramos)), strings.Join(faltando, " "))
	}

	if *matriz {
		fmt.Println()
		fmt.Printf("%-6s", "RAMO")
		for i := range resultados {
			fmt.Printf(" S%-2d", i+1)
		}
		fmt.Println()
		for _, ramo := range ramos {
			fmt.Printf("%-6s", ramo)
			for _, r := range resultados {
				marca := "."
				if r.cobertos[ramo] {
					marca = "x"
				}
				fmt.Printf(" %-3s", marca)
			}
			fmt.Println()
		}
		for i, r := range resultados {
			fmt.Printf("S%d = %s\n", i+1, nomeCurto(r.pacote))
		}
	}
}

func raizDoModulo() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		pai := filepath.Dir(dir)
		if pai == dir {
			return "", fmt.Errorf("go.mod não encontrado; rode a partir de gilded-rose-go/")
		}
		dir = pai
	}
}

// expandirPacotes resolve padrões como ./prompt_tests/... em pacotes individuais,
// para que cada suíte seja medida separadamente.
func expandirPacotes(raiz string, padroes []string) ([]string, error) {
	args := append([]string{"list"}, padroes...)
	cmd := exec.Command("go", args...)
	cmd.Dir = raiz
	saida, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("go list %v: %w", padroes, err)
	}
	return strings.Fields(string(saida)), nil
}

func nomeCurto(pacote string) string {
	return strings.TrimPrefix(pacote, "github.com/emilybache/gildedrose-refactoring-kata/")
}

// copiarModulo copia os arquivos .go e o go.mod do módulo (sem _evidence e
// outros arquivos grandes, que não participam da compilação).
func copiarModulo(origem, destino string) error {
	return filepath.WalkDir(origem, func(caminho string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(origem, caminho)
		if d.IsDir() {
			if rel == "_evidence" || strings.HasPrefix(d.Name(), ".") && rel != "." {
				return filepath.SkipDir
			}
			return os.MkdirAll(filepath.Join(destino, rel), 0o755)
		}
		if !strings.HasSuffix(caminho, ".go") && d.Name() != "go.mod" && d.Name() != "go.sum" {
			return nil
		}
		dados, err := os.ReadFile(caminho)
		if err != nil {
			return err
		}
		return os.WriteFile(filepath.Join(destino, rel), dados, 0o644)
	})
}

func lerMarcadores(caminho string) (map[string]posicao, error) {
	f, err := os.Open(caminho)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	marcadores := map[string]posicao{}
	scanner := bufio.NewScanner(f)
	for linha := 1; scanner.Scan(); linha++ {
		if m := reMarcador.FindStringSubmatchIndex(scanner.Text()); m != nil {
			ramo := scanner.Text()[m[2]:m[3]]
			// coluna 1-based, como no coverprofile
			marcadores[ramo] = posicao{linha, m[0] + 1}
		}
	}
	return marcadores, scanner.Err()
}

func lerPerfil(caminho string) ([]bloco, error) {
	f, err := os.Open(caminho)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var blocos []bloco
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		m := reBloco.FindStringSubmatch(scanner.Text())
		if m == nil || !strings.HasSuffix(m[1], "/"+arquivoProducao) {
			continue
		}
		n := make([]int, 6)
		for i := range n {
			n[i], _ = strconv.Atoi(m[i+2])
		}
		blocos = append(blocos, bloco{posicao{n[0], n[1]}, posicao{n[2], n[3]}, n[5]})
	}
	return blocos, scanner.Err()
}

func antes(a, b posicao) bool {
	return a.linha < b.linha || (a.linha == b.linha && a.coluna < b.coluna)
}

// ordenarRamos devolve D1T, D1F, D2T, D2F, ..., D17T, D17F.
func ordenarRamos(marcadores map[string]posicao) []string {
	var ramos []string
	for ramo := range marcadores {
		ramos = append(ramos, ramo)
	}
	sort.Slice(ramos, func(i, j int) bool {
		ni, _ := strconv.Atoi(ramos[i][1 : len(ramos[i])-1])
		nj, _ := strconv.Atoi(ramos[j][1 : len(ramos[j])-1])
		if ni != nj {
			return ni < nj
		}
		return ramos[i][len(ramos[i])-1] == 'T'
	})
	return ramos
}

func checar(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
