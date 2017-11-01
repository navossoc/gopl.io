# Solução dos exercícios

**Exercício 1.1:** Modifique o programa `echo` para exibir também `os.Args[0]`, que é o nome do comando que o chamou.

**Exercício 1.2:** Modifique o programa `echo` para exibir o índice e o valor de cada um de seus argumentos, um por linha.

**Exercício 1.3:** Experimente medir a diferença de tempo de execução entre nossas versões potencialmente ineficientes e a versão que usa `strings.Join`. (A seção 1.6 mostra parte do pacote `time`, e a seção 11.4 mostra como escrever testes comparativos para uma avaliação sistemática de desempenho.)

**Exercício 1.4:** Modifique `dup2` para que exiba os nomes de todos os arquivos em que cada linha duplicada ocorre.

**Exercício 1.5:** Altere a paleta de cores do programa Lissajous para verde sobre preto, para maior autenticidade. Para criar a cor web `#RRGGBB`, use `color.RGBA{0xRR, 0xGG, 0xBB, 0xff}`, em que cada par de dígitos hexadecimais representa a intensidade do componente vermelho, verde ou azul do pixel.

**Exercício 1.6:** Modifique o programa Lissajous para gerar imagens em várias cores adicionando mais valores a `palette` para então exibi-las alterando o terceiro argumento de `SetColorIndex` de alguma maneira interessante.

**Exercício 1.7:** A chamada de função `io.Copy(dst, src)` lê de `src` e escreve em `dst`. Use-a no lugar de `ioutil.ReadAll` para copiar o corpo da resposta para `os.Stdout` sem exigir um buffer grande o suficiente para armazenar todo o stream. Não se esqueça de verificar o resultado de erro de `io.Copy`.

**Exercício 1.8:** Modifique `fetch` para que o prefixo `http://` seja acrescentado a cada URL de argumento, caso esteja faltando. Você pode usar `strings.HasPrefix`.

**Exercício 1.9:** Modifique `fetch` para exibir também o código de status HTTP encontrado em `resp.Status`.

**Exercício 1.10:** Encontre um site que gere uma grande quantidade de dados. Investigue o caching executando `fetchall` duas vezes sucessivamente para ver se o tempo informado sofre muita alteração. Você sempre obtém o mesmo conteúdo? Modifique `fetchall` para exibir sua saída em um arquivo para que ela possa ser examinada.

**Exercício 1.11:** Experimente usar `fetchall` com listas mais longas de argumentos, por exemplo, amostras de sites disponíveis em `alexa.com` que fazem parte do primeiro milhão (top million). Como o programa se comporta se um site simplesmente não responder? (A seção 8.9 descreve maneiras de lidar com esses casos.)

**Exercício 1.12:** Modifique o servidor Lissajous para ler valores de parâmetros do URL. Por exemplo, você pode organizá-lo de modo que um URL como http://localhost:8000/?cycles=20 defina o número de ciclos para 20, em vez de usar o default igual a 5. Utilize a função `strconv.Atoi` para converter o parâmetro do tipo string em um inteiro. Você pode ver a documentação da função usando `go doc strconv.Atoi`.

----------

**Exercício 2.1:** Acrescente tipos, constantes e funções em `tempconv` para processar temperaturas na escala Kelvin, em que zero Kelvin corresponde a -273,15 °C e uma diferença de 1 K tem a mesma magnitude de 1 °C.
