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
	Contenido string
}
