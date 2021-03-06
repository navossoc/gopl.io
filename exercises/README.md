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

**Exercício 2.2:** Escreva um programa de conversão de unidades de propósito geral, análogo ao `cf`, que leia números de seus argumentos de linha de comando ou da entrada-padrão se não houver argumentos, e converta cada número em unidades como temperatura em Celsius e em Fahrenheit, comprimento em pés e metros, peso em libras e quilogramas e operações semelhantes.

**Exercício 2.3:** Reescreva `PopCount` para que use um loop no lugar de uma expressão única. Compare o desempenho das duas versões. (A seção 11.4 mostra como comparar o desempenho de diferentes implementações de forma sistemática.)

**Exercício 2.4:** Escreva uma versão de `PopCount` que conte bits deslocando seu argumento pelas 64 posições dos bits, testando o bit mais à direita a cada vez. Compare seu desempenho com a versão que faz consultas na tabela.

**Exercício 2.5:** A expressão `x&(x-1)` limpa o bit diferente de zero mais à direita de x. Escreva uma versão de `PopCount`que conte bits usando esse fato e avalie seu desempenho.

----------

**Exercício 3.1:** Se a função f devolver um valor `float64` não finito, o arquivo SVG conterá elementos `<polygon>` inválidos (embora muitos renderizadores SVG tratem essa situação com elegância). Modifique o programa para ignorar polígonos inválidos que forem gerados.

**Exercício 3.2:** Faça experimentos com visualizações de outras funções do pacote `math`. Você pode gerar padrões como caixa de ovo, morrinhos (moguls) ou uma sela?

**Exercício 3.3:** Pinte cada polígono de acordo com sua altura, de modo que os picos tenham a cor vermelha (#ff0000) e os vales sejam azuis (#0000ff).

**Exercício 3.4:** Seguindo a abordagem do exemplo Lissajous na seção 1.7, crie um servidor web que calcule superfícies e escreva dados SVG ao cliente. O servidor deve definir o cabeçalho `Content-Type` assim:

```go
w.Header().Set("Content-Type", "image/svg+xml")
```

**Exercício 3.5:** Implemente o conjunto de Mandelbrot todo colorido usando a função `image.NewRGBA` e o tipo `color.RGBA` ou `color.YCbCr`.

**Exercício 3.6:** Superamostragem (supersampling) é uma técnica para reduzir o efeito de *pixelation*, calculando o valor da cor em vários pontos em cada pixel e tirando a média. O método mais simples é dividir cada pixel em quatro "subpixels". Implemente isso.

**Exercício 3.7:** Outro fractal simples usa o método de Newton para encontrar soluções complexas a uma função como z⁴-1 = 0. Sombreie cada ponto de partida de acordo com o número de iterações necessárias para se aproximar de uma das quatro raízes. Pinte cada ponto segundo a raiz da qual ele se aproxima.

**Exercício 3.8:** Renderizar fractais com níveis altos de zoom exige alta precisão aritmética. Implemente o mesmo fractal usando quatro representações numéricas diferentes: `complex64`, `complex128`, `big.Float` e `big.Rat`. (Os dois últimos tipos encontram-se no pacote `math/big`. `Float` usa números de ponto flutuante quaisquer, porém com precisão limitada; `Rat` usa números racionais com precisão ilimitada.) Como eles se comparam quanto ao desempenho e ao uso de memória? Em que níveis de zoom os artefatos de renderização tornam-se visíveis?

**Exercício 3.9:** Escreva um servidor web que renderize fractais e escreva os dados da imagem ao cliente. Permita que o cliente especifique os valores de x, y e de zoom como parâmetros da requisição HTTP.

**Exercício 3.10:** Escreva uma versão não recursiva de `comma` usando `bytes.Buffer` no lugar de concatenação de strings.

**Exercício 3.11:** Melhore `comma` de modo que ela trate corretamente números de ponto flutuante e um sinal opcional.

**Exercício 3.12:** Escreva uma função que informe se duas strings são anagramas uma da outra, isto é, se eles contêm as mesmas letras em ordem diferente.

**Exercício 3.13:** Escreva declarações `const` para KB, MB, até YB, da forma mais compacta que você puder.

----------

**Exercício 4.1:** Escreva uma função que conte o número de bits diferentes em dois hashes SHA256. (Veja `PopCount` na seção 2.6.2.)

**Exercício 4.2:** Escreva um programa que exiba o hash SHA256 de sua entrada-padrão por default, mas aceite uma flag de linha de comando para exibir o hash SHA384 ou SHA512 em seu lugar.

**Exercício 4.3:** Reescreva `reverse` usando um ponteiro de array no lugar de uma fatia.

**Exercício 4.4:** Escreva uma versão de `rotate` que funcione com um único passo.

**Exercício 4.5:** Escreva uma função in-place para eliminar duplicatas adjacentes em uma fatia `[]string`.

**Exercício 4.6:** Escreva uma função in-place que transforme toda sequência de espaços Unicode adjacentes (veja `unicode.IsSpace`) de uma fatia `[]byte` codificada em UTF-8 em um único espaço ASCII.

**Exercício 4.7:** Modifique `reverse` para inverter os caracteres de uma fatia `[]byte` que representa uma string codificada em UTF-8, in-place. Você é capaz de fazer isso sem alocar uma nova memória?

**Exercício 4.8:** Modifique `charcount` para contar letras, dígitos e assim por diante de acordo com suas categorias Unicode, usando funções como `unicode.IsLetter`.

**Exercício 4.9:** Escreva um programa `wordfreq` para informar a frequência de cada palavra em um arquivo-texto de entrada. Chame `input.Split(bufio.ScanWords)` antes da primeira chamada a `Scan` para separar a entrada em palavras, e não em linhas.

**Exercício 4.10:** Modifique `issues` para informar os resultados em termos de idade, por exemplo, menos de um mês, menos de um ano e mais de um ano.
