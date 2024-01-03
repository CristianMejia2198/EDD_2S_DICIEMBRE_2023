package Peticiones

type PeticionLogin struct {
	UserName string
	Password string
	Tutor    bool
}

type PeticionRegistroTutor struct {
	Carnet   int
	Nombre   string
	Curso    string
	Password string
}

type PeticionRegistroAlumno struct {
	Carnet   int
	Nombre   string
	Password string
	Cursos   []string
}

type PeticionCursos struct {
	Cursos []Cursos
}

type Cursos struct {
	Codigo string
	Post   []string
}

type PeticionLibro struct {
	Carnet    int
	Nombre    string
	Contenido string
}

type PeticionPublicacion struct {
	Carnet    int
	Nombre    string
	Contenido string
}

type PeticionDecision struct {
	Accion string
	Nombre string
	Tutor  int
	Curso  string
}

type PeticionAlumnoSesion struct {
	Carnet string
	Cursos []string
}

type RespuestaImagen struct {
	Imagenbase64 string
	Nombre       string
}
