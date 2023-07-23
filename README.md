# massmail
## Send a separate e-mail to a list of known addresses

## Usage
The first run of the program will create a configuration file in the working directory. Fill in every field with the exception of optional file attachments. You'll need either a plaintext or html-encoded file which contains text to be sent as e-mail's body and a mail list file in which there's an e-mail address sitting on its own line.

If all the information is correct - you're all set to launch the binary again. It'll apply information written to the configuration file, look for specified files and start sending e-mails.

## GMail step-by-step
GMail address:password pair does not work out of the box; some toggles in the settings are required.

1. Enter your GMail box
2. Access settings
3. Proceed to "Forwarding and POP/IMAP" tab
4. Enable IMAP
5. Open Google Account/Security
6. Open 2-Step Verification tab
7. Scroll down and select "App passwords"
8. Generate a new app password
9. Use it **instead** of the account password

## License
AGPL

I have no responsibility for **your** actions

-----

## Использование
Первый запуск программы создаст конфигурационный файл в рабочей директории. Заполните каждое поле, с оговоркой в поле вложенных файлов, которое является необязательным. Дополнительно вам потребуется либо обычный текстовый, либо html файлик с текстом сообщения, которое будет отправлено; файл со списком адресов, по одному на каждой строчке.

## Шаги для использования с GMail
Просто Адрес:пароль не пройдёт, нужно чутка потыкать в настройках.

1. Зайдите в почтовый ящик
2. Пройдите в настройки
3. Выберите "Forwarding and POP/IMAP"
4. Включите IMAP
5. Затем откройте настройки самого аккаунта гугла, вкладка "Account/Security"
6. Откройте меню двухфакторной аутентификации
7. Пролистайте вниз и выберите "App passwords"
8. Создайте новый пароль для приложения
9. Используйте его **вместо** пароля аккаунта в конфигурационном файле 

## Лицензия
AGPL

Я не отвечаю за **ваши** действия