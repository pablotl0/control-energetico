## Descripción del problema: 

En la actualidad, los datos ambientales son cruciales para distintos perfiles como :
- Personas con problemas respiratorios
- Investigadores ambientales
- Responsables de políticas públicas
  
Los objetivos de dichas personas son entender la salud pública, la planificación urbana y el bienestar general. Sin embargo, se presenta una serie de desafíos al intentar obtener y analizar esta información:

### Acceso a Datos Dispersos: 
Para los investigadores ambientales es un incoveniente que los datos esten distribuidos en diferentes fuentes, plataformas o formatos, lo que dificulta su recolección y comprensión. 

### Falta de Contexto Comprensible: 
Para las personas con problemas respiratorios y educadores es un incoveniente que los informes no sean accesibles para el ciudadano común, debido a su lenguaje técnico o falta de explicación contextual.

### Variabilidad de los Datos: 
Para los urbanistas, así como responsables de políticas públicas es un incoveniente que los indicadores ambientales, como la calidad del aire o los niveles de polución, varían a lo largo del día, según las estaciones y dependiendo de las actividades humanas. Por ello necesitan datos en tiempo real para la planificación urbana y para análisis constante.

El proyecto se centrará en organizar y dar accesibilidad a información contrastada ,mejorando la toma de decisiones y el bienestar público.

## Documentación adicional
La documentación adicional se encuentra en el directorio [`/docs`](./docs/ob0/README.md).

[User stories](./docs/User-stories.md)

[Milestones](./docs/Milestones.md)

## Gestor de dependencias

[Go Modules](https://go.dev/ref/mod)

Elijo go mod desde Go 1.11, el lenguaje incluye Go Modules como gestor de dependencias integrado. Es el estándar oficial.

## Gestor de tareas

[Make](https://www.gnu.org/software/make/)
[Justificacion](./docs/Justificacion_gestor_tareas.md)

## Tareas automatizadas
- **make check**: verificar la sintaxis del código desarrollado para garantizar que cumple con los estándares y es funcional.  
- **make test**: ejecutar los tests que se implementarán más adelante para validar el correcto funcionamiento del código y asegurarse de su viabilidad.