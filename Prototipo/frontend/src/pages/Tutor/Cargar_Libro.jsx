import React, { useState } from "react";

function Cargar_Libro() {
  const [contenidoPDF, setContenidoPDF] = useState("");
  const uploadFileTutor = (event) => {
    const file = event.target.files[0];
    console.log(file);
    const reader = new FileReader();

    reader.onload = async (event) => {
      const content = event.target.result;
      console.log(content);
      setContenidoPDF(content);
      const valorLocal = localStorage.getItem("user");
      const response = await fetch("http://localhost:4000/registrar-libro", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Carnet: parseInt(valorLocal),
          Nombre: file.name,
          Contenido: content,
        }),
      });

      const result = await response.json();
    };

    reader.onerror = (error) => {
      console.error("Error al leer el archivo:", error);
    };

    reader.readAsDataURL(file);
  };

  const hacerpubli = (e) => {
    e.preventDefault();
    console.log("Listo");
    window.open("/principal/tutor/publicacion", "_self");
  };

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    localStorage.clear();
    window.open("/", "_self");
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <h1 className="h3 mb-3 fw-normal">Libros Tutor</h1>
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-primary"
              onClick={hacerpubli}
            >
              Publicacion
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-primary" onClick={salir}>
              Cerrar Sesion
            </button>
          </center>
          <br />
          <h4 className="h3 mb-3 fw-normal">Cargar PDF</h4>
          <br />
          <div className="input-group mb-3">
            <label className="input-group-text">Cargar Tutores</label>
            <input
              className="form-control"
              id="inputGroupFile01"
              type="file"
              accept=".pdf"
              onChange={uploadFileTutor}
            />
          </div>
          <div className="mb-3 fw-normal">
            <iframe src={contenidoPDF} width="800" height="800" />
          </div>
        </form>
      </div>
    </div>
  );
}

export default Cargar_Libro;
