import React, { useEffect, useState } from "react";

function Libros_Control() {
  const [libros, setLibros] = useState([]);
  const [eleccion, setEleccion] = useState(0);

  useEffect(() => {
    async function PedirLibros() {
      const response = await fetch("http://localhost:4000/enviar-libros-admin");
      const result = await response.json();
      console.log(result);
      if (result.status === 200) {
        setLibros(result.Arreglo);
      }
    }
    PedirLibros();
  }, []);

  const salir = (e) => {
    e.preventDefault();
    window.open("/principal/admin", "_self");
  };

  const handleChange = (e) => {
    var j = parseInt(e.target.value);
    setEleccion(j);
    console.log(j);
  };

  const LibrosObtenidos = () => {
    console.log(libros);
    return (
      <div>
        <select
          className="form-control"
          aria-label=".form-select-lg example"
          onChange={handleChange}
        >
          <option value={0}>Elegir Libro....</option>
          {libros.map((item, j) => (
            <option value={j} key={j}>
              {item.Nombre}
            </option>
          ))}
        </select>
      </div>
    );
  };

  const LibrosDefault = () => {
    console.log(libros);
    return (
      <div>
        <select
          className="form-control"
          aria-label=".form-select-lg example"
          onChange={handleChange}
        >
          <option value={0}>Elegir Libro....</option>
        </select>
      </div>
    );
  };

  const aceptar = async (e) => {
    e.preventDefault();
    const valorLocal = localStorage.getItem("user");
    if (libros.length > 0) {
      const response = await fetch("http://localhost:4000/registrar-log", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Accion: "Aceptado",
          Nombre: libros[eleccion].Nombre,
          Tutor: libros[eleccion].Tutor,
          Curso: libros[eleccion].Curso,
        }),
      });

      const result = await response.json();
    }
  };

  const rechazar = async (e) => {
    e.preventDefault();
    const valorLocal = localStorage.getItem("user");
    if (libros.length > 0) {
      const response = await fetch("http://localhost:4000/registrar-log", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Accion: "Rechazado",
          Nombre: libros[eleccion].Nombre,
          Tutor: libros[eleccion].Tutor,
        }),
      });

      const result = await response.json();
    }
  };

  const finalizar = async (e) => {
    e.preventDefault();
    const response = await fetch("http://localhost:4000/finalizar-libros");
    const result = await response.json();
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <h1 className="h3 mb-3 fw-normal">
            Dashboard Empleado {localStorage.getItem("empleado")}
          </h1>
          <br />
          <h4 className="h3 mb-3 fw-normal">Elige un libro</h4>
          <br />
          <div className="col align-self-center">
            {libros.length > 0 ? <LibrosObtenidos /> : <LibrosDefault />}
          </div>
          <br />
          <div className="row align-items-start">
            <div className="col">
              <center>
                <button
                  className="w-50 btn btn-outline-primary"
                  onClick={rechazar}
                >
                  Rechazar
                </button>
              </center>
            </div>
            <div className="col">
              <center>
                <button
                  className="w-50 btn btn-outline-primary"
                  onClick={aceptar}
                >
                  Aceptar
                </button>
              </center>
            </div>
            <div className="col">
              <center>
                <button
                  className="w-50 btn btn-outline-primary"
                  onClick={finalizar}
                >
                  Finalizar
                </button>
              </center>
            </div>
          </div>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={salir}>
              Salir
            </button>
          </center>
          <br />
          <p className="mt-5 mb-3 text-muted">EDD 201700918</p>
          <br />
        </form>
      </div>
    </div>
  );
}

export default Libros_Control;
