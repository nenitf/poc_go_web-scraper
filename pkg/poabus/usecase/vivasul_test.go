package usecase

import (
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
)

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		//		w.Write([]byte(`<!DOCTYPE html>
		//<html>
		//<head>
		//<title>Test Page</title>
		//</head>
		//<body>
		//<h1>Hello World</h1>
		//</body>
		//</html>
		//		`))

		file := path.Join("server_test", "index.htm")
		http.ServeFile(w, r, file)
	})
	return httptest.NewServer(mux)
}

func TestDeveExtrairNomeDaPagina(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	sut := NovaBuscaPeloVivaSul()
	horario := sut.Execute(ts.URL)

	got := horario.Nome
	want := "Linha 289 - RINCAO / VIA OSCAR PEREIRA"

	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func TestDeveExtrairHorariosDaPagina(t *testing.T) {
	// TODO como deixar o código menos longo?

	ts := newTestServer()
	defer ts.Close()

	sut := NovaBuscaPeloVivaSul()
	h := sut.Execute(ts.URL)

	gotl := len(h.HorariosUteisInvernoBC)
	wantl := 25

	if gotl != wantl {
		t.Error("got:", gotl, "want:", wantl)
	}

	got := h.HorariosUteisInvernoBC[0].Horario
	want := "05:53"

	if got != want {
		t.Error("got:", got, "want:", want)
	}

	got = h.HorariosUteisInvernoBC[len(h.HorariosUteisInvernoBC)-1].Horario
	want = "18:00"

	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
