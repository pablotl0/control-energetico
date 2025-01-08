## Criterios

Elcriterio de elección de la herramienta que voy a trabajar es:
 
- Tiempo de ejecución de la tarea "gofmt -e ./internal/.." en mi proyecto.

## Análisis de criterios

Las herramientas que considero son :
### [Mage](https://github.com/magefile/mage)  
- Tiempo de ejecución: 0.723s.  

### [Task](https://github.com/go-task/task)  
- Tiempo de ejecución: 0.324s.

### [Goyek](https://github.com/goyek/goyek)  
- Tiempo de ejecución: 0.824s. 

### [Make](https://www.gnu.org/software/make/)  
- Tiempo de ejecución: 0.068s, el más rápido entre todas las herramientas.   

### [Just](https://github.com/casey/just)  
- Tiempo de ejecución: 0.091s, similar al de Make. 

### [Gotaskr](https://github.com/Roemer/gotaskr) 
- Tiempo de ejecución: 0.861s.

### [Taskrunner](https://github.com/samsarahq/taskrunner)  
- Tiempo de ejecución: 0.903s.

## Conclusión

Tras considerar el criterio de Tiempo de ejecución, Make es la herramienta que mejor cumple. Por otra parte just tiende un rendimiento casi idéntico y la siguiente más cercana a estas es task. El resto de herramientas las descartaré al tener un tiempo de ejecución mayor.

Make es una herramienta muy consolidada, con décadas de uso y además tiene un gran tiempo de ejecución en mi proyecto por lo que es una opción excelente.
Task y Just son opciones eficientes pero su trayectoria menor las hace más dependientes de actualizaciones y soporte en evolución. Por tanto, elijo Make por la estabilidad que ofrece.
