package main



func oMaisRapido(url1, url2, url3 string) string{
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// estrutura de controle especifica para concorrência

	select {
	case t1 := <- c1:
		return t1

	case r2 := <
	}
}