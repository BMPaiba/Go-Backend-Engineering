package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/BMPaiba/Go-Backend-Engineering/internal/store"
)

var usernames = []string{
	"Santiago", "Lucía", "Mateo", "Sofía", "Sebastián",
	"Martina", "Alejandro", "Valentina", "Diego", "Elena",
	"Nicolás", "Isabella", "Samuel", "Camila", "Benjamin",
	"Victoria", "Daniel", "Mariana", "Joaquín", "Gabriela",
	"Andrés", "Natalia", "Felipe", "Luciana", "Gabriel",
	"Ximena", "Leonardo", "Catalina", "Hugo", "Estefanía",
	"Ignacio", "Emilia", "Javier", "Paola", "Ricardo",
	"Fernanda", "Cristóbal", "Valeria", "Marcos", "Raquel",
	"Eduardo", "Andrea", "Tomás", "Silvia", "Adrián",
	"Mónica", "Bautista", "Clara", "Julián", "Vanessa",
}

var titles = []string{
	"El Futuro de Go", "Guía de Docker", "Tips de Postgres", "Clean Code 101", "Microservicios Hoy",
	"Punteros en Go", "Aprendiendo SQL", "Rutas con Chi", "JSON en Backend", "Testing en Go",
	"CI/CD Básico", "Interfaces en Go", "Variables de Entorno", "Context y Timeout", "Manejo de Errores",
	"Sólid en 5 Minutos", "Optimización de DB", "Redis como Cache", "Seguridad en APIs", "JWT vs Cookies",
	"Despliegue en AWS", "Linux para Devs", "Docker Compose Pro", "Logs Eficientes", "Middleware en Go",
	"Websockets Simples", "Goroutines Avanzado", "Canales y Mutex", "Git Flow Ideal", "VPC y Redes",
	"Kubernetes Local", "Scripts de Bash", "Makefiles para Go", "Refactorización", "Patrones de Diseño",
	"Arquitectura Hexagonal", "Domain Driven Design", "GORM vs SQLX", "Migraciones de DB", "Serverless en Go",
	"Grafana y Métricas", "Prometheus Setup", "Mensajería con MQ", "OAuth2 Explicado", "GraphQL en 2026",
	"Postgres Indexing", "Estrategias de Cache", "Unit Testing Tips", "Mocking en Go", "API Documentation",
}

var tags = []string{
	"go", "golang", "postgres", "docker", "kubernetes",
	"aws", "gcp", "azure", "devops", "backend",
	"frontend", "fullstack", "sql", "nosql", "redis",
	"api", "rest", "graphql", "grpc", "microservices",
	"security", "auth", "jwt", "oauth", "testing",
	"tdd", "ci-cd", "linux", "bash", "git",
	"architecture", "clean-code", "solid", "ddd", "performance",
	"scalability", "monitoring", "logging", "tracing", "prometheus",
	"docker-compose", "terraform", "serverless", "distributed-systems", "concurrency",
	"algorithms", "data-structures", "opensource", "web-development", "cloud-native",
}

var contents = []string{
	"Exploramos las nuevas características de Go y cómo mejoran el rendimiento de las APIs modernas.",
	"Configurar Docker desde cero no tiene por qué ser difícil. Aquí te mostramos los comandos esenciales.",
	"Descubre cómo los índices en Postgres pueden salvar tu aplicación de la lentitud extrema.",
	"Escribir código limpio es un arte. Repasamos los principios de Uncle Bob aplicados a Go.",
	"¿Valen la pena los microservicios para proyectos pequeños? Analizamos los pros y contras.",
	"Entender los punteros es el primer paso para dominar la gestión de memoria en aplicaciones críticas.",
	"Desde SELECT hasta JOIN complejos, esta es la guía definitiva para backend developers.",
	"Chi es un router ligero y potente. Aprende a organizar tus rutas de forma profesional.",
	"El manejo de JSON en Go puede ser truculento. Te enseñamos a usar Marshall y Unmarshal.",
	"Si no testeas tu código, no confíes en él. Introducción a go test y tablas de pruebas.",
	"Automatizar el despliegue es vital. Configuramos un pipeline básico con GitHub Actions.",
	"Las interfaces son la base del desacoplamiento en Go. Aprende a usarlas correctamente.",
	"Nunca dejes contraseñas en el código. Usa variables de entorno para mantener tu app segura.",
	"Evita que tus peticiones queden colgadas. Aprende a usar context.Context de manera eficiente.",
	"En Go, los errores son valores. Deja de usar try-catch y abraza el flujo de control nativo.",
	"Repasamos los 5 principios SOLID con ejemplos prácticos en lenguaje Go.",
	"A veces una consulta lenta es solo falta de una buena estrategia de Vacuum en Postgres.",
	"Acelera tu API usando Redis para guardar resultados de consultas frecuentes.",
	"Implementar TLS y validar inputs son los pilares de una API robusta frente a ataques.",
	"¿Cuándo usar JWT y cuándo cookies? Comparamos ambos métodos de autenticación.",
	"Desplegar en la nube puede ser costoso si no eliges bien la instancia de EC2 adecuada.",
	"La terminal es tu mejor amiga. Comandos de Linux que todo programador debe conocer.",
	"Orquestar múltiples contenedores es sencillo si dominas el archivo docker-compose.yml.",
	"Los logs estructurados en formato JSON facilitan mucho la búsqueda de bugs en producción.",
	"Crea middleware para logging, recuperación de pánicos y autenticación en tus servicios.",
	"La comunicación en tiempo real es posible gracias a los protocolos de Websockets.",
	"Las goroutines permiten concurrencia masiva con muy pocos recursos del sistema.",
	"Usa canales para comunicar goroutines de forma segura sin compartir memoria directamente.",
	"Dominar Git es más que solo hacer commit y push. Repasamos el flujo de trabajo en equipo.",
	"Segmentar tu red con VPCs es el primer paso para una infraestructura segura en la nube.",
	"Instalar un cluster de K8s en tu máquina local para desarrollo usando Kind o Minikube.",
	"Automatiza tareas repetitivas en tu servidor usando scripts de Bash bien estructurados.",
	"Un buen Makefile te ahorra escribir comandos largos de Docker y migraciones cada día.",
	"Refactorizar no es cambiar por cambiar, es mejorar la estructura sin alterar el comportamiento.",
	"Patrones como Singleton, Factory y Decorator aplicados de forma idiomática en Go.",
	"Separa tu lógica de negocio de la base de datos usando arquitectura hexagonal.",
	"Modelar el dominio es lo más importante en aplicaciones empresariales complejas.",
	"Comparamos el uso de ORMs contra escribir SQL puro con la librería sqlx.",
	"Las migraciones son el historial de versiones de tu base de datos. No las ignores.",
	"Cómo ejecutar funciones pequeñas en la nube sin preocuparte por el servidor.",
	"Visualiza el estado de tu aplicación con paneles de Grafana conectados a tu DB.",
	"Mide la latencia y el uso de CPU de tus servicios usando métricas de Prometheus.",
	"Usa colas de mensajes para procesar tareas pesadas en segundo plano de forma asíncrona.",
	"Entender el flujo de redirecciones y tokens en el protocolo de autorización OAuth2.",
	"GraphQL permite al frontend pedir exactamente los datos que necesita y nada más.",
	"Los índices B-Tree y GIN en Postgres pueden acelerar tus búsquedas de forma increíble.",
	"Elige entre caché de escritura directa o escritura posterior según tus necesidades.",
	"Los tests unitarios deben ser rápidos, aislados y repetibles para ser efectivos.",
	"Aprende a simular dependencias externas usando interfaces y librerías de mocking.",
	"Generar documentación automática con Swagger (OpenAPI) directamente desde tu código.",
}

var comments = []string{
	"¡Excelente artículo! Me ayudó mucho a entender el concepto.",
	"¿Podrías profundizar más en la parte de la configuración?",
	"Justo lo que estaba buscando para mi proyecto final.",
	"Tengo un error al intentar ejecutar el código, ¿alguien más?",
	"Muy bien explicado, directo al grano.",
	"¿Este enfoque funciona igual en entornos de producción?",
	"Gracias por compartir, no conocía esa librería.",
	"¿Hay alguna alternativa más ligera a esta solución?",
	"Me encanta cómo explicas conceptos complejos de forma simple.",
	"¡Buenísimo! Esperando la segunda parte.",
	"¿Cómo afectaría esto al rendimiento si escala mucho?",
	"He probado esto en local y funciona de maravilla.",
	"Interesante punto de vista, aunque no estoy del todo de acuerdo.",
	"¿Podrías subir el código completo a GitHub?",
	"Excelente guía, muy completa y detallada.",
	"¿Qué versión de Go estás utilizando para este ejemplo?",
	"Me salvaste la vida, llevaba horas bloqueado con esto.",
	"¿Es seguro usar esto junto con un WAF?",
	"La sección sobre concurrencia me voló la cabeza.",
	"Mejor que la documentación oficial, sinceramente.",
	"¿Planeas hacer un video sobre este tema?",
	"Buen post, pero creo que falta mencionar el manejo de señales.",
	"¿Alguna recomendación de lectura adicional sobre esto?",
	"¡Seguí así! Tus posts son de mucha calidad.",
	"¿Esto es compatible con versiones antiguas de Postgres?",
	"No me queda claro el uso de las interfaces aquí.",
	"¡Increíble! Lo voy a implementar mañana mismo.",
	"¿Cómo manejas las migraciones en un equipo grande?",
	"Un poco complejo para principiantes, pero muy útil.",
	"El diseño del blog también está genial, por cierto.",
	"¿Has probado usar NATS en lugar de Redis para esto?",
	"Clarísimo, muchas gracias por el aporte.",
	"¿Qué tal es la curva de aprendizaje de esta herramienta?",
	"Me gustaría ver una comparativa con Node.js.",
	"¡Top! Uno de los mejores recursos que he encontrado este mes.",
	"¿Cómo se integra esto con un pipeline de CI/CD?",
	"Tuve que leerlo dos veces, pero al final lo entendí.",
	"¿Hay algún benchmark disponible para esta implementación?",
	"Muy didáctico, me sirvió para una entrevista técnica.",
	"¿Se puede aplicar este patrón en una arquitectura serverless?",
	"¿Recomiendas alguna certificación para profundizar en esto?",
	"El ejemplo de los punteros es el más claro que he visto.",
	"¿Funciona bien en entornos con alta latencia?",
	"¡Gracias! Me suscribo al feed para no perderme nada.",
	"¿Qué opinas de usar esto con Docker Swarm?",
	"La explicación de los middlewares es clave.",
	"Me dio un par de ideas para mejorar mi arquitectura actual.",
	"¿Es necesario usar una base de datos relacional para esto?",
	"¡Hacía falta un post así en español!",
	"Excelente trabajo, se nota la experiencia tras las palabras.",
}

func Seed(store store.Storage) error {

	ctx := context.Background()

	users := generateUsers(50)

	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user", user)
		}
	}

	posts := generatePosts(50, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post", post)
		}
	}

	comments := generateComments(50, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment", comment)
		}
	}

	log.Println("Seeding complete")

	return nil
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {

	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}
func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {

	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  posts[rand.Intn(len(posts))].UserID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms

}
