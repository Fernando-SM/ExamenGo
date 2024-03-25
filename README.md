<<<<<<< HEAD
# ExamenGo
=======
Uso
Detalles sobre cómo utilizar la API, incluyendo los endpoints disponibles.

Registro de Usuarios
Para registrar un nuevo usuario, envía una solicitud POST a /register con el siguiente cuerpo JSON:

{
  "username": "nuevoUsuario",
  "email": "nuevoUsuario@example.com",
  "phone": "1234567890",
  "password": "C0ntraseña1$&"
}

username: El nombre de usuario deseado para la cuenta.
email: La dirección de correo electrónico del usuario.
phone: El número de teléfono del usuario (debe tener 10 dígitos).
password: La contraseña para la cuenta (debe tener entre 6 y 12 caracteres e incluir al menos una mayúscula, una minúscula, un número y un carácter especial).
Inicio de Sesión
Para iniciar sesión, envía una solicitud POST a /login con el siguiente cuerpo JSON:

{
  "email": "nuevoUsuario@example.com",
  "password": "C0ntraseña1$&"
}

email: La dirección de correo electrónico del usuario registrado.
password: La contraseña del usuario registrado.
La respuesta a una solicitud de inicio de sesión exitosa incluirá un token JWT que puedes usar para autenticarte en solicitudes subsiguientes a otros endpoints que requieran autenticación.

Respuestas:

{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzExNDI3ODM1fQ.jNlaFtwAi_nD8tlFz5naN6lkBoSLaXt6-QCc2sMw858"}
>>>>>>> a7e0b6c (Initial commit, se agrega las funciones necesarias para hacer el login y registro a una bd local)
