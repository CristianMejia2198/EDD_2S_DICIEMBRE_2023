import React from "react";

function PrinicipalAdmin() {
  const archivos = (e) => {
    e.preventDefault();
    window.open("/principal/admin/carga-archivos", "_self");
  };

  const alumnos = (e) => {
    e.preventDefault();
    window.open("/principal/admin/alumnos", "_self");
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
          <center>
            <button className="w-50 btn btn-outline-success" onClick={archivos}>
              Carga Archivos
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={alumnos}>
              Ver Alumnos
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={salir}>
              Reportes
            </button>
          </center>
        </form>
      </div>
    </div>
  );
}

export default PrinicipalAdmin;
