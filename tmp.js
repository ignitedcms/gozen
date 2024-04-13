var jsonDocument = 
    [
     {
       "type": "table",
       "id": "63c0f27a-716e-804c-6873-cd99b945b63f",
       "x": 80,
       "y": 59,
       "name": "users",
       "entities": [
         {
           "text": "id",
           "id": "49be7d78-4dcf-38ab-3733-b4108701f1"
         },
         {
           "text": "name",
           "id": "49be7d78-4dcf-38ab-3733-b4108701fce4"
         }
       ]
     },
     {
       "type": "table",
       "id": "3253ff2a-a920-09d5-f033-ca759a778e19",
       "x": 255,
       "y": 246,
       "name": "migrations",
       "entities": [
         {
           "text": "id",
           "id": "e97f6f8a-4306-0667-3a95-0a5310a2c15c"
         },
         {
           "text": "firstName",
           "id": "357e132c-aa47-978f-a1fa-d13da6736989"
         },
         {
           "text": "company_fk",
           "id": "8d410fef-5c6e-286d-c9c3-c152d5bd9c52"
         }
       ]
     },
       {
       "type": "connection",
       "join": "inner",
       "id": "81cb3b59-66d1-ffc4-3cb7-0bad52ace43b",
       "source": {
         "table": "63c0f27a-716e-804c-6873-cd99b945b63f",
         "port": "49be7d78-4dcf-38ab-3733-b4108701fce4"
       },
       "target": {
         "table": "3253ff2a-a920-09d5-f033-ca759a778e19",
         "port": "e97f6f8a-4306-0667-3a95-0a5310a2c15c"
       }
     }
   ];
