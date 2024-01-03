import React, { useState } from "react";

function Carga_Publicacion() {
  const [contenidoPublicacion, setContenidoPublicacion] = useState("");
  const CargarPublicacionTutor = async (e) => {
    e.preventDefault();
    const valorLocal = localStorage.getItem("user");
    const response = await fetch(
      "http://localhost:4000/registrar-publicacion",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Carnet: parseInt(valorLocal),
          Nombre: valorLocal,
          Contenido: contenidoPublicacion,
        }),
      }
    );

    const result = await response.json();
  };

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    window.open("/principal/tutor", "_self");
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <h1 className="h3 mb-3 fw-normal">Publicacion Tutor</h1>
          <br />
          <h4 className="h3 mb-3 fw-normal">Cargar PDF</h4>
          <br />
          <center>
            <button className="w-50 btn btn-outline-primary" onClick={salir}>
              PDF
            </button>
          </center>
          <br />
          <div className="input-group mb-3">
            <div className="col align-self-center">
              <textarea
                name="textarea"
                rows="10"
                cols="50"
                value={contenidoPublicacion}
                onChange={(e) => setContenidoPublicacion(e.target.value)}
              ></textarea>
            </div>
          </div>
          <div className="mb-3 fw-normal">
            <center>
              <button
                className="w-50 btn btn-outline-primary"
                onClick={CargarPublicacionTutor}
              >
                Publicar
              </button>
            </center>
          </div>
        </form>
      </div>
    </div>
  );
}

export default Carga_Publicacion;
