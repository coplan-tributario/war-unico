# War Unico

Para resolver a necessidade de se ter um web.xml mapeando todos os objetos com todos os municipios, foi criado este software.

## Como funciona

Para gerar o web.xml é necessário seguir os seguintes passos:

1. Inicialmente passar o web.xml original por alguma formatação (evitar problema de leitura) - Recomendado a utilização do https://www.freeformatter.com/xml-formatter.html
2. Baixar o executavel (pasta dist)
3. Executa-lo passando os seguintes parâmetros:
3.1. --import "caminho até o web.xml de origem"
3.2. --export "caminho até onde o web.xml de destino deve ser gerado"
3.3. --sistema "nome do sistema"

Exemplo de execução

```sh
.\warunico.exe --import C:\war\web.xml --export C:\war\web_new.xml --sistema transparencia
```

## Sistemas que tem definido os municipios

- [ ] Administrativo
- [ ] Aplic
- [ ] Central
- [ ] Contábil
- [ ] Educação
- [ ] Planejamento
- [ ] Protocolo
- [ ] RH
- [x] Tributário
- [x] Transparencia

## Executando o código

Para executar via código fonte é necessário que tenha instalado go em sua maquina (https://go.dev/doc/install)

Assim que estiver instalado é possível executar com:

```sh
cd .\src\
go run .\main.go
```

Mas como é utilizado CLI é recomendado que seja gerado o executavel:

```sh
cd .\src\
go build -o warunico.exe
```