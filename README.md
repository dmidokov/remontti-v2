# remontti-v2
## Remontti CRM system 


## Установка и запуск

### Для Linux

1. Клонируйте репозиторий, для этого в терминате в каталоге где будет размещаться ваш проект выполните команду:

    ```bash
    git clone https://github.com/dmidokov/remontti-v2.git
    ```

2. Перейти в каталог `remontti-v2`. Для этого выполните команду:

    ```bash
    cd remontti-v2
    ```

3. Заполнить поля файла быстрого запуска `runapp_local.sh`, предварительно создав его путем копирования файла `runapp.sh`, для этого выполите команду:

    ```bash
    cp ./runapp.sh ./runapp_local.sh
    ```

    установите разрешение на исполнение для данного файла:

    ```bash
    chmod +x ./runapp_local.sh
    ```

    установите значения переменных окружения в файле `runapp_local.sh`, вместо значений в треугольных скобках установите значения для своего подключения:

    ```bash
    export DB_USER_NAME=<database_user_name>
    export DB_USER_PASSWORD=<database_password>
    export DB_HOST=<database_host>
    export DB_PORT=<database_port>
    export DB_NAME=<database_name>
    export ROOT_PATH=<path_to_root_of_application>
    export ADMIN_PASSWORD=<crm_admin_password>
    export SESSION_SECRET=<session_secrec_key>
    ```

    > `ADMIN_PASSWORD` - пароль для пользователя `admin`, <br> `SESSION_SECRET` - секрет которым будут шифроваться данные пользовательской сессии  

4. Для запуска приложения в консоли выполнить команду:

    ```bash
    ./runapp_local.sh
    ```

    > Во время запуска приложения, также будут созданы, если не существуют, таблицы требующиеся для работы и пользователь `admin` с паролем указанным в конфигурационном файле

5. В терминале вы должны увидеть лог со статусом запуска приложения.  

    ```bash
    2022/08/07 16:18:49 Starting the service...
    2022/08/07 16:18:49 Trying to load configuration
    2022/08/07 16:18:49 Trying to connect to database
    2022/08/07 16:18:49 Registrate handlers
    2022/08/07 16:18:49 The service is ready to listen and serve
    ```
    > **Примечание**: при возникновении ошибок связанных с запуском вы увидите в терминале сообщения об ошибках.


### Для Windows

1. Клонируйте репозиторий, для этого в терминате в каталоге где будет размещаться ваш проект выполните команду:

    ```cmd
    git clone https://github.com/RemonttiCRM/remontti-v2.git
    ```

2. Перейти в каталог `remontti-v2`. Для этого выполните команду:

    ```cmd
    chdir remontti-v2
    ```

3. Заполнить поля файла быстрого запуска `runapp_local.ps1`, предварительно создав его путем копирования файла `runapp.ps1`, для этого выполите команду:

    ```cmd
    copy runapp.ps1 runapp_local.ps1
    ```

    установите значения переменных окружения в файле `runapp_local.sh`, вместо значений в треугольных скобках установите значения для своего подключения:

    ```powershell
    $Env:DB_USER_NAME=<database_user_name>
    $Env:DB_USER_PASSWORD=<database_password> 
    $Env:DB_HOST=<database_host>
    $Env:DB_PORT=<database_port>
    $Env:DB_NAME=<database_name> 
    $Env:ROOT_PATH=<path_to_root_of_application> 
    $Env:ADMIN_PASSWORD=<crm_admin_password>
    $Env:SESSION_SECRET=<session_secrec_key>
    go run .\main.go
    ```

    > `ADMIN_PASSWORD` - пароль для пользователя `admin`, <br> `SESSION_SECRET` - секрет которым будут шифроваться данные пользовательской сессии  

4. Для запуска приложения в консоли выполнить команду:

    ```cmd
    ./runapp_local.ps1
    ```

    > Во время запуска приложения, также будут созданы, если не существуют, таблицы требующиеся для работы и пользователь `admin` с паролем указанным в конфигурационном файле

5. В терминале вы должны увидеть лог со статусом запуска приложения.  

    ```cmd
    2022/08/17 20:39:56 Starting the service...
    2022/08/17 20:39:56 Trying to load configuration
    2022/08/17 20:39:56 Trying to connect to database
    2022/08/17 20:39:57 Prepare sessions storage
    2022/08/17 20:39:57 Registrate handlers
    2022/08/17 20:39:57 The service is ready to listen and serve
    ```
    > **Примечание**: при возникновении ошибок связанных с запуском вы увидите в терминале сообщения об ошибках.

    
## Установка дополнительно ПО 

Для запуска приложения требуется иметь на ПК следующие программы

* :fire: - обязательно
* :droplet: - не обязательно
-------------------------
- :fire: golang - компилятор и утилиты для сборки приложений на golang
- :fire: postgres - СУБД, также помимо установки требуется иметь пользователя БД с установленым паролем и базу данных
- :fire: git - утилита командной строки для работы с git репозиториями 
- :droplet: PGAdmin - GUI для управления базами Postgres
- :droplet: Postman - средство разработки API для веб-приложений, позволяет выполнять запросы к API приложений, анализировать, тестировать результаты.

## Создание пользователя и БД postgres

Для того чтобы заполнить файл `runapp_local` вам потребуется сперва создать базу данных и пользователя для нее в `postgres`. Для этого, после установки БД подключитесь к ней используя стандартного пользователя `postgres` и утилиту `psql`.

Для добавления нового пользователя в postgres выполните запрос:

```sql
create user <username> with password '<password>';
```

 Обратите внимание что запрос заканчивается `;`, она является обязательной, так как если вы пропустите этот символ `postgres` будет ожидать от вас дальнейшего ввода, до появления `;`.

`username` и `password` имя и пароль для вашего нового пользователя, правильный запрос может выглядеть следующим образом:
```sql 
create user 'someusername' with password 'someuserpassword';
```

Следующим шагом выполните запрос на создание базы данных для вашего нового пользователя:

```sql
create database <databasename> with owner=<username>;
```

`databasename` это имя вашей новой базы данных, а `username` имя пользователя созданного на предыдущем шаге.
Запрос может выглядеть следующим образом:

```sql
create database remonttidb with owner='someusername';
```

Теперь на вашем сервере `postgres` есть новая база данных и пользователь обладающий правами на нее.