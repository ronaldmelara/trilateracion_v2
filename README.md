# Trilateración v.2

Esta versión se basa en el entregable del desafío técnico que me solicitó Mercado Libre. Con esta versión busco mejorar la estructura de proyecto a nivel arquitectura de capas. Las capas que incluyo aca esta basado en la arquitectura que se maneja para proyectos GO en la organización donde trabajo actualmente.

* assets: contiene todos los elementos estáticos como imagenes.
* build: es donde se aloja eel compilado de todo el proyecto
* cmd: contiene el main donde arranca la aplicación
* config: Permite alojar archivos de configuración de la aplicación
* docs: Folder destinado a alojar toda la documentación necesaria del proyecto
* internal: este aloja todos los subfoldes donde se encuentra nuestra lógica de negocio, etc.
  * handles: Estos son los manejadores para las peticiones que vienen desde el router
  * router: contiene los servicios y routers para el EndPoint de la API
  * model: contiene las clases que representan elementos de BD o estructuras de datos en general
  * repository: contiene todas las funciones que permiten conectar a un objecto de BD.
  * utils: contiene funciones utilitarias que son usadas en toda la aplicacion
* test: contiene los test unitarios de la aplicacion
* web: incluye los templates o sitio web que implemente la WEB API
