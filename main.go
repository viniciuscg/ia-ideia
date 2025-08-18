package main

import "github.com/viniciuscg/ia-ideia.git/aplicationCases/search"

func main() {
	search := search.NewSearch("vinic7us, Quero pesquisar todas as musicas escutadas nos ultimos 7 dias")

	json := search.SearchPrompt()
	if json.Type == "" {
		println("Failed to generate a valid search query.")
		return
	}
}
