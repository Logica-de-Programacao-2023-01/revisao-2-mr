package bonus

import "errors"

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	cidadesPartida := make(map[string]bool)

	for _, caminho := range caminhos {
		cidadesPartida[caminho[0]] = true
	}

	for _, caminho := range caminhos {
		if !cidadesPartida[caminho[1]] {
			return caminho[1], nil
		}
	}

	return "", errors.New("not implemented yet")
}
