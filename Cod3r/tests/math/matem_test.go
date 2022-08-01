package matem

import (
	"testing"
)

func TestMedia(t *testing.T) {
	valorEsperado := 5.0
	media := Media(5.0, 5.0, 5.0, 5.0)

	if media != valorEsperado {
		t.Errorf("valor esperado: %f \n valor recebido: %f", valorEsperado, media)
	}
	t.Logf("Test passed, value spected: %v\n value received: %v", valorEsperado, media)
}

// go test --coverprofile=resultado.out
// go tool cover --html=resultado.out
