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
    
