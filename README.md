# urlShortener
**urlShortner** - сервис, который предоставляет API по созданию сокращённых ссылок следующего формата:
- Ссылка должна быть длинной 10 символов
- Ссылка должна состоять из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа '_'
# Сервис принимает следующие запросы по http:
1. Метод `POST`, который сохраняет оригинальный URL в базе и возврает сокращённый
----
* **URL**: /encode/
*  **URL Params**: None 
* **Data Params**
   **Required:**
   ```json
  {"longUrl":"ссылка"}
  ```
* **Success Response:**
  * **Code:** 200 <br />
    **Content:** `{"укороченная ссылка"}`
    
  OR
  
  * **Code:** 200 <br />
    **Content:** `{"longURL is already in base укороченная ссылка"}`
    
* **Error Response:**
  * **Code:** 400 BAD REQUEST <br />
    **Content:** `{"Bad JSON"}`
    
  OR
  
   * **Code:** 400 BAD REQUEST <br />
    **Content:** `{"Missing field 'longUrl' from JSON object"}`
    
  OR
  
   * **Code:** 400 BAD REQUEST <br />
    **Content:** `{"Extraneous data after JSON object"}`
    
  OR
  
   * **Code:** 400 BAD REQUEST <br />
    **Content:** `{"Invalid URI for request"}`
2. Метод Get, который принимает сокращённый URL и выполняет редирект на оригинальный URL
----
* **URL**: /укороченная_ссылка
*  **URL Params**: None
* **Data Params**
   **Required:** None
   
* **Success Response:**
  * **Code:** 303 SEE OTHER <br />
    **Content:** `{"<a href="исходная ссылка">See Other</a>."}`
    
* **Error Response:**
  * **Code:** 404 NOT FOUND <br />
    **Content:** `{"short URL doesn't exist in base"}`
# Хранилище  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;В качестве хранилища возможно использование in-memory решения и базы данных - postgresql. Какое хранилище использовать указывается параметром при запуске сервиса. 
![image](https://user-images.githubusercontent.com/80648065/155390687-8f427f70-a635-4e98-98f9-ee1aca628551.png)
# Usage
По умолчанию поднимается контейнер в котором работает сервис

    make

Поднимается контерйнер с сервисом и контейнер с базой данных postgresql

    make psql
    
Выполняются тесты для inmemory

    make test_inmem
Выполняются тесты для psql

    make test_psql
Выполняются все тесты

    make testall


