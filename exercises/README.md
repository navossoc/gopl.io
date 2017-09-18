# Solução dos exercícios

**Exercício 1.1:** Modifique o programa `echo` para exibir também `os.Args[0]`, que é o nome do comando que o chamou.

**Exercício 1.2:** Modifique o programa `echo` para exibir o índice e o valor de cada um de seus argumentos, um por linha.

**Exercício 1.3:** Experimente medir a diferença de tempo de execução entre nossas versões potencialmente ineficientes e a versão que usa `strings.Join`. (A seção 1.6 mostra parte do pacote `time`, e a seção 11.4 mostra como escrever testes comparativos para uma avaliação sistemática de desempenho.)

**Exercício 1.4:** Modifique `dup2` para que exiba os nomes de todos os arquivos em que cada linha duplicada ocorre.

**Exercício 1.5:** Altere a paleta de cores do programa Lissajous para verde sobre preto, para maior autenticidade. Para criar a cor web `#RRGGBB`, use `color.RGBA{0xRR, 0xGG, 0xBB, 0xff}`, em que cada par de dígitos hexadecimais representa a intensidade do componente vermelho, verde ou azul do pixel.

**Exercício 1.6:** Modifique o programa Lissajous para gerar imagens em várias cores adicionando mais valores a `palette` para então exibi-las alterando o terceiro argumento de `SetColorIndex` de alguma maneira interessante.

**Exercício 1.7:** A chamada de função `io.Copy(dst, src)` lê de `src` e escreve em `dst`. Use-a no lugar de `ioutil.ReadAll` para copiar o corpo da resposta para `os.Stdout` sem exigir um buffer grande o suficiente para armazenar todo o stream. Não se esqueça de verificar o resultado de erro de `io.Copy`.

**Exercício 1.8:** Modifique `fetch` para que o prefixo `http://` seja acrescentado a cada URL de argumento, caso esteja faltando. Você pode usar `strings.HasPrefix`.
