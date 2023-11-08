# CICommandDispatcher

O CICommandDispatcher é um serviço em Go que recebe notificações de repositórios do Bitbucket ou GitHub e executa uma série de comandos pré-configurados com base em regras definidas no arquivo `rules.json`.

## Configuração

### Arquivo `rules.json`

O arquivo `rules.json` contém as regras que determinam quais comandos devem ser executados com base nas notificações recebidas. Certifique-se de configurar corretamente as regras no arquivo `rules.json`. Um exemplo de formato de arquivo `rules.json` é o seguinte:

```json
{
    "registry/app:v(.*)": [
        "sh ./config/templates/sshList.sh"
    ]
}
```

Personalize as regras de acordo com as necessidades do seu projeto.

### Mapeamento da pasta /config/templates no Docker
Este projeto requer a pasta `/config/templates` mapeada no contêiner Docker para que as configurações e templates necessários estejam disponíveis para a execução dos comandos. Certifique-se de incluir essa mapeação ao iniciar o contêiner Docker.
