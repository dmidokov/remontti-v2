# remontti-v2
## Remontti CRM system 


## Установка и запуск
1. Скачать пакет из пепозитория 

    ```
    git clone https://github.com/RemonttiCRM/remontti-v2.git
    ```

2. Перейти в каталог `remontti-v2`. Для этого выполните команду:

    ```
    cd remontti-v2
    ```

3. Заполнить поля файла быстрого запуска `runapp_local.sh`, предварительно создав его из файла `runapp.sh`, для этого выполите команду 

    ```
    cp ./runapp.sh ./runapp_local.sh
    ```

    затем установите разрешение на исполнение для данного файла 

    ```
    chmod +x ./runapp_local.sh
    ```

    далее заполните поля

    ```
    export DB_USER_NAME=<database_user_name>
    export DB_USER_PASSWORD=<database_password>
    export DB_HOST=<datav=base_host>
    export DB_PORT=<database_port>
    export DB_NAME=<database_name>
    export ROOT_PATH=<path_to_root_of_application>
    ```

4. В консоли выполнить команду:

    ```
    ./runapp_local.sh
    ```

5. В терминале вы должны увидеть лог со статусом запуска приложения. 

    ```
    2022/08/07 16:18:49 Starting the service...
    2022/08/07 16:18:49 Trying to load configuration
    2022/08/07 16:18:49 Trying to connect to database
    2022/08/07 16:18:49 Registrate handlers
    2022/08/07 16:18:49 The service is ready to listen and serve
    ```

    
## Установка дополнительно ПО 

Для запуска приложения требуется иметь на ПК следующие программы

* :fire: - обязательно
* :droplet: - не обязательно
-------------------------
- :fire: golang - компилятор и уилиты для сборки приложений на golang
- :fire: postgres - СУБД, также помимо установки требуется иметь пользователя БД с установленым паролем и базу данных
- :droplet: PGAdmin - GIU для управления базами Postgres
- :droplet: Postman - средство разработки API для веб-приложений, позволяет выполнять запросы к API приложений, анализировать, тестировать результаты.

