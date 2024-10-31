# School Management System API

Este es un sistema de gestión escolar que utiliza Go, Gin, GORM y MySQL. La API permite gestionar estudiantes, materias y calificaciones.

## Tabla de Contenidos

- [Requisitos](#requisitos)
- [Configuración del Proyecto](#configuración-del-proyecto)
- [Ejecución del Proyecto](#ejecución-del-proyecto)
- [Uso de la API](#uso-de-la-api)
- [Estructura de la Base de Datos](#estructura-de-la-base-de-datos)
- [Endpoints](#endpoints)

## Requisitos

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Git](https://git-scm.com/downloads)

## Configuración del Proyecto

1. Clona el repositorio:
    ```bash
    git clone https://github.com/TheAnonymousDarck/final-project-go.git
    cd final-project-go
    ```

2. Modifica el archivo `config.go`, dentro de `/cmd/config` acorde a tus credenciales, en caso de desearlo.

3. Asegúrate de que el archivo `docker-compose.yml` esté correctamente configurado para tu entorno.

## Ejecución del Proyecto

Para iniciar el proyecto, sigue estos pasos:

1. **Construir e iniciar los contenedores**:
    ```bash
    docker-compose up --build
    ```

   Este comando construirá las imágenes de Docker para la aplicación y la base de datos, y luego iniciará los contenedores.

2. La API estará disponible en `http://localhost:8000` y MySQL en el puerto `3306`.

3. **Detener los contenedores**:
    ```bash
    docker-compose down
    ```

## Uso de la API

Utiliza Postman, Curl u otra herramienta para probar los endpoints de la API. Asegúrate de que los contenedores estén en ejecución antes de realizar las solicitudes.

### Endpoints Disponibles

Aquí se listan algunos de los endpoints para gestionar estudiantes, materias y calificaciones:

- **Estudiantes**
    - `POST /students`: Crear un estudiante
    - `GET /students/:id`: Obtener detalles de un estudiante
    - `PUT /students/:id`: Actualizar un estudiante
    - `DELETE /students/:id`: Eliminar un estudiante

- **Materias**
    - `POST /subjects`: Crear una materia
    - `GET /subjects/:id`: Obtener detalles de una materia
    - `PUT /subjects/:id`: Actualizar una materia
    - `DELETE /subjects/:id`: Eliminar una materia

- **Calificaciones**
    - `POST /grades`: Crear una calificación
    - `GET /grades/:student_id`: Obtener calificaciones de un estudiante
    - `PUT /grades/:grade_id`: Actualizar una calificación
    - `DELETE /grades/:grade_id`: Eliminar una calificación

## Estructura de la Base de Datos

Las tablas principales en la base de datos `school_db` son:

- `students`: Gestión de estudiantes, con campos como `student_id`, `name`, `group`, y `email`.
- `subjects`: Gestión de materias, con campos como `subject_id` y `name`.
- `grades`: Gestión de calificaciones, con campos como `grade_id`, `student_id`, `subject_id`, y `grade`.

## Notas Adicionales

- **Migraciones Automáticas**: El sistema migrará automáticamente las tablas en MySQL al iniciar la aplicación.
- **Problemas Comunes**:
    - Si `MySQL` no se inicia correctamente, asegúrate de que las variables de entorno estén configuradas y que no existan conflictos de puertos.
    - Si tu API no se conecta a la base de datos, verifica la configuración en el archivo `config.go` y el bloque `environment` en `docker-compose.yml`.

---

¡Gracias por usar el School Management System API! 
Si tienes alguna pregunta o problema, no dudes en abrir un `issue` en el repositorio.
