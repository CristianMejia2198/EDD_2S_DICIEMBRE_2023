import React, { useState, useEffect } from "react";

function TablaAlumnos() {
  const [alumnosRegistrados, SetAlumnosRegistrados] = useState([]);
  useEffect(() => {
    async function peticion() {
      const response = await fetch("http://localhost:4000/tabla-alumnos");
      const result = await response.json();
      SetAlumnosRegistrados(result.Arreglo);
    }
    peticion();
  }, []);

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <h1 className="h3 mb-3 fw-normal">Administrador</h1>
          <br />
          <h4 className="h3 mb-3 fw-normal">Cargar Archivos</h4>
          <br />
          <br />
          <table className="table table-dark table-striped">
            <thead>
              <tr>
                <th scope="col">#</th>
                <th scope="col">Posicion</th>
                <th scope="col">Carnet </th>
                <th scope="col">Password </th>
              </tr>
            </thead>
            <tbody>
              {alumnosRegistrados.map((element, j) => {
                if (element.Id_Cliente != "") {
                  return (
                    <>
                      <tr key={"alum" + j}>
                        <th scope="row">{j + 1}</th>
                        <td>{element.Llave}</td>
                        <td>{element.Persona.Carnet}</td>
                        <td>{element.Persona.Password}</td>
                      </tr>
                    </>
                  );
                }
              })}
            </tbody>
          </table>
          <br />
          <p className="mt-5 mb-3 text-muted">EDD 201700918</p>
          <br />
        </form>
      </div>
    </div>
  );
}

export default TablaAlumnos;
