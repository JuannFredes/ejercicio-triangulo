# Triángulos - Ejercicio

---

Este es un proyecto en el cual resuelvo un ejercicio de lógica utilizando el lenguaje Go. El ejercicio consiste en recibir como entrada un archivo txt que contiene un conjunto de puntos cardinales que se deben leer desde el archivo. Una vez que se tienen estos puntos, se deben formar la mayor cantidad de triángulos posibles y determinar cuáles de ellos contienen puntos interiores y cuáles no. Posteriormente, se debe buscar el triángulo de mayor superficie dentro del subconjunto de triángulos resultantes.

## Instalación

---

**Para poder probar el código asegúrate de:**

1. Tener instalado [Go](https://go.dev/) en tu máquina.
2. Clonar el repositorio en tu máquina: `git clone https://github.com/JuannFredes/ejercicio-triangulo.git`.
3. Para probar el código sin compilar, escribe en la terminal: `go run main.go`.
4. Para compilar y probar, en la termina escribe `go build` y para ejecutar escribe `./ejercicio-triangulos`.

## Consigna del ejercicio y resolución 

---

[Consigna del ejercicio en pdf](https://github.com/JuannFredes/ejercicio-triangulo/blob/main/enunciado.pdf)

En mi caso, he dado una respuesta más completa de la que pedía la consigna del ejercicio, ya que he encontrado cuáles son los triángulos de mayor superficie en ambos casos, tanto aquellos que contienen puntos interiores como aquellos que no los contienen.

**Resolución:**

1. Creé dos carpetas: `models` y `util`. En `models` cree las estructuras de datos para los puntos y triángulos, y en `util` cree las funciones necesarias para la lógica del programa.
2. En la estructura de datos para el triángulo agregué un método que calcula la superficie del mismo.
3. En `util`, implementé las siguientes funciones, cada una con sus respectivos tests:
- Una función que lee un archivo txt y, siguiendo el formato que especifica la consigna, devuelve los puntos en una estructura de datos.
- Una función que genera todos los triángulos posibles con los puntos dados.
- Una función que determina, mediante la orientación, cuáles triángulos contienen puntos interiores y cuáles no. Devuelve dos subconjuntos de triángulos: uno que contiene puntos interiores y otro que no los contiene.
- Una función que encuentra el triángulo de mayor superficie en un conjunto de triángulos.
4. En el archivo `main.go` utilicé las funciones definidas en `util` para obtener los dos subconjuntos de triángulos que contienen puntos interiores y aquellos que no los contienen. Luego, busqué el triángulo de mayor superficie en cada uno de estos subconjuntos.

