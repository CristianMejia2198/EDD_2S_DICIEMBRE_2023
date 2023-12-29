import React from "react";

function PrinicipalAdmin() {
  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    window.open("/", "_self");
  };
  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <center>
            <button className="w-50 btn btn-outline-success" onClick={salir}>
              Carga Archivos
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={salir}>
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
