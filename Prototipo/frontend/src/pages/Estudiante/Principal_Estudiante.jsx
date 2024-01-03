import React, { useEffect, useState } from "react";

function Principal_Estudiante() {
  const [cursos, setCursos] = useState([]);

  useEffect(() => {
    async function PedirCursos() {
      const valorLocal = localStorage.getItem("user");
      const response = await fetch("http://localhost:4000/obtener-clases", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Carnet: valorLocal,
        }),
      });
      const result = await response.json();
      console.log(result);
      setCursos(result.Arreglo);
    }
    PedirCursos();
  }, []);

  const Palabra = () => {
    return (
      <div className="row">
        <div className="row align-items-start">
          {cursos.map((item, i) => (
            <div className="form-signin1 col" key={"CursoEstudiante" + i}>
              <div className="text-center">
                <div className="card card-body">
                  <h1 className="text-left" key={"album" + i} value={i}>
                    {item}
                  </h1>
                  <div>
                    <span
                      className="input-group-text"
                      id="validationTooltipUsernamePrepend"
                    ></span>{" "}
                    <br />
                  </div>
                </div>
                <br />
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  };

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    localStorage.clear();
    window.open("/", "_self");
  };

  const publicaciones = (e) => {
    e.preventDefault();
    window.open("/principal/estudiante/publicacion", "_self");
    localStorage.setItem("cursos", JSON.stringify(miArreglo));
  };

  const libro = (e) => {
    e.preventDefault();
    window.open("/principal/estudiante/libro", "_self");
    localStorage.setItem("cursos", JSON.stringify(miArreglo));
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <h1 className="h3 mb-3 fw-normal">Alumno - Clases</h1>
          <br />
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={salir}>
              Salir
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={libro}>
              Ver Libros
            </button>
          </center>
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-success"
              onClick={publicaciones}
            >
              Ver Publicaciones
            </button>
          </center>
          <br />
          {cursos.length > 0 ? <Palabra /> : null}
          <br />
          <p className="mt-5 mb-3 text-muted">EDD 201700918</p>
          <br />
        </form>
      </div>
    </div>
  );
}

export default Principal_Estudiante;
