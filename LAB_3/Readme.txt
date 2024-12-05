_________________________________________________________________________________________________________________

INTEGRANTES:

- Javiera Bobadilla, ROL = 202173584-4
- Francisco Rebolledo, ROL = 202004556-9

_________________________________________________________________________________________________________________

CONSIDERACIONES:
* 

_________________________________________________________________________________________________________________

DATOS IMPORTANTES:

0. ENTIDADES POR CADA MAQUINA VIRTUAL:

- VM097: Broker
- VM098: Jayce
- VM099: Supervisores Hexgate
- VM100: Servidores Hextech

0. CONTRASEÑAS DE LAS VM (Por si se pierden):

- VM097: PgFc6y4w23wi
- VM098: Jpo6q5MnzYuj
- VM099: iXFQgxpmuNFf
- VM100: W9ZDdZw6z7Jo

0. ACTUALIZACION DE PROTOS(por posibles cambios)(debes estar en la carpeta de la entidad):

- protoc --go_out=. --go-grpc_out=. <nombre>.proto
* ejemplo: protoc --go_out=. --go-grpc_out=. broker.proto 

_________________________________________________________________________________________________________________

INSTRUCCIONES:

1. ACCESO A DIRECTORIOS CORRESPONDIENTES:

- VM097: cd /home/dist/grupo-25/LAB3
- VM098: cd /home/dist/grupo-25/LAB3
- VM099: cd /home/dist/grupo-25/LAB3
- VM100: cd /home/dist/grupo-25/grupo-25/LAB3
*por alguna razon hay dos carpetas iguales seguidas en la VM100

2. COMANDOS DE INICIALIZACION DE CONTENEDORES POR MV(seguir orden acendente):

- MV097: make buildvm097
         make startvm097

- MV098: make buildvm098
         make startvm098

- MV099: make buildvm099
         make startvm099

- MV100: make buildvm100
         make startvm100

3. COMANDOS PARA VER LOS LOGS CORRESPONDIENTES(SI TIENE):

- MV097: make logsvm097
- MV098: make logsvm098
- MV099: make logsvm099
- MV100: make logsvm100

4. COMANDOS DE CERRADO POR SI NO TERMINAN SU EJECUCION:

- MV097: make stopvm097
- MV098: make stopvm098
- MV099: make stopvm099
- MV100: make stopvm100
         
_________________________________________________________________________________________________________________
