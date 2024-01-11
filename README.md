## Descrição

O projeto **upload-go** é uma aplicação desenvolvida em Go (Golang) para facilitar o processo de upload de arquivos. Utilizando Docker para simplificar o ambiente de execução, a aplicação oferece uma solução eficiente e segura para transferência de dados para um servidor.

## Utilização com Docker

Certifique-se de ter o Docker instalado no seu sistema antes de prosseguir.

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/upload-go.git
cd upload-go
```

2. Gere o token de autenticação conforme descrito anteriormente.

3. Crie o arquivo `.env` e insira o token gerado:

```env
AUTH_TOKEN=SEU_TOKEN_AQUI
```

4. Execute a aplicação com Docker:

```bash
docker compose up -d
```

A aplicação estará disponível em `http://sua_url/upload`.

## Endpoints

- **URL de Upload:** `http://sua_url/upload`

  **Retorno:**
  ```json
  {
    "filename": "arquivo.json"
  }
  ```

- **Visualizar Arquivo:** `http://sua_url/file/arquivo.json`

Lembre-se de substituir "sua_url" pela URL real onde a aplicação está hospedada.

Este projeto é mantido por [Glemiso C. Dutra](https://github.com/DTunnel0). Sinta-se à vontade para contribuir, relatar problemas ou enviar sugestões!