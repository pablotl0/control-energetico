## Criterios

Los criterios de elección de la herramienta que voy a trabajar son:

- Mantenimiento: Si elegimos una herramienta que no se mantiene más adelante se tendrá que sustituir por otra herramienta, con  lo que ello conlleva (pérdida de tiempo y pérdida de dinero). 
Para ello vamos a fijarnos en los repositorios de GitHub o páginas web de los proyectos para encontrar: 
	- Frecuencia de releases

- Comunidad: Una comunidad  asegura soporte a largo plazo, resolución rápida de problemas y mejoras continuas.
Para medirlo analizaremos:
	- Número de contribuidores
	- Pull requests abiertos y cerrados
	- Issues resueltos vs pendientes

- Seguridad: La seguridad es un criterio crucial ya que herramientas de task running tienen acceso directo al código fuente, pueden ejecutar comandos en el sistema y pueden tener acceso a credenciales y configuraciones sensibles.
Los aspectos a evaluar para la seguridad:
	- Tiempo promedio de respuesta a issues 
	- Vulnerabilidades del código


## Análisis de criterios

De los criterios anteriores hay mayor diferencia en la parte de mantenimiento y comunidad donde el desempeño de los gestores es:

1. Mage

- Mantenimiento:
  - Frecuencia de releases: Aproximadamente cada 6-8 meses.  
- Comunidad :  
  - 4.2k estrellas en GitHub.  
  - 76 contribuidores.  
  - Issues abiertos mesualmente mientras que no se cierran issues muy a menudo al igual ocurre con los pull request.


2. Make

- Mantenimiento: 
  - Frecuencia de releases: Aproximadamente cada 1-2 años.  
- Comunidad :  
  - Gran comunidad  con muchas conversaciones donde aprender sobre makefile.  

3. Task

- Mantenimiento:   
  - Frecuencia de releases alta, con actualizaciones menores constantes cada 1-3 meses.  
- Comunidad :  
  - 11.6k estrellas en GitHub.  
  - 190 contribuidores.  
  - Varios issues abiertos y cerrados semanalmente al igual que con los pull request.  

4. Goyek

- Mantenimiento:
  - Frecuencia de releases: Pocas versiones últimamente, su frecuencia es anual.
- Comunidad :  
  - 300 estrellas en GitHub.  
  - 8 contribuidores.     
  - Varios issues abiertos y cerrados semanalmente al igual que con los pull request. 

5. Just

- Mantenimiento:
  - Frecuencia de releases: Actualizaciones frecuentes, generalmente cada 1-2 meses.
- Comunidad :  
  - 10.5k estrellas en GitHub.
  - 240 contribuidores.     
  - Varios issues abiertos y cerrados semanalmente al igual que con los pull request. 

6. Invoke

- Mantenimiento:
  - Frecuencia de releases: Menor actividad en los últimos tiempos, con lanzamientos cada 1-2 años.
- Comunidad :  
  - 3.8k estrellas en GitHub.  
  - 72 contribuidores.     
  - Issues abiertos mesualmente mientras que no se cierran issues muy a menudo al igual ocurre con los pull request.

En cuanto al criterio de seguridad hay una diferencia menos significativa:

Seguridad:

Un proyecto bien mantenido y respaldado por una comunidad amplia y  proporciona mayor garantía de seguridad, ya que está en constante evolución para adaptarse a los desafíos y estándares actuales.
En este caso los proyectos mas grandes y reconocidos como Task, Make y Just son favorecidos en este aspecto.

En cuanto a vulnerabilidades, no hay grandes diferencias entre las distintas opciones. 
Tras analizar bases de datos que recopilan vulnerabilidades como:
- CVE Details (Common Vulnerabilities and Exposures):https://www.cvedetails.com/.
- National Vulnerability Database (NVD): https://nvd.nist.gov/.
- Exploit Database: https://www.exploit-db.com/.

No se han encontrado vulnerabilidades aparentemente en Task, Mage, Goyek, Invoke o Just. 
Para make se han encontrado alguna vulnerabilidad en versiones antiguas pero no muy relevantes.


## Conclusión

En comparación con otras herramientas, Task y Just destacan por combinar un mantenimiento constante manteniendo al día con las necesidades de los usuarios y las tendencias tecnológicas. Con unas comunidades que aseguran que el conocimiento, soporte y recursos sobre la herramienta están ampliamente disponibles. 
Sin embargo, la comunidad de task está enfocada en Go y las actualizaciones frecuentes aseguran mejor soporte y adaptabilidad. Aunque Just es sólido, Task es más adecuado para automatización avanzada en este ecosistema.
Por ello elijo el gestor de tareas Task.