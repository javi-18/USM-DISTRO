_________________________________________________________________________________________________________________

INTEGRANTES:

- Javiera Bobadilla, ROL = 202173584-4
- Francisco Rebolledo, ROL = 202004556-9

_________________________________________________________________________________________________________________

CONSIDERACIONES:
* Algunas funciones no estan implementadas debido a la complejidad / falta de experiencia / tiempo por lo que
es muy posible que el programa nunca termine y se quede en un loop infinito en practicamente todos los casos.

_________________________________________________________________________________________________________________

DATOS IMPORTANTES:

0. ENTIDADES POR CADA MAQUINA VIRTUAL:

- VM097: DataNode1/ DataNode2
- VM098: PrimaryNode / Diaboromon
- VM099: ContinenteServer/ IslaFile
- VM100: ContinenteFolder/ NodoTai

0. CONTRASEÑAS DE LAS VM (Por si se pierden):

- VM097: PgFc6y4w23wi
- VM098: Jpo6q5MnzYuj
- VM099: iXFQgxpmuNFf
- VM100: W9ZDdZw6z7Jo

0. INICIALIZACION DE LA RED DE SUPERPOSICION DE DOCKER (Por si se cierra):

- En VM097 usar el comando: docker swarm init 
* este comando devuelve el token que debe ser usado en las demas VM.
- En las demas VM usar el comando: docker swarm join --token <token> 10.35.168.107:2377
* ejemplo: docker swarm join --token SWMTKN-1-3njuzaevzn04qaukk1fhhgjvfyofolftvtiug5pn2y54g8hu1d-6ci4tp6ucz21iq44orrnxecn8 10.35.168.107:2377

0. ACTUALIZACION DE PROTOS(por posibles cambios)(debes estar en la carpeta de la entidad):

- protoc --go_out=. --go-grpc_out=. <nombre>.proto
* ejemplo: protoc --go_out=. --go-grpc_out=. isla_file.proto 

_________________________________________________________________________________________________________________

INSTRUCCIONES:

1. ACCESO A DIRECTORIOS CORRESPONDIENTES:

- VM097: cd /home/dist/grupo-25/LAB2/MV097/
- VM098: cd /home/dist/grupo-25/LAB2/MV098/
- VM099: cd /home/dist/grupo-25/LAB2/MV099/
- VM100: cd /home/dist/grupo-25/grupo-25/LAB2/MV100/ 
*por alguna razon hay dos carpetas iguales seguidas en la VM100

2. COMANDOS DE INICIALIZACION DE CONTENEDORES POR MV(seguir orden acendente):

- MV097: make build
         make up
- MV098: make build
         make up
- MV099: make build
         make up
- MV100: make build
         make up

3. COMANDOS DE CERRADO POR SI NO TERMINAN SU EJECUCION:

- MV097: make down
         make clean
- MV098: make down
         make clean
- MV099: make down
         make clean
- MV100: make down
         make clean
         
_________________________________________________________________________________________________________________