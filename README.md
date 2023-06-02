# Project Title
Delivery API

## Description

A delivery API mangment system built on top of the http protocol and adheres to the MVC pattern (Although the Views were not implemented).
The API was built with careful abstracted designing and therefore allowing other developers to hook different behaviors easily.

## API Structure
1. `app`
  - The directory where all API behaviors are being implemented
      - `http`
        - Responsible for the API management and execution
          - `controllers`
            - Responsible for all API request and response management
          - `rouets`
            - Responsible for API endpoint declaration
          - `globals`
            - The place where all global variables values reside  
      - `models`
        - Responsible for the DB transaction management and the response returning objects
      - `config`
        - The place where all API configuration resides
      - `database`
        - Responsible for the database management
          - `migrations`
            - Responsible for the DB tables creation
          - `seeds`
            - Responsible for the DB tables fulfillment
          - `connections`
            - Responsible for the connection of all DB types  
      - `externalApis`
        - Responsible for the communication with external apis
          - `apis`
            - The place where all apis behaviors implementations reside
              - `requests`
                - Responsible for the external communication with the apis
                  - `formats`
                    - Responsible for the API communication format (url query, json, etc.)
                  - `prorocols`
                    - Responsible for the communication protocol management
              - `resposnes`
                - Responsible for the external response management
                  - `formats`
                    - Responsible for the response communication format (json, xmk, etc.)e
      - `files`
        - An indexing for the files paths
      - `storage`
        - The place where all system file resides
      - `utils`
        - The place where all app common behavior declaration reside        
               
## Notes
  - Although the assignment specifications were entitled for and slim and oriented API, I've decided to add and abstraction as a good practice, so additional behavior can be hooked easily in the future
  - As I mentioned earlier, the Views mechanism wasn't applied in this project, due to lack of front-end , but it could easily act as a response parameters        handler. For instance: mapping what names of the parameters enter the API. Due to lack of time I haven't implemented this mechanism    
 
      
