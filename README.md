# URL Shortener

This project is just for learning purpose, It is a basic implementation of a url shortener. Which consists of three basic APIs

### 1. Create ShortURL : It create a short URL for a given long URL. It takes orignal URL as input and returns a short URL. Here is Example:

   <pre>  curl -H 'Content-Type: application/json'       -d '{ "orignal_url":"https://www.youtube.com/"}'       -X POST       http://localhost:8080/shorten </pre> 

### 2. Redirect ShortURL: This API's job is to redirect short URL to Orignal URL we got from above API. You can use it directly this in Browser. 
   
   <pre>http://localhost:8080/2Lf3m5</pre> 

### 3. Count of ShortURL: This API's Job is tell the count of how times a website has been access using a short.
    
   <pre>curl -X GET  http://localhost:8080/count </pre> 

   outputs: 
   <pre>
    {
    "https://www.udemy.com": 3,
    "https://www.youtube.com/": 3
    }</pre>

## Docker Build 

   <pre>docker build -t  short_url_image . </pre>

## Docker Run

   <pre>docker run -p 8080:8080 short_url_image:latest  </pre> 
