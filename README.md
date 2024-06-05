# Como usar o código:
Nesse mesmo diretório, crie um arquivo YAML,nomeado por exemplo de `application.yaml`.

Execute o programa passando o caminho para o arquivo YAML:
```bash
go run main.go /path/to/application.yaml
```

O programa irá ler o arquivo `application.yaml`, converter suas configurações para o formato de propriedades do Java e gerar um arquivo `.properties` no mesmo diretório.